package Cache

/*
	purpose:
	date: 20191205 fix
	author: stefan
*/

import (
	"github.com/Peakchen/xgameCommon/aktime"
	"container/list"
	"sync"
	"time"
)

type TCacheContent struct {
	key      string
	deadtime int64
}

type TCache struct {
	c sync.Map
	//ckey 	chan string
	td int64
	cl *list.List
}

func (this *TCache) Set(key string, data interface{}) {
	d := &TCacheContent{
		key:      key,
		deadtime: aktime.Now().Unix() + this.td,
	}
	this.cl.PushBack(d)
	this.c.Store(key, data)
}

func (this *TCache) Get(key string) interface{} {
	val, ok := this.c.Load(key)
	if !ok {
		return nil
	}
	return val
}

func (this *TCache) Remove(key string) {
	this.c.Delete(key)
}

func (this *TCache) Init(td int64) {
	this.td = td
	this.cl = list.New()
}

func (this *TCache) Run() {
	go this.loopcheck()
}

func (this *TCache) exit() {

}

func (this *TCache) loopcheck() {
	t := time.NewTicker(time.Duration(this.td) * time.Second)
	for {
		select {
		case <-t.C:
			if this.cl.Len() == 0 {
				continue
			}
			e := this.cl.Front()
			data := e.Value.(*TCacheContent)
			if data.deadtime > aktime.Now().Unix() {
				continue
			}
			this.cl.Remove(e)
			this.c.Delete(data.key)
		}
	}
}
