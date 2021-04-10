package entityMgr

import (
	"go-snake/logicModule/manager/entity"
)

type EntityHandler func(entity *entity.BaseEntity)

func Call(handler EntityHandler){
	
}