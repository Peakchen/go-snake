package gameServer

import (
	"go-snake/akmessage"
	"go-snake/app/in"
	"go-snake/app/application"
	"go-snake/common/akOrm"
	"go-snake/common/evtAsync"
	"go-snake/common/messageBase"
	"go-snake/common/mixNet"
	"go-snake/common/tcpNet"
	"go-snake/gameServer/app"
	"go-snake/gameServer/logic"
	"go-snake/gameServer/rpcBase"

	"github.com/Peakchen/xgameCommon/utils"

)

type Game struct {
	
}

func New(name string) *Game {
	
	application.SetAppName(name)
	
	return &Game{}

}

func (this *Game) Init() {
	//load config...
	//...
	app.Init()
	logic.Init()
	messageBase.InitCodec(&utils.CodecProtobuf{})
	mixNet.NewSessionMgr(mixNet.GetApp())
	evtAsync.NewMainEvtMgr()
}

func (this *Game) Type() akmessage.ServerType {
	return akmessage.ServerType_Game
}

func (this *Game) Run(d *in.Input) {

	akOrm.OpenDB(d.Scfg.MysqlUser, d.Scfg.MysqlPwd, d.Scfg.MysqlHost, d.Scfg.MysqlDataBase)

	rpcBase.RunRpc(d.Scfg.EtcdIP, d.Scfg.EtcdNodeIP)
	
	tcpNet.NewTcpClient(
		d.TCPHost,
		this.Type(),
		tcpNet.WithSSHeartBeat(messageBase.SS_HeatBeatMsg),
		tcpNet.WithMessageHandler(tcpNet.ServerMsgProc))

}
