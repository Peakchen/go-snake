// add by stefan

package akNet

import (
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/define"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_MainModule"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_Player"
	"github.com/Peakchen/xgameCommon/pprof"
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"
	"sync/atomic"
)

type TcpServer struct {
	sync.Mutex

	host      string
	pprofAddr string
	listener  *net.TCPListener
	cancel    context.CancelFunc
	off       chan *SvrTcpSession
	session   *SvrTcpSession
	// person online
	person  int32
	SvrType define.ERouteId
	pack    IMessagePack
	// session id
	SessionID uint64
	procName  string
}

func NewTcpServer(listenAddr, pprofAddr string, SvrType define.ERouteId, procName string) *TcpServer {
	return &TcpServer{
		host:      listenAddr,
		pprofAddr: pprofAddr,
		procName:  procName,
		SvrType:   SvrType,
		SessionID: ESessionBeginNum,
		off:       make(chan *SvrTcpSession, maxOfflineSize),
	}
}

func (this *TcpServer) Run() {
	os.Setenv("GOTRACEBACK", "crash")
	tcpAddr, err := net.ResolveTCPAddr("tcp4", this.host)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	this.listener = listener

	var (
		ctx context.Context
		sw  = sync.WaitGroup{}
	)

	ctx, this.cancel = context.WithCancel(context.Background())
	pprof.Run(ctx)

	this.pack = &ServerProtocol{}
	sw.Add(3)
	go this.loop(ctx, &sw)
	go this.loopoff(ctx, &sw)
	go func() {
		akLog.FmtPrintln("[server] run http server, host: ", this.pprofAddr)
		http.ListenAndServe(this.pprofAddr, nil)
	}()
	sw.Wait()
}

func (this *TcpServer) loop(ctx context.Context, sw *sync.WaitGroup) {
	defer func() {
		sw.Done()
		this.Exit(sw)
	}()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			c, err := this.listener.AcceptTCP()
			if err != nil || c == nil {
				akLog.Error("can not accept tcp.")
				continue
			}

			c.SetNoDelay(true)
			c.SetKeepAlive(true)
			atomic.AddUint64(&this.SessionID, 1)
			akLog.FmtPrintf("[server] accept connect here addr: %v, SessionID: %v.", c.RemoteAddr(), this.SessionID)
			this.session = NewSvrSession(c.RemoteAddr().String(), c, ctx, this.SvrType, this.off, this.pack, this.procName)
			this.session.HandleSession(sw)
			this.online()
		}
	}
}

func (this *TcpServer) loopoff(ctx context.Context, sw *sync.WaitGroup) {
	defer func() {
		sw.Done()
		this.Exit(sw)
	}()
	for {
		select {
		case offs, ok := <-this.off:
			if !ok {
				return
			}
			this.offline(offs)
		case <-ctx.Done():
			return
		}
	}
}

func (this *TcpServer) online() {
	this.person++
	// rpc notify person online...

}

func (this *TcpServer) offline(offs *SvrTcpSession) {
	// notify person offline...
	if offs.IsUser() {
		this.person--
		ntf := &MSG_Player.CS_LeaveServer_Req{}
		_, err := offs.GetPack().PackInnerMsg(uint16(MSG_MainModule.MAINMSG_PLAYER), uint16(MSG_Player.SUBMSG_CS_LeaveServer), ntf)
		if err != nil {
			akLog.Error(err)
			return
		}
		sendInnerSvr(offs)
	}
}

func (this *TcpServer) SendMessage() {

}

func (this *TcpServer) Exit(sw *sync.WaitGroup) {
	this.cancel()
	this.listener.Close()
	pprof.Exit()
}

func (this *TcpServer) SessionType() (st ESessionType) {
	return ESessionType_Server
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
