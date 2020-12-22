package webNet

import (
	"fmt"
	"go-snake/akmessage"
	"go-snake/common/messageBase"

	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

type WsCallback func(ws *WebSession, msg proto.Message)

func SendMsg(mainID akmessage.MSG, msg proto.Message, ws *WebSession) {
	out := messageBase.CSPackMsg_pb(mainID, msg)
	ws.Write(websocket.BinaryMessage, out)
}

func RecvMessage(sess *WebSession, msg *wsMessage) {
	pack := messageBase.CSPackTool()
	err := pack.UnPack(msg.data)
	if err != nil {
		akLog.Error("message unpack fail.")
		return
	}
	akLog.FmtPrintln("receive msgid: ", pack.GetMsgID())
	if handler := messageBase.MsgHandler(pack.GetMsgID()); handler != nil {
		handler.Call_pb(pack.GetData())
	} else {
		fmt.Println("invalid msg id: ", pack.GetMsgID())
	}
}

func Route(sess *WebSession, msg *wsMessage) {

}
