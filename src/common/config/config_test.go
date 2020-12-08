package config

import (
	"testing"
)

func init() {

}

func TestLoadServerCfg(t *testing.T) {
	LoadServerConfig("gameServer")
	GetServerConfig().PrintAll()
}
