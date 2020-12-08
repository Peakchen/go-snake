package cache

import (
	"ak-remote/common/rediscache"
	"fmt"
	"sync"
	"time"
)

const (
	deadlineDur   = time.Duration(5 * 60 * time.Second)
	freshRedisDur = time.Duration(1 * 60 * time.Second)
)

//key +
type DetailData struct {
	t   int64
	val interface{}
}

type MemCache struct {
	data sync.Map
}

var (
	GMemCache *MemCache = &MemCache{}
)

func (this *MemCache) Run() {
	go this.loopcheck()
}

func (this *MemCache) Set(key string, v interface{}) {
	this.data.Store(key, &DetailData{
		val: v,
		t:   time.Now().Unix(),
	})
}

func (this *MemCache) Get(key string) (v interface{}) {
	valInterface, existed := this.data.Load(key)
	if !existed {
		v = nil
	} else {
		origin := valInterface.(*DetailData)
		v = origin.val
	}

	return
}

/*
	数据过期时间比刷新缓存时间要长，使得数据在及时更新时能够保持redis内数据是最新的
*/
func (this *MemCache) loopcheck() {
	deadlineTick := time.NewTicker(deadlineDur)
	freshRedisTick := time.NewTicker(freshRedisDur)

	for {
		select {
		case <-deadlineTick.C:
			// 到期时间清理数据
			nowt := time.Now().Unix()
			fmt.Println("clean mem cache data t: ", nowt)
			this.data.Range(func(k, v interface{}) bool {
				origin := v.(*DetailData)
				if nowt-origin.t >= 5*60 {
					this.data.Delete(k)
				}
				return true
			})

		case <-freshRedisTick.C:
			// 定时更新redis 数据
			//fmt.Println("refresh redis cache data ... ")
			this.data.Range(func(k, v interface{}) bool {
				origin := v.(*DetailData)
				key := k.(string)
				rediscache.SetCache(key, origin.val)
				return true
			})
		}
	}
}
