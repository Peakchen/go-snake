package user

import "google.golang.org/protobuf/proto"

type IMail interface{
	HandleMailInfo(pb proto.Message)
	HandleMailRead(pb proto.Message)
	HandleMailTake(pb proto.Message)
	HandleMailOneKeyRead(pb proto.Message)
	HandleMailOneKeyTake(pb proto.Message)
	HandleMailDelete(pb proto.Message)
	HandleMailOneKeyDelete(pb proto.Message)
}