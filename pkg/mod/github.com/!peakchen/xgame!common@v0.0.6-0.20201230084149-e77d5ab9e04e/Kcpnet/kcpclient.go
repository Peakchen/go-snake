package Kcpnet

// by udp

import (
	"context"
	"crypto/sha1"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/define"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_MainModule"
	"github.com/Peakchen/xgameCommon/msgProto/MSG_Server"
	"github.com/Peakchen/xgameCommon/pprof"
	cli "github.com/urfave/cli"
	"github.com/xtaci/kcp-go"
	"golang.org/x/crypto/pbkdf2"
)

type KcpClient struct {
	sw           sync.WaitGroup
	svrName      string
	pack         IMessagePack
	Addr         string
	ppAddr       string
	ctx          context.Context
	cancel       context.CancelFunc
	sesson       *KcpClientSession
	offCh        chan *KcpClientSession
	svrType      define.ERouteId
	exCollection *ExternalCollection
	versionNo    int32
}

func NewKcpClient(addr, pprofAddr string, name string, svrType define.ERouteId, ver int32, exCol *ExternalCollection) *KcpClient {
	return &KcpClient{
		svrName:      name,
		Addr:         addr,
		ppAddr:       pprofAddr,
		offCh:        make(chan *KcpClientSession, 1000),
		svrType:      svrType,
		pack:         &KcpClientProtocol{},
		exCollection: exCol,
		versionNo:    ver,
	}
}

func (this *KcpClient) Run() {
	os.Setenv("GOTRACEBACK", "crash")

	this.ctx, this.cancel = context.WithCancel(context.Background())
	pprof.Run(this.ctx)

	app := &cli.App{
		Name:    this.svrName,
		Usage:   "a server...",
		Version: "v1.0",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "listen",
				Value: this.Addr,
				Usage: "local listen address",
			},
			cli.StringFlag{
				Name:  "key",
				Value: "innergateway",
				Usage: "key",
			},
			cli.StringFlag{
				Name:  "crypt",
				Value: "aes-128",
				Usage: "crypt",
			},
			cli.DurationFlag{
				Name:  "tcp_readDeadline",
				Value: pingPeriod,
				Usage: "tcp_readDeadline",
			},
			cli.DurationFlag{
				Name:  "tcp_writeDeadline",
				Value: pingPeriod,
				Usage: "sockbuf",
			},
			cli.DurationFlag{
				Name:  "udp_readDeadline",
				Value: pingPeriod,
				Usage: "udp_readDeadline",
			},
			cli.DurationFlag{
				Name:  "udp_writeDeadline",
				Value: pingPeriod,
				Usage: "udp_writeDeadline",
			},
			cli.IntFlag{
				Name:  "tcp_sockbuf_w",
				Value: tcpBuffSize,
				Usage: "tcp_sockbuf_w",
			},
			cli.IntFlag{
				Name:  "tcp_sockbuf_r",
				Value: tcpBuffSize,
				Usage: "tcp_sockbuf_r",
			},
			cli.IntFlag{
				Name:  "udp_sockbuf_w",
				Value: udpBuffSize,
				Usage: "udp_sockbuf_w",
			},
			cli.IntFlag{
				Name:  "udp_sockbuf_r",
				Value: udpBuffSize,
				Usage: "udp_sockbuf_r",
			},
			cli.IntFlag{
				Name:  "queuelen",
				Value: queueSize,
				Usage: "queue length",
			},
			cli.IntFlag{
				Name:  "dscp",
				Value: dscp,
				Usage: "dscp",
			},
			cli.IntFlag{
				Name:  "udp-sndwnd",
				Value: 2048,
				Usage: "udp-sndwnd",
			},
			cli.IntFlag{
				Name:  "udp-rcvwnd",
				Value: 2048,
				Usage: "udp-rcvwnd",
			},
			cli.IntFlag{
				Name:  "udp-mtu",
				Value: 1400,
				Usage: "udp-mtu",
			},
			cli.IntFlag{
				Name:  "dataShard",
				Value: 10,
				Usage: "dataShard",
			},
			cli.IntFlag{
				Name:  "parityShards",
				Value: 3,
				Usage: "parityShards",
			},
			cli.IntFlag{
				Name:  "nodelay",
				Value: 1,
				Usage: "nodelay",
			},
			cli.IntFlag{
				Name:  "interval",
				Value: 40,
				Usage: "interval",
			},
			cli.IntFlag{
				Name:  "resend",
				Value: 2,
				Usage: "resend",
			},
			cli.IntFlag{
				Name:  "nc",
				Value: 1,
				Usage: "nc",
			},
		},
		Action: func(c *cli.Context) error {
			akLog.FmtPrintln("action begin...")

			//setup net param
			config := &KcpSvrConfig{
				listen:            c.String("listen"),
				key:               c.String("key"),
				crypt:             c.String("crypt"),
				tcp_readDeadline:  c.Duration("tcp_readDeadline"),
				tcp_writeDeadline: c.Duration("tcp_writeDeadline"),
				udp_readDeadline:  c.Duration("udp_readDeadline"),
				udp_writeDeadline: c.Duration("udp_writeDeadline"),
				tcp_sockbuf_w:     c.Int("tcp_sockbuf_w"),
				tcp_sockbuf_r:     c.Int("tcp_sockbuf_r"),
				udp_sockbuf_w:     c.Int("udp_sockbuf_w"),
				udp_sockbuf_r:     c.Int("udp_sockbuf_r"),
				queuelen:          c.Int("queuelen"),
				dscp:              c.Int("dscp"),
				sndwnd:            c.Int("udp-sndwnd"),
				rcvwnd:            c.Int("udp-rcvwnd"),
				mtu:               c.Int("udp-mtu"),
				dataShard:         c.Int("dataShard"),
				parityShards:      c.Int("parityShards"),
				nodelay:           c.Int("nodelay"),
				interval:          c.Int("interval"),
				resend:            c.Int("resend"),
				nc:                c.Int("nc"),
			}
			// start udp server...
			this.connect(config, this.ctx, &this.sw)
			this.sw.Add(1)
			go this.loopconnect(config, this.ctx, &this.sw)
			go this.loopOffline(this.ctx, &this.sw)
			go this.loopSignalCheck(this.ctx, &this.sw)
			go func() {
				akLog.FmtPrintln("[client] run http server, host: ", this.ppAddr)
				http.ListenAndServe(this.ppAddr, nil)
			}()
			this.sw.Wait()
			return nil
		},
	}

	app.Run(os.Args)
}

func (this *KcpClient) connect(c *KcpSvrConfig, ctx context.Context, sw *sync.WaitGroup) {
	var pass = pbkdf2.Key([]byte(c.key), []byte(c.crypt), 4096, 32, sha1.New)
	block, _ := kcp.NewAESBlockCrypt(pass)
	akLog.FmtPrintln("client addr: ", this.Addr)
	conn, err := kcp.DialWithOptions(c.listen, block, 10, 3)
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.SetStreamMode(true)
	conn.SetWindowSize(c.sndwnd, c.rcvwnd)
	conn.SetReadBuffer(c.udp_sockbuf_r)
	conn.SetWriteBuffer(c.udp_sockbuf_w)
	conn.SetNoDelay(c.nodelay, c.interval, c.resend, c.nodelay)
	conn.SetMtu(c.mtu)
	conn.SetACKNoDelay(false)
	conn.SetDeadline(time.Now().Add(time.Minute))
	this.sesson = NewKcpClientSession(conn, this.offCh, this.exCollection)
	this.sesson.Handler()
	this.sendRegisterMsg()
	if this.exCollection != nil {
		this.exCollection.SetCenterClient(this)
	}
}

func (this *KcpClient) sendRegisterMsg() {
	akLog.FmtPrintf("after dial, send point: %v register message to server.", this.svrType)
	req := &MSG_Server.CS_ServerRegister_Req{}
	req.ServerType = int32(this.svrType)
	req.Msgs = GetAllMessageIDs()
	req.Version = this.versionNo
	akLog.FmtPrintln("register context: ", req.Msgs)
	data, err := this.pack.PackClientMsg(uint16(MSG_MainModule.MAINMSG_SERVER), uint16(MSG_Server.SUBMSG_CS_ServerRegister), req)
	if err != nil {
		akLog.Error(err)
		return
	}
	this.Send(data)
}

func (this *KcpClient) Send(data []byte) {
	if !this.sesson.Alive() {
		return
	}
	this.sesson.SetSendCache(data)
}

func (this *KcpClient) loopconnect(c *KcpSvrConfig, ctx context.Context, sw *sync.WaitGroup) {
	defer func() {
		this.exit()
		sw.Done()
	}()

	tick := time.NewTicker(time.Duration(5) * time.Second)
	for {
		select {
		case <-tick.C:
			if this.sesson == nil || !this.sesson.Alive() {
				this.connect(c, ctx, sw)
			}
		case <-ctx.Done():
			return
		}
	}
}

func (this *KcpClient) loopOffline(ctx context.Context, sw *sync.WaitGroup) {
	defer func() {
		this.exit()
		sw.Done()
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case offsession := <-this.offCh:
			offsession.Offline()
		}
	}
}

func (this *KcpClient) exit() {
	this.cancel()
}

func (this *KcpClient) loopSignalCheck(ctx context.Context, sw *sync.WaitGroup) {
	defer func() {
		sw.Done()
		this.exit()
	}()

	chsignal := make(chan os.Signal, 1)
	signal.Notify(chsignal, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for {
		select {
		case <-ctx.Done():
			return
		case s := <-chsignal:
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				akLog.FmtPrintln("signal exit:", s)
				return
			default:
				akLog.FmtPrintln("other signal:", s)
			}
		}
	}
}
