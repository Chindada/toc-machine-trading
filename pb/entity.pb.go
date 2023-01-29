// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: entity.proto

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

// StockNumArr
type StockNumArr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StockNumArr []string `protobuf:"bytes,1,rep,name=stock_num_arr,json=stockNumArr,proto3" json:"stock_num_arr,omitempty"`
}

func (x *StockNumArr) Reset() {
	*x = StockNumArr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StockNumArr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StockNumArr) ProtoMessage() {}

func (x *StockNumArr) ProtoReflect() protoreflect.Message {
	mi := &file_entity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StockNumArr.ProtoReflect.Descriptor instead.
func (*StockNumArr) Descriptor() ([]byte, []int) {
	return file_entity_proto_rawDescGZIP(), []int{0}
}

func (x *StockNumArr) GetStockNumArr() []string {
	if x != nil {
		return x.StockNumArr
	}
	return nil
}

// FutureNumArr
type FutureCodeArr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FutureCodeArr []string `protobuf:"bytes,1,rep,name=future_code_arr,json=futureCodeArr,proto3" json:"future_code_arr,omitempty"`
}

func (x *FutureCodeArr) Reset() {
	*x = FutureCodeArr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FutureCodeArr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FutureCodeArr) ProtoMessage() {}

func (x *FutureCodeArr) ProtoReflect() protoreflect.Message {
	mi := &file_entity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FutureCodeArr.ProtoReflect.Descriptor instead.
func (*FutureCodeArr) Descriptor() ([]byte, []int) {
	return file_entity_proto_rawDescGZIP(), []int{1}
}

func (x *FutureCodeArr) GetFutureCodeArr() []string {
	if x != nil {
		return x.FutureCodeArr
	}
	return nil
}

// ErrorMssage
type ErrorMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err string `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *ErrorMessage) Reset() {
	*x = ErrorMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorMessage) ProtoMessage() {}

func (x *ErrorMessage) ProtoReflect() protoreflect.Message {
	mi := &file_entity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorMessage.ProtoReflect.Descriptor instead.
func (*ErrorMessage) Descriptor() ([]byte, []int) {
	return file_entity_proto_rawDescGZIP(), []int{2}
}

func (x *ErrorMessage) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

var File_entity_proto protoreflect.FileDescriptor

var file_entity_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14,
	0x74, 0x6f, 0x63, 0x5f, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x65, 0x72, 0x22, 0x31, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d,
	0x41, 0x72, 0x72, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d,
	0x5f, 0x61, 0x72, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x4e, 0x75, 0x6d, 0x41, 0x72, 0x72, 0x22, 0x37, 0x0a, 0x0d, 0x46, 0x75, 0x74, 0x75, 0x72,
	0x65, 0x43, 0x6f, 0x64, 0x65, 0x41, 0x72, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x66, 0x75, 0x74, 0x75,
	0x72, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x61, 0x72, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0d, 0x66, 0x75, 0x74, 0x75, 0x72, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x41, 0x72, 0x72,
	0x22, 0x20, 0x0a, 0x0c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65,
	0x72, 0x72, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_entity_proto_rawDescOnce sync.Once
	file_entity_proto_rawDescData = file_entity_proto_rawDesc
)

func file_entity_proto_rawDescGZIP() []byte {
	file_entity_proto_rawDescOnce.Do(func() {
		file_entity_proto_rawDescData = protoimpl.X.CompressGZIP(file_entity_proto_rawDescData)
	})
	return file_entity_proto_rawDescData
}

var file_entity_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_entity_proto_goTypes = []interface{}{
	(*StockNumArr)(nil),   // 0: toc_python_forwarder.StockNumArr
	(*FutureCodeArr)(nil), // 1: toc_python_forwarder.FutureCodeArr
	(*ErrorMessage)(nil),  // 2: toc_python_forwarder.ErrorMessage
}
var file_entity_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_entity_proto_init() }
func file_entity_proto_init() {
	if File_entity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_entity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StockNumArr); i {
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
		file_entity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FutureCodeArr); i {
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
		file_entity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorMessage); i {
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
			RawDescriptor: file_entity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entity_proto_goTypes,
		DependencyIndexes: file_entity_proto_depIdxs,
		MessageInfos:      file_entity_proto_msgTypes,
	}.Build()
	File_entity_proto = out.File
	file_entity_proto_rawDesc = nil
	file_entity_proto_goTypes = nil
	file_entity_proto_depIdxs = nil
}
