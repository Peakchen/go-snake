package config

import (
	"fmt"
	"reflect"

	"github.com/Peakchen/xgameCommon/akLog"
	"gopkg.in/ini.v1"
	"strings"
	"strconv"
)

type ServerConfig struct {

	Node string
	Name string
	ID   int

	TCPHost string
	WebHost string

	RedisHost  string
	RedisPwd   string
	RedisIndex int

	MysqlHost     string
	MysqlUser     string
	MysqlPwd      string
	MysqlDataBase string

	PprofHost string

	HasWechat bool
	WebHttp   string
	AppID     string
	AppSecret string

	EtcdIP  	string
	EtcdNodeIP  string

	ExtraParams string
	
	NatsHost string
}

func (this *ServerConfig) PrintAll() {
	akLog.FmtPrintf("TCPHost: %v, WebHost: %v, RedisHost: %v, RedisPwd: %v, RedisIndex: %v, MysqlHost: %v, MysqlUser: %v, MysqlPwd: %v, MysqlDataBase: %v.",
		this.TCPHost,
		this.WebHost,

		this.RedisHost,
		this.RedisPwd,
		this.RedisIndex,

		this.MysqlHost,
		this.MysqlUser,
		this.MysqlPwd,
		this.MysqlDataBase,
		this.PprofHost,

		this.HasWechat,
		this.WebHttp,
		this.AppID,
		this.AppSecret,
	)
}

var (
	_svrcfg *ServerConfig
)

func LoadServerConfig(s string) *ServerConfig {

	f, err := ini.Load("./ini/server.ini")
	if err != nil {
		panic(fmt.Errorf("Failed to parse config file: %s", err))
	}

	section, err := f.GetSection(s)
	if err != nil {
		panic(fmt.Errorf("invalid config section: %s", err))
	}

	if !strings.Contains(section.Name(), "_") {
		panic("invalid ini service node: " + section.Name())
	}

	_cfg := new(ServerConfig)
	
	_cfg.Node = section.Name()
	sectionNode := strings.Split(section.Name(), "_")
	_cfg.Name = sectionNode[0]
	_cfg.ID, _ = strconv.Atoi(sectionNode[1])

	for _, k := range section.Keys() {

		fv := reflect.ValueOf(_cfg).Elem().FieldByName(k.Name())

		switch fv.Kind() {
		case reflect.Bool:

			boolv, _ := k.Bool()
			fv.Set(reflect.ValueOf(boolv))

		case reflect.String:

			fv.Set(reflect.ValueOf(k.Value()))

		case reflect.Int:

			intv, _ := k.Int()
			fv.Set(reflect.ValueOf(intv))

		}
	}

	_svrcfg = _cfg

	return _svrcfg
}

func GetServerConfig() *ServerConfig {
	return _svrcfg
}

func init() {

}
