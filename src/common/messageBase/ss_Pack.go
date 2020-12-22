package messageBase

import (
	"encoding/binary"
	"errors"
)

//TSSPackProto pack proto struct to data or unpack...
type TSSPackProto struct {
	mainID   uint32
	sessid   string
	uid      int64
	dataSize uint32
	data     []byte
}

func (this *TSSPackProto) Init(sid string, uid int64, id uint32, src []byte) {
	this.mainID = id
	this.dataSize = uint32(len(src))
	this.data = src
	this.sessid = sid
	this.uid = uid
}

func (this *TSSPackProto) UnPack(msg []byte) (err error) {
	if len(msg) < SS_MSG_PACK_DATA_SIZE {
		err = errors.New("msg head id size invalid.")
		return
	}
	var pos int
	this.mainID = binary.LittleEndian.Uint32(msg[pos:])
	pos += 4

	this.sessid = string(msg[pos : pos+SS_MSG_PACK_SESSION_SIZE])
	pos += SS_MSG_PACK_SESSION_SIZE

	this.uid = int64(binary.LittleEndian.Uint64(msg[pos:]))
	pos += SS_MSG_PACK_UID_SIZE

	this.dataSize = binary.LittleEndian.Uint32(msg[pos:])
	pos += 4
	if len(msg[pos:]) < int(this.dataSize) {
		err = errors.New("msg data content size invalid.")
		return
	}
	this.data = msg[pos : pos+int(this.dataSize)]
	return
}

func (this *TSSPackProto) Pack(out []byte) {
	defer this.Reset()

	var pos int
	binary.LittleEndian.PutUint32(out[pos:], this.mainID)
	pos += 4

	copy(out[pos:], this.sessid)
	pos += len(this.sessid)

	binary.LittleEndian.PutUint64(out[pos:], uint64(this.uid))
	pos += SS_MSG_PACK_UID_SIZE

	binary.LittleEndian.PutUint32(out[pos:], this.dataSize)
	pos += 4

	copy(out[pos:], this.data)
}

func (this *TSSPackProto) GetMsgID() uint32 {
	return this.mainID
}

func (this *TSSPackProto) GetData() []byte {
	return this.data
}

func (this *TSSPackProto) GetUID() int64 {
	return this.uid
}

func (this *TSSPackProto) GetSessID() string {
	return this.sessid
}

func (this *TSSPackProto) Reset() {
	this.mainID = 0
	this.dataSize = 0
	this.data = nil
	this.sessid = ""
	this.uid = 0
}

var _sspt = new(TSSPackProto)

func SSPackTool() *TSSPackProto {
	return _sspt
}
