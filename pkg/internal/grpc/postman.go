package grpc

import (
	"context"
	"git.solsynth.dev/hydrogen/dealer/pkg/internal/services"
	"git.solsynth.dev/hydrogen/dealer/pkg/proto"
)

func (v *Server) DeliverNotification(ctx context.Context, request *proto.DeliverNotificationRequest) (*proto.DeliverResponse, error) {
	services.PublishDeliveryTask(request)
	return &proto.DeliverResponse{}, nil
}

func (v *Server) DeliverNotificationBatch(ctx context.Context, request *proto.DeliverNotificationBatchRequest) (*proto.DeliverResponse, error) {
	for idx, provider := range request.GetProviders() {
		token := request.GetDeviceTokens()[idx]
		services.PublishDeliveryTask(&proto.DeliverNotificationRequest{
			Provider:    provider,
			DeviceToken: token,
			Notify:      request.GetNotify(),
		})
	}

	return &proto.DeliverResponse{}, nil
}

func (v *Server) DeliverEmail(ctx context.Context, request *proto.DeliverEmailRequest) (*proto.DeliverResponse, error) {
	services.PublishDeliveryTask(request)
	return &proto.DeliverResponse{}, nil
}

func (v *Server) DeliverEmailBatch(ctx context.Context, request *proto.DeliverEmailBatchRequest) (*proto.DeliverResponse, error) {
	for _, to := range request.GetTo() {
		services.PublishDeliveryTask(&proto.DeliverEmailRequest{
			To:    to,
			Email: request.GetEmail(),
		})
	}

	return &proto.DeliverResponse{}, nil
}
