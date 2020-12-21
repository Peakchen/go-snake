package Entity

import (
	"go-snake/gameServer/logic/base"
)

type Entity struct {
	base.IEntity

	uid int64
	sid string
}

func NewEntity(uid int64) *Entity {
	return &Entity{
		uid: uid,
	}
}

func (this *Entity) GetID() int64 {
	return this.uid
}

func (this *Entity) SetID(id int64) {
	this.uid = id
}

func (this *Entity) GetSessionID() string {
	return this.sid
}

func (this *Entity) SetSessionID(id string) {
	this.sid = id
}
