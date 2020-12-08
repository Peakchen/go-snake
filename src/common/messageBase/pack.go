package messageBase

import (
	"encoding/binary"
	"errors"
)

//TPackProto pack proto struct to data or unpack...
type TPackProto struct {
	mainID   uint32
	dataSize uint32
	data     []byte
}

func (this *TPackProto) Init(id uint32, size uint32, src []byte) {
	this.mainID = id
	this.dataSize = size
	this.data = src
}

func (this *TPackProto) UnPack(msg []byte) (err error) {
	if len(msg) < 4 {
		err = errors.New("msg head id size invalid.")
		return
	}
	var pos int
	this.mainID = binary.LittleEndian.Uint32(msg[pos:])
	pos += 4

	if len(msg[pos:]) <= 4 {
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
	binary.LittleEndian.PutUint32(out[pos:], this.mainID)
	pos += 4

	binary.LittleEndian.PutUint32(out[pos:], this.dataSize)
	pos += 4
	copy(out[pos:], this.data)
}

func (this *TPackProto) GetMsgID() uint32 {
	return this.mainID
}

func (this *TPackProto) GetData() []byte {
	return this.data
}

func PackTool() *TPackProto {
	return new(TPackProto)
}
