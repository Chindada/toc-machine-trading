// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: history.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HistoryTickResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*HistoryTickMessage `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *HistoryTickResponse) Reset() {
	*x = HistoryTickResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_history_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryTickResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryTickResponse) ProtoMessage() {}

func (x *HistoryTickResponse) ProtoReflect() protoreflect.Message {
	mi := &file_history_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryTickResponse.ProtoReflect.Descriptor instead.
func (*HistoryTickResponse) Descriptor() ([]byte, []int) {
	return file_history_proto_rawDescGZIP(), []int{0}
}

func (x *HistoryTickResponse) GetData() []*HistoryTickMessage {
	if x != nil {
		return x.Data
	}
	return nil
}

type HistoryTickMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts        int64   `protobuf:"varint,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Close     float64 `protobuf:"fixed64,2,opt,name=close,proto3" json:"close,omitempty"`
	Volume    int64   `protobuf:"varint,3,opt,name=volume,proto3" json:"volume,omitempty"`
	BidPrice  float64 `protobuf:"fixed64,4,opt,name=bid_price,json=bidPrice,proto3" json:"bid_price,omitempty"`
	BidVolume int64   `protobuf:"varint,5,opt,name=bid_volume,json=bidVolume,proto3" json:"bid_volume,omitempty"`
	AskPrice  float64 `protobuf:"fixed64,6,opt,name=ask_price,json=askPrice,proto3" json:"ask_price,omitempty"`
	AskVolume int64   `protobuf:"varint,7,opt,name=ask_volume,json=askVolume,proto3" json:"ask_volume,omitempty"`
	TickType  int64   `protobuf:"varint,8,opt,name=tick_type,json=tickType,proto3" json:"tick_type,omitempty"`
	Code      string  `protobuf:"bytes,9,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *HistoryTickMessage) Reset() {
	*x = HistoryTickMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_history_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryTickMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryTickMessage) ProtoMessage() {}

func (x *HistoryTickMessage) ProtoReflect() protoreflect.Message {
	mi := &file_history_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryTickMessage.ProtoReflect.Descriptor instead.
func (*HistoryTickMessage) Descriptor() ([]byte, []int) {
	return file_history_proto_rawDescGZIP(), []int{1}
}

func (x *HistoryTickMessage) GetTs() int64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

func (x *HistoryTickMessage) GetClose() float64 {
	if x != nil {
		return x.Close
	}
	return 0
}

func (x *HistoryTickMessage) GetVolume() int64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *HistoryTickMessage) GetBidPrice() float64 {
	if x != nil {
		return x.BidPrice
	}
	return 0
}

func (x *HistoryTickMessage) GetBidVolume() int64 {
	if x != nil {
		return x.BidVolume
	}
	return 0
}

func (x *HistoryTickMessage) GetAskPrice() float64 {
	if x != nil {
		return x.AskPrice
	}
	return 0
}

func (x *HistoryTickMessage) GetAskVolume() int64 {
	if x != nil {
		return x.AskVolume
	}
	return 0
}

func (x *HistoryTickMessage) GetTickType() int64 {
	if x != nil {
		return x.TickType
	}
	return 0
}

func (x *HistoryTickMessage) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type HistoryKbarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*HistoryKbarMessage `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *HistoryKbarResponse) Reset() {
	*x = HistoryKbarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_history_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryKbarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryKbarResponse) ProtoMessage() {}

func (x *HistoryKbarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_history_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryKbarResponse.ProtoReflect.Descriptor instead.
func (*HistoryKbarResponse) Descriptor() ([]byte, []int) {
	return file_history_proto_rawDescGZIP(), []int{2}
}

func (x *HistoryKbarResponse) GetData() []*HistoryKbarMessage {
	if x != nil {
		return x.Data
	}
	return nil
}

type HistoryKbarMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts     int64   `protobuf:"varint,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Close  float64 `protobuf:"fixed64,2,opt,name=close,proto3" json:"close,omitempty"`
	Open   float64 `protobuf:"fixed64,3,opt,name=open,proto3" json:"open,omitempty"`
	High   float64 `protobuf:"fixed64,4,opt,name=high,proto3" json:"high,omitempty"`
	Low    float64 `protobuf:"fixed64,5,opt,name=low,proto3" json:"low,omitempty"`
	Volume int64   `protobuf:"varint,6,opt,name=volume,proto3" json:"volume,omitempty"`
	Code   string  `protobuf:"bytes,7,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *HistoryKbarMessage) Reset() {
	*x = HistoryKbarMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_history_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryKbarMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryKbarMessage) ProtoMessage() {}

func (x *HistoryKbarMessage) ProtoReflect() protoreflect.Message {
	mi := &file_history_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryKbarMessage.ProtoReflect.Descriptor instead.
func (*HistoryKbarMessage) Descriptor() ([]byte, []int) {
	return file_history_proto_rawDescGZIP(), []int{3}
}

func (x *HistoryKbarMessage) GetTs() int64 {
	if x != nil {
		return x.Ts
	}
	return 0
}

func (x *HistoryKbarMessage) GetClose() float64 {
	if x != nil {
		return x.Close
	}
	return 0
}

func (x *HistoryKbarMessage) GetOpen() float64 {
	if x != nil {
		return x.Open
	}
	return 0
}

func (x *HistoryKbarMessage) GetHigh() float64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *HistoryKbarMessage) GetLow() float64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *HistoryKbarMessage) GetVolume() int64 {
	if x != nil {
		return x.Volume
	}
	return 0
}

func (x *HistoryKbarMessage) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type HistoryCloseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*HistoryCloseMessage `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *HistoryCloseResponse) Reset() {
	*x = HistoryCloseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_history_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryCloseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryCloseResponse) ProtoMessage() {}

func (x *HistoryCloseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_history_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryCloseResponse.ProtoReflect.Descriptor instead.
func (*HistoryCloseResponse) Descriptor() ([]byte, []int) {
	return file_history_proto_rawDescGZIP(), []int{4}
}

func (x *HistoryCloseResponse) GetData() []*HistoryCloseMessage {
	if x != nil {
		return x.Data
	}
	return nil
}

type HistoryCloseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date  string  `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Code  string  `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Close float64 `protobuf:"fixed64,3,opt,name=close,proto3" json:"close,omitempty"`
}

func (x *HistoryCloseMessage) Reset() {
	*x = HistoryCloseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_history_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryCloseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryCloseMessage) ProtoMessage() {}

func (x *HistoryCloseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_history_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryCloseMessage.ProtoReflect.Descriptor instead.
func (*HistoryCloseMessage) Descriptor() ([]byte, []int) {
	return file_history_proto_rawDescGZIP(), []int{5}
}

func (x *HistoryCloseMessage) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *HistoryCloseMessage) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *HistoryCloseMessage) GetClose() float64 {
	if x != nil {
		return x.Close
	}
	return 0
}

var File_history_proto protoreflect.FileDescriptor

var file_history_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x11, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x65, 0x72, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x50, 0x0a, 0x13, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x54, 0x69, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f,
	0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x54, 0x69, 0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0xfb, 0x01, 0x0a, 0x12, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x54, 0x69,
	0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f,
	0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x69, 0x64, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x62, 0x69, 0x64, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x69, 0x64, 0x5f, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x69, 0x64, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x73, 0x6b, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x61, 0x73, 0x6b, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x73, 0x6b, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x73, 0x6b, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x74, 0x69, 0x63, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x22, 0x50, 0x0a, 0x13, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4b, 0x62, 0x61, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f,
	0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x4b, 0x62, 0x61, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0xa0, 0x01, 0x0a, 0x12, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4b, 0x62,
	0x61, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f,
	0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x6f,
	0x70, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x6c, 0x6f, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x52, 0x0a, 0x14, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x69,
	0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x53, 0x0a, 0x13, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x73,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x32, 0x81,
	0x09, 0x0a, 0x14, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x67, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x54, 0x69, 0x63, 0x6b, 0x12, 0x26,
	0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x41, 0x72, 0x72, 0x57, 0x69,
	0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x1a, 0x26, 0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63,
	0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x54, 0x69, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x67, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x4b, 0x62, 0x61, 0x72, 0x12, 0x26, 0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61,
	0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x4e, 0x75, 0x6d, 0x41, 0x72, 0x72, 0x57, 0x69, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x1a,
	0x26, 0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4b, 0x62, 0x61, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x6c, 0x6f, 0x73,
	0x65, 0x12, 0x26, 0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x41, 0x72,
	0x72, 0x57, 0x69, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x1a, 0x27, 0x2e, 0x73, 0x69, 0x6e, 0x6f,
	0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x75, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x42, 0x79, 0x44, 0x61,
	0x74, 0x65, 0x41, 0x72, 0x72, 0x12, 0x29, 0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f,
	0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4e,
	0x75, 0x6d, 0x41, 0x72, 0x72, 0x57, 0x69, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x41, 0x72, 0x72,
	0x1a, 0x27, 0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x6c, 0x6f, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x54, 0x53, 0x45, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x54, 0x69, 0x63, 0x6b, 0x12, 0x17, 0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f,
	0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x1a, 0x26,
	0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x54, 0x69, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5b, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x54, 0x53, 0x45, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4b, 0x62,
	0x61, 0x72, 0x12, 0x17, 0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x1a, 0x26, 0x2e, 0x73, 0x69,
	0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4b, 0x62, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x54, 0x53, 0x45, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x6c, 0x6f, 0x73, 0x65,
	0x12, 0x17, 0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x1a, 0x27, 0x2e, 0x73, 0x69, 0x6e, 0x6f,
	0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4f, 0x54, 0x43, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x4b, 0x62, 0x61, 0x72, 0x12, 0x17, 0x2e, 0x73, 0x69, 0x6e, 0x6f,
	0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x61,
	0x74, 0x65, 0x1a, 0x26, 0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4b, 0x62,
	0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6a, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x54, 0x69, 0x63, 0x6b, 0x12, 0x28, 0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f, 0x66,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x43,
	0x6f, 0x64, 0x65, 0x41, 0x72, 0x72, 0x57, 0x69, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x1a, 0x26,
	0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x54, 0x69, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6c, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x46,
	0x75, 0x74, 0x75, 0x72, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x6c, 0x6f, 0x73,
	0x65, 0x12, 0x28, 0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x41, 0x72, 0x72, 0x57, 0x69, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x1a, 0x27, 0x2e, 0x73, 0x69,
	0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x46, 0x75, 0x74,
	0x75, 0x72, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4b, 0x62, 0x61, 0x72, 0x12, 0x28,
	0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70, 0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x41, 0x72, 0x72,
	0x57, 0x69, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x1a, 0x26, 0x2e, 0x73, 0x69, 0x6e, 0x6f, 0x70,
	0x61, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x4b, 0x62, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_history_proto_rawDescOnce sync.Once
	file_history_proto_rawDescData = file_history_proto_rawDesc
)

func file_history_proto_rawDescGZIP() []byte {
	file_history_proto_rawDescOnce.Do(func() {
		file_history_proto_rawDescData = protoimpl.X.CompressGZIP(file_history_proto_rawDescData)
	})
	return file_history_proto_rawDescData
}

var file_history_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_history_proto_goTypes = []interface{}{
	(*HistoryTickResponse)(nil),    // 0: sinopac_forwarder.HistoryTickResponse
	(*HistoryTickMessage)(nil),     // 1: sinopac_forwarder.HistoryTickMessage
	(*HistoryKbarResponse)(nil),    // 2: sinopac_forwarder.HistoryKbarResponse
	(*HistoryKbarMessage)(nil),     // 3: sinopac_forwarder.HistoryKbarMessage
	(*HistoryCloseResponse)(nil),   // 4: sinopac_forwarder.HistoryCloseResponse
	(*HistoryCloseMessage)(nil),    // 5: sinopac_forwarder.HistoryCloseMessage
	(*StockNumArrWithDate)(nil),    // 6: sinopac_forwarder.StockNumArrWithDate
	(*StockNumArrWithDateArr)(nil), // 7: sinopac_forwarder.StockNumArrWithDateArr
	(*Date)(nil),                   // 8: sinopac_forwarder.Date
	(*FutureCodeArrWithDate)(nil),  // 9: sinopac_forwarder.FutureCodeArrWithDate
}
var file_history_proto_depIdxs = []int32{
	1,  // 0: sinopac_forwarder.HistoryTickResponse.data:type_name -> sinopac_forwarder.HistoryTickMessage
	3,  // 1: sinopac_forwarder.HistoryKbarResponse.data:type_name -> sinopac_forwarder.HistoryKbarMessage
	5,  // 2: sinopac_forwarder.HistoryCloseResponse.data:type_name -> sinopac_forwarder.HistoryCloseMessage
	6,  // 3: sinopac_forwarder.HistoryDataInterface.GetStockHistoryTick:input_type -> sinopac_forwarder.StockNumArrWithDate
	6,  // 4: sinopac_forwarder.HistoryDataInterface.GetStockHistoryKbar:input_type -> sinopac_forwarder.StockNumArrWithDate
	6,  // 5: sinopac_forwarder.HistoryDataInterface.GetStockHistoryClose:input_type -> sinopac_forwarder.StockNumArrWithDate
	7,  // 6: sinopac_forwarder.HistoryDataInterface.GetStockHistoryCloseByDateArr:input_type -> sinopac_forwarder.StockNumArrWithDateArr
	8,  // 7: sinopac_forwarder.HistoryDataInterface.GetStockTSEHistoryTick:input_type -> sinopac_forwarder.Date
	8,  // 8: sinopac_forwarder.HistoryDataInterface.GetStockTSEHistoryKbar:input_type -> sinopac_forwarder.Date
	8,  // 9: sinopac_forwarder.HistoryDataInterface.GetStockTSEHistoryClose:input_type -> sinopac_forwarder.Date
	8,  // 10: sinopac_forwarder.HistoryDataInterface.GetOTCHistoryKbar:input_type -> sinopac_forwarder.Date
	9,  // 11: sinopac_forwarder.HistoryDataInterface.GetFutureHistoryTick:input_type -> sinopac_forwarder.FutureCodeArrWithDate
	9,  // 12: sinopac_forwarder.HistoryDataInterface.GetFutureHistoryClose:input_type -> sinopac_forwarder.FutureCodeArrWithDate
	9,  // 13: sinopac_forwarder.HistoryDataInterface.GetFutureHistoryKbar:input_type -> sinopac_forwarder.FutureCodeArrWithDate
	0,  // 14: sinopac_forwarder.HistoryDataInterface.GetStockHistoryTick:output_type -> sinopac_forwarder.HistoryTickResponse
	2,  // 15: sinopac_forwarder.HistoryDataInterface.GetStockHistoryKbar:output_type -> sinopac_forwarder.HistoryKbarResponse
	4,  // 16: sinopac_forwarder.HistoryDataInterface.GetStockHistoryClose:output_type -> sinopac_forwarder.HistoryCloseResponse
	4,  // 17: sinopac_forwarder.HistoryDataInterface.GetStockHistoryCloseByDateArr:output_type -> sinopac_forwarder.HistoryCloseResponse
	0,  // 18: sinopac_forwarder.HistoryDataInterface.GetStockTSEHistoryTick:output_type -> sinopac_forwarder.HistoryTickResponse
	2,  // 19: sinopac_forwarder.HistoryDataInterface.GetStockTSEHistoryKbar:output_type -> sinopac_forwarder.HistoryKbarResponse
	4,  // 20: sinopac_forwarder.HistoryDataInterface.GetStockTSEHistoryClose:output_type -> sinopac_forwarder.HistoryCloseResponse
	2,  // 21: sinopac_forwarder.HistoryDataInterface.GetOTCHistoryKbar:output_type -> sinopac_forwarder.HistoryKbarResponse
	0,  // 22: sinopac_forwarder.HistoryDataInterface.GetFutureHistoryTick:output_type -> sinopac_forwarder.HistoryTickResponse
	4,  // 23: sinopac_forwarder.HistoryDataInterface.GetFutureHistoryClose:output_type -> sinopac_forwarder.HistoryCloseResponse
	2,  // 24: sinopac_forwarder.HistoryDataInterface.GetFutureHistoryKbar:output_type -> sinopac_forwarder.HistoryKbarResponse
	14, // [14:25] is the sub-list for method output_type
	3,  // [3:14] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_history_proto_init() }
func file_history_proto_init() {
	if File_history_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_history_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryTickResponse); i {
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
		file_history_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryTickMessage); i {
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
		file_history_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryKbarResponse); i {
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
		file_history_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryKbarMessage); i {
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
		file_history_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryCloseResponse); i {
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
		file_history_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryCloseMessage); i {
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
			RawDescriptor: file_history_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_history_proto_goTypes,
		DependencyIndexes: file_history_proto_depIdxs,
		MessageInfos:      file_history_proto_msgTypes,
	}.Build()
	File_history_proto = out.File
	file_history_proto_rawDesc = nil
	file_history_proto_goTypes = nil
	file_history_proto_depIdxs = nil
}
