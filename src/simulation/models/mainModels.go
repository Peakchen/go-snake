package simuModel

import (
	"strings"
)

type SMIF interface {
	Name() string
	Parse(params ...string)
	Exec()
}

var simuMap = map[string]SMIF{}

func Register(sm SMIF){
	simuMap[sm.Name()] = sm
}

func Run(simuCase string){

	srcs := strings.Split(simuCase, ",")

	model := simuMap[srcs[0]]
	if model != nil {
		model.Parse(srcs[1:]...)
		model.Exec()
	}

}



