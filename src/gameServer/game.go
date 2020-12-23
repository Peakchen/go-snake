package gameServer

import (
	"go-snake/akmessage"
	"go-snake/app/in"
	"go-snake/common/akOrm"
	"go-snake/common/messageBase"
	"go-snake/common/mixNet"
	"go-snake/common/tcpNet"
	"go-snake/gameServer/app"

	_ "go-snake/gameServer/logic"

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
	akOrm.OpenDB(d.Scfg.MysqlUser, d.Scfg.MysqlPwd, d.Scfg.MysqlHost, d.Scfg.MysqlDataBase)
	tcpNet.NewTcpClient(
		d.TCPHost,
		this.Type(),
		tcpNet.WithSSHeartBeat(messageBase.SS_HeatBeatMsg),
		tcpNet.WithMessageHandler(tcpNet.ServerMsgProc))
}
