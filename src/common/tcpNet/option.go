package tcpNet

import (
	"go-snake/akmessage"
	"go-snake/common/messageBase"
	"net"
)

type MsgCallBack func(string, *net.TCPConn, func(string, []byte)) bool

type OptionFn func(opts *ExtFnsOption)

type ExtFnsOption struct {
	SS_HeartBeat func(string) []byte
	CS_HeartBeat func(string) []byte
	Handler      MsgCallBack
}

func SortOptions(fns ...OptionFn) *ExtFnsOption {
	opts := new(ExtFnsOption)
	for _, optFn := range fns {
		optFn(opts)
	}
	return opts
}

func WithSSHeartBeat(fn func(string) []byte) OptionFn {
	return func(opts *ExtFnsOption) {
		opts.CS_HeartBeat = fn
	}
}

func WithMessageHandler(fn MsgCallBack) OptionFn {
	return func(opts *ExtFnsOption) {
		opts.Handler = fn
	}
}

func SS_HeatBeatMsg(sid string) []byte {
	return messageBase.SSPackMsg(sid, 0, akmessage.MSG_SS_HEARTBEAT_REQ, &akmessage.SS_HeartBeat_Req{})
}

func CS_HeatBeatMsg(sid string) []byte {
	return nil
}
