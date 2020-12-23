package login

import (
	"go-snake/akmessage"
	"go-snake/robot/RoboIF"
	"go-snake/robot/base"
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

	fns     []func()
	recvs   map[uint32]func(pb proto.Message)
	isreg   bool
	islogin bool
}

func (this *Login) Name() string {
	return reflect.TypeOf(*this).Name()
}

func (this *Login) Init(v reflect.Value) {
	this.fns = []func(){}
	this.fns = append(this.fns, this.register)
	this.fns = append(this.fns, this.login)

	this.recvs = make(map[uint32]func(pb proto.Message))
	// this.recvs[uint32(akmessage.MSG_SC_ACC_REGISTER)] = this.SC_ACC_REGISTER
	// this.recvs[uint32(akmessage.MSG_SC_LOGIN)] = this.SC_LOGIN

	c := v.Interface().(*wscli.WsNet)
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

func (this *Login) SC_ACC_REGISTER(pb *akmessage.SC_AccRegister) {
	akLog.FmtPrintln("SC_ACC_REGISTER....")
}

func (this *Login) SC_LOGIN(pb *akmessage.SC_Login) {
	akLog.FmtPrintln("SC_LOGIN....")
}

func (this *Login) Recv(fn base.ModelRecvFn) bool {
	return fn(this)
}

func (this *Login) Left() {

}
