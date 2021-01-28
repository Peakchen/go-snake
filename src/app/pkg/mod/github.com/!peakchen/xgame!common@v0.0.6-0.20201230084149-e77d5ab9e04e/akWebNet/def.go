package akWebNet

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

/*
	外网关客户端消息组合：
	|主消息ID|次消息ID|内容长度|内容数据|
		2  		2		4	 len(内容长度)

	网关内消息组合：
	|目标服务器|主消息ID|次消息ID|内容长度|内容数据|
		2			2  		2		4	 len(内容长度)

*/
const (
	MSG_PACK_HEADID_SIZE = 6
	MSG_PACK_DATA_SIZE   = 4
	MSG_PACK_NODATA_SIZE = (MSG_PACK_HEADID_SIZE + MSG_PACK_DATA_SIZE)
)

//打包方式
type PACK_TYPE uint8

const (
	PACK_PROTO = PACK_TYPE(1)
)

//服务角色类型
type ACTOR_TYPE uint8

const (
	ACTOR_FRONT = ACTOR_TYPE(1)
	ACTOR_BACK  = ACTOR_TYPE(2)
)

const cstKeepLiveHeartBeatSec = (60 * 9) / 10
