package loginServer

import (
	"go-snake/akmessage"
	"go-snake/app/in"
	"go-snake/common/messageBase"
	"go-snake/common/mixNet"
	"go-snake/common/tcpNet"
	"go-snake/loginServer/app"

	"github.com/Peakchen/xgameCommon/utils"
)

type Login struct {
}

func New() *Login {
	return &Login{}
}

func (this *Login) Init() {
	//load config...
	//...
	app.Init()
	messageBase.InitCodec(&utils.CodecProtobuf{})
	mixNet.NewSessionMgr(mixNet.GetApp())
}

func (this *Login) Type() akmessage.ServerType {
	return akmessage.ServerType_Login
}

func (this *Login) Run(d *in.Input) {
	tcpNet.NewTcpClient(
		d.TCPHost,
		this.Type(),
		tcpNet.WithSSHeartBeat(tcpNet.SS_HeatBeatMsg),
		tcpNet.WithMessageHandler(tcpNet.ServerMsgProc))
}
