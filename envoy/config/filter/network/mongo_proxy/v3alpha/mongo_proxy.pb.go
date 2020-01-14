// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/mongo_proxy/v3alpha/mongo_proxy.proto

package envoy_config_filter_network_mongo_proxy_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/config/filter/fault/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MongoProxy struct {
	StatPrefix           string              `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	AccessLog            string              `protobuf:"bytes,2,opt,name=access_log,json=accessLog,proto3" json:"access_log,omitempty"`
	Delay                *v3alpha.FaultDelay `protobuf:"bytes,3,opt,name=delay,proto3" json:"delay,omitempty"`
	EmitDynamicMetadata  bool                `protobuf:"varint,4,opt,name=emit_dynamic_metadata,json=emitDynamicMetadata,proto3" json:"emit_dynamic_metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *MongoProxy) Reset()         { *m = MongoProxy{} }
func (m *MongoProxy) String() string { return proto.CompactTextString(m) }
func (*MongoProxy) ProtoMessage()    {}
func (*MongoProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b8f056578a6890a, []int{0}
}

func (m *MongoProxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MongoProxy.Unmarshal(m, b)
}
func (m *MongoProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MongoProxy.Marshal(b, m, deterministic)
}
func (m *MongoProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MongoProxy.Merge(m, src)
}
func (m *MongoProxy) XXX_Size() int {
	return xxx_messageInfo_MongoProxy.Size(m)
}
func (m *MongoProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_MongoProxy.DiscardUnknown(m)
}

var xxx_messageInfo_MongoProxy proto.InternalMessageInfo

func (m *MongoProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *MongoProxy) GetAccessLog() string {
	if m != nil {
		return m.AccessLog
	}
	return ""
}

func (m *MongoProxy) GetDelay() *v3alpha.FaultDelay {
	if m != nil {
		return m.Delay
	}
	return nil
}

func (m *MongoProxy) GetEmitDynamicMetadata() bool {
	if m != nil {
		return m.EmitDynamicMetadata
	}
	return false
}

func init() {
	proto.RegisterType((*MongoProxy)(nil), "envoy.config.filter.network.mongo_proxy.v3alpha.MongoProxy")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/mongo_proxy/v3alpha/mongo_proxy.proto", fileDescriptor_2b8f056578a6890a)
}

var fileDescriptor_2b8f056578a6890a = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xc1, 0x4a, 0xf3, 0x40,
	0x10, 0xc7, 0x49, 0xbf, 0x7e, 0xdf, 0x67, 0xb7, 0x07, 0x21, 0x22, 0x86, 0x82, 0x12, 0x3d, 0xe5,
	0xd2, 0x5d, 0x68, 0x11, 0x41, 0xf4, 0x60, 0x2c, 0x9e, 0x2c, 0x84, 0xbc, 0x40, 0x18, 0x93, 0x4d,
	0x5c, 0x4c, 0x76, 0x42, 0xb2, 0x8d, 0xcd, 0x1b, 0x78, 0xf0, 0x09, 0x7c, 0x1f, 0x5f, 0xca, 0x93,
	0x6c, 0x36, 0xda, 0x1e, 0x7a, 0xd0, 0xdb, 0xce, 0xfc, 0x67, 0x7e, 0x33, 0xf3, 0x5f, 0x72, 0xc3,
	0x65, 0x83, 0x2d, 0x8b, 0x51, 0xa6, 0x22, 0x63, 0xa9, 0xc8, 0x15, 0xaf, 0x98, 0xe4, 0xea, 0x19,
	0xab, 0x27, 0x56, 0xa0, 0xcc, 0x30, 0x2a, 0x2b, 0x5c, 0xb7, 0xac, 0x99, 0x43, 0x5e, 0x3e, 0xc2,
	0x76, 0x8e, 0x96, 0x15, 0x2a, 0xb4, 0x59, 0x87, 0xa0, 0x06, 0x41, 0x0d, 0x82, 0xf6, 0x08, 0xba,
	0x5d, 0xde, 0x23, 0x26, 0xd3, 0x5d, 0x33, 0x53, 0x58, 0xe5, 0xea, 0x7b, 0x4a, 0x17, 0x19, 0xfe,
	0xe4, 0x74, 0x95, 0x94, 0xc0, 0x40, 0x4a, 0x54, 0xa0, 0x04, 0xca, 0x9a, 0x35, 0xbc, 0xaa, 0x05,
	0x4a, 0x21, 0xb3, 0xbe, 0xe4, 0xa8, 0x81, 0x5c, 0x24, 0xa0, 0x38, 0xfb, 0x7a, 0x18, 0xe1, 0xec,
	0x75, 0x40, 0xc8, 0x52, 0xaf, 0x10, 0xe8, 0x0d, 0x6c, 0x8f, 0x8c, 0x6b, 0x05, 0x2a, 0x2a, 0x2b,
	0x9e, 0x8a, 0xb5, 0x63, 0xb9, 0x96, 0x37, 0xf2, 0xff, 0x7f, 0xf8, 0xc3, 0x6a, 0xe0, 0x5a, 0x21,
	0xd1, 0x5a, 0xd0, 0x49, 0xf6, 0x31, 0x21, 0x10, 0xc7, 0xbc, 0xae, 0xa3, 0x1c, 0x33, 0x67, 0xa0,
	0x0b, 0xc3, 0x91, 0xc9, 0xdc, 0x63, 0x66, 0xdf, 0x92, 0xbf, 0x09, 0xcf, 0xa1, 0x75, 0xfe, 0xb8,
	0x96, 0x37, 0x9e, 0x4d, 0xe9, 0x2e, 0x0f, 0xcc, 0x11, 0xfd, 0x49, 0xf4, 0x4e, 0x47, 0x0b, 0xdd,
	0x14, 0x9a, 0x5e, 0x7b, 0x46, 0x0e, 0x79, 0x21, 0x54, 0x94, 0xb4, 0x12, 0x0a, 0x11, 0x47, 0x05,
	0x57, 0x90, 0x80, 0x02, 0x67, 0xe8, 0x5a, 0xde, 0x5e, 0x78, 0xa0, 0xc5, 0x85, 0xd1, 0x96, 0xbd,
	0x74, 0x79, 0xf5, 0xf6, 0xfe, 0x72, 0x72, 0x41, 0xce, 0x7f, 0xec, 0xf9, 0x8c, 0x6e, 0xee, 0xf7,
	0x43, 0x72, 0x2d, 0xd0, 0xec, 0x6a, 0xe4, 0x5f, 0x7e, 0x9d, 0xbf, 0xbf, 0x81, 0x05, 0xda, 0xe0,
	0xc0, 0x7a, 0xf8, 0xd7, 0x39, 0x3d, 0xff, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x33, 0xf6, 0xa3, 0xa6,
	0x4a, 0x02, 0x00, 0x00,
}