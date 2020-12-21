package webNet

//
//from https://www.godoc.org/github.com/gorilla/websocket
//

import (
	"fmt"
	"go-snake/common"
	"go-snake/common/mixNet"
	"net/http"
	"time"

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
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
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

	NewWebSession(wsSocket, this.sessmgr)
	fmt.Println("connect ws socket: ", time.Now().Unix())
}

func (this *WebSocketSvr) run() {
	http.HandleFunc("/echo", this.sessionHandler)
	common.DosafeRoutine(func() {
		http.ListenAndServe(this.Addr, nil)
	}, nil)
}
