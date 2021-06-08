package app

import (
	"fmt"
	"go-snake/akmessage"
	"go-snake/common/messageBase"
	"reflect"
	"testing"

	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/utils"

	"google.golang.org/protobuf/proto"
	"time"
	"sync"
)

func init() {
	messageBase.InitCodec(&utils.CodecProtobuf{})
}

type callbackFn func(int, proto.Message)

func FN(id int, pb *akmessage.CS_AccRegister) {

}

func TestMsg(t *testing.T) {
	rfhand := reflect.TypeOf(FN)
	dst := reflect.New(rfhand.In(1).Elem()).Interface().(proto.Message)
	fmt.Println("int:", dst, dst == nil)

	hb := &akmessage.CS_AccRegister{
		Acc: "111",
		Pwd: "222",
	}
	data := messageBase.CSPackMsg_pb(akmessage.MSG_CS_ACC_REGISTER, hb)
	if data == nil {
		return
	}
	messageBase.CSUnPackMsg_pb(data, dst)
}

func TestPackHeartBeat(t *testing.T) {
	data := messageBase.SSPackMsg_pb("111", 1, akmessage.MSG_SC_HEARTBEAT, &akmessage.SS_HeartBeat_Rsp{})
	if data == nil {
		return
	}

	sspt := messageBase.SSPackTool()
	err := sspt.UnPack(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	msgid := sspt.GetMsgID()
	//dstData := sspt.GetData()
	ssroute := messageBase.GetSSRoute()
	//hbrsp := &akmessage.SS_HeartBeat_Rsp{}
	var dstData []byte = nil
	err = messageBase.Codec().Unmarshal(dstData, ssroute)
	if err != nil {
		akLog.Error(fmt.Errorf("unmarshal message fail, err: %v.", err))
		return
	}
	if dstData == nil || len(dstData) == 0 {
		fmt.Println("ss recv invalid, id: ", msgid, dstData)
		return
	}
}

func TestMapOperate(t *testing.T){

	var Data []int64

	var wg sync.WaitGroup
	var mu sync.Mutex
	var cond = sync.NewCond(&mu)
	

	wg.Add(2)

	go func(){

		t := time.NewTicker(10*time.Millisecond)
		defer func(){
			wg.Done()	
			t.Stop()
		}()

		nowt := time.Now().Unix()
		for range t.C{

			mu.Lock()
			nowt += 1
			//for i:=1; i<=5;i++{
			fmt.Println("produce... ")
			Data = append(Data, nowt)
			//}
			mu.Unlock()

			cond.Signal()

		}

	}()

	go func(){

		//t := time.NewTicker(60*time.Millisecond)
		defer func(){
			wg.Done()	
			//t.Stop()
		}()

		for {

			cond.L.Lock()
			if len(Data) == 0 {
				cond.Wait()
			}

			//for range t.C{

			for i:=0; i<len(Data);i++{
				fmt.Println("consume: ", Data[i])
				Data = append(Data[:i], Data[i+1:]...)
			}

			//}

			cond.L.Unlock()
		}

	}()

	wg.Wait()

	fmt.Println("test end.")

}