package Kcpnet

import (
	"sort"
	"sync"

	"github.com/Peakchen/xgameCommon/define"
)

type SessionMgr struct {
	sessionMap sync.Map
}

var (
	GServer2ServerSession *SessionMgr
	GClient2ServerSession *SessionMgr
)

func (this *SessionMgr) AddSession(key interface{}, sess TSession) {
	this.sessionMap.Store(key, sess)
}

func (this *SessionMgr) GetSession(key interface{}) (sess TSession) {
	v, exist := this.sessionMap.Load(key)
	if exist {
		sess = v.(TSession)
	}
	return
}

type Sessionlist []TSession

func (s Sessionlist) Len() int      { return len(s) }
func (s Sessionlist) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Sessionlist) Less(i, j int) bool {
	if s[i].GetVer() < s[j].GetVer() {
		return false
	}
	return GPlayerStaticis.GetPlayers(s[i].GetRemoteAddr()) < GPlayerStaticis.GetPlayers(s[j].GetRemoteAddr())
}

func (this *SessionMgr) GetBalanceSession(key interface{}) (sess TSession) {
	var (
		sessions = []TSession{}
		//randIdx  int32
	)
	this.sessionMap.Range(func(k, v interface{}) bool {
		cs := v.(TSession)
		if cs.GetRegPoint() == key.(define.ERouteId) && cs.Alive() {
			sessions = append(sessions, cs)
		}
		return true
	})

	slen := int32(len(sessions))
	if slen == 0 {
		return
	} else if slen > 1 {
		sort.Sort(Sessionlist(sessions))
	}
	sess = sessions[0]
	// if slen > 1 {
	// 	randIdx = utils.RandInt32FromZero(slen)
	// } else if slen == 0 {
	// 	return
	// }
	//session = sessions[randIdx]
	return
}

func (this *SessionMgr) RemoveSession(key interface{}) (exist bool) {
	_, exist = this.sessionMap.Load(key)
	if exist {
		this.sessionMap.Delete(key)
	}
	return
}

func (this *SessionMgr) GetSessionByIdentify(key interface{}) (sess TSession) {
	val, exist := this.sessionMap.Load(key)
	if exist {
		sess = val.(TSession)
	}
	return
}

func (this *SessionMgr) GetAllSession() (sess sync.Map) {
	return this.sessionMap
}

func init() {
	GServer2ServerSession = &SessionMgr{}
	GClient2ServerSession = &SessionMgr{}
}
