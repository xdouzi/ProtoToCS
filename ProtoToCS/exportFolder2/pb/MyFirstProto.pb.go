// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: MyFirstProto.proto

package pb

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

type BuffData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // 配表id
	BuffSid   int32 `protobuf:"varint,2,opt,name=buffSid,proto3" json:"buffSid,omitempty"`
	Opt       int32 `protobuf:"varint,3,opt,name=opt,proto3" json:"opt,omitempty"`
	DelReason int32 `protobuf:"varint,4,opt,name=delReason,proto3" json:"delReason,omitempty"`
	LeftRound int32 `protobuf:"varint,5,opt,name=leftRound,proto3" json:"leftRound,omitempty"`
	Owner     int32 `protobuf:"varint,6,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *BuffData) Reset() {
	*x = BuffData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MyFirstProto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuffData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuffData) ProtoMessage() {}

func (x *BuffData) ProtoReflect() protoreflect.Message {
	mi := &file_MyFirstProto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuffData.ProtoReflect.Descriptor instead.
func (*BuffData) Descriptor() ([]byte, []int) {
	return file_MyFirstProto_proto_rawDescGZIP(), []int{0}
}

func (x *BuffData) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BuffData) GetBuffSid() int32 {
	if x != nil {
		return x.BuffSid
	}
	return 0
}

func (x *BuffData) GetOpt() int32 {
	if x != nil {
		return x.Opt
	}
	return 0
}

func (x *BuffData) GetDelReason() int32 {
	if x != nil {
		return x.DelReason
	}
	return 0
}

func (x *BuffData) GetLeftRound() int32 {
	if x != nil {
		return x.LeftRound
	}
	return 0
}

func (x *BuffData) GetOwner() int32 {
	if x != nil {
		return x.Owner
	}
	return 0
}

var File_MyFirstProto_proto protoreflect.FileDescriptor

var file_MyFirstProto_proto_rawDesc = []byte{
	0x0a, 0x12, 0x4d, 0x79, 0x46, 0x69, 0x72, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x48, 0x6f, 0x74, 0x66, 0x69, 0x78, 0x2e, 0x4e, 0x65, 0x74,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x6c, 0x22, 0x98, 0x01, 0x0a, 0x08, 0x42, 0x75, 0x66, 0x66,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x75, 0x66, 0x66, 0x53, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x62, 0x75, 0x66, 0x66, 0x53, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x6f, 0x70, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6f, 0x70, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1c,
	0x0a, 0x09, 0x6c, 0x65, 0x66, 0x74, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x6c, 0x65, 0x66, 0x74, 0x52, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_MyFirstProto_proto_rawDescOnce sync.Once
	file_MyFirstProto_proto_rawDescData = file_MyFirstProto_proto_rawDesc
)

func file_MyFirstProto_proto_rawDescGZIP() []byte {
	file_MyFirstProto_proto_rawDescOnce.Do(func() {
		file_MyFirstProto_proto_rawDescData = protoimpl.X.CompressGZIP(file_MyFirstProto_proto_rawDescData)
	})
	return file_MyFirstProto_proto_rawDescData
}

var file_MyFirstProto_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MyFirstProto_proto_goTypes = []interface{}{
	(*BuffData)(nil), // 0: Hotfix.Net.Protol.BuffData
}
var file_MyFirstProto_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MyFirstProto_proto_init() }
func file_MyFirstProto_proto_init() {
	if File_MyFirstProto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MyFirstProto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuffData); i {
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
			RawDescriptor: file_MyFirstProto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MyFirstProto_proto_goTypes,
		DependencyIndexes: file_MyFirstProto_proto_depIdxs,
		MessageInfos:      file_MyFirstProto_proto_msgTypes,
	}.Build()
	File_MyFirstProto_proto = out.File
	file_MyFirstProto_proto_rawDesc = nil
	file_MyFirstProto_proto_goTypes = nil
	file_MyFirstProto_proto_depIdxs = nil
}