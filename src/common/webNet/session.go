package webNet

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"

	//"encoding/binary"
	//"io"
	//"strings"
	"go-snake/common"
	"go-snake/common/messageBase"
	"go-snake/common/mixNet"

	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/utils"
)

type wsMessage struct {
	messageType int    //消息类型 TextMessage/BinaryMessage/CloseMessage/PingMessage/PongMessage
	data        []byte //消息内容
}

type WebSession struct {
	wsconn     *websocket.Conn
	RemoteAddr string
	writeCh    chan *wsMessage
	sessionID  string
	sessmgr    mixNet.SessionMgrIf
	msgprocs   map[uint32]*messageBase.TMessageProc
	uid        int64
	status     uint32
}

func NewWebSession(conn *websocket.Conn, mgr mixNet.SessionMgrIf) *WebSession {
	sess := &WebSession{
		wsconn:     conn,
		RemoteAddr: conn.RemoteAddr().String(),
		writeCh:    make(chan *wsMessage, maxWriteMsgSize),
		sessionID:  utils.GetUUID(),
		sessmgr:    mgr,
		msgprocs:   messageBase.GetMsgHandlers(),
	}
	sess.sessmgr.AddWebSession(sess.GetSessionID(), sess)
	sess.Handle()
	return sess
}

func (this *WebSession) GetSessionID() string {
	return this.sessionID
}

func (this *WebSession) Handle() {
	this.SetConnected()

	common.DosafeRoutine(this.readloop, func() {
		akLog.FmtPrintln("read exit.")
	})
	common.DosafeRoutine(this.writeloop, func() {
		akLog.FmtPrintln("write exit.")
	})
}

func (this *WebSession) close() {
	this.Stop()
	this.sessmgr.RemoveWebSession(this.sessionID)
}

func (this *WebSession) Stop() {
	this.SetClose()
	time.Sleep(time.Second)
	this.wsconn.Close()

	if len(this.writeCh) > 0 {
		close(this.writeCh)
	}
}

func (this *WebSession) readloop() {

	defer func() {
		this.close()
	}()

	this.wsconn.SetReadLimit(maxMessageSize)
	for {
		select {
		default:
			if this.GetStatus() == messageBase.CLOSED {
				akLog.Info("session close...")
				return
			}
			common.Dosafe(func() {
				this.wsconn.SetReadDeadline(time.Now().Add(pongWait))
				msgType, data, err := this.wsconn.ReadMessage()
				if err != nil {
					akLog.Error("msg read fail, err: ", err.Error())
					this.close()
					websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure)
					return
				}

				this.read(&wsMessage{
					messageType: msgType,
					data:        data,
				})
			}, nil)
		}
	}
}

func (this *WebSession) read(content *wsMessage) {
	if handler := GetMessageHandler(content.messageType); handler != nil {
		handler(this, content)
	} else {
		panic(fmt.Errorf("invalid message type: %v.", content.messageType))
	}
}

func (this *WebSession) writeloop() {
	ticker := time.NewTicker(pingPeriod)
	deadline := time.Duration(pingPeriod / 2)
	defer func() {
		ticker.Stop()
		this.close()
	}()

	for {
		select {
		case msg := <-this.writeCh:
			if this.GetStatus() == messageBase.CLOSED {
				akLog.Info("session close...")
				return
			}
			common.Dosafe(func() {
				this.wsconn.SetWriteDeadline(time.Now().Add(writeWait))
				if err := this.wsconn.WriteMessage(msg.messageType, msg.data); err != nil {
					akLog.Error("send msg fail, err: ", err.Error(), len(this.writeCh))
					this.close()
					return
				}
			}, nil)
		case <-ticker.C:
			common.Dosafe(func() {
				if err := this.wsconn.WriteControl(websocket.PingMessage, []byte("ping"), time.Now().Add(deadline)); err != nil {
					akLog.Error("send ping, err: ", err.Error())
					this.close()
					return
				}
			}, nil)
		}
	}
}

func (this *WebSession) sendOffline() {
	/* send message to close connetion... */
	if err := this.wsconn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "now closing..."), time.Now().Add(time.Second)); err != nil {
		akLog.FmtPrintln("send close fail, err: ", err)
		return
	}
}

func (this *WebSession) Write(msgtype int, data []byte) {
	akLog.FmtPrintln("session writed channel data len: ", len(this.writeCh), common.SizeVal(this.writeCh), time.Now().Unix())
	if this.GetStatus() == messageBase.CLOSED {
		akLog.Info("session close...")
		return
	}

	this.writeCh <- &wsMessage{
		messageType: msgtype,
		data:        data,
	}
}

func (this *WebSession) GetSessionMgr() mixNet.SessionMgrIf {
	return this.sessmgr
}

func (this *WebSession) Bind(uid int64) { this.uid = uid }

func (this *WebSession) GetUID() int64 { return this.uid }

func (this *WebSession) SetClose() {
	atomic.StoreUint32(&this.status, messageBase.CLOSED)
}

func (this *WebSession) SetConnected() {
	atomic.StoreUint32(&this.status, messageBase.CONNECTED)
}

func (this *WebSession) GetStatus() uint32 {
	return atomic.LoadUint32(&this.status)
}
