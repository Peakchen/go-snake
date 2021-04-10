package Kcpnet

// add by stefan

import (
	"encoding/binary"
	"fmt"
	"reflect"

	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/stacktrace"
	"github.com/Peakchen/xgameCommon/utils"
	"github.com/golang/protobuf/proto"
)

/*
	model: KcpClientProtocol
	Client to Server, message
*/
type KcpClientProtocol struct {
	routepoint uint16
	mainid     uint16
	subid      uint16
	length     uint32
	data       []byte
	srcdata    []byte
	identify   string
	remoteAddr string
	postType   uint16
}

func (this *KcpClientProtocol) PackAction(Output []byte) (err error) {
	var pos int

	binary.LittleEndian.PutUint16(Output[pos:], this.mainid)
	pos += 2

	binary.LittleEndian.PutUint16(Output[pos:], this.subid)
	pos += 2

	binary.LittleEndian.PutUint16(Output[pos:], this.postType)
	pos += 2

	if len(this.identify) == 0 {
		err = akLog.RetError("[client] identify invalid, mainid: %v, subid: %v.", this.mainid, this.subid)
		return
	}

	if len(this.remoteAddr) == 0 {
		err = akLog.RetError("[client] remoteAddr invalid, mainid: %v, subid: %v.", this.mainid, this.subid)
		return
	}

	Output[pos] = uint8(EnMessage_IdentifyLen)
	pos++

	copy(Output[pos:], []byte(this.identify))
	pos += EnMessage_IdentifyLen

	Output[pos] = uint8(EnMessage_RemoteAddrLen)
	pos++

	copy(Output[pos:], []byte(this.remoteAddr))
	pos += EnMessage_RemoteAddrLen

	binary.LittleEndian.PutUint32(Output[pos:], this.length)
	pos += 4
	//akLog.FmtPrintln("client PackAction-> data len: ", this.length)
	copy(Output[pos:], this.data)
	err = nil
	return
}

func (this *KcpClientProtocol) PackAction4Client(Output []byte) (err error) {
	var pos int

	binary.LittleEndian.PutUint16(Output[pos:], this.mainid)
	pos += 2

	binary.LittleEndian.PutUint16(Output[pos:], this.subid)
	pos += 2

	binary.LittleEndian.PutUint16(Output[pos:], this.postType)
	pos += 2

	binary.LittleEndian.PutUint32(Output[pos:], this.length)
	pos += 4
	//akLog.FmtPrintln("client PackAction4Client-> data len: ", this.length)
	copy(Output[pos:], this.data)
	err = nil
	return
}

func (this *KcpClientProtocol) UnPackData() (msg proto.Message, cb reflect.Value, err error, exist bool) {
	err = nil
	mt, finded := GetMessageInfo(this.mainid, this.subid)
	if !finded {
		err = fmt.Errorf("[client] can not regist message, mainid: %v, subid: %v.", this.mainid, this.subid)
		return
	}

	exist = true
	dst := reflect.New(mt.paramTypes[1].Elem()).Interface()
	err = proto.Unmarshal(this.data, dst.(proto.Message))
	if err != nil {
		err = fmt.Errorf("[client] unmarshal message fail, mainid: %v, subid: %v, err: %v.", this.mainid, this.subid, err)
		return
	}
	msg = dst.(proto.Message)
	cb = mt.proc
	return
}

func (this *KcpClientProtocol) PackData(msg proto.Message) (data []byte, err error) {
	data, err = proto.Marshal(msg)
	return
}

func (this *KcpClientProtocol) GetRouteID() (route uint16) {
	return this.routepoint
}

func (this *KcpClientProtocol) GetMessageID() (mainID uint16, subID uint16) {
	return this.mainid, this.subid
}

func (this *KcpClientProtocol) SetCmd(mainid, subid uint16, data []byte) {
	this.mainid = mainid
	this.subid = subid
	this.data = data
	this.length = uint32(len(data))
	//akLog.FmtPrintf("[client] mainid: %v, subid: %v, data len: %v.", mainid, subid, this.length)
}

func (this *KcpClientProtocol) Clean() {
	this.length = 0
	this.data = make([]byte, maxMessageSize)
	this.mainid = 0
	this.subid = 0
}

func (this *KcpClientProtocol) PackInnerMsg(mainid, subid uint16, msg proto.Message) (out []byte, err error) {
	data, err := proto.Marshal(msg)
	if err != nil {
		err = akLog.RetError("client proto marshal fail, data: %v.", err)
		return
	}

	this.SetCmd(mainid, subid, data)
	out = make([]byte, len(data)+EnMessage_SvrNoDataLen)
	err = this.PackAction(out)
	return
}

func (this *KcpClientProtocol) PackClientMsg(mainid, subid uint16, msg proto.Message) (out []byte, err error) {
	data, err := proto.Marshal(msg)
	if err != nil {
		err = akLog.RetError("client for reg proto marshal fail, data: %v.", err)
		return
	}

	this.SetCmd(mainid, subid, data)
	out = make([]byte, len(data)+EnMessage_NoDataLen)
	err = this.PackAction4Client(out)
	return
}

func (this *KcpClientProtocol) GetSrcMsg() (data []byte) {
	return this.srcdata
}

func (this *KcpClientProtocol) SetIdentify(identify string) {
	this.identify = identify
}

func (this *KcpClientProtocol) GetIdentify() string {
	return this.identify
}

func (this *KcpClientProtocol) GetDataLen() (datalen uint32) {
	return this.length
}

func (this *KcpClientProtocol) SetRemoteAddr(addr string) {
	this.remoteAddr = addr
}

func (this *KcpClientProtocol) GetRemoteAddr() (addr string) {
	addr = this.remoteAddr
	return
}

func (this *KcpClientProtocol) SetPostType(pt uint16) {
	this.postType = pt
}

func (this *KcpClientProtocol) GetPostType() (pt uint16) {
	return this.postType
}

func (this *KcpClientProtocol) UnPackMsg4Client(InData []byte) (pos int, err error) {
	defer stacktrace.Catchcrash()

	this.mainid = binary.LittleEndian.Uint16(InData[pos:])
	pos += 2

	this.subid = binary.LittleEndian.Uint16(InData[pos:])
	pos += 2

	this.postType = binary.LittleEndian.Uint16(InData[pos:])
	pos += 2

	this.length = binary.LittleEndian.Uint32(InData[pos:])
	pos += 4

	if utils.SliceBytesLength(InData) < int(pos+int(this.length)) {
		err = fmt.Errorf("client mainid: %v, subid: %v; err: InData len: %v, pos: %v, data len: %v.", this.mainid, this.subid, len(InData), pos, this.length)
		return
	}

	this.data = InData[pos : pos+int(this.length)]
	this.srcdata = InData
	return pos, nil
}

func (this *KcpClientProtocol) UnPackMsg4Svr(InData []byte) (pos int, err error) {
	defer stacktrace.Catchcrash()

	this.mainid = binary.LittleEndian.Uint16(InData[pos:])
	pos += 2

	this.subid = binary.LittleEndian.Uint16(InData[pos:])
	pos += 2

	this.postType = binary.LittleEndian.Uint16(InData[pos:])
	pos += 2

	datalen := utils.SliceBytesLength(InData)
	identifylength := uint8(InData[pos])
	pos++
	if identifylength > 1 {
		if datalen < (int(identifylength) + pos + int(this.length)) {
			return
		}

		this.identify = string(InData[pos : pos+int(identifylength)])
		pos += int(identifylength)
	}

	remoteAddrlength := uint8(InData[pos])
	pos++
	if remoteAddrlength > 1 {
		if datalen < (int(remoteAddrlength) + pos) {
			return
		}

		this.remoteAddr = string(InData[pos : pos+int(remoteAddrlength)])
		pos += int(remoteAddrlength)
	}

	this.length = binary.LittleEndian.Uint32(InData[pos:])
	pos += 4
	if datalen < int(pos+int(this.length)) {
		err = fmt.Errorf("client  mainid: %v, subid: %v; err: InData len: %v, pos: %v, data len: %v.", this.mainid, this.subid, len(InData), pos, this.length)
		return
	}

	this.data = InData[pos : pos+int(this.length)]
	this.srcdata = InData
	return pos, nil
}
