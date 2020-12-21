package tcpNet

import (
	"encoding/binary"
	"go-snake/common/messageBase"
	"io"
	"net"
	"syscall"

	"github.com/Peakchen/xgameCommon/akLog"
)

func ServerMsgProc(sid string, conn *net.TCPConn, fn func(string, []byte)) bool {
	buff := make([]byte, messageBase.SS_MSG_PACK_NODATA_SIZE)
	rdn1, err := io.ReadFull(conn, buff)
	if err != nil || rdn1 < messageBase.SS_MSG_PACK_NODATA_SIZE {
		switch {
		case (err == syscall.EAGAIN || err == syscall.EWOULDBLOCK):
		case rdn1 > 0 && rdn1 < messageBase.SS_MSG_PACK_NODATA_SIZE:
			akLog.Error("unexcept error: ", err)
			return false
		case rdn1 == 0 && err == nil:
			akLog.Error("net EOF")
		}

		return true
	}

	size := binary.LittleEndian.Uint32(buff[messageBase.SS_MSG_PACK_DATA_SIZE:messageBase.SS_MSG_PACK_NODATA_SIZE])
	if size > 1024 {
		akLog.Error("data size invalid, value: ", size)
		return false
	}

	data := make([]byte, messageBase.SS_MSG_PACK_NODATA_SIZE+size)
	rdn2, err := io.ReadFull(conn, data[messageBase.SS_MSG_PACK_NODATA_SIZE:])
	if rdn2 < int(size) || err != nil {
		akLog.Error("read real data fail, read data: ", rdn2, err)
		return false
	}

	copy(data[:messageBase.SS_MSG_PACK_NODATA_SIZE], buff[:])

	fn(sid, data)
	return true
}
