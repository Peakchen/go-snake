package base

import (
	"go-snake/loginServer/entityBase"
	"sync"
)

type EntityManager struct {
	sync.RWMutex

	entitys map[int64]entityBase.IEntityUser
}

var (
	entitymgr = &EntityManager{
		entitys: make(map[int64]entityBase.IEntityUser),
	}
)

func DBPreloading() {
	entitymgr.LoadAll()
}

func GetEntityMgr() *EntityManager {
	return entitymgr
}

func GetUserByID(uid int64) entityBase.IEntityUser {
	entitymgr.RLock()
	defer entitymgr.RUnlock()

	return entitymgr.entitys[uid]
}

func AddUser(uid int64, user entityBase.IEntityUser) {
	entitymgr.Lock()
	defer entitymgr.Unlock()

	if _, ok := entitymgr.entitys[uid]; ok {
		return
	}

	entitymgr.entitys[uid] = user
}

func (this *EntityManager) GetEnity(uid int64) entityBase.IEntityUser {
	this.RLock()
	defer this.RUnlock()

	return this.entitys[uid]
}

func (this *EntityManager) AddEnity(uid int64, user entityBase.IEntityUser) bool {
	this.Lock()
	defer this.Unlock()

	if _, ok := this.entitys[uid]; ok {
		return false
	}

	this.entitys[uid] = user
	return true
}

func (this *EntityManager) UpdateEntity(uid int64, user entityBase.IEntityUser) bool {
	this.Lock()
	defer this.Unlock()

	if _, ok := this.entitys[uid]; !ok {
		return false
	}

	this.entitys[uid] = user
	return true
}
