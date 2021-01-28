package akNet

// client connect server.
// add by stefan

import (
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/define"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_MainModule"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_Server"
	"github.com/Peakchen/xgameCommon/pprof"

	//"time"
	"context"
)

type TcpClient struct {
	sync.Mutex

	cancel    context.CancelFunc
	host      string
	pprofAddr string
	dialsess  *ClientTcpSession
	// person offline flag
	off chan *ClientTcpSession
	// person online
	person   int32
	SvrType  define.ERouteId
	Adacb    AfterDialAct
	mpobj    IMessagePack
	procName string
}

func NewClient(host, pprofAddr string, SvrType define.ERouteId, Ada AfterDialAct, procName string) *TcpClient {
	return &TcpClient{
		host:      host,
		pprofAddr: pprofAddr,
		SvrType:   SvrType,
		Adacb:     Ada,
		off:       make(chan *ClientTcpSession, maxOfflineSize),
		procName:  procName,
	}
}

func (this *TcpClient) Run() {
	os.Setenv("GOTRACEBACK", "crash")

	var (
		ctx context.Context
		sw  = sync.WaitGroup{}
	)

	ctx, this.cancel = context.WithCancel(context.Background())
	this.mpobj = &ClientProtocol{}
	this.connect(ctx, &sw)
	pprof.Run(ctx)
	sw.Add(3)
	go this.loopconn(ctx, &sw)
	go this.loopoff(ctx, &sw)
	go func() {
		akLog.FmtPrintln("[client] run http server, host: ", this.pprofAddr)
		http.ListenAndServe(this.pprofAddr, nil)
	}()
	sw.Wait()
}

func (this *TcpClient) connect(ctx context.Context, sw *sync.WaitGroup) (err error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", this.host)
	if err != nil {
		akLog.Error("resolve tcp error: %v.", err.Error())
		return err
	}

	c, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		akLog.Error("net dial err: %v.", err)
		return err
	}

	akLog.FmtPrintf("[----------client-----------] addr: %v, svrtype: %v.", c.RemoteAddr(), this.SvrType)
	c.SetNoDelay(true)
	this.dialsess = NewClientSession(c.RemoteAddr().String(), c, ctx, this.SvrType, this.off, this.mpobj, this.procName)
	this.dialsess.HandleSession(sw)
	this.afterDial()
	return nil
}

func (this *TcpClient) loopconn(ctx context.Context, sw *sync.WaitGroup) {
	defer func() {
		sw.Done()
		this.Exit(sw)
	}()

	ticker := time.NewTicker(time.Duration(cstClientSessionCheckMs) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if this.dialsess == nil || false == this.dialsess.isAlive {
				if err := this.connect(ctx, sw); err != nil {
					akLog.FmtPrintf("dail to server fail, host: %v.", this.host)
				} else {
					// recover send cache msg.

				}
			}
		}
	}
}

func (this *TcpClient) loopoff(ctx context.Context, sw *sync.WaitGroup) {
	defer func() {
		sw.Done()
		this.Exit(sw)
	}()

	for {
		select {
		case os, ok := <-this.off:
			if ok {
				this.offline(os)
			}

		case <-ctx.Done():
			return
		}
	}
}

func (this *TcpClient) offline(os *ClientTcpSession) {
	// process
	this.dialsess.isAlive = false
}

func (this *TcpClient) Send(data []byte) {
	if !this.dialsess.isAlive {
		return
	}
	this.dialsess.SetSendCache(data)
}

func (this *TcpClient) SendMessage() {

}

func (this *TcpClient) Exit(sw *sync.WaitGroup) {
	this.dialsess = nil
	this.cancel()
	pprof.Exit()
}

func (this *TcpClient) sendRegisterMsg() {
	akLog.FmtPrintf("after dial, send point: %v register message to server.", this.SvrType)
	req := &MSG_Server.CS_ServerRegister_Req{}
	req.ServerType = int32(this.SvrType)
	req.Msgs = GetAllMessageIDs()
	akLog.FmtPrintln("register context: ", req.Msgs)
	buff, err := this.mpobj.PackClientMsg(uint16(MSG_MainModule.MAINMSG_SERVER), uint16(MSG_Server.SUBMSG_CS_ServerRegister), req)
	if err != nil {
		akLog.Error(err)
		return
	}
	this.Send(buff)
}

func (this *TcpClient) afterDial() {
	if this.Adacb != nil {
		this.Adacb(this.dialsess)
	}
	this.sendRegisterMsg()
}

func (this *TcpClient) SessionType() (st ESessionType) {
	return ESessionType_Client
}
