// Code generated by protoc-gen-go. DO NOT EDIT.
// source: error_message.proto

package ErrMessage

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ErrMessageId int32

const (
	ErrMessageId_msg_err_invalid       ErrMessageId = 0
	ErrMessageId_msg_err_timeout       ErrMessageId = 1
	ErrMessageId_msg_err_disconnect    ErrMessageId = 2
	ErrMessageId_msg_err_login_faild   ErrMessageId = 3
	ErrMessageId_msg_err_not_find_user ErrMessageId = 4
)

var ErrMessageId_name = map[int32]string{
	0: "msg_err_invalid",
	1: "msg_err_timeout",
	2: "msg_err_disconnect",
	3: "msg_err_login_faild",
	4: "msg_err_not_find_user",
}
var ErrMessageId_value = map[string]int32{
	"msg_err_invalid":       0,
	"msg_err_timeout":       1,
	"msg_err_disconnect":    2,
	"msg_err_login_faild":   3,
	"msg_err_not_find_user": 4,
}

func (x ErrMessageId) Enum() *ErrMessageId {
	p := new(ErrMessageId)
	*p = x
	return p
}
func (x ErrMessageId) String() string {
	return proto.EnumName(ErrMessageId_name, int32(x))
}
func (x *ErrMessageId) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ErrMessageId_value, data, "ErrMessageId")
	if err != nil {
		return err
	}
	*x = ErrMessageId(value)
	return nil
}
func (ErrMessageId) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_error_message_5be0330f23d25c93, []int{0}
}

func init() {
	proto.RegisterEnum("ErrMessage.ErrMessageId", ErrMessageId_name, ErrMessageId_value)
}

func init() { proto.RegisterFile("error_message.proto", fileDescriptor_error_message_5be0330f23d25c93) }

var fileDescriptor_error_message_5be0330f23d25c93 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xca, 0x31, 0x0e, 0xc2, 0x30,
	0x0c, 0x40, 0x51, 0x0a, 0x4c, 0x1e, 0x20, 0x72, 0x05, 0x88, 0x2b, 0x30, 0x70, 0x8b, 0x8e, 0x9c,
	0xc1, 0xaa, 0x1a, 0xb7, 0xb2, 0x94, 0xd8, 0xc8, 0x49, 0x39, 0x02, 0xe7, 0x66, 0x40, 0x41, 0xac,
	0xef, 0x7f, 0xe8, 0xd9, 0xdd, 0x9c, 0x32, 0x97, 0x32, 0x2e, 0x7c, 0x7f, 0xba, 0x55, 0x43, 0x18,
	0xdc, 0x1f, 0x5f, 0xb9, 0xbd, 0x3b, 0x38, 0x0c, 0xfe, 0x3b, 0x48, 0x22, 0xf6, 0x70, 0xcc, 0x65,
	0x21, 0x76, 0x27, 0xd1, 0xd7, 0x98, 0x24, 0x86, 0xcd, 0x3f, 0x56, 0xc9, 0x6c, 0x6b, 0x0d, 0x1d,
	0x9e, 0x01, 0x1b, 0x46, 0x29, 0x93, 0xa9, 0xf2, 0x54, 0xc3, 0x16, 0x2f, 0xd0, 0x37, 0x4f, 0xb6,
	0x88, 0xd2, 0x3c, 0x4a, 0x8a, 0x61, 0x87, 0x57, 0x38, 0xb5, 0xa0, 0x56, 0x69, 0x16, 0x8d, 0xb4,
	0x16, 0xf6, 0xb0, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x5e, 0xe2, 0xe6, 0x83, 0xaa, 0x00, 0x00,
	0x00,
}