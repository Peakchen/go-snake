package login

import (
	"go-snake/akmessage"
	"go-snake/robot/RoboIF"
	"go-snake/robot/manager"
	"go-snake/robot/wscli"
	"reflect"
	"time"

	"github.com/Peakchen/xgameCommon/akLog"

	"google.golang.org/protobuf/proto"
)

func init() {
	manager.RegisterModel(&Login{})
}

type Login struct {
	RoboIF.RobotModel

	fns   []func()
	isreg bool
}

func (this *Login) Name() string {
	return reflect.ValueOf(this).Elem().String()
}

func (this *Login) Init(c *wscli.WsNet) {
	this.fns = []func(){}
	this.fns = append(this.fns, this.register)
	this.fns = append(this.fns, this.login)
	this.Dail(c)
}

func (this *Login) Enter() {
	this.Sends()
}

func (this *Login) Sends() {
	for _, fn := range this.fns {
		fn()
	}
}

func (this *Login) login() {
	this.SendMsg(akmessage.MSG_CS_LOGIN, &akmessage.CS_Login{
		Acc: "111",
		Pwd: "222",
	})
	akLog.FmtPrintln("user login....")
}

func (this *Login) register() {
	if this.isreg {
		return
	}
	this.isreg = true
	this.SendMsg(akmessage.MSG_CS_ACC_REGISTER, &akmessage.CS_AccRegister{
		Acc: "111",
		Pwd: "222",
	})
	time.Sleep(time.Second)
	akLog.FmtPrintln("user reg....")
}

func (this *Login) Recv(pb proto.Message) {

}

func (this *Login) Left() {

}
