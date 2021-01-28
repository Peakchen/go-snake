// Code generated by protoc-gen-go. DO NOT EDIT.
// source: MSG_CenterGate.proto

package MSG_CenterGate // import "msgProto/MSG_CenterGate"

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

// add by stefan
// server
type SUBMSG int32

const (
	SUBMSG_Begin                   SUBMSG = 0
	SUBMSG_CS_PlayerOnline         SUBMSG = 1
	SUBMSG_SC_PlayerOnline         SUBMSG = 2
	SUBMSG_CS_PlayerOffline        SUBMSG = 3
	SUBMSG_SC_PlayerOffline        SUBMSG = 4
	SUBMSG_CS_GetBroadCastSessions SUBMSG = 5
	SUBMSG_SC_GetBroadCastSessions SUBMSG = 6
)

var SUBMSG_name = map[int32]string{
	0: "Begin",
	1: "CS_PlayerOnline",
	2: "SC_PlayerOnline",
	3: "CS_PlayerOffline",
	4: "SC_PlayerOffline",
	5: "CS_GetBroadCastSessions",
	6: "SC_GetBroadCastSessions",
}
var SUBMSG_value = map[string]int32{
	"Begin":                   0,
	"CS_PlayerOnline":         1,
	"SC_PlayerOnline":         2,
	"CS_PlayerOffline":        3,
	"SC_PlayerOffline":        4,
	"CS_GetBroadCastSessions": 5,
	"SC_GetBroadCastSessions": 6,
}

func (x SUBMSG) String() string {
	return proto.EnumName(SUBMSG_name, int32(x))
}
func (SUBMSG) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_MSG_CenterGate_5ccf95e5cb8e5b1b, []int{0}
}

type ErrorCode int32

const (
	ErrorCode_Invalid ErrorCode = 0
	ErrorCode_Success ErrorCode = 1
	ErrorCode_Fail    ErrorCode = 2
)

var ErrorCode_name = map[int32]string{
	0: "Invalid",
	1: "Success",
	2: "Fail",
}
var ErrorCode_value = map[string]int32{
	"Invalid": 0,
	"Success": 1,
	"Fail":    2,
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_MSG_CenterGate_5ccf95e5cb8e5b1b, []int{1}
}

// CS_PlayerOnline
type CS_PlayerOnline_Req struct {
	PlayerIdentify       string   `protobuf:"bytes,1,opt,name=PlayerIdentify,proto3" json:"PlayerIdentify,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_PlayerOnline_Req) Reset()         { *m = CS_PlayerOnline_Req{} }
func (m *CS_PlayerOnline_Req) String() string { return proto.CompactTextString(m) }
func (*CS_PlayerOnline_Req) ProtoMessage()    {}
func (*CS_PlayerOnline_Req) Descriptor() ([]byte, []int) {
	return fileDescriptor_MSG_CenterGate_5ccf95e5cb8e5b1b, []int{0}
}
func (m *CS_PlayerOnline_Req) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_PlayerOnline_Req.Unmarshal(m, b)
}
func (m *CS_PlayerOnline_Req) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_PlayerOnline_Req.Marshal(b, m, deterministic)
}
func (dst *CS_PlayerOnline_Req) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_PlayerOnline_Req.Merge(dst, src)
}
func (m *CS_PlayerOnline_Req) XXX_Size() int {
	return xxx_messageInfo_CS_PlayerOnline_Req.Size(m)
}
func (m *CS_PlayerOnline_Req) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_PlayerOnline_Req.DiscardUnknown(m)
}

var xxx_messageInfo_CS_PlayerOnline_Req proto.InternalMessageInfo

func (m *CS_PlayerOnline_Req) GetPlayerIdentify() string {
	if m != nil {
		return m.PlayerIdentify
	}
	return ""
}

// SC_PlayerOnline
type SC_PlayerOnline_Rsp struct {
	Ret                  ErrorCode `protobuf:"varint,1,opt,name=Ret,proto3,enum=MSG_CenterGate.ErrorCode" json:"Ret,omitempty"`
	PlayerIdentify       string    `protobuf:"bytes,2,opt,name=PlayerIdentify,proto3" json:"PlayerIdentify,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SC_PlayerOnline_Rsp) Reset()         { *m = SC_PlayerOnline_Rsp{} }
func (m *SC_PlayerOnline_Rsp) String() string { return proto.CompactTextString(m) }
func (*SC_PlayerOnline_Rsp) ProtoMessage()    {}
func (*SC_PlayerOnline_Rsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_MSG_CenterGate_5ccf95e5cb8e5b1b, []int{1}
}
func (m *SC_PlayerOnline_Rsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SC_PlayerOnline_Rsp.Unmarshal(m, b)
}
func (m *SC_PlayerOnline_Rsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SC_PlayerOnline_Rsp.Marshal(b, m, deterministic)
}
func (dst *SC_PlayerOnline_Rsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SC_PlayerOnline_Rsp.Merge(dst, src)
}
func (m *SC_PlayerOnline_Rsp) XXX_Size() int {
	return xxx_messageInfo_SC_PlayerOnline_Rsp.Size(m)
}
func (m *SC_PlayerOnline_Rsp) XXX_DiscardUnknown() {
	xxx_messageInfo_SC_PlayerOnline_Rsp.DiscardUnknown(m)
}

var xxx_messageInfo_SC_PlayerOnline_Rsp proto.InternalMessageInfo

func (m *SC_PlayerOnline_Rsp) GetRet() ErrorCode {
	if m != nil {
		return m.Ret
	}
	return ErrorCode_Invalid
}

func (m *SC_PlayerOnline_Rsp) GetPlayerIdentify() string {
	if m != nil {
		return m.PlayerIdentify
	}
	return ""
}

// CS_PlayerOffline
type CS_PlayerOffline_Req struct {
	PlayerIdentify       string   `protobuf:"bytes,1,opt,name=PlayerIdentify,proto3" json:"PlayerIdentify,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_PlayerOffline_Req) Reset()         { *m = CS_PlayerOffline_Req{} }
func (m *CS_PlayerOffline_Req) String() string { return proto.CompactTextString(m) }
func (*CS_PlayerOffline_Req) ProtoMessage()    {}
func (*CS_PlayerOffline_Req) Descriptor() ([]byte, []int) {
	return fileDescriptor_MSG_CenterGate_5ccf95e5cb8e5b1b, []int{2}
}
func (m *CS_PlayerOffline_Req) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_PlayerOffline_Req.Unmarshal(m, b)
}
func (m *CS_PlayerOffline_Req) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_PlayerOffline_Req.Marshal(b, m, deterministic)
}
func (dst *CS_PlayerOffline_Req) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_PlayerOffline_Req.Merge(dst, src)
}
func (m *CS_PlayerOffline_Req) XXX_Size() int {
	return xxx_messageInfo_CS_PlayerOffline_Req.Size(m)
}
func (m *CS_PlayerOffline_Req) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_PlayerOffline_Req.DiscardUnknown(m)
}

var xxx_messageInfo_CS_PlayerOffline_Req proto.InternalMessageInfo

func (m *CS_PlayerOffline_Req) GetPlayerIdentify() string {
	if m != nil {
		return m.PlayerIdentify
	}
	return ""
}

// SC_PlayerOffline
type SC_PlayerOffline_Rsp struct {
	Ret                  ErrorCode `protobuf:"varint,1,opt,name=Ret,proto3,enum=MSG_CenterGate.ErrorCode" json:"Ret,omitempty"`
	PlayerIdentify       string    `protobuf:"bytes,2,opt,name=PlayerIdentify,proto3" json:"PlayerIdentify,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SC_PlayerOffline_Rsp) Reset()         { *m = SC_PlayerOffline_Rsp{} }
func (m *SC_PlayerOffline_Rsp) String() string { return proto.CompactTextString(m) }
func (*SC_PlayerOffline_Rsp) ProtoMessage()    {}
func (*SC_PlayerOffline_Rsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_MSG_CenterGate_5ccf95e5cb8e5b1b, []int{3}
}
func (m *SC_PlayerOffline_Rsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SC_PlayerOffline_Rsp.Unmarshal(m, b)
}
func (m *SC_PlayerOffline_Rsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SC_PlayerOffline_Rsp.Marshal(b, m, deterministic)
}
func (dst *SC_PlayerOffline_Rsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SC_PlayerOffline_Rsp.Merge(dst, src)
}
func (m *SC_PlayerOffline_Rsp) XXX_Size() int {
	return xxx_messageInfo_SC_PlayerOffline_Rsp.Size(m)
}
func (m *SC_PlayerOffline_Rsp) XXX_DiscardUnknown() {
	xxx_messageInfo_SC_PlayerOffline_Rsp.DiscardUnknown(m)
}

var xxx_messageInfo_SC_PlayerOffline_Rsp proto.InternalMessageInfo

func (m *SC_PlayerOffline_Rsp) GetRet() ErrorCode {
	if m != nil {
		return m.Ret
	}
	return ErrorCode_Invalid
}

func (m *SC_PlayerOffline_Rsp) GetPlayerIdentify() string {
	if m != nil {
		return m.PlayerIdentify
	}
	return ""
}

// CS_GetBroadCastSessions
type CS_GetBroadCastSessions_Req struct {
	PlayerIdentifys      []string `protobuf:"bytes,1,rep,name=PlayerIdentifys,proto3" json:"PlayerIdentifys,omitempty"`
	Data                 string   `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_GetBroadCastSessions_Req) Reset()         { *m = CS_GetBroadCastSessions_Req{} }
func (m *CS_GetBroadCastSessions_Req) String() string { return proto.CompactTextString(m) }
func (*CS_GetBroadCastSessions_Req) ProtoMessage()    {}
func (*CS_GetBroadCastSessions_Req) Descriptor() ([]byte, []int) {
	return fileDescriptor_MSG_CenterGate_5ccf95e5cb8e5b1b, []int{4}
}
func (m *CS_GetBroadCastSessions_Req) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_GetBroadCastSessions_Req.Unmarshal(m, b)
}
func (m *CS_GetBroadCastSessions_Req) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_GetBroadCastSessions_Req.Marshal(b, m, deterministic)
}
func (dst *CS_GetBroadCastSessions_Req) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_GetBroadCastSessions_Req.Merge(dst, src)
}
func (m *CS_GetBroadCastSessions_Req) XXX_Size() int {
	return xxx_messageInfo_CS_GetBroadCastSessions_Req.Size(m)
}
func (m *CS_GetBroadCastSessions_Req) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_GetBroadCastSessions_Req.DiscardUnknown(m)
}

var xxx_messageInfo_CS_GetBroadCastSessions_Req proto.InternalMessageInfo

func (m *CS_GetBroadCastSessions_Req) GetPlayerIdentifys() []string {
	if m != nil {
		return m.PlayerIdentifys
	}
	return nil
}

func (m *CS_GetBroadCastSessions_Req) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

// SC_GetBroadCastSessions
type SC_GetBroadCastSessions_Rsp struct {
	Ret                  ErrorCode `protobuf:"varint,1,opt,name=Ret,proto3,enum=MSG_CenterGate.ErrorCode" json:"Ret,omitempty"`
	PlayerIdentifys      []string  `protobuf:"bytes,2,rep,name=PlayerIdentifys,proto3" json:"PlayerIdentifys,omitempty"`
	Data                 string    `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SC_GetBroadCastSessions_Rsp) Reset()         { *m = SC_GetBroadCastSessions_Rsp{} }
func (m *SC_GetBroadCastSessions_Rsp) String() string { return proto.CompactTextString(m) }
func (*SC_GetBroadCastSessions_Rsp) ProtoMessage()    {}
func (*SC_GetBroadCastSessions_Rsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_MSG_CenterGate_5ccf95e5cb8e5b1b, []int{5}
}
func (m *SC_GetBroadCastSessions_Rsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SC_GetBroadCastSessions_Rsp.Unmarshal(m, b)
}
func (m *SC_GetBroadCastSessions_Rsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SC_GetBroadCastSessions_Rsp.Marshal(b, m, deterministic)
}
func (dst *SC_GetBroadCastSessions_Rsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SC_GetBroadCastSessions_Rsp.Merge(dst, src)
}
func (m *SC_GetBroadCastSessions_Rsp) XXX_Size() int {
	return xxx_messageInfo_SC_GetBroadCastSessions_Rsp.Size(m)
}
func (m *SC_GetBroadCastSessions_Rsp) XXX_DiscardUnknown() {
	xxx_messageInfo_SC_GetBroadCastSessions_Rsp.DiscardUnknown(m)
}

var xxx_messageInfo_SC_GetBroadCastSessions_Rsp proto.InternalMessageInfo

func (m *SC_GetBroadCastSessions_Rsp) GetRet() ErrorCode {
	if m != nil {
		return m.Ret
	}
	return ErrorCode_Invalid
}

func (m *SC_GetBroadCastSessions_Rsp) GetPlayerIdentifys() []string {
	if m != nil {
		return m.PlayerIdentifys
	}
	return nil
}

func (m *SC_GetBroadCastSessions_Rsp) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*CS_PlayerOnline_Req)(nil), "MSG_CenterGate.CS_PlayerOnline_Req")
	proto.RegisterType((*SC_PlayerOnline_Rsp)(nil), "MSG_CenterGate.SC_PlayerOnline_Rsp")
	proto.RegisterType((*CS_PlayerOffline_Req)(nil), "MSG_CenterGate.CS_PlayerOffline_Req")
	proto.RegisterType((*SC_PlayerOffline_Rsp)(nil), "MSG_CenterGate.SC_PlayerOffline_Rsp")
	proto.RegisterType((*CS_GetBroadCastSessions_Req)(nil), "MSG_CenterGate.CS_GetBroadCastSessions_Req")
	proto.RegisterType((*SC_GetBroadCastSessions_Rsp)(nil), "MSG_CenterGate.SC_GetBroadCastSessions_Rsp")
	proto.RegisterEnum("MSG_CenterGate.SUBMSG", SUBMSG_name, SUBMSG_value)
	proto.RegisterEnum("MSG_CenterGate.ErrorCode", ErrorCode_name, ErrorCode_value)
}

func init() {
	proto.RegisterFile("MSG_CenterGate.proto", fileDescriptor_MSG_CenterGate_5ccf95e5cb8e5b1b)
}

var fileDescriptor_MSG_CenterGate_5ccf95e5cb8e5b1b = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0x4d, 0x4f, 0xc2, 0x30,
	0x18, 0xc7, 0xd9, 0x06, 0xe8, 0x1e, 0x13, 0x68, 0xca, 0x12, 0x20, 0x5c, 0x08, 0x07, 0x43, 0x30,
	0x81, 0x44, 0xcf, 0x7a, 0x58, 0xd5, 0x85, 0x03, 0x91, 0xac, 0xf1, 0xa2, 0x87, 0xa5, 0xb2, 0x42,
	0xa6, 0xb3, 0xc3, 0xb6, 0x9a, 0xf0, 0x0d, 0xfc, 0x1e, 0x7e, 0x51, 0xb3, 0xf9, 0x42, 0x68, 0xc0,
	0x84, 0x83, 0xb7, 0xed, 0xdf, 0xfe, 0xfb, 0xfc, 0x9e, 0x37, 0xf0, 0x26, 0x34, 0x88, 0x08, 0x17,
	0x9a, 0xcb, 0x80, 0x69, 0x3e, 0x5c, 0xca, 0x4c, 0x67, 0xb8, 0xb6, 0xa9, 0xf6, 0xce, 0xa1, 0x41,
	0x68, 0x34, 0x4d, 0xd9, 0x8a, 0xcb, 0x1b, 0x91, 0x26, 0x82, 0x47, 0x21, 0x7f, 0xc1, 0xc7, 0x50,
	0xfb, 0xd2, 0xc6, 0x31, 0x17, 0x3a, 0x99, 0xaf, 0x5a, 0x56, 0xd7, 0xea, 0xbb, 0xa1, 0xa1, 0xf6,
	0x1e, 0xa1, 0x41, 0x89, 0x61, 0x57, 0x4b, 0x7c, 0x02, 0x4e, 0xc8, 0x75, 0xe1, 0xa9, 0x9d, 0xb6,
	0x87, 0x06, 0xc9, 0x95, 0x94, 0x99, 0x24, 0x59, 0xcc, 0xc3, 0xfc, 0xd6, 0x96, 0x58, 0xf6, 0xd6,
	0x58, 0x17, 0xe0, 0xad, 0x51, 0xe7, 0xf3, 0xbd, 0x59, 0x9f, 0xc0, 0x5b, 0xb3, 0xfe, 0xf8, 0xff,
	0x0b, 0xf6, 0x1e, 0x3a, 0x84, 0x46, 0x01, 0xd7, 0xbe, 0xcc, 0x58, 0x4c, 0x98, 0xd2, 0x94, 0x2b,
	0x95, 0x64, 0x42, 0x15, 0xcc, 0x7d, 0xa8, 0x6f, 0x1a, 0x54, 0xcb, 0xea, 0x3a, 0x7d, 0x37, 0x34,
	0x65, 0x8c, 0xa1, 0x7c, 0xc9, 0x34, 0xfb, 0x0e, 0x53, 0x7c, 0xf7, 0xde, 0x2d, 0xe8, 0x50, 0xb2,
	0xe3, 0xf5, 0x7d, 0x33, 0xda, 0x82, 0x62, 0xff, 0x8d, 0xe2, 0xac, 0x51, 0x06, 0x1f, 0x16, 0x54,
	0xe9, 0xad, 0x3f, 0xa1, 0x01, 0x76, 0xa1, 0xe2, 0xf3, 0x45, 0x22, 0x50, 0x09, 0x37, 0xa0, 0x6e,
	0x4c, 0x15, 0xb2, 0x72, 0xd1, 0x98, 0x15, 0x64, 0x63, 0x0f, 0x90, 0xd9, 0x54, 0xe4, 0xe4, 0xaa,
	0xd9, 0x2a, 0x54, 0xc6, 0x1d, 0x68, 0xee, 0xa8, 0x29, 0xaa, 0xe4, 0x87, 0x3b, 0x4a, 0x82, 0xaa,
	0x83, 0x11, 0xb8, 0xbf, 0x59, 0xe3, 0x23, 0x38, 0x18, 0x8b, 0x37, 0x96, 0x26, 0x31, 0x2a, 0xe5,
	0x3f, 0xf4, 0x75, 0x36, 0xe3, 0x4a, 0x21, 0x0b, 0x1f, 0x42, 0xf9, 0x9a, 0x25, 0x29, 0xb2, 0xfd,
	0xf6, 0x5d, 0xf3, 0x59, 0x2d, 0xa6, 0xf9, 0xca, 0x8c, 0x36, 0xcb, 0xf7, 0x50, 0x2d, 0x16, 0xe9,
	0xec, 0x33, 0x00, 0x00, 0xff, 0xff, 0xee, 0xbc, 0x76, 0x82, 0x60, 0x03, 0x00, 0x00,
}
