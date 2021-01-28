package akNet

// add by stefan

import (
	"github.com/Peakchen/xgameCommon/akLog"
	"github.com/Peakchen/xgameCommon/stacktrace"
	"github.com/Peakchen/xgameCommon/utils"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"reflect"
)

/*
	model: ServerProtocol
	server to server, message
*/
type ServerProtocol struct {
	routepoint uint16
	mainid     uint16
	subid      uint16
	length     uint32
	data       []byte //消息体
	srcdata    []byte //源消息（未解包）
	identify   string
	remoteAddr string
	postType   uint16
}

func (this *ServerProtocol) PackAction(Output []byte) (err error) {
	var pos int
	// binary.LittleEndian.PutUint16(Output[pos:], this.routepoint)
	// pos += 2

	binary.LittleEndian.PutUint16(Output[pos:], this.mainid)
	pos += 2

	binary.LittleEndian.PutUint16(Output[pos:], this.subid)
	pos += 2

	binary.LittleEndian.PutUint16(Output[pos:], this.postType)
	pos += 2

	if len(this.identify) == 0 {
		err = akLog.RetError("[server] identify invalid, mainid: %v, subid: %v.", this.mainid, this.subid)
		return
	}

	if len(this.remoteAddr) == 0 {
		err = akLog.RetError("[server] remoteAddr invalid, mainid: %v, subid: %v.", this.mainid, this.subid)
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

	copy(Output[pos:], this.data)
	//akLog.FmtPrintln("server PackAction-> data len: ", this.length, len(Output))
	return
}

func (this *ServerProtocol) PackAction4Client(Output []byte) (err error) {
	var pos int
	// binary.LittleEndian.PutUint16(Output[pos:], this.routepoint)
	// pos += 2

	binary.LittleEndian.PutUint16(Output[pos:], this.mainid)
	pos += 2

	binary.LittleEndian.PutUint16(Output[pos:], this.subid)
	pos += 2

	binary.LittleEndian.PutUint16(Output[pos:], this.postType)
	pos += 2

	binary.LittleEndian.PutUint32(Output[pos:], this.length)
	pos += 4
	//akLog.FmtPrintln("PackAction4Client PackAction-> data len: ", this.length)
	copy(Output[pos:], this.data)
	return
}

func (this *ServerProtocol) UnPackData() (msg proto.Message, cb reflect.Value, err error, exist bool) {
	err = nil
	mt, finded := GetMessageInfo(this.mainid, this.subid)
	if !finded {
		err = fmt.Errorf("[server] can not regist message, mainid: %v, subid: %v.", this.mainid, this.subid)
		return
	}

	exist = true
	dst := reflect.New(mt.paramTypes[1].Elem()).Interface()
	err = proto.Unmarshal(this.data, dst.(proto.Message))
	if err != nil {
		err = fmt.Errorf("[server] unmarshal message fail, mainid: %v, subid: %v, err: %v.", this.mainid, this.subid, err)
		return
	}
	msg = dst.(proto.Message)
	cb = mt.proc
	return
}

func (this *ServerProtocol) PackData(msg proto.Message) (data []byte, err error) {
	data, err = proto.Marshal(msg)
	return
}

func (this *ServerProtocol) GetRouteID() (route uint16) {
	return this.routepoint
}

func (this *ServerProtocol) GetMessageID() (mainID uint16, subID uint16) {
	return this.mainid, this.subid
}

func (this *ServerProtocol) Clean() {
	this.length = 0
	this.data = make([]byte, maxMessageSize)
	this.mainid = 0
	this.subid = 0
}

func (this *ServerProtocol) SetCmd(mainid, subid uint16, data []byte) {
	this.mainid = mainid
	this.subid = subid
	this.data = data
	this.length = uint32(len(data))

	//akLog.FmtPrintf("[server] SetCmd mainid: %v, subid: %v, data len: %v.", mainid, subid, this.length)
}

func (this *ServerProtocol) PackInnerMsg(mainid, subid uint16, msg proto.Message) (out []byte, err error) {
	data, err := proto.Marshal(msg)
	if err != nil {
		err = akLog.RetError("server proto marshal fail, data: %v.", err)
		return
	}

	this.SetCmd(mainid, subid, data)
	out = make([]byte, len(data)+EnMessage_SvrNoDataLen)
	err = this.PackAction(out)
	return
}

func (this *ServerProtocol) PackClientMsg(mainid, subid uint16, msg proto.Message) (out []byte, err error) {
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

func (this *ServerProtocol) GetSrcMsg() (data []byte) {
	return this.srcdata
}

func (this *ServerProtocol) SetIdentify(identify string) {
	this.identify = identify
}

func (this *ServerProtocol) GetIdentify() string {
	return this.identify
}

func (this *ServerProtocol) GetDataLen() (datalen uint32) {
	return this.length
}

func (this *ServerProtocol) SetRemoteAddr(addr string) {
	this.remoteAddr = addr
}

func (this *ServerProtocol) GetRemoteAddr() (addr string) {
	addr = this.remoteAddr
	return
}

func (this *ServerProtocol) SetPostType(pt uint16) {
	this.postType = pt
}

func (this *ServerProtocol) GetPostType() (pt uint16) {
	return this.postType
}

func (this *ServerProtocol) UnPackMsg4Client(InData []byte) (pos int, err error) {
	defer stacktrace.Catchcrash()

	// this.routepoint = binary.LittleEndian.Uint16(InData[pos:])
	// pos += 2

	this.mainid = binary.LittleEndian.Uint16(InData[pos:])
	pos += 2

	this.subid = binary.LittleEndian.Uint16(InData[pos:])
	pos += 2

	this.postType = binary.LittleEndian.Uint16(InData[pos:])
	pos += 2

	this.length = binary.LittleEndian.Uint32(InData[pos:])
	pos += 4

	datalen := utils.SliceBytesLength(InData)
	//akLog.FmtPrintln("server UnPackMsg4Client-> len: ", this.length, datalen)
	if datalen < int(pos+int(this.length)) {
		err = fmt.Errorf("server  mainid: %v, subid: %v; err: InData len: %v, pos: %v, data len: %v.", this.mainid, this.subid, len(InData), pos, this.length)
		return
	}

	this.data = InData[pos : pos+int(this.length)]
	this.srcdata = InData[:]
	//akLog.FmtPrintf("message head: mainid: %v, subid: %v, srcdata len: %v.", this.mainid, this.subid, len(this.srcdata))
	return pos, nil
}

func (this *ServerProtocol) UnPackMsg4Svr(InData []byte) (pos int, err error) {
	defer stacktrace.Catchcrash()

	// this.routepoint = binary.LittleEndian.Uint16(InData[pos:])
	// pos += 2

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
		if datalen < (int(identifylength) + pos) {
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

	//akLog.FmtPrintln("server UnPackMsg4Svr-> len: ", this.length)
	this.length = binary.LittleEndian.Uint32(InData[pos:])
	pos += 4

	if datalen < int(pos+int(this.length)) {
		err = fmt.Errorf("server mainid: %v, subid: %v; err: InData len: %v, pos: %v, data len: %v.", this.mainid, this.subid, len(InData), pos, this.length)
		return
	}

	//akLog.FmtPrintf("message head: mainid: %v, subid: %v.", this.mainid, this.subid)
	this.data = InData[pos : pos+int(this.length)]
	this.srcdata = InData[:]
	return pos, nil
}
