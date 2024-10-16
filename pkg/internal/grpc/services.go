package grpc

import (
	"context"
	"fmt"
	"time"

	"git.solsynth.dev/hydrogen/dealer/pkg/internal/directory"
	"git.solsynth.dev/hydrogen/dealer/pkg/proto"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"
)

func convertServiceToInfo(in *directory.ServiceInstance) *proto.ServiceInfo {
	if in == nil {
		return nil
	}
	return &proto.ServiceInfo{
		Id:       in.ID,
		Type:     in.Type,
		Label:    in.Label,
		GrpcAddr: in.GrpcAddr,
		HttpAddr: in.HttpAddr,
	}
}

func (v *Server) GetService(ctx context.Context, request *proto.GetServiceRequest) (*proto.GetServiceResponse, error) {
	if request.Id != nil {
		out := directory.GetServiceInstance(request.GetId())
		return &proto.GetServiceResponse{
			Data: convertServiceToInfo(out),
		}, nil
	}
	if request.Type != nil {
		out := directory.GetServiceInstanceByType(request.GetType())
		return &proto.GetServiceResponse{
			Data: convertServiceToInfo(out),
		}, nil
	}
	return nil, fmt.Errorf("no filter condition is provided")
}

func (v *Server) ListService(ctx context.Context, request *proto.ListServiceRequest) (*proto.ListServiceResponse, error) {
	var out []*directory.ServiceInstance
	if request.Type != nil {
		out = directory.ListServiceInstanceByType(request.GetType())
	} else {
		out = directory.ListServiceInstance()
	}
	return &proto.ListServiceResponse{
		Data: lo.Map(out, func(item *directory.ServiceInstance, index int) *proto.ServiceInfo {
			return convertServiceToInfo(item)
		}),
	}, nil
}

func (v *Server) AddService(ctx context.Context, info *proto.ServiceInfo) (*proto.AddServiceResponse, error) {
	in := &directory.ServiceInstance{
		ID:       info.GetId(),
		Type:     info.GetType(),
		Label:    info.GetLabel(),
		GrpcAddr: info.GetGrpcAddr(),
		HttpAddr: info.HttpAddr,
	}
	directory.AddServiceInstance(in)
	log.Info().Str("id", info.GetId()).Str("label", info.GetLabel()).Msg("New service added.")
	return &proto.AddServiceResponse{
		IsSuccess: true,
	}, nil
}

func (v *Server) RemoveService(ctx context.Context, request *proto.RemoveServiceRequest) (*proto.RemoveServiceResponse, error) {
	directory.RemoveServiceInstance(request.GetId())
	log.Info().Str("id", request.GetId()).Msg("A service removed.")
	return &proto.RemoveServiceResponse{
		IsSuccess: true,
	}, nil
}

func (v *Server) BroadcastDeletion(ctx context.Context, request *proto.DeletionRequest) (*proto.DeletionResponse, error) {
	for _, service := range directory.ListServiceInstance() {
		conn, err := service.GetGrpcConn()
		if err != nil {
			continue
		}
		pc := proto.NewServiceDirectoryClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		_, err = pc.BroadcastDeletion(ctx, request)
		cancel()
	}
	return &proto.DeletionResponse{}, nil
}
