// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: proxy/shadowsocks/simplified/config.proto

package simplified

import (
	net "github.com/v2fly/v2ray-core/v5/common/net"
	packetaddr "github.com/v2fly/v2ray-core/v5/common/net/packetaddr"
	_ "github.com/v2fly/v2ray-core/v5/common/protoext"
	shadowsocks "github.com/v2fly/v2ray-core/v5/proxy/shadowsocks"
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

type ServerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method         *CipherTypeWrapper        `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Password       string                    `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Networks       *net.NetworkList          `protobuf:"bytes,3,opt,name=networks,proto3" json:"networks,omitempty"`
	PacketEncoding packetaddr.PacketAddrType `protobuf:"varint,4,opt,name=packet_encoding,json=packetEncoding,proto3,enum=v2ray.core.net.packetaddr.PacketAddrType" json:"packet_encoding,omitempty"`
}

func (x *ServerConfig) Reset() {
	*x = ServerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_shadowsocks_simplified_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerConfig) ProtoMessage() {}

func (x *ServerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_shadowsocks_simplified_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerConfig.ProtoReflect.Descriptor instead.
func (*ServerConfig) Descriptor() ([]byte, []int) {
	return file_proxy_shadowsocks_simplified_config_proto_rawDescGZIP(), []int{0}
}

func (x *ServerConfig) GetMethod() *CipherTypeWrapper {
	if x != nil {
		return x.Method
	}
	return nil
}

func (x *ServerConfig) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ServerConfig) GetNetworks() *net.NetworkList {
	if x != nil {
		return x.Networks
	}
	return nil
}

func (x *ServerConfig) GetPacketEncoding() packetaddr.PacketAddrType {
	if x != nil {
		return x.PacketEncoding
	}
	return packetaddr.PacketAddrType(0)
}

type ClientConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address                        *net.IPOrDomain    `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Port                           uint32             `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Method                         *CipherTypeWrapper `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	Password                       string             `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	ExperimentReducedIvHeadEntropy bool               `protobuf:"varint,90001,opt,name=experiment_reduced_iv_head_entropy,json=experimentReducedIvHeadEntropy,proto3" json:"experiment_reduced_iv_head_entropy,omitempty"`
}

func (x *ClientConfig) Reset() {
	*x = ClientConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_shadowsocks_simplified_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConfig) ProtoMessage() {}

func (x *ClientConfig) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_shadowsocks_simplified_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConfig.ProtoReflect.Descriptor instead.
func (*ClientConfig) Descriptor() ([]byte, []int) {
	return file_proxy_shadowsocks_simplified_config_proto_rawDescGZIP(), []int{1}
}

func (x *ClientConfig) GetAddress() *net.IPOrDomain {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *ClientConfig) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *ClientConfig) GetMethod() *CipherTypeWrapper {
	if x != nil {
		return x.Method
	}
	return nil
}

func (x *ClientConfig) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ClientConfig) GetExperimentReducedIvHeadEntropy() bool {
	if x != nil {
		return x.ExperimentReducedIvHeadEntropy
	}
	return false
}

type CipherTypeWrapper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value shadowsocks.CipherType `protobuf:"varint,1,opt,name=value,proto3,enum=v2ray.core.proxy.shadowsocks.CipherType" json:"value,omitempty"`
}

func (x *CipherTypeWrapper) Reset() {
	*x = CipherTypeWrapper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proxy_shadowsocks_simplified_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CipherTypeWrapper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CipherTypeWrapper) ProtoMessage() {}

func (x *CipherTypeWrapper) ProtoReflect() protoreflect.Message {
	mi := &file_proxy_shadowsocks_simplified_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CipherTypeWrapper.ProtoReflect.Descriptor instead.
func (*CipherTypeWrapper) Descriptor() ([]byte, []int) {
	return file_proxy_shadowsocks_simplified_config_proto_rawDescGZIP(), []int{2}
}

func (x *CipherTypeWrapper) GetValue() shadowsocks.CipherType {
	if x != nil {
		return x.Value
	}
	return shadowsocks.CipherType(0)
}

var File_proxy_shadowsocks_simplified_config_proto protoreflect.FileDescriptor

var file_proxy_shadowsocks_simplified_config_proto_rawDesc = []byte{
	0x0a, 0x29, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x73, 0x6f,
	0x63, 0x6b, 0x73, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x27, 0x76, 0x32, 0x72,
	0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x73, 0x68,
	0x61, 0x64, 0x6f, 0x77, 0x73, 0x6f, 0x63, 0x6b, 0x73, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x1a, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x65, 0x78, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6e,
	0x65, 0x74, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6e, 0x65, 0x74, 0x2f, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x6e, 0x65, 0x74, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x61, 0x64, 0x64,
	0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x73, 0x6f, 0x63, 0x6b,
	0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae,
	0x02, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x52, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x3a, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x73, 0x6f, 0x63, 0x6b, 0x73, 0x2e, 0x73,
	0x69, 0x6d, 0x70, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x2e, 0x43, 0x69, 0x70, 0x68, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x52, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x3e, 0x0a, 0x08, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x08, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x12,
	0x52, 0x0a, 0x0f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69,
	0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x61, 0x64, 0x64, 0x72, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0e, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x45, 0x6e, 0x63, 0x6f, 0x64,
	0x69, 0x6e, 0x67, 0x3a, 0x1a, 0x82, 0xb5, 0x18, 0x16, 0x0a, 0x07, 0x69, 0x6e, 0x62, 0x6f, 0x75,
	0x6e, 0x64, 0x12, 0x0b, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x73, 0x6f, 0x63, 0x6b, 0x73, 0x22,
	0xba, 0x02, 0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x3b, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x49, 0x50, 0x4f, 0x72, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x52, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x3a, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x73, 0x6f, 0x63, 0x6b, 0x73,
	0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x2e, 0x43, 0x69, 0x70, 0x68,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x52, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x4c, 0x0a, 0x22, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x64, 0x5f, 0x69, 0x76, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x5f,
	0x65, 0x6e, 0x74, 0x72, 0x6f, 0x70, 0x79, 0x18, 0x91, 0xbf, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x1e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x64, 0x75, 0x63,
	0x65, 0x64, 0x49, 0x76, 0x48, 0x65, 0x61, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x6f, 0x70, 0x79, 0x3a,
	0x1b, 0x82, 0xb5, 0x18, 0x17, 0x0a, 0x08, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x12,
	0x0b, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x73, 0x6f, 0x63, 0x6b, 0x73, 0x22, 0x53, 0x0a, 0x11,
	0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x12, 0x3e, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x28, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x73, 0x6f, 0x63, 0x6b, 0x73, 0x2e,
	0x43, 0x69, 0x70, 0x68, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x96, 0x01, 0x0a, 0x2b, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x73, 0x68, 0x61, 0x64, 0x6f,
	0x77, 0x73, 0x6f, 0x63, 0x6b, 0x73, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x76, 0x32, 0x66, 0x6c, 0x79, 0x2f, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x76, 0x35, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x73, 0x68, 0x61, 0x64, 0x6f, 0x77,
	0x73, 0x6f, 0x63, 0x6b, 0x73, 0x2f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64,
	0xaa, 0x02, 0x27, 0x56, 0x32, 0x52, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x53, 0x68, 0x61, 0x64, 0x6f, 0x77, 0x73, 0x6f, 0x63, 0x6b, 0x73, 0x2e,
	0x53, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proxy_shadowsocks_simplified_config_proto_rawDescOnce sync.Once
	file_proxy_shadowsocks_simplified_config_proto_rawDescData = file_proxy_shadowsocks_simplified_config_proto_rawDesc
)

func file_proxy_shadowsocks_simplified_config_proto_rawDescGZIP() []byte {
	file_proxy_shadowsocks_simplified_config_proto_rawDescOnce.Do(func() {
		file_proxy_shadowsocks_simplified_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_proxy_shadowsocks_simplified_config_proto_rawDescData)
	})
	return file_proxy_shadowsocks_simplified_config_proto_rawDescData
}

var file_proxy_shadowsocks_simplified_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proxy_shadowsocks_simplified_config_proto_goTypes = []interface{}{
	(*ServerConfig)(nil),           // 0: v2ray.core.proxy.shadowsocks.simplified.ServerConfig
	(*ClientConfig)(nil),           // 1: v2ray.core.proxy.shadowsocks.simplified.ClientConfig
	(*CipherTypeWrapper)(nil),      // 2: v2ray.core.proxy.shadowsocks.simplified.CipherTypeWrapper
	(*net.NetworkList)(nil),        // 3: v2ray.core.common.net.NetworkList
	(packetaddr.PacketAddrType)(0), // 4: v2ray.core.net.packetaddr.PacketAddrType
	(*net.IPOrDomain)(nil),         // 5: v2ray.core.common.net.IPOrDomain
	(shadowsocks.CipherType)(0),    // 6: v2ray.core.proxy.shadowsocks.CipherType
}
var file_proxy_shadowsocks_simplified_config_proto_depIdxs = []int32{
	2, // 0: v2ray.core.proxy.shadowsocks.simplified.ServerConfig.method:type_name -> v2ray.core.proxy.shadowsocks.simplified.CipherTypeWrapper
	3, // 1: v2ray.core.proxy.shadowsocks.simplified.ServerConfig.networks:type_name -> v2ray.core.common.net.NetworkList
	4, // 2: v2ray.core.proxy.shadowsocks.simplified.ServerConfig.packet_encoding:type_name -> v2ray.core.net.packetaddr.PacketAddrType
	5, // 3: v2ray.core.proxy.shadowsocks.simplified.ClientConfig.address:type_name -> v2ray.core.common.net.IPOrDomain
	2, // 4: v2ray.core.proxy.shadowsocks.simplified.ClientConfig.method:type_name -> v2ray.core.proxy.shadowsocks.simplified.CipherTypeWrapper
	6, // 5: v2ray.core.proxy.shadowsocks.simplified.CipherTypeWrapper.value:type_name -> v2ray.core.proxy.shadowsocks.CipherType
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proxy_shadowsocks_simplified_config_proto_init() }
func file_proxy_shadowsocks_simplified_config_proto_init() {
	if File_proxy_shadowsocks_simplified_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proxy_shadowsocks_simplified_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerConfig); i {
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
		file_proxy_shadowsocks_simplified_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientConfig); i {
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
		file_proxy_shadowsocks_simplified_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CipherTypeWrapper); i {
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
			RawDescriptor: file_proxy_shadowsocks_simplified_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proxy_shadowsocks_simplified_config_proto_goTypes,
		DependencyIndexes: file_proxy_shadowsocks_simplified_config_proto_depIdxs,
		MessageInfos:      file_proxy_shadowsocks_simplified_config_proto_msgTypes,
	}.Build()
	File_proxy_shadowsocks_simplified_config_proto = out.File
	file_proxy_shadowsocks_simplified_config_proto_rawDesc = nil
	file_proxy_shadowsocks_simplified_config_proto_goTypes = nil
	file_proxy_shadowsocks_simplified_config_proto_depIdxs = nil
}