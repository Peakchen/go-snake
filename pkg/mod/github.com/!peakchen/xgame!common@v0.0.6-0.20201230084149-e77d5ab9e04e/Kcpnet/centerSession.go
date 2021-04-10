package Kcpnet

import "sync"

/*
	server session, key: remoteAddr, value: session obj
	relationship of player and server.
	[
		map1 key: player(addr or identify), value: server session, route server session by player key.
		map2 key: server session, value: arrary for players (addr or identify), if servers disconnect, will delete all value.
	]
*/

type CenterSessionMgr struct {
	playerSessions sync.Map //key,value:server addr
	serverSessions sync.Map //key,value:[]string
}

func (this *CenterSessionMgr) AddPlayerSession(pkey string, session TSession) {
	this.playerSessions.Store(pkey, session)
	this.AppendSvrSession(session, pkey)
}

func (this *CenterSessionMgr) AppendSvrSession(session TSession, vKey string) {
	data, exist := this.serverSessions.Load(session.GetRemoteAddr())
	if !exist {
		vKeys := []string{vKey}
		this.serverSessions.Store(session.GetRemoteAddr(), vKeys)
		return
	}
	vKeys := data.([]string)
	vKeys = append(vKeys, vKey)
	this.serverSessions.Store(session.GetRemoteAddr(), vKeys)
}

func (this *CenterSessionMgr) GetPlayerSession(pkey string) (session TSession) {
	data, exist := this.playerSessions.Load(pkey)
	if exist {
		session = data.(TSession)
	}
	return
}

func (this *CenterSessionMgr) GetSvrSession(skey string) (session TSession) {
	data, exist := this.serverSessions.Load(skey)
	if exist {
		session = data.(TSession)
	}
	return
}

func (this *CenterSessionMgr) ClearSvrSession(session TSession) {
	pkeys, exist := this.serverSessions.Load(session.GetRemoteAddr())
	if exist {
		for _, pkey := range pkeys.([]string) {
			this.ClearPlayerSession(pkey)
		}
	}
	this.serverSessions.Delete(session.GetRemoteAddr())
}

func (this *CenterSessionMgr) ClearPlayerSession(pkey string) {
	this.playerSessions.Delete(pkey)
}
