package messageBase

import (
	"encoding/binary"
	"errors"
	"fmt"
)

//TCSPackProto pack proto struct to data or unpack...
type TCSPackProto struct {
	mainID   uint32
	dataSize uint32
	data     []byte
}

func (this *TCSPackProto) Init(id uint32, src []byte) {
	this.mainID = id
	this.dataSize = uint32(len(src))
	this.data = src
}

func (this *TCSPackProto) UnPack(msg []byte) (err error) {
	if len(msg) < 4 {
		err = errors.New("msg head id size invalid.")
		return
	}
	var pos int
	this.mainID = binary.LittleEndian.Uint32(msg[pos:])
	pos += 4

	if len(msg[pos:]) < 4 {
		err = errors.New("msg data size field invalid.")
		return
	}

	this.dataSize = binary.LittleEndian.Uint32(msg[pos:])
	pos += 4
	if len(msg[pos:]) < int(this.dataSize) {
		err = errors.New(fmt.Sprintf("msg data content size invalid, info: %v, %v, %v.", msg, len(msg[pos:]), int(this.dataSize)))
		return
	}
	this.data = msg[pos : pos+int(this.dataSize)]
	return
}

func (this *TCSPackProto) Pack(out []byte) {
	defer this.Reset()

	var pos int
	binary.LittleEndian.PutUint32(out[pos:], this.mainID)
	pos += 4

	binary.LittleEndian.PutUint32(out[pos:], this.dataSize)
	pos += 4

	copy(out[pos:], this.data)
}

func (this *TCSPackProto) GetMsgID() uint32 {
	return this.mainID
}

func (this *TCSPackProto) GetData() []byte {
	return this.data
}

func (this *TCSPackProto) Reset() {
	this.mainID = 0
	this.dataSize = 0
	this.data = nil
}

var _cspt = new(TCSPackProto)

func CSPackTool() *TCSPackProto {
	return _cspt
}
