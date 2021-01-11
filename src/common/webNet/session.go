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
	readCh     chan *wsMessage
	sessionID  string
	sessmgr    mixNet.SessionMgrIf
	msgprocs   map[uint32]*messageBase.TMessageProc
	uid        int64
	status     uint32
	stop       chan bool
}

func NewWebSession(conn *websocket.Conn, mgr mixNet.SessionMgrIf) *WebSession {
	sess := &WebSession{
		wsconn:     conn,
		RemoteAddr: conn.RemoteAddr().String(),
		writeCh:    make(chan *wsMessage, maxWriteMsgSize),
		readCh:     make(chan *wsMessage, maxReadMsgSize),
		sessionID:  utils.GetUUID(),
		sessmgr:    mgr,
		msgprocs:   messageBase.GetMsgHandlers(),
		stop:       make(chan bool, 1),
	}
	akLog.Info("ws new session: ", sess.sessionID)
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
	if this.GetStatus() == messageBase.CLOSED {
		return
	}
	this.stop <- true
	this.SetClose()
	//this.wsconn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	time.Sleep(time.Second)
	this.wsconn.Close()
	if len(this.writeCh) > 0 {
		akLog.Info("write left: ", len(this.writeCh))
		close(this.writeCh)
	}
	if len(this.readCh) > 0 {
		akLog.Info("read left: ", len(this.readCh))
		close(this.readCh)
	}
}

func (this *WebSession) readloop() {

	defer func() {
		this.close()
	}()

	//gorilla websocket causes panic if number of errors is >= 1000
	var error_count = 0
	//reset connection before error count exceeds the websocket limit
readloop:
	for error_count <= 500 {
		select {
		case <-this.stop:
			break readloop
		case rd := <-this.readCh:
			common.Dosafe(func() { this.read(rd) }, this.close)
		default:
			common.Dosafe(func() {
				this.wsconn.SetReadLimit(maxMessageSize)
				this.wsconn.SetReadDeadline(time.Now().Add(pongWait * 2))

				msgType, data, err := this.wsconn.ReadMessage()
				if err != nil {
					error_count++
					this.close()
					akLog.Info("msg read fail:", err.Error(), this.GetSessionID())
					websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure)
					return
				}

				this.readCh <- &wsMessage{
					messageType: msgType,
					data:        data,
				}

			}, this.close)
		}
	}
	akLog.Error("error_count: ", error_count)
}

func (this *WebSession) read(content *wsMessage) {
	akLog.FmtPrintln("read: ", content.messageType)
	if handler := GetMessageHandler(content.messageType); handler != nil {
		handler(this, content)
	} else {
		panic(fmt.Errorf("invalid message type: %v.", content.messageType))
	}
}

func (this *WebSession) writeloop() {

	common.Dosafe(func() {
		defer this.close()

	writeloop:
		for {
			select {
			case <-this.stop:
				break writeloop
			case msg := <-this.writeCh:

				this.wsconn.SetWriteDeadline(time.Now().Add(writeWait))
				if err := this.wsconn.WriteMessage(msg.messageType, msg.data); err != nil {
					akLog.Info("send msg fail: ", this.GetSessionID(), err.Error(), len(this.writeCh))
					this.close()
					return
				}

			}
		}
	}, nil)
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
	//return atomic.LoadUint32(&this.status)
	return this.status
}
