// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v3alpha/eds.proto

package envoy_api_v3alpha

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/endpoint"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/type/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type ClusterLoadAssignment struct {
	ClusterName          string                          `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	Endpoints            []*endpoint.LocalityLbEndpoints `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	NamedEndpoints       map[string]*endpoint.Endpoint   `protobuf:"bytes,5,rep,name=named_endpoints,json=namedEndpoints,proto3" json:"named_endpoints,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Policy               *ClusterLoadAssignment_Policy   `protobuf:"bytes,4,opt,name=policy,proto3" json:"policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ClusterLoadAssignment) Reset()         { *m = ClusterLoadAssignment{} }
func (m *ClusterLoadAssignment) String() string { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment) ProtoMessage()    {}
func (*ClusterLoadAssignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_35014a248c9ef3f8, []int{0}
}

func (m *ClusterLoadAssignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterLoadAssignment.Unmarshal(m, b)
}
func (m *ClusterLoadAssignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterLoadAssignment.Marshal(b, m, deterministic)
}
func (m *ClusterLoadAssignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterLoadAssignment.Merge(m, src)
}
func (m *ClusterLoadAssignment) XXX_Size() int {
	return xxx_messageInfo_ClusterLoadAssignment.Size(m)
}
func (m *ClusterLoadAssignment) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterLoadAssignment.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterLoadAssignment proto.InternalMessageInfo

func (m *ClusterLoadAssignment) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *ClusterLoadAssignment) GetEndpoints() []*endpoint.LocalityLbEndpoints {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *ClusterLoadAssignment) GetNamedEndpoints() map[string]*endpoint.Endpoint {
	if m != nil {
		return m.NamedEndpoints
	}
	return nil
}

func (m *ClusterLoadAssignment) GetPolicy() *ClusterLoadAssignment_Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

type ClusterLoadAssignment_Policy struct {
	DropOverloads           []*ClusterLoadAssignment_Policy_DropOverload `protobuf:"bytes,2,rep,name=drop_overloads,json=dropOverloads,proto3" json:"drop_overloads,omitempty"`
	OverprovisioningFactor  *wrappers.UInt32Value                        `protobuf:"bytes,3,opt,name=overprovisioning_factor,json=overprovisioningFactor,proto3" json:"overprovisioning_factor,omitempty"`
	EndpointStaleAfter      *duration.Duration                           `protobuf:"bytes,4,opt,name=endpoint_stale_after,json=endpointStaleAfter,proto3" json:"endpoint_stale_after,omitempty"`
	DisableOverprovisioning bool                                         `protobuf:"varint,5,opt,name=disable_overprovisioning,json=disableOverprovisioning,proto3" json:"disable_overprovisioning,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                                     `json:"-"`
	XXX_unrecognized        []byte                                       `json:"-"`
	XXX_sizecache           int32                                        `json:"-"`
}

func (m *ClusterLoadAssignment_Policy) Reset()         { *m = ClusterLoadAssignment_Policy{} }
func (m *ClusterLoadAssignment_Policy) String() string { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment_Policy) ProtoMessage()    {}
func (*ClusterLoadAssignment_Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_35014a248c9ef3f8, []int{0, 0}
}

func (m *ClusterLoadAssignment_Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterLoadAssignment_Policy.Unmarshal(m, b)
}
func (m *ClusterLoadAssignment_Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterLoadAssignment_Policy.Marshal(b, m, deterministic)
}
func (m *ClusterLoadAssignment_Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterLoadAssignment_Policy.Merge(m, src)
}
func (m *ClusterLoadAssignment_Policy) XXX_Size() int {
	return xxx_messageInfo_ClusterLoadAssignment_Policy.Size(m)
}
func (m *ClusterLoadAssignment_Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterLoadAssignment_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterLoadAssignment_Policy proto.InternalMessageInfo

func (m *ClusterLoadAssignment_Policy) GetDropOverloads() []*ClusterLoadAssignment_Policy_DropOverload {
	if m != nil {
		return m.DropOverloads
	}
	return nil
}

func (m *ClusterLoadAssignment_Policy) GetOverprovisioningFactor() *wrappers.UInt32Value {
	if m != nil {
		return m.OverprovisioningFactor
	}
	return nil
}

func (m *ClusterLoadAssignment_Policy) GetEndpointStaleAfter() *duration.Duration {
	if m != nil {
		return m.EndpointStaleAfter
	}
	return nil
}

func (m *ClusterLoadAssignment_Policy) GetDisableOverprovisioning() bool {
	if m != nil {
		return m.DisableOverprovisioning
	}
	return false
}

type ClusterLoadAssignment_Policy_DropOverload struct {
	Category             string                     `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	DropPercentage       *v3alpha.FractionalPercent `protobuf:"bytes,2,opt,name=drop_percentage,json=dropPercentage,proto3" json:"drop_percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ClusterLoadAssignment_Policy_DropOverload) Reset() {
	*m = ClusterLoadAssignment_Policy_DropOverload{}
}
func (m *ClusterLoadAssignment_Policy_DropOverload) String() string { return proto.CompactTextString(m) }
func (*ClusterLoadAssignment_Policy_DropOverload) ProtoMessage()    {}
func (*ClusterLoadAssignment_Policy_DropOverload) Descriptor() ([]byte, []int) {
	return fileDescriptor_35014a248c9ef3f8, []int{0, 0, 0}
}

func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Unmarshal(m, b)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Marshal(b, m, deterministic)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Merge(m, src)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_Size() int {
	return xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.Size(m)
}
func (m *ClusterLoadAssignment_Policy_DropOverload) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterLoadAssignment_Policy_DropOverload proto.InternalMessageInfo

func (m *ClusterLoadAssignment_Policy_DropOverload) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *ClusterLoadAssignment_Policy_DropOverload) GetDropPercentage() *v3alpha.FractionalPercent {
	if m != nil {
		return m.DropPercentage
	}
	return nil
}

func init() {
	proto.RegisterType((*ClusterLoadAssignment)(nil), "envoy.api.v3alpha.ClusterLoadAssignment")
	proto.RegisterMapType((map[string]*endpoint.Endpoint)(nil), "envoy.api.v3alpha.ClusterLoadAssignment.NamedEndpointsEntry")
	proto.RegisterType((*ClusterLoadAssignment_Policy)(nil), "envoy.api.v3alpha.ClusterLoadAssignment.Policy")
	proto.RegisterType((*ClusterLoadAssignment_Policy_DropOverload)(nil), "envoy.api.v3alpha.ClusterLoadAssignment.Policy.DropOverload")
}

func init() { proto.RegisterFile("envoy/api/v3alpha/eds.proto", fileDescriptor_35014a248c9ef3f8) }

var fileDescriptor_35014a248c9ef3f8 = []byte{
	// 764 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x4d, 0x53, 0xd3, 0x4c,
	0x1c, 0x67, 0xd3, 0x96, 0xa7, 0x2c, 0x3c, 0x85, 0x67, 0x1f, 0x95, 0x58, 0x19, 0x2c, 0x05, 0x67,
	0xda, 0x3a, 0xa6, 0x4c, 0x99, 0x71, 0xb0, 0x83, 0x07, 0x6a, 0xc1, 0xd1, 0x41, 0xe8, 0x84, 0xc1,
	0xa3, 0x9d, 0x6d, 0xb2, 0x94, 0x0c, 0x69, 0x76, 0xdd, 0x6c, 0xa3, 0xb9, 0x7a, 0x62, 0xbc, 0x7a,
	0xf3, 0x23, 0x78, 0xf7, 0x0b, 0x38, 0xe3, 0x27, 0xf0, 0xea, 0xd1, 0x0f, 0xe1, 0x70, 0x72, 0x92,
	0x6c, 0xd2, 0xd2, 0x16, 0xc4, 0x83, 0xb7, 0x34, 0xbf, 0x97, 0xff, 0x7b, 0x0a, 0xef, 0x10, 0xc7,
	0xa3, 0x7e, 0x15, 0x33, 0xab, 0xea, 0x6d, 0x60, 0x9b, 0x9d, 0xe0, 0x2a, 0x31, 0x5d, 0x8d, 0x71,
	0x2a, 0x28, 0xfa, 0x2f, 0x04, 0x35, 0xcc, 0x2c, 0x4d, 0x82, 0xf9, 0x95, 0x71, 0xbe, 0x69, 0xb9,
	0x06, 0xf5, 0x08, 0xf7, 0x23, 0x55, 0xbe, 0x3c, 0xc1, 0xd2, 0x31, 0x19, 0xb5, 0x1c, 0x91, 0x3c,
	0x48, 0x6a, 0x21, 0xa2, 0x0a, 0x9f, 0x91, 0x84, 0xcb, 0x08, 0x37, 0x48, 0xc2, 0x58, 0xea, 0x52,
	0xda, 0xb5, 0x49, 0xe8, 0x86, 0x1d, 0x87, 0x0a, 0x2c, 0x2c, 0xea, 0xc8, 0x04, 0xf3, 0xcb, 0x12,
	0x0d, 0x7f, 0x75, 0xfa, 0xc7, 0x55, 0xb3, 0xcf, 0x43, 0xc2, 0x65, 0xf8, 0x1b, 0x8e, 0x19, 0x23,
	0x3c, 0xd6, 0xaf, 0xf4, 0x4d, 0x86, 0x87, 0x7d, 0xab, 0x1e, 0xe1, 0xae, 0x45, 0x1d, 0xcb, 0xe9,
	0x4a, 0xca, 0xa2, 0x87, 0x6d, 0xcb, 0xc4, 0x82, 0x54, 0xe3, 0x87, 0x08, 0x28, 0x7e, 0xce, 0xc2,
	0x9b, 0x4f, 0xec, 0xbe, 0x2b, 0x08, 0xdf, 0xa3, 0xd8, 0xdc, 0x76, 0x5d, 0xab, 0xeb, 0xf4, 0x88,
	0x23, 0x50, 0x05, 0xce, 0x19, 0x11, 0xd0, 0x76, 0x70, 0x8f, 0xa8, 0xa0, 0x00, 0x4a, 0x33, 0x8d,
	0x7f, 0xce, 0x1b, 0x69, 0xae, 0x14, 0x80, 0x3e, 0x2b, 0xc1, 0x7d, 0xdc, 0x23, 0xe8, 0x05, 0x9c,
	0x89, 0x7b, 0xe2, 0xaa, 0x4a, 0x21, 0x55, 0x9a, 0xad, 0x55, 0xb5, 0xb1, 0xb6, 0x6b, 0x49, 0xdf,
	0xf6, 0xa8, 0x81, 0x6d, 0x4b, 0xf8, 0x7b, 0x9d, 0x9d, 0x58, 0xa6, 0x0f, 0x1c, 0x10, 0x81, 0xf3,
	0x41, 0x48, 0xb3, 0x3d, 0x30, 0xcd, 0x84, 0xa6, 0x5b, 0x13, 0x4c, 0x27, 0x66, 0xaf, 0x05, 0x69,
	0x99, 0x89, 0xf9, 0x8e, 0x23, 0xb8, 0xaf, 0xe7, 0x9c, 0x0b, 0x2f, 0xd1, 0x53, 0x38, 0xcd, 0xa8,
	0x6d, 0x19, 0xbe, 0x9a, 0x2e, 0x80, 0x4b, 0x52, 0x9e, 0xec, 0xde, 0x0a, 0x65, 0xba, 0x94, 0xe7,
	0xbf, 0xa7, 0xe1, 0x74, 0xf4, 0x0a, 0x19, 0x30, 0x67, 0x72, 0xca, 0xda, 0xc1, 0x2a, 0xd9, 0x14,
	0x9b, 0x71, 0x3b, 0xb6, 0xfe, 0xd0, 0x5b, 0x6b, 0x72, 0xca, 0x0e, 0xa4, 0x89, 0xfe, 0xaf, 0x39,
	0xf4, 0xcb, 0x45, 0xaf, 0xe0, 0x62, 0xe0, 0xcf, 0x38, 0xf5, 0x2c, 0x39, 0xe7, 0xf6, 0x31, 0x36,
	0x04, 0xe5, 0x6a, 0x2a, 0xac, 0x64, 0x49, 0x8b, 0x56, 0x46, 0x8b, 0x57, 0x46, 0x3b, 0x7a, 0xe6,
	0x88, 0x8d, 0xda, 0x4b, 0x6c, 0xf7, 0x49, 0x38, 0xc3, 0x8a, 0x52, 0x98, 0xd2, 0x6f, 0x8d, 0xba,
	0xec, 0x86, 0x26, 0xe8, 0x08, 0xde, 0x88, 0x3b, 0xdf, 0x76, 0x05, 0xb6, 0x49, 0x1b, 0x1f, 0x0b,
	0xc2, 0x65, 0x9b, 0x6e, 0x8f, 0x99, 0x37, 0xe5, 0xbe, 0x36, 0xb2, 0xe7, 0x8d, 0xcc, 0x27, 0xa0,
	0x54, 0xa6, 0x74, 0x14, 0x1b, 0x1c, 0x06, 0xfa, 0xed, 0x40, 0x8e, 0x1e, 0x41, 0xd5, 0xb4, 0x5c,
	0xdc, 0xb1, 0x49, 0x7b, 0x34, 0xb0, 0x9a, 0x29, 0x80, 0x52, 0x56, 0x5f, 0x94, 0xf8, 0xc1, 0x08,
	0x9c, 0xff, 0x02, 0xe0, 0xdc, 0x70, 0x47, 0xd0, 0x2a, 0xcc, 0x1a, 0x58, 0x90, 0x2e, 0xe5, 0xfe,
	0xe8, 0x66, 0x26, 0x00, 0xda, 0x87, 0xf3, 0xe1, 0x30, 0xe4, 0x31, 0xe2, 0x2e, 0x51, 0x95, 0xb0,
	0x84, 0x7b, 0x72, 0x1a, 0xc1, 0xc9, 0x26, 0xe3, 0xd8, 0xe5, 0xd8, 0x08, 0xaa, 0xc0, 0x76, 0x2b,
	0xe2, 0xeb, 0xe1, 0x28, 0x5b, 0x89, 0xb8, 0xfe, 0xf8, 0xe3, 0xd7, 0xb3, 0xe5, 0x4d, 0xf8, 0x70,
	0x68, 0x94, 0xb5, 0xeb, 0x4f, 0xb1, 0xbe, 0x1e, 0xc8, 0xef, 0xc3, 0xf2, 0xb5, 0xe5, 0xcf, 0xd3,
	0x59, 0xb0, 0xa0, 0xe4, 0xbb, 0xf0, 0xff, 0x09, 0xeb, 0x8c, 0x16, 0x60, 0xea, 0x94, 0xc8, 0xea,
	0xf5, 0xe0, 0x11, 0xd5, 0x61, 0xc6, 0x0b, 0x26, 0x2c, 0xab, 0x5c, 0xbb, 0xea, 0x04, 0x63, 0x33,
	0x3d, 0x92, 0xd4, 0x95, 0x4d, 0x50, 0x2f, 0x07, 0x09, 0xae, 0xc1, 0xe2, 0xef, 0x13, 0xac, 0xfd,
	0x54, 0xa0, 0x1a, 0x5b, 0x34, 0xe3, 0x4f, 0xe7, 0x21, 0xe1, 0x9e, 0x65, 0x10, 0xd4, 0x81, 0xf3,
	0x87, 0x82, 0x13, 0xdc, 0x1b, 0xdc, 0xda, 0xea, 0x84, 0x5c, 0x12, 0x9d, 0x4e, 0x5e, 0xf7, 0x89,
	0x2b, 0xf2, 0x6b, 0x57, 0x93, 0x5c, 0x46, 0x1d, 0x97, 0x14, 0xa7, 0x4a, 0x60, 0x1d, 0xa0, 0x53,
	0x98, 0x6b, 0x12, 0x5b, 0xe0, 0x41, 0x88, 0xd2, 0x24, 0x75, 0x40, 0x19, 0x8b, 0x53, 0xbe, 0x06,
	0xf3, 0x42, 0xb0, 0xf7, 0x00, 0xe6, 0x76, 0x89, 0x30, 0x4e, 0xfe, 0x4a, 0x41, 0x0f, 0xde, 0x7d,
	0xfb, 0xf1, 0x41, 0x59, 0x2e, 0x2e, 0x8d, 0xff, 0x19, 0xd5, 0x93, 0x2f, 0x60, 0xc8, 0x49, 0xd5,
	0x41, 0xa5, 0xb1, 0x0e, 0xef, 0x5a, 0x34, 0x32, 0x66, 0x9c, 0xbe, 0xf5, 0xc7, 0x63, 0x34, 0xb2,
	0x3b, 0xa6, 0xdb, 0x0a, 0xae, 0xb3, 0x05, 0xce, 0x00, 0xe8, 0x4c, 0x87, 0x97, 0xba, 0xf1, 0x2b,
	0x00, 0x00, 0xff, 0xff, 0x06, 0x39, 0x67, 0xee, 0x27, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EndpointDiscoveryServiceClient is the client API for EndpointDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EndpointDiscoveryServiceClient interface {
	StreamEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_StreamEndpointsClient, error)
	DeltaEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_DeltaEndpointsClient, error)
	FetchEndpoints(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
}

type endpointDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewEndpointDiscoveryServiceClient(cc *grpc.ClientConn) EndpointDiscoveryServiceClient {
	return &endpointDiscoveryServiceClient{cc}
}

func (c *endpointDiscoveryServiceClient) StreamEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_StreamEndpointsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EndpointDiscoveryService_serviceDesc.Streams[0], "/envoy.api.v3alpha.EndpointDiscoveryService/StreamEndpoints", opts...)
	if err != nil {
		return nil, err
	}
	x := &endpointDiscoveryServiceStreamEndpointsClient{stream}
	return x, nil
}

type EndpointDiscoveryService_StreamEndpointsClient interface {
	Send(*DiscoveryRequest) error
	Recv() (*DiscoveryResponse, error)
	grpc.ClientStream
}

type endpointDiscoveryServiceStreamEndpointsClient struct {
	grpc.ClientStream
}

func (x *endpointDiscoveryServiceStreamEndpointsClient) Send(m *DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceStreamEndpointsClient) Recv() (*DiscoveryResponse, error) {
	m := new(DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *endpointDiscoveryServiceClient) DeltaEndpoints(ctx context.Context, opts ...grpc.CallOption) (EndpointDiscoveryService_DeltaEndpointsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EndpointDiscoveryService_serviceDesc.Streams[1], "/envoy.api.v3alpha.EndpointDiscoveryService/DeltaEndpoints", opts...)
	if err != nil {
		return nil, err
	}
	x := &endpointDiscoveryServiceDeltaEndpointsClient{stream}
	return x, nil
}

type EndpointDiscoveryService_DeltaEndpointsClient interface {
	Send(*DeltaDiscoveryRequest) error
	Recv() (*DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type endpointDiscoveryServiceDeltaEndpointsClient struct {
	grpc.ClientStream
}

func (x *endpointDiscoveryServiceDeltaEndpointsClient) Send(m *DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceDeltaEndpointsClient) Recv() (*DeltaDiscoveryResponse, error) {
	m := new(DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *endpointDiscoveryServiceClient) FetchEndpoints(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.api.v3alpha.EndpointDiscoveryService/FetchEndpoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndpointDiscoveryServiceServer is the server API for EndpointDiscoveryService service.
type EndpointDiscoveryServiceServer interface {
	StreamEndpoints(EndpointDiscoveryService_StreamEndpointsServer) error
	DeltaEndpoints(EndpointDiscoveryService_DeltaEndpointsServer) error
	FetchEndpoints(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
}

// UnimplementedEndpointDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedEndpointDiscoveryServiceServer struct {
}

func (*UnimplementedEndpointDiscoveryServiceServer) StreamEndpoints(srv EndpointDiscoveryService_StreamEndpointsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamEndpoints not implemented")
}
func (*UnimplementedEndpointDiscoveryServiceServer) DeltaEndpoints(srv EndpointDiscoveryService_DeltaEndpointsServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaEndpoints not implemented")
}
func (*UnimplementedEndpointDiscoveryServiceServer) FetchEndpoints(ctx context.Context, req *DiscoveryRequest) (*DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchEndpoints not implemented")
}

func RegisterEndpointDiscoveryServiceServer(s *grpc.Server, srv EndpointDiscoveryServiceServer) {
	s.RegisterService(&_EndpointDiscoveryService_serviceDesc, srv)
}

func _EndpointDiscoveryService_StreamEndpoints_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EndpointDiscoveryServiceServer).StreamEndpoints(&endpointDiscoveryServiceStreamEndpointsServer{stream})
}

type EndpointDiscoveryService_StreamEndpointsServer interface {
	Send(*DiscoveryResponse) error
	Recv() (*DiscoveryRequest, error)
	grpc.ServerStream
}

type endpointDiscoveryServiceStreamEndpointsServer struct {
	grpc.ServerStream
}

func (x *endpointDiscoveryServiceStreamEndpointsServer) Send(m *DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceStreamEndpointsServer) Recv() (*DiscoveryRequest, error) {
	m := new(DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EndpointDiscoveryService_DeltaEndpoints_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EndpointDiscoveryServiceServer).DeltaEndpoints(&endpointDiscoveryServiceDeltaEndpointsServer{stream})
}

type EndpointDiscoveryService_DeltaEndpointsServer interface {
	Send(*DeltaDiscoveryResponse) error
	Recv() (*DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type endpointDiscoveryServiceDeltaEndpointsServer struct {
	grpc.ServerStream
}

func (x *endpointDiscoveryServiceDeltaEndpointsServer) Send(m *DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *endpointDiscoveryServiceDeltaEndpointsServer) Recv() (*DeltaDiscoveryRequest, error) {
	m := new(DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EndpointDiscoveryService_FetchEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointDiscoveryServiceServer).FetchEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.api.v3alpha.EndpointDiscoveryService/FetchEndpoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointDiscoveryServiceServer).FetchEndpoints(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EndpointDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.api.v3alpha.EndpointDiscoveryService",
	HandlerType: (*EndpointDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchEndpoints",
			Handler:    _EndpointDiscoveryService_FetchEndpoints_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamEndpoints",
			Handler:       _EndpointDiscoveryService_StreamEndpoints_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaEndpoints",
			Handler:       _EndpointDiscoveryService_DeltaEndpoints_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/api/v3alpha/eds.proto",
}
