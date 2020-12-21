package EntityUser

import (
	"go-snake/common/mixNet"
	"go-snake/loginServer/base"

	"google.golang.org/protobuf/proto"
)

func init() {
	base.SetEntityUser(NewEntityUser())
}

type EntityUser struct {
	UserID int64
}

func NewEntityUser() *EntityUser {
	return &EntityUser{}
}

func (this *EntityUser) GetID() int64 {
	return this.UserID
}

func (this *EntityUser) SetID(id int64) { this.UserID = id }

func (this *EntityUser) GetSessionID() string { return "" }

func (this *EntityUser) SetSessionID(id string) {}

func (this *EntityUser) SendMsg(id uint32, pb proto.Message) {
	data, err := proto.Marshal(pb)
	if err != nil {
		return
	}
	mixNet.GetApp().Bind(this.GetID())
	mixNet.GetApp().SendInner(this.GetSessionID(), id, data)
}
