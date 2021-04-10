package Kcpnet

import (
	"reflect"
	"time"

	"github.com/Peakchen/xgameCommon/define"
	"github.com/golang/protobuf/proto"
)

/*

- 协议格式，小端字节序

|主命令号 | 次命令号 | 发送类型 |包长度 | 包体
---------|-------- | -------- | -------- | ----
2字节    |2字节     |  2字节    | 4字节  | 包体
*/
const (
	EnMessage_MainIDPackLen = 2  //主命令
	EnMessage_SubIDPackLen  = 2  //次命令
	EnMessage_PostType      = 2  //发送类型
	EnMessage_DataPackLen   = 6  //真实数据长度 (主命令+次命令+发送类型)
	EnMessage_NoDataLen     = 10 //非data数据长度(包体之前的)->(主命令+次命令+发送类型+datalen)

	//2020.04.15 Identify 会考虑换成 userid 减小包长度
	EnMessage_SvrDataPackLen  = 50 //真实数据长度 (主命令+次命令+发送类型+ Identify长度 + Identify内容+RemoteAddr 长度+RemoteAddr 内容)
	EnMessage_SvrNoDataLen    = 54 //非data数据长度(包体之前的)->(主命令+次命令+发送类型+ Identify长度 + Identify内容+datalen+RemoteAddr 长度+RemoteAddr 内容)
	EnMessage_IdentifyEixst   = 1  //Identify 长度
	EnMessage_IdentifyLen     = 21 //Identify 内容
	EnMessage_RemoteAddrEixst = 1  //RemoteAddr 长度
	EnMessage_RemoteAddrLen   = 21 //RemoteAddr 内容
)

type IMessagePack interface {
	PackAction(Output []byte) (err error)
	PackAction4Client(Output []byte) (err error)
	PackData(msg proto.Message) (data []byte, err error)
	UnPackMsg4Client(InData []byte) (pos int, err error)
	UnPackData() (msg proto.Message, cb reflect.Value, err error, exist bool)
	GetRouteID() (route uint16)
	GetMessageID() (mainID uint16, subID uint16)
	Clean()
	SetCmd(mainid, subid uint16, data []byte)
	PackInnerMsg(mainid, subid uint16, msg proto.Message) (out []byte, err error)
	PackClientMsg(mainid, subid uint16, msg proto.Message) (out []byte, err error)
	GetSrcMsg() (data []byte)
	SetIdentify(identify string)
	GetIdentify() string
	UnPackMsg4Svr(InData []byte) (pos int, err error)
	GetDataLen() (datalen uint32)
	SetRemoteAddr(addr string)
	GetRemoteAddr() (addr string)
	SetPostType(pt uint16)
	GetPostType() (pt uint16)
}

type TSession interface {
	GetRemoteAddr() string
	GetRegPoint() (RegPoint define.ERouteId)
	GetIdentify() string
	SetSendCache(data []byte)
	Push(RegPoint define.ERouteId)
	SetIdentify(StrIdentify string)
	SendInnerSvrMsg(mainid, subid uint16, msg proto.Message) (succ bool, err error)
	SendSvrClientMsg(mainid, subid uint16, msg proto.Message) (succ bool, err error)
	SendInnerClientMsg(mainid, subid uint16, msg proto.Message) (succ bool, err error)
	SendInnerBroadcastMsg(mainid, subid uint16, msg proto.Message) (succ bool, err error)
	WriteMessage(data []byte) (succ bool)
	Alive() bool
	GetPack() (obj IMessagePack)
	IsUser() bool
	RefreshHeartBeat(mainid, subid uint16) bool
	GetModuleName() string
	GetExternalCollection() *ExternalCollection
	GetVer() int32
	SetVer(ver int32)
	GetSvrType() (t define.ERouteId)
}

const (
	cstKeepLiveHeartBeatSec     = 180 //180 3min
	cstCheckHeartBeatMonitorSec = cstKeepLiveHeartBeatSec / 2
	cstSvrDisconnectionSec      = 3 * cstKeepLiveHeartBeatSec //s
	cstClientDisconnectionSec   = 6 * cstKeepLiveHeartBeatSec //s
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 3 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 4096
	//offline session
	maxOfflineSize = 1024
	//
	udpBuffSize = 4 * 1024 * 1024
	tcpBuffSize = 1024
	queueSize   = 1000
	dscp        = 46
)

// message Post type
const (
	MsgPostType_Single    = uint16(1)
	MsgPostType_Broadcast = uint16(2)
)
