// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: imports/test_import_all.proto

package imports // import "github.com/gogo/protobuf/protoc-gen-gogo/testdata/imports"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import fmt1 "github.com/gogo/protobuf/protoc-gen-gogo/testdata/imports/fmt"
import test_a_1 "github.com/gogo/protobuf/protoc-gen-gogo/testdata/imports/test_a_1"
import test_a_2 "github.com/gogo/protobuf/protoc-gen-gogo/testdata/imports/test_a_2"
import test_b_1 "github.com/gogo/protobuf/protoc-gen-gogo/testdata/imports/test_b_1"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type All struct {
	Am1                  *test_a_1.M1 `protobuf:"bytes,1,opt,name=am1" json:"am1,omitempty"`
	Am2                  *test_a_1.M2 `protobuf:"bytes,2,opt,name=am2" json:"am2,omitempty"`
	Am3                  *test_a_2.M3 `protobuf:"bytes,3,opt,name=am3" json:"am3,omitempty"`
	Am4                  *test_a_2.M4 `protobuf:"bytes,4,opt,name=am4" json:"am4,omitempty"`
	Bm1                  *test_b_1.M1 `protobuf:"bytes,5,opt,name=bm1" json:"bm1,omitempty"`
	Bm2                  *test_b_1.M2 `protobuf:"bytes,6,opt,name=bm2" json:"bm2,omitempty"`
	Fmt                  *fmt1.M      `protobuf:"bytes,7,opt,name=fmt" json:"fmt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *All) Reset()         { *m = All{} }
func (m *All) String() string { return proto.CompactTextString(m) }
func (*All) ProtoMessage()    {}
func (*All) Descriptor() ([]byte, []int) {
	return fileDescriptor_test_import_all_a07d58de416f602a, []int{0}
}
func (m *All) XXX_Unmarshal(b []byte) error {
	if m, ok := (interface{})(m).(proto.Unmarshaler); ok {
		return m.Unmarshal(b)
	}
	return xxx_messageInfo_All.Unmarshal(m, b)
}
func (m *All) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if m, ok := (interface{})(m).(proto.Marshaler); ok {
		return m.Marshal()
	}
	return xxx_messageInfo_All.Marshal(b, m, deterministic)
}
func (dst *All) XXX_Merge(src proto.Message) {
	xxx_messageInfo_All.Merge(dst, src)
}
func (m *All) XXX_Size() int {
	if m, ok := (interface{})(m).(proto.Sizer); ok {
		return m.Size()
	}
	return xxx_messageInfo_All.Size(m)
}
func (m *All) XXX_DiscardUnknown() {
	xxx_messageInfo_All.DiscardUnknown(m)
}

var xxx_messageInfo_All proto.InternalMessageInfo

func (m *All) GetAm1() *test_a_1.M1 {
	if m != nil {
		return m.Am1
	}
	return nil
}

func (m *All) GetAm2() *test_a_1.M2 {
	if m != nil {
		return m.Am2
	}
	return nil
}

func (m *All) GetAm3() *test_a_2.M3 {
	if m != nil {
		return m.Am3
	}
	return nil
}

func (m *All) GetAm4() *test_a_2.M4 {
	if m != nil {
		return m.Am4
	}
	return nil
}

func (m *All) GetBm1() *test_b_1.M1 {
	if m != nil {
		return m.Bm1
	}
	return nil
}

func (m *All) GetBm2() *test_b_1.M2 {
	if m != nil {
		return m.Bm2
	}
	return nil
}

func (m *All) GetFmt() *fmt1.M {
	if m != nil {
		return m.Fmt
	}
	return nil
}

func init() {
	proto.RegisterType((*All)(nil), "test.All")
}

func init() {
	proto.RegisterFile("imports/test_import_all.proto", fileDescriptor_test_import_all_a07d58de416f602a)
}

var fileDescriptor_test_import_all_a07d58de416f602a = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xd0, 0x31, 0x4f, 0xf3, 0x30,
	0x10, 0x06, 0x60, 0xe5, 0x73, 0xbf, 0x20, 0x99, 0x05, 0x85, 0xc5, 0x20, 0x90, 0x50, 0x27, 0x96,
	0xda, 0xb2, 0x9d, 0x05, 0x31, 0xc1, 0xde, 0xa5, 0x23, 0x4b, 0xe4, 0x2b, 0x4d, 0xa8, 0x94, 0xc3,
	0x51, 0x7a, 0xfd, 0xbd, 0xfc, 0x15, 0x64, 0x1f, 0x48, 0x10, 0x9a, 0x2d, 0x79, 0x9f, 0xd7, 0x3e,
	0xdb, 0xf2, 0x76, 0x8f, 0x43, 0x1c, 0xe9, 0x60, 0x68, 0x77, 0xa0, 0x86, 0x7f, 0x9a, 0xd0, 0xf7,
	0x7a, 0x18, 0x23, 0xc5, 0x6a, 0x91, 0xe2, 0xeb, 0xab, 0x5f, 0xa5, 0xd0, 0x58, 0x83, 0x96, 0x0b,
	0xa7, 0xc8, 0xcd, 0x90, 0x33, 0xe8, 0xe7, 0xa9, 0x3e, 0x49, 0x30, 0x3f, 0x0b, 0x7e, 0xce, 0xba,
	0xfc, 0xa6, 0x16, 0xc9, 0x20, 0x87, 0xcb, 0x8f, 0x42, 0x8a, 0xa7, 0xbe, 0xaf, 0x6e, 0xa4, 0x08,
	0x68, 0x55, 0x71, 0x57, 0xdc, 0x9f, 0x3b, 0xa9, 0xd3, 0x6a, 0x1d, 0xf4, 0xda, 0x6e, 0x52, 0xcc,
	0xea, 0xd4, 0xbf, 0x89, 0xba, 0xa4, 0x8e, 0xd5, 0x2b, 0x31, 0x51, 0x9f, 0xd4, 0xb3, 0xd6, 0x6a,
	0x31, 0xd1, 0x3a, 0x69, 0x5d, 0x2d, 0xa5, 0x00, 0xb4, 0xea, 0x7f, 0xd6, 0x0b, 0x56, 0xd0, 0x43,
	0x18, 0xc9, 0xe6, 0xe9, 0x80, 0x96, 0x3b, 0x4e, 0x95, 0x7f, 0x3b, 0x2e, 0x9f, 0x01, 0xd0, 0x55,
	0x4a, 0x8a, 0x16, 0x49, 0x9d, 0xe5, 0x4e, 0xa9, 0x5b, 0x24, 0xbd, 0xde, 0xa4, 0xe8, 0xf9, 0xf1,
	0xe5, 0xa1, 0xdb, 0xd3, 0xdb, 0x11, 0xf4, 0x36, 0xa2, 0xe9, 0x62, 0x17, 0x4d, 0xbe, 0x3a, 0x1c,
	0x5b, 0xfe, 0xd8, 0xae, 0xba, 0xdd, 0xfb, 0x2a, 0x43, 0xda, 0xfa, 0x35, 0x50, 0x30, 0x5f, 0x4f,
	0x05, 0x65, 0x6e, 0xf8, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7f, 0xd4, 0x5c, 0x7f, 0x03, 0x02,
	0x00, 0x00,
}
