package gateServer

import (
	"go-snake/akmessage"
	"go-snake/app/in"
	"go-snake/common/messageBase"
	"go-snake/common/mixNet"
	"go-snake/common/tcpNet"
	"go-snake/common/webNet"
	"go-snake/gateServer/app"
	"math/rand"

	"github.com/Peakchen/xgameCommon/utils"

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
	messageBase.InitCodec(&utils.CodecProtobuf{})
	mixNet.NewSessionMgr(mixNet.GetApp())
}

func (this *Gate) Type() akmessage.ServerType {
	return akmessage.ServerType_Gate
}

func (this *Gate) Run(d *in.Input) {
	webNet.NewWebsocketSvr(d.WebHost)
	tcpNet.NewTcpServer(
		d.TCPHost,
		this.Type(),
		tcpNet.WithMessageHandler(tcpNet.ServerMsgProc))
}
