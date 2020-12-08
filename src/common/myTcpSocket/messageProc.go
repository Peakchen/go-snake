package myTcpSocket

import (
	"ak-remote/akmessage"
	"ak-remote/common/messageBase"
	"fmt"

	"google.golang.org/protobuf/proto"

	"github.com/Peakchen/xgameCommon/akLog"
)

//s->s
func ServerProc(data []byte) {
	pt := messageBase.SSPackTool()
	err := pt.UnPack(data)
	if err != nil {
		akLog.Error(err)
		return
	}
	handler := messageBase.MsgHandler(pt.GetMsgID())
	if handler != nil {
		handler.Call(pt.GetData())
	}
}

//c->s
func ClientProc(data []byte) {
	pt := messageBase.PackTool()
	err := pt.UnPack(data)
	if err != nil {
		akLog.Error(err)
		return
	}
	handler := messageBase.MsgHandler(pt.GetMsgID())
	if handler != nil {
		handler.Call(pt.GetData())
	}
}

func PackMsg(mainID akmessage.MSG, msg proto.Message) []byte {
	data, err := proto.Marshal(msg)
	if err != nil {
		akLog.Error(fmt.Errorf("proto marshal fail, mainid: %v, err: %v.", mainID, err))
		return nil
	}

	pack := messageBase.PackTool()
	pack.Init(uint32(mainID), uint32(len(data)), data)

	out := make([]byte, len(data)+messageBase.MSG_PACK_NODATA_SIZE)
	pack.Pack(out)
	return out
}
