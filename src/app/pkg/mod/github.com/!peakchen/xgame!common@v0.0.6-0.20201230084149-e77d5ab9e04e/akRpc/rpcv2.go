package akRpc

// server no block rpc, enter queue task call back for order.
// add by stefan 20190704 19:24
import (
	. "github.com/Peakchen/xgameCommon/RedisConn"
	"context"
	"fmt"
	"log"
	"sync"
)

type TAorpcV2 struct {
	c        *TAokoRedis
	callfuns map[string]reflect.Value
	acts     []*TModelActV2
}

var GAorpcV2obj *TAorpcV2

func CreateAorpcV2(ctx context.Context, wg *sync.WaitGroup, c *TAokoRedis) {
	GAorpcV2obj = &TAorpcV2{
		c: c,
	}
	wg.Add(1)
	go loop(ctx, wg)
}

func Exit(wg *sync.WaitGroup) {
	wg.Wait()
}

func loop(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			Exit(wg)
			return
		default:
			doRpcTasks()
		}
	}
}

func doRpcTasks() {
	for _, act := range GAorpcV2obj.acts {
		act.modf.Call(act.params)
	}
}

func CallV2(cn string, in []interface{}) (err error) {
	f, ok := GAorpcV2obj.callfuns[cn]
	if !ok {
		return fmt.Errorf("can not find call func: ", cn)
	}
	if GAorpcV2obj.acts == nil {
		GAorpcV2obj.acts = []*TModelActV2{}
	}
	if len(GAorpcV2obj.acts) >= ActChanMaxSize {
		return fmt.Errorf("rpc acts queue has arive max.")
	}
	var params []reflect.Value = []reflect.Value{}
	for _, param := range in {
		params = append(params, reflect.Value(param))
	}
	GAorpcV2obj.acts = append(GAorpcV2obj.acts, &TModelActV2{
		modf:   f,
		params: params,
	})
	return nil
}

func AorpcV2Register(cn string, cf interface{}) {
	GAorpcV2obj.callfuns[cn] = reflect.ValueOf(cf)
}

func init() {
	GAorpcV2obj = nil
}
