package webNet

import (
	"time"
)

const (
	// 允许等待的写入时间
	writeWait = 10 * time.Second

	//允许时间从对等方读取下一个pong消息。
	pongWait = 60 * time.Second

	//在此期间将ping发送给同级。 必须小于pongWait。
	pingPeriod = (pongWait * 5) / 10

	//允许来自对等方的最大信息大小。
	maxMessageSize = 512

	//发送通信管道最大信息量
	maxWriteMsgSize = 1000
)

const (
	MsgSendType_P2P       = 1 //点对点 直接发送
	MsgSendType_BroadCast = 2 //广播
)
