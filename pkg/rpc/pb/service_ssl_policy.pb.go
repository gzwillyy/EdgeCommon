// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: service_ssl_policy.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// 创建Policy
type CreateSSLPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Http2Enabled      bool     `protobuf:"varint,1,opt,name=http2Enabled,proto3" json:"http2Enabled,omitempty"`
	MinVersion        string   `protobuf:"bytes,2,opt,name=minVersion,proto3" json:"minVersion,omitempty"`
	CertsJSON         []byte   `protobuf:"bytes,3,opt,name=certsJSON,proto3" json:"certsJSON,omitempty"`
	HstsJSON          []byte   `protobuf:"bytes,4,opt,name=hstsJSON,proto3" json:"hstsJSON,omitempty"`
	ClientAuthType    int32    `protobuf:"varint,5,opt,name=clientAuthType,proto3" json:"clientAuthType,omitempty"`
	ClientCACertsJSON []byte   `protobuf:"bytes,6,opt,name=clientCACertsJSON,proto3" json:"clientCACertsJSON,omitempty"`
	CipherSuites      []string `protobuf:"bytes,7,rep,name=cipherSuites,proto3" json:"cipherSuites,omitempty"`
	CipherSuitesIsOn  bool     `protobuf:"varint,8,opt,name=cipherSuitesIsOn,proto3" json:"cipherSuitesIsOn,omitempty"`
}

func (x *CreateSSLPolicyRequest) Reset() {
	*x = CreateSSLPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ssl_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSSLPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSSLPolicyRequest) ProtoMessage() {}

func (x *CreateSSLPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ssl_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSSLPolicyRequest.ProtoReflect.Descriptor instead.
func (*CreateSSLPolicyRequest) Descriptor() ([]byte, []int) {
	return file_service_ssl_policy_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSSLPolicyRequest) GetHttp2Enabled() bool {
	if x != nil {
		return x.Http2Enabled
	}
	return false
}

func (x *CreateSSLPolicyRequest) GetMinVersion() string {
	if x != nil {
		return x.MinVersion
	}
	return ""
}

func (x *CreateSSLPolicyRequest) GetCertsJSON() []byte {
	if x != nil {
		return x.CertsJSON
	}
	return nil
}

func (x *CreateSSLPolicyRequest) GetHstsJSON() []byte {
	if x != nil {
		return x.HstsJSON
	}
	return nil
}

func (x *CreateSSLPolicyRequest) GetClientAuthType() int32 {
	if x != nil {
		return x.ClientAuthType
	}
	return 0
}

func (x *CreateSSLPolicyRequest) GetClientCACertsJSON() []byte {
	if x != nil {
		return x.ClientCACertsJSON
	}
	return nil
}

func (x *CreateSSLPolicyRequest) GetCipherSuites() []string {
	if x != nil {
		return x.CipherSuites
	}
	return nil
}

func (x *CreateSSLPolicyRequest) GetCipherSuitesIsOn() bool {
	if x != nil {
		return x.CipherSuitesIsOn
	}
	return false
}

type CreateSSLPolicyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SslPolicyId int64 `protobuf:"varint,1,opt,name=sslPolicyId,proto3" json:"sslPolicyId,omitempty"`
}

func (x *CreateSSLPolicyResponse) Reset() {
	*x = CreateSSLPolicyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ssl_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSSLPolicyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSSLPolicyResponse) ProtoMessage() {}

func (x *CreateSSLPolicyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_ssl_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSSLPolicyResponse.ProtoReflect.Descriptor instead.
func (*CreateSSLPolicyResponse) Descriptor() ([]byte, []int) {
	return file_service_ssl_policy_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSSLPolicyResponse) GetSslPolicyId() int64 {
	if x != nil {
		return x.SslPolicyId
	}
	return 0
}

// 修改Policy
type UpdateSSLPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SslPolicyId       int64    `protobuf:"varint,1,opt,name=sslPolicyId,proto3" json:"sslPolicyId,omitempty"`
	Http2Enabled      bool     `protobuf:"varint,2,opt,name=http2Enabled,proto3" json:"http2Enabled,omitempty"`
	MinVersion        string   `protobuf:"bytes,3,opt,name=minVersion,proto3" json:"minVersion,omitempty"`
	CertsJSON         []byte   `protobuf:"bytes,4,opt,name=certsJSON,proto3" json:"certsJSON,omitempty"`
	HstsJSON          []byte   `protobuf:"bytes,5,opt,name=hstsJSON,proto3" json:"hstsJSON,omitempty"`
	ClientAuthType    int32    `protobuf:"varint,6,opt,name=clientAuthType,proto3" json:"clientAuthType,omitempty"`
	ClientCACertsJSON []byte   `protobuf:"bytes,7,opt,name=clientCACertsJSON,proto3" json:"clientCACertsJSON,omitempty"`
	CipherSuites      []string `protobuf:"bytes,8,rep,name=cipherSuites,proto3" json:"cipherSuites,omitempty"`
	CipherSuitesIsOn  bool     `protobuf:"varint,9,opt,name=cipherSuitesIsOn,proto3" json:"cipherSuitesIsOn,omitempty"`
}

func (x *UpdateSSLPolicyRequest) Reset() {
	*x = UpdateSSLPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ssl_policy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSSLPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSSLPolicyRequest) ProtoMessage() {}

func (x *UpdateSSLPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ssl_policy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSSLPolicyRequest.ProtoReflect.Descriptor instead.
func (*UpdateSSLPolicyRequest) Descriptor() ([]byte, []int) {
	return file_service_ssl_policy_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateSSLPolicyRequest) GetSslPolicyId() int64 {
	if x != nil {
		return x.SslPolicyId
	}
	return 0
}

func (x *UpdateSSLPolicyRequest) GetHttp2Enabled() bool {
	if x != nil {
		return x.Http2Enabled
	}
	return false
}

func (x *UpdateSSLPolicyRequest) GetMinVersion() string {
	if x != nil {
		return x.MinVersion
	}
	return ""
}

func (x *UpdateSSLPolicyRequest) GetCertsJSON() []byte {
	if x != nil {
		return x.CertsJSON
	}
	return nil
}

func (x *UpdateSSLPolicyRequest) GetHstsJSON() []byte {
	if x != nil {
		return x.HstsJSON
	}
	return nil
}

func (x *UpdateSSLPolicyRequest) GetClientAuthType() int32 {
	if x != nil {
		return x.ClientAuthType
	}
	return 0
}

func (x *UpdateSSLPolicyRequest) GetClientCACertsJSON() []byte {
	if x != nil {
		return x.ClientCACertsJSON
	}
	return nil
}

func (x *UpdateSSLPolicyRequest) GetCipherSuites() []string {
	if x != nil {
		return x.CipherSuites
	}
	return nil
}

func (x *UpdateSSLPolicyRequest) GetCipherSuitesIsOn() bool {
	if x != nil {
		return x.CipherSuitesIsOn
	}
	return false
}

// 查找Policy
type FindEnabledSSLPolicyConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SslPolicyId int64 `protobuf:"varint,1,opt,name=sslPolicyId,proto3" json:"sslPolicyId,omitempty"`
}

func (x *FindEnabledSSLPolicyConfigRequest) Reset() {
	*x = FindEnabledSSLPolicyConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ssl_policy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledSSLPolicyConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledSSLPolicyConfigRequest) ProtoMessage() {}

func (x *FindEnabledSSLPolicyConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_ssl_policy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledSSLPolicyConfigRequest.ProtoReflect.Descriptor instead.
func (*FindEnabledSSLPolicyConfigRequest) Descriptor() ([]byte, []int) {
	return file_service_ssl_policy_proto_rawDescGZIP(), []int{3}
}

func (x *FindEnabledSSLPolicyConfigRequest) GetSslPolicyId() int64 {
	if x != nil {
		return x.SslPolicyId
	}
	return 0
}

type FindEnabledSSLPolicyConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SslPolicyJSON []byte `protobuf:"bytes,1,opt,name=sslPolicyJSON,proto3" json:"sslPolicyJSON,omitempty"`
}

func (x *FindEnabledSSLPolicyConfigResponse) Reset() {
	*x = FindEnabledSSLPolicyConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_ssl_policy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEnabledSSLPolicyConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEnabledSSLPolicyConfigResponse) ProtoMessage() {}

func (x *FindEnabledSSLPolicyConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_ssl_policy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEnabledSSLPolicyConfigResponse.ProtoReflect.Descriptor instead.
func (*FindEnabledSSLPolicyConfigResponse) Descriptor() ([]byte, []int) {
	return file_service_ssl_policy_proto_rawDescGZIP(), []int{4}
}

func (x *FindEnabledSSLPolicyConfigResponse) GetSslPolicyJSON() []byte {
	if x != nil {
		return x.SslPolicyJSON
	}
	return nil
}

var File_service_ssl_policy_proto protoreflect.FileDescriptor

var file_service_ssl_policy_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x73, 0x6c, 0x5f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x12,
	0x72, 0x70, 0x63, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xbc, 0x02, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x53, 0x4c,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a,
	0x0c, 0x68, 0x74, 0x74, 0x70, 0x32, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x68, 0x74, 0x74, 0x70, 0x32, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x69, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x69, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x65, 0x72, 0x74, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x65, 0x72, 0x74, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x12,
	0x1a, 0x0a, 0x08, 0x68, 0x73, 0x74, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x68, 0x73, 0x74, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x26, 0x0a, 0x0e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x41, 0x43,
	0x65, 0x72, 0x74, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x11,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x41, 0x43, 0x65, 0x72, 0x74, 0x73, 0x4a, 0x53, 0x4f,
	0x4e, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x53, 0x75, 0x69, 0x74, 0x65,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x53,
	0x75, 0x69, 0x74, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x53,
	0x75, 0x69, 0x74, 0x65, 0x73, 0x49, 0x73, 0x4f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x10, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x53, 0x75, 0x69, 0x74, 0x65, 0x73, 0x49, 0x73, 0x4f,
	0x6e, 0x22, 0x3b, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x53, 0x4c, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x73, 0x73, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x73, 0x73, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x22, 0xde,
	0x02, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x53, 0x4c, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x73, 0x6c,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x73, 0x73, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x68,
	0x74, 0x74, 0x70, 0x32, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0c, 0x68, 0x74, 0x74, 0x70, 0x32, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x6d, 0x69, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x69, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x65, 0x72, 0x74, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x63, 0x65, 0x72, 0x74, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x1a, 0x0a,
	0x08, 0x68, 0x73, 0x74, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x08, 0x68, 0x73, 0x74, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x2c, 0x0a, 0x11, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x41, 0x43, 0x65, 0x72,
	0x74, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x43, 0x41, 0x43, 0x65, 0x72, 0x74, 0x73, 0x4a, 0x53, 0x4f, 0x4e, 0x12,
	0x22, 0x0a, 0x0c, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x53, 0x75, 0x69, 0x74, 0x65, 0x73, 0x18,
	0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x53, 0x75, 0x69,
	0x74, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x69, 0x70, 0x68, 0x65, 0x72, 0x53, 0x75, 0x69,
	0x74, 0x65, 0x73, 0x49, 0x73, 0x4f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x63,
	0x69, 0x70, 0x68, 0x65, 0x72, 0x53, 0x75, 0x69, 0x74, 0x65, 0x73, 0x49, 0x73, 0x4f, 0x6e, 0x22,
	0x45, 0x0a, 0x21, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x53, 0x53,
	0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x73, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x73, 0x6c, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x22, 0x4a, 0x0a, 0x22, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x53, 0x53, 0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d,
	0x73, 0x73, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4a, 0x53, 0x4f, 0x4e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0d, 0x73, 0x73, 0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4a, 0x53,
	0x4f, 0x4e, 0x32, 0x8a, 0x02, 0x0a, 0x10, 0x53, 0x53, 0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x53, 0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x53, 0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x53, 0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x53, 0x4c,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x53, 0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x50, 0x43, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x6b, 0x0a, 0x1a, 0x66, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x53, 0x53, 0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x25, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x53, 0x53, 0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x53, 0x53, 0x4c, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_ssl_policy_proto_rawDescOnce sync.Once
	file_service_ssl_policy_proto_rawDescData = file_service_ssl_policy_proto_rawDesc
)

func file_service_ssl_policy_proto_rawDescGZIP() []byte {
	file_service_ssl_policy_proto_rawDescOnce.Do(func() {
		file_service_ssl_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_ssl_policy_proto_rawDescData)
	})
	return file_service_ssl_policy_proto_rawDescData
}

var file_service_ssl_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_service_ssl_policy_proto_goTypes = []interface{}{
	(*CreateSSLPolicyRequest)(nil),             // 0: pb.CreateSSLPolicyRequest
	(*CreateSSLPolicyResponse)(nil),            // 1: pb.CreateSSLPolicyResponse
	(*UpdateSSLPolicyRequest)(nil),             // 2: pb.UpdateSSLPolicyRequest
	(*FindEnabledSSLPolicyConfigRequest)(nil),  // 3: pb.FindEnabledSSLPolicyConfigRequest
	(*FindEnabledSSLPolicyConfigResponse)(nil), // 4: pb.FindEnabledSSLPolicyConfigResponse
	(*RPCSuccess)(nil),                         // 5: pb.RPCSuccess
}
var file_service_ssl_policy_proto_depIdxs = []int32{
	0, // 0: pb.SSLPolicyService.createSSLPolicy:input_type -> pb.CreateSSLPolicyRequest
	2, // 1: pb.SSLPolicyService.updateSSLPolicy:input_type -> pb.UpdateSSLPolicyRequest
	3, // 2: pb.SSLPolicyService.findEnabledSSLPolicyConfig:input_type -> pb.FindEnabledSSLPolicyConfigRequest
	1, // 3: pb.SSLPolicyService.createSSLPolicy:output_type -> pb.CreateSSLPolicyResponse
	5, // 4: pb.SSLPolicyService.updateSSLPolicy:output_type -> pb.RPCSuccess
	4, // 5: pb.SSLPolicyService.findEnabledSSLPolicyConfig:output_type -> pb.FindEnabledSSLPolicyConfigResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_ssl_policy_proto_init() }
func file_service_ssl_policy_proto_init() {
	if File_service_ssl_policy_proto != nil {
		return
	}
	file_rpc_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_ssl_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSSLPolicyRequest); i {
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
		file_service_ssl_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSSLPolicyResponse); i {
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
		file_service_ssl_policy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSSLPolicyRequest); i {
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
		file_service_ssl_policy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledSSLPolicyConfigRequest); i {
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
		file_service_ssl_policy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEnabledSSLPolicyConfigResponse); i {
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
			RawDescriptor: file_service_ssl_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_ssl_policy_proto_goTypes,
		DependencyIndexes: file_service_ssl_policy_proto_depIdxs,
		MessageInfos:      file_service_ssl_policy_proto_msgTypes,
	}.Build()
	File_service_ssl_policy_proto = out.File
	file_service_ssl_policy_proto_rawDesc = nil
	file_service_ssl_policy_proto_goTypes = nil
	file_service_ssl_policy_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SSLPolicyServiceClient is the client API for SSLPolicyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SSLPolicyServiceClient interface {
	// 创建Policy
	CreateSSLPolicy(ctx context.Context, in *CreateSSLPolicyRequest, opts ...grpc.CallOption) (*CreateSSLPolicyResponse, error)
	// 修改Policy
	UpdateSSLPolicy(ctx context.Context, in *UpdateSSLPolicyRequest, opts ...grpc.CallOption) (*RPCSuccess, error)
	// 查找Policy
	FindEnabledSSLPolicyConfig(ctx context.Context, in *FindEnabledSSLPolicyConfigRequest, opts ...grpc.CallOption) (*FindEnabledSSLPolicyConfigResponse, error)
}

type sSLPolicyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSSLPolicyServiceClient(cc grpc.ClientConnInterface) SSLPolicyServiceClient {
	return &sSLPolicyServiceClient{cc}
}

func (c *sSLPolicyServiceClient) CreateSSLPolicy(ctx context.Context, in *CreateSSLPolicyRequest, opts ...grpc.CallOption) (*CreateSSLPolicyResponse, error) {
	out := new(CreateSSLPolicyResponse)
	err := c.cc.Invoke(ctx, "/pb.SSLPolicyService/createSSLPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sSLPolicyServiceClient) UpdateSSLPolicy(ctx context.Context, in *UpdateSSLPolicyRequest, opts ...grpc.CallOption) (*RPCSuccess, error) {
	out := new(RPCSuccess)
	err := c.cc.Invoke(ctx, "/pb.SSLPolicyService/updateSSLPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sSLPolicyServiceClient) FindEnabledSSLPolicyConfig(ctx context.Context, in *FindEnabledSSLPolicyConfigRequest, opts ...grpc.CallOption) (*FindEnabledSSLPolicyConfigResponse, error) {
	out := new(FindEnabledSSLPolicyConfigResponse)
	err := c.cc.Invoke(ctx, "/pb.SSLPolicyService/findEnabledSSLPolicyConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SSLPolicyServiceServer is the server API for SSLPolicyService service.
type SSLPolicyServiceServer interface {
	// 创建Policy
	CreateSSLPolicy(context.Context, *CreateSSLPolicyRequest) (*CreateSSLPolicyResponse, error)
	// 修改Policy
	UpdateSSLPolicy(context.Context, *UpdateSSLPolicyRequest) (*RPCSuccess, error)
	// 查找Policy
	FindEnabledSSLPolicyConfig(context.Context, *FindEnabledSSLPolicyConfigRequest) (*FindEnabledSSLPolicyConfigResponse, error)
}

// UnimplementedSSLPolicyServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSSLPolicyServiceServer struct {
}

func (*UnimplementedSSLPolicyServiceServer) CreateSSLPolicy(context.Context, *CreateSSLPolicyRequest) (*CreateSSLPolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSSLPolicy not implemented")
}
func (*UnimplementedSSLPolicyServiceServer) UpdateSSLPolicy(context.Context, *UpdateSSLPolicyRequest) (*RPCSuccess, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSSLPolicy not implemented")
}
func (*UnimplementedSSLPolicyServiceServer) FindEnabledSSLPolicyConfig(context.Context, *FindEnabledSSLPolicyConfigRequest) (*FindEnabledSSLPolicyConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindEnabledSSLPolicyConfig not implemented")
}

func RegisterSSLPolicyServiceServer(s *grpc.Server, srv SSLPolicyServiceServer) {
	s.RegisterService(&_SSLPolicyService_serviceDesc, srv)
}

func _SSLPolicyService_CreateSSLPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSSLPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSLPolicyServiceServer).CreateSSLPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SSLPolicyService/CreateSSLPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSLPolicyServiceServer).CreateSSLPolicy(ctx, req.(*CreateSSLPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SSLPolicyService_UpdateSSLPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSSLPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSLPolicyServiceServer).UpdateSSLPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SSLPolicyService/UpdateSSLPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSLPolicyServiceServer).UpdateSSLPolicy(ctx, req.(*UpdateSSLPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SSLPolicyService_FindEnabledSSLPolicyConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindEnabledSSLPolicyConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSLPolicyServiceServer).FindEnabledSSLPolicyConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SSLPolicyService/FindEnabledSSLPolicyConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSLPolicyServiceServer).FindEnabledSSLPolicyConfig(ctx, req.(*FindEnabledSSLPolicyConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SSLPolicyService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.SSLPolicyService",
	HandlerType: (*SSLPolicyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createSSLPolicy",
			Handler:    _SSLPolicyService_CreateSSLPolicy_Handler,
		},
		{
			MethodName: "updateSSLPolicy",
			Handler:    _SSLPolicyService_UpdateSSLPolicy_Handler,
		},
		{
			MethodName: "findEnabledSSLPolicyConfig",
			Handler:    _SSLPolicyService_FindEnabledSSLPolicyConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_ssl_policy.proto",
}
