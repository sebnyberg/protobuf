// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: imports/test_a_1/m2.proto

package test_a_1 // import "github.com/gogo/protobuf/protoc-gen-gogo/testdata/imports/test_a_1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type M2 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M2) Reset()         { *m = M2{} }
func (m *M2) String() string { return proto.CompactTextString(m) }
func (*M2) ProtoMessage()    {}
func (*M2) Descriptor() ([]byte, []int) {
	return fileDescriptor_m2_d5c8bd8077345106, []int{0}
}
func (m *M2) XXX_Unmarshal(b []byte) error {
	if m, ok := (interface{})(m).(proto.Unmarshaler); ok {
		return m.Unmarshal(b)
	}
	return xxx_messageInfo_M2.Unmarshal(m, b)
}
func (m *M2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if m, ok := (interface{})(m).(proto.Marshaler); ok {
		return m.Marshal()
	}
	return xxx_messageInfo_M2.Marshal(b, m, deterministic)
}
func (dst *M2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M2.Merge(dst, src)
}
func (m *M2) XXX_Size() int {
	if m, ok := (interface{})(m).(proto.Sizer); ok {
		return m.Size()
	}
	return xxx_messageInfo_M2.Size(m)
}
func (m *M2) XXX_DiscardUnknown() {
	xxx_messageInfo_M2.DiscardUnknown(m)
}

var xxx_messageInfo_M2 proto.InternalMessageInfo

func init() {
	proto.RegisterType((*M2)(nil), "test.a.M2")
}

func init() { proto.RegisterFile("imports/test_a_1/m2.proto", fileDescriptor_m2_d5c8bd8077345106) }

var fileDescriptor_m2_d5c8bd8077345106 = []byte{
	// 112 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0xcc, 0x2d, 0xc8,
	0x2f, 0x2a, 0x29, 0xd6, 0x2f, 0x49, 0x2d, 0x2e, 0x89, 0x4f, 0x8c, 0x37, 0xd4, 0xcf, 0x35, 0xd2,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x09, 0xe9, 0x25, 0x2a, 0xb1, 0x70, 0x31, 0xf9,
	0x1a, 0x39, 0xb9, 0x44, 0x39, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea,
	0xa7, 0xe7, 0xa7, 0xe7, 0xeb, 0x83, 0x95, 0x25, 0x95, 0xa6, 0x41, 0x18, 0xc9, 0xba, 0xe9, 0xa9,
	0x79, 0xba, 0x60, 0x09, 0x90, 0xc6, 0x94, 0xc4, 0x92, 0x44, 0x7d, 0x74, 0xc3, 0x93, 0xd8, 0xc0,
	0x4a, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x9b, 0x89, 0x4c, 0x77, 0x00, 0x00, 0x00,
}
