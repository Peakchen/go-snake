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
