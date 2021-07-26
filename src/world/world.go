package world

import (
	"go-snake/akmessage"
	"go-snake/app/in"
	"go-snake/app/application"
	"go-snake/common/akOrm"
	"go-snake/common/evtAsync"
	"go-snake/common/messageBase"
	"go-snake/common/mixNet"
	"go-snake/common/tcpNet"
	"go-snake/world/app"
	"go-snake/world/logic"
	"go-snake/world/rpcBase"

	"github.com/Peakchen/xgameCommon/utils"

)


type World struct {
	
}

func New(name string) *World {
	
	application.SetAppName(name)
	
	return &World{}

}

func (this *World) Init() {
	//load config...
	//...
	app.Init()
	logic.Init()
	logic.LoadTab()
	messageBase.InitCodec(&utils.CodecProtobuf{})
	mixNet.NewSessionMgr(mixNet.GetApp())
	evtAsync.NewMainEvtMgr()
	
}

func (this *World) Type() akmessage.ServerType {
	return akmessage.ServerType_World
}

func (this *World) Run(d *in.Input) {

	akOrm.OpenDB(d.Scfg.MysqlUser, d.Scfg.MysqlPwd, d.Scfg.MysqlHost, d.Scfg.MysqlDataBase)

	rpcBase.RunRpc(d.Scfg.EtcdIP, d.Scfg.EtcdNodeIP)
	
	tcpNet.NewTcpClient(
		d.TCPHost,
		this.Type(),
		tcpNet.WithSSHeartBeat(messageBase.SS_HeatBeatMsg),
		tcpNet.WithMessageHandler(tcpNet.ServerMsgProc))

}