package root

import (
	"github.com/kataras/iris/v12"
	"go-snake/webcontrol/parser"
)


func (self *RootControl) login(){
	
	iris.New().Logger().Info("GET --> login")
	
	ret, err := parser.ParseURL(self.Ctx.Request().RequestURI)
	if err != nil {
		return
	}

	for k, v := range ret {
		iris.New().Logger().Info(k, v)
	}

}

func (self *RootControl) register(){
	
	iris.New().Logger().Info("GET --> register")

	ret, err := parser.ParseURL(self.Ctx.Request().RequestURI)
	if err != nil {
		return
	}

	for k, v := range ret {
		iris.New().Logger().Info(k, v)
	}

}