// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/admin/v3alpha/mutex_stats.proto

package envoy_admin_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

type MutexStats struct {
	NumContentions       uint64   `protobuf:"varint,1,opt,name=num_contentions,json=numContentions,proto3" json:"num_contentions,omitempty"`
	CurrentWaitCycles    uint64   `protobuf:"varint,2,opt,name=current_wait_cycles,json=currentWaitCycles,proto3" json:"current_wait_cycles,omitempty"`
	LifetimeWaitCycles   uint64   `protobuf:"varint,3,opt,name=lifetime_wait_cycles,json=lifetimeWaitCycles,proto3" json:"lifetime_wait_cycles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutexStats) Reset()         { *m = MutexStats{} }
func (m *MutexStats) String() string { return proto.CompactTextString(m) }
func (*MutexStats) ProtoMessage()    {}
func (*MutexStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0ad2bf347536df0, []int{0}
}

func (m *MutexStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutexStats.Unmarshal(m, b)
}
func (m *MutexStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutexStats.Marshal(b, m, deterministic)
}
func (m *MutexStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutexStats.Merge(m, src)
}
func (m *MutexStats) XXX_Size() int {
	return xxx_messageInfo_MutexStats.Size(m)
}
func (m *MutexStats) XXX_DiscardUnknown() {
	xxx_messageInfo_MutexStats.DiscardUnknown(m)
}

var xxx_messageInfo_MutexStats proto.InternalMessageInfo

func (m *MutexStats) GetNumContentions() uint64 {
	if m != nil {
		return m.NumContentions
	}
	return 0
}

func (m *MutexStats) GetCurrentWaitCycles() uint64 {
	if m != nil {
		return m.CurrentWaitCycles
	}
	return 0
}

func (m *MutexStats) GetLifetimeWaitCycles() uint64 {
	if m != nil {
		return m.LifetimeWaitCycles
	}
	return 0
}

func init() {
	proto.RegisterType((*MutexStats)(nil), "envoy.admin.v3alpha.MutexStats")
}

func init() {
	proto.RegisterFile("envoy/admin/v3alpha/mutex_stats.proto", fileDescriptor_d0ad2bf347536df0)
}

var fileDescriptor_d0ad2bf347536df0 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x4a, 0x03, 0x41,
	0x10, 0x86, 0x39, 0x15, 0x8b, 0x2d, 0x0c, 0x6e, 0x2c, 0x82, 0x45, 0x48, 0x84, 0xa0, 0xd5, 0xae,
	0x18, 0xb0, 0xb0, 0x4c, 0x6a, 0x21, 0x68, 0x61, 0x79, 0xac, 0x97, 0x55, 0x07, 0xb2, 0x33, 0xc7,
	0xee, 0xec, 0x99, 0x7b, 0x03, 0x9f, 0xc1, 0x87, 0xb1, 0xf3, 0xbd, 0xe4, 0xc6, 0xc8, 0x11, 0xb0,
	0xdd, 0xff, 0xfb, 0x7e, 0xf6, 0x1f, 0x35, 0xf3, 0xd8, 0x50, 0x6b, 0xdd, 0x3a, 0x00, 0xda, 0x66,
	0xee, 0x36, 0xf5, 0x9b, 0xb3, 0x21, 0xb3, 0xdf, 0x96, 0x89, 0x1d, 0x27, 0x53, 0x47, 0x62, 0xd2,
	0x43, 0xc1, 0x8c, 0x60, 0x66, 0x87, 0x9d, 0x4f, 0xf3, 0xba, 0x76, 0xd6, 0x21, 0x12, 0x3b, 0x06,
	0xc2, 0x64, 0x1b, 0x1f, 0x13, 0x10, 0x02, 0xbe, 0xfe, 0x7a, 0x17, 0x5f, 0x85, 0x52, 0xf7, 0x5d,
	0xdb, 0x63, 0x57, 0xa6, 0x2f, 0xd5, 0x00, 0x73, 0x28, 0x2b, 0x42, 0xf6, 0x28, 0xca, 0xa8, 0x98,
	0x14, 0x57, 0x47, 0x0f, 0x27, 0x98, 0xc3, 0xb2, 0x7f, 0xd5, 0x46, 0x0d, 0xab, 0x1c, 0xa3, 0x47,
	0x2e, 0xdf, 0x1d, 0x70, 0x59, 0xb5, 0xd5, 0xc6, 0xa7, 0xd1, 0x81, 0xc0, 0xa7, 0xbb, 0xe8, 0xc9,
	0x01, 0x2f, 0x25, 0xd0, 0xd7, 0xea, 0x6c, 0x03, 0x2f, 0x9e, 0x21, 0xf8, 0x3d, 0xe1, 0x50, 0x04,
	0xfd, 0x97, 0xf5, 0xc6, 0xdd, 0xec, 0xf3, 0xfb, 0x63, 0x3c, 0x51, 0xe3, 0xbd, 0x61, 0x37, 0x32,
	0xcc, 0xf4, 0x3f, 0x5e, 0xdc, 0xaa, 0x29, 0x90, 0x11, 0xa8, 0x8e, 0xb4, 0x6d, 0xcd, 0x3f, 0x87,
	0x58, 0x0c, 0x7a, 0x61, 0xd5, 0xcd, 0x5e, 0x15, 0xcf, 0xc7, 0xb2, 0x7f, 0xfe, 0x13, 0x00, 0x00,
	0xff, 0xff, 0x3a, 0x40, 0xaa, 0x9d, 0x60, 0x01, 0x00, 0x00,
}
