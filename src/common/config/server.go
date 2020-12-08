package config

import (
	"fmt"

	"github.com/Peakchen/xgameCommon/akLog"

	"gopkg.in/ini.v1"
)

type ServerConfig struct {
	Host   string
	DBHost string
	DBpwd  string
}

func (this *ServerConfig) PrintAll() {
	akLog.FmtPrintf("Host: %v, DBHost: %v, DBpwd: %v.", this.Host, this.DBHost, this.DBpwd)
}

var (
	_svrcfg = new(ServerConfig)
)

func LoadServerConfig(s string) {
	f, err := ini.Load("./server.ini")
	if err != nil {
		panic(fmt.Errorf("Failed to parse config file: %s", err))
	}
	section, err := f.GetSection(s)
	if err != nil {
		panic(fmt.Errorf("invalid config section: %s", err))
	}

	if section.HasKey("host") {
		k, err := section.GetKey("host")
		if err != nil {
			panic(fmt.Errorf("invalid config key: host"))
		}
		_svrcfg.Host = k.String()
	}
	if section.HasKey("DBHost") {
		k, err := section.GetKey("DBHost")
		if err != nil {
			panic(fmt.Errorf("invalid config key: DBHost"))
		}
		_svrcfg.DBHost = k.String()
	}
	if section.HasKey("DBpwd") {
		k, err := section.GetKey("DBpwd")
		if err != nil {
			panic(fmt.Errorf("invalid config key: DBpwd"))
		}
		_svrcfg.DBpwd = k.String()
	}
}

func GetServerConfig() *ServerConfig {
	return _svrcfg
}

func init() {

}
