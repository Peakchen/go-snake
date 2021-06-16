// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: chat.proto

package akmessage

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EChatType int32

const (
	EChatType_None        EChatType = 0 //非法类型
	EChatType_Private     EChatType = 1 //私有聊天
	EChatType_Global      EChatType = 2 //世界聊天
	EChatType_Guild       EChatType = 3 //公会/帮会
	EChatType_FriendGroup EChatType = 4 //朋友分组
)

// Enum value maps for EChatType.
var (
	EChatType_name = map[int32]string{
		0: "None",
		1: "Private",
		2: "Global",
		3: "Guild",
		4: "FriendGroup",
	}
	EChatType_value = map[string]int32{
		"None":        0,
		"Private":     1,
		"Global":      2,
		"Guild":       3,
		"FriendGroup": 4,
	}
)

func (x EChatType) Enum() *EChatType {
	p := new(EChatType)
	*p = x
	return p
}

func (x EChatType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EChatType) Descriptor() protoreflect.EnumDescriptor {
	return file_chat_proto_enumTypes[0].Descriptor()
}

func (EChatType) Type() protoreflect.EnumType {
	return &file_chat_proto_enumTypes[0]
}

func (x EChatType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EChatType.Descriptor instead.
func (EChatType) EnumDescriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

type CS_Chat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    EChatType `protobuf:"varint,1,opt,name=Type,proto3,enum=EChatType" json:"Type,omitempty"`
	Content string    `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *CS_Chat) Reset() {
	*x = CS_Chat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CS_Chat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CS_Chat) ProtoMessage() {}

func (x *CS_Chat) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CS_Chat.ProtoReflect.Descriptor instead.
func (*CS_Chat) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

func (x *CS_Chat) GetType() EChatType {
	if x != nil {
		return x.Type
	}
	return EChatType_None
}

func (x *CS_Chat) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type SC_Chat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ret     ErrorCode `protobuf:"varint,1,opt,name=Ret,proto3,enum=ErrorCode" json:"Ret,omitempty"`
	Type    EChatType `protobuf:"varint,2,opt,name=Type,proto3,enum=EChatType" json:"Type,omitempty"`
	Content string    `protobuf:"bytes,3,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *SC_Chat) Reset() {
	*x = SC_Chat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SC_Chat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_Chat) ProtoMessage() {}

func (x *SC_Chat) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_Chat.ProtoReflect.Descriptor instead.
func (*SC_Chat) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{1}
}

func (x *SC_Chat) GetRet() ErrorCode {
	if x != nil {
		return x.Ret
	}
	return ErrorCode_Invaild
}

func (x *SC_Chat) GetType() EChatType {
	if x != nil {
		return x.Type
	}
	return EChatType_None
}

func (x *SC_Chat) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type SS_Chat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    EChatType `protobuf:"varint,1,opt,name=Type,proto3,enum=EChatType" json:"Type,omitempty"`
	Content string    `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *SS_Chat) Reset() {
	*x = SS_Chat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SS_Chat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SS_Chat) ProtoMessage() {}

func (x *SS_Chat) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SS_Chat.ProtoReflect.Descriptor instead.
func (*SS_Chat) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{2}
}

func (x *SS_Chat) GetType() EChatType {
	if x != nil {
		return x.Type
	}
	return EChatType_None
}

func (x *SS_Chat) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a,
	0x07, 0x43, 0x53, 0x5f, 0x43, 0x68, 0x61, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x45, 0x43, 0x68, 0x61, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x22, 0x61, 0x0a, 0x07, 0x53, 0x43, 0x5f, 0x43, 0x68, 0x61, 0x74, 0x12, 0x1c, 0x0a,
	0x03, 0x52, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x03, 0x52, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x45, 0x43, 0x68, 0x61,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x43, 0x0a, 0x07, 0x53, 0x53, 0x5f, 0x43, 0x68, 0x61, 0x74,
	0x12, 0x1e, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a,
	0x2e, 0x45, 0x43, 0x68, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2a, 0x4a, 0x0a, 0x09, 0x45, 0x43,
	0x68, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f, 0x6e, 0x65, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x10, 0x01, 0x12, 0x0a,
	0x0a, 0x06, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x75,
	0x69, 0x6c, 0x64, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x10, 0x04, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2e, 0x2f, 0x61, 0x6b, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData = file_chat_proto_rawDesc
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_proto_rawDescData)
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_chat_proto_goTypes = []interface{}{
	(EChatType)(0),  // 0: EChatType
	(*CS_Chat)(nil), // 1: CS_Chat
	(*SC_Chat)(nil), // 2: SC_Chat
	(*SS_Chat)(nil), // 3: SS_Chat
	(ErrorCode)(0),  // 4: ErrorCode
}
var file_chat_proto_depIdxs = []int32{
	0, // 0: CS_Chat.Type:type_name -> EChatType
	4, // 1: SC_Chat.Ret:type_name -> ErrorCode
	0, // 2: SC_Chat.Type:type_name -> EChatType
	0, // 3: SS_Chat.Type:type_name -> EChatType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	file_errorcode_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CS_Chat); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SC_Chat); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SS_Chat); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chat_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		EnumInfos:         file_chat_proto_enumTypes,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_rawDesc = nil
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}
