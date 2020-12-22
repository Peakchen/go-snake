package config

import (
	"fmt"
	"reflect"
	"testing"

	"gopkg.in/ini.v1"
)

func init() {

}

func TestLoadServerCfg(t *testing.T) {
	LoadServerConfig("gameServer")
	GetServerConfig().PrintAll()
}

func TestRefCfg(t *testing.T) {
	f, err := ini.Load("../../ini/server.ini")
	if err != nil {
		panic(fmt.Errorf("Failed to parse config file: %s", err))
	}
	section, err := f.GetSection("game_1")
	if err != nil {
		panic(fmt.Errorf("invalid config section: %s", err))
	}
	var s = new(ServerConfig)
	for _, k := range section.Keys() {
		fv := reflect.ValueOf(s).Elem().FieldByName(k.Name())
		switch fv.Kind() {
		case reflect.String:
			fv.Set(reflect.ValueOf(k.Value()))
		case reflect.Int:
			intv, _ := k.Int()
			fv.Set(reflect.ValueOf(intv))
		}
	}
	s.PrintAll()
}
