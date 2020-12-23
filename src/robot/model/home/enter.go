package home

import (
	"go-snake/akmessage"
	"go-snake/robot/RoboIF"
	"go-snake/robot/base"
	"go-snake/robot/manager"
	"go-snake/robot/wscli"
	"reflect"

	"github.com/Peakchen/xgameCommon/akLog"
)

func init() {
	manager.RegisterModel(&GameHome{})
}

type GameHome struct {
	RoboIF.RobotModel

	fns []func()
}

func (this *GameHome) Name() string {
	return reflect.TypeOf(*this).Name()
}

func (this *GameHome) Init(v reflect.Value) {

	this.fns = []func(){}
	this.fns = append(this.fns, this.enter)

	c := v.Interface().(*wscli.WsNet)
	this.Dail(c)
}

func (this *GameHome) Enter() {
	this.Sends()
}

func (this *GameHome) Sends() {
	for _, fn := range this.fns {
		fn()
	}
}

func (this *GameHome) enter() {
	this.SendMsg(akmessage.MSG_CS_ENTER_GAME_SCENE, &akmessage.CS_EnterGameScene{})
	akLog.FmtPrintln("user enter game....")
}

func (this *GameHome) Recv(fn base.ModelRecvFn) bool {
	return fn(this)
}

func (this *GameHome) SC_EnterGameScene(pb *akmessage.SC_EnterGameScene) {
	akLog.FmtPrintln("SC_EnterGameScene....")
}

func (this *GameHome) Left() {

}
