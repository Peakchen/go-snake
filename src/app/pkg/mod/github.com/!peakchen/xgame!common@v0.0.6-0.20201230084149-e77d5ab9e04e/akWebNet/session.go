package akWebNet

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"

	//"encoding/binary"
	//"io"
	//"strings"
	"sync"

	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/stacktrace"
	"github.com/Peakchen/xgameCommon/utils"
)

type wsMessage struct {
	messageType int    //消息类型 TextMessage/BinaryMessage/CloseMessage/PingMessage/PongMessage
	data        []byte //消息内容
}

type WebSession struct {
	wsconn     *websocket.Conn
	offch      chan *WebSession //离线通道
	RemoteAddr string
	writeCh    chan *wsMessage //写通道
	readCh     chan *wsMessage //读通道
	IdCh       *uint32
	stopWrite  bool
	stopRead   bool
	one        sync.Once
	actor      *TActor
}

func NewWebSession(conn *websocket.Conn, off chan *WebSession, actor *TActor) *WebSession {
	return &WebSession{
		wsconn:     conn,
		offch:      off,
		RemoteAddr: conn.RemoteAddr().String(),
		writeCh:    make(chan *wsMessage, maxWriteMsgSize),
		readCh:     make(chan *wsMessage, maxWriteMsgSize),
		IdCh:       new(uint32),
		actor:      actor,
	}
}

func (this *WebSession) SetId(id uint32) {
	*this.IdCh = id
}

func (this *WebSession) GetId() uint32 {
	return *(this.IdCh)
}

func (this *WebSession) Handle() {
	akLog.FmtPrintln("read messageType: ", messageType, len(data), time.Now().Unix())
	go this.readloop()
	go this.writeloop()
	go this.heartbeatloop()
}

func (this *WebSession) heartbeatloop() {
	ticker := time.NewTicker(time.Duration(cstKeepLiveHeartBeatSec) * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			sendHeartBeat(this)
		}
	}
}

func (this *WebSession) offline() {
	akLog.FmtPrintf("exit ws socket, actor: %v, RemoteAddr: %v, time: %v.", this.actor.GetActorType(), this.RemoteAddr, time.Now().Unix())
	GwebSessionMgr.RemoveSession(this.RemoteAddr)
	//notify offline ... logout
}

func (this *WebSession) exit() {
	akLog.Error("exit log: ", stacktrace.NormalStackLog())
	//conn close
	this.one.Do(func() {
		this.offline()
		//this.offch <-this

		if _, noclosed := <-this.writeCh; noclosed {
			this.stopWrite = true
			close(this.writeCh)
		}

		if _, noclosed := <-this.readCh; noclosed {
			this.stopRead = true
			close(this.readCh)
		}

		this.wsconn.Close()
	})
}

func (this *WebSession) readloop() {

	defer func() {
		this.exit()
	}()

	this.wsconn.SetReadLimit(maxMessageSize)
	for {
		this.wsconn.SetReadDeadline(time.Now().Add(pongWait))
		msgType, data, err := this.wsconn.ReadMessage()
		if err != nil {
			websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure)
			akLog.Error("msg read fail, err: ", err.Error(), time.Now().Unix())
			return
		}

		if this.stopRead {
			return
		}

		go this.read(msgType, data)
	}
}

func (this *WebSession) read(messageType int, data []byte) {
	akLog.FmtPrintln("read messageType: ", messageType, len(data), time.Now().Unix())
	if handler := GetMessageHandler(messageType); handler != nil {
		handler(this, data)
	} else {
		panic(fmt.Errorf("invalid message type: %v.", messageType))
	}
}

func (this *WebSession) writeloop() {
	ticker := time.NewTicker(pingPeriod)
	deadline := time.Duration(pingPeriod / 2)
	defer func() {
		ticker.Stop()
		this.exit()
	}()

	for {
		select {
		case msg := <-this.writeCh:
			this.wsconn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := this.wsconn.WriteMessage(msg.messageType, msg.data); err != nil {
				akLog.Error("send msg fail, err: ", err.Error(), time.Now().Unix())
				return
			}
		case <-ticker.C:
			if err := this.wsconn.WriteControl(websocket.PingMessage, []byte("ping"), time.Now().Add(deadline)); err != nil {
				akLog.Error("send msg over time, err: ", err.Error(), time.Now().Unix())
				return
			}
		}
	}
}

func (this *WebSession) sendOffline() {
	/* send message to close connetion... */
	if err := this.wsconn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "now closing..."), time.Now().Add(time.Second)); err != nil {
		akLog.Error("send close fail, err: ", err)
		return
	}
}

func (this *WebSession) Write(msgtype int, data []byte) {
	if this.stopWrite {
		return
	}

	akLog.FmtPrintln("session writed channel data len: ", len(this.writeCh), utils.SizeVal(this.writeCh), time.Now().Unix())
	this.writeCh <- &wsMessage{
		messageType: msgtype,
		data:        data,
	}
}

func (this *WebSession) broadcast(msgtype int, data []byte) {
	sesses := GwebSessionMgr.GetSessions()
	sesses.Range(func(k, v interface{}) bool {
		if v != nil {
			sess := v.(*WebSession)
			sess.Write(msgtype, data)
		}

		return true
	})
}

func (this *WebSession) GetActor() *TActor {
	return this.actor
}
