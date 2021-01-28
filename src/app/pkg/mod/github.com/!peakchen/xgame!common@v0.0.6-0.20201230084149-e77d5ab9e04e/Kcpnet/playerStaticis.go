package Kcpnet

import "sync"

var (
	GPlayerStaticis = &PlayerStaticis{}
)

type PlayerStaticis struct {
	playersMap sync.Map
}

func (this *PlayerStaticis) AddPlayer(key interface{}) uint32 {
	cntIf, exist := this.playersMap.Load(key)
	if !exist {
		this.playersMap.Store(key, uint32(1))
		return 1
	}
	cnt := cntIf.(uint32)
	cnt++
	this.playersMap.Store(key, cnt)
	return cnt
}

func (this *PlayerStaticis) GetPlayers(key interface{}) uint32 {
	cntIf, exist := this.playersMap.Load(key)
	if !exist {
		return 0
	}
	return cntIf.(uint32)
}

func (this *PlayerStaticis) SubPlayer(key interface{}) uint32 {
	cntIf, exist := this.playersMap.Load(key)
	if !exist {
		return 0
	}
	cnt := cntIf.(uint32)
	if cnt <= 0 {
		return 0
	}
	cnt--
	this.playersMap.Store(key, cnt)
	return cnt
}
