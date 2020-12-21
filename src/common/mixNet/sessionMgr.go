package mixNet

// websocket, tcp socket session manager

import (
	"go-snake/common/messageBase"
	"sync"
)

type SessionMgr struct {
	svrSessions sync.Map
	cliSessions sync.Map

	App Application
}

var (
	mgr *SessionMgr
)

func NewSessionMgr(app Application) {
	mgr = &SessionMgr{
		App: app,
	}
}

func GetSessionMgr() *SessionMgr {
	return mgr
}

//--- websocket
func (this *SessionMgr) AddWebSession(id string, sess interface{}) {
	this.svrSessions.Store(id, sess)
	this.App.Online(messageBase.Net_WS, sess)
}

func (this *SessionMgr) GetWebSession(id string) (sess interface{}) {
	val, exist := this.svrSessions.Load(id)
	if exist {
		sess = val
	}
	return
}

func (this *SessionMgr) RemoveWebSession(id string) {
	this.svrSessions.Delete(id)
	this.App.Offline(messageBase.Net_WS, id)
}

func (this *SessionMgr) GetWebSessions() sync.Map {
	return this.svrSessions
}

//--- tcp

func (this *SessionMgr) AddTcpSession(id string, sess interface{}) {
	this.svrSessions.Store(id, sess)
	this.App.Online(messageBase.Net_TCP, sess)
}

func (this *SessionMgr) GetTcpSession(id string) (sess interface{}) {
	val, exist := this.svrSessions.Load(id)
	if exist {
		sess = val
	}
	return
}

func (this *SessionMgr) RemoveTcpSession(id string) {
	this.svrSessions.Delete(id)
	this.App.Offline(messageBase.Net_TCP, id)
}

func (this *SessionMgr) GetTcpSessions() sync.Map {
	return this.svrSessions
}

func (this *SessionMgr) SendInner(sid string, id uint32, data []byte) {
	this.App.SendInner(sid, id, data)
}

func (this *SessionMgr) SendClient(sid string, id uint32, data []byte) {
	this.App.SendClient(sid, id, data)
}

func (this *SessionMgr) Handler(sid string, data []byte) {
	this.App.Handler(sid, data)
}

func init() {

}
