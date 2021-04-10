package natsQueue

import (
	"go-snake/common/myNats"
	"go-snake/common/messageBase"
	"go-snake/robot/RoboIF"
	"go-snake/robot/base"
	"go-snake/robot/manager"
	"reflect"
	"go-snake/akmessage"
	"github.com/Peakchen/xgameCommon/akLog"
)

func init(){
	manager.RegisterModel(&NatsQueue{})
}

type NatsQueue struct {

	RoboIF.RobotModel

	fns []func()

}

func (this *NatsQueue) Name() string {
	return reflect.TypeOf(*this).Name()
}

func (this *NatsQueue) Init(v reflect.Value) {

	this.fns = []func(){}
	this.fns = append(this.fns, this.publish)

}

func (this *NatsQueue) Enter() {
	this.Sends()
}

func (this *NatsQueue) Sends() {
	for _, fn := range this.fns {
		fn()
	}
}

func (this *NatsQueue) publish() {

	akLog.FmtPrintln("nats queue publish...")

	myNats.Publish("nats", messageBase.CSPackMsg_pb(akmessage.MSG_CS_ENTER_GAME_SCENE, &akmessage.CS_EnterGameScene{}))

}

func (this *NatsQueue) Recv(fn base.ModelRecvFn) bool {
	return fn(this)
}

func (this *NatsQueue) Left() {
	
}