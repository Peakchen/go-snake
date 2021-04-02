package email

import (
	"go-snake/akmessage"
	"go-snake/app/in"
	"go-snake/app/application"
)

type Email struct {
}

func New(name string) *Email {
	
	application.SetAppName(name)

	return &Email{}
}

func (this *Email) Init() {

}

func (this *Email) Type() akmessage.ServerType {
	return akmessage.ServerType_Email
}

func (this *Email) Run(d *in.Input) {

}
