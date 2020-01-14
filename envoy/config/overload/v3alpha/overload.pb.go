// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/overload/v3alpha/overload.proto

package envoy_config_overload_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	duration "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/struct"
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

type ResourceMonitor struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to ConfigType:
	//	*ResourceMonitor_TypedConfig
	ConfigType           isResourceMonitor_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ResourceMonitor) Reset()         { *m = ResourceMonitor{} }
func (m *ResourceMonitor) String() string { return proto.CompactTextString(m) }
func (*ResourceMonitor) ProtoMessage()    {}
func (*ResourceMonitor) Descriptor() ([]byte, []int) {
	return fileDescriptor_542f910eb6259f77, []int{0}
}

func (m *ResourceMonitor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceMonitor.Unmarshal(m, b)
}
func (m *ResourceMonitor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceMonitor.Marshal(b, m, deterministic)
}
func (m *ResourceMonitor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceMonitor.Merge(m, src)
}
func (m *ResourceMonitor) XXX_Size() int {
	return xxx_messageInfo_ResourceMonitor.Size(m)
}
func (m *ResourceMonitor) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceMonitor.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceMonitor proto.InternalMessageInfo

func (m *ResourceMonitor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isResourceMonitor_ConfigType interface {
	isResourceMonitor_ConfigType()
}

type ResourceMonitor_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*ResourceMonitor_TypedConfig) isResourceMonitor_ConfigType() {}

func (m *ResourceMonitor) GetConfigType() isResourceMonitor_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *ResourceMonitor) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*ResourceMonitor_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ResourceMonitor) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ResourceMonitor_TypedConfig)(nil),
	}
}

type ThresholdTrigger struct {
	Value                float64  `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThresholdTrigger) Reset()         { *m = ThresholdTrigger{} }
func (m *ThresholdTrigger) String() string { return proto.CompactTextString(m) }
func (*ThresholdTrigger) ProtoMessage()    {}
func (*ThresholdTrigger) Descriptor() ([]byte, []int) {
	return fileDescriptor_542f910eb6259f77, []int{1}
}

func (m *ThresholdTrigger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThresholdTrigger.Unmarshal(m, b)
}
func (m *ThresholdTrigger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThresholdTrigger.Marshal(b, m, deterministic)
}
func (m *ThresholdTrigger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThresholdTrigger.Merge(m, src)
}
func (m *ThresholdTrigger) XXX_Size() int {
	return xxx_messageInfo_ThresholdTrigger.Size(m)
}
func (m *ThresholdTrigger) XXX_DiscardUnknown() {
	xxx_messageInfo_ThresholdTrigger.DiscardUnknown(m)
}

var xxx_messageInfo_ThresholdTrigger proto.InternalMessageInfo

func (m *ThresholdTrigger) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Trigger struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to TriggerOneof:
	//	*Trigger_Threshold
	TriggerOneof         isTrigger_TriggerOneof `protobuf_oneof:"trigger_oneof"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Trigger) Reset()         { *m = Trigger{} }
func (m *Trigger) String() string { return proto.CompactTextString(m) }
func (*Trigger) ProtoMessage()    {}
func (*Trigger) Descriptor() ([]byte, []int) {
	return fileDescriptor_542f910eb6259f77, []int{2}
}

func (m *Trigger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Trigger.Unmarshal(m, b)
}
func (m *Trigger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Trigger.Marshal(b, m, deterministic)
}
func (m *Trigger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trigger.Merge(m, src)
}
func (m *Trigger) XXX_Size() int {
	return xxx_messageInfo_Trigger.Size(m)
}
func (m *Trigger) XXX_DiscardUnknown() {
	xxx_messageInfo_Trigger.DiscardUnknown(m)
}

var xxx_messageInfo_Trigger proto.InternalMessageInfo

func (m *Trigger) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isTrigger_TriggerOneof interface {
	isTrigger_TriggerOneof()
}

type Trigger_Threshold struct {
	Threshold *ThresholdTrigger `protobuf:"bytes,2,opt,name=threshold,proto3,oneof"`
}

func (*Trigger_Threshold) isTrigger_TriggerOneof() {}

func (m *Trigger) GetTriggerOneof() isTrigger_TriggerOneof {
	if m != nil {
		return m.TriggerOneof
	}
	return nil
}

func (m *Trigger) GetThreshold() *ThresholdTrigger {
	if x, ok := m.GetTriggerOneof().(*Trigger_Threshold); ok {
		return x.Threshold
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Trigger) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Trigger_Threshold)(nil),
	}
}

type OverloadAction struct {
	Name                 string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Triggers             []*Trigger `protobuf:"bytes,2,rep,name=triggers,proto3" json:"triggers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *OverloadAction) Reset()         { *m = OverloadAction{} }
func (m *OverloadAction) String() string { return proto.CompactTextString(m) }
func (*OverloadAction) ProtoMessage()    {}
func (*OverloadAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_542f910eb6259f77, []int{3}
}

func (m *OverloadAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OverloadAction.Unmarshal(m, b)
}
func (m *OverloadAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OverloadAction.Marshal(b, m, deterministic)
}
func (m *OverloadAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OverloadAction.Merge(m, src)
}
func (m *OverloadAction) XXX_Size() int {
	return xxx_messageInfo_OverloadAction.Size(m)
}
func (m *OverloadAction) XXX_DiscardUnknown() {
	xxx_messageInfo_OverloadAction.DiscardUnknown(m)
}

var xxx_messageInfo_OverloadAction proto.InternalMessageInfo

func (m *OverloadAction) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OverloadAction) GetTriggers() []*Trigger {
	if m != nil {
		return m.Triggers
	}
	return nil
}

type OverloadManager struct {
	RefreshInterval      *duration.Duration `protobuf:"bytes,1,opt,name=refresh_interval,json=refreshInterval,proto3" json:"refresh_interval,omitempty"`
	ResourceMonitors     []*ResourceMonitor `protobuf:"bytes,2,rep,name=resource_monitors,json=resourceMonitors,proto3" json:"resource_monitors,omitempty"`
	Actions              []*OverloadAction  `protobuf:"bytes,3,rep,name=actions,proto3" json:"actions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *OverloadManager) Reset()         { *m = OverloadManager{} }
func (m *OverloadManager) String() string { return proto.CompactTextString(m) }
func (*OverloadManager) ProtoMessage()    {}
func (*OverloadManager) Descriptor() ([]byte, []int) {
	return fileDescriptor_542f910eb6259f77, []int{4}
}

func (m *OverloadManager) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OverloadManager.Unmarshal(m, b)
}
func (m *OverloadManager) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OverloadManager.Marshal(b, m, deterministic)
}
func (m *OverloadManager) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OverloadManager.Merge(m, src)
}
func (m *OverloadManager) XXX_Size() int {
	return xxx_messageInfo_OverloadManager.Size(m)
}
func (m *OverloadManager) XXX_DiscardUnknown() {
	xxx_messageInfo_OverloadManager.DiscardUnknown(m)
}

var xxx_messageInfo_OverloadManager proto.InternalMessageInfo

func (m *OverloadManager) GetRefreshInterval() *duration.Duration {
	if m != nil {
		return m.RefreshInterval
	}
	return nil
}

func (m *OverloadManager) GetResourceMonitors() []*ResourceMonitor {
	if m != nil {
		return m.ResourceMonitors
	}
	return nil
}

func (m *OverloadManager) GetActions() []*OverloadAction {
	if m != nil {
		return m.Actions
	}
	return nil
}

func init() {
	proto.RegisterType((*ResourceMonitor)(nil), "envoy.config.overload.v3alpha.ResourceMonitor")
	proto.RegisterType((*ThresholdTrigger)(nil), "envoy.config.overload.v3alpha.ThresholdTrigger")
	proto.RegisterType((*Trigger)(nil), "envoy.config.overload.v3alpha.Trigger")
	proto.RegisterType((*OverloadAction)(nil), "envoy.config.overload.v3alpha.OverloadAction")
	proto.RegisterType((*OverloadManager)(nil), "envoy.config.overload.v3alpha.OverloadManager")
}

func init() {
	proto.RegisterFile("envoy/config/overload/v3alpha/overload.proto", fileDescriptor_542f910eb6259f77)
}

var fileDescriptor_542f910eb6259f77 = []byte{
	// 567 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xae, 0xd3, 0x6e, 0xed, 0x5c, 0x4a, 0x4b, 0x54, 0x69, 0xed, 0x80, 0xa9, 0x54, 0x02, 0x8a,
	0x58, 0x1d, 0xd4, 0xc2, 0x81, 0x1e, 0x40, 0x0b, 0x93, 0x18, 0x88, 0x69, 0x53, 0xb4, 0x7b, 0xe5,
	0x35, 0x6e, 0x1a, 0x29, 0xb3, 0x2b, 0xc7, 0x89, 0xd6, 0x7f, 0xc0, 0x99, 0x23, 0xbf, 0x82, 0x3b,
	0x70, 0xe4, 0xff, 0x70, 0x44, 0x3d, 0xa1, 0xd8, 0x4e, 0xa7, 0x66, 0x5a, 0xba, 0x9c, 0xa2, 0xf7,
	0xbe, 0xef, 0xbd, 0xef, 0xfb, 0xf4, 0x0c, 0x0f, 0x08, 0x8d, 0xd9, 0xc2, 0x9a, 0x30, 0x3a, 0xf5,
	0x3d, 0x8b, 0xc5, 0x84, 0x07, 0x0c, 0xbb, 0x56, 0x3c, 0xc4, 0xc1, 0x7c, 0x86, 0x57, 0x05, 0x34,
	0xe7, 0x4c, 0x30, 0xf3, 0xb1, 0x44, 0x23, 0x85, 0x46, 0xab, 0xa6, 0x46, 0xef, 0xb5, 0x3d, 0xc6,
	0xbc, 0x80, 0x58, 0x12, 0x7c, 0x11, 0x4d, 0x2d, 0x4c, 0x17, 0x8a, 0xb9, 0xb7, 0x9f, 0x6d, 0xb9,
	0x11, 0xc7, 0xc2, 0x67, 0x54, 0xf7, 0x1f, 0x65, 0xfb, 0xa1, 0xe0, 0xd1, 0x44, 0xe8, 0xee, 0x93,
	0xc8, 0x9d, 0x63, 0x0b, 0x53, 0xca, 0x84, 0x24, 0x85, 0x56, 0x4c, 0x78, 0xe8, 0x33, 0xea, 0x53,
	0x4f, 0x43, 0x76, 0x63, 0x1c, 0xf8, 0x2e, 0x16, 0xc4, 0x4a, 0x7f, 0x54, 0xa3, 0xfb, 0x0b, 0xc0,
	0xba, 0x43, 0x42, 0x16, 0xf1, 0x09, 0x39, 0x61, 0xd4, 0x17, 0x8c, 0x9b, 0x0f, 0x61, 0x89, 0xe2,
	0x4b, 0xd2, 0x02, 0x1d, 0xd0, 0xdb, 0xb1, 0xcb, 0x4b, 0xbb, 0xc4, 0x8d, 0x0e, 0x70, 0x64, 0xd1,
	0x7c, 0x0b, 0xef, 0x89, 0xc5, 0x9c, 0xb8, 0x63, 0x65, 0xb3, 0x55, 0xec, 0x80, 0x5e, 0x75, 0xd0,
	0x44, 0x4a, 0x21, 0x4a, 0x15, 0xa2, 0x43, 0xba, 0x38, 0x2e, 0x38, 0x55, 0x89, 0xfd, 0x20, 0xa1,
	0xa3, 0xd7, 0xdf, 0xff, 0x7c, 0xdd, 0xb7, 0x60, 0xff, 0x96, 0x98, 0x06, 0x32, 0x26, 0x94, 0x51,
	0x63, 0xd7, 0x60, 0x55, 0x41, 0xc7, 0xc9, 0xac, 0xcf, 0xa5, 0x8a, 0xd1, 0x28, 0x3a, 0xdb, 0xaa,
	0xd4, 0xbd, 0x82, 0x8d, 0xf3, 0x19, 0x27, 0xe1, 0x8c, 0x05, 0xee, 0x39, 0xf7, 0x3d, 0x8f, 0x70,
	0xb3, 0x0f, 0xb7, 0x62, 0x1c, 0x44, 0x4a, 0x3f, 0xb0, 0x77, 0x97, 0x76, 0xd3, 0x34, 0xdb, 0x05,
	0xf9, 0xfd, 0x7d, 0xff, 0xa2, 0xa0, 0x3f, 0x47, 0xa1, 0x46, 0x6f, 0x12, 0x55, 0xaf, 0x20, 0xca,
	0x57, 0x95, 0xdd, 0xd2, 0xfd, 0x09, 0x60, 0x39, 0xdd, 0x98, 0x1b, 0xd8, 0x29, 0xdc, 0x11, 0x29,
	0xb9, 0x65, 0xc8, 0xb4, 0x2c, 0x94, 0x7b, 0x29, 0x37, 0x96, 0x1d, 0x17, 0x9c, 0xeb, 0x19, 0xa3,
	0x83, 0x44, 0xf0, 0x73, 0xf8, 0x74, 0x83, 0x60, 0x45, 0xb5, 0x9b, 0xb0, 0x26, 0xd4, 0xef, 0x98,
	0x51, 0xc2, 0xa6, 0x66, 0xf1, 0x9f, 0x0d, 0xba, 0x3f, 0x00, 0xbc, 0x7f, 0xaa, 0x29, 0x87, 0x93,
	0xe4, 0x68, 0xf2, 0x4d, 0x7c, 0x81, 0x15, 0x3d, 0x25, 0x6c, 0x19, 0x9d, 0x62, 0xaf, 0x3a, 0x78,
	0xb6, 0xc9, 0x83, 0xde, 0x5f, 0x59, 0xda, 0x5b, 0xdf, 0x80, 0x51, 0x01, 0xce, 0x6a, 0xc2, 0x68,
	0x98, 0x38, 0x40, 0xfa, 0x75, 0xdd, 0xea, 0x60, 0x5d, 0x5f, 0xf7, 0xb7, 0x01, 0xeb, 0x69, 0xe9,
	0x04, 0x53, 0x9c, 0x04, 0x7f, 0x04, 0x1b, 0x9c, 0x4c, 0x93, 0x60, 0xc6, 0x3e, 0x15, 0x84, 0xc7,
	0x38, 0x90, 0xfa, 0xab, 0x83, 0xf6, 0x8d, 0x83, 0x3c, 0xd2, 0x4f, 0xca, 0xa9, 0x6b, 0xca, 0x27,
	0xcd, 0x30, 0x09, 0x7c, 0xc0, 0xf5, 0xd1, 0x8d, 0x2f, 0xd5, 0xd5, 0xa5, 0x2e, 0xd1, 0x06, 0x97,
	0xd9, 0x63, 0xbd, 0x76, 0xdb, 0xe0, 0xeb, 0xad, 0xd0, 0xfc, 0x08, 0xcb, 0x58, 0x5a, 0x09, 0x5b,
	0x45, 0x39, 0xbc, 0xbf, 0x61, 0xf8, 0x7a, 0x00, 0x4e, 0xca, 0xbe, 0xe3, 0x3b, 0xca, 0x64, 0x65,
	0xbf, 0x83, 0x2f, 0x7d, 0xa6, 0x36, 0xce, 0x39, 0xbb, 0x5a, 0xe4, 0x2f, 0xb7, 0x6b, 0x29, 0xff,
	0x2c, 0x09, 0xf0, 0x0c, 0x5c, 0x6c, 0xcb, 0x24, 0x87, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0xac,
	0xe1, 0x87, 0x25, 0x14, 0x05, 0x00, 0x00,
}
