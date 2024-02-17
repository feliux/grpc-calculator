// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: primes.proto

package proto

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

type PrimesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *PrimesRequest) Reset() {
	*x = PrimesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_primes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimesRequest) ProtoMessage() {}

func (x *PrimesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_primes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimesRequest.ProtoReflect.Descriptor instead.
func (*PrimesRequest) Descriptor() ([]byte, []int) {
	return file_primes_proto_rawDescGZIP(), []int{0}
}

func (x *PrimesRequest) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type PrimesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *PrimesResponse) Reset() {
	*x = PrimesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_primes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimesResponse) ProtoMessage() {}

func (x *PrimesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_primes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimesResponse.ProtoReflect.Descriptor instead.
func (*PrimesResponse) Descriptor() ([]byte, []int) {
	return file_primes_proto_rawDescGZIP(), []int{1}
}

func (x *PrimesResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_primes_proto protoreflect.FileDescriptor

var file_primes_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x25, 0x0a, 0x0d, 0x50, 0x72,
	0x69, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x28, 0x0a, 0x0e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x29, 0x5a, 0x27, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x65, 0x6c, 0x69, 0x75, 0x78,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_primes_proto_rawDescOnce sync.Once
	file_primes_proto_rawDescData = file_primes_proto_rawDesc
)

func file_primes_proto_rawDescGZIP() []byte {
	file_primes_proto_rawDescOnce.Do(func() {
		file_primes_proto_rawDescData = protoimpl.X.CompressGZIP(file_primes_proto_rawDescData)
	})
	return file_primes_proto_rawDescData
}

var file_primes_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_primes_proto_goTypes = []interface{}{
	(*PrimesRequest)(nil),  // 0: calculator.PrimesRequest
	(*PrimesResponse)(nil), // 1: calculator.PrimesResponse
}
var file_primes_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_primes_proto_init() }
func file_primes_proto_init() {
	if File_primes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_primes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimesRequest); i {
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
		file_primes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimesResponse); i {
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
			RawDescriptor: file_primes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_primes_proto_goTypes,
		DependencyIndexes: file_primes_proto_depIdxs,
		MessageInfos:      file_primes_proto_msgTypes,
	}.Build()
	File_primes_proto = out.File
	file_primes_proto_rawDesc = nil
	file_primes_proto_goTypes = nil
	file_primes_proto_depIdxs = nil
}
