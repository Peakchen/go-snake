package webcontrol

import (
	"go-snake/akmessage"
	"go-snake/app/in"
)

type WebControl struct {
}

func New() *WebControl {
	return &WebControl{}
}

func (this *WebControl) Init() {

}

func (this *WebControl) Type() akmessage.ServerType {
	return akmessage.ServerType_WebControl
}

func (this *WebControl) Run(d *in.Input) {

}
