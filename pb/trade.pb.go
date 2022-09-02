// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: pb/trade.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StockOrderDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StockNum string  `protobuf:"bytes,1,opt,name=stock_num,json=stockNum,proto3" json:"stock_num,omitempty"`
	Price    float64 `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
	Quantity int64   `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Simulate bool    `protobuf:"varint,4,opt,name=simulate,proto3" json:"simulate,omitempty"`
}

func (x *StockOrderDetail) Reset() {
	*x = StockOrderDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_trade_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockOrderDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockOrderDetail) ProtoMessage() {}

func (x *StockOrderDetail) ProtoReflect() protoreflect.Message {
	mi := &file_pb_trade_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockOrderDetail.ProtoReflect.Descriptor instead.
func (*StockOrderDetail) Descriptor() ([]byte, []int) {
	return file_pb_trade_proto_rawDescGZIP(), []int{0}
}

func (x *StockOrderDetail) GetStockNum() string {
	if x != nil {
		return x.StockNum
	}
	return ""
}

func (x *StockOrderDetail) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *StockOrderDetail) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *StockOrderDetail) GetSimulate() bool {
	if x != nil {
		return x.Simulate
	}
	return false
}

type FutureOrderDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     string  `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Price    float64 `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
	Quantity int64   `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *FutureOrderDetail) Reset() {
	*x = FutureOrderDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_trade_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FutureOrderDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FutureOrderDetail) ProtoMessage() {}

func (x *FutureOrderDetail) ProtoReflect() protoreflect.Message {
	mi := &file_pb_trade_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FutureOrderDetail.ProtoReflect.Descriptor instead.
func (*FutureOrderDetail) Descriptor() ([]byte, []int) {
	return file_pb_trade_proto_rawDescGZIP(), []int{1}
}

func (x *FutureOrderDetail) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *FutureOrderDetail) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *FutureOrderDetail) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type TradeResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Status  string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Error   string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *TradeResult) Reset() {
	*x = TradeResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_trade_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TradeResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TradeResult) ProtoMessage() {}

func (x *TradeResult) ProtoReflect() protoreflect.Message {
	mi := &file_pb_trade_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TradeResult.ProtoReflect.Descriptor instead.
func (*TradeResult) Descriptor() ([]byte, []int) {
	return file_pb_trade_proto_rawDescGZIP(), []int{2}
}

func (x *TradeResult) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *TradeResult) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *TradeResult) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type OrderID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId  string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Simulate bool   `protobuf:"varint,2,opt,name=simulate,proto3" json:"simulate,omitempty"`
}

func (x *OrderID) Reset() {
	*x = OrderID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_trade_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderID) ProtoMessage() {}

func (x *OrderID) ProtoReflect() protoreflect.Message {
	mi := &file_pb_trade_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderID.ProtoReflect.Descriptor instead.
func (*OrderID) Descriptor() ([]byte, []int) {
	return file_pb_trade_proto_rawDescGZIP(), []int{3}
}

func (x *OrderID) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *OrderID) GetSimulate() bool {
	if x != nil {
		return x.Simulate
	}
	return false
}

type FutureOrderID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *FutureOrderID) Reset() {
	*x = FutureOrderID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_trade_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FutureOrderID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FutureOrderID) ProtoMessage() {}

func (x *FutureOrderID) ProtoReflect() protoreflect.Message {
	mi := &file_pb_trade_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FutureOrderID.ProtoReflect.Descriptor instead.
func (*FutureOrderID) Descriptor() ([]byte, []int) {
	return file_pb_trade_proto_rawDescGZIP(), []int{4}
}

func (x *FutureOrderID) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type StockOrderStatusArr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*StockOrderStatus `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *StockOrderStatusArr) Reset() {
	*x = StockOrderStatusArr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_trade_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockOrderStatusArr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockOrderStatusArr) ProtoMessage() {}

func (x *StockOrderStatusArr) ProtoReflect() protoreflect.Message {
	mi := &file_pb_trade_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockOrderStatusArr.ProtoReflect.Descriptor instead.
func (*StockOrderStatusArr) Descriptor() ([]byte, []int) {
	return file_pb_trade_proto_rawDescGZIP(), []int{5}
}

func (x *StockOrderStatusArr) GetData() []*StockOrderStatus {
	if x != nil {
		return x.Data
	}
	return nil
}

type StockOrderStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    string  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Code      string  `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Action    string  `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	Price     float64 `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	Quantity  int64   `protobuf:"varint,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
	OrderId   string  `protobuf:"bytes,6,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	OrderTime string  `protobuf:"bytes,7,opt,name=order_time,json=orderTime,proto3" json:"order_time,omitempty"`
}

func (x *StockOrderStatus) Reset() {
	*x = StockOrderStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_trade_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockOrderStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockOrderStatus) ProtoMessage() {}

func (x *StockOrderStatus) ProtoReflect() protoreflect.Message {
	mi := &file_pb_trade_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockOrderStatus.ProtoReflect.Descriptor instead.
func (*StockOrderStatus) Descriptor() ([]byte, []int) {
	return file_pb_trade_proto_rawDescGZIP(), []int{6}
}

func (x *StockOrderStatus) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *StockOrderStatus) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *StockOrderStatus) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *StockOrderStatus) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *StockOrderStatus) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *StockOrderStatus) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *StockOrderStatus) GetOrderTime() string {
	if x != nil {
		return x.OrderTime
	}
	return ""
}

var File_pb_trade_proto protoreflect.FileDescriptor

var file_pb_trade_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x62, 0x2f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0f, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x7d, 0x0a, 0x10, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x6e,
	0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x4e,
	0x75, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x22, 0x59, 0x0a, 0x11, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x56, 0x0a, 0x0b, 0x54,
	0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x22, 0x40, 0x0a, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x69, 0x6d,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x73, 0x69, 0x6d,
	0x75, 0x6c, 0x61, 0x74, 0x65, 0x22, 0x2a, 0x0a, 0x0d, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x4e, 0x0a, 0x13, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x72, 0x72, 0x12, 0x37, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63,
	0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0xc2, 0x01, 0x0a, 0x10, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x19, 0x0a, 0x08,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xba, 0x07, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x64, 0x65,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x08, 0x42, 0x75, 0x79,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x23, 0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f,
	0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x1a, 0x1e, 0x2e, 0x73, 0x69, 0x6e,
	0x6f, 0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x54,
	0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x09,
	0x53, 0x65, 0x6c, 0x6c, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x23, 0x2e, 0x73, 0x69, 0x6e, 0x6f,
	0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x1a, 0x1e,
	0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00,
	0x12, 0x57, 0x0a, 0x0e, 0x53, 0x65, 0x6c, 0x6c, 0x46, 0x69, 0x72, 0x73, 0x74, 0x53, 0x74, 0x6f,
	0x63, 0x6b, 0x12, 0x23, 0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x1a, 0x1e, 0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61,
	0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x64,
	0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0b, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x1a, 0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70,
	0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x44, 0x1a, 0x1e, 0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f, 0x66,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1a, 0x2e, 0x73,
	0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x1e, 0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70,
	0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x72, 0x72, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x26, 0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61,
	0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x72, 0x72, 0x22,
	0x00, 0x12, 0x56, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x41, 0x72, 0x72, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63,
	0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x09, 0x42, 0x75, 0x79,
	0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x12, 0x24, 0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63,
	0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x46, 0x75, 0x74, 0x75, 0x72,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x1a, 0x1e, 0x2e, 0x73,
	0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x54,
	0x0a, 0x0a, 0x53, 0x65, 0x6c, 0x6c, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x12, 0x24, 0x2e, 0x73,
	0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x1a, 0x1e, 0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x0f, 0x53, 0x65, 0x6c, 0x6c, 0x46, 0x69, 0x72, 0x73,
	0x74, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x12, 0x24, 0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61,
	0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x46, 0x75, 0x74, 0x75,
	0x72, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x1a, 0x1e, 0x2e,
	0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12,
	0x52, 0x0a, 0x0c, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x20, 0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x44, 0x1a, 0x1e, 0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pb_trade_proto_rawDescOnce sync.Once
	file_pb_trade_proto_rawDescData = file_pb_trade_proto_rawDesc
)

func file_pb_trade_proto_rawDescGZIP() []byte {
	file_pb_trade_proto_rawDescOnce.Do(func() {
		file_pb_trade_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_trade_proto_rawDescData)
	})
	return file_pb_trade_proto_rawDescData
}

var file_pb_trade_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pb_trade_proto_goTypes = []interface{}{
	(*StockOrderDetail)(nil),    // 0: sinopac_forwarder.StockOrderDetail
	(*FutureOrderDetail)(nil),   // 1: sinopac_forwarder.FutureOrderDetail
	(*TradeResult)(nil),         // 2: sinopac_forwarder.TradeResult
	(*OrderID)(nil),             // 3: sinopac_forwarder.OrderID
	(*FutureOrderID)(nil),       // 4: sinopac_forwarder.FutureOrderID
	(*StockOrderStatusArr)(nil), // 5: sinopac_forwarder.StockOrderStatusArr
	(*StockOrderStatus)(nil),    // 6: sinopac_forwarder.StockOrderStatus
	(*emptypb.Empty)(nil),       // 7: google.protobuf.Empty
	(*ErrorMessage)(nil),        // 8: sinopac_forwarder.ErrorMessage
}
var file_pb_trade_proto_depIdxs = []int32{
	6,  // 0: sinopac_forwarder.StockOrderStatusArr.data:type_name -> sinopac_forwarder.StockOrderStatus
	0,  // 1: sinopac_forwarder.TradeInterface.BuyStock:input_type -> sinopac_forwarder.StockOrderDetail
	0,  // 2: sinopac_forwarder.TradeInterface.SellStock:input_type -> sinopac_forwarder.StockOrderDetail
	0,  // 3: sinopac_forwarder.TradeInterface.SellFirstStock:input_type -> sinopac_forwarder.StockOrderDetail
	3,  // 4: sinopac_forwarder.TradeInterface.CancelStock:input_type -> sinopac_forwarder.OrderID
	3,  // 5: sinopac_forwarder.TradeInterface.GetOrderStatusByID:input_type -> sinopac_forwarder.OrderID
	7,  // 6: sinopac_forwarder.TradeInterface.GetOrderStatusArr:input_type -> google.protobuf.Empty
	7,  // 7: sinopac_forwarder.TradeInterface.GetNonBlockOrderStatusArr:input_type -> google.protobuf.Empty
	1,  // 8: sinopac_forwarder.TradeInterface.BuyFuture:input_type -> sinopac_forwarder.FutureOrderDetail
	1,  // 9: sinopac_forwarder.TradeInterface.SellFuture:input_type -> sinopac_forwarder.FutureOrderDetail
	1,  // 10: sinopac_forwarder.TradeInterface.SellFirstFuture:input_type -> sinopac_forwarder.FutureOrderDetail
	4,  // 11: sinopac_forwarder.TradeInterface.CancelFuture:input_type -> sinopac_forwarder.FutureOrderID
	2,  // 12: sinopac_forwarder.TradeInterface.BuyStock:output_type -> sinopac_forwarder.TradeResult
	2,  // 13: sinopac_forwarder.TradeInterface.SellStock:output_type -> sinopac_forwarder.TradeResult
	2,  // 14: sinopac_forwarder.TradeInterface.SellFirstStock:output_type -> sinopac_forwarder.TradeResult
	2,  // 15: sinopac_forwarder.TradeInterface.CancelStock:output_type -> sinopac_forwarder.TradeResult
	2,  // 16: sinopac_forwarder.TradeInterface.GetOrderStatusByID:output_type -> sinopac_forwarder.TradeResult
	5,  // 17: sinopac_forwarder.TradeInterface.GetOrderStatusArr:output_type -> sinopac_forwarder.StockOrderStatusArr
	8,  // 18: sinopac_forwarder.TradeInterface.GetNonBlockOrderStatusArr:output_type -> sinopac_forwarder.ErrorMessage
	2,  // 19: sinopac_forwarder.TradeInterface.BuyFuture:output_type -> sinopac_forwarder.TradeResult
	2,  // 20: sinopac_forwarder.TradeInterface.SellFuture:output_type -> sinopac_forwarder.TradeResult
	2,  // 21: sinopac_forwarder.TradeInterface.SellFirstFuture:output_type -> sinopac_forwarder.TradeResult
	2,  // 22: sinopac_forwarder.TradeInterface.CancelFuture:output_type -> sinopac_forwarder.TradeResult
	12, // [12:23] is the sub-list for method output_type
	1,  // [1:12] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_pb_trade_proto_init() }
func file_pb_trade_proto_init() {
	if File_pb_trade_proto != nil {
		return
	}
	file_pb_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pb_trade_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockOrderDetail); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_trade_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FutureOrderDetail); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_trade_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TradeResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_trade_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_trade_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FutureOrderID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_trade_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockOrderStatusArr); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pb_trade_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockOrderStatus); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_trade_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_trade_proto_goTypes,
		DependencyIndexes: file_pb_trade_proto_depIdxs,
		MessageInfos:      file_pb_trade_proto_msgTypes,
	}.Build()
	File_pb_trade_proto = out.File
	file_pb_trade_proto_rawDesc = nil
	file_pb_trade_proto_goTypes = nil
	file_pb_trade_proto_depIdxs = nil
}
