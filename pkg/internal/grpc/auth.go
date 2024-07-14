package grpc

import (
	"context"
	"fmt"
	"git.solsynth.dev/hydrogen/dealer/pkg/hyper"
	"git.solsynth.dev/hydrogen/dealer/pkg/internal/directory"
	"git.solsynth.dev/hydrogen/dealer/pkg/proto"
	"time"
)

func (v *Server) Authenticate(ctx context.Context, request *proto.AuthRequest) (*proto.AuthReply, error) {
	instance := directory.GetServiceInstanceByType(hyper.ServiceTypeAuthProvider)
	if instance == nil {
		return &proto.AuthReply{}, fmt.Errorf("no available service %s found", hyper.ServiceTypeAuthProvider)
	}

	conn, err := instance.GetGrpcConn()
	if err != nil {
		return nil, fmt.Errorf("service is down: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	out, err := proto.NewAuthClient(conn).Authenticate(ctx, request)
	return out, err
}

func (v *Server) EnsurePermGranted(ctx context.Context, request *proto.CheckPermRequest) (*proto.CheckPermReply, error) {
	instance := directory.GetServiceInstanceByType(hyper.ServiceTypeAuthProvider)
	if instance == nil {
		return &proto.CheckPermReply{}, fmt.Errorf("no available service %s found", hyper.ServiceTypeAuthProvider)
	}

	conn, err := instance.GetGrpcConn()
	if err != nil {
		return nil, fmt.Errorf("service is down: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	out, err := proto.NewAuthClient(conn).EnsurePermGranted(ctx, request)
	return out, err
}
