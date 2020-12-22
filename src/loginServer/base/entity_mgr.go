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

func GetUserByID(id int64) entityMgr.IEntityUser {
	return entitys.Users[id]
}

func (this *EntityManager) GetEntityByID(rid int64) entityMgr.IEntityUser {
	return this.Users[rid]
}

func (this *EntityManager) SetEntityByID(rid int64, user entityMgr.IEntityUser) {
	user.RegModels()
	this.Users[rid] = user
}
