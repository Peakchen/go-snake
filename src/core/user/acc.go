package user

import (
	"go-snake/dbmodel/acc_model"

	"google.golang.org/protobuf/proto"
)

type IAcc interface {
	LoadAcc(acc *accdb.Acc)

	HandlerRegister(pb proto.Message)
	HandlerLogin(pb proto.Message)
	HandlerLogout(pb proto.Message)
}
