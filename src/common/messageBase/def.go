package messageBase

/*
	网关客户端消息组合：
	|主消息ID|内容长度|内容数据|
		4  		4		4	 len(内容长度)

*/
const (
	MSG_PACK_HEADID_SIZE = 4
	MSG_PACK_DATA_SIZE   = 8
	MSG_PACK_NODATA_SIZE = (MSG_PACK_HEADID_SIZE + MSG_PACK_DATA_SIZE)
)

type NetType uint16

const (
	Net_TCP NetType = 1
	Net_WS  NetType = 2
)
