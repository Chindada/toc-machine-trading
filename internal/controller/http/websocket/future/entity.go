package future

import "tmt/internal/entity"

type AutomationType int

func (a AutomationType) String() string {
	switch a {
	case AutomationByBalance:
		return "By balance"
	case AutomationByTimePeriod:
		return "By time period"
	case AutomationByTimePeriodAndBalance:
		return "By time period and balance"
	}
	return "unknown"
}

const (
	AutomationNone AutomationType = iota
	AutomationByBalance
	AutomationByTimePeriod
	AutomationByTimePeriodAndBalance
)

type clientOrder struct {
	Code   string               `json:"code"`
	Action entity.OrderAction   `json:"action"`
	Price  float64              `json:"price"`
	Qty    int64                `json:"qty"`
	Option halfAutomationOption `json:"option"`
}

type halfAutomationOption struct {
	AutomationType AutomationType `json:"automation_type"`
	ByBalanceHigh  float64        `json:"by_balance_high"`
	ByBalanceLow   float64        `json:"by_balance_low"`
	ByTimePeriod   int64          `json:"by_time_period"`
}

func (f *clientOrder) toFutureOrder() *entity.FutureOrder {
	return &entity.FutureOrder{
		Code: f.Code,
		BaseOrder: entity.BaseOrder{
			Action:   f.Action,
			Quantity: f.Qty,
			Price:    f.Price,
		},
	}
}
