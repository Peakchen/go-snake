package app

import (
	"go-snake/common/messageBase"
	"go-snake/common/mixNet"
	"go-snake/common/tcpNet"
)

func Init() {
	mixNet.SetApp(NewApp())
}

//gate2 <-> game server
type GameApp struct {
	roles uint32

	session *tcpNet.TcpSession
}

func NewApp() *GameApp {
	return &GameApp{
		session: nil,
	}
}

func (this *GameApp) Online(nt messageBase.NetType, sess interface{}) {
	switch nt {
	case messageBase.Net_WS:

	case messageBase.Net_TCP:
		this.session = sess.(*tcpNet.TcpSession)
		this.session.SendMsg(messageBase.GetActorRegisterReq(this.session.GetSessionID(), this.session.GetType()))
	}
}

func (this *GameApp) Offline(nt messageBase.NetType, id string) {
	switch nt {
	case messageBase.Net_WS:

	case messageBase.Net_TCP:
		this.session = nil
	}
}

func (this *GameApp) Bind(id int64) {

}

func (this *GameApp) CS_SendInner(sid string, id uint32, data []byte) {

}

func (this *GameApp) SendClient(sid string, id uint32, data []byte) {

}

func (this *GameApp) Handler(sid string, data []byte) {

}

func (this *GameApp) SS_SendInner(sid string, id uint32, data []byte) {

}
