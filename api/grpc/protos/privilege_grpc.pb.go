// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PrivilegeClient is the client API for Privilege service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PrivilegeClient interface {
	GetUserPrivilegeAll(ctx context.Context, in *GetUserPrivilegeAllParam, opts ...grpc.CallOption) (*GetUserPrivilegeAllReply, error)
}

type privilegeClient struct {
	cc grpc.ClientConnInterface
}

func NewPrivilegeClient(cc grpc.ClientConnInterface) PrivilegeClient {
	return &privilegeClient{cc}
}

func (c *privilegeClient) GetUserPrivilegeAll(ctx context.Context, in *GetUserPrivilegeAllParam, opts ...grpc.CallOption) (*GetUserPrivilegeAllReply, error) {
	out := new(GetUserPrivilegeAllReply)
	err := c.cc.Invoke(ctx, "/protos.Privilege/GetUserPrivilegeAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrivilegeServer is the server API for Privilege service.
// All implementations should embed UnimplementedPrivilegeServer
// for forward compatibility
type PrivilegeServer interface {
	GetUserPrivilegeAll(context.Context, *GetUserPrivilegeAllParam) (*GetUserPrivilegeAllReply, error)
}

// UnimplementedPrivilegeServer should be embedded to have forward compatible implementations.
type UnimplementedPrivilegeServer struct {
}

func (UnimplementedPrivilegeServer) GetUserPrivilegeAll(context.Context, *GetUserPrivilegeAllParam) (*GetUserPrivilegeAllReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPrivilegeAll not implemented")
}

// UnsafePrivilegeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PrivilegeServer will
// result in compilation errors.
type UnsafePrivilegeServer interface {
	mustEmbedUnimplementedPrivilegeServer()
}

func RegisterPrivilegeServer(s grpc.ServiceRegistrar, srv PrivilegeServer) {
	s.RegisterService(&Privilege_ServiceDesc, srv)
}

func _Privilege_GetUserPrivilegeAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserPrivilegeAllParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrivilegeServer).GetUserPrivilegeAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Privilege/GetUserPrivilegeAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrivilegeServer).GetUserPrivilegeAll(ctx, req.(*GetUserPrivilegeAllParam))
	}
	return interceptor(ctx, in, info, handler)
}

// Privilege_ServiceDesc is the grpc.ServiceDesc for Privilege service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Privilege_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Privilege",
	HandlerType: (*PrivilegeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserPrivilegeAll",
			Handler:    _Privilege_GetUserPrivilegeAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/grpc/protos/privilege.proto",
}
