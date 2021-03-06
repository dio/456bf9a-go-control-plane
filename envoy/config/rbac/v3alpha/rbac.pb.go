// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/rbac/v3alpha/rbac.proto

package envoy_config_rbac_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/core"
	route "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/route"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	v1alpha1 "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
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

type RBAC_Action int32

const (
	RBAC_ALLOW RBAC_Action = 0
	RBAC_DENY  RBAC_Action = 1
)

var RBAC_Action_name = map[int32]string{
	0: "ALLOW",
	1: "DENY",
}

var RBAC_Action_value = map[string]int32{
	"ALLOW": 0,
	"DENY":  1,
}

func (x RBAC_Action) String() string {
	return proto.EnumName(RBAC_Action_name, int32(x))
}

func (RBAC_Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eb7d22ec66d90520, []int{0, 0}
}

type RBAC struct {
	Action               RBAC_Action        `protobuf:"varint,1,opt,name=action,proto3,enum=envoy.config.rbac.v3alpha.RBAC_Action" json:"action,omitempty"`
	Policies             map[string]*Policy `protobuf:"bytes,2,rep,name=policies,proto3" json:"policies,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RBAC) Reset()         { *m = RBAC{} }
func (m *RBAC) String() string { return proto.CompactTextString(m) }
func (*RBAC) ProtoMessage()    {}
func (*RBAC) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb7d22ec66d90520, []int{0}
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

func (m *RBAC) GetAction() RBAC_Action {
	if m != nil {
		return m.Action
	}
	return RBAC_ALLOW
}

func (m *RBAC) GetPolicies() map[string]*Policy {
	if m != nil {
		return m.Policies
	}
	return nil
}

type Policy struct {
	Permissions          []*Permission  `protobuf:"bytes,1,rep,name=permissions,proto3" json:"permissions,omitempty"`
	Principals           []*Principal   `protobuf:"bytes,2,rep,name=principals,proto3" json:"principals,omitempty"`
	Condition            *v1alpha1.Expr `protobuf:"bytes,3,opt,name=condition,proto3" json:"condition,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb7d22ec66d90520, []int{1}
}

func (m *Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy.Unmarshal(m, b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
}
func (m *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(m, src)
}
func (m *Policy) XXX_Size() int {
	return xxx_messageInfo_Policy.Size(m)
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetPermissions() []*Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *Policy) GetPrincipals() []*Principal {
	if m != nil {
		return m.Principals
	}
	return nil
}

func (m *Policy) GetCondition() *v1alpha1.Expr {
	if m != nil {
		return m.Condition
	}
	return nil
}

type Permission struct {
	// Types that are valid to be assigned to Rule:
	//	*Permission_AndRules
	//	*Permission_OrRules
	//	*Permission_Any
	//	*Permission_Header
	//	*Permission_DestinationIp
	//	*Permission_DestinationPort
	//	*Permission_Metadata
	//	*Permission_NotRule
	//	*Permission_RequestedServerName
	Rule                 isPermission_Rule `protobuf_oneof:"rule"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Permission) Reset()         { *m = Permission{} }
func (m *Permission) String() string { return proto.CompactTextString(m) }
func (*Permission) ProtoMessage()    {}
func (*Permission) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb7d22ec66d90520, []int{2}
}

func (m *Permission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permission.Unmarshal(m, b)
}
func (m *Permission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permission.Marshal(b, m, deterministic)
}
func (m *Permission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permission.Merge(m, src)
}
func (m *Permission) XXX_Size() int {
	return xxx_messageInfo_Permission.Size(m)
}
func (m *Permission) XXX_DiscardUnknown() {
	xxx_messageInfo_Permission.DiscardUnknown(m)
}

var xxx_messageInfo_Permission proto.InternalMessageInfo

type isPermission_Rule interface {
	isPermission_Rule()
}

type Permission_AndRules struct {
	AndRules *Permission_Set `protobuf:"bytes,1,opt,name=and_rules,json=andRules,proto3,oneof"`
}

type Permission_OrRules struct {
	OrRules *Permission_Set `protobuf:"bytes,2,opt,name=or_rules,json=orRules,proto3,oneof"`
}

type Permission_Any struct {
	Any bool `protobuf:"varint,3,opt,name=any,proto3,oneof"`
}

type Permission_Header struct {
	Header *route.HeaderMatcher `protobuf:"bytes,4,opt,name=header,proto3,oneof"`
}

type Permission_DestinationIp struct {
	DestinationIp *core.CidrRange `protobuf:"bytes,5,opt,name=destination_ip,json=destinationIp,proto3,oneof"`
}

type Permission_DestinationPort struct {
	DestinationPort uint32 `protobuf:"varint,6,opt,name=destination_port,json=destinationPort,proto3,oneof"`
}

type Permission_Metadata struct {
	Metadata *v3alpha.MetadataMatcher `protobuf:"bytes,7,opt,name=metadata,proto3,oneof"`
}

type Permission_NotRule struct {
	NotRule *Permission `protobuf:"bytes,8,opt,name=not_rule,json=notRule,proto3,oneof"`
}

type Permission_RequestedServerName struct {
	RequestedServerName *v3alpha.StringMatcher `protobuf:"bytes,9,opt,name=requested_server_name,json=requestedServerName,proto3,oneof"`
}

func (*Permission_AndRules) isPermission_Rule() {}

func (*Permission_OrRules) isPermission_Rule() {}

func (*Permission_Any) isPermission_Rule() {}

func (*Permission_Header) isPermission_Rule() {}

func (*Permission_DestinationIp) isPermission_Rule() {}

func (*Permission_DestinationPort) isPermission_Rule() {}

func (*Permission_Metadata) isPermission_Rule() {}

func (*Permission_NotRule) isPermission_Rule() {}

func (*Permission_RequestedServerName) isPermission_Rule() {}

func (m *Permission) GetRule() isPermission_Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

func (m *Permission) GetAndRules() *Permission_Set {
	if x, ok := m.GetRule().(*Permission_AndRules); ok {
		return x.AndRules
	}
	return nil
}

func (m *Permission) GetOrRules() *Permission_Set {
	if x, ok := m.GetRule().(*Permission_OrRules); ok {
		return x.OrRules
	}
	return nil
}

func (m *Permission) GetAny() bool {
	if x, ok := m.GetRule().(*Permission_Any); ok {
		return x.Any
	}
	return false
}

func (m *Permission) GetHeader() *route.HeaderMatcher {
	if x, ok := m.GetRule().(*Permission_Header); ok {
		return x.Header
	}
	return nil
}

func (m *Permission) GetDestinationIp() *core.CidrRange {
	if x, ok := m.GetRule().(*Permission_DestinationIp); ok {
		return x.DestinationIp
	}
	return nil
}

func (m *Permission) GetDestinationPort() uint32 {
	if x, ok := m.GetRule().(*Permission_DestinationPort); ok {
		return x.DestinationPort
	}
	return 0
}

func (m *Permission) GetMetadata() *v3alpha.MetadataMatcher {
	if x, ok := m.GetRule().(*Permission_Metadata); ok {
		return x.Metadata
	}
	return nil
}

func (m *Permission) GetNotRule() *Permission {
	if x, ok := m.GetRule().(*Permission_NotRule); ok {
		return x.NotRule
	}
	return nil
}

func (m *Permission) GetRequestedServerName() *v3alpha.StringMatcher {
	if x, ok := m.GetRule().(*Permission_RequestedServerName); ok {
		return x.RequestedServerName
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Permission) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Permission_AndRules)(nil),
		(*Permission_OrRules)(nil),
		(*Permission_Any)(nil),
		(*Permission_Header)(nil),
		(*Permission_DestinationIp)(nil),
		(*Permission_DestinationPort)(nil),
		(*Permission_Metadata)(nil),
		(*Permission_NotRule)(nil),
		(*Permission_RequestedServerName)(nil),
	}
}

type Permission_Set struct {
	Rules                []*Permission `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Permission_Set) Reset()         { *m = Permission_Set{} }
func (m *Permission_Set) String() string { return proto.CompactTextString(m) }
func (*Permission_Set) ProtoMessage()    {}
func (*Permission_Set) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb7d22ec66d90520, []int{2, 0}
}

func (m *Permission_Set) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permission_Set.Unmarshal(m, b)
}
func (m *Permission_Set) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permission_Set.Marshal(b, m, deterministic)
}
func (m *Permission_Set) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permission_Set.Merge(m, src)
}
func (m *Permission_Set) XXX_Size() int {
	return xxx_messageInfo_Permission_Set.Size(m)
}
func (m *Permission_Set) XXX_DiscardUnknown() {
	xxx_messageInfo_Permission_Set.DiscardUnknown(m)
}

var xxx_messageInfo_Permission_Set proto.InternalMessageInfo

func (m *Permission_Set) GetRules() []*Permission {
	if m != nil {
		return m.Rules
	}
	return nil
}

type Principal struct {
	// Types that are valid to be assigned to Identifier:
	//	*Principal_AndIds
	//	*Principal_OrIds
	//	*Principal_Any
	//	*Principal_Authenticated_
	//	*Principal_SourceIp
	//	*Principal_Header
	//	*Principal_Metadata
	//	*Principal_NotId
	Identifier           isPrincipal_Identifier `protobuf_oneof:"identifier"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Principal) Reset()         { *m = Principal{} }
func (m *Principal) String() string { return proto.CompactTextString(m) }
func (*Principal) ProtoMessage()    {}
func (*Principal) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb7d22ec66d90520, []int{3}
}

func (m *Principal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Principal.Unmarshal(m, b)
}
func (m *Principal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Principal.Marshal(b, m, deterministic)
}
func (m *Principal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Principal.Merge(m, src)
}
func (m *Principal) XXX_Size() int {
	return xxx_messageInfo_Principal.Size(m)
}
func (m *Principal) XXX_DiscardUnknown() {
	xxx_messageInfo_Principal.DiscardUnknown(m)
}

var xxx_messageInfo_Principal proto.InternalMessageInfo

type isPrincipal_Identifier interface {
	isPrincipal_Identifier()
}

type Principal_AndIds struct {
	AndIds *Principal_Set `protobuf:"bytes,1,opt,name=and_ids,json=andIds,proto3,oneof"`
}

type Principal_OrIds struct {
	OrIds *Principal_Set `protobuf:"bytes,2,opt,name=or_ids,json=orIds,proto3,oneof"`
}

type Principal_Any struct {
	Any bool `protobuf:"varint,3,opt,name=any,proto3,oneof"`
}

type Principal_Authenticated_ struct {
	Authenticated *Principal_Authenticated `protobuf:"bytes,4,opt,name=authenticated,proto3,oneof"`
}

type Principal_SourceIp struct {
	SourceIp *core.CidrRange `protobuf:"bytes,5,opt,name=source_ip,json=sourceIp,proto3,oneof"`
}

type Principal_Header struct {
	Header *route.HeaderMatcher `protobuf:"bytes,6,opt,name=header,proto3,oneof"`
}

type Principal_Metadata struct {
	Metadata *v3alpha.MetadataMatcher `protobuf:"bytes,7,opt,name=metadata,proto3,oneof"`
}

type Principal_NotId struct {
	NotId *Principal `protobuf:"bytes,8,opt,name=not_id,json=notId,proto3,oneof"`
}

func (*Principal_AndIds) isPrincipal_Identifier() {}

func (*Principal_OrIds) isPrincipal_Identifier() {}

func (*Principal_Any) isPrincipal_Identifier() {}

func (*Principal_Authenticated_) isPrincipal_Identifier() {}

func (*Principal_SourceIp) isPrincipal_Identifier() {}

func (*Principal_Header) isPrincipal_Identifier() {}

func (*Principal_Metadata) isPrincipal_Identifier() {}

func (*Principal_NotId) isPrincipal_Identifier() {}

func (m *Principal) GetIdentifier() isPrincipal_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *Principal) GetAndIds() *Principal_Set {
	if x, ok := m.GetIdentifier().(*Principal_AndIds); ok {
		return x.AndIds
	}
	return nil
}

func (m *Principal) GetOrIds() *Principal_Set {
	if x, ok := m.GetIdentifier().(*Principal_OrIds); ok {
		return x.OrIds
	}
	return nil
}

func (m *Principal) GetAny() bool {
	if x, ok := m.GetIdentifier().(*Principal_Any); ok {
		return x.Any
	}
	return false
}

func (m *Principal) GetAuthenticated() *Principal_Authenticated {
	if x, ok := m.GetIdentifier().(*Principal_Authenticated_); ok {
		return x.Authenticated
	}
	return nil
}

func (m *Principal) GetSourceIp() *core.CidrRange {
	if x, ok := m.GetIdentifier().(*Principal_SourceIp); ok {
		return x.SourceIp
	}
	return nil
}

func (m *Principal) GetHeader() *route.HeaderMatcher {
	if x, ok := m.GetIdentifier().(*Principal_Header); ok {
		return x.Header
	}
	return nil
}

func (m *Principal) GetMetadata() *v3alpha.MetadataMatcher {
	if x, ok := m.GetIdentifier().(*Principal_Metadata); ok {
		return x.Metadata
	}
	return nil
}

func (m *Principal) GetNotId() *Principal {
	if x, ok := m.GetIdentifier().(*Principal_NotId); ok {
		return x.NotId
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Principal) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Principal_AndIds)(nil),
		(*Principal_OrIds)(nil),
		(*Principal_Any)(nil),
		(*Principal_Authenticated_)(nil),
		(*Principal_SourceIp)(nil),
		(*Principal_Header)(nil),
		(*Principal_Metadata)(nil),
		(*Principal_NotId)(nil),
	}
}

type Principal_Set struct {
	Ids                  []*Principal `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Principal_Set) Reset()         { *m = Principal_Set{} }
func (m *Principal_Set) String() string { return proto.CompactTextString(m) }
func (*Principal_Set) ProtoMessage()    {}
func (*Principal_Set) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb7d22ec66d90520, []int{3, 0}
}

func (m *Principal_Set) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Principal_Set.Unmarshal(m, b)
}
func (m *Principal_Set) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Principal_Set.Marshal(b, m, deterministic)
}
func (m *Principal_Set) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Principal_Set.Merge(m, src)
}
func (m *Principal_Set) XXX_Size() int {
	return xxx_messageInfo_Principal_Set.Size(m)
}
func (m *Principal_Set) XXX_DiscardUnknown() {
	xxx_messageInfo_Principal_Set.DiscardUnknown(m)
}

var xxx_messageInfo_Principal_Set proto.InternalMessageInfo

func (m *Principal_Set) GetIds() []*Principal {
	if m != nil {
		return m.Ids
	}
	return nil
}

type Principal_Authenticated struct {
	PrincipalName        *v3alpha.StringMatcher `protobuf:"bytes,2,opt,name=principal_name,json=principalName,proto3" json:"principal_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Principal_Authenticated) Reset()         { *m = Principal_Authenticated{} }
func (m *Principal_Authenticated) String() string { return proto.CompactTextString(m) }
func (*Principal_Authenticated) ProtoMessage()    {}
func (*Principal_Authenticated) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb7d22ec66d90520, []int{3, 1}
}

func (m *Principal_Authenticated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Principal_Authenticated.Unmarshal(m, b)
}
func (m *Principal_Authenticated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Principal_Authenticated.Marshal(b, m, deterministic)
}
func (m *Principal_Authenticated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Principal_Authenticated.Merge(m, src)
}
func (m *Principal_Authenticated) XXX_Size() int {
	return xxx_messageInfo_Principal_Authenticated.Size(m)
}
func (m *Principal_Authenticated) XXX_DiscardUnknown() {
	xxx_messageInfo_Principal_Authenticated.DiscardUnknown(m)
}

var xxx_messageInfo_Principal_Authenticated proto.InternalMessageInfo

func (m *Principal_Authenticated) GetPrincipalName() *v3alpha.StringMatcher {
	if m != nil {
		return m.PrincipalName
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.config.rbac.v3alpha.RBAC_Action", RBAC_Action_name, RBAC_Action_value)
	proto.RegisterType((*RBAC)(nil), "envoy.config.rbac.v3alpha.RBAC")
	proto.RegisterMapType((map[string]*Policy)(nil), "envoy.config.rbac.v3alpha.RBAC.PoliciesEntry")
	proto.RegisterType((*Policy)(nil), "envoy.config.rbac.v3alpha.Policy")
	proto.RegisterType((*Permission)(nil), "envoy.config.rbac.v3alpha.Permission")
	proto.RegisterType((*Permission_Set)(nil), "envoy.config.rbac.v3alpha.Permission.Set")
	proto.RegisterType((*Principal)(nil), "envoy.config.rbac.v3alpha.Principal")
	proto.RegisterType((*Principal_Set)(nil), "envoy.config.rbac.v3alpha.Principal.Set")
	proto.RegisterType((*Principal_Authenticated)(nil), "envoy.config.rbac.v3alpha.Principal.Authenticated")
}

func init() {
	proto.RegisterFile("envoy/config/rbac/v3alpha/rbac.proto", fileDescriptor_eb7d22ec66d90520)
}

var fileDescriptor_eb7d22ec66d90520 = []byte{
	// 998 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xd1, 0x6e, 0xe3, 0x44,
	0x1b, 0x8d, 0x9d, 0xc4, 0x75, 0xbe, 0x2a, 0xfd, 0xfb, 0x0f, 0x42, 0x98, 0x54, 0x5b, 0xd2, 0x6c,
	0xbb, 0xdb, 0x2e, 0x60, 0x6b, 0x53, 0x69, 0x41, 0x11, 0xa0, 0xd6, 0xa5, 0x28, 0x5d, 0xed, 0x96,
	0xe2, 0x5e, 0x20, 0xb8, 0xa0, 0x9a, 0xda, 0xb3, 0xed, 0x40, 0x32, 0x63, 0xc6, 0x93, 0x28, 0x79,
	0x03, 0x6e, 0xb8, 0xe1, 0x12, 0x71, 0xc1, 0x33, 0xf0, 0x0e, 0x3c, 0x0b, 0xaf, 0x80, 0x7a, 0xb3,
	0x68, 0xc6, 0x8e, 0xe3, 0x88, 0x6c, 0x9a, 0x56, 0xdc, 0x44, 0x4e, 0xe6, 0x9c, 0x33, 0xdf, 0x7c,
	0x73, 0xce, 0x17, 0xc3, 0x36, 0x61, 0x43, 0x3e, 0xf6, 0x42, 0xce, 0x5e, 0xd1, 0x2b, 0x4f, 0x5c,
	0xe2, 0xd0, 0x1b, 0xee, 0xe3, 0x5e, 0x7c, 0x8d, 0xf5, 0x17, 0x37, 0x16, 0x5c, 0x72, 0xf4, 0xae,
	0x46, 0xb9, 0x29, 0xca, 0xd5, 0x0b, 0x19, 0xaa, 0x91, 0x09, 0xe0, 0x98, 0xe6, 0xc4, 0x90, 0x0b,
	0xe2, 0xe1, 0x28, 0x12, 0x24, 0x49, 0x52, 0x81, 0xc6, 0xc3, 0x7f, 0xa3, 0x04, 0x1f, 0x48, 0x92,
	0x7e, 0x66, 0xa0, 0xbd, 0x14, 0x24, 0xc7, 0x31, 0xf1, 0xfa, 0x58, 0x86, 0xd7, 0x44, 0xe4, 0xe8,
	0x3e, 0x91, 0x38, 0xc2, 0x12, 0x67, 0xd0, 0xc7, 0x0b, 0xa0, 0x89, 0x14, 0x94, 0x5d, 0x65, 0xc0,
	0x9d, 0x2b, 0xce, 0xaf, 0x7a, 0x44, 0xef, 0x4c, 0x46, 0xb1, 0xf0, 0x86, 0x4f, 0x35, 0xea, 0xa9,
	0x97, 0x8c, 0x99, 0xc4, 0xa3, 0x0c, 0xb6, 0x35, 0x88, 0x62, 0xec, 0x61, 0xc6, 0xb8, 0xc4, 0x92,
	0x72, 0x96, 0x78, 0x43, 0x22, 0x12, 0xca, 0xd9, 0x54, 0xe9, 0x9d, 0x21, 0xee, 0xd1, 0x08, 0x4b,
	0xe2, 0x4d, 0x1e, 0xd2, 0x85, 0xd6, 0x1f, 0x26, 0x54, 0x02, 0xff, 0xf0, 0x08, 0x7d, 0x06, 0x16,
	0x0e, 0x15, 0xdb, 0x31, 0x9a, 0xc6, 0xee, 0x5a, 0xfb, 0x91, 0xfb, 0xc6, 0xb6, 0xb9, 0x8a, 0xe0,
	0x1e, 0x6a, 0x74, 0x90, 0xb1, 0xd0, 0x09, 0xd8, 0x31, 0xef, 0xd1, 0x90, 0x92, 0xc4, 0x31, 0x9b,
	0xe5, 0xdd, 0xd5, 0xf6, 0x87, 0xb7, 0x29, 0x9c, 0x65, 0xf8, 0x63, 0x26, 0xc5, 0x38, 0xc8, 0xe9,
	0x8d, 0xef, 0xa0, 0x3e, 0xb3, 0x84, 0xd6, 0xa1, 0xfc, 0x03, 0x19, 0xeb, 0xc2, 0x6a, 0x81, 0x7a,
	0x44, 0x1f, 0x41, 0x75, 0x88, 0x7b, 0x03, 0xe2, 0x98, 0x4d, 0x63, 0x77, 0xb5, 0xbd, 0xb5, 0x60,
	0x2b, 0x2d, 0x35, 0x0e, 0x52, 0x7c, 0xc7, 0xfc, 0xd8, 0x68, 0x3d, 0x00, 0x2b, 0x2d, 0x1e, 0xd5,
	0xa0, 0x7a, 0xf8, 0xe2, 0xc5, 0x97, 0x5f, 0xaf, 0x97, 0x90, 0x0d, 0x95, 0xcf, 0x8f, 0x4f, 0xbf,
	0x59, 0x37, 0x3a, 0xcd, 0x5f, 0xff, 0xfc, 0x69, 0x73, 0x03, 0xe6, 0xd9, 0xa6, 0xad, 0x0b, 0x6f,
	0xfd, 0x6c, 0x82, 0x95, 0xca, 0xa2, 0xaf, 0x60, 0x35, 0x26, 0xa2, 0x4f, 0x13, 0xd5, 0xef, 0xc4,
	0x31, 0xf4, 0xc9, 0x77, 0x16, 0x95, 0x93, 0xa3, 0x7d, 0xfb, 0xc6, 0xaf, 0xfe, 0x62, 0x98, 0xb6,
	0x11, 0x14, 0x35, 0xd0, 0x29, 0x40, 0x2c, 0x28, 0x0b, 0x69, 0x8c, 0x7b, 0x93, 0x5e, 0x6e, 0x2f,
	0x52, 0x9c, 0x80, 0x0b, 0x82, 0x05, 0x05, 0xf4, 0x09, 0xd4, 0x42, 0xce, 0x22, 0xaa, 0x2f, 0xb7,
	0xac, 0xfb, 0xb5, 0xe9, 0xa6, 0xce, 0x72, 0x71, 0x4c, 0x5d, 0xe5, 0x2c, 0x77, 0xe2, 0x2c, 0xf7,
	0x78, 0x14, 0x8b, 0x60, 0x4a, 0xe8, 0xb4, 0x54, 0x37, 0x1e, 0xc0, 0xc6, 0xdc, 0x6e, 0xa4, 0x4d,
	0x68, 0xfd, 0x6e, 0x01, 0x4c, 0xcf, 0x85, 0xba, 0x50, 0xc3, 0x2c, 0xba, 0x10, 0x83, 0x1e, 0x49,
	0xf4, 0xa5, 0xad, 0xb6, 0xf7, 0x96, 0xea, 0x88, 0x7b, 0x4e, 0x64, 0xb7, 0x14, 0xd8, 0x98, 0x45,
	0x81, 0x22, 0xa3, 0x2f, 0xc0, 0xe6, 0x22, 0x13, 0x32, 0xef, 0x2e, 0xb4, 0xc2, 0x45, 0xaa, 0xb3,
	0x01, 0x65, 0xcc, 0xc6, 0xfa, 0xf0, 0xb6, 0xbf, 0x72, 0xe3, 0x57, 0xbe, 0x37, 0x6d, 0xa3, 0x5b,
	0x0a, 0xd4, 0xaf, 0xe8, 0x00, 0xac, 0x6b, 0x82, 0x23, 0x22, 0x9c, 0x8a, 0xde, 0x62, 0xe2, 0x7c,
	0xd5, 0x9b, 0x89, 0x74, 0x9a, 0xf4, 0xae, 0x86, 0xbd, 0x4c, 0x23, 0xdb, 0x2d, 0x05, 0x19, 0x0f,
	0x3d, 0x87, 0xb5, 0x88, 0x24, 0x92, 0x32, 0x1d, 0xbf, 0x0b, 0x1a, 0x3b, 0xd5, 0x19, 0x5b, 0x16,
	0x95, 0xd4, 0x7c, 0x71, 0x8f, 0x68, 0x24, 0x02, 0xcc, 0xae, 0x48, 0xb7, 0x14, 0xd4, 0x0b, 0xd4,
	0x93, 0x18, 0x3d, 0x83, 0xf5, 0xa2, 0x56, 0xcc, 0x85, 0x74, 0xac, 0xa6, 0xb1, 0x5b, 0xf7, 0x6b,
	0x37, 0xbe, 0xf5, 0xa4, 0xe2, 0xbc, 0x7e, 0x5d, 0xee, 0x96, 0x82, 0xff, 0x15, 0x40, 0x67, 0x5c,
	0x48, 0x95, 0xbf, 0xc9, 0x98, 0x71, 0x56, 0xf4, 0xee, 0xef, 0x67, 0xbb, 0xab, 0x39, 0xe3, 0x66,
	0x73, 0x26, 0x2f, 0xe3, 0x65, 0x86, 0x9d, 0x1e, 0x26, 0xa7, 0x23, 0x1f, 0x6c, 0xc6, 0xa5, 0x6e,
	0xbb, 0x63, 0x6b, 0xa9, 0xe5, 0x0c, 0xad, 0x3a, 0xce, 0xb8, 0x54, 0x2d, 0x47, 0x17, 0xf0, 0xb6,
	0x20, 0x3f, 0x0e, 0x48, 0x22, 0x49, 0x74, 0x91, 0x10, 0x31, 0x24, 0xe2, 0x82, 0xe1, 0x3e, 0x71,
	0x6a, 0x33, 0xd7, 0x38, 0xb7, 0xb6, 0x73, 0x3d, 0x03, 0xa7, 0x95, 0xbd, 0x95, 0x2b, 0x9d, 0x6b,
	0xa1, 0x53, 0xdc, 0x27, 0x8d, 0x11, 0x94, 0xcf, 0x89, 0x44, 0xc7, 0x50, 0x9d, 0xf8, 0xec, 0x5e,
	0xc9, 0x4b, 0xd9, 0x9d, 0x27, 0xca, 0xe5, 0x3b, 0xf0, 0x70, 0xbe, 0xcb, 0x67, 0x7c, 0xd5, 0x79,
	0xa4, 0xb0, 0x5b, 0xf0, 0xde, 0x2d, 0x58, 0x7f, 0x15, 0x2a, 0x4a, 0x1c, 0x95, 0xff, 0xf6, 0x8d,
	0xd6, 0x5f, 0x16, 0xd4, 0xf2, 0xa0, 0xa2, 0x23, 0x58, 0x51, 0x09, 0xa1, 0xd1, 0x24, 0x1f, 0xbb,
	0xcb, 0xe4, 0x3b, 0x73, 0xb5, 0x85, 0x59, 0x74, 0x12, 0x25, 0xe8, 0x10, 0x2c, 0x2e, 0xb4, 0x86,
	0x79, 0x67, 0x8d, 0x2a, 0x17, 0x4a, 0x62, 0x61, 0x2e, 0xbe, 0x85, 0x3a, 0x1e, 0xc8, 0x6b, 0xc2,
	0x24, 0x0d, 0xb1, 0x24, 0x51, 0x16, 0x8f, 0xf6, 0x52, 0xdb, 0x1c, 0x16, 0x99, 0xca, 0xe5, 0x33,
	0x52, 0xe8, 0x00, 0x6a, 0x09, 0x1f, 0x88, 0x90, 0xdc, 0x31, 0x2c, 0x76, 0xca, 0x3a, 0x89, 0x0b,
	0xa9, 0xb5, 0xee, 0x99, 0xda, 0xff, 0x30, 0x31, 0x9f, 0x82, 0xa5, 0x12, 0x43, 0xa3, 0x2c, 0x2f,
	0x4b, 0x8d, 0x6b, 0x75, 0x0d, 0x8c, 0xcb, 0x93, 0xa8, 0x21, 0x52, 0x2f, 0x1f, 0x40, 0x39, 0x75,
	0xc4, 0x7d, 0x26, 0xbe, 0xa2, 0x76, 0xf6, 0x94, 0x35, 0xb7, 0xa1, 0x35, 0xdf, 0x9a, 0x45, 0x0f,
	0x34, 0x7e, 0x33, 0xa0, 0x3e, 0x73, 0x49, 0xe8, 0x0c, 0xd6, 0xf2, 0x7f, 0x8d, 0x34, 0xab, 0xe6,
	0x1d, 0xb3, 0x1a, 0xd4, 0x73, 0x01, 0x95, 0xd1, 0xce, 0xbe, 0x2a, 0xc7, 0x85, 0x0f, 0x6e, 0x29,
	0x67, 0xa6, 0x8c, 0xe7, 0x15, 0xdb, 0x58, 0x37, 0x3b, 0x3b, 0x8a, 0xda, 0x84, 0xcd, 0xc5, 0x54,
	0xff, 0xff, 0x00, 0x34, 0x52, 0xdc, 0x57, 0x94, 0x08, 0x9d, 0x34, 0xff, 0x19, 0x3c, 0xa6, 0x3c,
	0x2d, 0x39, 0x16, 0x7c, 0x34, 0x7e, 0x73, 0x1f, 0xfd, 0x5a, 0x70, 0x89, 0xc3, 0x33, 0xf5, 0x1e,
	0x74, 0x66, 0x5c, 0x5a, 0xfa, 0x85, 0x68, 0xff, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x54, 0x69,
	0xa2, 0x87, 0x55, 0x0a, 0x00, 0x00,
}
