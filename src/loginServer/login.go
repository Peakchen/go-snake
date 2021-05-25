package loginServer

import (
	"go-snake/akmessage"
	"go-snake/app/in"
	"go-snake/common/akOrm"
	"go-snake/common/evtAsync"
	"go-snake/common/messageBase"
	"go-snake/common/mixNet"
	"go-snake/common/tcpNet"
	"go-snake/loginServer/app"
	"go-snake/model/sdk_wechat"
	"go-snake/loginServer/logic"
	"go-snake/loginServer/rpcBase"
	"go-snake/app/application"
	"github.com/Peakchen/xgameCommon/utils"
)

type Login struct {
}

func New(name string) *Login {
	
	application.SetAppName(name)

	return &Login{}
}

func (this *Login) Init() {
	//load config...
	//...
	app.Init()
	messageBase.InitCodec(&utils.CodecProtobuf{})
	mixNet.NewSessionMgr(mixNet.GetApp())
	evtAsync.NewMainEvtMgr()

}

func (this *Login) Type() akmessage.ServerType {
	return akmessage.ServerType_Login
}

func (this *Login) Run(d *in.Input) {
	//启动db链接
	akOrm.OpenDB(d.Scfg.MysqlUser, d.Scfg.MysqlPwd, d.Scfg.MysqlHost, d.Scfg.MysqlDataBase)
	//db数据加载至内存
	logic.LoadDB()
	//初始化etcd rpc
	rpcBase.RunRpc(d.Scfg.EtcdIP, d.Scfg.EtcdNodeIP)
	//启动tcp链接
	tcpNet.NewTcpClient(
		d.TCPHost,
		this.Type(),
		tcpNet.WithSSHeartBeat(messageBase.SS_HeatBeatMsg),
		tcpNet.WithMessageHandler(tcpNet.ServerMsgProc))
	//判断第三方服务器启动
	if d.Scfg.HasWechat {
		sdk_wechat.Run(d.Scfg.WebHttp, d.Scfg.AppID, d.Scfg.AppSecret)
	}
}
