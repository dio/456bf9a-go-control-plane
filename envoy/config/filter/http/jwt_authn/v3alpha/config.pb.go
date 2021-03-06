// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/jwt_authn/v3alpha/config.proto

package envoy_config_filter_http_jwt_authn_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/core"
	route "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/route"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type JwtProvider struct {
	Issuer    string   `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Audiences []string `protobuf:"bytes,2,rep,name=audiences,proto3" json:"audiences,omitempty"`
	// Types that are valid to be assigned to JwksSourceSpecifier:
	//	*JwtProvider_RemoteJwks
	//	*JwtProvider_LocalJwks
	JwksSourceSpecifier  isJwtProvider_JwksSourceSpecifier `protobuf_oneof:"jwks_source_specifier"`
	Forward              bool                              `protobuf:"varint,5,opt,name=forward,proto3" json:"forward,omitempty"`
	FromHeaders          []*JwtHeader                      `protobuf:"bytes,6,rep,name=from_headers,json=fromHeaders,proto3" json:"from_headers,omitempty"`
	FromParams           []string                          `protobuf:"bytes,7,rep,name=from_params,json=fromParams,proto3" json:"from_params,omitempty"`
	ForwardPayloadHeader string                            `protobuf:"bytes,8,opt,name=forward_payload_header,json=forwardPayloadHeader,proto3" json:"forward_payload_header,omitempty"`
	PayloadInMetadata    string                            `protobuf:"bytes,9,opt,name=payload_in_metadata,json=payloadInMetadata,proto3" json:"payload_in_metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *JwtProvider) Reset()         { *m = JwtProvider{} }
func (m *JwtProvider) String() string { return proto.CompactTextString(m) }
func (*JwtProvider) ProtoMessage()    {}
func (*JwtProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_13f616927a7040fd, []int{0}
}

func (m *JwtProvider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JwtProvider.Unmarshal(m, b)
}
func (m *JwtProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JwtProvider.Marshal(b, m, deterministic)
}
func (m *JwtProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JwtProvider.Merge(m, src)
}
func (m *JwtProvider) XXX_Size() int {
	return xxx_messageInfo_JwtProvider.Size(m)
}
func (m *JwtProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_JwtProvider.DiscardUnknown(m)
}

var xxx_messageInfo_JwtProvider proto.InternalMessageInfo

func (m *JwtProvider) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *JwtProvider) GetAudiences() []string {
	if m != nil {
		return m.Audiences
	}
	return nil
}

type isJwtProvider_JwksSourceSpecifier interface {
	isJwtProvider_JwksSourceSpecifier()
}

type JwtProvider_RemoteJwks struct {
	RemoteJwks *RemoteJwks `protobuf:"bytes,3,opt,name=remote_jwks,json=remoteJwks,proto3,oneof"`
}

type JwtProvider_LocalJwks struct {
	LocalJwks *core.DataSource `protobuf:"bytes,4,opt,name=local_jwks,json=localJwks,proto3,oneof"`
}

func (*JwtProvider_RemoteJwks) isJwtProvider_JwksSourceSpecifier() {}

func (*JwtProvider_LocalJwks) isJwtProvider_JwksSourceSpecifier() {}

func (m *JwtProvider) GetJwksSourceSpecifier() isJwtProvider_JwksSourceSpecifier {
	if m != nil {
		return m.JwksSourceSpecifier
	}
	return nil
}

func (m *JwtProvider) GetRemoteJwks() *RemoteJwks {
	if x, ok := m.GetJwksSourceSpecifier().(*JwtProvider_RemoteJwks); ok {
		return x.RemoteJwks
	}
	return nil
}

func (m *JwtProvider) GetLocalJwks() *core.DataSource {
	if x, ok := m.GetJwksSourceSpecifier().(*JwtProvider_LocalJwks); ok {
		return x.LocalJwks
	}
	return nil
}

func (m *JwtProvider) GetForward() bool {
	if m != nil {
		return m.Forward
	}
	return false
}

func (m *JwtProvider) GetFromHeaders() []*JwtHeader {
	if m != nil {
		return m.FromHeaders
	}
	return nil
}

func (m *JwtProvider) GetFromParams() []string {
	if m != nil {
		return m.FromParams
	}
	return nil
}

func (m *JwtProvider) GetForwardPayloadHeader() string {
	if m != nil {
		return m.ForwardPayloadHeader
	}
	return ""
}

func (m *JwtProvider) GetPayloadInMetadata() string {
	if m != nil {
		return m.PayloadInMetadata
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*JwtProvider) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*JwtProvider_RemoteJwks)(nil),
		(*JwtProvider_LocalJwks)(nil),
	}
}

type RemoteJwks struct {
	HttpUri              *core.HttpUri      `protobuf:"bytes,1,opt,name=http_uri,json=httpUri,proto3" json:"http_uri,omitempty"`
	CacheDuration        *duration.Duration `protobuf:"bytes,2,opt,name=cache_duration,json=cacheDuration,proto3" json:"cache_duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RemoteJwks) Reset()         { *m = RemoteJwks{} }
func (m *RemoteJwks) String() string { return proto.CompactTextString(m) }
func (*RemoteJwks) ProtoMessage()    {}
func (*RemoteJwks) Descriptor() ([]byte, []int) {
	return fileDescriptor_13f616927a7040fd, []int{1}
}

func (m *RemoteJwks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteJwks.Unmarshal(m, b)
}
func (m *RemoteJwks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteJwks.Marshal(b, m, deterministic)
}
func (m *RemoteJwks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteJwks.Merge(m, src)
}
func (m *RemoteJwks) XXX_Size() int {
	return xxx_messageInfo_RemoteJwks.Size(m)
}
func (m *RemoteJwks) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteJwks.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteJwks proto.InternalMessageInfo

func (m *RemoteJwks) GetHttpUri() *core.HttpUri {
	if m != nil {
		return m.HttpUri
	}
	return nil
}

func (m *RemoteJwks) GetCacheDuration() *duration.Duration {
	if m != nil {
		return m.CacheDuration
	}
	return nil
}

type JwtHeader struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ValuePrefix          string   `protobuf:"bytes,2,opt,name=value_prefix,json=valuePrefix,proto3" json:"value_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JwtHeader) Reset()         { *m = JwtHeader{} }
func (m *JwtHeader) String() string { return proto.CompactTextString(m) }
func (*JwtHeader) ProtoMessage()    {}
func (*JwtHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_13f616927a7040fd, []int{2}
}

func (m *JwtHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JwtHeader.Unmarshal(m, b)
}
func (m *JwtHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JwtHeader.Marshal(b, m, deterministic)
}
func (m *JwtHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JwtHeader.Merge(m, src)
}
func (m *JwtHeader) XXX_Size() int {
	return xxx_messageInfo_JwtHeader.Size(m)
}
func (m *JwtHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_JwtHeader.DiscardUnknown(m)
}

var xxx_messageInfo_JwtHeader proto.InternalMessageInfo

func (m *JwtHeader) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *JwtHeader) GetValuePrefix() string {
	if m != nil {
		return m.ValuePrefix
	}
	return ""
}

type ProviderWithAudiences struct {
	ProviderName         string   `protobuf:"bytes,1,opt,name=provider_name,json=providerName,proto3" json:"provider_name,omitempty"`
	Audiences            []string `protobuf:"bytes,2,rep,name=audiences,proto3" json:"audiences,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProviderWithAudiences) Reset()         { *m = ProviderWithAudiences{} }
func (m *ProviderWithAudiences) String() string { return proto.CompactTextString(m) }
func (*ProviderWithAudiences) ProtoMessage()    {}
func (*ProviderWithAudiences) Descriptor() ([]byte, []int) {
	return fileDescriptor_13f616927a7040fd, []int{3}
}

func (m *ProviderWithAudiences) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProviderWithAudiences.Unmarshal(m, b)
}
func (m *ProviderWithAudiences) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProviderWithAudiences.Marshal(b, m, deterministic)
}
func (m *ProviderWithAudiences) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProviderWithAudiences.Merge(m, src)
}
func (m *ProviderWithAudiences) XXX_Size() int {
	return xxx_messageInfo_ProviderWithAudiences.Size(m)
}
func (m *ProviderWithAudiences) XXX_DiscardUnknown() {
	xxx_messageInfo_ProviderWithAudiences.DiscardUnknown(m)
}

var xxx_messageInfo_ProviderWithAudiences proto.InternalMessageInfo

func (m *ProviderWithAudiences) GetProviderName() string {
	if m != nil {
		return m.ProviderName
	}
	return ""
}

func (m *ProviderWithAudiences) GetAudiences() []string {
	if m != nil {
		return m.Audiences
	}
	return nil
}

type JwtRequirement struct {
	// Types that are valid to be assigned to RequiresType:
	//	*JwtRequirement_ProviderName
	//	*JwtRequirement_ProviderAndAudiences
	//	*JwtRequirement_RequiresAny
	//	*JwtRequirement_RequiresAll
	//	*JwtRequirement_AllowMissingOrFailed
	//	*JwtRequirement_AllowMissing
	RequiresType         isJwtRequirement_RequiresType `protobuf_oneof:"requires_type"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *JwtRequirement) Reset()         { *m = JwtRequirement{} }
func (m *JwtRequirement) String() string { return proto.CompactTextString(m) }
func (*JwtRequirement) ProtoMessage()    {}
func (*JwtRequirement) Descriptor() ([]byte, []int) {
	return fileDescriptor_13f616927a7040fd, []int{4}
}

func (m *JwtRequirement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JwtRequirement.Unmarshal(m, b)
}
func (m *JwtRequirement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JwtRequirement.Marshal(b, m, deterministic)
}
func (m *JwtRequirement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JwtRequirement.Merge(m, src)
}
func (m *JwtRequirement) XXX_Size() int {
	return xxx_messageInfo_JwtRequirement.Size(m)
}
func (m *JwtRequirement) XXX_DiscardUnknown() {
	xxx_messageInfo_JwtRequirement.DiscardUnknown(m)
}

var xxx_messageInfo_JwtRequirement proto.InternalMessageInfo

type isJwtRequirement_RequiresType interface {
	isJwtRequirement_RequiresType()
}

type JwtRequirement_ProviderName struct {
	ProviderName string `protobuf:"bytes,1,opt,name=provider_name,json=providerName,proto3,oneof"`
}

type JwtRequirement_ProviderAndAudiences struct {
	ProviderAndAudiences *ProviderWithAudiences `protobuf:"bytes,2,opt,name=provider_and_audiences,json=providerAndAudiences,proto3,oneof"`
}

type JwtRequirement_RequiresAny struct {
	RequiresAny *JwtRequirementOrList `protobuf:"bytes,3,opt,name=requires_any,json=requiresAny,proto3,oneof"`
}

type JwtRequirement_RequiresAll struct {
	RequiresAll *JwtRequirementAndList `protobuf:"bytes,4,opt,name=requires_all,json=requiresAll,proto3,oneof"`
}

type JwtRequirement_AllowMissingOrFailed struct {
	AllowMissingOrFailed *empty.Empty `protobuf:"bytes,5,opt,name=allow_missing_or_failed,json=allowMissingOrFailed,proto3,oneof"`
}

type JwtRequirement_AllowMissing struct {
	AllowMissing *empty.Empty `protobuf:"bytes,6,opt,name=allow_missing,json=allowMissing,proto3,oneof"`
}

func (*JwtRequirement_ProviderName) isJwtRequirement_RequiresType() {}

func (*JwtRequirement_ProviderAndAudiences) isJwtRequirement_RequiresType() {}

func (*JwtRequirement_RequiresAny) isJwtRequirement_RequiresType() {}

func (*JwtRequirement_RequiresAll) isJwtRequirement_RequiresType() {}

func (*JwtRequirement_AllowMissingOrFailed) isJwtRequirement_RequiresType() {}

func (*JwtRequirement_AllowMissing) isJwtRequirement_RequiresType() {}

func (m *JwtRequirement) GetRequiresType() isJwtRequirement_RequiresType {
	if m != nil {
		return m.RequiresType
	}
	return nil
}

func (m *JwtRequirement) GetProviderName() string {
	if x, ok := m.GetRequiresType().(*JwtRequirement_ProviderName); ok {
		return x.ProviderName
	}
	return ""
}

func (m *JwtRequirement) GetProviderAndAudiences() *ProviderWithAudiences {
	if x, ok := m.GetRequiresType().(*JwtRequirement_ProviderAndAudiences); ok {
		return x.ProviderAndAudiences
	}
	return nil
}

func (m *JwtRequirement) GetRequiresAny() *JwtRequirementOrList {
	if x, ok := m.GetRequiresType().(*JwtRequirement_RequiresAny); ok {
		return x.RequiresAny
	}
	return nil
}

func (m *JwtRequirement) GetRequiresAll() *JwtRequirementAndList {
	if x, ok := m.GetRequiresType().(*JwtRequirement_RequiresAll); ok {
		return x.RequiresAll
	}
	return nil
}

func (m *JwtRequirement) GetAllowMissingOrFailed() *empty.Empty {
	if x, ok := m.GetRequiresType().(*JwtRequirement_AllowMissingOrFailed); ok {
		return x.AllowMissingOrFailed
	}
	return nil
}

func (m *JwtRequirement) GetAllowMissing() *empty.Empty {
	if x, ok := m.GetRequiresType().(*JwtRequirement_AllowMissing); ok {
		return x.AllowMissing
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*JwtRequirement) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*JwtRequirement_ProviderName)(nil),
		(*JwtRequirement_ProviderAndAudiences)(nil),
		(*JwtRequirement_RequiresAny)(nil),
		(*JwtRequirement_RequiresAll)(nil),
		(*JwtRequirement_AllowMissingOrFailed)(nil),
		(*JwtRequirement_AllowMissing)(nil),
	}
}

type JwtRequirementOrList struct {
	Requirements         []*JwtRequirement `protobuf:"bytes,1,rep,name=requirements,proto3" json:"requirements,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *JwtRequirementOrList) Reset()         { *m = JwtRequirementOrList{} }
func (m *JwtRequirementOrList) String() string { return proto.CompactTextString(m) }
func (*JwtRequirementOrList) ProtoMessage()    {}
func (*JwtRequirementOrList) Descriptor() ([]byte, []int) {
	return fileDescriptor_13f616927a7040fd, []int{5}
}

func (m *JwtRequirementOrList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JwtRequirementOrList.Unmarshal(m, b)
}
func (m *JwtRequirementOrList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JwtRequirementOrList.Marshal(b, m, deterministic)
}
func (m *JwtRequirementOrList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JwtRequirementOrList.Merge(m, src)
}
func (m *JwtRequirementOrList) XXX_Size() int {
	return xxx_messageInfo_JwtRequirementOrList.Size(m)
}
func (m *JwtRequirementOrList) XXX_DiscardUnknown() {
	xxx_messageInfo_JwtRequirementOrList.DiscardUnknown(m)
}

var xxx_messageInfo_JwtRequirementOrList proto.InternalMessageInfo

func (m *JwtRequirementOrList) GetRequirements() []*JwtRequirement {
	if m != nil {
		return m.Requirements
	}
	return nil
}

type JwtRequirementAndList struct {
	Requirements         []*JwtRequirement `protobuf:"bytes,1,rep,name=requirements,proto3" json:"requirements,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *JwtRequirementAndList) Reset()         { *m = JwtRequirementAndList{} }
func (m *JwtRequirementAndList) String() string { return proto.CompactTextString(m) }
func (*JwtRequirementAndList) ProtoMessage()    {}
func (*JwtRequirementAndList) Descriptor() ([]byte, []int) {
	return fileDescriptor_13f616927a7040fd, []int{6}
}

func (m *JwtRequirementAndList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JwtRequirementAndList.Unmarshal(m, b)
}
func (m *JwtRequirementAndList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JwtRequirementAndList.Marshal(b, m, deterministic)
}
func (m *JwtRequirementAndList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JwtRequirementAndList.Merge(m, src)
}
func (m *JwtRequirementAndList) XXX_Size() int {
	return xxx_messageInfo_JwtRequirementAndList.Size(m)
}
func (m *JwtRequirementAndList) XXX_DiscardUnknown() {
	xxx_messageInfo_JwtRequirementAndList.DiscardUnknown(m)
}

var xxx_messageInfo_JwtRequirementAndList proto.InternalMessageInfo

func (m *JwtRequirementAndList) GetRequirements() []*JwtRequirement {
	if m != nil {
		return m.Requirements
	}
	return nil
}

type RequirementRule struct {
	Match                *route.RouteMatch `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	Requires             *JwtRequirement   `protobuf:"bytes,2,opt,name=requires,proto3" json:"requires,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RequirementRule) Reset()         { *m = RequirementRule{} }
func (m *RequirementRule) String() string { return proto.CompactTextString(m) }
func (*RequirementRule) ProtoMessage()    {}
func (*RequirementRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_13f616927a7040fd, []int{7}
}

func (m *RequirementRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequirementRule.Unmarshal(m, b)
}
func (m *RequirementRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequirementRule.Marshal(b, m, deterministic)
}
func (m *RequirementRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequirementRule.Merge(m, src)
}
func (m *RequirementRule) XXX_Size() int {
	return xxx_messageInfo_RequirementRule.Size(m)
}
func (m *RequirementRule) XXX_DiscardUnknown() {
	xxx_messageInfo_RequirementRule.DiscardUnknown(m)
}

var xxx_messageInfo_RequirementRule proto.InternalMessageInfo

func (m *RequirementRule) GetMatch() *route.RouteMatch {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *RequirementRule) GetRequires() *JwtRequirement {
	if m != nil {
		return m.Requires
	}
	return nil
}

type FilterStateRule struct {
	Name                 string                     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Requires             map[string]*JwtRequirement `protobuf:"bytes,3,rep,name=requires,proto3" json:"requires,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *FilterStateRule) Reset()         { *m = FilterStateRule{} }
func (m *FilterStateRule) String() string { return proto.CompactTextString(m) }
func (*FilterStateRule) ProtoMessage()    {}
func (*FilterStateRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_13f616927a7040fd, []int{8}
}

func (m *FilterStateRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterStateRule.Unmarshal(m, b)
}
func (m *FilterStateRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterStateRule.Marshal(b, m, deterministic)
}
func (m *FilterStateRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterStateRule.Merge(m, src)
}
func (m *FilterStateRule) XXX_Size() int {
	return xxx_messageInfo_FilterStateRule.Size(m)
}
func (m *FilterStateRule) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterStateRule.DiscardUnknown(m)
}

var xxx_messageInfo_FilterStateRule proto.InternalMessageInfo

func (m *FilterStateRule) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FilterStateRule) GetRequires() map[string]*JwtRequirement {
	if m != nil {
		return m.Requires
	}
	return nil
}

type JwtAuthentication struct {
	Providers            map[string]*JwtProvider `protobuf:"bytes,1,rep,name=providers,proto3" json:"providers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Rules                []*RequirementRule      `protobuf:"bytes,2,rep,name=rules,proto3" json:"rules,omitempty"`
	FilterStateRules     *FilterStateRule        `protobuf:"bytes,3,opt,name=filter_state_rules,json=filterStateRules,proto3" json:"filter_state_rules,omitempty"`
	BypassCorsPreflight  bool                    `protobuf:"varint,4,opt,name=bypass_cors_preflight,json=bypassCorsPreflight,proto3" json:"bypass_cors_preflight,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *JwtAuthentication) Reset()         { *m = JwtAuthentication{} }
func (m *JwtAuthentication) String() string { return proto.CompactTextString(m) }
func (*JwtAuthentication) ProtoMessage()    {}
func (*JwtAuthentication) Descriptor() ([]byte, []int) {
	return fileDescriptor_13f616927a7040fd, []int{9}
}

func (m *JwtAuthentication) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JwtAuthentication.Unmarshal(m, b)
}
func (m *JwtAuthentication) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JwtAuthentication.Marshal(b, m, deterministic)
}
func (m *JwtAuthentication) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JwtAuthentication.Merge(m, src)
}
func (m *JwtAuthentication) XXX_Size() int {
	return xxx_messageInfo_JwtAuthentication.Size(m)
}
func (m *JwtAuthentication) XXX_DiscardUnknown() {
	xxx_messageInfo_JwtAuthentication.DiscardUnknown(m)
}

var xxx_messageInfo_JwtAuthentication proto.InternalMessageInfo

func (m *JwtAuthentication) GetProviders() map[string]*JwtProvider {
	if m != nil {
		return m.Providers
	}
	return nil
}

func (m *JwtAuthentication) GetRules() []*RequirementRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *JwtAuthentication) GetFilterStateRules() *FilterStateRule {
	if m != nil {
		return m.FilterStateRules
	}
	return nil
}

func (m *JwtAuthentication) GetBypassCorsPreflight() bool {
	if m != nil {
		return m.BypassCorsPreflight
	}
	return false
}

func init() {
	proto.RegisterType((*JwtProvider)(nil), "envoy.config.filter.http.jwt_authn.v3alpha.JwtProvider")
	proto.RegisterType((*RemoteJwks)(nil), "envoy.config.filter.http.jwt_authn.v3alpha.RemoteJwks")
	proto.RegisterType((*JwtHeader)(nil), "envoy.config.filter.http.jwt_authn.v3alpha.JwtHeader")
	proto.RegisterType((*ProviderWithAudiences)(nil), "envoy.config.filter.http.jwt_authn.v3alpha.ProviderWithAudiences")
	proto.RegisterType((*JwtRequirement)(nil), "envoy.config.filter.http.jwt_authn.v3alpha.JwtRequirement")
	proto.RegisterType((*JwtRequirementOrList)(nil), "envoy.config.filter.http.jwt_authn.v3alpha.JwtRequirementOrList")
	proto.RegisterType((*JwtRequirementAndList)(nil), "envoy.config.filter.http.jwt_authn.v3alpha.JwtRequirementAndList")
	proto.RegisterType((*RequirementRule)(nil), "envoy.config.filter.http.jwt_authn.v3alpha.RequirementRule")
	proto.RegisterType((*FilterStateRule)(nil), "envoy.config.filter.http.jwt_authn.v3alpha.FilterStateRule")
	proto.RegisterMapType((map[string]*JwtRequirement)(nil), "envoy.config.filter.http.jwt_authn.v3alpha.FilterStateRule.RequiresEntry")
	proto.RegisterType((*JwtAuthentication)(nil), "envoy.config.filter.http.jwt_authn.v3alpha.JwtAuthentication")
	proto.RegisterMapType((map[string]*JwtProvider)(nil), "envoy.config.filter.http.jwt_authn.v3alpha.JwtAuthentication.ProvidersEntry")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/jwt_authn/v3alpha/config.proto", fileDescriptor_13f616927a7040fd)
}

var fileDescriptor_13f616927a7040fd = []byte{
	// 1214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xcb, 0x6e, 0xdb, 0x46,
	0x17, 0x36, 0x25, 0xd9, 0x96, 0x8e, 0xec, 0x5c, 0x26, 0x71, 0xc2, 0xdf, 0x09, 0x12, 0x45, 0x41,
	0x00, 0xe3, 0x5f, 0x50, 0x80, 0x72, 0xad, 0x9c, 0xb4, 0x96, 0x9c, 0xa4, 0x8a, 0x11, 0x37, 0x2a,
	0x83, 0xde, 0x56, 0xc4, 0x58, 0x1c, 0x59, 0x93, 0x50, 0x24, 0x3b, 0x33, 0x94, 0xa2, 0x37, 0x08,
	0xba, 0xec, 0xb2, 0xdb, 0xae, 0xfb, 0x16, 0x2d, 0x50, 0x74, 0xd7, 0x77, 0xe8, 0x0b, 0x74, 0x57,
	0x64, 0x55, 0xcc, 0x85, 0xba, 0xd8, 0x4a, 0x1a, 0x3a, 0x45, 0x37, 0x02, 0xe7, 0x9c, 0x39, 0xdf,
	0x7c, 0xe7, 0x2e, 0xb8, 0x4b, 0xc2, 0x61, 0x34, 0xae, 0x75, 0xa3, 0xb0, 0x47, 0x0f, 0x6b, 0x3d,
	0x1a, 0x08, 0xc2, 0x6a, 0x7d, 0x21, 0xe2, 0xda, 0x8b, 0x91, 0xf0, 0x70, 0x22, 0xfa, 0x61, 0x6d,
	0x78, 0x13, 0x07, 0x71, 0x1f, 0x9b, 0x4b, 0x4e, 0xcc, 0x22, 0x11, 0xa1, 0xff, 0x2b, 0x43, 0xc7,
	0xc8, 0xb4, 0xa1, 0x23, 0x0d, 0x9d, 0x89, 0xa1, 0x63, 0x0c, 0x37, 0xaf, 0xe9, 0x47, 0x70, 0x4c,
	0x67, 0xb0, 0x18, 0xa9, 0x1d, 0x60, 0x4e, 0x34, 0xdc, 0xe6, 0x8d, 0xb7, 0x5c, 0x91, 0x88, 0x5e,
	0xc2, 0xa8, 0xb9, 0x76, 0xfd, 0xf8, 0x35, 0x16, 0x25, 0x82, 0xe8, 0x5f, 0x73, 0xe9, 0xca, 0x61,
	0x14, 0x1d, 0x06, 0xa4, 0xa6, 0x4e, 0x07, 0x49, 0xaf, 0xe6, 0x27, 0x0c, 0x0b, 0x1a, 0x85, 0x46,
	0x7f, 0xe9, 0xa8, 0x9e, 0x0c, 0x62, 0x31, 0x36, 0xca, 0x6b, 0x89, 0x1f, 0xe3, 0x1a, 0x0e, 0xc3,
	0x48, 0x28, 0x1b, 0x5e, 0x1b, 0x12, 0xc6, 0x69, 0x14, 0xd2, 0xd0, 0xb8, 0xbe, 0x79, 0x71, 0x88,
	0x03, 0xea, 0x63, 0x41, 0x6a, 0xe9, 0x87, 0x56, 0x54, 0x7f, 0x2e, 0x40, 0x79, 0x6f, 0x24, 0x3a,
	0x2c, 0x1a, 0x52, 0x9f, 0x30, 0x74, 0x15, 0x56, 0x28, 0xe7, 0x09, 0x61, 0xb6, 0x55, 0xb1, 0xb6,
	0x4a, 0xad, 0xd5, 0x37, 0xad, 0x02, 0xcb, 0x55, 0x2c, 0xd7, 0x88, 0xd1, 0x65, 0x28, 0xe1, 0xc4,
	0xa7, 0x24, 0xec, 0x12, 0x6e, 0xe7, 0x2a, 0xf9, 0xad, 0x92, 0x3b, 0x15, 0xa0, 0x6f, 0xa0, 0xcc,
	0xc8, 0x20, 0x12, 0xc4, 0x7b, 0x31, 0x7a, 0xc9, 0xed, 0x7c, 0xc5, 0xda, 0x2a, 0xd7, 0xef, 0x38,
	0xef, 0x1f, 0x78, 0xc7, 0x55, 0xe6, 0x7b, 0xa3, 0x97, 0xbc, 0xbd, 0xe4, 0x02, 0x9b, 0x9c, 0xd0,
	0x2e, 0x40, 0x10, 0x75, 0x71, 0xa0, 0x91, 0x0b, 0x0a, 0xb9, 0x6a, 0x90, 0x71, 0x4c, 0x27, 0x00,
	0x32, 0x07, 0xce, 0x43, 0x2c, 0xf0, 0xf3, 0x28, 0x61, 0x5d, 0xd2, 0x5e, 0x72, 0x4b, 0xca, 0x4e,
	0x81, 0xd8, 0xb0, 0xda, 0x8b, 0xd8, 0x08, 0x33, 0xdf, 0x5e, 0xae, 0x58, 0x5b, 0x45, 0x37, 0x3d,
	0xa2, 0xaf, 0x61, 0xad, 0xc7, 0xa2, 0x81, 0xd7, 0x27, 0xd8, 0x27, 0x8c, 0xdb, 0x2b, 0x95, 0xfc,
	0x56, 0xb9, 0x7e, 0x3b, 0x0b, 0xf5, 0xbd, 0x91, 0x68, 0x2b, 0x6b, 0xb7, 0x2c, 0xa1, 0xf4, 0x37,
	0x47, 0x57, 0x41, 0x1d, 0xbd, 0x18, 0x33, 0x3c, 0xe0, 0xf6, 0xaa, 0x8a, 0x19, 0x48, 0x51, 0x47,
	0x49, 0xd0, 0x2d, 0xb8, 0x60, 0x58, 0x78, 0x31, 0x1e, 0x07, 0x11, 0xf6, 0x0d, 0x0b, 0xbb, 0x28,
	0x73, 0xe0, 0x9e, 0x37, 0xda, 0x8e, 0x56, 0x6a, 0x5c, 0xe4, 0xc0, 0xb9, 0xf4, 0x36, 0x0d, 0xbd,
	0x01, 0x11, 0xd8, 0xc7, 0x02, 0xdb, 0x25, 0x65, 0x72, 0xd6, 0xa8, 0x9e, 0x84, 0xfb, 0x46, 0xd1,
	0x78, 0xf0, 0xc3, 0x2f, 0xaf, 0xaf, 0xdc, 0x83, 0x3b, 0xef, 0xe3, 0x50, 0x7d, 0xe2, 0x50, 0x5a,
	0x18, 0xad, 0xcb, 0xb0, 0x21, 0x03, 0xef, 0x71, 0x15, 0x55, 0x8f, 0xc7, 0xa4, 0x4b, 0x7b, 0x94,
	0x30, 0x94, 0xff, 0xab, 0x65, 0x55, 0x7f, 0xb5, 0x00, 0xa6, 0x99, 0x43, 0x0d, 0x28, 0xa6, 0x5d,
	0xa0, 0xea, 0xa8, 0x5c, 0xbf, 0xfa, 0xb6, 0x4c, 0xb5, 0x85, 0x88, 0xbf, 0x60, 0xd4, 0x5d, 0xed,
	0xeb, 0x0f, 0xb4, 0x03, 0xa7, 0xba, 0xb8, 0xdb, 0x27, 0x5e, 0xda, 0x02, 0x76, 0x4e, 0x21, 0xfc,
	0xcf, 0xd1, 0x3d, 0xe0, 0xa4, 0x3d, 0xe0, 0x3c, 0x34, 0x17, 0xdc, 0x75, 0x65, 0x90, 0x1e, 0x1b,
	0xf7, 0xa5, 0xa7, 0x77, 0xe1, 0x76, 0x06, 0x4f, 0xa7, 0xdc, 0xab, 0xaf, 0x2d, 0x28, 0x4d, 0x32,
	0x89, 0x2e, 0x41, 0x21, 0xc4, 0x03, 0x72, 0xb4, 0x1b, 0x94, 0x10, 0x5d, 0x83, 0xb5, 0x21, 0x0e,
	0x12, 0xe2, 0xc5, 0x8c, 0xf4, 0xe8, 0x2b, 0x45, 0xb4, 0xe4, 0x96, 0x95, 0xac, 0xa3, 0x44, 0x8d,
	0x6d, 0xc9, 0xe5, 0x0e, 0xdc, 0xca, 0x16, 0x75, 0xfd, 0x78, 0xf5, 0x47, 0x0b, 0x36, 0xd2, 0x04,
	0x7c, 0x45, 0x45, 0xbf, 0x39, 0xe9, 0xb3, 0xeb, 0xb0, 0x1e, 0x1b, 0x85, 0x37, 0xe5, 0xe7, 0xae,
	0xa5, 0xc2, 0xcf, 0x24, 0xbd, 0x77, 0xb6, 0x6a, 0xe3, 0x53, 0xc9, 0xac, 0x05, 0x3b, 0x19, 0x98,
	0x2d, 0xe4, 0x52, 0xfd, 0xa3, 0x00, 0xa7, 0xf6, 0x46, 0xc2, 0x25, 0xdf, 0x26, 0x94, 0x91, 0x01,
	0x09, 0x05, 0xba, 0xb1, 0x90, 0x5e, 0x7b, 0xe9, 0x08, 0xc1, 0x31, 0x5c, 0x98, 0x5c, 0xc3, 0xa1,
	0xef, 0xcd, 0xb2, 0x95, 0x29, 0x6f, 0x66, 0xe9, 0xbe, 0x85, 0xe4, 0xda, 0x4b, 0xee, 0xf9, 0xf4,
	0x89, 0x66, 0xe8, 0x4f, 0x03, 0x48, 0x60, 0x8d, 0x69, 0xc2, 0xdc, 0xc3, 0xe1, 0xd8, 0x4c, 0xaa,
	0x9d, 0x8c, 0xed, 0x3e, 0xe3, 0xf3, 0x33, 0xf6, 0x94, 0x72, 0xd1, 0x5e, 0x72, 0xcb, 0x29, 0x6e,
	0x33, 0x1c, 0xa3, 0xde, 0xec, 0x33, 0x41, 0x60, 0xc6, 0x56, 0xf3, 0xe4, 0xcf, 0x34, 0x43, 0xff,
	0xd8, 0x3b, 0x41, 0x80, 0x9e, 0xc1, 0x45, 0x1c, 0x04, 0xd1, 0xc8, 0x1b, 0x50, 0xce, 0x69, 0x78,
	0xe8, 0x45, 0xcc, 0xeb, 0x61, 0x1a, 0x10, 0x3d, 0xe7, 0xca, 0xf5, 0x0b, 0xc7, 0xba, 0xe7, 0x91,
	0xdc, 0x20, 0x32, 0x3e, 0xca, 0x70, 0x5f, 0xdb, 0x3d, 0x63, 0x8f, 0x95, 0x15, 0x7a, 0x00, 0xeb,
	0x73, 0x80, 0xf6, 0xca, 0x3f, 0xc0, 0xac, 0xcd, 0xc2, 0x34, 0x76, 0x64, 0x71, 0x6d, 0xc3, 0x47,
	0xd9, 0xca, 0x7e, 0xc6, 0xcf, 0xd6, 0x69, 0x58, 0x9f, 0x44, 0x4e, 0x8c, 0x63, 0x22, 0x47, 0xcc,
	0xf9, 0x45, 0x21, 0x47, 0xfd, 0x49, 0x8c, 0xa5, 0x90, 0xdb, 0x96, 0x9a, 0xdc, 0x8d, 0x93, 0xc7,
	0xb8, 0x55, 0x7c, 0xd3, 0x5a, 0xfe, 0xde, 0xca, 0x15, 0x73, 0xee, 0x1c, 0x72, 0xe3, 0xb1, 0xf4,
	0xaa, 0x09, 0x9f, 0x9c, 0xd8, 0x2b, 0xcd, 0xb8, 0xfa, 0x9b, 0x05, 0x1b, 0x0b, 0xd3, 0xfa, 0x1f,
	0xfa, 0x72, 0x82, 0xf6, 0x5f, 0x48, 0xb9, 0xfa, 0xa7, 0x05, 0xa7, 0x67, 0xc4, 0x6e, 0x12, 0x10,
	0xb4, 0x0b, 0xcb, 0x03, 0x2c, 0xba, 0x7d, 0x33, 0xfc, 0xaf, 0x2f, 0x18, 0xfe, 0xfa, 0xdf, 0x8f,
	0x2b, 0x7f, 0xf7, 0xe5, 0x55, 0x45, 0xf4, 0x3b, 0x2b, 0x77, 0xc6, 0x72, 0xb5, 0x2d, 0xfa, 0x12,
	0x8a, 0x69, 0x05, 0x98, 0x79, 0xf0, 0x01, 0x71, 0x70, 0x27, 0x58, 0x8d, 0xa6, 0xf4, 0xfc, 0x3e,
	0x34, 0x32, 0xad, 0x87, 0x39, 0xff, 0xaa, 0xbf, 0xe7, 0xe0, 0xf4, 0x63, 0x65, 0xf1, 0x5c, 0x60,
	0x41, 0x94, 0xcf, 0xef, 0xdc, 0x14, 0x64, 0xc6, 0x97, 0xbc, 0xca, 0xe9, 0x93, 0x2c, 0xbe, 0x1c,
	0x79, 0x2b, 0xe5, 0xc3, 0x1f, 0x85, 0x82, 0x8d, 0xa7, 0xae, 0x6d, 0x8e, 0x60, 0x7d, 0x4e, 0x85,
	0xce, 0x40, 0xfe, 0x25, 0x19, 0x9b, 0xed, 0x20, 0x3f, 0x51, 0x07, 0x96, 0xd5, 0x7e, 0xfa, 0x17,
	0x42, 0xaa, 0x81, 0x1a, 0xb9, 0x7b, 0xd6, 0x49, 0x62, 0x7a, 0xc4, 0xa7, 0xea, 0x4f, 0x05, 0x38,
	0xbb, 0x37, 0x12, 0xcd, 0x44, 0xf4, 0x49, 0x28, 0x68, 0x57, 0xed, 0x72, 0xf4, 0x02, 0x4a, 0xe9,
	0xfc, 0x4e, 0xbb, 0xe1, 0x69, 0x46, 0xca, 0xf3, 0x88, 0x93, 0x3d, 0x61, 0x82, 0x37, 0x85, 0x47,
	0x9f, 0xc3, 0x32, 0x4b, 0x02, 0xb3, 0x2b, 0xcb, 0xf5, 0xed, 0x6c, 0x7f, 0x5b, 0xe7, 0x2a, 0xc4,
	0xd5, 0x48, 0x88, 0x02, 0xd2, 0x76, 0x1e, 0x97, 0x8e, 0x7a, 0x1a, 0x5f, 0x2f, 0x9b, 0xed, 0x0f,
	0xa8, 0x00, 0xf7, 0x4c, 0x6f, 0x5e, 0xc0, 0x51, 0x1d, 0x36, 0x0e, 0xc6, 0x31, 0xe6, 0xdc, 0xeb,
	0x46, 0x8c, 0xab, 0xbf, 0x24, 0x01, 0x3d, 0xec, 0x0b, 0xb5, 0x73, 0x8a, 0xee, 0x39, 0xad, 0xdc,
	0x8d, 0x18, 0xef, 0xa4, 0xaa, 0xcd, 0x04, 0x4e, 0xcd, 0x87, 0x63, 0x41, 0xc1, 0xec, 0xcf, 0x17,
	0xcc, 0xdd, 0x8c, 0xd1, 0x4f, 0xf1, 0x67, 0xab, 0x65, 0x57, 0x56, 0xcb, 0xc7, 0x70, 0x3f, 0xdb,
	0xec, 0x99, 0xcf, 0x63, 0xeb, 0x09, 0xdc, 0xa3, 0x91, 0x26, 0x13, 0xb3, 0xe8, 0xd5, 0x38, 0x03,
	0xaf, 0x56, 0x79, 0x57, 0xdd, 0xea, 0xc8, 0x25, 0xd6, 0xb1, 0x0e, 0x56, 0xd4, 0x36, 0xbb, 0xf9,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x30, 0x61, 0xb2, 0x5a, 0x0e, 0x00, 0x00,
}
