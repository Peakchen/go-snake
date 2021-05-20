package logic

import (
	"go-snake/model/inner"
	"go-snake/model/role"
)

func Init(){

	inner.Register()
	role.Register()
	
}
