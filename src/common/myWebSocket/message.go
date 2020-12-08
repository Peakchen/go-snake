package myWebSocket

import (
	"ak-remote/akmessage"
	"ak-remote/common/messageBase"
	"fmt"

	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

type WsCallback func(ws *WebSession, msg proto.Message)

func SendMsg(mainID akmessage.MSG, msg proto.Message, ws *WebSession) {
	data, err := proto.Marshal(msg)
	if err != nil {
		akLog.Error(fmt.Errorf("proto marshal fail, mainid: %v, err: %v.", mainID, err))
		return
	}

	pack := messageBase.PackTool()
	pack.Init(uint32(mainID), uint32(len(data)), data)

	out := make([]byte, len(data)+messageBase.MSG_PACK_NODATA_SIZE)
	pack.Pack(out)
	ws.Write(websocket.BinaryMessage, out)
}

func RecvMessage(sess *WebSession, msg *wsMessage) {
	pack := messageBase.PackTool()
	err := pack.UnPack(msg.data)
	if err != nil {
		akLog.Error("message unpack fail.")
		return
	}

	akLog.FmtPrintln("receive msgid: ", pack.GetMsgID())
	if handler := messageBase.MsgHandler(pack.GetMsgID()); handler != nil {
		handler.Call(pack.GetData())
	} else {
		fmt.Println("invalid msg id: ", pack.GetMsgID())
	}
}

func Route(sess *WebSession, msg *wsMessage) {

}
