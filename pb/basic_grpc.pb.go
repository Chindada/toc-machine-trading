// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: basic.proto

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

// BasicDataInterfaceClient is the client API for BasicDataInterface service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BasicDataInterfaceClient interface {
	GetAllStockDetail(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StockDetailResponse, error)
	GetAllFutureDetail(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FutureDetailResponse, error)
}

type basicDataInterfaceClient struct {
	cc grpc.ClientConnInterface
}

func NewBasicDataInterfaceClient(cc grpc.ClientConnInterface) BasicDataInterfaceClient {
	return &basicDataInterfaceClient{cc}
}

func (c *basicDataInterfaceClient) GetAllStockDetail(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*StockDetailResponse, error) {
	out := new(StockDetailResponse)
	err := c.cc.Invoke(ctx, "/sinopac_forwarder.BasicDataInterface/GetAllStockDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *basicDataInterfaceClient) GetAllFutureDetail(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FutureDetailResponse, error) {
	out := new(FutureDetailResponse)
	err := c.cc.Invoke(ctx, "/sinopac_forwarder.BasicDataInterface/GetAllFutureDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BasicDataInterfaceServer is the server API for BasicDataInterface service.
// All implementations must embed UnimplementedBasicDataInterfaceServer
// for forward compatibility
type BasicDataInterfaceServer interface {
	GetAllStockDetail(context.Context, *emptypb.Empty) (*StockDetailResponse, error)
	GetAllFutureDetail(context.Context, *emptypb.Empty) (*FutureDetailResponse, error)
	mustEmbedUnimplementedBasicDataInterfaceServer()
}

// UnimplementedBasicDataInterfaceServer must be embedded to have forward compatible implementations.
type UnimplementedBasicDataInterfaceServer struct {
}

func (UnimplementedBasicDataInterfaceServer) GetAllStockDetail(context.Context, *emptypb.Empty) (*StockDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllStockDetail not implemented")
}
func (UnimplementedBasicDataInterfaceServer) GetAllFutureDetail(context.Context, *emptypb.Empty) (*FutureDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllFutureDetail not implemented")
}
func (UnimplementedBasicDataInterfaceServer) mustEmbedUnimplementedBasicDataInterfaceServer() {}

// UnsafeBasicDataInterfaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BasicDataInterfaceServer will
// result in compilation errors.
type UnsafeBasicDataInterfaceServer interface {
	mustEmbedUnimplementedBasicDataInterfaceServer()
}

func RegisterBasicDataInterfaceServer(s grpc.ServiceRegistrar, srv BasicDataInterfaceServer) {
	s.RegisterService(&BasicDataInterface_ServiceDesc, srv)
}

func _BasicDataInterface_GetAllStockDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasicDataInterfaceServer).GetAllStockDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sinopac_forwarder.BasicDataInterface/GetAllStockDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasicDataInterfaceServer).GetAllStockDetail(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BasicDataInterface_GetAllFutureDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BasicDataInterfaceServer).GetAllFutureDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sinopac_forwarder.BasicDataInterface/GetAllFutureDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BasicDataInterfaceServer).GetAllFutureDetail(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// BasicDataInterface_ServiceDesc is the grpc.ServiceDesc for BasicDataInterface service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BasicDataInterface_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sinopac_forwarder.BasicDataInterface",
	HandlerType: (*BasicDataInterfaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllStockDetail",
			Handler:    _BasicDataInterface_GetAllStockDetail_Handler,
		},
		{
			MethodName: "GetAllFutureDetail",
			Handler:    _BasicDataInterface_GetAllFutureDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "basic.proto",
}
