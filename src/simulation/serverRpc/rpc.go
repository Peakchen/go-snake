package serverRpc

import (
	"go-snake/common/myetcd"
	"go-snake/akmessage"
	"go-snake/common/logicBase"
	"go-snake/simulation/sx"
	"go-snake/simulation/models"
	//"go-snake/common"
	"time"
	"github.com/Peakchen/xgameCommon/akLog"
)

type SMServerRpc struct {
	Params string
}

func (self *SMServerRpc) Name()string{
	return sx.SM_Discovery
}

func (self *SMServerRpc) Parse(params ...string){

	if len(params) == 0 {
		return
	}

	//var dst []interface{}
	//common.ParseJson(params, &dst, "simulation json extraparams.")

}

func (self *SMServerRpc) Exec() {

	t := time.NewTicker(2*time.Second)
	defer t.Stop()

	for range t.C{
		akLog.FmtPrintln("enter game scene....")
		msg := &akmessage.L2G_Get_Role_Num_Req{}
		_, err := myetcd.Call(logicBase.RPC_GAME, akmessage.RPCMSG_L2G_GET_ROLE_NUM_REQ, msg)
		akLog.FmtPrintln("err: ", err)
	}

}

func init(){
	simuModel.Register(&SMServerRpc{})
}
