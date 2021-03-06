// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/discovery/v3alpha/hds.proto

package envoy_service_discovery_v3alpha

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/core"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/endpoint"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Capability_Protocol int32

const (
	Capability_HTTP  Capability_Protocol = 0
	Capability_TCP   Capability_Protocol = 1
	Capability_REDIS Capability_Protocol = 2
)

var Capability_Protocol_name = map[int32]string{
	0: "HTTP",
	1: "TCP",
	2: "REDIS",
}

var Capability_Protocol_value = map[string]int32{
	"HTTP":  0,
	"TCP":   1,
	"REDIS": 2,
}

func (x Capability_Protocol) String() string {
	return proto.EnumName(Capability_Protocol_name, int32(x))
}

func (Capability_Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_298093248f9d2656, []int{0, 0}
}

type Capability struct {
	HealthCheckProtocols []Capability_Protocol `protobuf:"varint,1,rep,packed,name=health_check_protocols,json=healthCheckProtocols,proto3,enum=envoy.service.discovery.v3alpha.Capability_Protocol" json:"health_check_protocols,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Capability) Reset()         { *m = Capability{} }
func (m *Capability) String() string { return proto.CompactTextString(m) }
func (*Capability) ProtoMessage()    {}
func (*Capability) Descriptor() ([]byte, []int) {
	return fileDescriptor_298093248f9d2656, []int{0}
}

func (m *Capability) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Capability.Unmarshal(m, b)
}
func (m *Capability) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Capability.Marshal(b, m, deterministic)
}
func (m *Capability) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Capability.Merge(m, src)
}
func (m *Capability) XXX_Size() int {
	return xxx_messageInfo_Capability.Size(m)
}
func (m *Capability) XXX_DiscardUnknown() {
	xxx_messageInfo_Capability.DiscardUnknown(m)
}

var xxx_messageInfo_Capability proto.InternalMessageInfo

func (m *Capability) GetHealthCheckProtocols() []Capability_Protocol {
	if m != nil {
		return m.HealthCheckProtocols
	}
	return nil
}

type HealthCheckRequest struct {
	Node                 *core.Node  `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Capability           *Capability `protobuf:"bytes,2,opt,name=capability,proto3" json:"capability,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *HealthCheckRequest) Reset()         { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string { return proto.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()    {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_298093248f9d2656, []int{1}
}

func (m *HealthCheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckRequest.Unmarshal(m, b)
}
func (m *HealthCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckRequest.Marshal(b, m, deterministic)
}
func (m *HealthCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckRequest.Merge(m, src)
}
func (m *HealthCheckRequest) XXX_Size() int {
	return xxx_messageInfo_HealthCheckRequest.Size(m)
}
func (m *HealthCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckRequest proto.InternalMessageInfo

func (m *HealthCheckRequest) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *HealthCheckRequest) GetCapability() *Capability {
	if m != nil {
		return m.Capability
	}
	return nil
}

type EndpointHealth struct {
	Endpoint             *endpoint.Endpoint `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	HealthStatus         core.HealthStatus  `protobuf:"varint,2,opt,name=health_status,json=healthStatus,proto3,enum=envoy.api.v3alpha.core.HealthStatus" json:"health_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *EndpointHealth) Reset()         { *m = EndpointHealth{} }
func (m *EndpointHealth) String() string { return proto.CompactTextString(m) }
func (*EndpointHealth) ProtoMessage()    {}
func (*EndpointHealth) Descriptor() ([]byte, []int) {
	return fileDescriptor_298093248f9d2656, []int{2}
}

func (m *EndpointHealth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointHealth.Unmarshal(m, b)
}
func (m *EndpointHealth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointHealth.Marshal(b, m, deterministic)
}
func (m *EndpointHealth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointHealth.Merge(m, src)
}
func (m *EndpointHealth) XXX_Size() int {
	return xxx_messageInfo_EndpointHealth.Size(m)
}
func (m *EndpointHealth) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointHealth.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointHealth proto.InternalMessageInfo

func (m *EndpointHealth) GetEndpoint() *endpoint.Endpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *EndpointHealth) GetHealthStatus() core.HealthStatus {
	if m != nil {
		return m.HealthStatus
	}
	return core.HealthStatus_UNKNOWN
}

type EndpointHealthResponse struct {
	EndpointsHealth      []*EndpointHealth `protobuf:"bytes,1,rep,name=endpoints_health,json=endpointsHealth,proto3" json:"endpoints_health,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *EndpointHealthResponse) Reset()         { *m = EndpointHealthResponse{} }
func (m *EndpointHealthResponse) String() string { return proto.CompactTextString(m) }
func (*EndpointHealthResponse) ProtoMessage()    {}
func (*EndpointHealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_298093248f9d2656, []int{3}
}

func (m *EndpointHealthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndpointHealthResponse.Unmarshal(m, b)
}
func (m *EndpointHealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndpointHealthResponse.Marshal(b, m, deterministic)
}
func (m *EndpointHealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndpointHealthResponse.Merge(m, src)
}
func (m *EndpointHealthResponse) XXX_Size() int {
	return xxx_messageInfo_EndpointHealthResponse.Size(m)
}
func (m *EndpointHealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EndpointHealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EndpointHealthResponse proto.InternalMessageInfo

func (m *EndpointHealthResponse) GetEndpointsHealth() []*EndpointHealth {
	if m != nil {
		return m.EndpointsHealth
	}
	return nil
}

type HealthCheckRequestOrEndpointHealthResponse struct {
	// Types that are valid to be assigned to RequestType:
	//	*HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest
	//	*HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse
	RequestType          isHealthCheckRequestOrEndpointHealthResponse_RequestType `protobuf_oneof:"request_type"`
	XXX_NoUnkeyedLiteral struct{}                                                 `json:"-"`
	XXX_unrecognized     []byte                                                   `json:"-"`
	XXX_sizecache        int32                                                    `json:"-"`
}

func (m *HealthCheckRequestOrEndpointHealthResponse) Reset() {
	*m = HealthCheckRequestOrEndpointHealthResponse{}
}
func (m *HealthCheckRequestOrEndpointHealthResponse) String() string {
	return proto.CompactTextString(m)
}
func (*HealthCheckRequestOrEndpointHealthResponse) ProtoMessage() {}
func (*HealthCheckRequestOrEndpointHealthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_298093248f9d2656, []int{4}
}

func (m *HealthCheckRequestOrEndpointHealthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckRequestOrEndpointHealthResponse.Unmarshal(m, b)
}
func (m *HealthCheckRequestOrEndpointHealthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckRequestOrEndpointHealthResponse.Marshal(b, m, deterministic)
}
func (m *HealthCheckRequestOrEndpointHealthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckRequestOrEndpointHealthResponse.Merge(m, src)
}
func (m *HealthCheckRequestOrEndpointHealthResponse) XXX_Size() int {
	return xxx_messageInfo_HealthCheckRequestOrEndpointHealthResponse.Size(m)
}
func (m *HealthCheckRequestOrEndpointHealthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckRequestOrEndpointHealthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckRequestOrEndpointHealthResponse proto.InternalMessageInfo

type isHealthCheckRequestOrEndpointHealthResponse_RequestType interface {
	isHealthCheckRequestOrEndpointHealthResponse_RequestType()
}

type HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest struct {
	HealthCheckRequest *HealthCheckRequest `protobuf:"bytes,1,opt,name=health_check_request,json=healthCheckRequest,proto3,oneof"`
}

type HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse struct {
	EndpointHealthResponse *EndpointHealthResponse `protobuf:"bytes,2,opt,name=endpoint_health_response,json=endpointHealthResponse,proto3,oneof"`
}

func (*HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest) isHealthCheckRequestOrEndpointHealthResponse_RequestType() {
}

func (*HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse) isHealthCheckRequestOrEndpointHealthResponse_RequestType() {
}

func (m *HealthCheckRequestOrEndpointHealthResponse) GetRequestType() isHealthCheckRequestOrEndpointHealthResponse_RequestType {
	if m != nil {
		return m.RequestType
	}
	return nil
}

func (m *HealthCheckRequestOrEndpointHealthResponse) GetHealthCheckRequest() *HealthCheckRequest {
	if x, ok := m.GetRequestType().(*HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest); ok {
		return x.HealthCheckRequest
	}
	return nil
}

func (m *HealthCheckRequestOrEndpointHealthResponse) GetEndpointHealthResponse() *EndpointHealthResponse {
	if x, ok := m.GetRequestType().(*HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse); ok {
		return x.EndpointHealthResponse
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HealthCheckRequestOrEndpointHealthResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest)(nil),
		(*HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse)(nil),
	}
}

type LocalityEndpoints struct {
	Locality             *core.Locality       `protobuf:"bytes,1,opt,name=locality,proto3" json:"locality,omitempty"`
	Endpoints            []*endpoint.Endpoint `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LocalityEndpoints) Reset()         { *m = LocalityEndpoints{} }
func (m *LocalityEndpoints) String() string { return proto.CompactTextString(m) }
func (*LocalityEndpoints) ProtoMessage()    {}
func (*LocalityEndpoints) Descriptor() ([]byte, []int) {
	return fileDescriptor_298093248f9d2656, []int{5}
}

func (m *LocalityEndpoints) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalityEndpoints.Unmarshal(m, b)
}
func (m *LocalityEndpoints) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalityEndpoints.Marshal(b, m, deterministic)
}
func (m *LocalityEndpoints) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalityEndpoints.Merge(m, src)
}
func (m *LocalityEndpoints) XXX_Size() int {
	return xxx_messageInfo_LocalityEndpoints.Size(m)
}
func (m *LocalityEndpoints) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalityEndpoints.DiscardUnknown(m)
}

var xxx_messageInfo_LocalityEndpoints proto.InternalMessageInfo

func (m *LocalityEndpoints) GetLocality() *core.Locality {
	if m != nil {
		return m.Locality
	}
	return nil
}

func (m *LocalityEndpoints) GetEndpoints() []*endpoint.Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

type ClusterHealthCheck struct {
	ClusterName          string               `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	HealthChecks         []*core.HealthCheck  `protobuf:"bytes,2,rep,name=health_checks,json=healthChecks,proto3" json:"health_checks,omitempty"`
	LocalityEndpoints    []*LocalityEndpoints `protobuf:"bytes,3,rep,name=locality_endpoints,json=localityEndpoints,proto3" json:"locality_endpoints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ClusterHealthCheck) Reset()         { *m = ClusterHealthCheck{} }
func (m *ClusterHealthCheck) String() string { return proto.CompactTextString(m) }
func (*ClusterHealthCheck) ProtoMessage()    {}
func (*ClusterHealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_298093248f9d2656, []int{6}
}

func (m *ClusterHealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterHealthCheck.Unmarshal(m, b)
}
func (m *ClusterHealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterHealthCheck.Marshal(b, m, deterministic)
}
func (m *ClusterHealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterHealthCheck.Merge(m, src)
}
func (m *ClusterHealthCheck) XXX_Size() int {
	return xxx_messageInfo_ClusterHealthCheck.Size(m)
}
func (m *ClusterHealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterHealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterHealthCheck proto.InternalMessageInfo

func (m *ClusterHealthCheck) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ClusterHealthCheck) GetHealthChecks() []*core.HealthCheck {
	if m != nil {
		return m.HealthChecks
	}
	return nil
}

func (m *ClusterHealthCheck) GetLocalityEndpoints() []*LocalityEndpoints {
	if m != nil {
		return m.LocalityEndpoints
	}
	return nil
}

type HealthCheckSpecifier struct {
	ClusterHealthChecks  []*ClusterHealthCheck `protobuf:"bytes,1,rep,name=cluster_health_checks,json=clusterHealthChecks,proto3" json:"cluster_health_checks,omitempty"`
	Interval             *duration.Duration    `protobuf:"bytes,2,opt,name=interval,proto3" json:"interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *HealthCheckSpecifier) Reset()         { *m = HealthCheckSpecifier{} }
func (m *HealthCheckSpecifier) String() string { return proto.CompactTextString(m) }
func (*HealthCheckSpecifier) ProtoMessage()    {}
func (*HealthCheckSpecifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_298093248f9d2656, []int{7}
}

func (m *HealthCheckSpecifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckSpecifier.Unmarshal(m, b)
}
func (m *HealthCheckSpecifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckSpecifier.Marshal(b, m, deterministic)
}
func (m *HealthCheckSpecifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckSpecifier.Merge(m, src)
}
func (m *HealthCheckSpecifier) XXX_Size() int {
	return xxx_messageInfo_HealthCheckSpecifier.Size(m)
}
func (m *HealthCheckSpecifier) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckSpecifier.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckSpecifier proto.InternalMessageInfo

func (m *HealthCheckSpecifier) GetClusterHealthChecks() []*ClusterHealthCheck {
	if m != nil {
		return m.ClusterHealthChecks
	}
	return nil
}

func (m *HealthCheckSpecifier) GetInterval() *duration.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.service.discovery.v3alpha.Capability_Protocol", Capability_Protocol_name, Capability_Protocol_value)
	proto.RegisterType((*Capability)(nil), "envoy.service.discovery.v3alpha.Capability")
	proto.RegisterType((*HealthCheckRequest)(nil), "envoy.service.discovery.v3alpha.HealthCheckRequest")
	proto.RegisterType((*EndpointHealth)(nil), "envoy.service.discovery.v3alpha.EndpointHealth")
	proto.RegisterType((*EndpointHealthResponse)(nil), "envoy.service.discovery.v3alpha.EndpointHealthResponse")
	proto.RegisterType((*HealthCheckRequestOrEndpointHealthResponse)(nil), "envoy.service.discovery.v3alpha.HealthCheckRequestOrEndpointHealthResponse")
	proto.RegisterType((*LocalityEndpoints)(nil), "envoy.service.discovery.v3alpha.LocalityEndpoints")
	proto.RegisterType((*ClusterHealthCheck)(nil), "envoy.service.discovery.v3alpha.ClusterHealthCheck")
	proto.RegisterType((*HealthCheckSpecifier)(nil), "envoy.service.discovery.v3alpha.HealthCheckSpecifier")
}

func init() {
	proto.RegisterFile("envoy/service/discovery/v3alpha/hds.proto", fileDescriptor_298093248f9d2656)
}

var fileDescriptor_298093248f9d2656 = []byte{
	// 876 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0x4d, 0x6f, 0xe3, 0x44,
	0x18, 0xee, 0x24, 0x0b, 0xa4, 0x6f, 0x4b, 0x49, 0x87, 0x52, 0x85, 0x6a, 0xb5, 0xdb, 0x1a, 0x10,
	0xdd, 0x2f, 0xbb, 0x9b, 0xec, 0x02, 0xb2, 0x10, 0x42, 0x4d, 0x0b, 0x59, 0xed, 0x6a, 0x89, 0x9c,
	0x9e, 0xb8, 0x44, 0x13, 0x7b, 0x36, 0x36, 0xb8, 0x1e, 0xe3, 0x71, 0x22, 0x72, 0x85, 0xcb, 0xde,
	0xb9, 0xad, 0xc4, 0x5f, 0xe0, 0x5f, 0x20, 0x01, 0x17, 0x90, 0xf8, 0x0b, 0x1c, 0x38, 0xf1, 0x1b,
	0x90, 0x67, 0xc6, 0x13, 0x67, 0x9d, 0x34, 0x09, 0x27, 0x6e, 0xf1, 0x3b, 0xcf, 0xfb, 0xf1, 0x3c,
	0x7e, 0xde, 0x71, 0xe0, 0x16, 0x8d, 0xc6, 0x6c, 0x62, 0x71, 0x9a, 0x8c, 0x03, 0x97, 0x5a, 0x5e,
	0xc0, 0x5d, 0x36, 0xa6, 0xc9, 0xc4, 0x1a, 0xb7, 0x48, 0x18, 0xfb, 0xc4, 0xf2, 0x3d, 0x6e, 0xc6,
	0x09, 0x4b, 0x19, 0xbe, 0x29, 0xa0, 0xa6, 0x82, 0x9a, 0x1a, 0x6a, 0x2a, 0xe8, 0xc1, 0x91, 0xac,
	0x45, 0xe2, 0x40, 0x67, 0xbb, 0x2c, 0xa1, 0xd6, 0x80, 0x70, 0x2a, 0x6b, 0x1c, 0xdc, 0x5a, 0x00,
	0xf1, 0x29, 0x09, 0x53, 0xbf, 0xef, 0xfa, 0xd4, 0xfd, 0x7a, 0x31, 0x94, 0x46, 0x5e, 0xcc, 0x82,
	0x28, 0xd5, 0x3f, 0x14, 0xf4, 0xfa, 0x90, 0xb1, 0x61, 0x48, 0x05, 0x96, 0x44, 0x11, 0x4b, 0x49,
	0x1a, 0xb0, 0x48, 0xcd, 0x7d, 0x70, 0x43, 0x9d, 0x8a, 0xa7, 0xc1, 0xe8, 0x99, 0xe5, 0x8d, 0x12,
	0x01, 0x50, 0xe7, 0x47, 0x23, 0x2f, 0x26, 0xc5, 0x3c, 0x6b, 0x4c, 0x13, 0x1e, 0xb0, 0x28, 0x88,
	0x86, 0x12, 0x62, 0xfc, 0x81, 0x00, 0xda, 0x24, 0x26, 0x83, 0x20, 0x0c, 0xd2, 0x09, 0xfe, 0x0a,
	0xf6, 0x8b, 0x03, 0xf7, 0x05, 0xc8, 0x65, 0x21, 0x6f, 0xa0, 0xc3, 0xea, 0xf1, 0x4e, 0xf3, 0x81,
	0xb9, 0x44, 0x2a, 0x73, 0x5a, 0xcc, 0xec, 0xaa, 0x64, 0x67, 0x4f, 0xd6, 0x6c, 0x67, 0x25, 0xf3,
	0x20, 0x37, 0x8e, 0xa1, 0x96, 0x3f, 0xe0, 0x1a, 0x5c, 0xeb, 0x5c, 0x5c, 0x74, 0xeb, 0x1b, 0xf8,
	0x35, 0xa8, 0x5e, 0xb4, 0xbb, 0x75, 0x84, 0x37, 0xe1, 0x15, 0xe7, 0xfc, 0xec, 0x51, 0xaf, 0x5e,
	0xb1, 0xef, 0xbe, 0xf8, 0xf9, 0xf9, 0x8d, 0xf7, 0xe1, 0xbd, 0x85, 0xbd, 0x9b, 0x85, 0xb6, 0xc6,
	0xaf, 0x08, 0x70, 0x67, 0xda, 0xd0, 0xa1, 0xdf, 0x8c, 0x28, 0x4f, 0xf1, 0x09, 0x5c, 0x8b, 0x98,
	0x47, 0x1b, 0xe8, 0x10, 0x1d, 0x6f, 0x35, 0xaf, 0x2b, 0x22, 0x24, 0x0e, 0xf4, 0xe8, 0xd9, 0xfb,
	0x32, 0x9f, 0x32, 0x8f, 0x3a, 0x02, 0x89, 0x1f, 0x03, 0xb8, 0xba, 0x6c, 0xa3, 0x22, 0xf2, 0xee,
	0xac, 0x21, 0x80, 0x53, 0x48, 0xb7, 0x1f, 0x64, 0x1c, 0x2c, 0xb8, 0x77, 0x05, 0x87, 0xf2, 0xd0,
	0xc6, 0xef, 0x08, 0x76, 0xce, 0x95, 0x25, 0xe4, 0x31, 0xfe, 0x14, 0x6a, 0xb9, 0x49, 0x14, 0x97,
	0x77, 0xe7, 0x70, 0xd1, 0x3e, 0xca, 0xb3, 0x1d, 0x9d, 0x85, 0x1f, 0xc1, 0xeb, 0xea, 0x25, 0xf3,
	0x94, 0xa4, 0x23, 0x2e, 0xa8, 0xed, 0xcc, 0x2d, 0x23, 0x24, 0x91, 0x8d, 0x7b, 0x02, 0xeb, 0x6c,
	0xfb, 0x85, 0x27, 0xfb, 0x24, 0x63, 0x75, 0x47, 0xed, 0xda, 0x7c, 0x56, 0xb3, 0xe3, 0x1b, 0x3f,
	0x21, 0xd8, 0x9f, 0x0d, 0x39, 0x94, 0xc7, 0x2c, 0xe2, 0x14, 0x7f, 0x09, 0xf5, 0x7c, 0x46, 0xde,
	0x97, 0x6d, 0x84, 0xed, 0xb6, 0x9a, 0xd6, 0x52, 0xd5, 0x5f, 0x2a, 0xf9, 0x86, 0x2e, 0x24, 0x03,
	0xf6, 0x47, 0xd9, 0xa0, 0x2d, 0xb8, 0xbf, 0xf2, 0xa0, 0xf9, 0x54, 0xc6, 0xdf, 0x15, 0xb8, 0x5d,
	0x7e, 0x33, 0x5f, 0x24, 0x0b, 0x48, 0x0c, 0x61, 0x6f, 0x66, 0x83, 0x12, 0x89, 0x57, 0xaf, 0xaa,
	0xb5, 0x94, 0x48, 0xb9, 0x55, 0x67, 0xc3, 0xc1, 0x7e, 0xd9, 0xcf, 0x1c, 0x1a, 0x39, 0x49, 0x25,
	0x56, 0x3f, 0x51, 0x43, 0x28, 0xaf, 0x7e, 0xb8, 0xae, 0x6a, 0x2a, 0xbd, 0xb3, 0xe1, 0xec, 0xd3,
	0xb9, 0x27, 0xf6, 0x93, 0x4c, 0xc6, 0xcf, 0xe1, 0x7c, 0x2d, 0x17, 0x2f, 0xd2, 0xea, 0x74, 0x07,
	0xb6, 0x95, 0x3c, 0xfd, 0x74, 0x12, 0x53, 0xe3, 0x37, 0x04, 0xbb, 0x4f, 0x98, 0x4b, 0xb2, 0x85,
	0xc9, 0x53, 0x38, 0xfe, 0x18, 0x6a, 0xa1, 0x0a, 0x2a, 0x15, 0x0f, 0x17, 0x39, 0x35, 0x4f, 0x76,
	0x74, 0x06, 0x3e, 0x85, 0x4d, 0xed, 0x85, 0x46, 0x45, 0xb8, 0x69, 0xb5, 0x7d, 0x99, 0xa6, 0xd9,
	0xad, 0x8c, 0xb5, 0x09, 0x77, 0xaf, 0x60, 0x5d, 0x1a, 0xdb, 0x78, 0x51, 0x01, 0xdc, 0x0e, 0x47,
	0x3c, 0xa5, 0x49, 0x41, 0x12, 0x7c, 0x04, 0xdb, 0xae, 0x8c, 0xf6, 0x23, 0x72, 0x29, 0xaf, 0xa3,
	0x4d, 0x67, 0x4b, 0xc5, 0x9e, 0x92, 0x4b, 0x8a, 0x3b, 0x7a, 0x3f, 0x85, 0x85, 0xf2, 0xb1, 0xdf,
	0xb9, 0x7a, 0x3f, 0xa5, 0xe2, 0xdb, 0x05, 0xa7, 0x70, 0x4c, 0x00, 0xe7, 0x42, 0xf4, 0xa7, 0x2a,
	0x54, 0x45, 0xb9, 0xe6, 0x52, 0x77, 0x94, 0x38, 0x39, 0xbb, 0xe1, 0xcb, 0xa1, 0x55, 0xee, 0xb5,
	0xb2, 0x0a, 0xc6, 0x3f, 0x08, 0xf6, 0x0a, 0xcf, 0xbd, 0x98, 0xba, 0xc1, 0xb3, 0x80, 0x26, 0x78,
	0x08, 0x6f, 0xe5, 0xf2, 0xcc, 0x6a, 0x20, 0x2f, 0x82, 0xe5, 0xfb, 0x53, 0x6e, 0xe6, 0xbc, 0xe9,
	0x96, 0x62, 0x1c, 0x3f, 0x84, 0x5a, 0x10, 0xa5, 0x34, 0x19, 0x93, 0x50, 0xad, 0xcb, 0xdb, 0xa6,
	0xfc, 0x9c, 0x9a, 0xf9, 0xe7, 0xd4, 0x3c, 0x53, 0x9f, 0x53, 0x47, 0x43, 0xed, 0x0f, 0x32, 0xba,
	0xf7, 0xc1, 0x5a, 0x6d, 0x01, 0x34, 0xaf, 0xe6, 0xf7, 0x55, 0xd8, 0x97, 0x07, 0x67, 0x39, 0xb6,
	0x27, 0x93, 0xf1, 0x8f, 0x08, 0x76, 0x7b, 0x69, 0x42, 0xc9, 0x65, 0xd1, 0x27, 0x8f, 0xff, 0xc3,
	0x4d, 0xb1, 0x68, 0xd1, 0x0e, 0x1e, 0xae, 0x53, 0x4c, 0x0f, 0x6d, 0x6c, 0x1c, 0xa3, 0x13, 0x84,
	0x7f, 0x41, 0x50, 0xff, 0x8c, 0xa6, 0xae, 0xff, 0x7f, 0x1b, 0xef, 0xe4, 0xbb, 0x3f, 0xff, 0xfa,
	0xa1, 0x72, 0x64, 0xdc, 0xd4, 0x7f, 0xa3, 0x74, 0xa2, 0x5d, 0x34, 0x8f, 0x80, 0x55, 0x6d, 0x74,
	0xfb, 0xf4, 0x13, 0xb8, 0x17, 0x30, 0xd9, 0x2c, 0x4e, 0xd8, 0xb7, 0x93, 0x65, 0x7d, 0x4f, 0x6b,
	0x1d, 0x8f, 0x8b, 0x3f, 0x29, 0x5d, 0xf4, 0x1c, 0xa1, 0xc1, 0xab, 0xc2, 0x1a, 0xad, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xee, 0x58, 0x2e, 0x4e, 0x5c, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HealthDiscoveryServiceClient is the client API for HealthDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HealthDiscoveryServiceClient interface {
	StreamHealthCheck(ctx context.Context, opts ...grpc.CallOption) (HealthDiscoveryService_StreamHealthCheckClient, error)
	FetchHealthCheck(ctx context.Context, in *HealthCheckRequestOrEndpointHealthResponse, opts ...grpc.CallOption) (*HealthCheckSpecifier, error)
}

type healthDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewHealthDiscoveryServiceClient(cc *grpc.ClientConn) HealthDiscoveryServiceClient {
	return &healthDiscoveryServiceClient{cc}
}

func (c *healthDiscoveryServiceClient) StreamHealthCheck(ctx context.Context, opts ...grpc.CallOption) (HealthDiscoveryService_StreamHealthCheckClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HealthDiscoveryService_serviceDesc.Streams[0], "/envoy.service.discovery.v3alpha.HealthDiscoveryService/StreamHealthCheck", opts...)
	if err != nil {
		return nil, err
	}
	x := &healthDiscoveryServiceStreamHealthCheckClient{stream}
	return x, nil
}

type HealthDiscoveryService_StreamHealthCheckClient interface {
	Send(*HealthCheckRequestOrEndpointHealthResponse) error
	Recv() (*HealthCheckSpecifier, error)
	grpc.ClientStream
}

type healthDiscoveryServiceStreamHealthCheckClient struct {
	grpc.ClientStream
}

func (x *healthDiscoveryServiceStreamHealthCheckClient) Send(m *HealthCheckRequestOrEndpointHealthResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *healthDiscoveryServiceStreamHealthCheckClient) Recv() (*HealthCheckSpecifier, error) {
	m := new(HealthCheckSpecifier)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *healthDiscoveryServiceClient) FetchHealthCheck(ctx context.Context, in *HealthCheckRequestOrEndpointHealthResponse, opts ...grpc.CallOption) (*HealthCheckSpecifier, error) {
	out := new(HealthCheckSpecifier)
	err := c.cc.Invoke(ctx, "/envoy.service.discovery.v3alpha.HealthDiscoveryService/FetchHealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthDiscoveryServiceServer is the server API for HealthDiscoveryService service.
type HealthDiscoveryServiceServer interface {
	StreamHealthCheck(HealthDiscoveryService_StreamHealthCheckServer) error
	FetchHealthCheck(context.Context, *HealthCheckRequestOrEndpointHealthResponse) (*HealthCheckSpecifier, error)
}

// UnimplementedHealthDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHealthDiscoveryServiceServer struct {
}

func (*UnimplementedHealthDiscoveryServiceServer) StreamHealthCheck(srv HealthDiscoveryService_StreamHealthCheckServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamHealthCheck not implemented")
}
func (*UnimplementedHealthDiscoveryServiceServer) FetchHealthCheck(ctx context.Context, req *HealthCheckRequestOrEndpointHealthResponse) (*HealthCheckSpecifier, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchHealthCheck not implemented")
}

func RegisterHealthDiscoveryServiceServer(s *grpc.Server, srv HealthDiscoveryServiceServer) {
	s.RegisterService(&_HealthDiscoveryService_serviceDesc, srv)
}

func _HealthDiscoveryService_StreamHealthCheck_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HealthDiscoveryServiceServer).StreamHealthCheck(&healthDiscoveryServiceStreamHealthCheckServer{stream})
}

type HealthDiscoveryService_StreamHealthCheckServer interface {
	Send(*HealthCheckSpecifier) error
	Recv() (*HealthCheckRequestOrEndpointHealthResponse, error)
	grpc.ServerStream
}

type healthDiscoveryServiceStreamHealthCheckServer struct {
	grpc.ServerStream
}

func (x *healthDiscoveryServiceStreamHealthCheckServer) Send(m *HealthCheckSpecifier) error {
	return x.ServerStream.SendMsg(m)
}

func (x *healthDiscoveryServiceStreamHealthCheckServer) Recv() (*HealthCheckRequestOrEndpointHealthResponse, error) {
	m := new(HealthCheckRequestOrEndpointHealthResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _HealthDiscoveryService_FetchHealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequestOrEndpointHealthResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthDiscoveryServiceServer).FetchHealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.discovery.v3alpha.HealthDiscoveryService/FetchHealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthDiscoveryServiceServer).FetchHealthCheck(ctx, req.(*HealthCheckRequestOrEndpointHealthResponse))
	}
	return interceptor(ctx, in, info, handler)
}

var _HealthDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.discovery.v3alpha.HealthDiscoveryService",
	HandlerType: (*HealthDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchHealthCheck",
			Handler:    _HealthDiscoveryService_FetchHealthCheck_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamHealthCheck",
			Handler:       _HealthDiscoveryService_StreamHealthCheck_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/discovery/v3alpha/hds.proto",
}
