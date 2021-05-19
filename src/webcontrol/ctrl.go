package webcontrol

import (
	"go-snake/akmessage"
	"go-snake/app/in"
	"go-snake/app/application"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"go-snake/webcontrol/route"
	"go-snake/webcontrol/controller"
	"go-snake/webcontrol/rpcBase"
	"github.com/kataras/iris/v12/middleware/logger"
    "github.com/kataras/iris/v12/middleware/recover"
)

type WebControl struct {
}

func New(name string) *WebControl {
	
	application.SetAppName(name)

	return &WebControl{}
}

func (this *WebControl) Init() {

}

func (this *WebControl) Type() akmessage.ServerType {
	return akmessage.ServerType_WebControl
}

func (this *WebControl) Run(d *in.Input) {
	
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.Use(recover.New())
    app.Use(logger.New())
	
	tmpl := iris.HTML("./webDir", ".html") //Layout("layout.html") 看情况加
	app.RegisterView(tmpl)
	app.HandleDir("/webDir", "./webDir")

	route.Register(app)
	mvc.Configure(app, controller.Basic)

	rpcBase.RunRpc(d.Scfg.EtcdIP, d.Scfg.EtcdNodeIP)

	app.Run(iris.Addr(d.Scfg.TCPHost),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
    )

}
