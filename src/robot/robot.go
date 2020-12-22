package robot

import (
	"go-snake/akmessage"
	"go-snake/app/in"
	"go-snake/common/messageBase"
	"go-snake/common/mixNet"
	_ "go-snake/robot/login"
	"go-snake/robot/manager"
	"go-snake/robot/wscli"

	"github.com/Peakchen/xgameCommon/utils"
)

type Robot struct {
}

func New() *Robot {
	return &Robot{}
}

func (this *Robot) Init() {
	//load config...
	//...
	messageBase.InitCodec(&utils.CodecProtobuf{})
	mixNet.NewSessionMgr(mixNet.GetApp())
}

func (this *Robot) Type() akmessage.ServerType {
	return akmessage.ServerType_Robot
}

func (this *Robot) Run(d *in.Input) {
	for i := 0; i < d.Clis; i++ {
		wscli.NewClient(d.WebHost, manager.RangeModels)
	}
}
