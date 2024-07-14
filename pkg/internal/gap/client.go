package gap

import (
	"git.solsynth.dev/hydrogen/dealer/pkg/hyper"
	"git.solsynth.dev/hydrogen/dealer/pkg/proto"
	"github.com/spf13/viper"
)

var H *hyper.HyperConn

func NewHyperClient(info *proto.ServiceInfo) (err error) {
	H, err = hyper.NewHyperConn(viper.GetString("dealer.addr"), info)
	H.KeepRegisterService()
	return
}
