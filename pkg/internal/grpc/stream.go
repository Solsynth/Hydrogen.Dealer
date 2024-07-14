package grpc

import (
	"context"
	"fmt"
	"git.solsynth.dev/hydrogen/dealer/pkg/internal/services"
	"git.solsynth.dev/hydrogen/dealer/pkg/proto"
)

func (v *Server) PushStream(ctx context.Context, request *proto.PushStreamRequest) (*proto.PushStreamResponse, error) {
	cnt, success, errs := services.WebsocketPush(uint(request.GetUserId()), request.GetBody())
	if len(errs) > 0 {
		// Partial fail
		return &proto.PushStreamResponse{
			IsAllSuccess:  false,
			AffectedCount: int64(success),
			FailedCount:   int64(cnt - success),
		}, nil
	} else if cnt > 0 && success == 0 {
		// All fail
		return nil, fmt.Errorf("all push request failed: %v", errs)
	}

	return &proto.PushStreamResponse{
		IsAllSuccess:  true,
		AffectedCount: int64(success),
		FailedCount:   int64(cnt - success),
	}, nil
}
