package Entity

import (
	"ak-remote/gameServer/logic/base"

	"github.com/globalsign/mgo/bson"
)

type Entity struct {
	base.IEntity

	RID       string
	SessionID string
}

func NewEntity() *Entity {
	return &Entity{
		RID: bson.NewObjectId().String(),
	}
}

func (this *Entity) GetID() string {
	return this.RID
}

func (this *Entity) SetID(id string) {
	this.RID = id
}

func (this *Entity) GetSessionID() string {
	return this.SessionID
}

func (this *Entity) SetSessionID(id string) {
	this.SessionID = id
}
