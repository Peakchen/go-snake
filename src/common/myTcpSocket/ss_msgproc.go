package myTcpSocket

import (
	"encoding/binary"
	"io"
	"net"
	"syscall"

	"github.com/Peakchen/xgameCommon/akLog"
)

func ServerMsgProc(conn *net.TCPConn) bool {
	buff := make([]byte, 10)
	rdn1, err := io.ReadFull(conn, buff)
	if err != nil || rdn1 < 10 {
		switch {
		case (err == syscall.EAGAIN || err == syscall.EWOULDBLOCK):
		case rdn1 > 0 && rdn1 < 10:
			akLog.Error("unexcept error: ", err)
			return false
		case rdn1 == 0 && err == nil:
			akLog.Error("net EOF")
		}

		return true
	}

	size := binary.LittleEndian.Uint32(buff[6:10])
	if size > 1024 {
		akLog.Error("data size invalid, value: ", size)
		return false
	}

	data := make([]byte, 10+size)
	rdn2, err := io.ReadFull(conn, data[10:])
	if rdn2 < int(size) || err != nil {
		akLog.Error("read real data fail, read data: ", rdn2, err)
		return false
	}

	copy(data[:10], buff[:])

	ServerProc(data)
	return true
}
