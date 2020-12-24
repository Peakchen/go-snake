package chat

import (
	"go-snake/akmessage"
	"go-snake/app/in"
)

type Chat struct {
}

func New() *Chat {
	return &Chat{}
}

func (this *Chat) Init() {

}

func (this *Chat) Type() akmessage.ServerType {
	return akmessage.ServerType_Chat
}

func (this *Chat) Run(d *in.Input) {

}
