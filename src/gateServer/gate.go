package gateServer

import (
	"go-snake/akmessage"
	"go-snake/app/in"
	"go-snake/common/mixNet"
	"go-snake/common/tcpNet"
	"go-snake/common/webNet"
	"go-snake/gateServer/app"
	"math/rand"

	"github.com/Peakchen/xgameCommon/aktime"
)

func init() {
	t := aktime.Now().Unix()
	s := rand.NewSource(t)
	rand.New(s).Seed(t)
}

type Gate struct {
}

func New() *Gate {
	return &Gate{}
}

func (this *Gate) Init() {
	app.Init()
}

func (this *Gate) Type() akmessage.ServerType {
	return akmessage.ServerType_Gate
}

func (this *Gate) Run(d *in.Input) {
	mixNet.NewSessionMgr(mixNet.GetApp())

	webNet.NewWebsocketSvr(d.WebHost)
	tcpNet.NewTcpServer(
		d.TCPHost,
		this.Type(),
		tcpNet.WithMessageHandler(tcpNet.ServerMsgProc))
}
