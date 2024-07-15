package grpc

import (
	"context"
	"git.solsynth.dev/hydrogen/dealer/pkg/hyper"
	"git.solsynth.dev/hydrogen/dealer/pkg/proto"
	"google.golang.org/grpc"
)

func (v *Server) RecordEvent(ctx context.Context, request *proto.RecordEventRequest) (*proto.RecordEventResponse, error) {
	return forwardInvokeRequest(
		hyper.ServiceTypeAuthProvider,
		func(ctx context.Context, conn *grpc.ClientConn) (*proto.RecordEventResponse, error) {
			out, err := proto.NewEventRecorderClient(conn).RecordEvent(ctx, request)
			return out, err
		},
	)
}
