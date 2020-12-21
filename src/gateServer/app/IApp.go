package app

import (
	"go-snake/common/messageBase"
	"go-snake/common/mixNet"
	"reflect"

	"google.golang.org/protobuf/proto"

	"github.com/Peakchen/xgameCommon/akLog"
)

type IMessageActor interface {
	SetSession(sid string)
	IGateMessage
}

type FnMessageProc func(IMessageActor, proto.Message)

type MessageContent struct {
	refPb   reflect.Type
	refProc reflect.Value
}

var mps = map[uint32]*MessageContent{}

func RegisterActorMessageProc(id uint32, pb interface{}, fn FnMessageProc) {
	mps[id] = &MessageContent{
		refPb:   reflect.TypeOf(pb),
		refProc: reflect.ValueOf(fn),
	}
}

func GetActorMessageProc(id uint32) *MessageContent {
	return mps[id]
}

type MessageActor struct {
	sid string
}

func (this *MessageActor) SetSession(sid string) {
	this.sid = sid
}

func (this *MessageActor) SendMsg(id uint32, pb proto.Message) {
	src, err := proto.Marshal(pb)
	if err != nil {
		akLog.Error("pb marshal fail, err: ", err)
		return
	}
	cspt := messageBase.CSPackTool()
	cspt.Init(id, len(src), src)
	packdata := make([]byte, messageBase.CS_MSG_PACK_DATA_SIZE+len(src))
	cspt.UnPack(packdata)
	mixNet.GetApp().SendClient(this.sid, id, packdata)
}
