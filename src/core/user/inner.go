package user

import "google.golang.org/protobuf/proto"

type IInner interface {
	HandlerRegisterResp(pb proto.Message)
	HandlerHeartBeatResp(pb proto.Message)
}
