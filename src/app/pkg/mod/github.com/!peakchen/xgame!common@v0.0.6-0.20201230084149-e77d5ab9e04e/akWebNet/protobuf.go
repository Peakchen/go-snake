package akWebNet

import (
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"

	"github.com/Peakchen/xgameCommon/akNet"
	"github.com/golang/protobuf/proto"
)

//TPackProto pack proto struct to data or unpack...
type TPackProto struct {
	mainID   uint16
	subID    uint16
	dataSize uint32
	data     []byte
}

func (this *TPackProto) UnPack(msg []byte) (err error) {
	if len(msg) <= MSG_PACK_HEADID_SIZE {
		err = errors.New("msg head id size invalid.")
		return
	}
	var pos int
	this.mainID = binary.LittleEndian.Uint16(msg[pos:])
	pos += 2
	this.subID = binary.LittleEndian.Uint16(msg[pos:])
	pos += 2

	if len(msg[pos:]) <= MSG_PACK_DATA_SIZE {
		err = errors.New("msg data size field invalid.")
		return
	}

	this.dataSize = binary.LittleEndian.Uint32(msg[pos:])
	pos += 4
	if len(msg[pos:]) < int(this.dataSize) {
		err = errors.New("msg data content size invalid.")
		return
	}
	this.data = msg[pos : pos+int(this.dataSize)]
	return
}

func (this *TPackProto) Pack(out []byte) {
	var pos int
	binary.LittleEndian.PutUint16(out[pos:], this.mainID)
	pos += 2

	binary.LittleEndian.PutUint16(out[pos:], this.subID)
	pos += 2

	binary.LittleEndian.PutUint32(out[pos:], this.dataSize)
	pos += 4
	copy(out[pos:], this.data)
}

type ProtoBufMsgOp struct {
	pp *TPackProto
}

func (this *ProtoBufMsgOp) Init() {
	if this.pp == nil {
		this.pp = &TPackProto{}
	}
}

//protoBufPack for protobuf pack data
func (this *ProtoBufMsgOp) Pack(mainid, subid uint16, msg proto.Message) (out []byte, err error) {
	data, err := proto.Marshal(msg)
	if err != nil {
		err = fmt.Errorf("proto marshal fail, mainid: %v, subid: %v, err: %v.", mainid, subid, err)
		return
	}

	this.pp.mainID = mainid
	this.pp.subID = subid
	this.pp.dataSize = uint32(len(data))
	this.pp.data = data

	out = make([]byte, len(data)+MSG_PACK_NODATA_SIZE)
	this.pp.Pack(out)
	return
}

func (this *ProtoBufMsgOp) UnPack() (msg proto.Message, cb reflect.Value, err error) {
	msg, cb, err = akNet.UnPackMsg(this.pp.mainID, this.pp.subID, this.pp.data)
	return
}

func (this *ProtoBufMsgOp) GetMsgInfo(data []byte) (info *BaseMsgInfo, err error) {
	err = this.pp.UnPack(data)
	if err != nil {
		return
	}
	info = &BaseMsgInfo{
		MainID: this.pp.mainID,
		SubID:  this.pp.subID,
	}
	return
}
