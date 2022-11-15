package websocket

import (
	"context"
	"time"

	"tmt/internal/entity"
)

type futureOrder struct {
	Code   string             `json:"code"`
	Action entity.OrderAction `json:"action"`
	Price  float64            `json:"price"`
	Qty    int64              `json:"qty"`
}

type tradeRate struct {
	OutRate int64 `json:"out_rate"`
	InRate  int64 `json:"in_rate"`
}

func (w *WSRouter) processTrade(clientMsg msg) {
	if clientMsg.FutureOrder == nil {
		return
	}

	order := &entity.FutureOrder{
		Code: clientMsg.FutureOrder.Code,
		BaseOrder: entity.BaseOrder{
			Action:   clientMsg.FutureOrder.Action,
			Quantity: clientMsg.FutureOrder.Qty,
			Price:    clientMsg.FutureOrder.Price,
		},
	}

	switch clientMsg.FutureOrder.Action {
	case entity.ActionBuy:
		_, _, err := w.o.BuyFuture(order)
		if err != nil {
			w.msgChan <- errMsg{ErrMsg: err.Error()}
			log.Error(err)
		}

	case entity.ActionSell:
		_, _, err := w.o.SellFuture(order)
		if err != nil {
			w.msgChan <- errMsg{ErrMsg: err.Error()}
			log.Error(err)
		}
	}
}

func (w *WSRouter) sendFuture(ctx context.Context) {
	timestamp := time.Now().UnixNano()
	tickChan := make(chan *entity.RealTimeFutureTick)

	snapshot, err := w.s.GetFutureSnapshotByCode(w.s.GetMainFutureCode())
	if err != nil {
		w.msgChan <- errMsg{ErrMsg: err.Error()}
	}

	var tickType, chgType int64
	switch snapshot.TickType {
	case "Sell":
		tickType = 1
	case "Buy":
		tickType = 2
	}

	switch snapshot.ChgType {
	case "LimitUp":
		chgType = 1
	case "Up":
		chgType = 2
	case "Unchanged":
		chgType = 3
	case "Dowm":
		chgType = 4
	case "LimitDown":
		chgType = 5
	}

	w.msgChan <- &entity.RealTimeFutureTick{
		Code:        snapshot.Code,
		TickTime:    snapshot.SnapTime,
		Open:        snapshot.Open,
		Close:       snapshot.Close,
		High:        snapshot.High,
		Low:         snapshot.Low,
		Amount:      float64(snapshot.Amount),
		TotalAmount: float64(snapshot.AmountSum),
		Volume:      snapshot.Volume,
		TotalVolume: snapshot.VolumeSum,
		TickType:    tickType,
		ChgType:     chgType,
		PriceChg:    snapshot.PriceChg,
		PctChg:      snapshot.PctChg,
	}
	go w.processTickArr(tickChan)

	defer w.s.DeleteFutureRealTimeConnection(timestamp)
	w.s.NewFutureRealTimeConnection(timestamp, tickChan)
	<-ctx.Done()
}

func (w *WSRouter) processTickArr(tickChan chan *entity.RealTimeFutureTick) {
	var tickArr []*entity.RealTimeFutureTick
	for {
		tick, ok := <-tickChan
		if !ok {
			close(w.msgChan)
			return
		}
		tickArr = append(tickArr, tick)

		var outVolume, inVolume int64
		for i := len(tickArr) - 1; i >= 0; i-- {
			if time.Since(tickArr[i].TickTime) > 30*time.Second {
				tickArr = tickArr[i+1:]
				break
			}

			switch tickArr[i].TickType {
			case 1:
				outVolume += tickArr[i].Volume
			case 2:
				inVolume += tickArr[i].Volume
			}
		}

		w.msgChan <- &tradeRate{
			OutRate: outVolume,
			InRate:  inVolume,
		}
		w.msgChan <- tick
	}
}
