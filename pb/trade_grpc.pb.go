// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: forwarder/trade.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TradeInterface_BuyStock_FullMethodName                  = "/forwarder.TradeInterface/BuyStock"
	TradeInterface_SellStock_FullMethodName                 = "/forwarder.TradeInterface/SellStock"
	TradeInterface_BuyOddStock_FullMethodName               = "/forwarder.TradeInterface/BuyOddStock"
	TradeInterface_SellOddStock_FullMethodName              = "/forwarder.TradeInterface/SellOddStock"
	TradeInterface_SellFirstStock_FullMethodName            = "/forwarder.TradeInterface/SellFirstStock"
	TradeInterface_CancelStock_FullMethodName               = "/forwarder.TradeInterface/CancelStock"
	TradeInterface_BuyFuture_FullMethodName                 = "/forwarder.TradeInterface/BuyFuture"
	TradeInterface_SellFuture_FullMethodName                = "/forwarder.TradeInterface/SellFuture"
	TradeInterface_SellFirstFuture_FullMethodName           = "/forwarder.TradeInterface/SellFirstFuture"
	TradeInterface_CancelFuture_FullMethodName              = "/forwarder.TradeInterface/CancelFuture"
	TradeInterface_BuyOption_FullMethodName                 = "/forwarder.TradeInterface/BuyOption"
	TradeInterface_SellOption_FullMethodName                = "/forwarder.TradeInterface/SellOption"
	TradeInterface_SellFirstOption_FullMethodName           = "/forwarder.TradeInterface/SellFirstOption"
	TradeInterface_CancelOption_FullMethodName              = "/forwarder.TradeInterface/CancelOption"
	TradeInterface_GetLocalOrderStatusArr_FullMethodName    = "/forwarder.TradeInterface/GetLocalOrderStatusArr"
	TradeInterface_GetSimulateOrderStatusArr_FullMethodName = "/forwarder.TradeInterface/GetSimulateOrderStatusArr"
	TradeInterface_GetFuturePosition_FullMethodName         = "/forwarder.TradeInterface/GetFuturePosition"
	TradeInterface_GetStockPosition_FullMethodName          = "/forwarder.TradeInterface/GetStockPosition"
	TradeInterface_GetSettlement_FullMethodName             = "/forwarder.TradeInterface/GetSettlement"
	TradeInterface_GetAccountBalance_FullMethodName         = "/forwarder.TradeInterface/GetAccountBalance"
	TradeInterface_GetMargin_FullMethodName                 = "/forwarder.TradeInterface/GetMargin"
)

// TradeInterfaceClient is the client API for TradeInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TradeInterfaceClient interface {
	BuyStock(ctx context.Context, in *StockOrderDetail, opts ...grpc.CallOption) (*TradeResult, error)
	SellStock(ctx context.Context, in *StockOrderDetail, opts ...grpc.CallOption) (*TradeResult, error)
	BuyOddStock(ctx context.Context, in *OddStockOrderDetail, opts ...grpc.CallOption) (*TradeResult, error)
	SellOddStock(ctx context.Context, in *OddStockOrderDetail, opts ...grpc.CallOption) (*TradeResult, error)
	SellFirstStock(ctx context.Context, in *StockOrderDetail, opts ...grpc.CallOption) (*TradeResult, error)
	CancelStock(ctx context.Context, in *OrderID, opts ...grpc.CallOption) (*TradeResult, error)
	BuyFuture(ctx context.Context, in *FutureOrderDetail, opts ...grpc.CallOption) (*TradeResult, error)
	SellFuture(ctx context.Context, in *FutureOrderDetail, opts ...grpc.CallOption) (*TradeResult, error)
	SellFirstFuture(ctx context.Context, in *FutureOrderDetail, opts ...grpc.CallOption) (*TradeResult, error)
	CancelFuture(ctx context.Context, in *FutureOrderID, opts ...grpc.CallOption) (*TradeResult, error)
	BuyOption(ctx context.Context, in *OptionOrderDetail, opts ...grpc.CallOption) (*TradeResult, error)
	SellOption(ctx context.Context, in *OptionOrderDetail, opts ...grpc.CallOption) (*TradeResult, error)
	SellFirstOption(ctx context.Context, in *OptionOrderDetail, opts ...grpc.CallOption) (*TradeResult, error)
	CancelOption(ctx context.Context, in *OptionOrderID, opts ...grpc.CallOption) (*TradeResult, error)
	GetLocalOrderStatusArr(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetSimulateOrderStatusArr(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetFuturePosition(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FuturePositionArr, error)
	GetStockPosition(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StockPositionArr, error)
	GetSettlement(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SettlementList, error)
	GetAccountBalance(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AccountBalance, error)
	GetMargin(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Margin, error)
}

type tradeInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewTradeInterfaceClient(cc grpc.ClientConnInterface) TradeInterfaceClient {
	return &tradeInterfaceClient{cc}
}

func (c *tradeInterfaceClient) BuyStock(ctx context.Context, in *StockOrderDetail, opts ...grpc.CallOption) (*TradeResult, error) {
	out := new(TradeResult)
	err := c.cc.Invoke(ctx, TradeInterface_BuyStock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeInterfaceClient) SellStock(ctx context.Context, in *StockOrderDetail, opts ...grpc.CallOption) (*TradeResult, error) {
	out := new(TradeResult)
	err := c.cc.Invoke(ctx, TradeInterface_SellStock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeInterfaceClient) BuyOddStock(ctx context.Context, in *OddStockOrderDetail, opts ...grpc.CallOption) (*TradeResult, error) {
	out := new(TradeResult)
	err := c.cc.Invoke(ctx, TradeInterface_BuyOddStock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeInterfaceClient) SellOddStock(ctx context.Context, in *OddStockOrderDetail, opts ...grpc.CallOption) (*TradeResult, error) {
	out := new(TradeResult)
	err := c.cc.Invoke(ctx, TradeInterface_SellOddStock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeInterfaceClient) SellFirstStock(ctx context.Context, in *StockOrderDetail, opts ...grpc.CallOption) (*TradeResult, error) {
	out := new(TradeResult)
	err := c.cc.Invoke(ctx, TradeInterface_SellFirstStock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeInterfaceClient) CancelStock(ctx context.Context, in *OrderID, opts ...grpc.CallOption) (*TradeResult, error) {
	out := new(TradeResult)
	err := c.cc.Invoke(ctx, TradeInterface_CancelStock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeInterfaceClient) BuyFuture(ctx context.Context, in *FutureOrderDetail, opts ...grpc.CallOption) (*TradeResult, error) {
	out := new(TradeResult)
	err := c.cc.Invoke(ctx, TradeInterface_BuyFuture_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeInterfaceClient) SellFuture(ctx context.Context, in *FutureOrderDetail, opts ...grpc.CallOption) (*TradeResult, error) {
	out := new(TradeResult)
	err := c.cc.Invoke(ctx, TradeInterface_SellFuture_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeInterfaceClient) SellFirstFuture(ctx context.Context, in *FutureOrderDetail, opts ...grpc.CallOption) (*TradeResult, error) {
	out := new(TradeResult)
	err := c.cc.Invoke(ctx, TradeInterface_SellFirstFuture_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeInterfaceClient) CancelFuture(ctx context.Context, in *FutureOrderID, opts ...grpc.CallOption) (*TradeResult, error) {
	out := new(TradeResult)
	err := c.cc.Invoke(ctx, TradeInterface_CancelFuture_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeInterfaceClient) BuyOption(ctx context.Context, in *OptionOrderDetail, opts ...grpc.CallOption) (*TradeResult, error) {
	out := new(TradeResult)
	err := c.cc.Invoke(ctx, TradeInterface_BuyOption_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeInterfaceClient) SellOption(ctx context.Context, in *OptionOrderDetail, opts ...grpc.CallOption) (*TradeResult, error) {
	out := new(TradeResult)
	err := c.cc.Invoke(ctx, TradeInterface_SellOption_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeInterfaceClient) SellFirstOption(ctx context.Context, in *OptionOrderDetail, opts ...grpc.CallOption) (*TradeResult, error) {
	out := new(TradeResult)
	err := c.cc.Invoke(ctx, TradeInterface_SellFirstOption_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeInterfaceClient) CancelOption(ctx context.Context, in *OptionOrderID, opts ...grpc.CallOption) (*TradeResult, error) {
	out := new(TradeResult)
	err := c.cc.Invoke(ctx, TradeInterface_CancelOption_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeInterfaceClient) GetLocalOrderStatusArr(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TradeInterface_GetLocalOrderStatusArr_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeInterfaceClient) GetSimulateOrderStatusArr(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TradeInterface_GetSimulateOrderStatusArr_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeInterfaceClient) GetFuturePosition(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FuturePositionArr, error) {
	out := new(FuturePositionArr)
	err := c.cc.Invoke(ctx, TradeInterface_GetFuturePosition_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeInterfaceClient) GetStockPosition(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StockPositionArr, error) {
	out := new(StockPositionArr)
	err := c.cc.Invoke(ctx, TradeInterface_GetStockPosition_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeInterfaceClient) GetSettlement(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SettlementList, error) {
	out := new(SettlementList)
	err := c.cc.Invoke(ctx, TradeInterface_GetSettlement_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeInterfaceClient) GetAccountBalance(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AccountBalance, error) {
	out := new(AccountBalance)
	err := c.cc.Invoke(ctx, TradeInterface_GetAccountBalance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeInterfaceClient) GetMargin(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Margin, error) {
	out := new(Margin)
	err := c.cc.Invoke(ctx, TradeInterface_GetMargin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TradeInterfaceServer is the server API for TradeInterface service.
// All implementations must embed UnimplementedTradeInterfaceServer
// for forward compatibility
type TradeInterfaceServer interface {
	BuyStock(context.Context, *StockOrderDetail) (*TradeResult, error)
	SellStock(context.Context, *StockOrderDetail) (*TradeResult, error)
	BuyOddStock(context.Context, *OddStockOrderDetail) (*TradeResult, error)
	SellOddStock(context.Context, *OddStockOrderDetail) (*TradeResult, error)
	SellFirstStock(context.Context, *StockOrderDetail) (*TradeResult, error)
	CancelStock(context.Context, *OrderID) (*TradeResult, error)
	BuyFuture(context.Context, *FutureOrderDetail) (*TradeResult, error)
	SellFuture(context.Context, *FutureOrderDetail) (*TradeResult, error)
	SellFirstFuture(context.Context, *FutureOrderDetail) (*TradeResult, error)
	CancelFuture(context.Context, *FutureOrderID) (*TradeResult, error)
	BuyOption(context.Context, *OptionOrderDetail) (*TradeResult, error)
	SellOption(context.Context, *OptionOrderDetail) (*TradeResult, error)
	SellFirstOption(context.Context, *OptionOrderDetail) (*TradeResult, error)
	CancelOption(context.Context, *OptionOrderID) (*TradeResult, error)
	GetLocalOrderStatusArr(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	GetSimulateOrderStatusArr(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	GetFuturePosition(context.Context, *emptypb.Empty) (*FuturePositionArr, error)
	GetStockPosition(context.Context, *emptypb.Empty) (*StockPositionArr, error)
	GetSettlement(context.Context, *emptypb.Empty) (*SettlementList, error)
	GetAccountBalance(context.Context, *emptypb.Empty) (*AccountBalance, error)
	GetMargin(context.Context, *emptypb.Empty) (*Margin, error)
	mustEmbedUnimplementedTradeInterfaceServer()
}

// UnimplementedTradeInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedTradeInterfaceServer struct {
}

func (UnimplementedTradeInterfaceServer) BuyStock(context.Context, *StockOrderDetail) (*TradeResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuyStock not implemented")
}
func (UnimplementedTradeInterfaceServer) SellStock(context.Context, *StockOrderDetail) (*TradeResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SellStock not implemented")
}
func (UnimplementedTradeInterfaceServer) BuyOddStock(context.Context, *OddStockOrderDetail) (*TradeResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuyOddStock not implemented")
}
func (UnimplementedTradeInterfaceServer) SellOddStock(context.Context, *OddStockOrderDetail) (*TradeResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SellOddStock not implemented")
}
func (UnimplementedTradeInterfaceServer) SellFirstStock(context.Context, *StockOrderDetail) (*TradeResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SellFirstStock not implemented")
}
func (UnimplementedTradeInterfaceServer) CancelStock(context.Context, *OrderID) (*TradeResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelStock not implemented")
}
func (UnimplementedTradeInterfaceServer) BuyFuture(context.Context, *FutureOrderDetail) (*TradeResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuyFuture not implemented")
}
func (UnimplementedTradeInterfaceServer) SellFuture(context.Context, *FutureOrderDetail) (*TradeResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SellFuture not implemented")
}
func (UnimplementedTradeInterfaceServer) SellFirstFuture(context.Context, *FutureOrderDetail) (*TradeResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SellFirstFuture not implemented")
}
func (UnimplementedTradeInterfaceServer) CancelFuture(context.Context, *FutureOrderID) (*TradeResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelFuture not implemented")
}
func (UnimplementedTradeInterfaceServer) BuyOption(context.Context, *OptionOrderDetail) (*TradeResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuyOption not implemented")
}
func (UnimplementedTradeInterfaceServer) SellOption(context.Context, *OptionOrderDetail) (*TradeResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SellOption not implemented")
}
func (UnimplementedTradeInterfaceServer) SellFirstOption(context.Context, *OptionOrderDetail) (*TradeResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SellFirstOption not implemented")
}
func (UnimplementedTradeInterfaceServer) CancelOption(context.Context, *OptionOrderID) (*TradeResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelOption not implemented")
}
func (UnimplementedTradeInterfaceServer) GetLocalOrderStatusArr(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocalOrderStatusArr not implemented")
}
func (UnimplementedTradeInterfaceServer) GetSimulateOrderStatusArr(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSimulateOrderStatusArr not implemented")
}
func (UnimplementedTradeInterfaceServer) GetFuturePosition(context.Context, *emptypb.Empty) (*FuturePositionArr, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFuturePosition not implemented")
}
func (UnimplementedTradeInterfaceServer) GetStockPosition(context.Context, *emptypb.Empty) (*StockPositionArr, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStockPosition not implemented")
}
func (UnimplementedTradeInterfaceServer) GetSettlement(context.Context, *emptypb.Empty) (*SettlementList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSettlement not implemented")
}
func (UnimplementedTradeInterfaceServer) GetAccountBalance(context.Context, *emptypb.Empty) (*AccountBalance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountBalance not implemented")
}
func (UnimplementedTradeInterfaceServer) GetMargin(context.Context, *emptypb.Empty) (*Margin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMargin not implemented")
}
func (UnimplementedTradeInterfaceServer) mustEmbedUnimplementedTradeInterfaceServer() {}

// UnsafeTradeInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TradeInterfaceServer will
// result in compilation errors.
type UnsafeTradeInterfaceServer interface {
	mustEmbedUnimplementedTradeInterfaceServer()
}

func RegisterTradeInterfaceServer(s grpc.ServiceRegistrar, srv TradeInterfaceServer) {
	s.RegisterService(&TradeInterface_ServiceDesc, srv)
}

func _TradeInterface_BuyStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockOrderDetail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeInterfaceServer).BuyStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeInterface_BuyStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeInterfaceServer).BuyStock(ctx, req.(*StockOrderDetail))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeInterface_SellStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockOrderDetail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeInterfaceServer).SellStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeInterface_SellStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeInterfaceServer).SellStock(ctx, req.(*StockOrderDetail))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeInterface_BuyOddStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OddStockOrderDetail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeInterfaceServer).BuyOddStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeInterface_BuyOddStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeInterfaceServer).BuyOddStock(ctx, req.(*OddStockOrderDetail))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeInterface_SellOddStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OddStockOrderDetail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeInterfaceServer).SellOddStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeInterface_SellOddStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeInterfaceServer).SellOddStock(ctx, req.(*OddStockOrderDetail))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeInterface_SellFirstStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockOrderDetail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeInterfaceServer).SellFirstStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeInterface_SellFirstStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeInterfaceServer).SellFirstStock(ctx, req.(*StockOrderDetail))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeInterface_CancelStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeInterfaceServer).CancelStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeInterface_CancelStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeInterfaceServer).CancelStock(ctx, req.(*OrderID))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeInterface_BuyFuture_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FutureOrderDetail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeInterfaceServer).BuyFuture(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeInterface_BuyFuture_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeInterfaceServer).BuyFuture(ctx, req.(*FutureOrderDetail))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeInterface_SellFuture_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FutureOrderDetail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeInterfaceServer).SellFuture(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeInterface_SellFuture_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeInterfaceServer).SellFuture(ctx, req.(*FutureOrderDetail))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeInterface_SellFirstFuture_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FutureOrderDetail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeInterfaceServer).SellFirstFuture(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeInterface_SellFirstFuture_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeInterfaceServer).SellFirstFuture(ctx, req.(*FutureOrderDetail))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeInterface_CancelFuture_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FutureOrderID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeInterfaceServer).CancelFuture(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeInterface_CancelFuture_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeInterfaceServer).CancelFuture(ctx, req.(*FutureOrderID))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeInterface_BuyOption_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OptionOrderDetail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeInterfaceServer).BuyOption(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeInterface_BuyOption_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeInterfaceServer).BuyOption(ctx, req.(*OptionOrderDetail))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeInterface_SellOption_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OptionOrderDetail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeInterfaceServer).SellOption(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeInterface_SellOption_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeInterfaceServer).SellOption(ctx, req.(*OptionOrderDetail))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeInterface_SellFirstOption_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OptionOrderDetail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeInterfaceServer).SellFirstOption(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeInterface_SellFirstOption_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeInterfaceServer).SellFirstOption(ctx, req.(*OptionOrderDetail))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeInterface_CancelOption_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OptionOrderID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeInterfaceServer).CancelOption(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeInterface_CancelOption_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeInterfaceServer).CancelOption(ctx, req.(*OptionOrderID))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeInterface_GetLocalOrderStatusArr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeInterfaceServer).GetLocalOrderStatusArr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeInterface_GetLocalOrderStatusArr_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeInterfaceServer).GetLocalOrderStatusArr(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeInterface_GetSimulateOrderStatusArr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeInterfaceServer).GetSimulateOrderStatusArr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeInterface_GetSimulateOrderStatusArr_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeInterfaceServer).GetSimulateOrderStatusArr(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeInterface_GetFuturePosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeInterfaceServer).GetFuturePosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeInterface_GetFuturePosition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeInterfaceServer).GetFuturePosition(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeInterface_GetStockPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeInterfaceServer).GetStockPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeInterface_GetStockPosition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeInterfaceServer).GetStockPosition(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeInterface_GetSettlement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeInterfaceServer).GetSettlement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeInterface_GetSettlement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeInterfaceServer).GetSettlement(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeInterface_GetAccountBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeInterfaceServer).GetAccountBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeInterface_GetAccountBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeInterfaceServer).GetAccountBalance(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeInterface_GetMargin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeInterfaceServer).GetMargin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TradeInterface_GetMargin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeInterfaceServer).GetMargin(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// TradeInterface_ServiceDesc is the grpc.ServiceDesc for TradeInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TradeInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "forwarder.TradeInterface",
	HandlerType: (*TradeInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BuyStock",
			Handler:    _TradeInterface_BuyStock_Handler,
		},
		{
			MethodName: "SellStock",
			Handler:    _TradeInterface_SellStock_Handler,
		},
		{
			MethodName: "BuyOddStock",
			Handler:    _TradeInterface_BuyOddStock_Handler,
		},
		{
			MethodName: "SellOddStock",
			Handler:    _TradeInterface_SellOddStock_Handler,
		},
		{
			MethodName: "SellFirstStock",
			Handler:    _TradeInterface_SellFirstStock_Handler,
		},
		{
			MethodName: "CancelStock",
			Handler:    _TradeInterface_CancelStock_Handler,
		},
		{
			MethodName: "BuyFuture",
			Handler:    _TradeInterface_BuyFuture_Handler,
		},
		{
			MethodName: "SellFuture",
			Handler:    _TradeInterface_SellFuture_Handler,
		},
		{
			MethodName: "SellFirstFuture",
			Handler:    _TradeInterface_SellFirstFuture_Handler,
		},
		{
			MethodName: "CancelFuture",
			Handler:    _TradeInterface_CancelFuture_Handler,
		},
		{
			MethodName: "BuyOption",
			Handler:    _TradeInterface_BuyOption_Handler,
		},
		{
			MethodName: "SellOption",
			Handler:    _TradeInterface_SellOption_Handler,
		},
		{
			MethodName: "SellFirstOption",
			Handler:    _TradeInterface_SellFirstOption_Handler,
		},
		{
			MethodName: "CancelOption",
			Handler:    _TradeInterface_CancelOption_Handler,
		},
		{
			MethodName: "GetLocalOrderStatusArr",
			Handler:    _TradeInterface_GetLocalOrderStatusArr_Handler,
		},
		{
			MethodName: "GetSimulateOrderStatusArr",
			Handler:    _TradeInterface_GetSimulateOrderStatusArr_Handler,
		},
		{
			MethodName: "GetFuturePosition",
			Handler:    _TradeInterface_GetFuturePosition_Handler,
		},
		{
			MethodName: "GetStockPosition",
			Handler:    _TradeInterface_GetStockPosition_Handler,
		},
		{
			MethodName: "GetSettlement",
			Handler:    _TradeInterface_GetSettlement_Handler,
		},
		{
			MethodName: "GetAccountBalance",
			Handler:    _TradeInterface_GetAccountBalance_Handler,
		},
		{
			MethodName: "GetMargin",
			Handler:    _TradeInterface_GetMargin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "forwarder/trade.proto",
}
