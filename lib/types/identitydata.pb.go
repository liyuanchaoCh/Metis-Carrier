// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lib/types/identitydata.proto

package types

import (
	fmt "fmt"
	common "github.com/RosettaFlow/Carrier-Go/lib/common"
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

// IdentityData represents the stored data structure.
type IdentityData struct {
	IdentityId string `protobuf:"bytes,1,opt,name=identity_id,json=identityId,proto3" json:"identity_id,omitempty"`
	NodeId     string `protobuf:"bytes,2,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	NodeName   string `protobuf:"bytes,3,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	DataId     string `protobuf:"bytes,4,opt,name=data_id,json=dataId,proto3" json:"data_id,omitempty"`
	// N means normal, D means deleted
	DataStatus common.DataStatus `protobuf:"varint,5,opt,name=data_status,json=dataStatus,proto3,enum=api.protobuf.DataStatus" json:"data_status,omitempty"`
	// Y : normal, N non-normal
	Status common.CommonStatus `protobuf:"varint,6,opt,name=status,proto3,enum=api.protobuf.CommonStatus" json:"status,omitempty"`
	// json format for credential
	Credential           string   `protobuf:"bytes,7,opt,name=credential,proto3" json:"credential,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IdentityData) Reset()         { *m = IdentityData{} }
func (m *IdentityData) String() string { return proto.CompactTextString(m) }
func (*IdentityData) ProtoMessage()    {}
func (*IdentityData) Descriptor() ([]byte, []int) {
	return fileDescriptor_fdff2e4768e45a01, []int{0}
}
func (m *IdentityData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IdentityData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IdentityData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IdentityData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IdentityData.Merge(m, src)
}
func (m *IdentityData) XXX_Size() int {
	return m.Size()
}
func (m *IdentityData) XXX_DiscardUnknown() {
	xxx_messageInfo_IdentityData.DiscardUnknown(m)
}

var xxx_messageInfo_IdentityData proto.InternalMessageInfo

func (m *IdentityData) GetIdentityId() string {
	if m != nil {
		return m.IdentityId
	}
	return ""
}

func (m *IdentityData) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *IdentityData) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *IdentityData) GetDataId() string {
	if m != nil {
		return m.DataId
	}
	return ""
}

func (m *IdentityData) GetDataStatus() common.DataStatus {
	if m != nil {
		return m.DataStatus
	}
	return common.DataStatus_DataStatus_Unknown
}

func (m *IdentityData) GetStatus() common.CommonStatus {
	if m != nil {
		return m.Status
	}
	return common.CommonStatus_CommonStatus_Unknown
}

func (m *IdentityData) GetCredential() string {
	if m != nil {
		return m.Credential
	}
	return ""
}

func init() {
	proto.RegisterType((*IdentityData)(nil), "types.IdentityData")
}

func init() { proto.RegisterFile("lib/types/identitydata.proto", fileDescriptor_fdff2e4768e45a01) }

var fileDescriptor_fdff2e4768e45a01 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xc1, 0x4a, 0xf3, 0x40,
	0x10, 0xc7, 0xd9, 0x7e, 0x5f, 0x53, 0xbb, 0x15, 0x0f, 0x0b, 0x62, 0xa8, 0x12, 0x8b, 0x5e, 0x0a,
	0x62, 0x16, 0xea, 0xa9, 0x57, 0x2b, 0x4a, 0x2e, 0x1e, 0xea, 0xcd, 0x8b, 0x4c, 0xba, 0x6b, 0x5d,
	0x68, 0x32, 0x61, 0x77, 0x8b, 0xf4, 0x21, 0x7c, 0x2f, 0x8f, 0x3e, 0x82, 0xe4, 0x49, 0x64, 0x27,
	0x49, 0xc5, 0xdb, 0xfe, 0xe7, 0x37, 0xbf, 0xec, 0x4c, 0x96, 0x9f, 0x6d, 0x4c, 0x2e, 0xfd, 0xae,
	0xd2, 0x4e, 0x1a, 0xa5, 0x4b, 0x6f, 0xfc, 0x4e, 0x81, 0x87, 0xb4, 0xb2, 0xe8, 0x51, 0xf4, 0x89,
	0x8c, 0x8f, 0x43, 0xd3, 0x0a, 0x8b, 0x02, 0x4b, 0x99, 0x83, 0xd3, 0x0d, 0x1d, 0x5f, 0x5a, 0x5d,
	0xa1, 0x93, 0x14, 0xf2, 0xed, 0xab, 0x5c, 0xe3, 0x1a, 0x29, 0xd0, 0xa9, 0x69, 0xba, 0xf8, 0xe8,
	0xf1, 0xc3, 0xac, 0xfd, 0xf2, 0x1d, 0x78, 0x10, 0xe7, 0x7c, 0xd4, 0xdd, 0xf4, 0x62, 0x54, 0xcc,
	0x26, 0x6c, 0x3a, 0x5c, 0xf2, 0xae, 0x94, 0x29, 0x71, 0xc2, 0x07, 0x25, 0x2a, 0x1d, 0x60, 0x8f,
	0x60, 0x14, 0x62, 0xa6, 0xc4, 0x29, 0x1f, 0x12, 0x28, 0xa1, 0xd0, 0xf1, 0x3f, 0x42, 0x07, 0xa1,
	0xf0, 0x08, 0x85, 0x0e, 0x56, 0x18, 0x3c, 0x58, 0xff, 0x1b, 0x2b, 0xc4, 0x4c, 0x89, 0x39, 0x1f,
	0x11, 0x70, 0x1e, 0xfc, 0xd6, 0xc5, 0xfd, 0x09, 0x9b, 0x1e, 0xcd, 0xe2, 0x14, 0x2a, 0x93, 0x76,
	0x93, 0xa7, 0x61, 0xb0, 0x27, 0xe2, 0x4b, 0xae, 0xf6, 0x67, 0x31, 0xe3, 0x51, 0x6b, 0x45, 0x64,
	0x8d, 0xff, 0x5a, 0x0b, 0xfa, 0x23, 0xad, 0xd7, 0x76, 0x8a, 0x84, 0xf3, 0x95, 0xd5, 0xb4, 0x0d,
	0x6c, 0xe2, 0x41, 0xb3, 0xdd, 0x6f, 0xe5, 0x76, 0xfe, 0x59, 0x27, 0xec, 0xab, 0x4e, 0xd8, 0x77,
	0x9d, 0xb0, 0xe7, 0xab, 0xb5, 0xf1, 0x6f, 0xdb, 0x3c, 0x5d, 0x61, 0x21, 0x97, 0xe8, 0xb4, 0xf7,
	0x70, 0xbf, 0xc1, 0x77, 0xb9, 0x00, 0x6b, 0x8d, 0xb6, 0xd7, 0x0f, 0x28, 0xf7, 0x0f, 0x94, 0x47,
	0x74, 0xf3, 0xcd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x53, 0x6d, 0xc3, 0x41, 0xb4, 0x01, 0x00,
	0x00,
}

func (m *IdentityData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IdentityData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IdentityData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Credential) > 0 {
		i -= len(m.Credential)
		copy(dAtA[i:], m.Credential)
		i = encodeVarintIdentitydata(dAtA, i, uint64(len(m.Credential)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Status != 0 {
		i = encodeVarintIdentitydata(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x30
	}
	if m.DataStatus != 0 {
		i = encodeVarintIdentitydata(dAtA, i, uint64(m.DataStatus))
		i--
		dAtA[i] = 0x28
	}
	if len(m.DataId) > 0 {
		i -= len(m.DataId)
		copy(dAtA[i:], m.DataId)
		i = encodeVarintIdentitydata(dAtA, i, uint64(len(m.DataId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.NodeName) > 0 {
		i -= len(m.NodeName)
		copy(dAtA[i:], m.NodeName)
		i = encodeVarintIdentitydata(dAtA, i, uint64(len(m.NodeName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.NodeId) > 0 {
		i -= len(m.NodeId)
		copy(dAtA[i:], m.NodeId)
		i = encodeVarintIdentitydata(dAtA, i, uint64(len(m.NodeId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.IdentityId) > 0 {
		i -= len(m.IdentityId)
		copy(dAtA[i:], m.IdentityId)
		i = encodeVarintIdentitydata(dAtA, i, uint64(len(m.IdentityId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIdentitydata(dAtA []byte, offset int, v uint64) int {
	offset -= sovIdentitydata(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IdentityData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.IdentityId)
	if l > 0 {
		n += 1 + l + sovIdentitydata(uint64(l))
	}
	l = len(m.NodeId)
	if l > 0 {
		n += 1 + l + sovIdentitydata(uint64(l))
	}
	l = len(m.NodeName)
	if l > 0 {
		n += 1 + l + sovIdentitydata(uint64(l))
	}
	l = len(m.DataId)
	if l > 0 {
		n += 1 + l + sovIdentitydata(uint64(l))
	}
	if m.DataStatus != 0 {
		n += 1 + sovIdentitydata(uint64(m.DataStatus))
	}
	if m.Status != 0 {
		n += 1 + sovIdentitydata(uint64(m.Status))
	}
	l = len(m.Credential)
	if l > 0 {
		n += 1 + l + sovIdentitydata(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovIdentitydata(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIdentitydata(x uint64) (n int) {
	return sovIdentitydata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IdentityData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIdentitydata
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
			return fmt.Errorf("proto: IdentityData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IdentityData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdentityId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentitydata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIdentitydata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentitydata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IdentityId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentitydata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIdentitydata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentitydata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentitydata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIdentitydata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentitydata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentitydata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIdentitydata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentitydata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataStatus", wireType)
			}
			m.DataStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentitydata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DataStatus |= common.DataStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentitydata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= common.CommonStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Credential", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentitydata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIdentitydata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentitydata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Credential = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIdentitydata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIdentitydata
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
func skipIdentitydata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIdentitydata
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
					return 0, ErrIntOverflowIdentitydata
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
					return 0, ErrIntOverflowIdentitydata
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
				return 0, ErrInvalidLengthIdentitydata
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIdentitydata
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIdentitydata
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIdentitydata        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIdentitydata          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIdentitydata = fmt.Errorf("proto: unexpected end of group")
)
