package robot

import (
	"go-snake/akmessage"
	"go-snake/app/in"
	_ "go-snake/robot/login"
	"go-snake/robot/manager"
	"go-snake/robot/wscli"
)

type Robot struct {
}

func New() *Robot {
	return &Robot{}
}

func (this *Robot) Init() {
	//load config...
	//...
}

func (this *Robot) Type() akmessage.ServerType {
	return akmessage.ServerType_Robot
}

func (this *Robot) Run(d *in.Input) {
	for i := 0; i < d.Clis; i++ {
		wscli.NewClient(d.WebHost, manager.RangeModels)
	}
}
