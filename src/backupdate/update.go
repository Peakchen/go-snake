package backupdate

import (
	"go-snake/akmessage"
	"go-snake/app/in"
)

type BackUpdate struct {
}

func New() *BackUpdate {
	return &BackUpdate{}
}

func (this *BackUpdate) Init() {

}

func (this *BackUpdate) Type() akmessage.ServerType {
	return akmessage.ServerType_BackUpdate
}

func (this *BackUpdate) Run(d *in.Input) {

}
