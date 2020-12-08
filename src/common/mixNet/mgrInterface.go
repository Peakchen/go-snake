package mixNet

import (
	"ak-remote/common/messageBase"
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
}

type Application interface {
	Online(nt messageBase.NetType, sess interface{})
	Offline(nt messageBase.NetType, id string)
}
