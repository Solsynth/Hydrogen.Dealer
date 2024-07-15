package grpc

import (
	"context"
	"fmt"
	"time"

	"git.solsynth.dev/hydrogen/dealer/pkg/hyper"
	"git.solsynth.dev/hydrogen/dealer/pkg/internal/directory"
	"git.solsynth.dev/hydrogen/dealer/pkg/proto"
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

func (v *Server) EnsurePermGranted(ctx context.Context, request *proto.CheckPermRequest) (*proto.CheckPermResponse, error) {
	instance := directory.GetServiceInstanceByType(hyper.ServiceTypeAuthProvider)
	if instance == nil {
		return &proto.CheckPermResponse{}, fmt.Errorf("no available service %s found", hyper.ServiceTypeAuthProvider)
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

func (v *Server) EnsureUserPermGranted(ctx context.Context, request *proto.CheckUserPermRequest) (*proto.CheckUserPermResponse, error) {
	instance := directory.GetServiceInstance(hyper.ServiceTypeAuthProvider)
	if instance == nil {
		return &proto.CheckUserPermResponse{}, fmt.Errorf("no available service %s found", hyper.ServiceTypeAuthProvider)
	}

	conn, err := instance.GetGrpcConn()
	if err != nil {
		return nil, fmt.Errorf("service is down: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	out, err := proto.NewAuthClient(conn).EnsureUserPermGranted(ctx, request)
	return out, err
}
