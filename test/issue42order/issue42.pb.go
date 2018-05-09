// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: issue42.proto

package issue42

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import encoding_binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type UnorderedFields struct {
	A                    *int64   `protobuf:"varint,10,opt,name=A" json:"A,omitempty"`
	B                    *uint64  `protobuf:"fixed64,1,opt,name=B" json:"B,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnorderedFields) Reset()         { *m = UnorderedFields{} }
func (m *UnorderedFields) String() string { return proto.CompactTextString(m) }
func (*UnorderedFields) ProtoMessage()    {}
func (*UnorderedFields) Descriptor() ([]byte, []int) {
	return fileDescriptor_issue42_6157ac17a2848d4f, []int{0}
}
func (m *UnorderedFields) XXX_Unmarshal(b []byte) error {
	if m, ok := (interface{})(m).(proto.Unmarshaler); ok {
		return m.Unmarshal(b)
	}
	return xxx_messageInfo_UnorderedFields.Unmarshal(m, b)
}
func (m *UnorderedFields) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if m, ok := (interface{})(m).(proto.Marshaler); ok {
		return m.Marshal()
	}
	return xxx_messageInfo_UnorderedFields.Marshal(b, m, deterministic)
}
func (dst *UnorderedFields) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnorderedFields.Merge(dst, src)
}
func (m *UnorderedFields) XXX_Size() int {
	if m, ok := (interface{})(m).(proto.Sizer); ok {
		return m.Size()
	}
	return xxx_messageInfo_UnorderedFields.Size(m)
}
func (m *UnorderedFields) XXX_DiscardUnknown() {
	xxx_messageInfo_UnorderedFields.DiscardUnknown(m)
}

var xxx_messageInfo_UnorderedFields proto.InternalMessageInfo

func (m *UnorderedFields) GetA() int64 {
	if m != nil && m.A != nil {
		return *m.A
	}
	return 0
}

func (m *UnorderedFields) GetB() uint64 {
	if m != nil && m.B != nil {
		return *m.B
	}
	return 0
}

type OrderedFields struct {
	B                    *uint64  `protobuf:"fixed64,1,opt,name=B" json:"B,omitempty"`
	A                    *int64   `protobuf:"varint,10,opt,name=A" json:"A,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderedFields) Reset()         { *m = OrderedFields{} }
func (m *OrderedFields) String() string { return proto.CompactTextString(m) }
func (*OrderedFields) ProtoMessage()    {}
func (*OrderedFields) Descriptor() ([]byte, []int) {
	return fileDescriptor_issue42_6157ac17a2848d4f, []int{1}
}
func (m *OrderedFields) XXX_Unmarshal(b []byte) error {
	if m, ok := (interface{})(m).(proto.Unmarshaler); ok {
		return m.Unmarshal(b)
	}
	return xxx_messageInfo_OrderedFields.Unmarshal(m, b)
}
func (m *OrderedFields) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if m, ok := (interface{})(m).(proto.Marshaler); ok {
		return m.Marshal()
	}
	return xxx_messageInfo_OrderedFields.Marshal(b, m, deterministic)
}
func (dst *OrderedFields) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderedFields.Merge(dst, src)
}
func (m *OrderedFields) XXX_Size() int {
	if m, ok := (interface{})(m).(proto.Sizer); ok {
		return m.Size()
	}
	return xxx_messageInfo_OrderedFields.Size(m)
}
func (m *OrderedFields) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderedFields.DiscardUnknown(m)
}

var xxx_messageInfo_OrderedFields proto.InternalMessageInfo

func (m *OrderedFields) GetB() uint64 {
	if m != nil && m.B != nil {
		return *m.B
	}
	return 0
}

func (m *OrderedFields) GetA() int64 {
	if m != nil && m.A != nil {
		return *m.A
	}
	return 0
}

func init() {
	proto.RegisterType((*UnorderedFields)(nil), "issue42.UnorderedFields")
	proto.RegisterType((*OrderedFields)(nil), "issue42.OrderedFields")
}
func (m *UnorderedFields) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnorderedFields) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.B != nil {
		dAtA[i] = 0x9
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(*m.B))
		i += 8
	}
	if m.A != nil {
		dAtA[i] = 0x50
		i++
		i = encodeVarintIssue42(dAtA, i, uint64(*m.A))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *OrderedFields) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OrderedFields) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.B != nil {
		dAtA[i] = 0x9
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(*m.B))
		i += 8
	}
	if m.A != nil {
		dAtA[i] = 0x50
		i++
		i = encodeVarintIssue42(dAtA, i, uint64(*m.A))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintIssue42(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedUnorderedFields(r randyIssue42, easy bool) *UnorderedFields {
	this := &UnorderedFields{}
	if r.Intn(10) != 0 {
		v1 := uint64(uint64(r.Uint32()))
		this.B = &v1
	}
	if r.Intn(10) != 0 {
		v2 := int64(r.Int63())
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		this.A = &v2
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedIssue42(r, 11)
	}
	return this
}

func NewPopulatedOrderedFields(r randyIssue42, easy bool) *OrderedFields {
	this := &OrderedFields{}
	if r.Intn(10) != 0 {
		v3 := uint64(uint64(r.Uint32()))
		this.B = &v3
	}
	if r.Intn(10) != 0 {
		v4 := int64(r.Int63())
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		this.A = &v4
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedIssue42(r, 11)
	}
	return this
}

type randyIssue42 interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneIssue42(r randyIssue42) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringIssue42(r randyIssue42) string {
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
		tmps[i] = randUTF8RuneIssue42(r)
	}
	return string(tmps)
}
func randUnrecognizedIssue42(r randyIssue42, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldIssue42(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldIssue42(dAtA []byte, r randyIssue42, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateIssue42(dAtA, uint64(key))
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		dAtA = encodeVarintPopulateIssue42(dAtA, uint64(v6))
	case 1:
		dAtA = encodeVarintPopulateIssue42(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateIssue42(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateIssue42(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateIssue42(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateIssue42(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *UnorderedFields) Size() (n int) {
	var l int
	_ = l
	if m.B != nil {
		n += 9
	}
	if m.A != nil {
		n += 1 + sovIssue42(uint64(*m.A))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *OrderedFields) Size() (n int) {
	var l int
	_ = l
	if m.B != nil {
		n += 9
	}
	if m.A != nil {
		n += 1 + sovIssue42(uint64(*m.A))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovIssue42(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozIssue42(x uint64) (n int) {
	return sovIssue42(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UnorderedFields) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIssue42
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UnorderedFields: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnorderedFields: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field B", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.B = &v
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue42
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.A = &v
		default:
			iNdEx = preIndex
			skippy, err := skipIssue42(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIssue42
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OrderedFields) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIssue42
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OrderedFields: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OrderedFields: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field B", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.B = &v
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field A", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue42
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.A = &v
		default:
			iNdEx = preIndex
			skippy, err := skipIssue42(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIssue42
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipIssue42(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIssue42
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowIssue42
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowIssue42
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthIssue42
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowIssue42
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipIssue42(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthIssue42 = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIssue42   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("issue42.proto", fileDescriptor_issue42_6157ac17a2848d4f) }

var fileDescriptor_issue42_6157ac17a2848d4f = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x2c, 0x2e, 0x2e,
	0x4d, 0x35, 0x31, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0xa5, 0x74, 0xd3,
	0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0xd3, 0xf3, 0xf5, 0xc1,
	0xf2, 0x49, 0xa5, 0x69, 0x60, 0x1e, 0x98, 0x03, 0x66, 0x41, 0xf4, 0x29, 0xe9, 0x72, 0xf1, 0x87,
	0xe6, 0xe5, 0x17, 0xa5, 0xa4, 0x16, 0xa5, 0xa6, 0xb8, 0x65, 0xa6, 0xe6, 0xa4, 0x14, 0x0b, 0xf1,
	0x70, 0x31, 0x3a, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x05, 0x31, 0x3a, 0x81, 0x78, 0x8e, 0x12,
	0x5c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x8c, 0x8e, 0x4a, 0xda, 0x5c, 0xbc, 0xfe, 0xc4, 0x2a, 0x76,
	0x12, 0xf8, 0xf1, 0x50, 0x8e, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c,
	0x92, 0x63, 0x04, 0x04, 0x00, 0x00, 0xff, 0xff, 0x94, 0xa9, 0xfd, 0x9c, 0xb5, 0x00, 0x00, 0x00,
}
