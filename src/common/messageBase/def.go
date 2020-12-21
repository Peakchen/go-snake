package messageBase

/*
	网关客户端消息组合：
	|主消息ID|内容长度|内容数据|
		4  		4		4	 len(内容长度)

*/
const (
	CS_MSG_PACK_HEADID_SIZE = 4
	CS_MSG_PACK_DATA_SIZE   = 8
	CS_MSG_PACK_NODATA_SIZE = (CS_MSG_PACK_HEADID_SIZE + CS_MSG_PACK_DATA_SIZE)
)

/*
	网关服务器消息组合：
	|主消息ID|会话ID|uid|内容长度|内容数据|
		4  		36	 8	  4	 len(内容长度)

*/
const (
	SS_MSG_PACK_HEADID_SIZE  = 4
	SS_MSG_PACK_DATA_SIZE    = 48
	SS_MSG_PACK_NODATA_SIZE  = (SS_MSG_PACK_HEADID_SIZE + SS_MSG_PACK_DATA_SIZE)
	SS_MSG_PACK_SESSION_SIZE = 36 // format: 5e144ec0-d7d4-4190-b38b-fcaf09b18246
	SS_MSG_PACK_UID_SIZE     = 8  // user id
)

type NetType uint16

const (
	Net_TCP NetType = 1
	Net_WS  NetType = 2
)
