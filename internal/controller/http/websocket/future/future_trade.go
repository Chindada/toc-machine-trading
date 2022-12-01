// Package future package future
package future

import (
	"encoding/json"
	"sync"
	"time"

	"tmt/internal/entity"
	"tmt/internal/usecase"
	"tmt/internal/usecase/modules/event"

	"tmt/internal/controller/http/websocket"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WSFutureTrade struct {
	*websocket.WSRouter // ws router
	*event.Bus          // event bus

	s usecase.Stream // stream
	o usecase.Order  // order

	// save tick chan for assist
	assistTickChanMap     map[string]chan *entity.RealTimeFutureTick
	assistTickChanMapLock sync.RWMutex

	// if waiting manual is not nil, will not accept new order
	waitingManualOrder *entity.FutureOrder

	// save manual order or order from assist
	orderMap           map[string]*entity.FutureOrder
	orderMapLock       sync.Mutex
	cancelOrderMap     map[string]*entity.FutureOrder
	cancelOrderMapLock sync.Mutex

	// limit one process at a time
	processLock sync.Mutex

	// save assist target for assist, if assist order is done, will start to send tick to assist target
	// then delete from map
	assistTargetWaitingMap     map[string]*assistTarget
	assistTargetWaitingMapLock sync.Mutex
}

// StartWSFutureTrade - Start ws future trade with one time bus
func StartWSFutureTrade(c *gin.Context, s usecase.Stream, o usecase.Order) {
	w := &WSFutureTrade{
		s:                      s,
		o:                      o,
		assistTickChanMap:      make(map[string]chan *entity.RealTimeFutureTick),
		assistTargetWaitingMap: make(map[string]*assistTarget),
		orderMap:               make(map[string]*entity.FutureOrder),
		cancelOrderMap:         make(map[string]*entity.FutureOrder),
		WSRouter:               websocket.NewWSRouter(c),
		Bus:                    event.Get(true),
	}

	forwardChan := make(chan []byte)
	go func() {
		for {
			msg, ok := <-forwardChan
			if !ok {
				return
			}

			var fMsg clientOrder
			if err := json.Unmarshal(msg, &fMsg); err != nil {
				w.SendErrToClient(errUnmarshal)
				continue
			}
			w.processClientOrder(fMsg)
		}
	}()

	go w.sendFuture()
	go w.checkAssistTargetStatus()

	w.SubscribeTopic(topicAssistDone, w.closeDoneChan)
	w.SubscribeTopic(topicPlaceOrder, w.addOrderFromAssist)

	w.ReadFromClient(forwardChan)
}

func (w *WSFutureTrade) processClientOrder(client clientOrder) {
	defer w.processLock.Unlock()
	w.processLock.Lock()

	switch {
	case !w.o.IsFutureTradeTime():
		w.SendErrToClient(errNotTradeTime)
		return
	case w.waitingManualOrder != nil:
		w.SendErrToClient(errNotFilled)
		return
	case w.isAssistingFull():
		w.SendErrToClient(errAssitingIsFull)
		return
	case client.Option.AutomationType != AutomationNone && client.Qty > 1:
		w.SendErrToClient(errAssistNotSupport)
		return
	}

	o := w.placeOrder(client.toFutureOrder())
	if o == nil {
		w.SendErrToClient(errPlaceOrder)
		return
	}
	w.waitingManualOrder = o

	if client.Option.AutomationType != AutomationNone {
		// save assist target, wait for order status update
		w.assistTargetWaitingMapLock.Lock()
		w.assistTargetWaitingMap[o.OrderID] = &assistTarget{
			WSFutureTrade:        w,
			FutureOrder:          o,
			halfAutomationOption: client.Option,
		}
		w.assistTargetWaitingMapLock.Unlock()
	} else {
		// save manual order, it has timeout
		w.orderMapLock.Lock()
		w.orderMap[o.OrderID] = o
		w.orderMapLock.Unlock()
	}
}

func (w *WSFutureTrade) isAssistingFull() bool {
	defer w.assistTickChanMapLock.RUnlock()
	w.assistTickChanMapLock.RLock()
	return len(w.assistTickChanMap) > 0
}

func (w *WSFutureTrade) closeDoneChan(orderID string) {
	w.assistTickChanMapLock.Lock()
	close(w.assistTickChanMap[orderID])
	delete(w.assistTickChanMap, orderID)
	w.assistTickChanMapLock.Unlock()
}

func (w *WSFutureTrade) addOrderFromAssist(o *entity.FutureOrder) {
	w.orderMapLock.Lock()
	w.orderMap[o.OrderID] = o
	w.orderMapLock.Unlock()
}

func (w *WSFutureTrade) sendFuture() {
	if err := w.sendFutureSnapshot(); err != nil {
		w.SendErrToClient(err)
		return
	}

	tickChan := make(chan *entity.RealTimeFutureTick)
	orderStatusChan := make(chan interface{})
	go w.processTickArr(tickChan)
	go w.processOrderStatus(orderStatusChan)

	go w.sendTradeIndex()
	go w.sendPosition()

	connectionID := uuid.New().String()
	w.s.NewFutureRealTimeConnection(tickChan, connectionID)
	w.s.NewOrderStatusConnection(orderStatusChan, connectionID)

	<-w.Ctx().Done()

	w.s.DeleteFutureRealTimeConnection(connectionID)
	w.s.DeleteOrderStatusConnection(connectionID)
}

func (w *WSFutureTrade) checkAssistTargetStatus() {
	for {
		select {
		case <-w.Ctx().Done():
			return

		case <-time.After(1 * time.Second):
			w.assistTargetWaitingMapLock.Lock()
			for orderID, a := range w.assistTargetWaitingMap {
				if a.Status == entity.StatusFilled {
					assist := newAssistTrader(w.Ctx(), a)
					w.assistTickChanMapLock.Lock()
					w.assistTickChanMap[orderID] = assist.getTickChan()
					w.assistTickChanMapLock.Unlock()
				}

				if !a.Cancellable() {
					delete(w.assistTargetWaitingMap, orderID)
				}
			}
			w.assistTargetWaitingMapLock.Unlock()
		}
	}
}

func (w *WSFutureTrade) placeOrder(order *entity.FutureOrder) *entity.FutureOrder {
	var err error
	switch order.Action {
	case entity.ActionBuy:
		order.OrderID, order.Status, err = w.o.BuyFuture(order)
		if err != nil {
			return nil
		}

	case entity.ActionSell:
		order.OrderID, order.Status, err = w.o.SellFuture(order)
		if err != nil {
			return nil
		}
	}

	// order trade time is not set by order service, set it here
	order.TradeTime = time.Now()
	return order
}

func (w *WSFutureTrade) processTickArr(tickChan chan *entity.RealTimeFutureTick) {
	baseDuration := 10 * time.Second
	var tickArr entity.RealTimeFutureTickArr
	for {
		tick, ok := <-tickChan
		if !ok {
			return
		}
		w.sendTickToAssit(tick)
		w.SendToClient(tick)
		tickArr = append(tickArr, tick)
		var firstPeriod, secondPeriod, thirdPeriod, fourthPeriod entity.RealTimeFutureTickArr
	L:
		for i := len(tickArr) - 1; i >= 0; i-- {
			switch {
			case time.Since(tickArr[i].TickTime) <= baseDuration*1:
				fourthPeriod = append(fourthPeriod, tickArr[i])
				thirdPeriod = append(thirdPeriod, tickArr[i])
				secondPeriod = append(secondPeriod, tickArr[i])
				firstPeriod = append(firstPeriod, tickArr[i])

			case time.Since(tickArr[i].TickTime) <= baseDuration*2:
				fourthPeriod = append(fourthPeriod, tickArr[i])
				thirdPeriod = append(thirdPeriod, tickArr[i])
				secondPeriod = append(secondPeriod, tickArr[i])

			case time.Since(tickArr[i].TickTime) <= baseDuration*3:
				fourthPeriod = append(fourthPeriod, tickArr[i])
				thirdPeriod = append(thirdPeriod, tickArr[i])

			case time.Since(tickArr[i].TickTime) <= baseDuration*4:
				fourthPeriod = append(fourthPeriod, tickArr[i])

			default:
				tickArr = tickArr[i+1:]
				break L
			}
		}

		w.SendToClient(periodTradeVolume{
			FirstPeriod:  firstPeriod.GetOutInVolume(),
			SecondPeriod: secondPeriod.GetOutInVolume(),
			ThirdPeriod:  thirdPeriod.GetOutInVolume(),
			FourthPeriod: fourthPeriod.GetOutInVolume(),
		})
	}
}

func (w *WSFutureTrade) processOrderStatus(orderStatusChan chan interface{}) {
	finishedOrderMap := make(map[string]*entity.FutureOrder)
	for {
		order, ok := <-orderStatusChan
		if !ok {
			return
		}

		if o, ok := order.(*entity.FutureOrder); ok {
			if finishedOrderMap[o.OrderID] != nil {
				continue
			}

			w.updateCacheOrder(o)

			if !o.Cancellable() {
				finishedOrderMap[o.OrderID] = o
				if w.waitingManualOrder != nil && w.waitingManualOrder.OrderID == o.OrderID {
					w.waitingManualOrder = nil
				}
			} else {
				w.cancelOverTimeOrder(o)
			}
		}
	}
}

// cancelOverTimeOrder cancel order if it is not cancelled or filled, and also update order from assist
func (w *WSFutureTrade) updateCacheOrder(o *entity.FutureOrder) {
	w.orderMapLock.Lock()
	cache, ok := w.orderMap[o.OrderID]
	if ok {
		if cache.Status != o.Status {
			w.SendToClient(o)
		}
		o.TradeTime = cache.TradeTime
		w.orderMap[o.OrderID] = o
		w.PublishTopicEvent(topicOrderStatus, o) // publish updated order to assist
	}
	w.orderMapLock.Unlock()
	if !ok {
		w.updateAssistTargetWaitingOrder(o)
	}
}

// updateAssistTargetWaitingOrder if filled, assist trader will start
func (w *WSFutureTrade) updateAssistTargetWaitingOrder(o *entity.FutureOrder) {
	w.assistTargetWaitingMapLock.Lock()
	if a, ok := w.assistTargetWaitingMap[o.OrderID]; ok {
		if a.Status != o.Status {
			w.SendToClient(o)
		}
		o.TradeTime = a.TradeTime
		a.FutureOrder = o
		w.assistTargetWaitingMap[o.OrderID] = a
	}
	w.assistTargetWaitingMapLock.Unlock()
}

func (w *WSFutureTrade) cancelOverTimeOrder(o *entity.FutureOrder) {
	defer w.cancelOrderMapLock.Unlock()
	w.cancelOrderMapLock.Lock()

	if o.Cancellable() && time.Since(o.TradeTime) > 10*time.Second && w.cancelOrderMap[o.OrderID] == nil {
		id, s, err := w.o.CancelFutureOrderID(o.OrderID)
		if err != nil || s != entity.StatusCancelled || id == "" {
			w.SendErrToClient(errCancelOrderFailed)
			return
		}
		w.cancelOrderMap[o.OrderID] = o
	}
}

func (w *WSFutureTrade) sendTickToAssit(tick *entity.RealTimeFutureTick) {
	w.assistTickChanMapLock.RLock()
	for _, v := range w.assistTickChanMap {
		v <- tick
	}
	w.assistTickChanMapLock.RUnlock()
}

func (w *WSFutureTrade) sendFutureSnapshot() error {
	snapshot, err := w.s.GetFutureSnapshotByCode(w.s.GetMainFutureCode())
	if err != nil {
		return errGetSnapshot
	} else {
		w.SendToClient(snapshot.ToRealTimeFutureTick())
	}
	return nil
}

func (w *WSFutureTrade) sendTradeIndex() {
	w.SendToClient(w.generateTradeIndex())

	for {
		select {
		case <-w.Ctx().Done():
			return

		case <-time.After(5 * time.Second):
			w.SendToClient(w.generateTradeIndex())
		}
	}
}

func (w *WSFutureTrade) sendPosition() {
	if position, err := w.generatePosition(); err != nil {
		w.SendErrToClient(errGetPosition)
	} else {
		w.SendToClient(&futurePosition{position})
	}

	for {
		select {
		case <-w.Ctx().Done():
			return

		case <-time.After(5 * time.Second):
			if position, err := w.generatePosition(); err != nil {
				w.SendErrToClient(errGetPosition)
			} else {
				w.SendToClient(&futurePosition{position})
			}
		}
	}
}

func (w *WSFutureTrade) generateTradeIndex() *entity.TradeIndex {
	t := w.s.GetTradeIndex()
	switch {
	case t.Nasdaq == nil, t.NF == nil, t.TSE == nil, t.OTC == nil:
		time.Sleep(time.Second)
		return w.generateTradeIndex()
	default:
		return t
	}
}

func (w *WSFutureTrade) generatePosition() ([]*entity.FuturePosition, error) {
	position, err := w.o.GetFuturePosition()
	if err != nil {
		return nil, err
	}
	return position, nil
}
