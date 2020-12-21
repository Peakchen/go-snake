package loginServer

import (
	"go-snake/akmessage"
	"go-snake/app/in"
	"go-snake/common/mixNet"
	"go-snake/common/tcpNet"
	"go-snake/loginServer/app"
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
}

func (this *Login) Type() akmessage.ServerType {
	return akmessage.ServerType_Login
}

func (this *Login) Run(d *in.Input) {
	mixNet.NewSessionMgr(mixNet.GetApp())
	tcpNet.NewTcpClient(
		d.TCPHost,
		this.Type(),
		tcpNet.WithSSHeartBeat(tcpNet.SS_HeatBeatMsg),
		tcpNet.WithMessageHandler(tcpNet.ServerMsgProc))
}
