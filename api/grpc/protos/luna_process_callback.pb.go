// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0--rc1
// source: api/grpc/protos/luna_process_callback.proto

package protos

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

type InstanceStatus int32

const (
	InstanceStatus_DOING    InstanceStatus = 0
	InstanceStatus_REJECTED InstanceStatus = 1
	InstanceStatus_APPROVED InstanceStatus = 2
	InstanceStatus_WITHDRAW InstanceStatus = 3
)

// Enum value maps for InstanceStatus.
var (
	InstanceStatus_name = map[int32]string{
		0: "DOING",
		1: "REJECTED",
		2: "APPROVED",
		3: "WITHDRAW",
	}
	InstanceStatus_value = map[string]int32{
		"DOING":    0,
		"REJECTED": 1,
		"APPROVED": 2,
		"WITHDRAW": 3,
	}
)

func (x InstanceStatus) Enum() *InstanceStatus {
	p := new(InstanceStatus)
	*p = x
	return p
}

func (x InstanceStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InstanceStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_grpc_protos_luna_process_callback_proto_enumTypes[0].Descriptor()
}

func (InstanceStatus) Type() protoreflect.EnumType {
	return &file_api_grpc_protos_luna_process_callback_proto_enumTypes[0]
}

func (x InstanceStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InstanceStatus.Descriptor instead.
func (InstanceStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_grpc_protos_luna_process_callback_proto_rawDescGZIP(), []int{0}
}

type CallbackRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//  流程实例ID, 36位UUID, 建议用来做幂等
	InstanceId string `protobuf:"bytes,1,opt,name=instanceId,proto3" json:"instanceId,omitempty"`
	// 业务外键 可能为空
	BusinessKey string `protobuf:"bytes,2,opt,name=businessKey,proto3" json:"businessKey,omitempty"`
	// 流程定义id
	ProcessId int64 `protobuf:"varint,3,opt,name=processId,proto3" json:"processId,omitempty"`
	// 发起人域账号
	Starter string `protobuf:"bytes,4,opt,name=starter,proto3" json:"starter,omitempty"`
	// 发起时间
	StartTime int64 `protobuf:"varint,5,opt,name=startTime,proto3" json:"startTime,omitempty"`
	// 流程状态
	Status InstanceStatus `protobuf:"varint,6,opt,name=status,proto3,enum=proto.luna.process.callback.InstanceStatus" json:"status,omitempty"`
	// 表单内容
	FormValue string `protobuf:"bytes,7,opt,name=formValue,proto3" json:"formValue,omitempty"`
}

func (x *CallbackRequest) Reset() {
	*x = CallbackRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_protos_luna_process_callback_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallbackRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallbackRequest) ProtoMessage() {}

func (x *CallbackRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_protos_luna_process_callback_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallbackRequest.ProtoReflect.Descriptor instead.
func (*CallbackRequest) Descriptor() ([]byte, []int) {
	return file_api_grpc_protos_luna_process_callback_proto_rawDescGZIP(), []int{0}
}

func (x *CallbackRequest) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

func (x *CallbackRequest) GetBusinessKey() string {
	if x != nil {
		return x.BusinessKey
	}
	return ""
}

func (x *CallbackRequest) GetProcessId() int64 {
	if x != nil {
		return x.ProcessId
	}
	return 0
}

func (x *CallbackRequest) GetStarter() string {
	if x != nil {
		return x.Starter
	}
	return ""
}

func (x *CallbackRequest) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *CallbackRequest) GetStatus() InstanceStatus {
	if x != nil {
		return x.Status
	}
	return InstanceStatus_DOING
}

func (x *CallbackRequest) GetFormValue() string {
	if x != nil {
		return x.FormValue
	}
	return ""
}

type CallbackReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Succeed bool   `protobuf:"varint,1,opt,name=succeed,proto3" json:"succeed,omitempty"`
	Msg     string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *CallbackReply) Reset() {
	*x = CallbackReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_protos_luna_process_callback_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallbackReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallbackReply) ProtoMessage() {}

func (x *CallbackReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_protos_luna_process_callback_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallbackReply.ProtoReflect.Descriptor instead.
func (*CallbackReply) Descriptor() ([]byte, []int) {
	return file_api_grpc_protos_luna_process_callback_proto_rawDescGZIP(), []int{1}
}

func (x *CallbackReply) GetSucceed() bool {
	if x != nil {
		return x.Succeed
	}
	return false
}

func (x *CallbackReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_api_grpc_protos_luna_process_callback_proto protoreflect.FileDescriptor

var file_api_grpc_protos_luna_process_callback_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x6c, 0x75, 0x6e, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x63,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x75, 0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x2e, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x22, 0x8c, 0x02, 0x0a, 0x0f, 0x43,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c,
	0x75, 0x6e, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x66,
	0x6f, 0x72, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x66, 0x6f, 0x72, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3b, 0x0a, 0x0d, 0x43, 0x61, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x2a, 0x45, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x4f, 0x49, 0x4e,
	0x47, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x44, 0x10, 0x02, 0x12,
	0x0c, 0x0a, 0x08, 0x57, 0x49, 0x54, 0x48, 0x44, 0x52, 0x41, 0x57, 0x10, 0x03, 0x32, 0x7d, 0x0a,
	0x13, 0x4c, 0x75, 0x6e, 0x61, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x66, 0x0a, 0x0a, 0x44, 0x6f, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x12, 0x2c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x75, 0x6e, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6c, 0x75, 0x6e, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x2e, 0x43,
	0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x18, 0x5a, 0x16,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_grpc_protos_luna_process_callback_proto_rawDescOnce sync.Once
	file_api_grpc_protos_luna_process_callback_proto_rawDescData = file_api_grpc_protos_luna_process_callback_proto_rawDesc
)

func file_api_grpc_protos_luna_process_callback_proto_rawDescGZIP() []byte {
	file_api_grpc_protos_luna_process_callback_proto_rawDescOnce.Do(func() {
		file_api_grpc_protos_luna_process_callback_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_grpc_protos_luna_process_callback_proto_rawDescData)
	})
	return file_api_grpc_protos_luna_process_callback_proto_rawDescData
}

var file_api_grpc_protos_luna_process_callback_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_grpc_protos_luna_process_callback_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_grpc_protos_luna_process_callback_proto_goTypes = []interface{}{
	(InstanceStatus)(0),     // 0: proto.luna.process.callback.InstanceStatus
	(*CallbackRequest)(nil), // 1: proto.luna.process.callback.CallbackRequest
	(*CallbackReply)(nil),   // 2: proto.luna.process.callback.CallbackReply
}
var file_api_grpc_protos_luna_process_callback_proto_depIdxs = []int32{
	0, // 0: proto.luna.process.callback.CallbackRequest.status:type_name -> proto.luna.process.callback.InstanceStatus
	1, // 1: proto.luna.process.callback.LunaCallbackService.DoCallback:input_type -> proto.luna.process.callback.CallbackRequest
	2, // 2: proto.luna.process.callback.LunaCallbackService.DoCallback:output_type -> proto.luna.process.callback.CallbackReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_grpc_protos_luna_process_callback_proto_init() }
func file_api_grpc_protos_luna_process_callback_proto_init() {
	if File_api_grpc_protos_luna_process_callback_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_grpc_protos_luna_process_callback_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallbackRequest); i {
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
		file_api_grpc_protos_luna_process_callback_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallbackReply); i {
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
			RawDescriptor: file_api_grpc_protos_luna_process_callback_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_grpc_protos_luna_process_callback_proto_goTypes,
		DependencyIndexes: file_api_grpc_protos_luna_process_callback_proto_depIdxs,
		EnumInfos:         file_api_grpc_protos_luna_process_callback_proto_enumTypes,
		MessageInfos:      file_api_grpc_protos_luna_process_callback_proto_msgTypes,
	}.Build()
	File_api_grpc_protos_luna_process_callback_proto = out.File
	file_api_grpc_protos_luna_process_callback_proto_rawDesc = nil
	file_api_grpc_protos_luna_process_callback_proto_goTypes = nil
	file_api_grpc_protos_luna_process_callback_proto_depIdxs = nil
}
