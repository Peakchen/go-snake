package RoboIF

import (
	"go-snake/akmessage"
	"go-snake/common/messageBase"
	"go-snake/robot/wscli"

	"github.com/Peakchen/xgameCommon/akLog"

	"google.golang.org/protobuf/proto"
)

type IRobotModel interface {
	Name() string
	Init(c *wscli.WsNet)
	Enter()
	Recv(pb proto.Message)
	Left()
}

type RobotModel struct {
	cli *wscli.WsNet
}

func (this *RobotModel) Dail(cli *wscli.WsNet) {
	this.cli = cli
}

func (this *RobotModel) SendMsg(id akmessage.MSG, pb proto.Message) {
	data := messageBase.CSPackMsg(id, pb)
	if data == nil {
		akLog.Error("pack msg fail, mid: ", id)
		return
	}
	this.cli.SendMsg(data)
}
