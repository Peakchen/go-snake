package user

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

	dbid int64
	sid  string
}

func newEntity(sid string) *Entity {
	return &Entity{
		dbid: utils.NewInt64_v1(),
		sid:  sid,
	}
}

func (this *Entity) GetID() int64 { return this.dbid }

func (this *Entity) SetID(id int64) { this.dbid = id }

func (this *Entity) GetSessionID() string { return this.sid }

func (this *Entity) SetSessionID(id string) { this.sid = id }

func (this *Entity) SendMsg(id akmessage.MSG, pb proto.Message) {
	if len(this.GetSessionID()) == 0 {
		akLog.Info("session is null.")
	}
	data := messageBase.SSPackMsg_pb(this.GetSessionID(), this.GetID(), id, pb)
	if data == nil {
		akLog.Error("invalid msg pack fail, id: ", id)
		return
	}
	mixNet.GetApp().SS_SendInner(this.GetSessionID(), uint32(id), data)
}

type IEntityUser interface {
	IEntity

	//RegModels()

	IAcc
	IInner
	IWxRole
	IRole
}

type EntityUser struct {
	*Entity

	IAcc
	IInner
	IWxRole
	IRole
}

func InitEntity(dbid int64) IEntityUser {
	user := &EntityUser{
		Entity: &Entity{
			dbid: dbid,
		},
	}
	user.RegModels()
	return user
}

func NewEntity(sid string, uid int64) IEntityUser {
	return &EntityUser{
		Entity: &Entity{
			dbid: uid,
			sid:  sid,
		},
	}
}

func NewEntityBySid(sid string) IEntityUser {
	user := &EntityUser{
		Entity: newEntity(sid),
	}
	user.RegModels()
	return user
}

func (this *EntityUser) RegModels() {
	RangeModels(func(id int, cb ModelFn) {
		entity := cb(this)
		switch id {
		case M_ACCOUNT:
			this.IAcc = entity.(IAcc)
		case M_SERVERINNER:
			this.IInner = entity.(IInner)
		case M_WXROLE:
			this.IWxRole = entity.(IWxRole)
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
