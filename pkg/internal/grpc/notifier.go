package grpc

import (
	"context"
	"git.solsynth.dev/hydrogen/dealer/pkg/hyper"
	"git.solsynth.dev/hydrogen/dealer/pkg/proto"
	"google.golang.org/grpc"
)

func (v *Server) NotifyUser(ctx context.Context, request *proto.NotifyUserRequest) (*proto.NotifyResponse, error) {
	return forwardInvokeRequest(
		hyper.ServiceTypeAuthProvider,
		func(ctx context.Context, conn *grpc.ClientConn) (*proto.NotifyResponse, error) {
			out, err := proto.NewNotifierClient(conn).NotifyUser(ctx, request)
			return out, err
		},
	)
}

func (v *Server) NotifyAllUser(ctx context.Context, request *proto.NotifyRequest) (*proto.NotifyResponse, error) {
	return forwardInvokeRequest(
		hyper.ServiceTypeAuthProvider,
		func(ctx context.Context, conn *grpc.ClientConn) (*proto.NotifyResponse, error) {
			out, err := proto.NewNotifierClient(conn).NotifyAllUser(ctx, request)
			return out, err
		},
	)
}
