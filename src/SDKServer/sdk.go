package SDKServer

import (
	"go-snake/akmessage"
	"go-snake/app/in"
)

type SDK struct {
}

func New() *SDK {
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
