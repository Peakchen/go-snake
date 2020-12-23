package webNet

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"

	//"encoding/binary"
	//"io"
	//"strings"
	"go-snake/common"
	"go-snake/common/messageBase"
	"go-snake/common/mixNet"
	"sync"

	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/utils"
)

type wsMessage struct {
	messageType int    //消息类型 TextMessage/BinaryMessage/CloseMessage/PingMessage/PongMessage
	data        []byte //消息内容
}

type WebSession struct {
	wsconn     *websocket.Conn
	offch      chan bool
	RemoteAddr string
	writeCh    chan *wsMessage
	sessionID  string
	wg         sync.WaitGroup
	sessmgr    mixNet.SessionMgrIf
	msgprocs   map[uint32]*messageBase.TMessageProc
	uid        int64
}

func NewWebSession(conn *websocket.Conn, mgr mixNet.SessionMgrIf) *WebSession {
	sess := &WebSession{
		wsconn:     conn,
		offch:      make(chan bool, 1),
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
	this.wg.Add(2)
	common.DosafeRoutine(func() { this.readloop() }, func() {
		akLog.FmtPrintln("read exit.")
	})
	common.DosafeRoutine(func() { this.writeloop() }, func() {
		akLog.FmtPrintln("write exit.")
	})
}

func (this *WebSession) exit() {
	this.offch <- true

	this.wg.Wait()
	if len(this.writeCh) > 0 {
		close(this.writeCh)
	}

	this.sessmgr.RemoveWebSession(this.sessionID)
	this.wsconn.Close()
}

func (this *WebSession) readloop() {

	defer func() {
		this.wg.Done()
		this.exit()
	}()

	this.wsconn.SetReadLimit(maxMessageSize)
	for {
		select {
		case <-this.offch:
			return
		default:
			this.wsconn.SetReadDeadline(time.Now().Add(pongWait))
			msgType, data, err := this.wsconn.ReadMessage()
			if err != nil {
				websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure)
				fmt.Println("msg read fail, err: ", err.Error(), time.Now().Unix())
				return
			}

			this.read(&wsMessage{
				messageType: msgType,
				data:        data,
			})
		}
	}
}

func (this *WebSession) read(content *wsMessage) {
	akLog.FmtPrintln("read messageType: ", content.messageType, len(content.data), time.Now().Unix())
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
		this.wg.Done()
		this.exit()
	}()

	for {
		select {
		case <-this.offch:
			return
		case msg := <-this.writeCh:
			this.wsconn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := this.wsconn.WriteMessage(msg.messageType, msg.data); err != nil {
				fmt.Println("send msg fail, err: ", err.Error(), time.Now().Unix())
				return
			}
		case <-ticker.C:
			if err := this.wsconn.WriteControl(websocket.PingMessage, []byte("ping"), time.Now().Add(deadline)); err != nil {
				fmt.Println("send msg over time, err: ", err.Error(), time.Now().Unix())
				return
			}
		}
	}
}

func (this *WebSession) sendOffline() {
	/* send message to close connetion... */
	if err := this.wsconn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "now closing..."), time.Now().Add(time.Second)); err != nil {
		fmt.Println("send close fail, err: ", err)
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
