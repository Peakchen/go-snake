package natsQueue


import (
	"go-snake/common/myNats"
	"github.com/nats-io/nats.go"
	"go-snake/common/messageBase"
	"go-snake/simulation/sx"
	"go-snake/simulation/models"
	//"go-snake/common"
	"github.com/Peakchen/xgameCommon/akLog"
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

}

func init(){
	simuModel.Register(&SMQueue{})
}
