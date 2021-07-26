package mapSys

import (
	"go-snake/core/user"
)

var mapEntity *Map

func init() {
	
	user.RegisterModel(user.M_Map, func(entity user.IEntityUser) interface{} { 

		mapEntity = newMap(entity) 
		return mapEntity
	})
}

func newMap(entity user.IEntityUser)*Map{
	return &Map{
		IEntityUser: entity,
	}
}

type Map struct {

	user.IEntityUser

}



