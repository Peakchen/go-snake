package messageBase

import (
	"fmt"
	"go-snake/akmessage"

	"google.golang.org/protobuf/proto"

	"github.com/Peakchen/xgameCommon/akLog"
)

// ------------------------------------------
//s->s
func ServerProc_pb(data []byte) {
	pt := SSPackTool()
	err := pt.UnPack(data)
	if err != nil {
		akLog.Error(err)
		return
	}
	handler := MsgHandler(pt.GetMsgID())
	if handler != nil {
		handler.Call_pb(pt.GetData())
	}
}

func SSPackMsg_pb(session string, rid int64, mainID akmessage.MSG, msg proto.Message) []byte {
	data, err := Codec().Marshal(msg)
	if err != nil {
		akLog.Error(fmt.Errorf("proto marshal fail, mainid: %v, err: %v.", mainID, err))
		return nil
	}

	pack := SSPackTool()
	pack.Init(session, rid, uint32(mainID), data)

	out := make([]byte, len(data)+SS_MSG_PACK_NODATA_SIZE)
	pack.Pack(out)
	akLog.FmtPrintln("ss pack msg: ", out, len(out))
	return out
}

func SSUnPackMsg_pb(src []byte, pb proto.Message) (sid string, id uint32) {
	pack := SSPackTool()
	err := pack.UnPack(src)
	if err != nil {
		akLog.Error("ss un pack fail, err: ", err)
		return
	}
	err = Codec().Unmarshal(pack.GetData(), pb)
	if err != nil {
		akLog.Error(fmt.Errorf("unmarshal message fail, err: %v.", err))
		return
	}
	return pack.GetSessID(), pack.GetMsgID()
}

// -----------------------------------------
//c->s
func ClientProc_pb(data []byte) {
	pt := CSPackTool()
	err := pt.UnPack(data)
	if err != nil {
		akLog.Error(err)
		return
	}
	handler := MsgHandler(pt.GetMsgID())
	if handler != nil {
		handler.Call_pb(pt.GetData())
	}
}

func CSPackMsg_pb(mainID akmessage.MSG, msg proto.Message) []byte {
	data, err := Codec().Marshal(msg)
	if err != nil {
		akLog.Error(fmt.Errorf("proto marshal fail, mainid: %v, err: %v.", mainID, err))
		return nil
	}

	pack := CSPackTool()
	pack.Init(uint32(mainID), len(data), data)

	out := make([]byte, len(data)+CS_MSG_PACK_NODATA_SIZE)
	pack.Pack(out)
	return out
}

func CSUnPackMsg_pb(data []byte, pb proto.Message) (id uint32) {
	pack := CSPackTool()
	err := pack.UnPack(data)
	if err != nil {
		akLog.Error(err)
		return
	}
	err = Codec().Unmarshal(pack.GetData(), pb)
	if err != nil {
		akLog.Error(fmt.Errorf("unmarshal message fail, err: %v.", err))
		return
	}
	return pack.GetMsgID()
}
