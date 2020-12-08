package in

import (
	"flag"
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
	flag.String("host", "", "host content")
	flag.String("pprof", "", "pprof content")
	flag.Int("log", 0, "log content")

	flag.Parse()
}

type Input struct {
	AppName string
	Host    string
	PProfIP string
	Log     int
}

func ParseInput() *Input {
	in := &Input{}
	appName := flag.Lookup("app")
	if appName != nil {
		in.AppName = appName.Value.String()
	}
	host := flag.Lookup("host")
	if host != nil {
		in.Host = host.Value.String()
	}
	pprof := flag.Lookup("pprof")
	if pprof != nil {
		in.PProfIP = pprof.Value.String()
	}
	log := flag.Lookup("log")
	if log != nil {
		in.Log = log.Value.(flag.Getter).Get().(int)
	}
	return in
}
