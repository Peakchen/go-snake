package base

import "go-snake/loginServer/entityMgr"

type EntityManager struct {
	Users map[int64]entityMgr.IEntityUser
}

var (
	entitys = &EntityManager{
		Users: make(map[int64]entityMgr.IEntityUser),
	}
)

func GetEntityMgr() *EntityManager {
	return entitys
}

func GetUserByID(uid int64) entityMgr.IEntityUser {
	return entitys.Users[uid]
}

func (this *EntityManager) GetEntityByID(uid int64) entityMgr.IEntityUser {
	return this.Users[uid]
}

func (this *EntityManager) SetEntityByID(uid int64, user entityMgr.IEntityUser) {
	user.RegModels()
	this.Users[uid] = user
}
