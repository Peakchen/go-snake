package logic

import (
	"go-snake/model/inner"
	"go-snake/model/role"
	"go-snake/model/mail"
)

func Init(){

	inner.Register()
	role.Register()
	mail.Register()
	
}
