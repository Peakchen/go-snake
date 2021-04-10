package akRpc

// add by stefan 20190614 19:49
// add aorpc for between server and server conmunication.
// server block rpc.
import (
	"container/list"
	"context"
	"fmt"
	"reflect"
	"sync"
	"time"
)

type TAorpcV1 struct {
	models  map[string]interface{}
	wg      sync.WaitGroup
	ctx     context.Context
	cancel  context.CancelFunc
	acts    *list.List
	actchan chan *TModelActV1
	mutex   sync.Mutex
	retchan chan *TActRet
}

var Aorpc *TAorpcV1 = nil

func init() {
	Aorpc = &TAorpcV1{}
	Aorpc.Init()
}

func (this *TAorpcV1) Init() {
	this.models = map[string]interface{}{}
	this.acts = list.New()
}

func (this *TAorpcV1) Run() {
	this.ctx, this.cancel = context.WithCancel(context.Background())
	this.wg.Add(2)
	go this.loop()
	go this.loopAct()
}

/*
	take model and func witch func in params,
*/
func (this *TAorpcV1) Call(key, modelname, funcName string, ins []interface{}, outs []interface{}) error {
	m, ok := this.models[modelname]
	if !ok {
		return fmt.Errorf("can not find model, input model name: %v.", modelname)
	}
	v := reflect.ValueOf(m)
	t := fmt.Sprintf("%s", reflect.TypeOf(m))
	f := v.MethodByName(funcName)
	rv := []reflect.Value{}
	for _, in := range ins {
		rv = append(rv, reflect.ValueOf(in))
	}
	//f.Call(rv)
	actkey := key + ":" + modelname + ":" + funcName
	this.actchan <- &TModelActV1{
		actid:  actkey,
		modf:   f,
		params: rv,
		mod:    m,
		modt:   t,
	}
	var twg sync.WaitGroup
	twg.Add(1)
	go this.loopRet(actkey, outs, &twg)
	twg.Wait()
	return nil
}

func (this *TAorpcV1) loopRet(actkey string, outs []interface{}, twg *sync.WaitGroup) {
	t := time.NewTicker(time.Duration(rpcdealline))
	for {
		select {
		case ar := <-this.retchan:
			if ar.actid == actkey {
				for i, ret := range ar.rets {
					reflect.ValueOf(outs[i]).Set(ret)
				}
				twg.Done()
				return
			}
		case <-t.C:
			// beyond return time, then return nothing.
			twg.Done()
		}
	}
}

func (this *TAorpcV1) loop() {
	defer this.wg.Done()
	//t := time.NewTicker(time.Duration(rpcdealline))
	for {
		select {
		case <-this.ctx.Done():
			this.Exit()
			return
		//case <-t.C:

		case act := <-this.actchan:
			if act == nil {
				return
			}
			if this.acts.Len() >= ActChanMaxSize {
				fmt.Println("has enough acts in chan.")
				return
			}
			this.acts.PushBack(act)
		}
	}
}

func (this *TAorpcV1) loopAct() {
	defer this.wg.Done()
	for {
		if this.acts.Len() == 0 {
			continue
		}
		this.mutex.Lock()
		e := this.acts.Front()
		act := e.Value.(*TModelActV1)
		if act == nil {
			fmt.Println("act value invalid: ", e.Value)
			continue
		}
		mrts := act.modf.Call(act.params)
		this.retchan <- &TActRet{
			actid: act.actid,
			rets:  mrts,
		}
		this.acts.Remove(e)
		this.mutex.Unlock()
	}
}

func (this *TAorpcV1) Exit() {
	this.cancel()
	this.wg.Wait()
}

func Register(name string, model interface{}) {
	_, ok := Aorpc.models[name]
	if ok {
		return
	}
	Aorpc.models[name] = model
}
