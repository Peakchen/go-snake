package akNet

import (
	"net"
	"strconv"
	"strings"
)

/*
	func: EncodeCmd
	purpose: Encode message mainid and subid to cmd.
*/
func EncodeCmd(mainID, subID uint16) uint32 {
	return (uint32(mainID) << 16) | uint32(subID)
}

/*
	func: DecodeCmd
	purpose: DecodeCmd message cmd to mainid and subid.
*/
func DecodeCmd(cmd uint32) (uint16, uint16) {
	return uint16(cmd >> 16), uint16(cmd)
}

//ip地址string转int64
func IpToInt64(ip net.IP) int64 {
	bits := strings.Split(ip.String(), ".")
	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}

//ip地址int64转string
func Int64Toip(ipn int64) net.IP {
	var bytes [4]byte
	bytes[0] = byte(ipn & 0xFF)
	bytes[1] = byte((ipn >> 8) & 0xFF)
	bytes[2] = byte((ipn >> 16) & 0xFF)
	bytes[3] = byte((ipn >> 24) & 0xFF)

	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
}
