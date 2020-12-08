package app

import (
	"ak-remote/common/messageBase"
	"ak-remote/common/myTcpSocket"
	"ak-remote/common/myWebSocket"

	"github.com/Peakchen/xgameCommon/akLog"
)

type S2SContext struct {
	roles   uint32
	session *myTcpSocket.AkTcpSession
}

type C2SContext struct {
	SID     string
	session *myWebSocket.WebSession
}

type GateApp struct {
	roles uint32
	c2s   map[string]*C2SContext
	s2s   map[string]*S2SContext
	o2i   map[string]string
}

var (
	app *GateApp
)

func New() *GateApp {
	app = &GateApp{
		c2s: make(map[string]*C2SContext),
		s2s: make(map[string]*S2SContext),
		o2i: make(map[string]string),
	}
	return app
}

func GetApp() *GateApp {
	return app
}

// rule 1: get max roles server and role not big equal 5w
func (this *GateApp) getMaxRoleSvr() string {
	var max uint32
	var dst string
	for id, c := range this.s2s {
		if c.roles > max && c.roles < 50000 {
			max = c.roles
			dst = id
		}
	}
	return dst
}

func (this *GateApp) Online(nt messageBase.NetType, sess interface{}) {
	switch nt {
	case messageBase.Net_WS:
		if len(this.s2s) == 0 {
			akLog.Fail("not inner server can be connect.")
			return
		}

		selectID := this.getMaxRoleSvr()
		c, ok := this.s2s[selectID]
		if ok {
			this.roles++

			c.roles++

			s := sess.(*myWebSocket.WebSession)
			this.c2s[s.GetSessionID()] = &C2SContext{
				SID:     selectID,
				session: s,
			}
		}

	case messageBase.Net_TCP:
		s := sess.(*myTcpSocket.AkTcpSession)
		this.s2s[s.GetSessionID()] = &S2SContext{
			session: s,
		}
	}
}

func (this *GateApp) Offline(nt messageBase.NetType, id string) {
	switch nt {
	case messageBase.Net_WS:
		this.roles--
		delete(this.c2s, id)
	case messageBase.Net_TCP:
		delete(this.s2s, id)
	}
}

//c->gate1<=>gate2->s
func (this *GateApp) SendInner(id string, data []byte) {
	c, ok := this.c2s[id]
	if !ok {
		akLog.Error("can not find client session, id: ", id)
		return
	}
	s, ok := this.s2s[c.SID]
	if !ok {
		akLog.Error("can not find server session, id: ", id)
		return
	}
	s.session.SendMsg(data)
}

//s->gate2<=>gate2->c
//s-> gate2 rid
func (this *GateApp) SendClient(id string, data []byte) {

}
