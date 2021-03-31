package simulation

import (
	"go-snake/akmessage"
	"go-snake/app/in"
	"go-snake/simulation/rpcBase"
	"go-snake/simulation/models"
	_ "go-snake/simulation/reg"
	"go-snake/common/myNats"
	"github.com/Peakchen/xgameCommon/utils"
)

type Simulation struct {
}

func New() *Simulation {
	return &Simulation{}
}

func (this *Simulation) Init() {

}

func (this *Simulation) Type() akmessage.ServerType {
	return akmessage.ServerType_Simulation
}

func (this *Simulation) Run(d *in.Input) {

	myNats.Register(d.Scfg.NatsHost, utils.ENCodecType_Pb)
	rpcBase.RunRpc(d.Scfg.EtcdIP, d.Scfg.EtcdNodeIP)

	simuModel.Run(d.Scfg.ExtraParams)
}
