// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: pb/shortener.proto

package pb

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

type Url struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=Url,proto3" json:"Url,omitempty"`
}

func (x *Url) Reset() {
	*x = Url{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_shortener_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Url) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Url) ProtoMessage() {}

func (x *Url) ProtoReflect() protoreflect.Message {
	mi := &file_pb_shortener_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Url.ProtoReflect.Descriptor instead.
func (*Url) Descriptor() ([]byte, []int) {
	return file_pb_shortener_proto_rawDescGZIP(), []int{0}
}

func (x *Url) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ShortenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url           string `protobuf:"bytes,1,opt,name=Url,proto3" json:"Url,omitempty"`
	RequestedStub string `protobuf:"bytes,2,opt,name=RequestedStub,proto3" json:"RequestedStub,omitempty"`
}

func (x *ShortenReq) Reset() {
	*x = ShortenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_shortener_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenReq) ProtoMessage() {}

func (x *ShortenReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_shortener_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenReq.ProtoReflect.Descriptor instead.
func (*ShortenReq) Descriptor() ([]byte, []int) {
	return file_pb_shortener_proto_rawDescGZIP(), []int{1}
}

func (x *ShortenReq) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ShortenReq) GetRequestedStub() string {
	if x != nil {
		return x.RequestedStub
	}
	return ""
}

type Short struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stub string `protobuf:"bytes,1,opt,name=stub,proto3" json:"stub,omitempty"`
}

func (x *Short) Reset() {
	*x = Short{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_shortener_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Short) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Short) ProtoMessage() {}

func (x *Short) ProtoReflect() protoreflect.Message {
	mi := &file_pb_shortener_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Short.ProtoReflect.Descriptor instead.
func (*Short) Descriptor() ([]byte, []int) {
	return file_pb_shortener_proto_rawDescGZIP(), []int{2}
}

func (x *Short) GetStub() string {
	if x != nil {
		return x.Stub
	}
	return ""
}

var File_pb_shortener_proto protoreflect.FileDescriptor

var file_pb_shortener_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x62, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x17, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x12,
	0x10, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x72,
	0x6c, 0x22, 0x44, 0x0a, 0x0a, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12,
	0x10, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x72,
	0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x53, 0x74,
	0x75, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x53, 0x74, 0x75, 0x62, 0x22, 0x1b, 0x0a, 0x05, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x75, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x73, 0x74, 0x75, 0x62, 0x32, 0x56, 0x0a, 0x09, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x12, 0x26, 0x0a, 0x07, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x12, 0x0e, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x09, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x22, 0x00, 0x12, 0x21, 0x0a, 0x09, 0x55, 0x6e, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x12, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x1a, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x72, 0x6c, 0x22, 0x00, 0x42, 0x29, 0x5a, 0x27,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x61, 0x76, 0x61, 0x7a,
	0x2d, 0x61, 0x6c, 0x61, 0x6e, 0x69, 0x2f, 0x75, 0x72, 0x6c, 0x2d, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_shortener_proto_rawDescOnce sync.Once
	file_pb_shortener_proto_rawDescData = file_pb_shortener_proto_rawDesc
)

func file_pb_shortener_proto_rawDescGZIP() []byte {
	file_pb_shortener_proto_rawDescOnce.Do(func() {
		file_pb_shortener_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_shortener_proto_rawDescData)
	})
	return file_pb_shortener_proto_rawDescData
}

var file_pb_shortener_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pb_shortener_proto_goTypes = []interface{}{
	(*Url)(nil),        // 0: pb.Url
	(*ShortenReq)(nil), // 1: pb.ShortenReq
	(*Short)(nil),      // 2: pb.Short
}
var file_pb_shortener_proto_depIdxs = []int32{
	1, // 0: pb.Shortener.Shorten:input_type -> pb.ShortenReq
	2, // 1: pb.Shortener.UnShorten:input_type -> pb.Short
	2, // 2: pb.Shortener.Shorten:output_type -> pb.Short
	0, // 3: pb.Shortener.UnShorten:output_type -> pb.Url
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_shortener_proto_init() }
func file_pb_shortener_proto_init() {
	if File_pb_shortener_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_shortener_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Url); i {
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
		file_pb_shortener_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortenReq); i {
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
		file_pb_shortener_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Short); i {
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
			RawDescriptor: file_pb_shortener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_shortener_proto_goTypes,
		DependencyIndexes: file_pb_shortener_proto_depIdxs,
		MessageInfos:      file_pb_shortener_proto_msgTypes,
	}.Build()
	File_pb_shortener_proto = out.File
	file_pb_shortener_proto_rawDesc = nil
	file_pb_shortener_proto_goTypes = nil
	file_pb_shortener_proto_depIdxs = nil
}