// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/wasm/v3alpha/wasm.proto

package envoy_config_filter_network_wasm_v3alpha

import (
	fmt "fmt"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/config/wasm/v3alpha"
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

type Wasm struct {
	Config               *v3alpha.PluginConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Wasm) Reset()         { *m = Wasm{} }
func (m *Wasm) String() string { return proto.CompactTextString(m) }
func (*Wasm) ProtoMessage()    {}
func (*Wasm) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1945f9ba3a50209, []int{0}
}

func (m *Wasm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Wasm.Unmarshal(m, b)
}
func (m *Wasm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Wasm.Marshal(b, m, deterministic)
}
func (m *Wasm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Wasm.Merge(m, src)
}
func (m *Wasm) XXX_Size() int {
	return xxx_messageInfo_Wasm.Size(m)
}
func (m *Wasm) XXX_DiscardUnknown() {
	xxx_messageInfo_Wasm.DiscardUnknown(m)
}

var xxx_messageInfo_Wasm proto.InternalMessageInfo

func (m *Wasm) GetConfig() *v3alpha.PluginConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func init() {
	proto.RegisterType((*Wasm)(nil), "envoy.config.filter.network.wasm.v3alpha.Wasm")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/wasm/v3alpha/wasm.proto", fileDescriptor_d1945f9ba3a50209)
}

var fileDescriptor_d1945f9ba3a50209 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4e, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2,
	0xcf, 0x4b, 0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x2f, 0x4f, 0x2c, 0xce, 0xd5, 0x2f, 0x33, 0x4e,
	0xcc, 0x29, 0xc8, 0x48, 0x04, 0x73, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x34, 0xc0, 0x9a,
	0xf4, 0x20, 0x9a, 0xf4, 0x20, 0x9a, 0xf4, 0xa0, 0x9a, 0xf4, 0xc0, 0xea, 0xa0, 0x9a, 0xa4, 0xc4,
	0xcb, 0x12, 0x73, 0x32, 0x53, 0x12, 0x4b, 0x52, 0xf5, 0x61, 0x0c, 0x88, 0x11, 0x52, 0x2a, 0x28,
	0xf6, 0xe2, 0xb0, 0x48, 0xc9, 0x9d, 0x8b, 0x25, 0x3c, 0xb1, 0x38, 0x57, 0xc8, 0x9e, 0x8b, 0x0d,
	0xa2, 0x52, 0x82, 0x51, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x5d, 0x0f, 0xc5, 0x05, 0xc8, 0x56, 0xea,
	0x05, 0xe4, 0x94, 0xa6, 0x67, 0xe6, 0x39, 0x83, 0x65, 0x82, 0xa0, 0xda, 0x9c, 0x5c, 0xb9, 0xcc,
	0x32, 0xf3, 0x21, 0x9a, 0x0a, 0x8a, 0xf2, 0x2b, 0x2a, 0xf5, 0x88, 0xf5, 0x81, 0x13, 0x27, 0xc8,
	0x01, 0x01, 0x20, 0xd7, 0x04, 0x30, 0x26, 0xb1, 0x81, 0x9d, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0xf9, 0x25, 0x87, 0xd1, 0x36, 0x01, 0x00, 0x00,
}