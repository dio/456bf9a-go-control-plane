// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v3alpha/endpoint/endpoint.proto

package envoy_api_v3alpha_endpoint

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/core"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

type Endpoint struct {
	Address              *core.Address               `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	HealthCheckConfig    *Endpoint_HealthCheckConfig `protobuf:"bytes,2,opt,name=health_check_config,json=healthCheckConfig,proto3" json:"health_check_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e420efb9b41886f, []int{0}
}

func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint.Unmarshal(m, b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
}
func (m *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(m, src)
}
func (m *Endpoint) XXX_Size() int {
	return xxx_messageInfo_Endpoint.Size(m)
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetAddress() *core.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Endpoint) GetHealthCheckConfig() *Endpoint_HealthCheckConfig {
	if m != nil {
		return m.HealthCheckConfig
	}
	return nil
}

type Endpoint_HealthCheckConfig struct {
	PortValue            uint32   `protobuf:"varint,1,opt,name=port_value,json=portValue,proto3" json:"port_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Endpoint_HealthCheckConfig) Reset()         { *m = Endpoint_HealthCheckConfig{} }
func (m *Endpoint_HealthCheckConfig) String() string { return proto.CompactTextString(m) }
func (*Endpoint_HealthCheckConfig) ProtoMessage()    {}
func (*Endpoint_HealthCheckConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e420efb9b41886f, []int{0, 0}
}

func (m *Endpoint_HealthCheckConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint_HealthCheckConfig.Unmarshal(m, b)
}
func (m *Endpoint_HealthCheckConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint_HealthCheckConfig.Marshal(b, m, deterministic)
}
func (m *Endpoint_HealthCheckConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint_HealthCheckConfig.Merge(m, src)
}
func (m *Endpoint_HealthCheckConfig) XXX_Size() int {
	return xxx_messageInfo_Endpoint_HealthCheckConfig.Size(m)
}
func (m *Endpoint_HealthCheckConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint_HealthCheckConfig.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint_HealthCheckConfig proto.InternalMessageInfo

func (m *Endpoint_HealthCheckConfig) GetPortValue() uint32 {
	if m != nil {
		return m.PortValue
	}
	return 0
}

type LbEndpoint struct {
	// Types that are valid to be assigned to HostIdentifier:
	//	*LbEndpoint_Endpoint
	//	*LbEndpoint_EndpointName
	HostIdentifier       isLbEndpoint_HostIdentifier `protobuf_oneof:"host_identifier"`
	HealthStatus         core.HealthStatus           `protobuf:"varint,2,opt,name=health_status,json=healthStatus,proto3,enum=envoy.api.v3alpha.core.HealthStatus" json:"health_status,omitempty"`
	Metadata             *core.Metadata              `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	LoadBalancingWeight  *wrappers.UInt32Value       `protobuf:"bytes,4,opt,name=load_balancing_weight,json=loadBalancingWeight,proto3" json:"load_balancing_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *LbEndpoint) Reset()         { *m = LbEndpoint{} }
func (m *LbEndpoint) String() string { return proto.CompactTextString(m) }
func (*LbEndpoint) ProtoMessage()    {}
func (*LbEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e420efb9b41886f, []int{1}
}

func (m *LbEndpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LbEndpoint.Unmarshal(m, b)
}
func (m *LbEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LbEndpoint.Marshal(b, m, deterministic)
}
func (m *LbEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LbEndpoint.Merge(m, src)
}
func (m *LbEndpoint) XXX_Size() int {
	return xxx_messageInfo_LbEndpoint.Size(m)
}
func (m *LbEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_LbEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_LbEndpoint proto.InternalMessageInfo

type isLbEndpoint_HostIdentifier interface {
	isLbEndpoint_HostIdentifier()
}

type LbEndpoint_Endpoint struct {
	Endpoint *Endpoint `protobuf:"bytes,1,opt,name=endpoint,proto3,oneof"`
}

type LbEndpoint_EndpointName struct {
	EndpointName string `protobuf:"bytes,5,opt,name=endpoint_name,json=endpointName,proto3,oneof"`
}

func (*LbEndpoint_Endpoint) isLbEndpoint_HostIdentifier() {}

func (*LbEndpoint_EndpointName) isLbEndpoint_HostIdentifier() {}

func (m *LbEndpoint) GetHostIdentifier() isLbEndpoint_HostIdentifier {
	if m != nil {
		return m.HostIdentifier
	}
	return nil
}

func (m *LbEndpoint) GetEndpoint() *Endpoint {
	if x, ok := m.GetHostIdentifier().(*LbEndpoint_Endpoint); ok {
		return x.Endpoint
	}
	return nil
}

func (m *LbEndpoint) GetEndpointName() string {
	if x, ok := m.GetHostIdentifier().(*LbEndpoint_EndpointName); ok {
		return x.EndpointName
	}
	return ""
}

func (m *LbEndpoint) GetHealthStatus() core.HealthStatus {
	if m != nil {
		return m.HealthStatus
	}
	return core.HealthStatus_UNKNOWN
}

func (m *LbEndpoint) GetMetadata() *core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *LbEndpoint) GetLoadBalancingWeight() *wrappers.UInt32Value {
	if m != nil {
		return m.LoadBalancingWeight
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*LbEndpoint) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*LbEndpoint_Endpoint)(nil),
		(*LbEndpoint_EndpointName)(nil),
	}
}

type LocalityLbEndpoints struct {
	Locality             *core.Locality        `protobuf:"bytes,1,opt,name=locality,proto3" json:"locality,omitempty"`
	LbEndpoints          []*LbEndpoint         `protobuf:"bytes,2,rep,name=lb_endpoints,json=lbEndpoints,proto3" json:"lb_endpoints,omitempty"`
	LoadBalancingWeight  *wrappers.UInt32Value `protobuf:"bytes,3,opt,name=load_balancing_weight,json=loadBalancingWeight,proto3" json:"load_balancing_weight,omitempty"`
	Priority             uint32                `protobuf:"varint,5,opt,name=priority,proto3" json:"priority,omitempty"`
	Proximity            *wrappers.UInt32Value `protobuf:"bytes,6,opt,name=proximity,proto3" json:"proximity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *LocalityLbEndpoints) Reset()         { *m = LocalityLbEndpoints{} }
func (m *LocalityLbEndpoints) String() string { return proto.CompactTextString(m) }
func (*LocalityLbEndpoints) ProtoMessage()    {}
func (*LocalityLbEndpoints) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e420efb9b41886f, []int{2}
}

func (m *LocalityLbEndpoints) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalityLbEndpoints.Unmarshal(m, b)
}
func (m *LocalityLbEndpoints) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalityLbEndpoints.Marshal(b, m, deterministic)
}
func (m *LocalityLbEndpoints) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalityLbEndpoints.Merge(m, src)
}
func (m *LocalityLbEndpoints) XXX_Size() int {
	return xxx_messageInfo_LocalityLbEndpoints.Size(m)
}
func (m *LocalityLbEndpoints) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalityLbEndpoints.DiscardUnknown(m)
}

var xxx_messageInfo_LocalityLbEndpoints proto.InternalMessageInfo

func (m *LocalityLbEndpoints) GetLocality() *core.Locality {
	if m != nil {
		return m.Locality
	}
	return nil
}

func (m *LocalityLbEndpoints) GetLbEndpoints() []*LbEndpoint {
	if m != nil {
		return m.LbEndpoints
	}
	return nil
}

func (m *LocalityLbEndpoints) GetLoadBalancingWeight() *wrappers.UInt32Value {
	if m != nil {
		return m.LoadBalancingWeight
	}
	return nil
}

func (m *LocalityLbEndpoints) GetPriority() uint32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *LocalityLbEndpoints) GetProximity() *wrappers.UInt32Value {
	if m != nil {
		return m.Proximity
	}
	return nil
}

func init() {
	proto.RegisterType((*Endpoint)(nil), "envoy.api.v3alpha.endpoint.Endpoint")
	proto.RegisterType((*Endpoint_HealthCheckConfig)(nil), "envoy.api.v3alpha.endpoint.Endpoint.HealthCheckConfig")
	proto.RegisterType((*LbEndpoint)(nil), "envoy.api.v3alpha.endpoint.LbEndpoint")
	proto.RegisterType((*LocalityLbEndpoints)(nil), "envoy.api.v3alpha.endpoint.LocalityLbEndpoints")
}

func init() {
	proto.RegisterFile("envoy/api/v3alpha/endpoint/endpoint.proto", fileDescriptor_9e420efb9b41886f)
}

var fileDescriptor_9e420efb9b41886f = []byte{
	// 636 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcf, 0x6b, 0x13, 0x4f,
	0x1c, 0x6d, 0xba, 0xfd, 0x91, 0x4c, 0x9b, 0xef, 0x97, 0x6e, 0x11, 0x97, 0x20, 0x35, 0x86, 0x56,
	0xd3, 0x0a, 0xb3, 0x25, 0x05, 0xc5, 0xe8, 0xc5, 0x2d, 0x42, 0x0b, 0x55, 0xca, 0x88, 0x8a, 0xa7,
	0xe5, 0x93, 0xec, 0x24, 0x3b, 0xb8, 0x99, 0x59, 0x76, 0x27, 0xa9, 0xbd, 0x79, 0xf4, 0xe2, 0x59,
	0xf0, 0xff, 0xf1, 0xe6, 0x5f, 0xd4, 0x4b, 0x65, 0x26, 0x33, 0x9b, 0x60, 0x9a, 0x58, 0xf0, 0xb6,
	0x3b, 0xf3, 0xde, 0xfb, 0xbc, 0xcf, 0x7b, 0x9b, 0xa0, 0x7d, 0xca, 0x47, 0xe2, 0xd2, 0x87, 0x94,
	0xf9, 0xa3, 0x23, 0x48, 0xd2, 0x18, 0x7c, 0xca, 0xa3, 0x54, 0x30, 0x2e, 0x8b, 0x07, 0x9c, 0x66,
	0x42, 0x0a, 0xb7, 0xa6, 0xa1, 0x18, 0x52, 0x86, 0x0d, 0x14, 0x5b, 0x44, 0x6d, 0x77, 0x56, 0xa6,
	0x2b, 0x32, 0xea, 0x43, 0x14, 0x65, 0x34, 0xcf, 0xc7, 0x0a, 0xb5, 0x07, 0x73, 0x50, 0x1d, 0xc8,
	0xa9, 0x81, 0xec, 0xcf, 0x81, 0xc4, 0x14, 0x12, 0x19, 0x87, 0xdd, 0x98, 0x76, 0x3f, 0x19, 0xe8,
	0x4e, 0x5f, 0x88, 0x7e, 0x42, 0x7d, 0xfd, 0xd6, 0x19, 0xf6, 0xfc, 0x8b, 0x0c, 0xd2, 0x94, 0x66,
	0xc5, 0xb4, 0x61, 0x94, 0x82, 0x0f, 0x9c, 0x0b, 0x09, 0x92, 0x09, 0x9e, 0xfb, 0x23, 0x9a, 0xe5,
	0x4c, 0x70, 0xc6, 0xfb, 0x06, 0x72, 0x77, 0x04, 0x09, 0x8b, 0x40, 0x52, 0xdf, 0x3e, 0x8c, 0x2f,
	0x1a, 0xbf, 0x96, 0x51, 0xf9, 0x95, 0x59, 0xce, 0x7d, 0x86, 0xd6, 0xcd, 0x1e, 0x5e, 0xa9, 0x5e,
	0x6a, 0x6e, 0xb4, 0xee, 0xe3, 0xd9, 0x28, 0x94, 0x4b, 0xfc, 0x72, 0x0c, 0x23, 0x16, 0xef, 0xf6,
	0xd0, 0xf6, 0xb4, 0xf3, 0xb0, 0x2b, 0x78, 0x8f, 0xf5, 0xbd, 0x65, 0x2d, 0xf3, 0x04, 0xcf, 0x4f,
	0x14, 0xdb, 0xe9, 0xf8, 0x44, 0xf3, 0x8f, 0x15, 0xfd, 0x58, 0xb3, 0xc9, 0x56, 0xfc, 0xe7, 0x51,
	0x6d, 0x84, 0xb6, 0x66, 0x70, 0x6e, 0x13, 0xa1, 0x54, 0x64, 0x32, 0x1c, 0x41, 0x32, 0xa4, 0xda,
	0x7a, 0x35, 0xa8, 0x5c, 0x05, 0x6b, 0x07, 0x2b, 0xde, 0xf5, 0xb5, 0x43, 0x2a, 0xea, 0xf2, 0xbd,
	0xba, 0x6b, 0x3f, 0xfd, 0xf1, 0xf3, 0xeb, 0x4e, 0x0b, 0x1d, 0x4e, 0xf9, 0x69, 0xdd, 0xc6, 0x4a,
	0x7b, 0x4f, 0x11, 0xeb, 0x68, 0x67, 0x31, 0xb1, 0xf1, 0xdd, 0x41, 0xe8, 0xac, 0x53, 0x04, 0x1a,
	0xa0, 0xb2, 0xc5, 0x98, 0x44, 0x77, 0x6f, 0x13, 0xc5, 0xc9, 0x12, 0x29, 0x78, 0xee, 0x1e, 0xaa,
	0xda, 0xe7, 0x90, 0xc3, 0x80, 0x7a, 0xab, 0xf5, 0x52, 0xb3, 0x72, 0xb2, 0x44, 0x36, 0xed, 0xf1,
	0x1b, 0x18, 0x50, 0xf7, 0x14, 0x55, 0x4d, 0x01, 0xb9, 0x04, 0x39, 0xcc, 0x75, 0xf4, 0xff, 0xdd,
	0x38, 0x4f, 0x37, 0x38, 0x5e, 0xf1, 0xad, 0xc6, 0x92, 0xcd, 0x78, 0xea, 0xcd, 0x7d, 0x81, 0xca,
	0x03, 0x2a, 0x21, 0x02, 0x09, 0x9e, 0xa3, 0x5d, 0xd7, 0xe7, 0xa9, 0xbc, 0x36, 0x38, 0x52, 0x30,
	0xdc, 0x8f, 0xe8, 0x4e, 0x22, 0x20, 0x0a, 0x3b, 0x90, 0x00, 0xef, 0x32, 0xde, 0x0f, 0x2f, 0x28,
	0xeb, 0xc7, 0xd2, 0x5b, 0xd1, 0x52, 0xf7, 0xf0, 0xf8, 0x6b, 0xc6, 0xf6, 0x6b, 0xc6, 0xef, 0x4e,
	0xb9, 0x3c, 0x6a, 0xe9, 0x7e, 0x82, 0xf5, 0xab, 0x60, 0xe5, 0x60, 0xb9, 0x59, 0x22, 0xdb, 0x4a,
	0x23, 0xb0, 0x12, 0x1f, 0xb4, 0x42, 0xfb, 0x91, 0x2a, 0xa1, 0x81, 0xea, 0x37, 0x97, 0x30, 0xc9,
	0x3d, 0xd8, 0x42, 0xff, 0xc7, 0x22, 0x97, 0x21, 0x8b, 0x28, 0x97, 0xac, 0xc7, 0x68, 0xd6, 0xf8,
	0xe6, 0xa0, 0xed, 0x33, 0xd1, 0x85, 0x84, 0xc9, 0xcb, 0x09, 0x52, 0x2f, 0x9b, 0x98, 0x63, 0x53,
	0xd1, 0xdc, 0x65, 0x2d, 0x9d, 0x14, 0x0c, 0xf7, 0x14, 0x6d, 0x26, 0x9d, 0xd0, 0x5a, 0x50, 0xa1,
	0x3b, 0xcd, 0x8d, 0xd6, 0xc3, 0x45, 0x25, 0x4f, 0x86, 0x93, 0x8d, 0x64, 0xca, 0xc8, 0xdc, 0xdc,
	0x9c, 0x7f, 0xcd, 0xcd, 0xdd, 0x45, 0xe5, 0x34, 0x63, 0x22, 0x53, 0x3b, 0xae, 0xea, 0x5f, 0x47,
	0xf9, 0x2a, 0x58, 0x3d, 0x70, 0xbc, 0x2f, 0x25, 0x52, 0xdc, 0xb8, 0x6d, 0x54, 0x49, 0x33, 0xf1,
	0x99, 0x0d, 0x14, 0x6c, 0xed, 0xef, 0x43, 0xc9, 0x04, 0xde, 0x3e, 0x54, 0xcd, 0x3c, 0x36, 0x7f,
	0xb2, 0xb3, 0xcd, 0xcc, 0xe6, 0x1e, 0x3c, 0x47, 0x4d, 0x26, 0xc6, 0x39, 0x29, 0x99, 0xcb, 0x05,
	0x91, 0x05, 0x55, 0x4b, 0x3b, 0x57, 0x36, 0xce, 0x4b, 0x9d, 0x35, 0xed, 0xe7, 0xe8, 0x77, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xcf, 0xbf, 0xc5, 0xfe, 0xd5, 0x05, 0x00, 0x00,
}
