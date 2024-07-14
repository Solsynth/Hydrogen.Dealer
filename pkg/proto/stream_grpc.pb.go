// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: stream.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	StreamController_CountStreamConnection_FullMethodName = "/proto.StreamController/CountStreamConnection"
	StreamController_PushStream_FullMethodName            = "/proto.StreamController/PushStream"
)

// StreamControllerClient is the client API for StreamController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamControllerClient interface {
	CountStreamConnection(ctx context.Context, in *CountConnectionRequest, opts ...grpc.CallOption) (*CountConnectionResponse, error)
	PushStream(ctx context.Context, in *PushStreamRequest, opts ...grpc.CallOption) (*PushStreamResponse, error)
}

type streamControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamControllerClient(cc grpc.ClientConnInterface) StreamControllerClient {
	return &streamControllerClient{cc}
}

func (c *streamControllerClient) CountStreamConnection(ctx context.Context, in *CountConnectionRequest, opts ...grpc.CallOption) (*CountConnectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CountConnectionResponse)
	err := c.cc.Invoke(ctx, StreamController_CountStreamConnection_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamControllerClient) PushStream(ctx context.Context, in *PushStreamRequest, opts ...grpc.CallOption) (*PushStreamResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PushStreamResponse)
	err := c.cc.Invoke(ctx, StreamController_PushStream_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StreamControllerServer is the server API for StreamController service.
// All implementations must embed UnimplementedStreamControllerServer
// for forward compatibility
type StreamControllerServer interface {
	CountStreamConnection(context.Context, *CountConnectionRequest) (*CountConnectionResponse, error)
	PushStream(context.Context, *PushStreamRequest) (*PushStreamResponse, error)
	mustEmbedUnimplementedStreamControllerServer()
}

// UnimplementedStreamControllerServer must be embedded to have forward compatible implementations.
type UnimplementedStreamControllerServer struct {
}

func (UnimplementedStreamControllerServer) CountStreamConnection(context.Context, *CountConnectionRequest) (*CountConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountStreamConnection not implemented")
}
func (UnimplementedStreamControllerServer) PushStream(context.Context, *PushStreamRequest) (*PushStreamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushStream not implemented")
}
func (UnimplementedStreamControllerServer) mustEmbedUnimplementedStreamControllerServer() {}

// UnsafeStreamControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamControllerServer will
// result in compilation errors.
type UnsafeStreamControllerServer interface {
	mustEmbedUnimplementedStreamControllerServer()
}

func RegisterStreamControllerServer(s grpc.ServiceRegistrar, srv StreamControllerServer) {
	s.RegisterService(&StreamController_ServiceDesc, srv)
}

func _StreamController_CountStreamConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamControllerServer).CountStreamConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamController_CountStreamConnection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamControllerServer).CountStreamConnection(ctx, req.(*CountConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamController_PushStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamControllerServer).PushStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamController_PushStream_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamControllerServer).PushStream(ctx, req.(*PushStreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StreamController_ServiceDesc is the grpc.ServiceDesc for StreamController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.StreamController",
	HandlerType: (*StreamControllerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CountStreamConnection",
			Handler:    _StreamController_CountStreamConnection_Handler,
		},
		{
			MethodName: "PushStream",
			Handler:    _StreamController_PushStream_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stream.proto",
}
