// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: inout.proto

package hive_go

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SecretResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Created int64  `protobuf:"varint,2,opt,name=created,proto3" json:"created,omitempty"`
	Value   []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SecretResponse) Reset() {
	*x = SecretResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inout_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretResponse) ProtoMessage() {}

func (x *SecretResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inout_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretResponse.ProtoReflect.Descriptor instead.
func (*SecretResponse) Descriptor() ([]byte, []int) {
	return file_inout_proto_rawDescGZIP(), []int{0}
}

func (x *SecretResponse) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *SecretResponse) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *SecretResponse) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type GetSecretResponseV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *SecretResponse `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetSecretResponseV1) Reset() {
	*x = GetSecretResponseV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inout_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSecretResponseV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSecretResponseV1) ProtoMessage() {}

func (x *GetSecretResponseV1) ProtoReflect() protoreflect.Message {
	mi := &file_inout_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSecretResponseV1.ProtoReflect.Descriptor instead.
func (*GetSecretResponseV1) Descriptor() ([]byte, []int) {
	return file_inout_proto_rawDescGZIP(), []int{1}
}

func (x *GetSecretResponseV1) GetData() *SecretResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

type SecretCreatedV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Created int64  `protobuf:"varint,2,opt,name=created,proto3" json:"created,omitempty"`
	Value   []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SecretCreatedV1) Reset() {
	*x = SecretCreatedV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inout_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretCreatedV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretCreatedV1) ProtoMessage() {}

func (x *SecretCreatedV1) ProtoReflect() protoreflect.Message {
	mi := &file_inout_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretCreatedV1.ProtoReflect.Descriptor instead.
func (*SecretCreatedV1) Descriptor() ([]byte, []int) {
	return file_inout_proto_rawDescGZIP(), []int{2}
}

func (x *SecretCreatedV1) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *SecretCreatedV1) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *SecretCreatedV1) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_inout_proto protoreflect.FileDescriptor

var file_inout_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x69, 0x6e, 0x6f, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x68,
	0x69, 0x76, 0x65, 0x22, 0x50, 0x0a, 0x0e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x31, 0x12, 0x28, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x68, 0x69, 0x76,
	0x65, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x51, 0x0a, 0x0f, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x56, 0x31, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x3b, 0x68,
	0x69, 0x76, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inout_proto_rawDescOnce sync.Once
	file_inout_proto_rawDescData = file_inout_proto_rawDesc
)

func file_inout_proto_rawDescGZIP() []byte {
	file_inout_proto_rawDescOnce.Do(func() {
		file_inout_proto_rawDescData = protoimpl.X.CompressGZIP(file_inout_proto_rawDescData)
	})
	return file_inout_proto_rawDescData
}

var file_inout_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_inout_proto_goTypes = []interface{}{
	(*SecretResponse)(nil),      // 0: hive.SecretResponse
	(*GetSecretResponseV1)(nil), // 1: hive.GetSecretResponseV1
	(*SecretCreatedV1)(nil),     // 2: hive.SecretCreatedV1
}
var file_inout_proto_depIdxs = []int32{
	0, // 0: hive.GetSecretResponseV1.data:type_name -> hive.SecretResponse
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_inout_proto_init() }
func file_inout_proto_init() {
	if File_inout_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_inout_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretResponse); i {
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
		file_inout_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSecretResponseV1); i {
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
		file_inout_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretCreatedV1); i {
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
			RawDescriptor: file_inout_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_inout_proto_goTypes,
		DependencyIndexes: file_inout_proto_depIdxs,
		MessageInfos:      file_inout_proto_msgTypes,
	}.Build()
	File_inout_proto = out.File
	file_inout_proto_rawDesc = nil
	file_inout_proto_goTypes = nil
	file_inout_proto_depIdxs = nil
}