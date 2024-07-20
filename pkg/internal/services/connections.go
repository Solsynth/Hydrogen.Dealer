package services

import (
	"context"
	"sync"

	"git.solsynth.dev/hydrogen/dealer/pkg/hyper"
	"git.solsynth.dev/hydrogen/dealer/pkg/internal/directory"
	"git.solsynth.dev/hydrogen/dealer/pkg/internal/models"
	"git.solsynth.dev/hydrogen/dealer/pkg/proto"
	"github.com/gofiber/contrib/websocket"
)

var (
	wsMutex sync.Mutex
	wsConn  = make(map[uint]map[*websocket.Conn]bool)
)

func ClientRegister(user models.Account, conn *websocket.Conn) {
	wsMutex.Lock()
	if wsConn[user.ID] == nil {
		wsConn[user.ID] = make(map[*websocket.Conn]bool)
	}
	wsConn[user.ID][conn] = true
	wsMutex.Unlock()

	pc, err := directory.GetServiceInstanceByType(hyper.ServiceTypeAuthProvider).GetGrpcConn()
	if err == nil {
		proto.NewStreamControllerClient(pc).EmitStreamEvent(context.Background(), &proto.StreamEventRequest{
			Event:  "ClientRegister",
			UserId: uint64(user.ID),
		})
	}
}

func ClientUnregister(user models.Account, conn *websocket.Conn) {
	wsMutex.Lock()
	if wsConn[user.ID] == nil {
		wsConn[user.ID] = make(map[*websocket.Conn]bool)
	}
	delete(wsConn[user.ID], conn)
	wsMutex.Unlock()

	pc, err := directory.GetServiceInstanceByType(hyper.ServiceTypeAuthProvider).GetGrpcConn()
	if err == nil {
		proto.NewStreamControllerClient(pc).EmitStreamEvent(context.Background(), &proto.StreamEventRequest{
			Event:  "ClientUnregister",
			UserId: uint64(user.ID),
		})
	}
}

func ClientCount(uid uint) int {
	return len(wsConn[uid])
}

func WebsocketPush(uid uint, body []byte) (count int, success int, errs []error) {
	for conn := range wsConn[uid] {
		if err := conn.WriteMessage(1, body); err != nil {
			errs = append(errs, err)
		} else {
			success++
		}
		count++
	}
	return
}

func WebsocketPushBatch(uidList []uint, body []byte) (count int, success int, errs []error) {
	for _, uid := range uidList {
		for conn := range wsConn[uid] {
			if err := conn.WriteMessage(1, body); err != nil {
				errs = append(errs, err)
			} else {
				success++
			}
			count++
		}
	}
	return
}
