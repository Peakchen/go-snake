package akWebNet

/*
	初级消息处理
*/

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type receiveMsgProc func(sess *WebSession, data []byte)

var (
	_procMsgs = map[int]receiveMsgProc{
		websocket.TextMessage:   TextMessageFunc,
		websocket.BinaryMessage: BinaryMessageFunc,
		websocket.CloseMessage:  CloseMessageFunc,
		websocket.PingMessage:   PingMessageFunc,
		websocket.PongMessage:   PongMessageFunc,
	}
)

func GetMessageHandler(id int) receiveMsgProc {
	return _procMsgs[id]
}

func TextMessageFunc(sess *WebSession, data []byte) {
	fmt.Println("read TextMessage data: ", string(data))
	sess.Write(websocket.TextMessage, []byte("hello,too!"))
}

func BinaryMessageFunc(sess *WebSession, data []byte) {
	MsgProc(sess, data, PACK_PROTO)
}

func CloseMessageFunc(sess *WebSession, data []byte) {
	fmt.Println("read CloseMessage.")
	sess.offch <- sess
}

func PingMessageFunc(sess *WebSession, data []byte) {
	fmt.Println("read PingMessage.")
}

func PongMessageFunc(sess *WebSession, data []byte) {
	fmt.Println("read PongMessage.")
}
