package app

import (
	"go-snake/common/messageBase"
	"go-snake/common/mixNet"
	"go-snake/common/tcpNet"
)

func Init() {
	mixNet.SetApp(NewApp())
}

type S2SContext struct {
	roles   uint32
	session *tcpNet.TcpSession
}

//gate2 <-> game server
type GameApp struct {
	roles uint32

	s2s map[string]*S2SContext
}

func NewApp() *GameApp {
	return &GameApp{
		s2s: make(map[string]*S2SContext),
	}
}

func (this *GameApp) Online(nt messageBase.NetType, sess interface{}) {
	switch nt {
	case messageBase.Net_WS:

	case messageBase.Net_TCP:
		s := sess.(*tcpNet.TcpSession)
		this.s2s[s.GetSessionID()] = &S2SContext{
			session: s,
		}
	}
}

func (this *GameApp) Offline(nt messageBase.NetType, id string) {
	switch nt {
	case messageBase.Net_WS:

	case messageBase.Net_TCP:
		delete(this.s2s, id)
	}
}

func (this *GameApp) Bind(id int64) {

}

func (this *GameApp) SendInner(sid string, id uint32, data []byte) {

}

func (this *GameApp) SendClient(sid string, id uint32, data []byte) {

}

func (this *GameApp) Handler(sid string, data []byte) {

}
