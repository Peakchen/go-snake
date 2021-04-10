package entityMgr

import (
	"go-snake/logicModule/manager/entity"
)

type BaseEntityMgr struct {
	
	entitys map[int64]*entity.BaseEntity

}

func newEntityMgr()*BaseEntityMgr{
	
	return &BaseEntityMgr{
		entitys: make(map[int64]*entity.BaseEntity),
	}
}

func (this *BaseEntityMgr) GetEntity(uid int64)*entity.BaseEntity{

	return this.entitys[uid]

}

func (this *BaseEntityMgr) SetEntity(uid int64){
	
	if this.entitys[uid] != nil {
		return
	}

	this.entitys[uid] = entity.NewEntity(uid)

}

