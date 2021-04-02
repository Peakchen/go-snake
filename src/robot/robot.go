package robot

import (
	"go-snake/akmessage"
	"go-snake/app/in"
	"go-snake/common/messageBase"
	"go-snake/common/mixNet"
	"go-snake/robot/manager"
	_ "go-snake/robot/model"
	"go-snake/robot/option"
	"go-snake/robot/wscli"
	"time"
	"go-snake/common/myNats"
	"github.com/Peakchen/xgameCommon/utils"
	"go-snake/app/application"
)

type Robot struct {
}

func New(name string) *Robot {
	
	application.SetAppName(name)

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

	myNats.Register(d.Scfg.NatsHost, utils.ENCodecType_Pb)

	for i := 0; i < d.Clis; i++ {
		time.Sleep(10 * time.Millisecond)

		wscli.NewClient(
			i,
			d.WebHost,
			option.WithModelsRun(manager.RangeModels),
			option.WithModelRecv(manager.RangeRecv),
		)
	}
}
