package entityBase

import (
	"go-snake/loginServer/logic/account/acc_model"

	"google.golang.org/protobuf/proto"
)

type IAcc interface {
	LoadAcc(acc *acc_model.Acc)

	HandlerRegister(pb proto.Message)
	HandlerLogin(pb proto.Message)
	HandlerLogout(pb proto.Message)
}
