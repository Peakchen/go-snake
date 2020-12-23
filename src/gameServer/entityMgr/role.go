package entityMgr

import "google.golang.org/protobuf/proto"

type IRole interface {
	HandlerEnter(pb proto.Message)
}
