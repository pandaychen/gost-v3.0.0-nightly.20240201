// protoc --go_out=. --go_opt=paths=source_relative \
//	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
//	observer.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: observer.proto

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

type ObserveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *ObserveRequest) Reset() {
	*x = ObserveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_observer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObserveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObserveRequest) ProtoMessage() {}

func (x *ObserveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_observer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObserveRequest.ProtoReflect.Descriptor instead.
func (*ObserveRequest) Descriptor() ([]byte, []int) {
	return file_observer_proto_rawDescGZIP(), []int{0}
}

func (x *ObserveRequest) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind    string         `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Service string         `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Client  string         `protobuf:"bytes,3,opt,name=client,proto3" json:"client,omitempty"`
	Type    string         `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Stats   *Stats         `protobuf:"bytes,5,opt,name=stats,proto3" json:"stats,omitempty"`
	Status  *ServiceStatus `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_observer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_observer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_observer_proto_rawDescGZIP(), []int{1}
}

func (x *Event) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Event) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *Event) GetClient() string {
	if x != nil {
		return x.Client
	}
	return ""
}

func (x *Event) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Event) GetStats() *Stats {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *Event) GetStatus() *ServiceStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type ServiceStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State string `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Msg   string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ServiceStatus) Reset() {
	*x = ServiceStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_observer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceStatus) ProtoMessage() {}

func (x *ServiceStatus) ProtoReflect() protoreflect.Message {
	mi := &file_observer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceStatus.ProtoReflect.Descriptor instead.
func (*ServiceStatus) Descriptor() ([]byte, []int) {
	return file_observer_proto_rawDescGZIP(), []int{2}
}

func (x *ServiceStatus) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *ServiceStatus) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type Stats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalConns   uint64 `protobuf:"varint,1,opt,name=totalConns,proto3" json:"totalConns,omitempty"`
	CurrentConns uint64 `protobuf:"varint,2,opt,name=currentConns,proto3" json:"currentConns,omitempty"`
	InputBytes   uint64 `protobuf:"varint,3,opt,name=inputBytes,proto3" json:"inputBytes,omitempty"`
	OutputBytes  uint64 `protobuf:"varint,4,opt,name=outputBytes,proto3" json:"outputBytes,omitempty"`
	TotalErrs    uint64 `protobuf:"varint,5,opt,name=totalErrs,proto3" json:"totalErrs,omitempty"`
}

func (x *Stats) Reset() {
	*x = Stats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_observer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stats) ProtoMessage() {}

func (x *Stats) ProtoReflect() protoreflect.Message {
	mi := &file_observer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stats.ProtoReflect.Descriptor instead.
func (*Stats) Descriptor() ([]byte, []int) {
	return file_observer_proto_rawDescGZIP(), []int{3}
}

func (x *Stats) GetTotalConns() uint64 {
	if x != nil {
		return x.TotalConns
	}
	return 0
}

func (x *Stats) GetCurrentConns() uint64 {
	if x != nil {
		return x.CurrentConns
	}
	return 0
}

func (x *Stats) GetInputBytes() uint64 {
	if x != nil {
		return x.InputBytes
	}
	return 0
}

func (x *Stats) GetOutputBytes() uint64 {
	if x != nil {
		return x.OutputBytes
	}
	return 0
}

func (x *Stats) GetTotalErrs() uint64 {
	if x != nil {
		return x.TotalErrs
	}
	return 0
}

type ObserveReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *ObserveReply) Reset() {
	*x = ObserveReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_observer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObserveReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObserveReply) ProtoMessage() {}

func (x *ObserveReply) ProtoReflect() protoreflect.Message {
	mi := &file_observer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObserveReply.ProtoReflect.Descriptor instead.
func (*ObserveReply) Descriptor() ([]byte, []int) {
	return file_observer_proto_rawDescGZIP(), []int{4}
}

func (x *ObserveReply) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_observer_proto protoreflect.FileDescriptor

var file_observer_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x0e, 0x4f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x06, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22,
	0xb3, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x37, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0xab,
	0x01, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x43, 0x6f, 0x6e, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x6e, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x73, 0x12, 0x1e, 0x0a, 0x0a,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0b, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x73, 0x22, 0x1e, 0x0a, 0x0c,
	0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x32, 0x41, 0x0a, 0x08,
	0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x07, 0x4f, 0x62, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x62, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42,
	0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2d, 0x67, 0x6f, 0x73, 0x74, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x6f, 0x62, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_observer_proto_rawDescOnce sync.Once
	file_observer_proto_rawDescData = file_observer_proto_rawDesc
)

func file_observer_proto_rawDescGZIP() []byte {
	file_observer_proto_rawDescOnce.Do(func() {
		file_observer_proto_rawDescData = protoimpl.X.CompressGZIP(file_observer_proto_rawDescData)
	})
	return file_observer_proto_rawDescData
}

var file_observer_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_observer_proto_goTypes = []interface{}{
	(*ObserveRequest)(nil), // 0: proto.ObserveRequest
	(*Event)(nil),          // 1: proto.Event
	(*ServiceStatus)(nil),  // 2: proto.ServiceStatus
	(*Stats)(nil),          // 3: proto.Stats
	(*ObserveReply)(nil),   // 4: proto.ObserveReply
}
var file_observer_proto_depIdxs = []int32{
	1, // 0: proto.ObserveRequest.events:type_name -> proto.Event
	3, // 1: proto.Event.stats:type_name -> proto.Stats
	2, // 2: proto.Event.status:type_name -> proto.ServiceStatus
	0, // 3: proto.Observer.Observe:input_type -> proto.ObserveRequest
	4, // 4: proto.Observer.Observe:output_type -> proto.ObserveReply
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_observer_proto_init() }
func file_observer_proto_init() {
	if File_observer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_observer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObserveRequest); i {
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
		file_observer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_observer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceStatus); i {
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
		file_observer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stats); i {
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
		file_observer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObserveReply); i {
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
			RawDescriptor: file_observer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_observer_proto_goTypes,
		DependencyIndexes: file_observer_proto_depIdxs,
		MessageInfos:      file_observer_proto_msgTypes,
	}.Build()
	File_observer_proto = out.File
	file_observer_proto_rawDesc = nil
	file_observer_proto_goTypes = nil
	file_observer_proto_depIdxs = nil
}
