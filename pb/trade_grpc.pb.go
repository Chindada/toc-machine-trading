// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: trade.proto

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

// TradeInterfaceClient is the client API for TradeInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TradeInterfaceClient interface {
	BuyStock(ctx context.Context, in *StockOrderDetail, opts ...grpc.CallOption) (*TradeResult, error)
	SellStock(ctx context.Context, in *StockOrderDetail, opts ...grpc.CallOption) (*TradeResult, error)
	SellFirstStock(ctx context.Context, in *StockOrderDetail, opts ...grpc.CallOption) (*TradeResult, error)
	CancelStock(ctx context.Context, in *OrderID, opts ...grpc.CallOption) (*TradeResult, error)
	GetOrderStatusByID(ctx context.Context, in *OrderID, opts ...grpc.CallOption) (*TradeResult, error)
	GetOrderStatusArr(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StockOrderStatusArr, error)
	GetNonBlockOrderStatusArr(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ErrorMessage, error)
	BuyFuture(ctx context.Context, in *FutureOrderDetail, opts ...grpc.CallOption) (*TradeResult, error)
	SellFuture(ctx context.Context, in *FutureOrderDetail, opts ...grpc.CallOption) (*TradeResult, error)
	SellFirstFuture(ctx context.Context, in *FutureOrderDetail, opts ...grpc.CallOption) (*TradeResult, error)
	CancelFuture(ctx context.Context, in *FutureOrderID, opts ...grpc.CallOption) (*TradeResult, error)
}

type tradeInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewTradeInterfaceClient(cc grpc.ClientConnInterface) TradeInterfaceClient {
	return &tradeInterfaceClient{cc}
}

func (c *tradeInterfaceClient) BuyStock(ctx context.Context, in *StockOrderDetail, opts ...grpc.CallOption) (*TradeResult, error) {
	out := new(TradeResult)
	err := c.cc.Invoke(ctx, "/sinopac_forwarder.TradeInterface/BuyStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeInterfaceClient) SellStock(ctx context.Context, in *StockOrderDetail, opts ...grpc.CallOption) (*TradeResult, error) {
	out := new(TradeResult)
	err := c.cc.Invoke(ctx, "/sinopac_forwarder.TradeInterface/SellStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeInterfaceClient) SellFirstStock(ctx context.Context, in *StockOrderDetail, opts ...grpc.CallOption) (*TradeResult, error) {
	out := new(TradeResult)
	err := c.cc.Invoke(ctx, "/sinopac_forwarder.TradeInterface/SellFirstStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeInterfaceClient) CancelStock(ctx context.Context, in *OrderID, opts ...grpc.CallOption) (*TradeResult, error) {
	out := new(TradeResult)
	err := c.cc.Invoke(ctx, "/sinopac_forwarder.TradeInterface/CancelStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeInterfaceClient) GetOrderStatusByID(ctx context.Context, in *OrderID, opts ...grpc.CallOption) (*TradeResult, error) {
	out := new(TradeResult)
	err := c.cc.Invoke(ctx, "/sinopac_forwarder.TradeInterface/GetOrderStatusByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeInterfaceClient) GetOrderStatusArr(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StockOrderStatusArr, error) {
	out := new(StockOrderStatusArr)
	err := c.cc.Invoke(ctx, "/sinopac_forwarder.TradeInterface/GetOrderStatusArr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeInterfaceClient) GetNonBlockOrderStatusArr(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ErrorMessage, error) {
	out := new(ErrorMessage)
	err := c.cc.Invoke(ctx, "/sinopac_forwarder.TradeInterface/GetNonBlockOrderStatusArr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeInterfaceClient) BuyFuture(ctx context.Context, in *FutureOrderDetail, opts ...grpc.CallOption) (*TradeResult, error) {
	out := new(TradeResult)
	err := c.cc.Invoke(ctx, "/sinopac_forwarder.TradeInterface/BuyFuture", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeInterfaceClient) SellFuture(ctx context.Context, in *FutureOrderDetail, opts ...grpc.CallOption) (*TradeResult, error) {
	out := new(TradeResult)
	err := c.cc.Invoke(ctx, "/sinopac_forwarder.TradeInterface/SellFuture", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeInterfaceClient) SellFirstFuture(ctx context.Context, in *FutureOrderDetail, opts ...grpc.CallOption) (*TradeResult, error) {
	out := new(TradeResult)
	err := c.cc.Invoke(ctx, "/sinopac_forwarder.TradeInterface/SellFirstFuture", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tradeInterfaceClient) CancelFuture(ctx context.Context, in *FutureOrderID, opts ...grpc.CallOption) (*TradeResult, error) {
	out := new(TradeResult)
	err := c.cc.Invoke(ctx, "/sinopac_forwarder.TradeInterface/CancelFuture", in, out, opts...)
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
	SellFirstStock(context.Context, *StockOrderDetail) (*TradeResult, error)
	CancelStock(context.Context, *OrderID) (*TradeResult, error)
	GetOrderStatusByID(context.Context, *OrderID) (*TradeResult, error)
	GetOrderStatusArr(context.Context, *emptypb.Empty) (*StockOrderStatusArr, error)
	GetNonBlockOrderStatusArr(context.Context, *emptypb.Empty) (*ErrorMessage, error)
	BuyFuture(context.Context, *FutureOrderDetail) (*TradeResult, error)
	SellFuture(context.Context, *FutureOrderDetail) (*TradeResult, error)
	SellFirstFuture(context.Context, *FutureOrderDetail) (*TradeResult, error)
	CancelFuture(context.Context, *FutureOrderID) (*TradeResult, error)
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
func (UnimplementedTradeInterfaceServer) SellFirstStock(context.Context, *StockOrderDetail) (*TradeResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SellFirstStock not implemented")
}
func (UnimplementedTradeInterfaceServer) CancelStock(context.Context, *OrderID) (*TradeResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelStock not implemented")
}
func (UnimplementedTradeInterfaceServer) GetOrderStatusByID(context.Context, *OrderID) (*TradeResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderStatusByID not implemented")
}
func (UnimplementedTradeInterfaceServer) GetOrderStatusArr(context.Context, *emptypb.Empty) (*StockOrderStatusArr, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderStatusArr not implemented")
}
func (UnimplementedTradeInterfaceServer) GetNonBlockOrderStatusArr(context.Context, *emptypb.Empty) (*ErrorMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNonBlockOrderStatusArr not implemented")
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
		FullMethod: "/sinopac_forwarder.TradeInterface/BuyStock",
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
		FullMethod: "/sinopac_forwarder.TradeInterface/SellStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeInterfaceServer).SellStock(ctx, req.(*StockOrderDetail))
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
		FullMethod: "/sinopac_forwarder.TradeInterface/SellFirstStock",
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
		FullMethod: "/sinopac_forwarder.TradeInterface/CancelStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeInterfaceServer).CancelStock(ctx, req.(*OrderID))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeInterface_GetOrderStatusByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeInterfaceServer).GetOrderStatusByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sinopac_forwarder.TradeInterface/GetOrderStatusByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeInterfaceServer).GetOrderStatusByID(ctx, req.(*OrderID))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeInterface_GetOrderStatusArr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeInterfaceServer).GetOrderStatusArr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sinopac_forwarder.TradeInterface/GetOrderStatusArr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeInterfaceServer).GetOrderStatusArr(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TradeInterface_GetNonBlockOrderStatusArr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TradeInterfaceServer).GetNonBlockOrderStatusArr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sinopac_forwarder.TradeInterface/GetNonBlockOrderStatusArr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeInterfaceServer).GetNonBlockOrderStatusArr(ctx, req.(*emptypb.Empty))
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
		FullMethod: "/sinopac_forwarder.TradeInterface/BuyFuture",
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
		FullMethod: "/sinopac_forwarder.TradeInterface/SellFuture",
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
		FullMethod: "/sinopac_forwarder.TradeInterface/SellFirstFuture",
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
		FullMethod: "/sinopac_forwarder.TradeInterface/CancelFuture",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TradeInterfaceServer).CancelFuture(ctx, req.(*FutureOrderID))
	}
	return interceptor(ctx, in, info, handler)
}

// TradeInterface_ServiceDesc is the grpc.ServiceDesc for TradeInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TradeInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sinopac_forwarder.TradeInterface",
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
			MethodName: "SellFirstStock",
			Handler:    _TradeInterface_SellFirstStock_Handler,
		},
		{
			MethodName: "CancelStock",
			Handler:    _TradeInterface_CancelStock_Handler,
		},
		{
			MethodName: "GetOrderStatusByID",
			Handler:    _TradeInterface_GetOrderStatusByID_Handler,
		},
		{
			MethodName: "GetOrderStatusArr",
			Handler:    _TradeInterface_GetOrderStatusArr_Handler,
		},
		{
			MethodName: "GetNonBlockOrderStatusArr",
			Handler:    _TradeInterface_GetNonBlockOrderStatusArr_Handler,
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "trade.proto",
}
