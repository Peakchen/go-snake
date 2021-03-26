package rpcBase

import (
	"google.golang.org/protobuf/proto"
	"go-snake/akmessage"
	"github.com/Peakchen/xgameCommon/akLog"
	"errors"
	"fmt"
)

var (
	gameRpc = newGameRpc()
)

func HandlerGetRoleNum(pb proto.Message) (*akmessage.RpcResponse, error){

	req := pb.(*akmessage.L2G_Get_Role_Num_Req)
	akLog.FmtPrintln("get role num message.", req.String())

	ack := &akmessage.G2L_Get_Role_Num_Rsp{
		Roles: 1,
	}

	data, err := proto.Marshal(ack)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("proto marshal fail,msg: %v.", ack.String()))
	}

	return &akmessage.RpcResponse{RespData: data}, nil
}

func init(){
	gameRpc.Register(akmessage.RPCMSG_L2G_GET_ROLE_NUM_REQ, (*akmessage.L2G_Get_Role_Num_Req)(nil), HandlerGetRoleNum)
}