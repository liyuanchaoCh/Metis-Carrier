// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: carrier/rpc/debug/v1/messages.proto

package v1

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// test struct for gossip
type GossipTestData struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty" ssz-size:"?,32" ssz-max:"16777216"`
	Count                uint64   `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Step                 uint64   `protobuf:"varint,3,opt,name=step,proto3" json:"step,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GossipTestData) Reset()         { *m = GossipTestData{} }
func (m *GossipTestData) String() string { return proto.CompactTextString(m) }
func (*GossipTestData) ProtoMessage()    {}
func (*GossipTestData) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a20826b6fdab274, []int{0}
}
func (m *GossipTestData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GossipTestData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GossipTestData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GossipTestData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GossipTestData.Merge(m, src)
}
func (m *GossipTestData) XXX_Size() int {
	return m.Size()
}
func (m *GossipTestData) XXX_DiscardUnknown() {
	xxx_messageInfo_GossipTestData.DiscardUnknown(m)
}

var xxx_messageInfo_GossipTestData proto.InternalMessageInfo

func (m *GossipTestData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GossipTestData) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *GossipTestData) GetStep() uint64 {
	if m != nil {
		return m.Step
	}
	return 0
}

// test struct for signed gossip
type SignedGossipTestData struct {
	Data                 *GossipTestData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Signature            []byte          `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SignedGossipTestData) Reset()         { *m = SignedGossipTestData{} }
func (m *SignedGossipTestData) String() string { return proto.CompactTextString(m) }
func (*SignedGossipTestData) ProtoMessage()    {}
func (*SignedGossipTestData) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a20826b6fdab274, []int{1}
}
func (m *SignedGossipTestData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignedGossipTestData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignedGossipTestData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignedGossipTestData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedGossipTestData.Merge(m, src)
}
func (m *SignedGossipTestData) XXX_Size() int {
	return m.Size()
}
func (m *SignedGossipTestData) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedGossipTestData.DiscardUnknown(m)
}

var xxx_messageInfo_SignedGossipTestData proto.InternalMessageInfo

func (m *SignedGossipTestData) GetData() *GossipTestData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SignedGossipTestData) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*GossipTestData)(nil), "carrier.rpc.debug.v1.GossipTestData")
	proto.RegisterType((*SignedGossipTestData)(nil), "carrier.rpc.debug.v1.SignedGossipTestData")
}

func init() {
	proto.RegisterFile("carrier/rpc/debug/v1/messages.proto", fileDescriptor_2a20826b6fdab274)
}

var fileDescriptor_2a20826b6fdab274 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x4a, 0x03, 0x31,
	0x18, 0xc4, 0x59, 0xad, 0x82, 0xb1, 0x78, 0x58, 0x7a, 0x28, 0x22, 0xb5, 0xa4, 0x22, 0x3d, 0xd8,
	0x84, 0x6d, 0xa1, 0x95, 0x5e, 0x84, 0x2a, 0x78, 0xb6, 0x7a, 0xf2, 0x96, 0xcd, 0x7e, 0xa6, 0x41,
	0x76, 0x13, 0xf2, 0x25, 0x55, 0xfa, 0x84, 0x1e, 0x7d, 0x02, 0x91, 0x3e, 0x82, 0x4f, 0x20, 0xa6,
	0x15, 0xff, 0xd0, 0xdb, 0xcc, 0xc7, 0x0c, 0x3f, 0x32, 0x21, 0x1d, 0x29, 0x9c, 0xd3, 0xe0, 0xb8,
	0xb3, 0x92, 0x17, 0x90, 0x07, 0xc5, 0xe7, 0x19, 0x2f, 0x01, 0x51, 0x28, 0x40, 0x66, 0x9d, 0xf1,
	0x26, 0x6d, 0xac, 0x43, 0xcc, 0x59, 0xc9, 0x62, 0x88, 0xcd, 0xb3, 0xc3, 0x8e, 0x03, 0x6b, 0x90,
	0xc7, 0x48, 0x1e, 0x1e, 0xb8, 0x32, 0xca, 0x44, 0x13, 0xd5, 0xaa, 0x4a, 0xe7, 0xe4, 0xe0, 0xda,
	0x20, 0x6a, 0x7b, 0x07, 0xe8, 0xaf, 0x84, 0x17, 0xe9, 0x98, 0xd4, 0x0a, 0xe1, 0x45, 0x33, 0x69,
	0x27, 0xdd, 0xfa, 0xe4, 0xf4, 0xe3, 0xed, 0x98, 0x22, 0x2e, 0x7a, 0xa8, 0x17, 0x30, 0xa6, 0x17,
	0x67, 0x83, 0x3e, 0x6d, 0x7f, 0xf9, 0x52, 0x3c, 0x8f, 0x69, 0x36, 0x1c, 0x8d, 0x46, 0xfd, 0x6c,
	0x48, 0xa7, 0xb1, 0x93, 0x36, 0xc8, 0x8e, 0x34, 0xa1, 0xf2, 0xcd, 0xad, 0x76, 0xd2, 0xad, 0x4d,
	0x57, 0x26, 0x4d, 0x49, 0x0d, 0x3d, 0xd8, 0xe6, 0x76, 0x3c, 0x46, 0x4d, 0x2b, 0xd2, 0xb8, 0xd5,
	0xaa, 0x82, 0xe2, 0x1f, 0xfd, 0xfc, 0x17, 0x7d, 0xbf, 0x7f, 0xc2, 0x36, 0xbd, 0x8c, 0xfd, 0xed,
	0xac, 0xd9, 0x47, 0x64, 0x0f, 0xb5, 0xaa, 0x84, 0x0f, 0x0e, 0x22, 0xbf, 0x3e, 0xfd, 0x39, 0x4c,
	0x6e, 0x5e, 0x96, 0xad, 0xe4, 0x75, 0xd9, 0x4a, 0xde, 0x97, 0xad, 0xe4, 0xfe, 0x52, 0x69, 0x3f,
	0x0b, 0x39, 0x93, 0xa6, 0xe4, 0x85, 0xf0, 0xa1, 0xf4, 0x20, 0x67, 0xb8, 0x92, 0xbd, 0x0a, 0xfc,
	0x93, 0x71, 0x8f, 0xbd, 0xef, 0xe5, 0x6d, 0xce, 0x37, 0x7d, 0x42, 0xbe, 0x1b, 0x17, 0x1c, 0x7c,
	0x06, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x61, 0x7a, 0x0d, 0xa3, 0x01, 0x00, 0x00,
}

func (m *GossipTestData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GossipTestData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GossipTestData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Step != 0 {
		i = encodeVarintMessages(dAtA, i, uint64(m.Step))
		i--
		dAtA[i] = 0x18
	}
	if m.Count != 0 {
		i = encodeVarintMessages(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SignedGossipTestData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignedGossipTestData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignedGossipTestData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x12
	}
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessages(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessages(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessages(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GossipTestData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	if m.Count != 0 {
		n += 1 + sovMessages(uint64(m.Count))
	}
	if m.Step != 0 {
		n += 1 + sovMessages(uint64(m.Step))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SignedGossipTestData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMessages(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessages(x uint64) (n int) {
	return sovMessages(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GossipTestData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GossipTestData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GossipTestData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Step", wireType)
			}
			m.Step = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Step |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessages
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
func (m *SignedGossipTestData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SignedGossipTestData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignedGossipTestData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &GossipTestData{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessages
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
func skipMessages(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessages
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
					return 0, ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessages
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
			if length < 0 {
				return 0, ErrInvalidLengthMessages
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessages
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessages
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessages        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessages          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessages = fmt.Errorf("proto: unexpected end of group")
)