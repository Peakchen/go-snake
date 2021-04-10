package akWebNet

import (
	"sync"

	"github.com/Peakchen/xgameCommon/define"
	"github.com/Peakchen/xgameCommon/utils"
)

type ActorSession struct {
	Actor define.ERouteId
	Sess  *WebSession
}

type wsClientSession struct {
	sessMap sync.Map
}

var (
	GwebSessionMgr = &wsClientSession{}
)

func (this *wsClientSession) AddSession(sess *WebSession, actor define.ERouteId) {
	this.sessMap.Store(sess.RemoteAddr, &ActorSession{
		Actor: actor,
		Sess:  sess,
	})
}

func (this *wsClientSession) GetSession(addr string) *WebSession {
	val, exist := this.sessMap.Load(addr)
	if !exist {
		return nil
	}
	actorSess := val.(*ActorSession)
	return actorSess.Sess
}

func (this *wsClientSession) GetSessionByActor(actor define.ERouteId) (sess *WebSession) {
	var (
		slen      int32
		randIdx   int32
		websesses = []*WebSession{}
	)
	this.sessMap.Range(func(k, v interface{}) bool {
		actorSess := v.(*ActorSession)
		if actorSess.Sess.GetActor().GetActorType() == actor {
			websesses = append(websesses, sess)
		}
		return true
	})

	slen = int32(len(websesses))
	if slen > 1 {
		randIdx = utils.RandInt32FromZero(slen)
	} else if slen == 0 {
		return
	}

	sess = websesses[randIdx]
	return
}

func (this *wsClientSession) RemoveSession(addr string) {
	this.sessMap.Delete(addr)
}

func (this *wsClientSession) GetSessions() sync.Map {
	return this.sessMap
}
