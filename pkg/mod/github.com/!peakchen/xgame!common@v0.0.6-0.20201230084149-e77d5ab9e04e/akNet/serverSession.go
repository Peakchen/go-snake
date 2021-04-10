package akNet

// add by stefan

import (
	"sync"

	"github.com/Peakchen/xgameCommon/define"
	"github.com/Peakchen/xgameCommon/utils"
)

var (
	GServer2ServerSession *TSvr2SvrSession
)

type TSvr2SvrSession struct {
	s2sSession sync.Map
}

func (this *TSvr2SvrSession) RemoveSession(key interface{}) {
	this.s2sSession.Delete(key)
}

func (this *TSvr2SvrSession) AddSession(key interface{}, session TcpSession) {
	this.s2sSession.Store(key, session)
}

func (this *TSvr2SvrSession) GetSession(key interface{}) (session TcpSession) {
	var (
		sessions = []TcpSession{}
		slen     int32
		randIdx  int32
	)
	this.s2sSession.Range(func(k, v interface{}) bool {
		cs := v.(TcpSession)
		if cs.GetRegPoint() == key.(define.ERouteId) && cs.Alive() {
			sessions = append(sessions, cs)
		}
		return true
	})

	slen = int32(len(sessions))
	if slen > 1 {
		randIdx = utils.RandInt32FromZero(slen)
	} else if slen == 0 {
		return
	}

	session = sessions[randIdx]
	return
}

func (this *TSvr2SvrSession) GetSessionByIdentify(key interface{}) (session TcpSession) {
	val, exist := this.s2sSession.Load(key)
	if exist {
		session = val.(TcpSession)
	}
	return
}

func (this *TSvr2SvrSession) GetAllSession() (sessions sync.Map) {
	return this.s2sSession
}

func init() {
	GServer2ServerSession = &TSvr2SvrSession{}
}
