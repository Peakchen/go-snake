package email

import (
	"go-snake/akmessage"
	"go-snake/app/in"
)

type Email struct {
}

func New() *Email {
	return &Email{}
}

func (this *Email) Init() {

}

func (this *Email) Type() akmessage.ServerType {
	return akmessage.ServerType_Email
}

func (this *Email) Run(d *in.Input) {

}
