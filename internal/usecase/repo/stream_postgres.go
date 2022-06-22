package repo

import (
	"context"
	"errors"

	"toc-machine-trading/internal/entity"
	"toc-machine-trading/pkg/postgres"

	"github.com/Masterminds/squirrel"
	"github.com/google/go-cmp/cmp"
	"github.com/jackc/pgx/v4"
)

// StreamRepo -.
type StreamRepo struct {
	*postgres.Postgres
}

// NewStream -.
func NewStream(pg *postgres.Postgres) *StreamRepo {
	return &StreamRepo{pg}
}

// InsertEvent -.
func (r *StreamRepo) InsertEvent(ctx context.Context, t *entity.SinopacEvent) error {
	builder := r.Builder.Insert(tableNameEvent).
		Columns("event, event_code, info, response, event_time").
		Values(t.Event, t.EventCode, t.Info, t.Response, t.EventTime)

	if sql, args, err := builder.ToSql(); err != nil {
		return err
	} else if _, err := r.Pool.Exec(ctx, sql, args...); err != nil {
		return err
	}

	return nil
}

// InserOrUpdatetOrder -.
func (r *StreamRepo) InserOrUpdatetOrder(ctx context.Context, t *entity.Order) error {
	dbOrder, err := r.QueryOrderByID(ctx, t.OrderID)
	if err != nil {
		return err
	}

	builder := r.Builder.Insert(tableNameTradeOrder).
		Columns("order_id, status, order_time, stock_num, action, price, quantity, trade_time")
	if dbOrder == nil {
		builder = builder.Values(t.OrderID, t.Status, t.OrderTime, t.StockNum, t.Action, t.Price, t.Quantity, t.TradeTime)
		if sql, args, err := builder.ToSql(); err != nil {
			return err
		} else if _, err := r.Pool.Exec(ctx, sql, args...); err != nil {
			return err
		}
	} else if !cmp.Equal(t, dbOrder) {
		b := r.Builder.
			Update(tableNameTradeOrder).
			Set("order_id", t.OrderID).
			Set("status", t.Status).
			Set("order_time", t.OrderTime).
			Set("stock_num", t.StockNum).
			Set("action", t.Action).
			Set("price", t.Price).
			Set("quantity", t.Quantity).
			Set("trade_time", t.TradeTime).
			Where(squirrel.Eq{"order_id": t.OrderID})
		if sql, args, err := b.ToSql(); err != nil {
			return err
		} else if _, err := r.Pool.Exec(ctx, sql, args...); err != nil {
			return err
		}
	}
	return nil
}

// QueryOrderByID -.
func (r *StreamRepo) QueryOrderByID(ctx context.Context, orderID string) (*entity.Order, error) {
	sql, _, err := r.Builder.
		Select("order_id, status, order_time, stock_num, action, price, quantity, trade_time, number, name, exchange, category, day_trade, last_close").
		From(tableNameTradeOrder).
		Where(squirrel.Eq{"order_id": orderID}).
		Join("basic_stock ON trade_order.stock_num = basic_stock.number").ToSql()
	if err != nil {
		return nil, err
	}
	row := r.Pool.QueryRow(ctx, sql)
	e := entity.Order{Stock: new(entity.Stock)}
	if err := row.Scan(&e.OrderID, &e.Status, &e.OrderTime, &e.StockNum, &e.Action, &e.Price, &e.Quantity, &e.TradeTime,
		&e.Stock.Number, &e.Stock.Name, &e.Stock.Exchange, &e.Stock.Category, &e.Stock.DayTrade, &e.Stock.LastClose); err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}
	return &e, nil
}
