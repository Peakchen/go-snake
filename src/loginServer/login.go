package loginServer

import (
	"go-snake/akmessage"
	"go-snake/app/in"
	"go-snake/common/akOrm"
	"go-snake/common/messageBase"
	"go-snake/common/mixNet"
	"go-snake/common/tcpNet"
	"go-snake/loginServer/app"

	_ "go-snake/loginServer/logic"

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
	akOrm.OpenDB(d.Scfg.MysqlUser, d.Scfg.MysqlPwd, d.Scfg.MysqlHost, d.Scfg.MysqlDataBase)
	tcpNet.NewTcpClient(
		d.TCPHost,
		this.Type(),
		tcpNet.WithSSHeartBeat(messageBase.SS_HeatBeatMsg),
		tcpNet.WithMessageHandler(tcpNet.ServerMsgProc))
}
