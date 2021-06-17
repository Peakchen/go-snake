package user

import "google.golang.org/protobuf/proto"

type IFriend interface {
	HandleFriendList(pb proto.Message)
	HandleFriendAdd(pb proto.Message)
	HandleFriendDelete(pb proto.Message)
	HandleFriendQuery(pb proto.Message)
	HandleFriendBlack(pb proto.Message)
}