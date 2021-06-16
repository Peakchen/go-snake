package user

import "google.golang.org/protobuf/proto"

type IChat interface {
	HandleChat(pb proto.Message)
	HandleChat_ss(pb proto.Message)
}