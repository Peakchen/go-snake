package tcpNet

import (
	"encoding/binary"
	"go-snake/common/messageBase"
	"io"
	"net"
	"syscall"

	"github.com/Peakchen/xgameCommon/akLog"
)

func ClientMsgProc(sid string, conn *net.TCPConn, fn func(string, []byte)) bool {
	buff := make([]byte, messageBase.CS_MSG_PACK_DATA_SIZE)
	rdn1, err := io.ReadFull(conn, buff)
	if err != nil || rdn1 < messageBase.CS_MSG_PACK_DATA_SIZE {
		switch {
		case (err == syscall.EAGAIN || err == syscall.EWOULDBLOCK):
		case rdn1 > 0 && rdn1 < messageBase.CS_MSG_PACK_DATA_SIZE:
			akLog.Error("unexcept error: ", err)
			return false
		case rdn1 == 0 && err == nil:
			akLog.Info("net EOF")
		}

		return true
	}

	size := binary.LittleEndian.Uint32(buff[messageBase.CS_MSG_PACK_DATA_SIZE:messageBase.CS_MSG_PACK_NODATA_SIZE])
	if size > 1024 {
		akLog.Error("data size invalid, value: ", size)
		return false
	}

	data := make([]byte, messageBase.CS_MSG_PACK_DATA_SIZE+size)
	rdn2, err := io.ReadFull(conn, data[messageBase.CS_MSG_PACK_DATA_SIZE:])
	if rdn2 < int(size) || err != nil {
		akLog.Error("read real data fail, read data: ", rdn2, err)
		return false
	}

	copy(data[:messageBase.CS_MSG_PACK_DATA_SIZE], buff[:])

	fn(sid, data)
	return true
}
