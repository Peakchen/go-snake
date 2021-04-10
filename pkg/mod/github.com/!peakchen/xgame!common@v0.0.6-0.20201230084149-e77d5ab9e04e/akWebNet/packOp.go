package akWebNet

import (
	"fmt"
	"reflect"

	"github.com/Peakchen/xgameCommon/define"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_MainModule"
	"github.com/golang/protobuf/proto"
)

type BaseMsgInfo struct {
	Actor  uint16
	MainID uint16
	SubID  uint16
}

type MsgOperIF interface {
	Init()
	Pack(mainid, subid uint16, msg proto.Message) (out []byte, err error)
	UnPack() (msg proto.Message, cb reflect.Value, err error)
	GetMsgInfo(data []byte) (info *BaseMsgInfo, err error)
}

var (
	msgOps = map[PACK_TYPE]MsgOperIF{
		PACK_PROTO: &ProtoBufMsgOp{},
	}
)

func PackMsgOp(mainid, subid uint16, msg proto.Message, pt PACK_TYPE) (out []byte, err error) {
	op, exist := msgOps[pt]
	if !exist {
		err = fmt.Errorf("can not find msg op, src info[mainid: %v, subid: %v, pt: %v].", mainid, subid, pt)
		return
	}
	op.Init()
	out, err = op.Pack(mainid, subid, msg)
	return
}

func UnPackMsgOp(pt PACK_TYPE) (msg proto.Message, cb reflect.Value, err error) {
	op, exist := msgOps[pt]
	if !exist {
		err = fmt.Errorf("can not find msg op, src info[pt: %v].", pt)
		return
	}
	op.Init()
	msg, cb, err = op.UnPack()
	return
}

func GetUnPackMsgInfo(data []byte, pt PACK_TYPE) (info *BaseMsgInfo, err error) {
	op, exist := msgOps[pt]
	if !exist {
		err = fmt.Errorf("can not find msg op, src info[pt: %v].", pt)
		return
	}
	op.Init()
	info, err = op.GetMsgInfo(data)
	if info.MainID == uint16(MSG_MainModule.MAINMSG_SERVER) {
		info.Actor = uint16(define.ERouteId_ER_SG)
	} else if info.MainID == uint16(MSG_MainModule.MAINMSG_LOGIN) {
		info.Actor = uint16(define.ERouteId_ER_Login)
	} else if info.MainID == uint16(MSG_MainModule.MAINMSG_HEARTBEAT) || info.MainID == uint16(MSG_MainModule.MAINMSG_RPC) {
		return
	} else if info.MainID >= uint16(MSG_MainModule.MAINMSG_PLAYER) {
		info.Actor = uint16(define.ERouteId_ER_Game)
	}
	return
}
