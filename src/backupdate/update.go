package backupdate

import (
	"go-snake/akmessage"
	"go-snake/app/in"
	"go-snake/app/application"
)

type BackUpdate struct {
}

func New(name string) *BackUpdate {
	
	application.SetAppName(name)

	return &BackUpdate{}
}

func (this *BackUpdate) Init() {

}

func (this *BackUpdate) Type() akmessage.ServerType {
	return akmessage.ServerType_BackUpdate
}

func (this *BackUpdate) Run(d *in.Input) {

}
