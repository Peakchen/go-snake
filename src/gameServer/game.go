package gameServer

import (
	"go-snake/akmessage"
	"go-snake/app/in"
	"go-snake/common/messageBase"
	"go-snake/common/mixNet"
	"go-snake/common/tcpNet"
	"go-snake/gameServer/app"

	"github.com/Peakchen/xgameCommon/utils"
)

type Game struct {
}

func New() *Game {
	return &Game{}
}

func (this *Game) Init() {
	//load config...
	//...
	app.Init()
	messageBase.InitCodec(&utils.CodecProtobuf{})
	mixNet.NewSessionMgr(mixNet.GetApp())
}

func (this *Game) Type() akmessage.ServerType {
	return akmessage.ServerType_Game
}

func (this *Game) Run(d *in.Input) {
	tcpNet.NewTcpClient(
		d.TCPHost,
		this.Type(),
		tcpNet.WithSSHeartBeat(tcpNet.SS_HeatBeatMsg),
		tcpNet.WithMessageHandler(tcpNet.ServerMsgProc))
}
