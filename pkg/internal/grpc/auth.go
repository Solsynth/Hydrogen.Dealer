package grpc

import (
	"context"
	"git.solsynth.dev/hydrogen/dealer/pkg/hyper"
	"git.solsynth.dev/hydrogen/dealer/pkg/proto"
	"google.golang.org/grpc"
)

func (v *Server) Authenticate(ctx context.Context, request *proto.AuthRequest) (*proto.AuthReply, error) {
	return forwardInvokeRequest(
		hyper.ServiceTypeAuthProvider,
		func(ctx context.Context, conn *grpc.ClientConn) (*proto.AuthReply, error) {
			out, err := proto.NewAuthClient(conn).Authenticate(ctx, request)
			return out, err
		},
	)
}

func (v *Server) EnsurePermGranted(ctx context.Context, request *proto.CheckPermRequest) (*proto.CheckPermResponse, error) {
	return forwardInvokeRequest(
		hyper.ServiceTypeAuthProvider,
		func(ctx context.Context, conn *grpc.ClientConn) (*proto.CheckPermResponse, error) {
			out, err := proto.NewAuthClient(conn).EnsurePermGranted(ctx, request)
			return out, err
		},
	)
}

func (v *Server) EnsureUserPermGranted(ctx context.Context, request *proto.CheckUserPermRequest) (*proto.CheckUserPermResponse, error) {
	return forwardInvokeRequest(
		hyper.ServiceTypeAuthProvider,
		func(ctx context.Context, conn *grpc.ClientConn) (*proto.CheckUserPermResponse, error) {
			out, err := proto.NewAuthClient(conn).EnsureUserPermGranted(ctx, request)
			return out, err
		},
	)
}
