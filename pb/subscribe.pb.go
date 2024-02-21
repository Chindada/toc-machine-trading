// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: forwarder/subscribe.proto

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

type SubscribeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FailArr []string `protobuf:"bytes,1,rep,name=fail_arr,json=failArr,proto3" json:"fail_arr,omitempty"`
}

func (x *SubscribeResponse) Reset() {
	*x = SubscribeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_forwarder_subscribe_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeResponse) ProtoMessage() {}

func (x *SubscribeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_forwarder_subscribe_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeResponse.ProtoReflect.Descriptor instead.
func (*SubscribeResponse) Descriptor() ([]byte, []int) {
	return file_forwarder_subscribe_proto_rawDescGZIP(), []int{0}
}

func (x *SubscribeResponse) GetFailArr() []string {
	if x != nil {
		return x.FailArr
	}
	return nil
}

var File_forwarder_subscribe_proto protoreflect.FileDescriptor

var file_forwarder_subscribe_proto_rawDesc = []byte{
	0x0a, 0x19, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x66, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x11, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x66, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x72, 0x72, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x41, 0x72, 0x72, 0x32, 0xb8, 0x06, 0x0a, 0x16,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x44, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x12, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x63, 0x6b, 0x12, 0x16, 0x2e, 0x66,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4e, 0x75,
	0x6d, 0x41, 0x72, 0x72, 0x1a, 0x1c, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x14, 0x55, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x63, 0x6b, 0x12, 0x16, 0x2e, 0x66,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4e, 0x75,
	0x6d, 0x41, 0x72, 0x72, 0x1a, 0x1c, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x14, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x42, 0x69, 0x64, 0x41, 0x73, 0x6b, 0x12, 0x16, 0x2e, 0x66,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x4e, 0x75,
	0x6d, 0x41, 0x72, 0x72, 0x1a, 0x1c, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x16, 0x55, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x42, 0x69, 0x64, 0x41, 0x73, 0x6b, 0x12, 0x16,
	0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x4e, 0x75, 0x6d, 0x41, 0x72, 0x72, 0x1a, 0x1c, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x13, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x12, 0x18, 0x2e,
	0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65,
	0x43, 0x6f, 0x64, 0x65, 0x41, 0x72, 0x72, 0x1a, 0x1c, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x15, 0x55, 0x6e, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x54, 0x69, 0x63, 0x6b,
	0x12, 0x18, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x46, 0x75, 0x74,
	0x75, 0x72, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x41, 0x72, 0x72, 0x1a, 0x1c, 0x2e, 0x66, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x15, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x42, 0x69, 0x64,
	0x41, 0x73, 0x6b, 0x12, 0x18, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x41, 0x72, 0x72, 0x1a, 0x1c, 0x2e,
	0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a,
	0x17, 0x55, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x46, 0x75, 0x74, 0x75,
	0x72, 0x65, 0x42, 0x69, 0x64, 0x41, 0x73, 0x6b, 0x12, 0x18, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x46, 0x75, 0x74, 0x75, 0x72, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x41,
	0x72, 0x72, 0x1a, 0x1c, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x47, 0x0a, 0x12, 0x55, 0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x41, 0x6c, 0x6c, 0x54, 0x69, 0x63, 0x6b, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x17, 0x2e, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x14, 0x55,
	0x6e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x6c, 0x6c, 0x42, 0x69, 0x64,
	0x41, 0x73, 0x6b, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x66, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_forwarder_subscribe_proto_rawDescOnce sync.Once
	file_forwarder_subscribe_proto_rawDescData = file_forwarder_subscribe_proto_rawDesc
)

func file_forwarder_subscribe_proto_rawDescGZIP() []byte {
	file_forwarder_subscribe_proto_rawDescOnce.Do(func() {
		file_forwarder_subscribe_proto_rawDescData = protoimpl.X.CompressGZIP(file_forwarder_subscribe_proto_rawDescData)
	})
	return file_forwarder_subscribe_proto_rawDescData
}

var file_forwarder_subscribe_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_forwarder_subscribe_proto_goTypes = []interface{}{
	(*SubscribeResponse)(nil), // 0: forwarder.SubscribeResponse
	(*StockNumArr)(nil),       // 1: forwarder.StockNumArr
	(*FutureCodeArr)(nil),     // 2: forwarder.FutureCodeArr
	(*emptypb.Empty)(nil),     // 3: google.protobuf.Empty
	(*ErrorMessage)(nil),      // 4: forwarder.ErrorMessage
}
var file_forwarder_subscribe_proto_depIdxs = []int32{
	1,  // 0: forwarder.SubscribeDataInterface.SubscribeStockTick:input_type -> forwarder.StockNumArr
	1,  // 1: forwarder.SubscribeDataInterface.UnSubscribeStockTick:input_type -> forwarder.StockNumArr
	1,  // 2: forwarder.SubscribeDataInterface.SubscribeStockBidAsk:input_type -> forwarder.StockNumArr
	1,  // 3: forwarder.SubscribeDataInterface.UnSubscribeStockBidAsk:input_type -> forwarder.StockNumArr
	2,  // 4: forwarder.SubscribeDataInterface.SubscribeFutureTick:input_type -> forwarder.FutureCodeArr
	2,  // 5: forwarder.SubscribeDataInterface.UnSubscribeFutureTick:input_type -> forwarder.FutureCodeArr
	2,  // 6: forwarder.SubscribeDataInterface.SubscribeFutureBidAsk:input_type -> forwarder.FutureCodeArr
	2,  // 7: forwarder.SubscribeDataInterface.UnSubscribeFutureBidAsk:input_type -> forwarder.FutureCodeArr
	3,  // 8: forwarder.SubscribeDataInterface.UnSubscribeAllTick:input_type -> google.protobuf.Empty
	3,  // 9: forwarder.SubscribeDataInterface.UnSubscribeAllBidAsk:input_type -> google.protobuf.Empty
	0,  // 10: forwarder.SubscribeDataInterface.SubscribeStockTick:output_type -> forwarder.SubscribeResponse
	0,  // 11: forwarder.SubscribeDataInterface.UnSubscribeStockTick:output_type -> forwarder.SubscribeResponse
	0,  // 12: forwarder.SubscribeDataInterface.SubscribeStockBidAsk:output_type -> forwarder.SubscribeResponse
	0,  // 13: forwarder.SubscribeDataInterface.UnSubscribeStockBidAsk:output_type -> forwarder.SubscribeResponse
	0,  // 14: forwarder.SubscribeDataInterface.SubscribeFutureTick:output_type -> forwarder.SubscribeResponse
	0,  // 15: forwarder.SubscribeDataInterface.UnSubscribeFutureTick:output_type -> forwarder.SubscribeResponse
	0,  // 16: forwarder.SubscribeDataInterface.SubscribeFutureBidAsk:output_type -> forwarder.SubscribeResponse
	0,  // 17: forwarder.SubscribeDataInterface.UnSubscribeFutureBidAsk:output_type -> forwarder.SubscribeResponse
	4,  // 18: forwarder.SubscribeDataInterface.UnSubscribeAllTick:output_type -> forwarder.ErrorMessage
	4,  // 19: forwarder.SubscribeDataInterface.UnSubscribeAllBidAsk:output_type -> forwarder.ErrorMessage
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_forwarder_subscribe_proto_init() }
func file_forwarder_subscribe_proto_init() {
	if File_forwarder_subscribe_proto != nil {
		return
	}
	file_forwarder_entity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_forwarder_subscribe_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeResponse); i {
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
			RawDescriptor: file_forwarder_subscribe_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_forwarder_subscribe_proto_goTypes,
		DependencyIndexes: file_forwarder_subscribe_proto_depIdxs,
		MessageInfos:      file_forwarder_subscribe_proto_msgTypes,
	}.Build()
	File_forwarder_subscribe_proto = out.File
	file_forwarder_subscribe_proto_rawDesc = nil
	file_forwarder_subscribe_proto_goTypes = nil
	file_forwarder_subscribe_proto_depIdxs = nil
}
