package application

import (
	"go-snake/akmessage"
	"go-snake/app/in"
)

type ApplicationIF interface {
	Init()
	Type() akmessage.ServerType
	Run(d *in.Input)
}
