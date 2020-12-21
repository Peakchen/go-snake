package login

import (
	"go-snake/akmessage"
	"go-snake/robot/RoboIF"
	"go-snake/robot/manager"
	"go-snake/robot/wscli"

	"google.golang.org/protobuf/proto"
)

func init() {
	manager.RegisterModel(&Login{})
}

type Login struct {
	RoboIF.RobotModel

	fns []func()
}

func (this *Login) Name() string {
	return "login"
}

func (this *Login) Init(c *wscli.WsNet) {
	this.fns = []func(){}
	this.fns = append(this.fns, func() { this.login() })
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
}

func (this *Login) register() {

}

func (this *Login) Recv(pb proto.Message) {

}

func (this *Login) Left() {

}
