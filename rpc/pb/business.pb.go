// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: business.proto

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

// 换电状态
type BusinessExchangeStatus int32

const (
	// 处理中
	BusinessExchangeStatus_Processing BusinessExchangeStatus = 0
	// 成功
	BusinessExchangeStatus_Success BusinessExchangeStatus = 1
	// 失败
	BusinessExchangeStatus_Failed BusinessExchangeStatus = 2
)

// Enum value maps for BusinessExchangeStatus.
var (
	BusinessExchangeStatus_name = map[int32]string{
		0: "Processing",
		1: "Success",
		2: "Failed",
	}
	BusinessExchangeStatus_value = map[string]int32{
		"Processing": 0,
		"Success":    1,
		"Failed":     2,
	}
)

func (x BusinessExchangeStatus) Enum() *BusinessExchangeStatus {
	p := new(BusinessExchangeStatus)
	*p = x
	return p
}

func (x BusinessExchangeStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BusinessExchangeStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_business_proto_enumTypes[0].Descriptor()
}

func (BusinessExchangeStatus) Type() protoreflect.EnumType {
	return &file_business_proto_enumTypes[0]
}

func (x BusinessExchangeStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BusinessExchangeStatus.Descriptor instead.
func (BusinessExchangeStatus) EnumDescriptor() ([]byte, []int) {
	return file_business_proto_rawDescGZIP(), []int{0}
}

type BusinessExchangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid    string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Serial  string `protobuf:"bytes,2,opt,name=serial,proto3" json:"serial,omitempty"`
	Battery string `protobuf:"bytes,3,opt,name=battery,proto3" json:"battery,omitempty"` //  bool alternative = 4; // 是否使用备用方案
}

func (x *BusinessExchangeRequest) Reset() {
	*x = BusinessExchangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_business_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusinessExchangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusinessExchangeRequest) ProtoMessage() {}

func (x *BusinessExchangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_business_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusinessExchangeRequest.ProtoReflect.Descriptor instead.
func (*BusinessExchangeRequest) Descriptor() ([]byte, []int) {
	return file_business_proto_rawDescGZIP(), []int{0}
}

func (x *BusinessExchangeRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *BusinessExchangeRequest) GetSerial() string {
	if x != nil {
		return x.Serial
	}
	return ""
}

func (x *BusinessExchangeRequest) GetBattery() string {
	if x != nil {
		return x.Battery
	}
	return ""
}

type BusinessExchangeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 操作步骤 1:开空电仓 2:放旧电池 3:开满电仓 4:取新电池
	Step uint32 `protobuf:"varint,1,opt,name=step,proto3" json:"step,omitempty"`
	// 状态 1:处理中 2:成功 3:失败
	Status BusinessExchangeStatus `protobuf:"varint,2,opt,name=status,proto3,enum=pb.BusinessExchangeStatus" json:"status,omitempty"`
	// 消息
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// 步骤是否终止
	Stop bool `protobuf:"varint,4,opt,name=stop,proto3" json:"stop,omitempty"`
}

func (x *BusinessExchangeResponse) Reset() {
	*x = BusinessExchangeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_business_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusinessExchangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusinessExchangeResponse) ProtoMessage() {}

func (x *BusinessExchangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_business_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusinessExchangeResponse.ProtoReflect.Descriptor instead.
func (*BusinessExchangeResponse) Descriptor() ([]byte, []int) {
	return file_business_proto_rawDescGZIP(), []int{1}
}

func (x *BusinessExchangeResponse) GetStep() uint32 {
	if x != nil {
		return x.Step
	}
	return 0
}

func (x *BusinessExchangeResponse) GetStatus() BusinessExchangeStatus {
	if x != nil {
		return x.Status
	}
	return BusinessExchangeStatus_Processing
}

func (x *BusinessExchangeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *BusinessExchangeResponse) GetStop() bool {
	if x != nil {
		return x.Stop
	}
	return false
}

var File_business_proto protoreflect.FileDescriptor

var file_business_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0x5f, 0x0a, 0x17, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x62,
	0x61, 0x74, 0x74, 0x65, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61,
	0x74, 0x74, 0x65, 0x72, 0x79, 0x22, 0x90, 0x01, 0x0a, 0x18, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x6f, 0x70, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x73, 0x74, 0x6f, 0x70, 0x2a, 0x41, 0x0a, 0x16, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x02, 0x32, 0x55, 0x0a, 0x08, 0x42,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x49, 0x0a, 0x08, 0x45, 0x78, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x45, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x30, 0x01, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x75, 0x72, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x64, 0x65, 0x2f, 0x61, 0x64, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_business_proto_rawDescOnce sync.Once
	file_business_proto_rawDescData = file_business_proto_rawDesc
)

func file_business_proto_rawDescGZIP() []byte {
	file_business_proto_rawDescOnce.Do(func() {
		file_business_proto_rawDescData = protoimpl.X.CompressGZIP(file_business_proto_rawDescData)
	})
	return file_business_proto_rawDescData
}

var file_business_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_business_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_business_proto_goTypes = []interface{}{
	(BusinessExchangeStatus)(0),      // 0: pb.BusinessExchangeStatus
	(*BusinessExchangeRequest)(nil),  // 1: pb.BusinessExchangeRequest
	(*BusinessExchangeResponse)(nil), // 2: pb.BusinessExchangeResponse
}
var file_business_proto_depIdxs = []int32{
	0, // 0: pb.BusinessExchangeResponse.status:type_name -> pb.BusinessExchangeStatus
	1, // 1: pb.Business.Exchange:input_type -> pb.BusinessExchangeRequest
	2, // 2: pb.Business.Exchange:output_type -> pb.BusinessExchangeResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_business_proto_init() }
func file_business_proto_init() {
	if File_business_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_business_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusinessExchangeRequest); i {
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
		file_business_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusinessExchangeResponse); i {
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
			RawDescriptor: file_business_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_business_proto_goTypes,
		DependencyIndexes: file_business_proto_depIdxs,
		EnumInfos:         file_business_proto_enumTypes,
		MessageInfos:      file_business_proto_msgTypes,
	}.Build()
	File_business_proto = out.File
	file_business_proto_rawDesc = nil
	file_business_proto_goTypes = nil
	file_business_proto_depIdxs = nil
}
