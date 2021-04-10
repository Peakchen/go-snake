package entity

import (
	"go-snake/dbModule"
)

type BaseEntity struct {
	
	uid int64
	role *dbModule.Role

}

func NewEntity(uid int64)*BaseEntity{
	
	return &BaseEntity{
		uid: uid,
	}

}

func (this *BaseEntity) SetUID(uid int64) {
	this.uid = uid
}


