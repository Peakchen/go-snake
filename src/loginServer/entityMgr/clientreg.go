package entityMgr

import "google.golang.org/protobuf/proto"

type IClientRegister interface {
	HandlerRegisterResp(pb proto.Message)
}
