// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: notify.proto

package proto

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

// NotifierClient is the client API for Notifier service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotifierClient interface {
	NotifyUser(ctx context.Context, in *NotifyUserRequest, opts ...grpc.CallOption) (*NotifyResponse, error)
	NotifyUserBatch(ctx context.Context, in *NotifyUserBatchRequest, opts ...grpc.CallOption) (*NotifyResponse, error)
	NotifyAllUser(ctx context.Context, in *NotifyRequest, opts ...grpc.CallOption) (*NotifyResponse, error)
}

type notifierClient struct {
	cc grpc.ClientConnInterface
}

func NewNotifierClient(cc grpc.ClientConnInterface) NotifierClient {
	return &notifierClient{cc}
}

func (c *notifierClient) NotifyUser(ctx context.Context, in *NotifyUserRequest, opts ...grpc.CallOption) (*NotifyResponse, error) {
	out := new(NotifyResponse)
	err := c.cc.Invoke(ctx, "/proto.Notifier/NotifyUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifierClient) NotifyUserBatch(ctx context.Context, in *NotifyUserBatchRequest, opts ...grpc.CallOption) (*NotifyResponse, error) {
	out := new(NotifyResponse)
	err := c.cc.Invoke(ctx, "/proto.Notifier/NotifyUserBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notifierClient) NotifyAllUser(ctx context.Context, in *NotifyRequest, opts ...grpc.CallOption) (*NotifyResponse, error) {
	out := new(NotifyResponse)
	err := c.cc.Invoke(ctx, "/proto.Notifier/NotifyAllUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotifierServer is the server API for Notifier service.
// All implementations must embed UnimplementedNotifierServer
// for forward compatibility
type NotifierServer interface {
	NotifyUser(context.Context, *NotifyUserRequest) (*NotifyResponse, error)
	NotifyUserBatch(context.Context, *NotifyUserBatchRequest) (*NotifyResponse, error)
	NotifyAllUser(context.Context, *NotifyRequest) (*NotifyResponse, error)
	mustEmbedUnimplementedNotifierServer()
}

// UnimplementedNotifierServer must be embedded to have forward compatible implementations.
type UnimplementedNotifierServer struct {
}

func (UnimplementedNotifierServer) NotifyUser(context.Context, *NotifyUserRequest) (*NotifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyUser not implemented")
}
func (UnimplementedNotifierServer) NotifyUserBatch(context.Context, *NotifyUserBatchRequest) (*NotifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyUserBatch not implemented")
}
func (UnimplementedNotifierServer) NotifyAllUser(context.Context, *NotifyRequest) (*NotifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyAllUser not implemented")
}
func (UnimplementedNotifierServer) mustEmbedUnimplementedNotifierServer() {}

// UnsafeNotifierServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotifierServer will
// result in compilation errors.
type UnsafeNotifierServer interface {
	mustEmbedUnimplementedNotifierServer()
}

func RegisterNotifierServer(s grpc.ServiceRegistrar, srv NotifierServer) {
	s.RegisterService(&Notifier_ServiceDesc, srv)
}

func _Notifier_NotifyUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).NotifyUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Notifier/NotifyUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).NotifyUser(ctx, req.(*NotifyUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifier_NotifyUserBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyUserBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).NotifyUserBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Notifier/NotifyUserBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).NotifyUserBatch(ctx, req.(*NotifyUserBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifier_NotifyAllUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).NotifyAllUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Notifier/NotifyAllUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).NotifyAllUser(ctx, req.(*NotifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Notifier_ServiceDesc is the grpc.ServiceDesc for Notifier service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Notifier_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Notifier",
	HandlerType: (*NotifierServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NotifyUser",
			Handler:    _Notifier_NotifyUser_Handler,
		},
		{
			MethodName: "NotifyUserBatch",
			Handler:    _Notifier_NotifyUserBatch_Handler,
		},
		{
			MethodName: "NotifyAllUser",
			Handler:    _Notifier_NotifyAllUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notify.proto",
}
