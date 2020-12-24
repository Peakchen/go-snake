package wscli

import (
	"go-snake/akmessage"
	"go-snake/common"
	"go-snake/common/messageBase"
	"go-snake/robot/option"
	"net/url"
	"reflect"
	"sync/atomic"
	"time"

	"github.com/Peakchen/xgameCommon/akLog"

	"github.com/gorilla/websocket"
	"github.com/panjf2000/ants/v2"
)

type WsNet struct {
	host     string
	c        *websocket.Conn
	fnCh     chan func()
	sendCh   chan []byte
	status   uint32
	lostData chan []byte

	opts *option.RobotOption
}

func NewClient(host string, optFns ...option.RobotOptionFn) {
	ws := &WsNet{host: host,
		sendCh:   make(chan []byte, 100),
		lostData: make(chan []byte, 10),
		fnCh:     make(chan func(), 100),
		opts:     option.SortRobotOptions(optFns...),
	}

	err := ws.dail()
	if err != nil {
		akLog.Fail("dial fail:", err)
		return
	}

	common.DosafeRoutine(func() { ws.readloop() }, nil)
	common.DosafeRoutine(func() { ws.writeloop() }, nil)
	common.DosafeRoutine(func() { ws.checkconnect() }, nil)
	common.DosafeRoutine(func() { ws.heartbeat() }, nil)
}

func (this *WsNet) dail() error {
	this.SetConnecting()

	u := url.URL{Scheme: "ws", Host: this.host, Path: "/echo"}
	akLog.FmtPrintf("connecting to %s", u.String())

	dstc, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}

	this.c = dstc
	this.SetConnected()
	return nil
}

func (this *WsNet) close() {
	if this.GetStatus() == messageBase.CLOSED {
		return
	}
	this.c.Close()
	this.SetClose()
}

func (this *WsNet) readloop() {
	for {
		select {
		case fn := <-this.fnCh:
			_ = ants.Submit(fn)
			//common.Dosafe(fn, nil)
		default:
			if this.GetStatus() == messageBase.CONNECTED {

				this.c.SetReadDeadline(time.Now().Add(40 * time.Second))

				_, data, err := this.c.ReadMessage()
				if err != nil {
					akLog.Error("read:", err)
					this.close()
					continue
				}

				akLog.FmtPrintln("ReadMessage...")
				if this.opts.ModelsRecv != nil {
					this.fnCh <- func() {
						cspt := messageBase.CSPackTool()
						err = cspt.UnPack(data)
						if err != nil {
							akLog.Error("cs unpack fail.")
							return
						}
						this.opts.ModelsRecv([]reflect.Value{
							reflect.ValueOf(cspt.GetMsgID()),
							reflect.ValueOf(cspt.GetData())})
					}
				}
			}

		}
	}
}

func (this *WsNet) writeloop() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			akLog.FmtPrintln("send test, status: ", this.GetStatus())

			if this.GetStatus() == messageBase.CONNECTED && this.opts.ModelsRun != nil {
				this.opts.ModelsRun(reflect.ValueOf(this))
			}

		case data := <-this.sendCh:

			if this.GetStatus() == messageBase.CONNECTED {
				this.c.SetWriteDeadline(time.Now().Add(15 * time.Second))
				err := this.c.WriteMessage(websocket.BinaryMessage, data)
				if err != nil {
					akLog.Error("write fail: ", err)
					this.close()
					if len(this.lostData) < cap(this.lostData) {
						this.lostData <- data
					}
				}
			}
		}
	}
}

func (this *WsNet) checkconnect() {
	tick := time.NewTicker(3 * time.Second)
	defer tick.Stop()

	for {
		select {
		case <-tick.C:
			akLog.FmtPrintln("checkconnect status: ", this.GetStatus())
			if this.GetStatus() == messageBase.CLOSED || this.GetStatus() == messageBase.CONNECTING {
				this.dail()
			}
			if this.GetStatus() == messageBase.CONNECTED {
				if len(this.lostData) > 0 {
					for d := range this.lostData {
						this.sendCh <- d
					}
				}
			}
		}
	}
}

func (this *WsNet) heartbeat() {
	tick := time.NewTicker(30 * time.Second)
	defer tick.Stop()

	for {
		select {
		case <-tick.C:
			if this.GetStatus() == messageBase.CONNECTED {
				this.sendHeartBeatMsg()
			}
		}
	}
}

func (this *WsNet) sendHeartBeatMsg() {
	hb := &akmessage.CS_HeartBeat{}
	data := messageBase.CSPackMsg_pb(akmessage.MSG_CS_HEARTBEAT, hb)
	if data == nil {
		akLog.Error("pack heart beat msg fail.")
		return
	}
	akLog.FmtPrintf("heart beat, id: %v.", uint32(akmessage.MSG_CS_HEARTBEAT))
	this.SendMsg(data)
}

func (this *WsNet) SendMsg(data []byte) {
	if len(this.sendCh) >= cap(this.sendCh) {
		this.lostData <- data
	} else {
		this.sendCh <- data
	}
}

func (this *WsNet) SetClose() {
	atomic.StoreUint32(&this.status, messageBase.CLOSED)
}

func (this *WsNet) SetConnected() {
	atomic.StoreUint32(&this.status, messageBase.CONNECTED)
}

func (this *WsNet) SetConnecting() {
	atomic.StoreUint32(&this.status, messageBase.CONNECTING)
}

func (this *WsNet) GetStatus() uint32 {
	return atomic.LoadUint32(&this.status)
}
