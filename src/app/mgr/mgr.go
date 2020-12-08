package mgr

import (
	"ak-remote/accountServer"
	"ak-remote/common/application"
	"ak-remote/gameServer"
	"ak-remote/gateServer"
	"ak-remote/loginServer"
)

var (
	apps = map[string]application.ApplicationIF{
		CstGame:    gameServer.New(),
		CstGate:    gateServer.New(),
		CstLogin:   loginServer.New(),
		CstAccount: accountServer.New(),
	}
)

func GetApp(sn string) application.ApplicationIF {
	return apps[sn]
}
