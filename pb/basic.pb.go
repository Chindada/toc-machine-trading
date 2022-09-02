// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: pb/basic.proto

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

type StockDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stock []*StockDetailMessage `protobuf:"bytes,1,rep,name=stock,proto3" json:"stock,omitempty"`
}

func (x *StockDetailResponse) Reset() {
	*x = StockDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_basic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockDetailResponse) ProtoMessage() {}

func (x *StockDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_basic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockDetailResponse.ProtoReflect.Descriptor instead.
func (*StockDetailResponse) Descriptor() ([]byte, []int) {
	return file_pb_basic_proto_rawDescGZIP(), []int{0}
}

func (x *StockDetailResponse) GetStock() []*StockDetailMessage {
	if x != nil {
		return x.Stock
	}
	return nil
}

type StockDetailMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exchange   string  `protobuf:"bytes,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Category   string  `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Code       string  `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Name       string  `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Reference  float64 `protobuf:"fixed64,5,opt,name=reference,proto3" json:"reference,omitempty"`
	UpdateDate string  `protobuf:"bytes,6,opt,name=update_date,json=updateDate,proto3" json:"update_date,omitempty"`
	DayTrade   string  `protobuf:"bytes,7,opt,name=day_trade,json=dayTrade,proto3" json:"day_trade,omitempty"`
}

func (x *StockDetailMessage) Reset() {
	*x = StockDetailMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_basic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockDetailMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockDetailMessage) ProtoMessage() {}

func (x *StockDetailMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pb_basic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockDetailMessage.ProtoReflect.Descriptor instead.
func (*StockDetailMessage) Descriptor() ([]byte, []int) {
	return file_pb_basic_proto_rawDescGZIP(), []int{1}
}

func (x *StockDetailMessage) GetExchange() string {
	if x != nil {
		return x.Exchange
	}
	return ""
}

func (x *StockDetailMessage) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *StockDetailMessage) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *StockDetailMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StockDetailMessage) GetReference() float64 {
	if x != nil {
		return x.Reference
	}
	return 0
}

func (x *StockDetailMessage) GetUpdateDate() string {
	if x != nil {
		return x.UpdateDate
	}
	return ""
}

func (x *StockDetailMessage) GetDayTrade() string {
	if x != nil {
		return x.DayTrade
	}
	return ""
}

type FutureDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Future []*FutureDetailMessage `protobuf:"bytes,1,rep,name=future,proto3" json:"future,omitempty"`
}

func (x *FutureDetailResponse) Reset() {
	*x = FutureDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_basic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FutureDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FutureDetailResponse) ProtoMessage() {}

func (x *FutureDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_basic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FutureDetailResponse.ProtoReflect.Descriptor instead.
func (*FutureDetailResponse) Descriptor() ([]byte, []int) {
	return file_pb_basic_proto_rawDescGZIP(), []int{2}
}

func (x *FutureDetailResponse) GetFuture() []*FutureDetailMessage {
	if x != nil {
		return x.Future
	}
	return nil
}

type FutureDetailMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code           string  `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Symbol         string  `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Name           string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Category       string  `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
	DeliveryMonth  string  `protobuf:"bytes,5,opt,name=delivery_month,json=deliveryMonth,proto3" json:"delivery_month,omitempty"`
	DeliveryDate   string  `protobuf:"bytes,6,opt,name=delivery_date,json=deliveryDate,proto3" json:"delivery_date,omitempty"`
	UnderlyingKind string  `protobuf:"bytes,7,opt,name=underlying_kind,json=underlyingKind,proto3" json:"underlying_kind,omitempty"`
	Unit           int64   `protobuf:"varint,8,opt,name=unit,proto3" json:"unit,omitempty"`
	LimitUp        float64 `protobuf:"fixed64,9,opt,name=limit_up,json=limitUp,proto3" json:"limit_up,omitempty"`
	LimitDown      float64 `protobuf:"fixed64,10,opt,name=limit_down,json=limitDown,proto3" json:"limit_down,omitempty"`
	Reference      float64 `protobuf:"fixed64,11,opt,name=reference,proto3" json:"reference,omitempty"`
	UpdateDate     string  `protobuf:"bytes,12,opt,name=update_date,json=updateDate,proto3" json:"update_date,omitempty"`
}

func (x *FutureDetailMessage) Reset() {
	*x = FutureDetailMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_basic_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FutureDetailMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FutureDetailMessage) ProtoMessage() {}

func (x *FutureDetailMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pb_basic_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FutureDetailMessage.ProtoReflect.Descriptor instead.
func (*FutureDetailMessage) Descriptor() ([]byte, []int) {
	return file_pb_basic_proto_rawDescGZIP(), []int{3}
}

func (x *FutureDetailMessage) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *FutureDetailMessage) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *FutureDetailMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FutureDetailMessage) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *FutureDetailMessage) GetDeliveryMonth() string {
	if x != nil {
		return x.DeliveryMonth
	}
	return ""
}

func (x *FutureDetailMessage) GetDeliveryDate() string {
	if x != nil {
		return x.DeliveryDate
	}
	return ""
}

func (x *FutureDetailMessage) GetUnderlyingKind() string {
	if x != nil {
		return x.UnderlyingKind
	}
	return ""
}

func (x *FutureDetailMessage) GetUnit() int64 {
	if x != nil {
		return x.Unit
	}
	return 0
}

func (x *FutureDetailMessage) GetLimitUp() float64 {
	if x != nil {
		return x.LimitUp
	}
	return 0
}

func (x *FutureDetailMessage) GetLimitDown() float64 {
	if x != nil {
		return x.LimitDown
	}
	return 0
}

func (x *FutureDetailMessage) GetReference() float64 {
	if x != nil {
		return x.Reference
	}
	return 0
}

func (x *FutureDetailMessage) GetUpdateDate() string {
	if x != nil {
		return x.UpdateDate
	}
	return ""
}

var File_pb_basic_proto protoreflect.FileDescriptor

var file_pb_basic_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x62, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x52, 0x0a, 0x13, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63,
	0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x73,
	0x74, 0x6f, 0x63, 0x6b, 0x22, 0xd0, 0x01, 0x0a, 0x12, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65,
	0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72,
	0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61,
	0x79, 0x5f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x61, 0x79, 0x54, 0x72, 0x61, 0x64, 0x65, 0x22, 0x56, 0x0a, 0x14, 0x46, 0x75, 0x74, 0x75, 0x72,
	0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3e, 0x0a, 0x06, 0x66, 0x75, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x06, 0x66, 0x75, 0x74, 0x75, 0x72, 0x65, 0x22,
	0xf3, 0x02, 0x0a, 0x13, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f,
	0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x27, 0x0a, 0x0f, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x6c, 0x79, 0x69, 0x6e, 0x67, 0x5f, 0x6b, 0x69,
	0x6e, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x75, 0x6e, 0x64, 0x65, 0x72, 0x6c,
	0x79, 0x69, 0x6e, 0x67, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x75, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x55, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x61, 0x74, 0x65, 0x32, 0xc4, 0x01, 0x0a, 0x12, 0x42, 0x61, 0x73, 0x69, 0x63, 0x44,
	0x61, 0x74, 0x61, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x26, 0x2e, 0x73, 0x69, 0x6e, 0x6f,
	0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46, 0x75, 0x74,
	0x75, 0x72, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x27, 0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04,
	0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_basic_proto_rawDescOnce sync.Once
	file_pb_basic_proto_rawDescData = file_pb_basic_proto_rawDesc
)

func file_pb_basic_proto_rawDescGZIP() []byte {
	file_pb_basic_proto_rawDescOnce.Do(func() {
		file_pb_basic_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_basic_proto_rawDescData)
	})
	return file_pb_basic_proto_rawDescData
}

var file_pb_basic_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_basic_proto_goTypes = []interface{}{
	(*StockDetailResponse)(nil),  // 0: sinopac_forwarder.StockDetailResponse
	(*StockDetailMessage)(nil),   // 1: sinopac_forwarder.StockDetailMessage
	(*FutureDetailResponse)(nil), // 2: sinopac_forwarder.FutureDetailResponse
	(*FutureDetailMessage)(nil),  // 3: sinopac_forwarder.FutureDetailMessage
	(*emptypb.Empty)(nil),        // 4: google.protobuf.Empty
}
var file_pb_basic_proto_depIdxs = []int32{
	1, // 0: sinopac_forwarder.StockDetailResponse.stock:type_name -> sinopac_forwarder.StockDetailMessage
	3, // 1: sinopac_forwarder.FutureDetailResponse.future:type_name -> sinopac_forwarder.FutureDetailMessage
	4, // 2: sinopac_forwarder.BasicDataInterface.GetAllStockDetail:input_type -> google.protobuf.Empty
	4, // 3: sinopac_forwarder.BasicDataInterface.GetAllFutureDetail:input_type -> google.protobuf.Empty
	0, // 4: sinopac_forwarder.BasicDataInterface.GetAllStockDetail:output_type -> sinopac_forwarder.StockDetailResponse
	2, // 5: sinopac_forwarder.BasicDataInterface.GetAllFutureDetail:output_type -> sinopac_forwarder.FutureDetailResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pb_basic_proto_init() }
func file_pb_basic_proto_init() {
	if File_pb_basic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_basic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockDetailResponse); i {
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
		file_pb_basic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockDetailMessage); i {
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
		file_pb_basic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FutureDetailResponse); i {
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
		file_pb_basic_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FutureDetailMessage); i {
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
			RawDescriptor: file_pb_basic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_basic_proto_goTypes,
		DependencyIndexes: file_pb_basic_proto_depIdxs,
		MessageInfos:      file_pb_basic_proto_msgTypes,
	}.Build()
	File_pb_basic_proto = out.File
	file_pb_basic_proto_rawDesc = nil
	file_pb_basic_proto_goTypes = nil
	file_pb_basic_proto_depIdxs = nil
}
