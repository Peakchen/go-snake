package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func Basic(context *mvc.Application){

    iris.New().Logger().Info("root controller.")

}