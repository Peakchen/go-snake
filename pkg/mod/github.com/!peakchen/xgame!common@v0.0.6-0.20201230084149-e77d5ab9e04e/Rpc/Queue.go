package Rpc

import (
	"github.com/Peakchen/xgameCommon/akLog"
	. "github.com/Peakchen/xgameCommon/RedisConn"
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"sync"
	"time"
)

// add by stefan 20190715 19:39
// purpose: each model timer call back refresh data.

// note: fix 20200108 11:23

type TAokoCallBackParam struct {
	cb     reflect.Value   //call back func
	params []reflect.Value //func params
}

type TDataPack struct {
	Key  string
	Data string
}

type TAokoQueueRpc struct {
	tmo  []string //
	conn *TAokoRedis.Conn
}

var (
	GAokoQueueRpc = &TAokoQueueRpc{}
)

func NewQueueRpc(ctx context.Context, wg *sync.WaitGroup, c *TAokoRedis) {
	GAokoQueueRpc.conn = c.Conn
	GAokoQueueRpc.tmo = []string{}
	wg.Add(1)
	go this.loop(ctx, wg)

}

func (this *TAokoQueueRpc) Call(Identify, name string, model interface{}, args ...interface{}) {

	refVals := []reflect.Value{}
	refVals = append(refVals, args)
	for _, arg := range args {
		refVals = append(refVals, reflect.ValueOf(arg))
	}

	cbdata := &TAokoCallBackParam{
		cb:     reflect.ValueOf(model),
		params: refVals,
	}

	bydata, err := json.Marshal(cbdata)
	if err != nil {
		fmt.Println("register marshal fail: ", err)
		return
	}
	this.tmo = append(this.tmo, name)
	pack = &TDataPack{
		Key:  Identify,
		Data: string(bydata),
	}
	bypack, err := json.Marshal(pack)
	if err != nil {
		fmt.Println("register marshal fail: ", err)
		return
	}
	// Identify + value
	_, err := this.conn.Do("RPUSH", name, bypack...)
	if err != nil {
		akLog.Error("RPUSH data: %v, err: %v.\n", Ret, err)
		return
	}
}

func (this *TAokoQueueRpc) loop(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	tick := time.NewTicker(time.Duration(1) * time.Second)
	for {
		select {
		case <-tick.C:
			this.handler()
		}
	}
}

func (this *TAokoQueueRpc) handler() {
	for _, name := range this.tmo {
		data, err := this.conn.Do("LPOP", name)
		if err != nil {
			akLog.Error("LPOP data: %v, err: %v.\n", data, err)
			continue
		}

		if len(string(data)) == 0 {
			continue
		}

		info := &TDataPack{}
		if err := json.Unmarshal(data, info); err != nil {
			akLog.Error("unmarshal pack fail, model name: ", name)
			continue
		}

		cbdata := &TAokoCallBackParam{}
		if err := json.Unmarshal(info.Data, cbdata); err != nil {
			akLog.Error("unmarshal callback fail, model name: ", name)
			continue
		}

		cbdata.Value.Call(cbdata.params)
	}
}

func (this *TAokoQueueRpc) exit(wg *sync.WaitGroup) {
	this.tmo = nil
	wg.Wait()
}
