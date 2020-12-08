package gateServer

import (
	"math/rand"

	"github.com/Peakchen/xgameCommon/aktime"
)

func init() {
	t := aktime.Now().Unix()
	s := rand.NewSource(t)
	rand.New(s).Seed(t)
}

type Gate struct {
}

func New() *Gate {
	return &Gate{}
}

func (this *Gate) Init() {

}

func (this *Gate) Run() {

}
