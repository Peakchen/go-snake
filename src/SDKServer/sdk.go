package SDKServer

import (
	"go-snake/akmessage"
	"go-snake/app/in"
	"go-snake/app/application"
)

type SDK struct {
}

func New(name string) *SDK {
	
	application.SetAppName(name)

	return &SDK{}
}

func (this *SDK) Init() {
	//load config...
	//...

}

func (this *SDK) Type() akmessage.ServerType {
	return akmessage.ServerType_SDK
}

func (this *SDK) Run(d *in.Input) {

}
