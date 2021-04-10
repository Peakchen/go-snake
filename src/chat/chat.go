package chat

import (
	"go-snake/akmessage"
	"go-snake/app/in"
	"go-snake/app/application"
)

type Chat struct {
}

func New(name string) *Chat {
	
	application.SetAppName(name)

	return &Chat{}
}

func (this *Chat) Init() {

}

func (this *Chat) Type() akmessage.ServerType {
	return akmessage.ServerType_Chat
}

func (this *Chat) Run(d *in.Input) {

}
