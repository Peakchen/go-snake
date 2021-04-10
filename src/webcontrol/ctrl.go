package webcontrol

import (
	"go-snake/akmessage"
	"go-snake/app/in"
	"go-snake/app/application"
)

type WebControl struct {
}

func New(name string) *WebControl {
	
	application.SetAppName(name)

	return &WebControl{}
}

func (this *WebControl) Init() {

}

func (this *WebControl) Type() akmessage.ServerType {
	return akmessage.ServerType_WebControl
}

func (this *WebControl) Run(d *in.Input) {
	
}
