package entityMgr

import "google.golang.org/protobuf/proto"

type IAcc interface {
	HandlerRegister(pb proto.Message)
	HandlerLogin(pb proto.Message)
	HandlerLogout(pb proto.Message)
}
