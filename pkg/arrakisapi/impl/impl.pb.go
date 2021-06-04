// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: REDACTEDapi/impl/impl.proto

package zookie

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

type DecodedZookie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version uint32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Types that are assignable to VersionOneof:
	//	*DecodedZookie_V1
	VersionOneof isDecodedZookie_VersionOneof `protobuf_oneof:"version_oneof"`
}

func (x *DecodedZookie) Reset() {
	*x = DecodedZookie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_REDACTEDapi_impl_impl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecodedZookie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecodedZookie) ProtoMessage() {}

func (x *DecodedZookie) ProtoReflect() protoreflect.Message {
	mi := &file_REDACTEDapi_impl_impl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecodedZookie.ProtoReflect.Descriptor instead.
func (*DecodedZookie) Descriptor() ([]byte, []int) {
	return file_REDACTEDapi_impl_impl_proto_rawDescGZIP(), []int{0}
}

func (x *DecodedZookie) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (m *DecodedZookie) GetVersionOneof() isDecodedZookie_VersionOneof {
	if m != nil {
		return m.VersionOneof
	}
	return nil
}

func (x *DecodedZookie) GetV1() *DecodedZookie_V1Zookie {
	if x, ok := x.GetVersionOneof().(*DecodedZookie_V1); ok {
		return x.V1
	}
	return nil
}

type isDecodedZookie_VersionOneof interface {
	isDecodedZookie_VersionOneof()
}

type DecodedZookie_V1 struct {
	V1 *DecodedZookie_V1Zookie `protobuf:"bytes,2,opt,name=v1,proto3,oneof"`
}

func (*DecodedZookie_V1) isDecodedZookie_VersionOneof() {}

type DecodedZookie_V1Zookie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Revision uint64 `protobuf:"varint,1,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *DecodedZookie_V1Zookie) Reset() {
	*x = DecodedZookie_V1Zookie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_REDACTEDapi_impl_impl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecodedZookie_V1Zookie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecodedZookie_V1Zookie) ProtoMessage() {}

func (x *DecodedZookie_V1Zookie) ProtoReflect() protoreflect.Message {
	mi := &file_REDACTEDapi_impl_impl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecodedZookie_V1Zookie.ProtoReflect.Descriptor instead.
func (*DecodedZookie_V1Zookie) Descriptor() ([]byte, []int) {
	return file_REDACTEDapi_impl_impl_proto_rawDescGZIP(), []int{0, 0}
}

func (x *DecodedZookie_V1Zookie) GetRevision() uint64 {
	if x != nil {
		return x.Revision
	}
	return 0
}

var File_REDACTEDapi_impl_impl_proto protoreflect.FileDescriptor

var file_REDACTEDapi_impl_impl_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x72, 0x72, 0x61, 0x6b, 0x69, 0x73, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6d, 0x70,
	0x6c, 0x2f, 0x69, 0x6d, 0x70, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a,
	0x0d, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x5a, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x02, 0x76, 0x31, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x5a, 0x6f,
	0x6f, 0x6b, 0x69, 0x65, 0x2e, 0x56, 0x31, 0x5a, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x48, 0x00, 0x52,
	0x02, 0x76, 0x31, 0x1a, 0x26, 0x0a, 0x08, 0x56, 0x31, 0x5a, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x0f, 0x0a, 0x0d, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x42, 0x22, 0x5a, 0x20,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2f, 0x7a, 0x6f, 0x6f, 0x6b, 0x69, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_REDACTEDapi_impl_impl_proto_rawDescOnce sync.Once
	file_REDACTEDapi_impl_impl_proto_rawDescData = file_REDACTEDapi_impl_impl_proto_rawDesc
)

func file_REDACTEDapi_impl_impl_proto_rawDescGZIP() []byte {
	file_REDACTEDapi_impl_impl_proto_rawDescOnce.Do(func() {
		file_REDACTEDapi_impl_impl_proto_rawDescData = protoimpl.X.CompressGZIP(file_REDACTEDapi_impl_impl_proto_rawDescData)
	})
	return file_REDACTEDapi_impl_impl_proto_rawDescData
}

var file_REDACTEDapi_impl_impl_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_REDACTEDapi_impl_impl_proto_goTypes = []interface{}{
	(*DecodedZookie)(nil),          // 0: DecodedZookie
	(*DecodedZookie_V1Zookie)(nil), // 1: DecodedZookie.V1Zookie
}
var file_REDACTEDapi_impl_impl_proto_depIdxs = []int32{
	1, // 0: DecodedZookie.v1:type_name -> DecodedZookie.V1Zookie
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_REDACTEDapi_impl_impl_proto_init() }
func file_REDACTEDapi_impl_impl_proto_init() {
	if File_REDACTEDapi_impl_impl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_REDACTEDapi_impl_impl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecodedZookie); i {
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
		file_REDACTEDapi_impl_impl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecodedZookie_V1Zookie); i {
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
	file_REDACTEDapi_impl_impl_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*DecodedZookie_V1)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_REDACTEDapi_impl_impl_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_REDACTEDapi_impl_impl_proto_goTypes,
		DependencyIndexes: file_REDACTEDapi_impl_impl_proto_depIdxs,
		MessageInfos:      file_REDACTEDapi_impl_impl_proto_msgTypes,
	}.Build()
	File_REDACTEDapi_impl_impl_proto = out.File
	file_REDACTEDapi_impl_impl_proto_rawDesc = nil
	file_REDACTEDapi_impl_impl_proto_goTypes = nil
	file_REDACTEDapi_impl_impl_proto_depIdxs = nil
}
