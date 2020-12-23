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

func newEntity(sid string, uid int64) *Entity {
	return &Entity{
		uid: uid,
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

	IRole
}

type EntityUser struct {
	*Entity

	IRole
}

func NewEntityBySid(sid string) IEntityUser {
	return &EntityUser{
		Entity: newEntity(sid, utils.NewInt64_v1()),
	}
}

func NewEntity(sid string, uid int64) IEntityUser {
	return &EntityUser{
		Entity: newEntity(sid, uid),
	}
}

func (this *EntityUser) RegModels() {
	RangeModels(func(id int, cb ModelFn) {
		entity := cb(this)
		switch id {
		case M_ROLE:
			this.IRole = entity.(IRole)
		}
	})
}

type IEntityAI interface {
	//self

	//child
	IEntity
}
