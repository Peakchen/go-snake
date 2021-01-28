package akWebNet

import (
	"sync"

	"github.com/Peakchen/xgameCommon/akNet"
)

type MsgRoute struct {
	msg sync.Map
}

func (this *MsgRoute) Register(cmd uint32, val int32) {
	this.msg.Store(cmd, val)
}

func (this *MsgRoute) IsMsg(mainid, subid uint16) (exist bool) {
	cmd := akNet.EncodeCmd(mainid, subid)
	_, exist = this.msg.Load(cmd)
	return
}

func (this *MsgRoute) UnRegisterMsgs() {
	this.msg.Range(func(cmd, v interface{}) bool {
		this.msg.Delete(cmd)
		return true
	})
}
