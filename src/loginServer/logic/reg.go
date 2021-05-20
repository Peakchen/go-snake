package logic

import (
	"go-snake/model/account"
	"go-snake/model/role"
)

func Init(){

	account.Register()
	role.Register()
	
}
