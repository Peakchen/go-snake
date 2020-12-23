package app

import (
	"go-snake/akmessage"
	"go-snake/common/messageBase"
	"go-snake/common/mixNet"
	"reflect"

	"google.golang.org/protobuf/proto"
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

type GateSession struct {
	sid string
}

func (this *GateSession) SetSession(sid string) {
	this.sid = sid
}

func (this *GateSession) SendMsg_cs(id uint32, pb proto.Message) {
	packdata := messageBase.CSPackMsg_pb(akmessage.MSG(id), pb)
	mixNet.GetApp().SendClient(this.sid, id, packdata)
}

func (this *GateSession) SendMsg_ss(id uint32, pb proto.Message) {
	packdata := messageBase.SSPackMsg_pb(this.sid, 0, akmessage.MSG(id), pb)
	mixNet.GetApp().SS_SendInner(this.sid, id, packdata)
}
