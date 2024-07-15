package grpc

import (
	"context"
	"fmt"
	"time"

	"git.solsynth.dev/hydrogen/dealer/pkg/hyper"
	"git.solsynth.dev/hydrogen/dealer/pkg/internal/directory"
	"git.solsynth.dev/hydrogen/dealer/pkg/proto"
)

func (v *Server) RecordEvent(ctx context.Context, request *proto.RecordEventRequest) (*proto.RecordEventResponse, error) {
	instance := directory.GetServiceInstance(hyper.ServiceTypeAuthProvider)
	if instance == nil {
		return &proto.RecordEventResponse{}, fmt.Errorf("no available service %s found", hyper.ServiceTypeAuthProvider)
	}

	conn, err := instance.GetGrpcConn()
	if err != nil {
		return nil, fmt.Errorf("service is down: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	out, err := proto.NewEventRecorderClient(conn).RecordEvent(ctx, request)
	return out, err
}
