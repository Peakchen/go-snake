package online

import (
	"ak-remote/common/messageBase"
	"ak-remote/common/myTcpSocket"
	"ak-remote/loginServer/app"
)

func init() {
	app.SetApp(NewApp())
}

type S2SContext struct {
	roles   uint32
	session *myTcpSocket.AkTcpSession
}

//gate2 <-> game server
type GameApp struct {
	roles uint32

	s2s map[string]*S2SContext
}

func NewApp() *GameApp {
	app = &GameApp{
		s2s: make(map[string]*S2SContext),
	}
	return app
}

func (this *GameApp) Online(nt messageBase.NetType, sess interface{}) {
	switch nt {
	case messageBase.Net_WS:

	case messageBase.Net_TCP:
		s := sess.(*myTcpSocket.AkTcpSession)
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
