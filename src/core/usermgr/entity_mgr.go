package usermgr

import (
	"go-snake/core/user"
	"sync"
)

type EntityManager struct {
	sync.RWMutex

	entitys map[int64]user.IEntityUser
	loads   []func(em *EntityManager)
}

var (
	entitymgr = &EntityManager{
		entitys: make(map[int64]user.IEntityUser),
	}
)

func DBPreLoad() {
	entitymgr.LoadAll()
}

func DBPreAddAndLoad(f ...func(em *EntityManager)){
	entitymgr.Add(f...)
	DBPreLoad()
}

func GetEntityMgr() *EntityManager {
	return entitymgr
}

func GetUserByID(uid int64) user.IEntityUser {
	entitymgr.RLock()
	defer entitymgr.RUnlock()

	return entitymgr.entitys[uid]
}

func AddUser(uid int64, user user.IEntityUser) {
	entitymgr.Lock()
	defer entitymgr.Unlock()

	if _, ok := entitymgr.entitys[uid]; ok {
		return
	}

	entitymgr.entitys[uid] = user
}

func (this *EntityManager) GetEnity(uid int64) user.IEntityUser {
	this.RLock()
	defer this.RUnlock()

	return this.entitys[uid]
}

func (this *EntityManager) AddEnity(uid int64, user user.IEntityUser) bool {
	this.Lock()
	defer this.Unlock()

	if _, ok := this.entitys[uid]; ok {
		return false
	}

	this.entitys[uid] = user
	return true
}

func (this *EntityManager) UpdateEntity(uid int64, user user.IEntityUser) bool {
	this.Lock()
	defer this.Unlock()

	if _, ok := this.entitys[uid]; !ok {
		return false
	}

	this.entitys[uid] = user
	return true
}
