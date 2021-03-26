// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: logic.proto

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

//进入游戏
type CS_EnterGameScene struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CS_EnterGameScene) Reset() {
	*x = CS_EnterGameScene{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CS_EnterGameScene) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CS_EnterGameScene) ProtoMessage() {}

func (x *CS_EnterGameScene) ProtoReflect() protoreflect.Message {
	mi := &file_logic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CS_EnterGameScene.ProtoReflect.Descriptor instead.
func (*CS_EnterGameScene) Descriptor() ([]byte, []int) {
	return file_logic_proto_rawDescGZIP(), []int{0}
}

//进入游戏
type SC_EnterGameScene struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SC_EnterGameScene) Reset() {
	*x = SC_EnterGameScene{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SC_EnterGameScene) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_EnterGameScene) ProtoMessage() {}

func (x *SC_EnterGameScene) ProtoReflect() protoreflect.Message {
	mi := &file_logic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_EnterGameScene.ProtoReflect.Descriptor instead.
func (*SC_EnterGameScene) Descriptor() ([]byte, []int) {
	return file_logic_proto_rawDescGZIP(), []int{1}
}

type CS_HeartBeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CS_HeartBeat) Reset() {
	*x = CS_HeartBeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CS_HeartBeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CS_HeartBeat) ProtoMessage() {}

func (x *CS_HeartBeat) ProtoReflect() protoreflect.Message {
	mi := &file_logic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CS_HeartBeat.ProtoReflect.Descriptor instead.
func (*CS_HeartBeat) Descriptor() ([]byte, []int) {
	return file_logic_proto_rawDescGZIP(), []int{2}
}

type SC_HeartBeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SC_HeartBeat) Reset() {
	*x = SC_HeartBeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logic_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SC_HeartBeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SC_HeartBeat) ProtoMessage() {}

func (x *SC_HeartBeat) ProtoReflect() protoreflect.Message {
	mi := &file_logic_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SC_HeartBeat.ProtoReflect.Descriptor instead.
func (*SC_HeartBeat) Descriptor() ([]byte, []int) {
	return file_logic_proto_rawDescGZIP(), []int{3}
}

var File_logic_proto protoreflect.FileDescriptor

var file_logic_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x13, 0x0a,
	0x11, 0x43, 0x53, 0x5f, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x63, 0x65,
	0x6e, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x53, 0x43, 0x5f, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x47, 0x61,
	0x6d, 0x65, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x43, 0x53, 0x5f, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x53, 0x43, 0x5f, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2e, 0x2f, 0x61, 0x6b,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_logic_proto_rawDescOnce sync.Once
	file_logic_proto_rawDescData = file_logic_proto_rawDesc
)

func file_logic_proto_rawDescGZIP() []byte {
	file_logic_proto_rawDescOnce.Do(func() {
		file_logic_proto_rawDescData = protoimpl.X.CompressGZIP(file_logic_proto_rawDescData)
	})
	return file_logic_proto_rawDescData
}

var file_logic_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_logic_proto_goTypes = []interface{}{
	(*CS_EnterGameScene)(nil), // 0: CS_EnterGameScene
	(*SC_EnterGameScene)(nil), // 1: SC_EnterGameScene
	(*CS_HeartBeat)(nil),      // 2: CS_HeartBeat
	(*SC_HeartBeat)(nil),      // 3: SC_HeartBeat
}
var file_logic_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_logic_proto_init() }
func file_logic_proto_init() {
	if File_logic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_logic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CS_EnterGameScene); i {
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
		file_logic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SC_EnterGameScene); i {
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
		file_logic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CS_HeartBeat); i {
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
		file_logic_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SC_HeartBeat); i {
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
			RawDescriptor: file_logic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_logic_proto_goTypes,
		DependencyIndexes: file_logic_proto_depIdxs,
		MessageInfos:      file_logic_proto_msgTypes,
	}.Build()
	File_logic_proto = out.File
	file_logic_proto_rawDesc = nil
	file_logic_proto_goTypes = nil
	file_logic_proto_depIdxs = nil
}
