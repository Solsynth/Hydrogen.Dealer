package hyper

import (
	"context"
	"git.solsynth.dev/hydrogen/dealer/pkg/proto"
	"time"
)

func (v *HyperConn) RecordAuditLog(user uint, action, target, ip, ua string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	in, err := v.GetServiceGrpcConn(ServiceTypeAuthProvider)
	if err != nil {
		return err
	}

	_, err = proto.NewEventRecorderClient(in).RecordEvent(ctx, &proto.RecordEventRequest{
		UserId:    uint64(user),
		Action:    action,
		Target:    target,
		Ip:        ip,
		UserAgent: ua,
	})
	if err != nil {
		return err
	}

	return nil
}
