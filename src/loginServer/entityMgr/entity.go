package entityMgr

import (
	"go-snake/akmessage"
	"go-snake/common/messageBase"
	"go-snake/common/mixNet"

	"github.com/Peakchen/xgameCommon/akLog"

	"github.com/Peakchen/xgameCommon/utils"

	"google.golang.org/protobuf/proto"
)

type IEntity interface {
	GetID() int64
	SetID(id int64)
	GetSessionID() string
	SetSessionID(sid string)
	SendMsg(id akmessage.MSG, pb proto.Message)
}

type Entity struct {
	IEntity

	uid int64
	sid string
}

func newEntity(sid string) *Entity {
	return &Entity{
		uid: utils.NewInt64_v1(),
		sid: sid,
	}
}

func (this *Entity) GetID() int64 { return this.uid }

func (this *Entity) SetID(id int64) { this.uid = id }

func (this *Entity) GetSessionID() string { return this.sid }

func (this *Entity) SetSessionID(id string) { this.sid = id }

func (this *Entity) SendMsg(id akmessage.MSG, pb proto.Message) {
	data := messageBase.SSPackMsg_pb(this.GetSessionID(), this.GetID(), id, pb)
	if data == nil {
		akLog.Error("invalid msg pack fail, id: ", id)
		return
	}
	mixNet.GetApp().SS_SendInner(this.GetSessionID(), uint32(id), data)
}

type IEntityUser interface {
	IEntity

	RegModels()

	IAcc
	IInner
}

type EntityUser struct {
	*Entity

	IAcc
	IInner
}

func NewEntity(sid string) IEntityUser {
	return &EntityUser{
		Entity: newEntity(sid),
	}
}

func (this *EntityUser) RegModels() {
	RangeModels(func(id int, cb ModelFn) {
		entity := cb(this)
		switch id {
		case M_ACCOUNT:
			this.IAcc = entity.(IAcc)
		case M_SERVERINNER:
			this.IInner = entity.(IInner)
		}
	})
}

type IEntityAI interface {
	//self

	//child
	IEntity
}
