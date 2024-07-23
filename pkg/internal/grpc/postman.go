package grpc

import (
	"context"
	"git.solsynth.dev/hydrogen/dealer/pkg/internal/services"
	"git.solsynth.dev/hydrogen/dealer/pkg/proto"
	"sync"
)

func (v *Server) DeliverNotification(ctx context.Context, request *proto.DeliverNotificationRequest) (*proto.DeliverResponse, error) {
	services.DealDeliveryTask(request)
	return &proto.DeliverResponse{}, nil
}

func (v *Server) DeliverNotificationBatch(ctx context.Context, request *proto.DeliverNotificationBatchRequest) (*proto.DeliverResponse, error) {
	var wg sync.WaitGroup
	for idx, provider := range request.GetProviders() {
		token := request.GetDeviceTokens()[idx]
		provider := provider
		go func() {
			wg.Add(1)
			services.DealDeliveryTask(&proto.DeliverNotificationRequest{
				Provider:    provider,
				DeviceToken: token,
				Notify:      request.GetNotify(),
			})
			wg.Done()
		}()
	}

	return &proto.DeliverResponse{}, nil
}

func (v *Server) DeliverEmail(ctx context.Context, request *proto.DeliverEmailRequest) (*proto.DeliverResponse, error) {
	services.DealDeliveryTask(request)
	return &proto.DeliverResponse{}, nil
}

func (v *Server) DeliverEmailBatch(ctx context.Context, request *proto.DeliverEmailBatchRequest) (*proto.DeliverResponse, error) {
	var wg sync.WaitGroup
	for _, to := range request.GetTo() {
		to := to
		go func() {
			wg.Add(1)
			services.DealDeliveryTask(&proto.DeliverEmailRequest{
				To:    to,
				Email: request.GetEmail(),
			})
			wg.Done()
		}()
	}

	return &proto.DeliverResponse{}, nil
}
