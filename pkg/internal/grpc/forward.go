package grpc

import (
	"context"
	"fmt"
	"git.solsynth.dev/hydrogen/dealer/pkg/hyper"
	"git.solsynth.dev/hydrogen/dealer/pkg/internal/directory"
	"google.golang.org/grpc"
	"time"
)

func forwardInvokeRequest[T any](serviceType string, executor func(context.Context, *grpc.ClientConn) (T, error)) (T, error) {
	var emptyResult T
	instance := directory.GetServiceInstance(serviceType)
	if instance == nil {
		return emptyResult, fmt.Errorf("no available service %s found", hyper.ServiceTypeAuthProvider)
	}

	conn, err := instance.GetGrpcConn()
	if err != nil {
		return emptyResult, fmt.Errorf("service is down: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	return executor(ctx, conn)
}
