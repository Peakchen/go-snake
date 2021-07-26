package mapSys

import (
	"go-snake/akmessage"
	"go-snake/common/myNats"
	"github.com/nats-io/nats.go"
	"go-snake/common/messageBase"

	"github.com/Peakchen/xgameCommon/akLog"
	//"google.golang.org/protobuf/proto"
	"fmt"
)


func Register(){

	myNats.Subscribe("mapEnter", func(m *nats.Msg) {
		
		var msg akmessage.SS_EnterMap_Req
		err := messageBase.Codec().Unmarshal(m.Data, &msg)
		if err != nil {
			akLog.Error(fmt.Errorf("unmarshal message fail, err: %v.", err))
			return
		}

		mapEntity.HandleEnterMap_ss(&msg)

	})

	// msg.RegisterActorMessageProc(uint32(akmessage.MSG_SS_ENTERMAP_REQ), (*akmessage.SS_EnterMap_Req)(nil),
	// func(actor user.IEntityUser, pb proto.Message) {
	// 	if actor != nil {	actor.HandleEnterMap_ss(pb)	}else{	akLog.Error("invalid actor object.")	}
	// })

}

func (this *Map) HandleEnterMap_ss(req *akmessage.SS_EnterMap_Req){

	akLog.FmtPrintln("enter map： ", req.RoleID, req.MapID)
	this.SendMsg(akmessage.MSG_SS_ENTERMAP_RSP, &akmessage.SS_EnterMap_Rsp{
		RoleID: req.RoleID,
		MapID: req.MapID,
	})

}
