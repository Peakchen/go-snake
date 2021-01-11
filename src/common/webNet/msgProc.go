package webNet

/*
	初级消息处理
*/

import (
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/gorilla/websocket"
)

type receiveMsgProc func(sess *WebSession, msg *wsMessage)

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

func TextMessageFunc(sess *WebSession, msg *wsMessage) {
	akLog.Info("read TextMessage data: ", string(msg.data))
	sess.Write(websocket.TextMessage, []byte("hello,too!"))
}

func BinaryMessageFunc(sess *WebSession, msg *wsMessage) {
	sess.GetSessionMgr().CS_SendInner(sess.GetSessionID(), 0, msg.data)
	//RecvMessage(sess, msg)
}

func CloseMessageFunc(sess *WebSession, msg *wsMessage) {
	akLog.Info("close CloseMessage.")
}

func PingMessageFunc(sess *WebSession, msg *wsMessage) {
	akLog.Info("ping PingMessage.")
}

func PongMessageFunc(sess *WebSession, msg *wsMessage) {
	akLog.Info("pong PongMessage.")
}
