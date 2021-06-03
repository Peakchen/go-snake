package natsQueue


import (
	"go-snake/common/myNats"
	"github.com/nats-io/nats.go"
	"go-snake/common/messageBase"
	"go-snake/simulation/sx"
	"go-snake/simulation/models"
	//"go-snake/common"
	"github.com/Peakchen/xgameCommon/akLog"
	"go-snake/akmessage"
	"fmt"
)

type SMQueue struct {
	Params string
}

func (self *SMQueue) Name()string{
	return sx.SM_NatsQueue
}

func (self *SMQueue) Parse(params ...string){

}

func (self *SMQueue) Exec() {

	akLog.FmtPrintln("Subscribe nats messaeg....")
	
	myNats.Subscribe("nats", func(m *nats.Msg) {
		
		pack := messageBase.CSPackTool()
		err := pack.UnPack(m.Data)
		if err != nil {
			akLog.Error(err)
			return
		}

		akLog.FmtPrintln("msg: ", pack.GetMsgID())
	})

	myNats.Subscribe("natstwo", func(m *nats.Msg){
		
		var msg akmessage.CS_AccRegister
		err := messageBase.Codec().Unmarshal(m.Data, &msg)
		if err != nil {
			akLog.Error(fmt.Errorf("unmarshal message fail, err: %v.", err))
			return
		}

		akLog.FmtPrintln("CS_AccRegister: ", msg.Acc)

		var rsp akmessage.SC_AccRegister
		rsp.Ret = 1001
		data := messageBase.CSPackMsg_pb(akmessage.MSG_SC_ACC_REGISTER, &rsp)
		m.Respond(data)

	})

}

func init(){
	simuModel.Register(&SMQueue{})
}
