// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v3alpha/listener/listener.proto

package envoy_api_v3alpha_listener

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/core"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_ "github.com/golang/protobuf/ptypes/struct"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type FilterChainMatch_ConnectionSourceType int32

const (
	FilterChainMatch_ANY      FilterChainMatch_ConnectionSourceType = 0
	FilterChainMatch_LOCAL    FilterChainMatch_ConnectionSourceType = 1
	FilterChainMatch_EXTERNAL FilterChainMatch_ConnectionSourceType = 2
)

var FilterChainMatch_ConnectionSourceType_name = map[int32]string{
	0: "ANY",
	1: "LOCAL",
	2: "EXTERNAL",
}

var FilterChainMatch_ConnectionSourceType_value = map[string]int32{
	"ANY":      0,
	"LOCAL":    1,
	"EXTERNAL": 2,
}

func (x FilterChainMatch_ConnectionSourceType) String() string {
	return proto.EnumName(FilterChainMatch_ConnectionSourceType_name, int32(x))
}

func (FilterChainMatch_ConnectionSourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_907f943165b7183f, []int{1, 0}
}

type Filter struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to ConfigType:
	//	*Filter_TypedConfig
	ConfigType           isFilter_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_907f943165b7183f, []int{0}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isFilter_ConfigType interface {
	isFilter_ConfigType()
}

type Filter_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,4,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*Filter_TypedConfig) isFilter_ConfigType() {}

func (m *Filter) GetConfigType() isFilter_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *Filter) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*Filter_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Filter) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Filter_TypedConfig)(nil),
	}
}

type FilterChainMatch struct {
	DestinationPort      *wrappers.UInt32Value                 `protobuf:"bytes,8,opt,name=destination_port,json=destinationPort,proto3" json:"destination_port,omitempty"`
	PrefixRanges         []*core.CidrRange                     `protobuf:"bytes,3,rep,name=prefix_ranges,json=prefixRanges,proto3" json:"prefix_ranges,omitempty"`
	AddressSuffix        string                                `protobuf:"bytes,4,opt,name=address_suffix,json=addressSuffix,proto3" json:"address_suffix,omitempty"`
	SuffixLen            *wrappers.UInt32Value                 `protobuf:"bytes,5,opt,name=suffix_len,json=suffixLen,proto3" json:"suffix_len,omitempty"`
	SourceType           FilterChainMatch_ConnectionSourceType `protobuf:"varint,12,opt,name=source_type,json=sourceType,proto3,enum=envoy.api.v3alpha.listener.FilterChainMatch_ConnectionSourceType" json:"source_type,omitempty"`
	SourcePrefixRanges   []*core.CidrRange                     `protobuf:"bytes,6,rep,name=source_prefix_ranges,json=sourcePrefixRanges,proto3" json:"source_prefix_ranges,omitempty"`
	SourcePorts          []uint32                              `protobuf:"varint,7,rep,packed,name=source_ports,json=sourcePorts,proto3" json:"source_ports,omitempty"`
	ServerNames          []string                              `protobuf:"bytes,11,rep,name=server_names,json=serverNames,proto3" json:"server_names,omitempty"`
	TransportProtocol    string                                `protobuf:"bytes,9,opt,name=transport_protocol,json=transportProtocol,proto3" json:"transport_protocol,omitempty"`
	ApplicationProtocols []string                              `protobuf:"bytes,10,rep,name=application_protocols,json=applicationProtocols,proto3" json:"application_protocols,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *FilterChainMatch) Reset()         { *m = FilterChainMatch{} }
func (m *FilterChainMatch) String() string { return proto.CompactTextString(m) }
func (*FilterChainMatch) ProtoMessage()    {}
func (*FilterChainMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_907f943165b7183f, []int{1}
}

func (m *FilterChainMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterChainMatch.Unmarshal(m, b)
}
func (m *FilterChainMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterChainMatch.Marshal(b, m, deterministic)
}
func (m *FilterChainMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterChainMatch.Merge(m, src)
}
func (m *FilterChainMatch) XXX_Size() int {
	return xxx_messageInfo_FilterChainMatch.Size(m)
}
func (m *FilterChainMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterChainMatch.DiscardUnknown(m)
}

var xxx_messageInfo_FilterChainMatch proto.InternalMessageInfo

func (m *FilterChainMatch) GetDestinationPort() *wrappers.UInt32Value {
	if m != nil {
		return m.DestinationPort
	}
	return nil
}

func (m *FilterChainMatch) GetPrefixRanges() []*core.CidrRange {
	if m != nil {
		return m.PrefixRanges
	}
	return nil
}

func (m *FilterChainMatch) GetAddressSuffix() string {
	if m != nil {
		return m.AddressSuffix
	}
	return ""
}

func (m *FilterChainMatch) GetSuffixLen() *wrappers.UInt32Value {
	if m != nil {
		return m.SuffixLen
	}
	return nil
}

func (m *FilterChainMatch) GetSourceType() FilterChainMatch_ConnectionSourceType {
	if m != nil {
		return m.SourceType
	}
	return FilterChainMatch_ANY
}

func (m *FilterChainMatch) GetSourcePrefixRanges() []*core.CidrRange {
	if m != nil {
		return m.SourcePrefixRanges
	}
	return nil
}

func (m *FilterChainMatch) GetSourcePorts() []uint32 {
	if m != nil {
		return m.SourcePorts
	}
	return nil
}

func (m *FilterChainMatch) GetServerNames() []string {
	if m != nil {
		return m.ServerNames
	}
	return nil
}

func (m *FilterChainMatch) GetTransportProtocol() string {
	if m != nil {
		return m.TransportProtocol
	}
	return ""
}

func (m *FilterChainMatch) GetApplicationProtocols() []string {
	if m != nil {
		return m.ApplicationProtocols
	}
	return nil
}

type FilterChain struct {
	FilterChainMatch     *FilterChainMatch     `protobuf:"bytes,1,opt,name=filter_chain_match,json=filterChainMatch,proto3" json:"filter_chain_match,omitempty"`
	Filters              []*Filter             `protobuf:"bytes,3,rep,name=filters,proto3" json:"filters,omitempty"`
	UseProxyProto        *wrappers.BoolValue   `protobuf:"bytes,4,opt,name=use_proxy_proto,json=useProxyProto,proto3" json:"use_proxy_proto,omitempty"`
	Metadata             *core.Metadata        `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	TransportSocket      *core.TransportSocket `protobuf:"bytes,6,opt,name=transport_socket,json=transportSocket,proto3" json:"transport_socket,omitempty"`
	Name                 string                `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *FilterChain) Reset()         { *m = FilterChain{} }
func (m *FilterChain) String() string { return proto.CompactTextString(m) }
func (*FilterChain) ProtoMessage()    {}
func (*FilterChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_907f943165b7183f, []int{2}
}

func (m *FilterChain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterChain.Unmarshal(m, b)
}
func (m *FilterChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterChain.Marshal(b, m, deterministic)
}
func (m *FilterChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterChain.Merge(m, src)
}
func (m *FilterChain) XXX_Size() int {
	return xxx_messageInfo_FilterChain.Size(m)
}
func (m *FilterChain) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterChain.DiscardUnknown(m)
}

var xxx_messageInfo_FilterChain proto.InternalMessageInfo

func (m *FilterChain) GetFilterChainMatch() *FilterChainMatch {
	if m != nil {
		return m.FilterChainMatch
	}
	return nil
}

func (m *FilterChain) GetFilters() []*Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *FilterChain) GetUseProxyProto() *wrappers.BoolValue {
	if m != nil {
		return m.UseProxyProto
	}
	return nil
}

func (m *FilterChain) GetMetadata() *core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *FilterChain) GetTransportSocket() *core.TransportSocket {
	if m != nil {
		return m.TransportSocket
	}
	return nil
}

func (m *FilterChain) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListenerFilter struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to ConfigType:
	//	*ListenerFilter_TypedConfig
	ConfigType           isListenerFilter_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ListenerFilter) Reset()         { *m = ListenerFilter{} }
func (m *ListenerFilter) String() string { return proto.CompactTextString(m) }
func (*ListenerFilter) ProtoMessage()    {}
func (*ListenerFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_907f943165b7183f, []int{3}
}

func (m *ListenerFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenerFilter.Unmarshal(m, b)
}
func (m *ListenerFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenerFilter.Marshal(b, m, deterministic)
}
func (m *ListenerFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenerFilter.Merge(m, src)
}
func (m *ListenerFilter) XXX_Size() int {
	return xxx_messageInfo_ListenerFilter.Size(m)
}
func (m *ListenerFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenerFilter.DiscardUnknown(m)
}

var xxx_messageInfo_ListenerFilter proto.InternalMessageInfo

func (m *ListenerFilter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isListenerFilter_ConfigType interface {
	isListenerFilter_ConfigType()
}

type ListenerFilter_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*ListenerFilter_TypedConfig) isListenerFilter_ConfigType() {}

func (m *ListenerFilter) GetConfigType() isListenerFilter_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *ListenerFilter) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*ListenerFilter_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ListenerFilter) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ListenerFilter_TypedConfig)(nil),
	}
}

func init() {
	proto.RegisterEnum("envoy.api.v3alpha.listener.FilterChainMatch_ConnectionSourceType", FilterChainMatch_ConnectionSourceType_name, FilterChainMatch_ConnectionSourceType_value)
	proto.RegisterType((*Filter)(nil), "envoy.api.v3alpha.listener.Filter")
	proto.RegisterType((*FilterChainMatch)(nil), "envoy.api.v3alpha.listener.FilterChainMatch")
	proto.RegisterType((*FilterChain)(nil), "envoy.api.v3alpha.listener.FilterChain")
	proto.RegisterType((*ListenerFilter)(nil), "envoy.api.v3alpha.listener.ListenerFilter")
}

func init() {
	proto.RegisterFile("envoy/api/v3alpha/listener/listener.proto", fileDescriptor_907f943165b7183f)
}

var fileDescriptor_907f943165b7183f = []byte{
	// 871 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6f, 0xdb, 0x46,
	0x10, 0x0d, 0x4d, 0x59, 0x1f, 0x43, 0xc9, 0x66, 0x17, 0x2e, 0xca, 0xba, 0x86, 0x21, 0x2b, 0x69,
	0xca, 0x7e, 0x84, 0x04, 0xa4, 0x4b, 0xeb, 0xe4, 0x22, 0x0a, 0x09, 0x5a, 0x43, 0x71, 0x05, 0xda,
	0x2d, 0xda, 0x5e, 0x88, 0x35, 0xb5, 0x92, 0x89, 0xd2, 0xbb, 0xc4, 0xee, 0x4a, 0xb5, 0xae, 0x39,
	0xf5, 0xdc, 0x63, 0x7f, 0x41, 0xcf, 0xfd, 0x0f, 0xbd, 0xf6, 0x07, 0xf9, 0x92, 0x82, 0xcb, 0xa5,
	0x2c, 0x5b, 0x71, 0x9c, 0xe6, 0xc6, 0x9d, 0x79, 0xf3, 0x38, 0xfb, 0xe6, 0xed, 0xc0, 0xe7, 0x84,
	0xce, 0xd9, 0xc2, 0xc7, 0x59, 0xe2, 0xcf, 0x7b, 0x38, 0xcd, 0xce, 0xb1, 0x9f, 0x26, 0x42, 0x12,
	0x4a, 0xf8, 0xf2, 0xc3, 0xcb, 0x38, 0x93, 0x0c, 0xed, 0x2a, 0xa8, 0x87, 0xb3, 0xc4, 0xd3, 0x50,
	0xaf, 0x44, 0xec, 0x3e, 0x5a, 0xa7, 0x89, 0x19, 0x27, 0x3e, 0x1e, 0x8f, 0x39, 0x11, 0xa2, 0x60,
	0xd8, 0x3d, 0xb8, 0x03, 0x75, 0x86, 0x05, 0xd1, 0x90, 0x8f, 0xa7, 0x8c, 0x4d, 0x53, 0xe2, 0xab,
	0xd3, 0xd9, 0x6c, 0xe2, 0x63, 0xba, 0xd0, 0xa9, 0xbd, 0xdb, 0x29, 0x21, 0xf9, 0x2c, 0x96, 0x3a,
	0xbb, 0x7f, 0x3b, 0xfb, 0x1b, 0xc7, 0x59, 0x46, 0xf8, 0xf2, 0xdf, 0xb3, 0x71, 0x86, 0x7d, 0x4c,
	0x29, 0x93, 0x58, 0x26, 0x8c, 0x0a, 0x7f, 0x4e, 0xb8, 0x48, 0x18, 0x4d, 0xe8, 0x54, 0x43, 0x3e,
	0x9a, 0xe3, 0x34, 0x19, 0x63, 0x49, 0xfc, 0xf2, 0xa3, 0x48, 0x74, 0xfe, 0x32, 0xa0, 0xfa, 0x22,
	0x49, 0x25, 0xe1, 0xe8, 0x13, 0xa8, 0x50, 0x7c, 0x41, 0x1c, 0xa3, 0x6d, 0xb8, 0x8d, 0xa0, 0x76,
	0x15, 0x54, 0xf8, 0x46, 0xdb, 0x08, 0x55, 0x10, 0x7d, 0x03, 0x4d, 0xb9, 0xc8, 0xc8, 0x38, 0x8a,
	0x19, 0x9d, 0x24, 0x53, 0xa7, 0xd2, 0x36, 0x5c, 0xab, 0xbb, 0xe3, 0x15, 0xad, 0x79, 0x65, 0x6b,
	0x5e, 0x9f, 0x2e, 0xbe, 0x7d, 0x10, 0x5a, 0x0a, 0x3b, 0x50, 0xd0, 0xc3, 0x87, 0x7f, 0xfe, 0xf3,
	0xfb, 0xfe, 0x3e, 0xec, 0xad, 0x68, 0xdc, 0x5d, 0xca, 0xeb, 0x15, 0x3f, 0x0f, 0x5a, 0x60, 0x15,
	0xcc, 0x51, 0x5e, 0x7a, 0x54, 0xa9, 0x9b, 0x76, 0xe5, 0xa8, 0x52, 0xdf, 0xb0, 0xcd, 0xb0, 0x5a,
	0x24, 0x3a, 0xaf, 0xaa, 0x60, 0x17, 0xe8, 0xc1, 0x39, 0x4e, 0xe8, 0x4b, 0x2c, 0xe3, 0x73, 0x74,
	0x0a, 0xf6, 0x98, 0x08, 0x99, 0x50, 0x75, 0xf3, 0x28, 0x63, 0x5c, 0x3a, 0x75, 0xd5, 0xdb, 0xde,
	0x5a, 0x6f, 0x3f, 0x7c, 0x47, 0x65, 0xaf, 0xfb, 0x23, 0x4e, 0x67, 0x24, 0xb0, 0xae, 0x82, 0xfa,
	0x17, 0x55, 0xe7, 0xf5, 0x6b, 0xd3, 0x35, 0xc2, 0xed, 0x15, 0x8a, 0x11, 0xe3, 0x12, 0xbd, 0x80,
	0x56, 0xc6, 0xc9, 0x24, 0xb9, 0x8c, 0x38, 0xa6, 0x53, 0x22, 0x1c, 0xb3, 0x6d, 0xba, 0x56, 0xf7,
	0xc0, 0x5b, 0xf7, 0x49, 0x3e, 0x65, 0x6f, 0x90, 0x8c, 0x79, 0x98, 0x23, 0xc3, 0x66, 0x51, 0xa7,
	0x0e, 0x02, 0x7d, 0x0a, 0x5b, 0xda, 0x26, 0x91, 0x98, 0x4d, 0x26, 0xc9, 0xa5, 0xd2, 0xad, 0x11,
	0xb6, 0x74, 0xf4, 0x44, 0x05, 0xd1, 0x53, 0x80, 0x22, 0x1d, 0xa5, 0x84, 0x3a, 0x9b, 0xf7, 0xb7,
	0x1f, 0x36, 0x0a, 0xfc, 0x90, 0x50, 0x94, 0x82, 0x25, 0xd8, 0x8c, 0xc7, 0x44, 0x29, 0xe7, 0x34,
	0xdb, 0x86, 0xbb, 0xd5, 0xed, 0x7b, 0x77, 0x3b, 0xda, 0xbb, 0x2d, 0xa2, 0x37, 0x60, 0x94, 0x92,
	0x38, 0xbf, 0xfd, 0x89, 0x62, 0x3a, 0x5d, 0x64, 0x24, 0xa8, 0x5f, 0x05, 0x9b, 0xaf, 0x8c, 0x0d,
	0xdb, 0x08, 0x41, 0x2c, 0xa3, 0xe8, 0x04, 0x76, 0xf4, 0xdf, 0x6e, 0x0a, 0x54, 0x7d, 0x57, 0x81,
	0x50, 0x51, 0x3e, 0x5a, 0x95, 0xa9, 0x07, 0xcd, 0x92, 0x94, 0x71, 0x29, 0x9c, 0x5a, 0xdb, 0x74,
	0x5b, 0x81, 0x7d, 0x15, 0xb4, 0xfe, 0x30, 0xa0, 0x73, 0x3d, 0x27, 0x7d, 0xd1, 0x7c, 0x44, 0x02,
	0x1d, 0x40, 0x53, 0x10, 0x3e, 0x27, 0x3c, 0xca, 0x0d, 0x2a, 0x1c, 0xab, 0x6d, 0xba, 0x8d, 0xd0,
	0x2a, 0x62, 0xc7, 0x79, 0x08, 0x3d, 0x01, 0x24, 0x39, 0xa6, 0x22, 0x67, 0x8d, 0x94, 0x8e, 0x31,
	0x4b, 0x9d, 0x86, 0x1a, 0xc1, 0x07, 0xcb, 0xcc, 0x48, 0x27, 0x50, 0x0f, 0x3e, 0xc4, 0x59, 0x96,
	0x26, 0xb1, 0xf6, 0x92, 0x8e, 0x0b, 0x07, 0x14, 0xf5, 0xce, 0x4a, 0xb2, 0xac, 0x11, 0x9d, 0xaf,
	0x61, 0xe7, 0x4d, 0xf2, 0xa1, 0x1a, 0x98, 0xfd, 0xe3, 0x9f, 0xed, 0x07, 0xa8, 0x01, 0x9b, 0xc3,
	0xef, 0x07, 0xfd, 0xa1, 0x6d, 0xa0, 0x26, 0xd4, 0x9f, 0xff, 0x74, 0xfa, 0x3c, 0x3c, 0xee, 0x0f,
	0xed, 0x8d, 0xc3, 0x27, 0xf9, 0xbb, 0x70, 0xe1, 0xf1, 0xdb, 0xde, 0xc5, 0xf5, 0x90, 0x8e, 0x2a,
	0x75, 0xc3, 0xde, 0xe8, 0xfc, 0x6b, 0x82, 0xb5, 0x92, 0x42, 0xbf, 0x00, 0x9a, 0xa8, 0x63, 0x14,
	0xe7, 0xe7, 0xe8, 0x22, 0xc7, 0xaa, 0x27, 0x6c, 0x75, 0xbf, 0xfa, 0x3f, 0x26, 0x08, 0xed, 0xc9,
	0xed, 0xb7, 0xf5, 0x0c, 0x6a, 0x45, 0xac, 0xf4, 0x7f, 0xe7, 0x7e, 0xc2, 0xb0, 0x2c, 0x41, 0x01,
	0x6c, 0xcf, 0x44, 0x6e, 0x13, 0x76, 0xb9, 0x28, 0xb4, 0xd4, 0x4b, 0x63, 0x77, 0xcd, 0xd9, 0x01,
	0x63, 0x69, 0xe1, 0xeb, 0xd6, 0x4c, 0x90, 0x51, 0x5e, 0xa1, 0x04, 0x46, 0xcf, 0xa0, 0x7e, 0x41,
	0x24, 0x1e, 0x63, 0x89, 0xf5, 0xb3, 0x68, 0xdf, 0xe5, 0xb0, 0x97, 0x1a, 0x17, 0x2e, 0x2b, 0x50,
	0x08, 0xf6, 0xf5, 0xf8, 0x05, 0x8b, 0x7f, 0x25, 0xd2, 0xa9, 0x2a, 0x96, 0xcf, 0xee, 0x62, 0x39,
	0x2d, 0xf1, 0x27, 0x0a, 0x1e, 0x6e, 0xcb, 0x9b, 0x01, 0x84, 0xf4, 0x92, 0xac, 0x29, 0x13, 0xa9,
	0xef, 0x43, 0x37, 0x1f, 0xe4, 0x43, 0x38, 0xb8, 0x77, 0x90, 0x7a, 0xa1, 0x59, 0x32, 0x15, 0xf9,
	0x1e, 0x95, 0xe4, 0x52, 0x76, 0xfe, 0x36, 0x60, 0x6b, 0xa8, 0xb1, 0xef, 0xb3, 0x88, 0xcd, 0x77,
	0x5f, 0xc4, 0x5f, 0xe6, 0x7d, 0x3e, 0x86, 0x47, 0x6f, 0xee, 0xf3, 0x66, 0x13, 0xeb, 0x0b, 0x79,
	0x65, 0x15, 0x07, 0x4f, 0xc1, 0x4d, 0x58, 0xa1, 0xa1, 0x9a, 0xef, 0x5b, 0x7c, 0x11, 0xb4, 0x4a,
	0x62, 0x35, 0xd2, 0x91, 0x71, 0x56, 0x55, 0x3d, 0xf6, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x5c,
	0x4b, 0x3e, 0xa7, 0xa0, 0x07, 0x00, 0x00,
}
