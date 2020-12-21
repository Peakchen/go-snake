package gameServer

import (
	"go-snake/akmessage"
	"go-snake/app/in"
	"go-snake/common/mixNet"
	"go-snake/common/tcpNet"
	"go-snake/gameServer/app"
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
}

func (this *Game) Type() akmessage.ServerType {
	return akmessage.ServerType_Game
}

func (this *Game) Run(d *in.Input) {
	mixNet.NewSessionMgr(mixNet.GetApp())
	tcpNet.NewTcpClient(
		d.TCPHost,
		this.Type(),
		tcpNet.WithSSHeartBeat(tcpNet.SS_HeatBeatMsg),
		tcpNet.WithMessageHandler(tcpNet.ServerMsgProc))
}
