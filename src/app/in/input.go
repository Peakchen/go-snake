package in

import (
	"flag"
	"fmt"
	"go-snake/common/config"
)

/*
params:
	1. -app xxx
	2. -host xxx
	3. -pprof xxx
	4. -log xxx
*/
func init() {
	flag.String("app", "", "app name")
	flag.String("ver", "", "ver name")
	flag.String("webhost", "", "webhost content")
	flag.String("tcphost", "", "tcphost content")
	flag.String("pprof", "", "pprof content")
	flag.Int("log", 0, "log content")
	flag.Int("clients", 1, "clients content")

	flag.Parse()
}

type Input struct {
	AppName string
	Ver     string
	WebHost string
	TCPHost string
	PProfIP string
	Log     int
	Clis    int

	Scfg *config.ServerConfig
}

func ParseInput() *Input {
	in := &Input{}
	appName := flag.Lookup("app")
	if appName != nil {
		in.AppName = appName.Value.String()
	}
	ver := flag.Lookup("ver")
	if ver != nil {
		in.Ver = ver.Value.String()
	}
	webhost := flag.Lookup("webhost")
	if webhost != nil {
		in.WebHost = webhost.Value.String()
	}
	tcphost := flag.Lookup("tcphost")
	if tcphost != nil {
		in.TCPHost = tcphost.Value.String()
	}
	pprof := flag.Lookup("pprof")
	if pprof != nil {
		in.PProfIP = pprof.Value.String()
	}
	log := flag.Lookup("log")
	if log != nil {
		in.Log = log.Value.(flag.Getter).Get().(int)
	}
	clients := flag.Lookup("clients")
	if clients != nil {
		in.Clis = clients.Value.(flag.Getter).Get().(int)
	}
	if len(in.Ver) > 0 && len(in.AppName) > 0 {
		in.Scfg = config.LoadServerConfig(fmt.Sprintf("%v_%v", in.AppName, in.Ver))
	}
	return in
}
