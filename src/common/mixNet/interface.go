package mixNet

import (
	"go-snake/common/messageBase"
	"sync"
)

type SessionMgrIf interface {
	AddWebSession(id string, sess interface{})
	GetWebSession(id string) (sess interface{})
	RemoveWebSession(id string)
	GetWebSessions() sync.Map

	AddTcpSession(id string, sess interface{})
	GetTcpSession(id string) (sess interface{})
	RemoveTcpSession(id string)
	GetTcpSessions() sync.Map

	CS_SendInner(sid string, id uint32, data []byte)
	SendClient(sid string, id uint32, data []byte)
	Handler(sid string, data []byte)
}

type Application interface {
	Online(nt messageBase.NetType, sess interface{})
	Offline(nt messageBase.NetType, id string)
	Bind(sid string, id int64)
	CS_SendInner(sid string, id uint32, data []byte)
	SendClient(sid string, id uint32, data []byte)
	Handler(sid string, data []byte)
	SS_SendInner(sid string, id uint32, data []byte)
}

var _app Application

func SetApp(g Application) { _app = g }
func GetApp() Application  { return _app }
