// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/rbac/v3alpha/rbac.proto

package envoy_config_filter_network_rbac_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/config/rbac/v3alpha"
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

type RBAC_EnforcementType int32

const (
	RBAC_ONE_TIME_ON_FIRST_BYTE RBAC_EnforcementType = 0
	RBAC_CONTINUOUS             RBAC_EnforcementType = 1
)

var RBAC_EnforcementType_name = map[int32]string{
	0: "ONE_TIME_ON_FIRST_BYTE",
	1: "CONTINUOUS",
}

var RBAC_EnforcementType_value = map[string]int32{
	"ONE_TIME_ON_FIRST_BYTE": 0,
	"CONTINUOUS":             1,
}

func (x RBAC_EnforcementType) String() string {
	return proto.EnumName(RBAC_EnforcementType_name, int32(x))
}

func (RBAC_EnforcementType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6858d3a47bb4f89e, []int{0, 0}
}

type RBAC struct {
	Rules                *v3alpha.RBAC        `protobuf:"bytes,1,opt,name=rules,proto3" json:"rules,omitempty"`
	ShadowRules          *v3alpha.RBAC        `protobuf:"bytes,2,opt,name=shadow_rules,json=shadowRules,proto3" json:"shadow_rules,omitempty"`
	StatPrefix           string               `protobuf:"bytes,3,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	EnforcementType      RBAC_EnforcementType `protobuf:"varint,4,opt,name=enforcement_type,json=enforcementType,proto3,enum=envoy.config.filter.network.rbac.v3alpha.RBAC_EnforcementType" json:"enforcement_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RBAC) Reset()         { *m = RBAC{} }
func (m *RBAC) String() string { return proto.CompactTextString(m) }
func (*RBAC) ProtoMessage()    {}
func (*RBAC) Descriptor() ([]byte, []int) {
	return fileDescriptor_6858d3a47bb4f89e, []int{0}
}

func (m *RBAC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RBAC.Unmarshal(m, b)
}
func (m *RBAC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RBAC.Marshal(b, m, deterministic)
}
func (m *RBAC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RBAC.Merge(m, src)
}
func (m *RBAC) XXX_Size() int {
	return xxx_messageInfo_RBAC.Size(m)
}
func (m *RBAC) XXX_DiscardUnknown() {
	xxx_messageInfo_RBAC.DiscardUnknown(m)
}

var xxx_messageInfo_RBAC proto.InternalMessageInfo

func (m *RBAC) GetRules() *v3alpha.RBAC {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *RBAC) GetShadowRules() *v3alpha.RBAC {
	if m != nil {
		return m.ShadowRules
	}
	return nil
}

func (m *RBAC) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *RBAC) GetEnforcementType() RBAC_EnforcementType {
	if m != nil {
		return m.EnforcementType
	}
	return RBAC_ONE_TIME_ON_FIRST_BYTE
}

func init() {
	proto.RegisterEnum("envoy.config.filter.network.rbac.v3alpha.RBAC_EnforcementType", RBAC_EnforcementType_name, RBAC_EnforcementType_value)
	proto.RegisterType((*RBAC)(nil), "envoy.config.filter.network.rbac.v3alpha.RBAC")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/rbac/v3alpha/rbac.proto", fileDescriptor_6858d3a47bb4f89e)
}

var fileDescriptor_6858d3a47bb4f89e = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x41, 0x4f, 0xf2, 0x30,
	0x1c, 0xc6, 0xdf, 0x01, 0xef, 0xfb, 0x86, 0x62, 0x80, 0xec, 0xa0, 0x84, 0x83, 0x4e, 0xe2, 0x61,
	0xf1, 0xd0, 0x26, 0x10, 0x3d, 0x98, 0x68, 0xe2, 0xc8, 0x4c, 0x38, 0xb8, 0x91, 0x31, 0x0e, 0x9e,
	0x96, 0x32, 0x3a, 0x68, 0x9c, 0xed, 0xd2, 0x95, 0xc1, 0xbe, 0x80, 0xf1, 0x33, 0xf8, 0x7d, 0xfc,
	0x52, 0x9e, 0x0c, 0x1d, 0x46, 0x31, 0x31, 0xd9, 0xad, 0xed, 0xff, 0xf9, 0x3d, 0x79, 0xfe, 0x7d,
	0xc0, 0x80, 0xb0, 0x8c, 0xe7, 0x28, 0xe4, 0x2c, 0xa2, 0x0b, 0x14, 0xd1, 0x58, 0x12, 0x81, 0x18,
	0x91, 0x6b, 0x2e, 0x1e, 0x91, 0x98, 0xe1, 0x10, 0x65, 0x03, 0x1c, 0x27, 0x4b, 0xac, 0x2e, 0x30,
	0x11, 0x5c, 0x72, 0xdd, 0x54, 0x10, 0x2c, 0x20, 0x58, 0x40, 0x70, 0x07, 0x41, 0xa5, 0xdb, 0x41,
	0xdd, 0xb3, 0x3d, 0xfb, 0x5f, 0xfc, 0xba, 0xa7, 0xab, 0x79, 0x82, 0x11, 0x66, 0x8c, 0x4b, 0x2c,
	0x29, 0x67, 0x29, 0xca, 0x88, 0x48, 0x29, 0x67, 0x94, 0x2d, 0x76, 0x92, 0xa3, 0x0c, 0xc7, 0x74,
	0x8e, 0x25, 0x41, 0x9f, 0x87, 0x62, 0xd0, 0x7b, 0xae, 0x82, 0x9a, 0x67, 0xdd, 0x0e, 0xf5, 0x0b,
	0xf0, 0x57, 0xac, 0x62, 0x92, 0x76, 0x34, 0x43, 0x33, 0x1b, 0xfd, 0x13, 0xb8, 0x17, 0xf2, 0x7b,
	0x2a, 0xb8, 0xd5, 0x7b, 0x85, 0x5a, 0xb7, 0xc0, 0x41, 0xba, 0xc4, 0x73, 0xbe, 0x0e, 0x0a, 0xba,
	0x52, 0x8e, 0x6e, 0x14, 0x90, 0xa7, 0x3c, 0x4c, 0xd0, 0x48, 0x25, 0x96, 0x41, 0x22, 0x48, 0x44,
	0x37, 0x9d, 0xaa, 0xa1, 0x99, 0x75, 0xeb, 0xff, 0xbb, 0x55, 0x13, 0x15, 0x43, 0xf3, 0xc0, 0x76,
	0x36, 0x56, 0x23, 0x9d, 0x82, 0x36, 0x61, 0x11, 0x17, 0x21, 0x79, 0x22, 0x4c, 0x06, 0x32, 0x4f,
	0x48, 0xa7, 0x66, 0x68, 0x66, 0xb3, 0x7f, 0x03, 0xcb, 0x7e, 0xaa, 0x0a, 0x00, 0xed, 0x2f, 0x1b,
	0x3f, 0x4f, 0x88, 0xd7, 0x22, 0xfb, 0x0f, 0xbd, 0x6b, 0xd0, 0xfa, 0xa1, 0xd1, 0xbb, 0xe0, 0xd0,
	0x75, 0xec, 0xc0, 0x1f, 0xdd, 0xdb, 0x81, 0xeb, 0x04, 0x77, 0x23, 0x6f, 0xe2, 0x07, 0xd6, 0x83,
	0x6f, 0xb7, 0xff, 0xe8, 0x4d, 0x00, 0x86, 0xae, 0xe3, 0x8f, 0x9c, 0xa9, 0x3b, 0x9d, 0xb4, 0xb5,
	0x2b, 0xf4, 0xfa, 0xf6, 0x72, 0x7c, 0x0e, 0x4a, 0x54, 0xdd, 0x57, 0x81, 0x2c, 0x1b, 0x5c, 0x52,
	0x5e, 0x2c, 0x91, 0x08, 0xbe, 0xc9, 0x4b, 0xef, 0x63, 0xd5, 0xbd, 0x19, 0x0e, 0xc7, 0xdb, 0x36,
	0xc7, 0xda, 0xec, 0x9f, 0xaa, 0x75, 0xf0, 0x11, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x2a, 0xbb, 0x13,
	0x99, 0x02, 0x00, 0x00,
}
