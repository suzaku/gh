// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: codespace/codespace_host_service.v1.proto

package codespace

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

type NotifyCodespaceOfClientActivityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId         string   `protobuf:"bytes,1,opt,name=ClientId,proto3" json:"ClientId,omitempty"`
	ClientActivities []string `protobuf:"bytes,2,rep,name=ClientActivities,proto3" json:"ClientActivities,omitempty"`
}

func (x *NotifyCodespaceOfClientActivityRequest) Reset() {
	*x = NotifyCodespaceOfClientActivityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codespace_codespace_host_service_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyCodespaceOfClientActivityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyCodespaceOfClientActivityRequest) ProtoMessage() {}

func (x *NotifyCodespaceOfClientActivityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_codespace_codespace_host_service_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyCodespaceOfClientActivityRequest.ProtoReflect.Descriptor instead.
func (*NotifyCodespaceOfClientActivityRequest) Descriptor() ([]byte, []int) {
	return file_codespace_codespace_host_service_v1_proto_rawDescGZIP(), []int{0}
}

func (x *NotifyCodespaceOfClientActivityRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *NotifyCodespaceOfClientActivityRequest) GetClientActivities() []string {
	if x != nil {
		return x.ClientActivities
	}
	return nil
}

type NotifyCodespaceOfClientActivityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result  bool   `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *NotifyCodespaceOfClientActivityResponse) Reset() {
	*x = NotifyCodespaceOfClientActivityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codespace_codespace_host_service_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyCodespaceOfClientActivityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyCodespaceOfClientActivityResponse) ProtoMessage() {}

func (x *NotifyCodespaceOfClientActivityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_codespace_codespace_host_service_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyCodespaceOfClientActivityResponse.ProtoReflect.Descriptor instead.
func (*NotifyCodespaceOfClientActivityResponse) Descriptor() ([]byte, []int) {
	return file_codespace_codespace_host_service_v1_proto_rawDescGZIP(), []int{1}
}

func (x *NotifyCodespaceOfClientActivityResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *NotifyCodespaceOfClientActivityResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type RebuildContainerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Incremental *bool `protobuf:"varint,1,opt,name=Incremental,proto3,oneof" json:"Incremental,omitempty"`
}

func (x *RebuildContainerRequest) Reset() {
	*x = RebuildContainerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codespace_codespace_host_service_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RebuildContainerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RebuildContainerRequest) ProtoMessage() {}

func (x *RebuildContainerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_codespace_codespace_host_service_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RebuildContainerRequest.ProtoReflect.Descriptor instead.
func (*RebuildContainerRequest) Descriptor() ([]byte, []int) {
	return file_codespace_codespace_host_service_v1_proto_rawDescGZIP(), []int{2}
}

func (x *RebuildContainerRequest) GetIncremental() bool {
	if x != nil && x.Incremental != nil {
		return *x.Incremental
	}
	return false
}

type RebuildContainerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RebuildContainer bool `protobuf:"varint,1,opt,name=RebuildContainer,proto3" json:"RebuildContainer,omitempty"`
}

func (x *RebuildContainerResponse) Reset() {
	*x = RebuildContainerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codespace_codespace_host_service_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RebuildContainerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RebuildContainerResponse) ProtoMessage() {}

func (x *RebuildContainerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_codespace_codespace_host_service_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RebuildContainerResponse.ProtoReflect.Descriptor instead.
func (*RebuildContainerResponse) Descriptor() ([]byte, []int) {
	return file_codespace_codespace_host_service_v1_proto_rawDescGZIP(), []int{3}
}

func (x *RebuildContainerResponse) GetRebuildContainer() bool {
	if x != nil {
		return x.RebuildContainer
	}
	return false
}

var File_codespace_codespace_host_service_v1_proto protoreflect.FileDescriptor

var file_codespace_codespace_host_service_v1_proto_rawDesc = []byte{
	0x0a, 0x29, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x64, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x27, 0x43, 0x6f, 0x64,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x22, 0x70, 0x0a, 0x26, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x43, 0x6f,
	0x64, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4f, 0x66, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x5b, 0x0a, 0x27, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x43, 0x6f, 0x64, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4f, 0x66, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x50, 0x0a, 0x17, 0x52, 0x65, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25,
	0x0a, 0x0b, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0b, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x61, 0x6c, 0x22, 0x46, 0x0a, 0x18, 0x52, 0x65, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2a, 0x0a, 0x10, 0x52, 0x65, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x52, 0x65, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x32, 0xf5, 0x02,
	0x0a, 0x0d, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x12,
	0xc4, 0x01, 0x0a, 0x1f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x4f, 0x66, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x12, 0x4f, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73,
	0x2e, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x48,
	0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4f, 0x66, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x50, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x73, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x48, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x4f, 0x66,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x9c, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x41, 0x73, 0x79, 0x6e, 0x63,
	0x12, 0x40, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x72,
	0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x48, 0x6f, 0x73, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x41, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2e,
	0x47, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x48, 0x6f,
	0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_codespace_codespace_host_service_v1_proto_rawDescOnce sync.Once
	file_codespace_codespace_host_service_v1_proto_rawDescData = file_codespace_codespace_host_service_v1_proto_rawDesc
)

func file_codespace_codespace_host_service_v1_proto_rawDescGZIP() []byte {
	file_codespace_codespace_host_service_v1_proto_rawDescOnce.Do(func() {
		file_codespace_codespace_host_service_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_codespace_codespace_host_service_v1_proto_rawDescData)
	})
	return file_codespace_codespace_host_service_v1_proto_rawDescData
}

var file_codespace_codespace_host_service_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_codespace_codespace_host_service_v1_proto_goTypes = []interface{}{
	(*NotifyCodespaceOfClientActivityRequest)(nil),  // 0: Codespaces.Grpc.CodespaceHostService.v1.NotifyCodespaceOfClientActivityRequest
	(*NotifyCodespaceOfClientActivityResponse)(nil), // 1: Codespaces.Grpc.CodespaceHostService.v1.NotifyCodespaceOfClientActivityResponse
	(*RebuildContainerRequest)(nil),                 // 2: Codespaces.Grpc.CodespaceHostService.v1.RebuildContainerRequest
	(*RebuildContainerResponse)(nil),                // 3: Codespaces.Grpc.CodespaceHostService.v1.RebuildContainerResponse
}
var file_codespace_codespace_host_service_v1_proto_depIdxs = []int32{
	0, // 0: Codespaces.Grpc.CodespaceHostService.v1.CodespaceHost.NotifyCodespaceOfClientActivity:input_type -> Codespaces.Grpc.CodespaceHostService.v1.NotifyCodespaceOfClientActivityRequest
	2, // 1: Codespaces.Grpc.CodespaceHostService.v1.CodespaceHost.RebuildContainerAsync:input_type -> Codespaces.Grpc.CodespaceHostService.v1.RebuildContainerRequest
	1, // 2: Codespaces.Grpc.CodespaceHostService.v1.CodespaceHost.NotifyCodespaceOfClientActivity:output_type -> Codespaces.Grpc.CodespaceHostService.v1.NotifyCodespaceOfClientActivityResponse
	3, // 3: Codespaces.Grpc.CodespaceHostService.v1.CodespaceHost.RebuildContainerAsync:output_type -> Codespaces.Grpc.CodespaceHostService.v1.RebuildContainerResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_codespace_codespace_host_service_v1_proto_init() }
func file_codespace_codespace_host_service_v1_proto_init() {
	if File_codespace_codespace_host_service_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_codespace_codespace_host_service_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyCodespaceOfClientActivityRequest); i {
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
		file_codespace_codespace_host_service_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyCodespaceOfClientActivityResponse); i {
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
		file_codespace_codespace_host_service_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RebuildContainerRequest); i {
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
		file_codespace_codespace_host_service_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RebuildContainerResponse); i {
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
	file_codespace_codespace_host_service_v1_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_codespace_codespace_host_service_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_codespace_codespace_host_service_v1_proto_goTypes,
		DependencyIndexes: file_codespace_codespace_host_service_v1_proto_depIdxs,
		MessageInfos:      file_codespace_codespace_host_service_v1_proto_msgTypes,
	}.Build()
	File_codespace_codespace_host_service_v1_proto = out.File
	file_codespace_codespace_host_service_v1_proto_rawDesc = nil
	file_codespace_codespace_host_service_v1_proto_goTypes = nil
	file_codespace_codespace_host_service_v1_proto_depIdxs = nil
}
