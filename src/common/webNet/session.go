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
	"sync"
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
	readStatus uint32
	writeStatus uint32
	stop       chan bool
	wg         sync.WaitGroup
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

	this.wg.Add(2)

	common.DosafeRoutine(this.readloop, func() {
		akLog.FmtPrintln("read exit.")
	})

	common.DosafeRoutine(this.writeloop, func() {
		akLog.FmtPrintln("write exit.")
	})

	this.wg.Wait()
	this.close()

}

func (this *WebSession) close() {
	this.Stop()
	this.sessmgr.RemoveWebSession(this.sessionID)
}

func (this *WebSession) Stop() {

	if this.writeStatus != messageBase.CLOSED || this.readStatus != messageBase.CLOSED{
		return
	}

	this.stop <- true
	
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
		
		this.wg.Done()

	}()

	//gorilla websocket causes panic if number of errors is >= 1000
	var error_count = 0
	//reset connection before error count exceeds the websocket limit

	for error_count <= 500 {

		select {

		case <-this.stop:
			return

		case rd := <-this.readCh:

			if this.readStatus == messageBase.CLOSED {
				return
			}

			common.Dosafe(func() { this.read(rd) }, nil)

		default:

			if this.readStatus == messageBase.CLOSED {
				return
			}

			common.Dosafe(func() {

				this.wsconn.SetReadLimit(maxMessageSize)
				this.wsconn.SetReadDeadline(time.Now().Add(pongWait * 2))

				msgType, data, err := this.wsconn.ReadMessage()
				if err != nil {

					error_count++
					this.SetReadClose()

					akLog.Info("msg read fail:", err.Error(), this.GetSessionID())
					websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure)
					return
				}

				this.readCh <- &wsMessage{
					messageType: msgType,
					data:        data,
				}

			}, nil)
		}

	}

	akLog.Error("error_count: ", error_count)

}

func (this *WebSession) read(content *wsMessage) {
	akLog.FmtPrintln("read: ", content.messageType, len(content.data))
	if handler := GetMessageHandler(content.messageType); handler != nil {
		handler(this, content)
	} else {
		panic(fmt.Errorf("invalid message type: %v.", content.messageType))
	}
}

func (this *WebSession) writeloop() {

	common.Dosafe(func() {
		
		defer func(){
			
			this.wg.Done()

		}()

	writeloop:
		for {
			select {
			case <-this.stop:

				break writeloop

			case msg := <-this.writeCh:

				if this.writeStatus == messageBase.CLOSED {
					return
				}

				this.wsconn.SetWriteDeadline(time.Now().Add(writeWait))
				if err := this.wsconn.WriteMessage(msg.messageType, msg.data); err != nil {
					akLog.Info("send msg fail: ", this.GetSessionID(), err.Error(), len(this.writeCh))
					this.SetWriteClose()
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

func (this *WebSession) SetReadClose() {
	atomic.StoreUint32(&this.readStatus, messageBase.CLOSED)
}

func (this *WebSession) SetWriteClose() {
	atomic.StoreUint32(&this.writeStatus, messageBase.CLOSED)
}

func (this *WebSession) SetConnected() {
	atomic.StoreUint32(&this.writeStatus, messageBase.CONNECTED)
	atomic.StoreUint32(&this.readStatus, messageBase.CONNECTED)
}
