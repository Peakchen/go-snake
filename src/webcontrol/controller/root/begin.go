package root

import (
	"github.com/kataras/iris/v12/mvc"
)

func RootBegin(ctx *mvc.Application){
	ctx.Handle(new(RootControl))
}
