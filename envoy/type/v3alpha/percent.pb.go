// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/v3alpha/percent.proto

package envoy_type_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

type FractionalPercent_DenominatorType int32

const (
	FractionalPercent_HUNDRED      FractionalPercent_DenominatorType = 0
	FractionalPercent_TEN_THOUSAND FractionalPercent_DenominatorType = 1
	FractionalPercent_MILLION      FractionalPercent_DenominatorType = 2
)

var FractionalPercent_DenominatorType_name = map[int32]string{
	0: "HUNDRED",
	1: "TEN_THOUSAND",
	2: "MILLION",
}

var FractionalPercent_DenominatorType_value = map[string]int32{
	"HUNDRED":      0,
	"TEN_THOUSAND": 1,
	"MILLION":      2,
}

func (x FractionalPercent_DenominatorType) String() string {
	return proto.EnumName(FractionalPercent_DenominatorType_name, int32(x))
}

func (FractionalPercent_DenominatorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2c71426cf14750f3, []int{1, 0}
}

type Percent struct {
	Value                float64  `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Percent) Reset()         { *m = Percent{} }
func (m *Percent) String() string { return proto.CompactTextString(m) }
func (*Percent) ProtoMessage()    {}
func (*Percent) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c71426cf14750f3, []int{0}
}

func (m *Percent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Percent.Unmarshal(m, b)
}
func (m *Percent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Percent.Marshal(b, m, deterministic)
}
func (m *Percent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Percent.Merge(m, src)
}
func (m *Percent) XXX_Size() int {
	return xxx_messageInfo_Percent.Size(m)
}
func (m *Percent) XXX_DiscardUnknown() {
	xxx_messageInfo_Percent.DiscardUnknown(m)
}

var xxx_messageInfo_Percent proto.InternalMessageInfo

func (m *Percent) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type FractionalPercent struct {
	Numerator            uint32                            `protobuf:"varint,1,opt,name=numerator,proto3" json:"numerator,omitempty"`
	Denominator          FractionalPercent_DenominatorType `protobuf:"varint,2,opt,name=denominator,proto3,enum=envoy.type.v3alpha.FractionalPercent_DenominatorType" json:"denominator,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *FractionalPercent) Reset()         { *m = FractionalPercent{} }
func (m *FractionalPercent) String() string { return proto.CompactTextString(m) }
func (*FractionalPercent) ProtoMessage()    {}
func (*FractionalPercent) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c71426cf14750f3, []int{1}
}

func (m *FractionalPercent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FractionalPercent.Unmarshal(m, b)
}
func (m *FractionalPercent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FractionalPercent.Marshal(b, m, deterministic)
}
func (m *FractionalPercent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FractionalPercent.Merge(m, src)
}
func (m *FractionalPercent) XXX_Size() int {
	return xxx_messageInfo_FractionalPercent.Size(m)
}
func (m *FractionalPercent) XXX_DiscardUnknown() {
	xxx_messageInfo_FractionalPercent.DiscardUnknown(m)
}

var xxx_messageInfo_FractionalPercent proto.InternalMessageInfo

func (m *FractionalPercent) GetNumerator() uint32 {
	if m != nil {
		return m.Numerator
	}
	return 0
}

func (m *FractionalPercent) GetDenominator() FractionalPercent_DenominatorType {
	if m != nil {
		return m.Denominator
	}
	return FractionalPercent_HUNDRED
}

func init() {
	proto.RegisterEnum("envoy.type.v3alpha.FractionalPercent_DenominatorType", FractionalPercent_DenominatorType_name, FractionalPercent_DenominatorType_value)
	proto.RegisterType((*Percent)(nil), "envoy.type.v3alpha.Percent")
	proto.RegisterType((*FractionalPercent)(nil), "envoy.type.v3alpha.FractionalPercent")
}

func init() { proto.RegisterFile("envoy/type/v3alpha/percent.proto", fileDescriptor_2c71426cf14750f3) }

var fileDescriptor_2c71426cf14750f3 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x4b, 0x02, 0x41,
	0x14, 0xc7, 0x1d, 0xc1, 0xac, 0xd1, 0x6a, 0x1b, 0x04, 0x53, 0x44, 0x36, 0xbb, 0xd8, 0xa1, 0x5d,
	0x50, 0xba, 0x08, 0x41, 0x2d, 0x6b, 0x28, 0xd8, 0x2a, 0xab, 0x1e, 0x3a, 0xc5, 0x4b, 0x87, 0x5a,
	0x58, 0x67, 0x86, 0x71, 0x5c, 0xda, 0x6b, 0xa7, 0xfe, 0x86, 0xfe, 0x9f, 0xfe, 0xa7, 0xf0, 0x14,
	0xfb, 0x43, 0x94, 0xec, 0x9d, 0x86, 0x79, 0x9f, 0xf7, 0x7d, 0xdf, 0x2f, 0x0f, 0xeb, 0x94, 0x05,
	0x3c, 0x34, 0x55, 0x28, 0xa8, 0x19, 0xb4, 0xc1, 0x17, 0x6f, 0x60, 0x0a, 0x2a, 0x67, 0x94, 0x29,
	0x43, 0x48, 0xae, 0x38, 0x21, 0x31, 0x61, 0x44, 0x84, 0x91, 0x12, 0xd5, 0x8b, 0xd5, 0x5c, 0x80,
	0x09, 0x8c, 0x71, 0x05, 0xca, 0xe3, 0x6c, 0x69, 0x06, 0x54, 0x2e, 0x3d, 0xce, 0x3c, 0xf6, 0x9a,
	0x8c, 0x55, 0xcb, 0x01, 0xf8, 0xde, 0x1c, 0x14, 0x35, 0x37, 0x8f, 0xa4, 0xd1, 0x18, 0xe3, 0xfc,
	0x28, 0x59, 0x40, 0xae, 0x71, 0x2e, 0x00, 0x7f, 0x45, 0xcf, 0x91, 0x8e, 0x9a, 0xc8, 0x2a, 0xaf,
	0xad, 0x12, 0x21, 0x95, 0x4c, 0x5c, 0x4f, 0x77, 0x57, 0x99, 0xb4, 0xdc, 0x84, 0xea, 0x54, 0xbe,
	0xbe, 0x3f, 0xeb, 0x25, 0xbc, 0x6b, 0x28, 0x55, 0x6a, 0xfc, 0x20, 0x7c, 0xf6, 0x20, 0x61, 0x16,
	0x99, 0x01, 0x7f, 0xa3, 0x5f, 0xc3, 0x47, 0x6c, 0xb5, 0xa0, 0x12, 0x14, 0x97, 0xf1, 0x8e, 0x63,
	0x77, 0xfb, 0x41, 0x00, 0x17, 0xe6, 0x94, 0xf1, 0x85, 0xc7, 0xe2, 0x7e, 0x56, 0x47, 0xcd, 0x93,
	0xd6, 0x8d, 0xb1, 0x1f, 0xd7, 0xd8, 0x53, 0x36, 0xec, 0xed, 0xe0, 0x24, 0x14, 0xd4, 0x3a, 0x5c,
	0x5b, 0xb9, 0x0f, 0x94, 0xd5, 0x90, 0xbb, 0xab, 0xd9, 0xb8, 0xc5, 0xa7, 0x7f, 0x48, 0x52, 0xc0,
	0xf9, 0xde, 0xd4, 0xb1, 0xdd, 0xae, 0xad, 0x65, 0x88, 0x86, 0x8b, 0x93, 0xae, 0xf3, 0x3c, 0xe9,
	0x0d, 0xa7, 0xe3, 0x7b, 0xc7, 0xd6, 0x50, 0xd4, 0x7e, 0xec, 0x0f, 0x06, 0xfd, 0xa1, 0xa3, 0x65,
	0x3b, 0x97, 0x51, 0xe0, 0x3a, 0xae, 0xed, 0x58, 0xda, 0xb3, 0x62, 0xb5, 0xb0, 0xee, 0xf1, 0xc4,
	0xb5, 0x90, 0xfc, 0x3d, 0xfc, 0x27, 0x80, 0x55, 0x4c, 0xe1, 0x51, 0x74, 0x81, 0x11, 0x7a, 0x39,
	0x88, 0x4f, 0xd1, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x34, 0x1d, 0xb1, 0xa5, 0xfe, 0x01, 0x00,
	0x00,
}
