package chat

import (
	"go-snake/akmessage"
	"go-snake/app/in"
	"go-snake/app/application"
	"go-snake/gameServer/app"
	"go-snake/chat/rpcBase"
	"go-snake/common/tcpNet"
	"go-snake/common/messageBase"
)

type Chat struct {
}

func New(name string) *Chat {
	
	application.SetAppName(name)

	return &Chat{}
}

func (this *Chat) Init() {

	app.Init()

}

func (this *Chat) Type() akmessage.ServerType {
	return akmessage.ServerType_Chat
}

func (this *Chat) Run(d *in.Input) {

	rpcBase.RunRpc(d.Scfg.EtcdIP, d.Scfg.EtcdNodeIP)

	tcpNet.NewTcpClient(
		d.TCPHost,
		this.Type(),
		tcpNet.WithSSHeartBeat(messageBase.SS_HeatBeatMsg),
		tcpNet.WithMessageHandler(tcpNet.ServerMsgProc))

}
