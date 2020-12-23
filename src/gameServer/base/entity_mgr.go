package base

import "go-snake/gameServer/entityMgr"

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

func AddUser(uid int64, user entityMgr.IEntityUser) {
	if _, ok := entitys.Users[uid]; ok {
		return
	}
	user.RegModels()
	entitys.Users[uid] = user
}

func (this *EntityManager) GetEntityByID(uid int64) entityMgr.IEntityUser {
	return this.Users[uid]
}

func (this *EntityManager) SetEntityByID(uid int64, user entityMgr.IEntityUser) {
	if _, ok := this.Users[uid]; ok {
		return
	}
	user.RegModels()
	this.Users[uid] = user
}
