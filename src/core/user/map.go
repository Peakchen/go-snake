package user

import "google.golang.org/protobuf/proto"


type IMap interface {
	HandleEnterMap(pb proto.Message)
}