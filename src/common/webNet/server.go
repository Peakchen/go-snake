package webNet

//
//from https://www.godoc.org/github.com/gorilla/websocket
//

import (
	"fmt"
	"go-snake/common"
	"go-snake/common/mixNet"
	"net/http"

	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/gorilla/websocket"
	//"strings"
	//"strconv"
)

type WebSocketSvr struct {
	Addr    string
	sessmgr mixNet.SessionMgrIf
}

func NewWebsocketSvr(addr string) {
	ws := &WebSocketSvr{
		Addr:    addr,
		sessmgr: mixNet.GetSessionMgr(),
	}
	ws.run()
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024 * 4,
	WriteBufferSize: 1024 * 4,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (this *WebSocketSvr) sessionHandler(resp http.ResponseWriter, req *http.Request) {

	wsSocket, err := upgrader.Upgrade(resp, req, nil)
	if err != nil {
		fmt.Println("upgrader websocket fail, err: ", err.Error())
		return
	}

	if this.sessmgr.IsClose() {
		akLog.Info("server close ws...")
		return
	}

	NewWebSession(wsSocket, this.sessmgr)
}

func (this *WebSocketSvr) run() {
	http.HandleFunc("/echo", this.sessionHandler)
	common.DosafeRoutine(func() {
		http.ListenAndServe(this.Addr, nil)
	}, nil)
}
