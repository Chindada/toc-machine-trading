// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces_usecase.go

// Package usecase_test is a generated GoMock package.
package usecase_test

import (
	context "context"
	reflect "reflect"
	time "time"
	config "tmt/cmd/config"
	entity "tmt/internal/entity"
	simulator "tmt/internal/usecase/module/simulator"

	gomock "github.com/golang/mock/gomock"
)

// MockBasic is a mock of Basic interface.
type MockBasic struct {
	ctrl     *gomock.Controller
	recorder *MockBasicMockRecorder
}

// MockBasicMockRecorder is the mock recorder for MockBasic.
type MockBasicMockRecorder struct {
	mock *MockBasic
}

// NewMockBasic creates a new mock instance.
func NewMockBasic(ctrl *gomock.Controller) *MockBasic {
	mock := &MockBasic{ctrl: ctrl}
	mock.recorder = &MockBasicMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBasic) EXPECT() *MockBasicMockRecorder {
	return m.recorder
}

// GetAllRepoStock mocks base method.
func (m *MockBasic) GetAllRepoStock(ctx context.Context) ([]*entity.Stock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllRepoStock", ctx)
	ret0, _ := ret[0].([]*entity.Stock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllRepoStock indicates an expected call of GetAllRepoStock.
func (mr *MockBasicMockRecorder) GetAllRepoStock(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllRepoStock", reflect.TypeOf((*MockBasic)(nil).GetAllRepoStock), ctx)
}

// GetConfig mocks base method.
func (m *MockBasic) GetConfig() *config.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig")
	ret0, _ := ret[0].(*config.Config)
	return ret0
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockBasicMockRecorder) GetConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockBasic)(nil).GetConfig))
}

// GetShioajiUsage mocks base method.
func (m *MockBasic) GetShioajiUsage() (*entity.ShioajiUsage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShioajiUsage")
	ret0, _ := ret[0].(*entity.ShioajiUsage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShioajiUsage indicates an expected call of GetShioajiUsage.
func (mr *MockBasicMockRecorder) GetShioajiUsage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShioajiUsage", reflect.TypeOf((*MockBasic)(nil).GetShioajiUsage))
}

// MockTarget is a mock of Target interface.
type MockTarget struct {
	ctrl     *gomock.Controller
	recorder *MockTargetMockRecorder
}

// MockTargetMockRecorder is the mock recorder for MockTarget.
type MockTargetMockRecorder struct {
	mock *MockTarget
}

// NewMockTarget creates a new mock instance.
func NewMockTarget(ctrl *gomock.Controller) *MockTarget {
	mock := &MockTarget{ctrl: ctrl}
	mock.recorder = &MockTargetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTarget) EXPECT() *MockTargetMockRecorder {
	return m.recorder
}

// GetTargets mocks base method.
func (m *MockTarget) GetTargets(ctx context.Context) []*entity.StockTarget {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTargets", ctx)
	ret0, _ := ret[0].([]*entity.StockTarget)
	return ret0
}

// GetTargets indicates an expected call of GetTargets.
func (mr *MockTargetMockRecorder) GetTargets(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTargets", reflect.TypeOf((*MockTarget)(nil).GetTargets), ctx)
}

// MockHistory is a mock of History interface.
type MockHistory struct {
	ctrl     *gomock.Controller
	recorder *MockHistoryMockRecorder
}

// MockHistoryMockRecorder is the mock recorder for MockHistory.
type MockHistoryMockRecorder struct {
	mock *MockHistory
}

// NewMockHistory creates a new mock instance.
func NewMockHistory(ctrl *gomock.Controller) *MockHistory {
	mock := &MockHistory{ctrl: ctrl}
	mock.recorder = &MockHistoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHistory) EXPECT() *MockHistoryMockRecorder {
	return m.recorder
}

// FetchFutureHistoryKbar mocks base method.
func (m *MockHistory) FetchFutureHistoryKbar(code string, date time.Time) ([]*entity.FutureHistoryKbar, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchFutureHistoryKbar", code, date)
	ret0, _ := ret[0].([]*entity.FutureHistoryKbar)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchFutureHistoryKbar indicates an expected call of FetchFutureHistoryKbar.
func (mr *MockHistoryMockRecorder) FetchFutureHistoryKbar(code, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchFutureHistoryKbar", reflect.TypeOf((*MockHistory)(nil).FetchFutureHistoryKbar), code, date)
}

// GetDayKbarByStockNumDate mocks base method.
func (m *MockHistory) GetDayKbarByStockNumDate(stockNum string, date time.Time) *entity.StockHistoryKbar {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDayKbarByStockNumDate", stockNum, date)
	ret0, _ := ret[0].(*entity.StockHistoryKbar)
	return ret0
}

// GetDayKbarByStockNumDate indicates an expected call of GetDayKbarByStockNumDate.
func (mr *MockHistoryMockRecorder) GetDayKbarByStockNumDate(stockNum, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDayKbarByStockNumDate", reflect.TypeOf((*MockHistory)(nil).GetDayKbarByStockNumDate), stockNum, date)
}

// GetTradeDay mocks base method.
func (m *MockHistory) GetTradeDay() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTradeDay")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetTradeDay indicates an expected call of GetTradeDay.
func (mr *MockHistoryMockRecorder) GetTradeDay() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTradeDay", reflect.TypeOf((*MockHistory)(nil).GetTradeDay))
}

// SimulateMulti mocks base method.
func (m *MockHistory) SimulateMulti(cond []*config.TradeFuture) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SimulateMulti", cond)
}

// SimulateMulti indicates an expected call of SimulateMulti.
func (mr *MockHistoryMockRecorder) SimulateMulti(cond interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SimulateMulti", reflect.TypeOf((*MockHistory)(nil).SimulateMulti), cond)
}

// SimulateOne mocks base method.
func (m *MockHistory) SimulateOne(cond *config.TradeFuture) *simulator.SimulateBalance {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SimulateOne", cond)
	ret0, _ := ret[0].(*simulator.SimulateBalance)
	return ret0
}

// SimulateOne indicates an expected call of SimulateOne.
func (mr *MockHistoryMockRecorder) SimulateOne(cond interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SimulateOne", reflect.TypeOf((*MockHistory)(nil).SimulateOne), cond)
}

// MockRealTime is a mock of RealTime interface.
type MockRealTime struct {
	ctrl     *gomock.Controller
	recorder *MockRealTimeMockRecorder
}

// MockRealTimeMockRecorder is the mock recorder for MockRealTime.
type MockRealTimeMockRecorder struct {
	mock *MockRealTime
}

// NewMockRealTime creates a new mock instance.
func NewMockRealTime(ctrl *gomock.Controller) *MockRealTime {
	mock := &MockRealTime{ctrl: ctrl}
	mock.recorder = &MockRealTimeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRealTime) EXPECT() *MockRealTimeMockRecorder {
	return m.recorder
}

// DeleteFutureRealTimeClient mocks base method.
func (m *MockRealTime) DeleteFutureRealTimeClient(connectionID string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteFutureRealTimeClient", connectionID)
}

// DeleteFutureRealTimeClient indicates an expected call of DeleteFutureRealTimeClient.
func (mr *MockRealTimeMockRecorder) DeleteFutureRealTimeClient(connectionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFutureRealTimeClient", reflect.TypeOf((*MockRealTime)(nil).DeleteFutureRealTimeClient), connectionID)
}

// GetFutureSnapshotByCode mocks base method.
func (m *MockRealTime) GetFutureSnapshotByCode(code string) (*entity.FutureSnapShot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFutureSnapshotByCode", code)
	ret0, _ := ret[0].(*entity.FutureSnapShot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFutureSnapshotByCode indicates an expected call of GetFutureSnapshotByCode.
func (mr *MockRealTimeMockRecorder) GetFutureSnapshotByCode(code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFutureSnapshotByCode", reflect.TypeOf((*MockRealTime)(nil).GetFutureSnapshotByCode), code)
}

// GetMainFuture mocks base method.
func (m *MockRealTime) GetMainFuture() *entity.Future {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMainFuture")
	ret0, _ := ret[0].(*entity.Future)
	return ret0
}

// GetMainFuture indicates an expected call of GetMainFuture.
func (mr *MockRealTimeMockRecorder) GetMainFuture() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMainFuture", reflect.TypeOf((*MockRealTime)(nil).GetMainFuture))
}

// GetOTCSnapshot mocks base method.
func (m *MockRealTime) GetOTCSnapshot(ctx context.Context) (*entity.StockSnapShot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOTCSnapshot", ctx)
	ret0, _ := ret[0].(*entity.StockSnapShot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOTCSnapshot indicates an expected call of GetOTCSnapshot.
func (mr *MockRealTimeMockRecorder) GetOTCSnapshot(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOTCSnapshot", reflect.TypeOf((*MockRealTime)(nil).GetOTCSnapshot), ctx)
}

// GetStockSnapshotByNumArr mocks base method.
func (m *MockRealTime) GetStockSnapshotByNumArr(stockNumArr []string) ([]*entity.StockSnapShot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStockSnapshotByNumArr", stockNumArr)
	ret0, _ := ret[0].([]*entity.StockSnapShot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStockSnapshotByNumArr indicates an expected call of GetStockSnapshotByNumArr.
func (mr *MockRealTimeMockRecorder) GetStockSnapshotByNumArr(stockNumArr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStockSnapshotByNumArr", reflect.TypeOf((*MockRealTime)(nil).GetStockSnapshotByNumArr), stockNumArr)
}

// GetTSESnapshot mocks base method.
func (m *MockRealTime) GetTSESnapshot(ctx context.Context) (*entity.StockSnapShot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTSESnapshot", ctx)
	ret0, _ := ret[0].(*entity.StockSnapShot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTSESnapshot indicates an expected call of GetTSESnapshot.
func (mr *MockRealTimeMockRecorder) GetTSESnapshot(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTSESnapshot", reflect.TypeOf((*MockRealTime)(nil).GetTSESnapshot), ctx)
}

// GetTradeIndex mocks base method.
func (m *MockRealTime) GetTradeIndex() *entity.TradeIndex {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTradeIndex")
	ret0, _ := ret[0].(*entity.TradeIndex)
	return ret0
}

// GetTradeIndex indicates an expected call of GetTradeIndex.
func (mr *MockRealTimeMockRecorder) GetTradeIndex() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTradeIndex", reflect.TypeOf((*MockRealTime)(nil).GetTradeIndex))
}

// NewFutureRealTimeClient mocks base method.
func (m *MockRealTime) NewFutureRealTimeClient(tickChan chan *entity.RealTimeFutureTick, orderStatusChan chan interface{}, connectionID string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NewFutureRealTimeClient", tickChan, orderStatusChan, connectionID)
}

// NewFutureRealTimeClient indicates an expected call of NewFutureRealTimeClient.
func (mr *MockRealTimeMockRecorder) NewFutureRealTimeClient(tickChan, orderStatusChan, connectionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewFutureRealTimeClient", reflect.TypeOf((*MockRealTime)(nil).NewFutureRealTimeClient), tickChan, orderStatusChan, connectionID)
}

// MockTrade is a mock of Trade interface.
type MockTrade struct {
	ctrl     *gomock.Controller
	recorder *MockTradeMockRecorder
}

// MockTradeMockRecorder is the mock recorder for MockTrade.
type MockTradeMockRecorder struct {
	mock *MockTrade
}

// NewMockTrade creates a new mock instance.
func NewMockTrade(ctrl *gomock.Controller) *MockTrade {
	mock := &MockTrade{ctrl: ctrl}
	mock.recorder = &MockTradeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrade) EXPECT() *MockTradeMockRecorder {
	return m.recorder
}

// BuyFuture mocks base method.
func (m *MockTrade) BuyFuture(order *entity.FutureOrder) (string, entity.OrderStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuyFuture", order)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(entity.OrderStatus)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// BuyFuture indicates an expected call of BuyFuture.
func (mr *MockTradeMockRecorder) BuyFuture(order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuyFuture", reflect.TypeOf((*MockTrade)(nil).BuyFuture), order)
}

// CalculateBuyCost mocks base method.
func (m *MockTrade) CalculateBuyCost(price float64, quantity int64) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculateBuyCost", price, quantity)
	ret0, _ := ret[0].(int64)
	return ret0
}

// CalculateBuyCost indicates an expected call of CalculateBuyCost.
func (mr *MockTradeMockRecorder) CalculateBuyCost(price, quantity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculateBuyCost", reflect.TypeOf((*MockTrade)(nil).CalculateBuyCost), price, quantity)
}

// CalculateSellCost mocks base method.
func (m *MockTrade) CalculateSellCost(price float64, quantity int64) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculateSellCost", price, quantity)
	ret0, _ := ret[0].(int64)
	return ret0
}

// CalculateSellCost indicates an expected call of CalculateSellCost.
func (mr *MockTradeMockRecorder) CalculateSellCost(price, quantity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculateSellCost", reflect.TypeOf((*MockTrade)(nil).CalculateSellCost), price, quantity)
}

// CalculateTradeDiscount mocks base method.
func (m *MockTrade) CalculateTradeDiscount(price float64, quantity int64) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculateTradeDiscount", price, quantity)
	ret0, _ := ret[0].(int64)
	return ret0
}

// CalculateTradeDiscount indicates an expected call of CalculateTradeDiscount.
func (mr *MockTradeMockRecorder) CalculateTradeDiscount(price, quantity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculateTradeDiscount", reflect.TypeOf((*MockTrade)(nil).CalculateTradeDiscount), price, quantity)
}

// CancelFutureOrderByID mocks base method.
func (m *MockTrade) CancelFutureOrderByID(orderID string) (string, entity.OrderStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelFutureOrderByID", orderID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(entity.OrderStatus)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CancelFutureOrderByID indicates an expected call of CancelFutureOrderByID.
func (mr *MockTradeMockRecorder) CancelFutureOrderByID(orderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelFutureOrderByID", reflect.TypeOf((*MockTrade)(nil).CancelFutureOrderByID), orderID)
}

// GetAccountBalance mocks base method.
func (m *MockTrade) GetAccountBalance(ctx context.Context) ([]*entity.AccountBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountBalance", ctx)
	ret0, _ := ret[0].([]*entity.AccountBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountBalance indicates an expected call of GetAccountBalance.
func (mr *MockTradeMockRecorder) GetAccountBalance(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountBalance", reflect.TypeOf((*MockTrade)(nil).GetAccountBalance), ctx)
}

// GetAllFutureOrder mocks base method.
func (m *MockTrade) GetAllFutureOrder(ctx context.Context) ([]*entity.FutureOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFutureOrder", ctx)
	ret0, _ := ret[0].([]*entity.FutureOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFutureOrder indicates an expected call of GetAllFutureOrder.
func (mr *MockTradeMockRecorder) GetAllFutureOrder(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFutureOrder", reflect.TypeOf((*MockTrade)(nil).GetAllFutureOrder), ctx)
}

// GetAllFutureTradeBalance mocks base method.
func (m *MockTrade) GetAllFutureTradeBalance(ctx context.Context) ([]*entity.FutureTradeBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFutureTradeBalance", ctx)
	ret0, _ := ret[0].([]*entity.FutureTradeBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFutureTradeBalance indicates an expected call of GetAllFutureTradeBalance.
func (mr *MockTradeMockRecorder) GetAllFutureTradeBalance(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFutureTradeBalance", reflect.TypeOf((*MockTrade)(nil).GetAllFutureTradeBalance), ctx)
}

// GetAllStockOrder mocks base method.
func (m *MockTrade) GetAllStockOrder(ctx context.Context) ([]*entity.StockOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllStockOrder", ctx)
	ret0, _ := ret[0].([]*entity.StockOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllStockOrder indicates an expected call of GetAllStockOrder.
func (mr *MockTradeMockRecorder) GetAllStockOrder(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllStockOrder", reflect.TypeOf((*MockTrade)(nil).GetAllStockOrder), ctx)
}

// GetAllStockTradeBalance mocks base method.
func (m *MockTrade) GetAllStockTradeBalance(ctx context.Context) ([]*entity.StockTradeBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllStockTradeBalance", ctx)
	ret0, _ := ret[0].([]*entity.StockTradeBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllStockTradeBalance indicates an expected call of GetAllStockTradeBalance.
func (mr *MockTradeMockRecorder) GetAllStockTradeBalance(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllStockTradeBalance", reflect.TypeOf((*MockTrade)(nil).GetAllStockTradeBalance), ctx)
}

// GetFutureOrderByTradeDay mocks base method.
func (m *MockTrade) GetFutureOrderByTradeDay(ctx context.Context, tradeDay string) ([]*entity.FutureOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFutureOrderByTradeDay", ctx, tradeDay)
	ret0, _ := ret[0].([]*entity.FutureOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFutureOrderByTradeDay indicates an expected call of GetFutureOrderByTradeDay.
func (mr *MockTradeMockRecorder) GetFutureOrderByTradeDay(ctx, tradeDay interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFutureOrderByTradeDay", reflect.TypeOf((*MockTrade)(nil).GetFutureOrderByTradeDay), ctx, tradeDay)
}

// GetFuturePosition mocks base method.
func (m *MockTrade) GetFuturePosition() ([]*entity.FuturePosition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFuturePosition")
	ret0, _ := ret[0].([]*entity.FuturePosition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFuturePosition indicates an expected call of GetFuturePosition.
func (mr *MockTradeMockRecorder) GetFuturePosition() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFuturePosition", reflect.TypeOf((*MockTrade)(nil).GetFuturePosition))
}

// GetLastFutureTradeBalance mocks base method.
func (m *MockTrade) GetLastFutureTradeBalance(ctx context.Context) (*entity.FutureTradeBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastFutureTradeBalance", ctx)
	ret0, _ := ret[0].(*entity.FutureTradeBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastFutureTradeBalance indicates an expected call of GetLastFutureTradeBalance.
func (mr *MockTradeMockRecorder) GetLastFutureTradeBalance(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastFutureTradeBalance", reflect.TypeOf((*MockTrade)(nil).GetLastFutureTradeBalance), ctx)
}

// GetLastStockTradeBalance mocks base method.
func (m *MockTrade) GetLastStockTradeBalance(ctx context.Context) (*entity.StockTradeBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastStockTradeBalance", ctx)
	ret0, _ := ret[0].(*entity.StockTradeBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastStockTradeBalance indicates an expected call of GetLastStockTradeBalance.
func (mr *MockTradeMockRecorder) GetLastStockTradeBalance(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastStockTradeBalance", reflect.TypeOf((*MockTrade)(nil).GetLastStockTradeBalance), ctx)
}

// IsFutureTradeTime mocks base method.
func (m *MockTrade) IsFutureTradeTime() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFutureTradeTime")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsFutureTradeTime indicates an expected call of IsFutureTradeTime.
func (mr *MockTradeMockRecorder) IsFutureTradeTime() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFutureTradeTime", reflect.TypeOf((*MockTrade)(nil).IsFutureTradeTime))
}

// ManualInsertFutureOrder mocks base method.
func (m *MockTrade) ManualInsertFutureOrder(ctx context.Context, order *entity.FutureOrder) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ManualInsertFutureOrder", ctx, order)
	ret0, _ := ret[0].(error)
	return ret0
}

// ManualInsertFutureOrder indicates an expected call of ManualInsertFutureOrder.
func (mr *MockTradeMockRecorder) ManualInsertFutureOrder(ctx, order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ManualInsertFutureOrder", reflect.TypeOf((*MockTrade)(nil).ManualInsertFutureOrder), ctx, order)
}

// MoveFutureOrderToLatestTradeDay mocks base method.
func (m *MockTrade) MoveFutureOrderToLatestTradeDay(ctx context.Context, orderID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MoveFutureOrderToLatestTradeDay", ctx, orderID)
	ret0, _ := ret[0].(error)
	return ret0
}

// MoveFutureOrderToLatestTradeDay indicates an expected call of MoveFutureOrderToLatestTradeDay.
func (mr *MockTradeMockRecorder) MoveFutureOrderToLatestTradeDay(ctx, orderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveFutureOrderToLatestTradeDay", reflect.TypeOf((*MockTrade)(nil).MoveFutureOrderToLatestTradeDay), ctx, orderID)
}

// MoveStockOrderToLatestTradeDay mocks base method.
func (m *MockTrade) MoveStockOrderToLatestTradeDay(ctx context.Context, orderID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MoveStockOrderToLatestTradeDay", ctx, orderID)
	ret0, _ := ret[0].(error)
	return ret0
}

// MoveStockOrderToLatestTradeDay indicates an expected call of MoveStockOrderToLatestTradeDay.
func (mr *MockTradeMockRecorder) MoveStockOrderToLatestTradeDay(ctx, orderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveStockOrderToLatestTradeDay", reflect.TypeOf((*MockTrade)(nil).MoveStockOrderToLatestTradeDay), ctx, orderID)
}

// SellFuture mocks base method.
func (m *MockTrade) SellFuture(order *entity.FutureOrder) (string, entity.OrderStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SellFuture", order)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(entity.OrderStatus)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SellFuture indicates an expected call of SellFuture.
func (mr *MockTradeMockRecorder) SellFuture(order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SellFuture", reflect.TypeOf((*MockTrade)(nil).SellFuture), order)
}

// UpdateTradeBalanceByTradeDay mocks base method.
func (m *MockTrade) UpdateTradeBalanceByTradeDay(ctx context.Context, date string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTradeBalanceByTradeDay", ctx, date)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTradeBalanceByTradeDay indicates an expected call of UpdateTradeBalanceByTradeDay.
func (mr *MockTradeMockRecorder) UpdateTradeBalanceByTradeDay(ctx, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTradeBalanceByTradeDay", reflect.TypeOf((*MockTrade)(nil).UpdateTradeBalanceByTradeDay), ctx, date)
}

// MockAnalyze is a mock of Analyze interface.
type MockAnalyze struct {
	ctrl     *gomock.Controller
	recorder *MockAnalyzeMockRecorder
}

// MockAnalyzeMockRecorder is the mock recorder for MockAnalyze.
type MockAnalyzeMockRecorder struct {
	mock *MockAnalyze
}

// NewMockAnalyze creates a new mock instance.
func NewMockAnalyze(ctrl *gomock.Controller) *MockAnalyze {
	mock := &MockAnalyze{ctrl: ctrl}
	mock.recorder = &MockAnalyzeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAnalyze) EXPECT() *MockAnalyzeMockRecorder {
	return m.recorder
}

// GetRebornMap mocks base method.
func (m *MockAnalyze) GetRebornMap(ctx context.Context) map[time.Time][]entity.Stock {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRebornMap", ctx)
	ret0, _ := ret[0].(map[time.Time][]entity.Stock)
	return ret0
}

// GetRebornMap indicates an expected call of GetRebornMap.
func (mr *MockAnalyzeMockRecorder) GetRebornMap(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRebornMap", reflect.TypeOf((*MockAnalyze)(nil).GetRebornMap), ctx)
}
