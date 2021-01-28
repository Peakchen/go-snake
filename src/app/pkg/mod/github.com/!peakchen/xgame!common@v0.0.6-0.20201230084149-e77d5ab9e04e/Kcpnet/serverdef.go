package Kcpnet

import "time"

type KcpSvrConfig struct {
	listen                        string
	key                           string
	crypt                         string
	tcp_readDeadline              time.Duration
	tcp_writeDeadline             time.Duration
	udp_readDeadline              time.Duration
	udp_writeDeadline             time.Duration
	tcp_sockbuf_w                 int
	tcp_sockbuf_r                 int
	udp_sockbuf_w                 int
	udp_sockbuf_r                 int
	queuelen                      int
	dscp                          int
	sndwnd                        int
	rcvwnd                        int
	mtu                           int
	dataShard                     int
	parityShards                  int
	nodelay, interval, resend, nc int
}
