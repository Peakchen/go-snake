package app

import (
	"fmt"
	"go-snake/akmessage"
	"go-snake/common/messageBase"
	"reflect"
	"testing"

	"github.com/Peakchen/xgameCommon/akLog"
	"google.golang.org/protobuf/proto"
)

type callbackFn func(int, proto.Message)

func FN(id int, pb proto.Message) {

}

func TestMsg(t *testing.T) {
	rfhand := reflect.TypeOf(FN)
	dst := reflect.New(rfhand.In(1)).Interface()
	fmt.Println("int:", dst, dst == nil)

	cspt := messageBase.CSPackTool()
	hb := &akmessage.CS_HeartBeat{}
	src, err := proto.Marshal(hb)
	if err != nil {
		akLog.Error("pb marshal heart beat msg fail.")
		return
	}
	cspt.Init(uint32(akmessage.MSG_CS_HEARTBEAT), len(src), src)
	data := make([]byte, len(src)+messageBase.CS_MSG_PACK_DATA_SIZE)
	cspt.Pack(data)

	err = proto.Unmarshal(data, dst.(proto.Message))
	if err != nil {
		akLog.Error(fmt.Errorf("unmarshal message fail, err: %v.", err))
		return
	}
}
