package msg

import (
	"go-snake/chat/entityBase"
	"reflect"

	"google.golang.org/protobuf/proto"
)

type FnMessageProc func(entityBase.IEntityUser, proto.Message)

type MessageContent struct {
	RefPb   reflect.Type
	RefProc reflect.Value
}

var mps = map[uint32]*MessageContent{}

func RegisterActorMessageProc(id uint32, pb interface{}, fn FnMessageProc) {
	mps[id] = &MessageContent{
		RefPb:   reflect.TypeOf(pb),
		RefProc: reflect.ValueOf(fn),
	}
}

func GetActorMessageProc(id uint32) *MessageContent {
	return mps[id]
}
