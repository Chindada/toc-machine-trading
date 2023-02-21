// Package dt package dt
package dt

import (
	"fmt"
	"sync"
	"time"

	"tmt/cmd/config"
	"tmt/internal/entity"
	"tmt/internal/usecase/grpcapi"
	"tmt/pkg/eventbus"

	"github.com/google/uuid"
)

type DTFuture struct {
	code          string
	orderQuantity int64

	lastTick *entity.RealTimeFutureTick
	tickArr  entity.RealTimeFutureTickArr

	sc       *grpcapi.TradegRPCAPI
	localBus *eventbus.Bus

	traderMap     map[string]*DTTraderFuture
	traderMapLock sync.RWMutex

	tickChan   chan *entity.RealTimeFutureTick
	notify     chan *entity.FutureOrder
	switchChan chan bool

	tradeConfig *config.TradeFuture

	lastTickRate float64
	isTradeTime  bool
}

func NewDTFuture(code string, s *grpcapi.TradegRPCAPI, tradeConfig *config.TradeFuture) *DTFuture {
	d := &DTFuture{
		code:          code,
		orderQuantity: tradeConfig.Quantity,
		tickChan:      make(chan *entity.RealTimeFutureTick),
		notify:        make(chan *entity.FutureOrder),
		switchChan:    make(chan bool),
		sc:            s,
		tickArr:       []*entity.RealTimeFutureTick{},
		localBus:      eventbus.Get(uuid.NewString()),
		tradeConfig:   tradeConfig,
		traderMap:     make(map[string]*DTTraderFuture),
	}

	d.localBus.SubscribeTopic(topicTraderDone, d.removeDoneTrader)

	d.processOrderStatusAndTradeSwitch()
	d.processTick()

	return d
}

func (d *DTFuture) processOrderStatusAndTradeSwitch() {
	notCancellableOrderMap := make(map[string]*entity.FutureOrder)
	go func() {
		for {
			select {
			case o := <-d.notify:
				if _, ok := notCancellableOrderMap[o.OrderID]; ok {
					continue
				}

				if !o.Cancellable() {
					notCancellableOrderMap[o.OrderID] = o
				} else {
					d.cancelOverTimeOrder(o)
				}

				d.localBus.PublishTopicEvent(topicUpdateOrder, o)

			case ts := <-d.switchChan:
				d.isTradeTime = ts
			}
		}
	}()
}

func (d *DTFuture) cancelOverTimeOrder(order *entity.FutureOrder) {
	if time.Since(order.OrderTime) < time.Duration(d.tradeConfig.BuySellWaitTime)*time.Second {
		return
	}

	result, err := d.sc.CancelFuture(order.OrderID)
	if err != nil {
		logger.Error(err)
		return
	}

	if entity.StringToOrderStatus(result.GetStatus()) != entity.StatusCancelled {
		logger.Error("Cancel future order failed", result.GetStatus())
		return
	}

	if e := d.sc.NotifyToSlack(fmt.Sprintf("Cancelled %s", order.String())); e != nil {
		logger.Errorf("Notify to slack failed: %v", e)
	}
}

func (d *DTFuture) processTick() {
	go func() {
		for {
			d.lastTick = <-d.tickChan
			d.tickArr = append(d.tickArr, d.lastTick)

			if len(d.tickArr) > 1 && d.tickArr[len(d.tickArr)-1].TickTime.Sub(d.tickArr[len(d.tickArr)-2].TickTime) > time.Second {
				d.tickArr = entity.RealTimeFutureTickArr{}
				d.lastTickRate = 0
				continue
			}

			if d.lastTick.TickTime.Sub(d.tickArr[0].TickTime) > time.Duration(d.tradeConfig.TickInterval)*time.Second {
				d.tickArr = d.tickArr[1:]
			}

			if o := d.generateOrder(); o != nil {
				var wg sync.WaitGroup
				for i := 0; i < int(d.orderQuantity); i++ {
					wg.Add(1)
					go d.addTrader(o, &wg)
				}
				wg.Wait()
			}

			d.sendTickToTrader(d.lastTick)
		}
	}()
}

func (d *DTFuture) sendTickToTrader(tick *entity.RealTimeFutureTick) {
	d.traderMapLock.RLock()
	defer d.traderMapLock.RUnlock()

	for _, trader := range d.traderMap {
		if ch := trader.TickChan(); ch != nil {
			ch <- tick
		}
	}
}

func (d *DTFuture) generateOrder() *entity.FutureOrder {
	if !d.tradeConfig.AllowTrade || !d.isTradeTime {
		return nil
	}

	if len(d.tickArr) < int(d.tradeConfig.TickInterval) {
		return nil
	}

	d.traderMapLock.RLock()
	defer d.traderMapLock.RUnlock()
	if len(d.traderMap) > 0 {
		return nil
	}

	outInRatio, tickRate := d.tickArr.GetOutInRatioAndRate()
	defer func() {
		d.lastTickRate = tickRate
	}()

	if d.lastTickRate < d.tradeConfig.RateLimit || d.lastTickRate == 0 || 100*(tickRate-d.lastTickRate)/d.lastTickRate < d.tradeConfig.RateChangeRatio {
		return nil
	}

	switch {
	case outInRatio > d.tradeConfig.OutInRatio:
		return &entity.FutureOrder{
			Code: d.code,
			BaseOrder: entity.BaseOrder{
				Action:   entity.ActionBuy,
				Price:    d.lastTick.Close - 1,
				Quantity: orderQtyUnit,
			},
		}
	case 100-outInRatio < d.tradeConfig.OutInRatio:
		return &entity.FutureOrder{
			Code: d.code,
			BaseOrder: entity.BaseOrder{
				Action:   entity.ActionSell,
				Price:    d.lastTick.Close + 1,
				Quantity: orderQtyUnit,
			},
		}
	default:
		return nil
	}
}

func (d *DTFuture) addTrader(order *entity.FutureOrder, wg *sync.WaitGroup) {
	defer wg.Done()

	orderWithCfg := orderWithCfg{
		order:            order,
		cfg:              d.tradeConfig,
		lastTradeOutTime: time.Now().Add(time.Duration(d.tradeConfig.MaxHoldTime) * time.Minute),
	}

	if trader := NewDTTraderFuture(orderWithCfg, d.sc, d.localBus); trader != nil {
		d.traderMapLock.Lock()
		d.traderMap[trader.id] = trader
		d.traderMapLock.Unlock()
	}
}

func (d *DTFuture) removeDoneTrader(id string) {
	if trader := d.traderMap[id]; trader != nil {
		d.traderMapLock.Lock()
		close(trader.tickChan)
		delete(d.traderMap, id)
		d.traderMapLock.Unlock()
	}
}

func (d *DTFuture) Notify() chan *entity.FutureOrder {
	return d.notify
}

func (d *DTFuture) TickChan() chan *entity.RealTimeFutureTick {
	return d.tickChan
}

func (d *DTFuture) SwitchChan() chan bool {
	return d.switchChan
}
