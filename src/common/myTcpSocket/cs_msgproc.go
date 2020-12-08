package myTcpSocket

import (
	"encoding/binary"
	"io"
	"net"
	"syscall"

	"github.com/Peakchen/xgameCommon/akLog"
)

func ClientMsgProc(conn *net.TCPConn) bool {
	buff := make([]byte, 8)
	rdn1, err := io.ReadFull(conn, buff)
	if err != nil || rdn1 < 8 {
		switch {
		case (err == syscall.EAGAIN || err == syscall.EWOULDBLOCK):
		case rdn1 > 0 && rdn1 < 8:
			akLog.Error("unexcept error: ", err)
			return false
		case rdn1 == 0 && err == nil:
			akLog.Error("net EOF")
		}

		return true
	}

	size := binary.LittleEndian.Uint32(buff[4:8])
	if size > 1024 {
		akLog.Error("data size invalid, value: ", size)
		return false
	}

	data := make([]byte, 8+size)
	rdn2, err := io.ReadFull(conn, data[8:])
	if rdn2 < int(size) || err != nil {
		akLog.Error("read real data fail, read data: ", rdn2, err)
		return false
	}

	copy(data[:8], buff[:])

	ClientProc(data)
	return true
}
