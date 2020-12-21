package manager

import (
	"go-snake/robot/RoboIF"
	"go-snake/robot/wscli"
)

var robotModels = make(map[string]RoboIF.IRobotModel)

func RegisterModel(m RoboIF.IRobotModel)   { robotModels[m.Name()] = m }
func GetModel(n string) RoboIF.IRobotModel { return robotModels[n] }
func RangeModels(c *wscli.WsNet) {
	for _, md := range robotModels {
		md.Init(c)
		md.Enter()
	}
}
