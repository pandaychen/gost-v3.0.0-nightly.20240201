// protoc --go_out=. --go_opt=paths=source_relative \
//	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
//	recorder.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: recorder.proto

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

type RecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data     []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Metadata []byte `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *RecordRequest) Reset() {
	*x = RecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recorder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordRequest) ProtoMessage() {}

func (x *RecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_recorder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordRequest.ProtoReflect.Descriptor instead.
func (*RecordRequest) Descriptor() ([]byte, []int) {
	return file_recorder_proto_rawDescGZIP(), []int{0}
}

func (x *RecordRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *RecordRequest) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type RecordReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *RecordReply) Reset() {
	*x = RecordReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recorder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordReply) ProtoMessage() {}

func (x *RecordReply) ProtoReflect() protoreflect.Message {
	mi := &file_recorder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordReply.ProtoReflect.Descriptor instead.
func (*RecordReply) Descriptor() ([]byte, []int) {
	return file_recorder_proto_rawDescGZIP(), []int{1}
}

func (x *RecordReply) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_recorder_proto protoreflect.FileDescriptor

var file_recorder_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x0d, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1d, 0x0a, 0x0b, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x32, 0x3e, 0x0a, 0x08, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x6f, 0x73, 0x74, 0x2f, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2f, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_recorder_proto_rawDescOnce sync.Once
	file_recorder_proto_rawDescData = file_recorder_proto_rawDesc
)

func file_recorder_proto_rawDescGZIP() []byte {
	file_recorder_proto_rawDescOnce.Do(func() {
		file_recorder_proto_rawDescData = protoimpl.X.CompressGZIP(file_recorder_proto_rawDescData)
	})
	return file_recorder_proto_rawDescData
}

var file_recorder_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_recorder_proto_goTypes = []interface{}{
	(*RecordRequest)(nil), // 0: proto.RecordRequest
	(*RecordReply)(nil),   // 1: proto.RecordReply
}
var file_recorder_proto_depIdxs = []int32{
	0, // 0: proto.Recorder.Record:input_type -> proto.RecordRequest
	1, // 1: proto.Recorder.Record:output_type -> proto.RecordReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_recorder_proto_init() }
func file_recorder_proto_init() {
	if File_recorder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_recorder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordRequest); i {
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
		file_recorder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordReply); i {
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
			RawDescriptor: file_recorder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_recorder_proto_goTypes,
		DependencyIndexes: file_recorder_proto_depIdxs,
		MessageInfos:      file_recorder_proto_msgTypes,
	}.Build()
	File_recorder_proto = out.File
	file_recorder_proto_rawDesc = nil
	file_recorder_proto_goTypes = nil
	file_recorder_proto_depIdxs = nil
}
