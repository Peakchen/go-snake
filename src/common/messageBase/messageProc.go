package messageBase

import (
	"fmt"
	"go-snake/akmessage"

	"google.golang.org/protobuf/proto"

	"github.com/Peakchen/xgameCommon/akLog"
)

// ------------------------------------------
//s->s
func ServerProc(data []byte) {
	pt := SSPackTool()
	err := pt.UnPack(data)
	if err != nil {
		akLog.Error(err)
		return
	}
	handler := MsgHandler(pt.GetMsgID())
	if handler != nil {
		handler.Call(pt.GetData())
	}
}

func SSPackMsg(session string, rid int64, mainID akmessage.MSG, msg proto.Message) []byte {
	data, err := proto.Marshal(msg)
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

func SSUnPackMsg(src []byte) (sid string, dst []byte, id uint32) {
	pack := SSPackTool()
	err := pack.UnPack(src)
	if err != nil {
		akLog.Error("ss un pack fail, err: ", err)
		return
	}
	return pack.GetSessID(), pack.GetData(), pack.GetMsgID()
}

// -----------------------------------------
//c->s
func ClientProc(data []byte) {
	pt := CSPackTool()
	err := pt.UnPack(data)
	if err != nil {
		akLog.Error(err)
		return
	}
	handler := MsgHandler(pt.GetMsgID())
	if handler != nil {
		handler.Call(pt.GetData())
	}
}

func CSPackMsg(mainID akmessage.MSG, msg proto.Message) []byte {
	data, err := proto.Marshal(msg)
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

func GetActorRegisterReq(sid string, st akmessage.ServerType) []byte {
	return SSPackMsg(sid, 0, akmessage.MSG_SS_REGISTER_REQ, &akmessage.SS_Register_Req{St: st})
}

func GetActorRegisterRsp(sid string, st akmessage.ServerType) []byte {
	return SSPackMsg(sid, 0, akmessage.MSG_SS_REGISTER_RSP, &akmessage.SS_Register_Resp{})
}
