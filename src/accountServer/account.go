package accountServer

import (
	"go-snake/akmessage"
	"go-snake/app/in"
	"go-snake/app/application"
)

type Account struct {
}

func New(name string) *Account {
	
	application.SetAppName(name)

	return &Account{}
}

func (this *Account) Init() {

}

func (this *Account) Type() akmessage.ServerType {
	return akmessage.ServerType_Account
}

func (this *Account) Run(d *in.Input) {

}
