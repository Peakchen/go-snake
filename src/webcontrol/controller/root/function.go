package root

import (
	"github.com/kataras/iris/v12"
	"go-snake/webcontrol/parser"
	"go-snake/webcontrol/controller/errorCode"
	"go-snake/dbmodel/backdb"
	"github.com/Peakchen/xgameCommon/utils"
)


func (this *RootControl) login(){
	
	iris.New().Logger().Info("GET --> login")
	
	ret, err := parser.ParseURL(this.Ctx.Request().RequestURI)
	if err != nil {
		return
	}

	for k, v := range ret {
		iris.New().Logger().Info(k, v)
	}

	acc := ret["account"].(string)
	pwd := ret["pwd"].(string)

	user := backdb.LoadUser(acc, pwd)
	if nil == user {
		this.Ctx.Writef("not this role.")
		return
	}

	//todo: ...

	this.Ctx.Writef(errorCode.SUCCESS)

}

func (this *RootControl) register(){
	
	iris.New().Logger().Info("GET --> register")

	ret, err := parser.ParseURL(this.Ctx.Request().RequestURI)
	if err != nil {
		return
	}

	for k, v := range ret {
		iris.New().Logger().Info(k, v)
	}

	acc := ret["account"].(string)
	pwd := ret["pwd"].(string)

	if backdb.FindUser(acc) {
		this.Ctx.Writef("this is old role.")
		return
	}

	user := backdb.NewUser(utils.NewInt64_v1(), acc, pwd)
	user.Update()

	this.Ctx.Writef(errorCode.SUCCESS)

}