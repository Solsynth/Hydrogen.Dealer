// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: record.proto

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
	EventRecorder_RecordEvent_FullMethodName = "/proto.EventRecorder/RecordEvent"
)

// EventRecorderClient is the client API for EventRecorder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventRecorderClient interface {
	RecordEvent(ctx context.Context, in *RecordEventRequest, opts ...grpc.CallOption) (*RecordEventResponse, error)
}

type eventRecorderClient struct {
	cc grpc.ClientConnInterface
}

func NewEventRecorderClient(cc grpc.ClientConnInterface) EventRecorderClient {
	return &eventRecorderClient{cc}
}

func (c *eventRecorderClient) RecordEvent(ctx context.Context, in *RecordEventRequest, opts ...grpc.CallOption) (*RecordEventResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RecordEventResponse)
	err := c.cc.Invoke(ctx, EventRecorder_RecordEvent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventRecorderServer is the server API for EventRecorder service.
// All implementations must embed UnimplementedEventRecorderServer
// for forward compatibility
type EventRecorderServer interface {
	RecordEvent(context.Context, *RecordEventRequest) (*RecordEventResponse, error)
	mustEmbedUnimplementedEventRecorderServer()
}

// UnimplementedEventRecorderServer must be embedded to have forward compatible implementations.
type UnimplementedEventRecorderServer struct {
}

func (UnimplementedEventRecorderServer) RecordEvent(context.Context, *RecordEventRequest) (*RecordEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecordEvent not implemented")
}
func (UnimplementedEventRecorderServer) mustEmbedUnimplementedEventRecorderServer() {}

// UnsafeEventRecorderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventRecorderServer will
// result in compilation errors.
type UnsafeEventRecorderServer interface {
	mustEmbedUnimplementedEventRecorderServer()
}

func RegisterEventRecorderServer(s grpc.ServiceRegistrar, srv EventRecorderServer) {
	s.RegisterService(&EventRecorder_ServiceDesc, srv)
}

func _EventRecorder_RecordEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventRecorderServer).RecordEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EventRecorder_RecordEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventRecorderServer).RecordEvent(ctx, req.(*RecordEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EventRecorder_ServiceDesc is the grpc.ServiceDesc for EventRecorder service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventRecorder_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.EventRecorder",
	HandlerType: (*EventRecorderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecordEvent",
			Handler:    _EventRecorder_RecordEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "record.proto",
}
