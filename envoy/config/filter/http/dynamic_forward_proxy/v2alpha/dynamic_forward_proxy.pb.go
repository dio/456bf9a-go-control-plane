// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/dynamic_forward_proxy/v2alpha/dynamic_forward_proxy.proto

package envoy_config_filter_http_dynamic_forward_proxy_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v2alpha "github.com/envoyproxy/go-control-plane/envoy/config/common/dynamic_forward_proxy/v2alpha"
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

type FilterConfig struct {
	DnsCacheConfig       *v2alpha.DnsCacheConfig `protobuf:"bytes,1,opt,name=dns_cache_config,json=dnsCacheConfig,proto3" json:"dns_cache_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *FilterConfig) Reset()         { *m = FilterConfig{} }
func (m *FilterConfig) String() string { return proto.CompactTextString(m) }
func (*FilterConfig) ProtoMessage()    {}
func (*FilterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_85a2356b260c47da, []int{0}
}

func (m *FilterConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterConfig.Unmarshal(m, b)
}
func (m *FilterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterConfig.Marshal(b, m, deterministic)
}
func (m *FilterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterConfig.Merge(m, src)
}
func (m *FilterConfig) XXX_Size() int {
	return xxx_messageInfo_FilterConfig.Size(m)
}
func (m *FilterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FilterConfig proto.InternalMessageInfo

func (m *FilterConfig) GetDnsCacheConfig() *v2alpha.DnsCacheConfig {
	if m != nil {
		return m.DnsCacheConfig
	}
	return nil
}

type PerRouteConfig struct {
	// Types that are valid to be assigned to HostRewriteSpecifier:
	//	*PerRouteConfig_HostRewrite
	//	*PerRouteConfig_AutoHostRewriteHeader
	HostRewriteSpecifier isPerRouteConfig_HostRewriteSpecifier `protobuf_oneof:"host_rewrite_specifier"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *PerRouteConfig) Reset()         { *m = PerRouteConfig{} }
func (m *PerRouteConfig) String() string { return proto.CompactTextString(m) }
func (*PerRouteConfig) ProtoMessage()    {}
func (*PerRouteConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_85a2356b260c47da, []int{1}
}

func (m *PerRouteConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PerRouteConfig.Unmarshal(m, b)
}
func (m *PerRouteConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PerRouteConfig.Marshal(b, m, deterministic)
}
func (m *PerRouteConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PerRouteConfig.Merge(m, src)
}
func (m *PerRouteConfig) XXX_Size() int {
	return xxx_messageInfo_PerRouteConfig.Size(m)
}
func (m *PerRouteConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PerRouteConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PerRouteConfig proto.InternalMessageInfo

type isPerRouteConfig_HostRewriteSpecifier interface {
	isPerRouteConfig_HostRewriteSpecifier()
}

type PerRouteConfig_HostRewrite struct {
	HostRewrite string `protobuf:"bytes,1,opt,name=host_rewrite,json=hostRewrite,proto3,oneof"`
}

type PerRouteConfig_AutoHostRewriteHeader struct {
	AutoHostRewriteHeader string `protobuf:"bytes,2,opt,name=auto_host_rewrite_header,json=autoHostRewriteHeader,proto3,oneof"`
}

func (*PerRouteConfig_HostRewrite) isPerRouteConfig_HostRewriteSpecifier() {}

func (*PerRouteConfig_AutoHostRewriteHeader) isPerRouteConfig_HostRewriteSpecifier() {}

func (m *PerRouteConfig) GetHostRewriteSpecifier() isPerRouteConfig_HostRewriteSpecifier {
	if m != nil {
		return m.HostRewriteSpecifier
	}
	return nil
}

func (m *PerRouteConfig) GetHostRewrite() string {
	if x, ok := m.GetHostRewriteSpecifier().(*PerRouteConfig_HostRewrite); ok {
		return x.HostRewrite
	}
	return ""
}

func (m *PerRouteConfig) GetAutoHostRewriteHeader() string {
	if x, ok := m.GetHostRewriteSpecifier().(*PerRouteConfig_AutoHostRewriteHeader); ok {
		return x.AutoHostRewriteHeader
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PerRouteConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PerRouteConfig_HostRewrite)(nil),
		(*PerRouteConfig_AutoHostRewriteHeader)(nil),
	}
}

func init() {
	proto.RegisterType((*FilterConfig)(nil), "envoy.config.filter.http.dynamic_forward_proxy.v2alpha.FilterConfig")
	proto.RegisterType((*PerRouteConfig)(nil), "envoy.config.filter.http.dynamic_forward_proxy.v2alpha.PerRouteConfig")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/dynamic_forward_proxy/v2alpha/dynamic_forward_proxy.proto", fileDescriptor_85a2356b260c47da)
}

var fileDescriptor_85a2356b260c47da = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x86, 0x9b, 0x42, 0x2f, 0xf7, 0x4e, 0x4b, 0x29, 0xb9, 0xb6, 0x86, 0x2a, 0x22, 0x5d, 0xb9,
	0xca, 0x40, 0x05, 0xf7, 0x4d, 0x4b, 0xc9, 0x32, 0x64, 0xe1, 0x76, 0x98, 0x26, 0x93, 0x66, 0x20,
	0x99, 0x13, 0x26, 0x93, 0xd6, 0x3e, 0x80, 0x1b, 0x37, 0x6e, 0x7d, 0x21, 0x5f, 0xc8, 0xa5, 0x0b,
	0x91, 0xcc, 0xb4, 0x6a, 0xa0, 0x28, 0xb8, 0x0b, 0xe7, 0x3b, 0xf3, 0xcd, 0x3f, 0xe7, 0x04, 0x85,
	0x4c, 0x6c, 0x60, 0x87, 0x23, 0x10, 0x09, 0x5f, 0xe3, 0x84, 0x67, 0x8a, 0x49, 0x9c, 0x2a, 0x55,
	0xe0, 0x78, 0x27, 0x68, 0xce, 0x23, 0x92, 0x80, 0xdc, 0x52, 0x19, 0x93, 0x42, 0xc2, 0xdd, 0x0e,
	0x6f, 0xa6, 0x34, 0x2b, 0x52, 0x7a, 0x9c, 0xba, 0x85, 0x04, 0x05, 0xf6, 0x8d, 0x76, 0xba, 0xc6,
	0xe9, 0x1a, 0xa7, 0x5b, 0x3b, 0xdd, 0xe3, 0xa7, 0xf6, 0xce, 0xf1, 0xac, 0x91, 0x25, 0x82, 0x3c,
	0x07, 0xf1, 0x53, 0x0c, 0x51, 0x92, 0x88, 0x46, 0x29, 0x33, 0x57, 0x8f, 0x2f, 0xaa, 0xb8, 0xa0,
	0x98, 0x0a, 0x01, 0x8a, 0x2a, 0x0e, 0xa2, 0xc4, 0x39, 0x5f, 0x4b, 0xaa, 0x0e, 0xfc, 0x74, 0x43,
	0x33, 0x1e, 0x53, 0xc5, 0xf0, 0xe1, 0xc3, 0x80, 0xc9, 0xbd, 0x85, 0x7a, 0x4b, 0x9d, 0x74, 0xae,
	0xaf, 0xb7, 0x2b, 0x34, 0xf8, 0x90, 0x13, 0x13, 0xc9, 0xb1, 0x2e, 0xad, 0xab, 0xee, 0x74, 0xe6,
	0x36, 0xde, 0x67, 0x72, 0x7e, 0xff, 0x34, 0x77, 0x21, 0xca, 0x79, 0x6d, 0x32, 0x72, 0xef, 0xef,
	0xab, 0xd7, 0x79, 0xb0, 0xda, 0x03, 0x2b, 0xec, 0xc7, 0x0d, 0x32, 0x79, 0xb6, 0x50, 0x3f, 0x60,
	0x32, 0x84, 0x4a, 0xed, 0x4b, 0xf6, 0x0c, 0xf5, 0x52, 0x28, 0x15, 0x91, 0x6c, 0x2b, 0xb9, 0x62,
	0x3a, 0xc5, 0x3f, 0xef, 0xfc, 0xe5, 0xe9, 0xed, 0xb1, 0x33, 0x42, 0x27, 0x5f, 0x19, 0xc9, 0xb8,
	0x62, 0x92, 0x66, 0x7e, 0x2b, 0xec, 0xd6, 0xf5, 0xd0, 0x94, 0xed, 0x5b, 0xe4, 0xd0, 0x4a, 0x01,
	0x69, 0xf4, 0xa6, 0x8c, 0xc6, 0x4c, 0x3a, 0x6d, 0xad, 0x3b, 0xd3, 0xba, 0x21, 0xfa, 0x7f, 0xa4,
	0xc5, 0x6f, 0x85, 0xc3, 0xfa, 0xb8, 0xff, 0x69, 0xf4, 0x35, 0xf0, 0x1c, 0x34, 0x6a, 0xf4, 0x97,
	0x05, 0x8b, 0x78, 0xc2, 0x99, 0xf4, 0x56, 0x68, 0xc1, 0xc1, 0x0c, 0xca, 0x4c, 0xe2, 0x77, 0xff,
	0x84, 0xe7, 0x2c, 0x0c, 0x5e, 0x1a, 0x1a, 0xd4, 0x30, 0xa8, 0x37, 0x16, 0x58, 0xab, 0x3f, 0x7a,
	0x75, 0xd7, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x89, 0x86, 0x8f, 0xc4, 0x02, 0x00, 0x00,
}