package root

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type RootControl struct {
    Ctx iris.Context
}

func (this *RootControl) BeforeActivation(ba mvc.BeforeActivation) {
	
	ba.Handle("GET", "/login", "login", func(ctx iris.Context) {
		ctx.Next()
	})

	ba.Handle("GET", "/register", "register", func(ctx iris.Context) {
		ctx.Next()
	})

}

