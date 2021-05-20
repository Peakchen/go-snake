package mgr

import (
	"go-snake/SDKServer"
	"go-snake/accountServer"
	"go-snake/app/application"
	"go-snake/gameServer"
	"go-snake/gateServer"
	"go-snake/loginServer"
	"go-snake/robot"
	"go-snake/simulation"
	"go-snake/webcontrol"
)

var (
	apps = map[string]application.ApplicationIF{
		CstGame:    gameServer.New(CstGame),
		CstGate:    gateServer.New(CstGate),
		CstLogin:   loginServer.New(CstLogin),
		CstAccount: accountServer.New(CstAccount),
		CstRobot:   robot.New(CstRobot),
		CstSDK:     SDKServer.New(CstSDK),
		CstSimu:    simulation.New(CstSimu),
		CstWebctrl: webcontrol.New(CstWebctrl),
	}
)

func GetApp(sn string) application.ApplicationIF {
	return apps[sn]
}
