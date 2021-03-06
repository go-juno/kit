// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0--rc1
// source: api/grpc/protos/role.proto

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

type SetUserRoleParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain  string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	RoleKey string `protobuf:"bytes,2,opt,name=roleKey,proto3" json:"roleKey,omitempty"`
	SiteKey string `protobuf:"bytes,3,opt,name=siteKey,proto3" json:"siteKey,omitempty"`
}

func (x *SetUserRoleParam) Reset() {
	*x = SetUserRoleParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_protos_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetUserRoleParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetUserRoleParam) ProtoMessage() {}

func (x *SetUserRoleParam) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_protos_role_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetUserRoleParam.ProtoReflect.Descriptor instead.
func (*SetUserRoleParam) Descriptor() ([]byte, []int) {
	return file_api_grpc_protos_role_proto_rawDescGZIP(), []int{0}
}

func (x *SetUserRoleParam) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *SetUserRoleParam) GetRoleKey() string {
	if x != nil {
		return x.RoleKey
	}
	return ""
}

func (x *SetUserRoleParam) GetSiteKey() string {
	if x != nil {
		return x.SiteKey
	}
	return ""
}

type SetUserRoleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetUserRoleReply) Reset() {
	*x = SetUserRoleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_protos_role_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetUserRoleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetUserRoleReply) ProtoMessage() {}

func (x *SetUserRoleReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_protos_role_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetUserRoleReply.ProtoReflect.Descriptor instead.
func (*SetUserRoleReply) Descriptor() ([]byte, []int) {
	return file_api_grpc_protos_role_proto_rawDescGZIP(), []int{1}
}

type GetUserRoleParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain  string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	SiteKey string `protobuf:"bytes,2,opt,name=siteKey,proto3" json:"siteKey,omitempty"`
}

func (x *GetUserRoleParam) Reset() {
	*x = GetUserRoleParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_protos_role_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRoleParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRoleParam) ProtoMessage() {}

func (x *GetUserRoleParam) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_protos_role_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRoleParam.ProtoReflect.Descriptor instead.
func (*GetUserRoleParam) Descriptor() ([]byte, []int) {
	return file_api_grpc_protos_role_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserRoleParam) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *GetUserRoleParam) GetSiteKey() string {
	if x != nil {
		return x.SiteKey
	}
	return ""
}

type GetUserRoleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role []*GetUserRoleReply_Role `protobuf:"bytes,1,rep,name=role,proto3" json:"role,omitempty"`
}

func (x *GetUserRoleReply) Reset() {
	*x = GetUserRoleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_protos_role_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRoleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRoleReply) ProtoMessage() {}

func (x *GetUserRoleReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_protos_role_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRoleReply.ProtoReflect.Descriptor instead.
func (*GetUserRoleReply) Descriptor() ([]byte, []int) {
	return file_api_grpc_protos_role_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserRoleReply) GetRole() []*GetUserRoleReply_Role {
	if x != nil {
		return x.Role
	}
	return nil
}

type GetUserListByRoleKeyParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SiteKey string `protobuf:"bytes,1,opt,name=siteKey,proto3" json:"siteKey,omitempty"`
	RoleKey string `protobuf:"bytes,2,opt,name=roleKey,proto3" json:"roleKey,omitempty"`
}

func (x *GetUserListByRoleKeyParam) Reset() {
	*x = GetUserListByRoleKeyParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_protos_role_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserListByRoleKeyParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserListByRoleKeyParam) ProtoMessage() {}

func (x *GetUserListByRoleKeyParam) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_protos_role_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserListByRoleKeyParam.ProtoReflect.Descriptor instead.
func (*GetUserListByRoleKeyParam) Descriptor() ([]byte, []int) {
	return file_api_grpc_protos_role_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserListByRoleKeyParam) GetSiteKey() string {
	if x != nil {
		return x.SiteKey
	}
	return ""
}

func (x *GetUserListByRoleKeyParam) GetRoleKey() string {
	if x != nil {
		return x.RoleKey
	}
	return ""
}

type GetUserListByRoleKeyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain []string `protobuf:"bytes,1,rep,name=domain,proto3" json:"domain,omitempty"`
}

func (x *GetUserListByRoleKeyReply) Reset() {
	*x = GetUserListByRoleKeyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_protos_role_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserListByRoleKeyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserListByRoleKeyReply) ProtoMessage() {}

func (x *GetUserListByRoleKeyReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_protos_role_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserListByRoleKeyReply.ProtoReflect.Descriptor instead.
func (*GetUserListByRoleKeyReply) Descriptor() ([]byte, []int) {
	return file_api_grpc_protos_role_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserListByRoleKeyReply) GetDomain() []string {
	if x != nil {
		return x.Domain
	}
	return nil
}

type GetUserRoleReply_Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleKey string `protobuf:"bytes,1,opt,name=roleKey,proto3" json:"roleKey,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetUserRoleReply_Role) Reset() {
	*x = GetUserRoleReply_Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_grpc_protos_role_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRoleReply_Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRoleReply_Role) ProtoMessage() {}

func (x *GetUserRoleReply_Role) ProtoReflect() protoreflect.Message {
	mi := &file_api_grpc_protos_role_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRoleReply_Role.ProtoReflect.Descriptor instead.
func (*GetUserRoleReply_Role) Descriptor() ([]byte, []int) {
	return file_api_grpc_protos_role_proto_rawDescGZIP(), []int{3, 0}
}

func (x *GetUserRoleReply_Role) GetRoleKey() string {
	if x != nil {
		return x.RoleKey
	}
	return ""
}

func (x *GetUserRoleReply_Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_api_grpc_protos_role_proto protoreflect.FileDescriptor

var file_api_grpc_protos_role_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x22, 0x5e, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x6f, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x69,
	0x74, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x69, 0x74,
	0x65, 0x4b, 0x65, 0x79, 0x22, 0x12, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x6f, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x44, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x16, 0x0a, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x69, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x69, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x22, 0x7b,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x31, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x1a, 0x34, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x6f, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x72, 0x6f, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x4f, 0x0a, 0x19, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x52, 0x6f, 0x6c, 0x65,
	0x4b, 0x65, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x69, 0x74, 0x65,
	0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x69, 0x74, 0x65, 0x4b,
	0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x22, 0x33, 0x0a, 0x19,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x52, 0x6f, 0x6c,
	0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x32, 0xf0, 0x01, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x43, 0x0a, 0x0b, 0x53, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x43, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x6f, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a,
	0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x52, 0x6f, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_grpc_protos_role_proto_rawDescOnce sync.Once
	file_api_grpc_protos_role_proto_rawDescData = file_api_grpc_protos_role_proto_rawDesc
)

func file_api_grpc_protos_role_proto_rawDescGZIP() []byte {
	file_api_grpc_protos_role_proto_rawDescOnce.Do(func() {
		file_api_grpc_protos_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_grpc_protos_role_proto_rawDescData)
	})
	return file_api_grpc_protos_role_proto_rawDescData
}

var file_api_grpc_protos_role_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_grpc_protos_role_proto_goTypes = []interface{}{
	(*SetUserRoleParam)(nil),          // 0: protos.SetUserRoleParam
	(*SetUserRoleReply)(nil),          // 1: protos.SetUserRoleReply
	(*GetUserRoleParam)(nil),          // 2: protos.GetUserRoleParam
	(*GetUserRoleReply)(nil),          // 3: protos.GetUserRoleReply
	(*GetUserListByRoleKeyParam)(nil), // 4: protos.GetUserListByRoleKeyParam
	(*GetUserListByRoleKeyReply)(nil), // 5: protos.GetUserListByRoleKeyReply
	(*GetUserRoleReply_Role)(nil),     // 6: protos.GetUserRoleReply.Role
}
var file_api_grpc_protos_role_proto_depIdxs = []int32{
	6, // 0: protos.GetUserRoleReply.role:type_name -> protos.GetUserRoleReply.Role
	0, // 1: protos.Role.SetUserRole:input_type -> protos.SetUserRoleParam
	2, // 2: protos.Role.GetUserRole:input_type -> protos.GetUserRoleParam
	4, // 3: protos.Role.GetUserListByRoleKey:input_type -> protos.GetUserListByRoleKeyParam
	1, // 4: protos.Role.SetUserRole:output_type -> protos.SetUserRoleReply
	3, // 5: protos.Role.GetUserRole:output_type -> protos.GetUserRoleReply
	5, // 6: protos.Role.GetUserListByRoleKey:output_type -> protos.GetUserListByRoleKeyReply
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_grpc_protos_role_proto_init() }
func file_api_grpc_protos_role_proto_init() {
	if File_api_grpc_protos_role_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_grpc_protos_role_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetUserRoleParam); i {
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
		file_api_grpc_protos_role_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetUserRoleReply); i {
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
		file_api_grpc_protos_role_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRoleParam); i {
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
		file_api_grpc_protos_role_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRoleReply); i {
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
		file_api_grpc_protos_role_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserListByRoleKeyParam); i {
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
		file_api_grpc_protos_role_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserListByRoleKeyReply); i {
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
		file_api_grpc_protos_role_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRoleReply_Role); i {
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
			RawDescriptor: file_api_grpc_protos_role_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_grpc_protos_role_proto_goTypes,
		DependencyIndexes: file_api_grpc_protos_role_proto_depIdxs,
		MessageInfos:      file_api_grpc_protos_role_proto_msgTypes,
	}.Build()
	File_api_grpc_protos_role_proto = out.File
	file_api_grpc_protos_role_proto_rawDesc = nil
	file_api_grpc_protos_role_proto_goTypes = nil
	file_api_grpc_protos_role_proto_depIdxs = nil
}
