package app

import (
	"fmt"
	"go-snake/akmessage"
	"go-snake/common/messageBase"
	"reflect"
	"testing"

	"google.golang.org/protobuf/proto"
)

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
