package base

import (
	"go-snake/chat/entityBase"
	"sync"
)

type EntityManager struct {
	sync.RWMutex

	entitys map[int64]entityBase.IEntityUser
}

var (
	entitys = &EntityManager{
		entitys: make(map[int64]entityBase.IEntityUser),
	}
)

func GetEntityMgr() *EntityManager {
	return entitys
}

func GetUserByID(uid int64) entityBase.IEntityUser {
	entitys.RLock()
	defer entitys.RUnlock()

	return entitys.entitys[uid]
}

func AddUser(uid int64, user entityBase.IEntityUser) {
	entitys.Lock()
	defer entitys.Unlock()

	if _, ok := entitys.entitys[uid]; ok {
		return
	}
	user.RegModels()
	entitys.entitys[uid] = user
}

func (this *EntityManager) GetEnity(uid int64) entityBase.IEntityUser {
	this.RLock()
	defer this.RUnlock()

	return this.entitys[uid]
}

func (this *EntityManager) AddEnity(uid int64, user entityBase.IEntityUser) {
	this.Lock()
	defer this.Unlock()

	if _, ok := this.entitys[uid]; ok {
		return
	}
	user.RegModels()
	this.entitys[uid] = user
}
