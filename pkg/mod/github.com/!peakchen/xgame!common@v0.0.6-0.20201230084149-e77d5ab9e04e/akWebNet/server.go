package akWebNet

//
//from https://www.godoc.org/github.com/gorilla/websocket
//

import (
	"context"
	"net/http"
	"sync"

	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/aktime"
	"github.com/Peakchen/xgameCommon/define"
	"github.com/Peakchen/xgameCommon/pprof"
	"github.com/gorilla/websocket"
	//"strings"
	//"strconv"
)

type WebSocketSvr struct {
	Addr      string
	pprofAddr string
	offch     chan *WebSession //离线通道
	cancel    context.CancelFunc
	actor     define.ERouteId
}

func NewWebsocketSvr(addr string, pprofAddr string, actor define.ERouteId) *WebSocketSvr {
	return &WebSocketSvr{
		Addr:      addr,
		offch:     make(chan *WebSession, 1024),
		pprofAddr: pprofAddr,
		actor:     actor,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (this *WebSocketSvr) wsSvrHandler(resp http.ResponseWriter, req *http.Request) {
	wsSocket, err := upgrader.Upgrade(resp, req, nil)
	if err != nil {
		akLog.Error("upgrader websocket fail, err: ", err.Error())
		return
	}

	sess := NewWebSession(wsSocket, this.offch, &TActor{
		Route:     &MsgRoute{},
		ActorType: this.actor,
	})
	sess.Handle()
	akLog.FmtPrintln("connect ws socket: ", sess.RemoteAddr, aktime.Now().Unix())
}

func (this *WebSocketSvr) Run() {
	http.HandleFunc("/echo", this.wsSvrHandler)
	var ctx context.Context
	ctx, this.cancel = context.WithCancel(context.Background())
	pprof.Run(ctx)
	var sw sync.WaitGroup
	sw.Add(1)
	go loopSignalCheck(ctx, &sw)
	go func() {
		akLog.FmtPrintln("run server, listen host: ", this.Addr)
		http.ListenAndServe(this.Addr, nil)
	}()
	go func() {
		akLog.FmtPrintln("run pprof http server, host: ", this.pprofAddr)
		http.ListenAndServe(this.pprofAddr, nil)
	}()
	sw.Wait()
	this.exit()
}

func (this *WebSocketSvr) exit() {
	close(this.offch)
	this.cancel()
}
