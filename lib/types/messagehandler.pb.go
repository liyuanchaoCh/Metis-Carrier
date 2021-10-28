// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lib/types/messagehandler.proto

package types

import (
	fmt "fmt"
	common "github.com/RosettaFlow/Carrier-Go/lib/common"
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

type PowerMsg struct {
	//*
	//PowerId   string `json:"powerId"`
	//JobNodeId string `json:"jobNodeId"`
	//CreateAt  uint64 `json:"createAt"`
	PowerId              string   `protobuf:"bytes,1,opt,name=power_id,json=powerId,proto3" json:"power_id,omitempty"`
	JobNodeId            string   `protobuf:"bytes,2,opt,name=job_node_id,json=jobNodeId,proto3" json:"job_node_id,omitempty"`
	CreateAt             uint64   `protobuf:"varint,3,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PowerMsg) Reset()         { *m = PowerMsg{} }
func (m *PowerMsg) String() string { return proto.CompactTextString(m) }
func (*PowerMsg) ProtoMessage()    {}
func (*PowerMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_77f381d9c6d35a5b, []int{0}
}
func (m *PowerMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PowerMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PowerMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PowerMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PowerMsg.Merge(m, src)
}
func (m *PowerMsg) XXX_Size() int {
	return m.Size()
}
func (m *PowerMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_PowerMsg.DiscardUnknown(m)
}

var xxx_messageInfo_PowerMsg proto.InternalMessageInfo

func (m *PowerMsg) GetPowerId() string {
	if m != nil {
		return m.PowerId
	}
	return ""
}

func (m *PowerMsg) GetJobNodeId() string {
	if m != nil {
		return m.JobNodeId
	}
	return ""
}

func (m *PowerMsg) GetCreateAt() uint64 {
	if m != nil {
		return m.CreateAt
	}
	return 0
}

type MetadataMsg struct {
	//*
	//MetadataId      string                    `json:"metadataId"`
	//MetadataSummary *libtypes.MetadataSummary `json:"metadataSummary"`
	//ColumnMetas     []*types.MetadataColumn   `json:"columnMetas"`
	//CreateAt        uint64                    `json:"createAt"`
	MetadataId           string            `protobuf:"bytes,1,opt,name=metadata_id,json=metadataId,proto3" json:"metadata_id,omitempty"`
	MetadataSummary      *MetadataSummary  `protobuf:"bytes,2,opt,name=metadata_summary,json=metadataSummary,proto3" json:"metadata_summary,omitempty"`
	ColumnMetas          []*MetadataColumn `protobuf:"bytes,3,rep,name=column_metas,json=columnMetas,proto3" json:"column_metas,omitempty"`
	CreateAt             uint64            `protobuf:"varint,4,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MetadataMsg) Reset()         { *m = MetadataMsg{} }
func (m *MetadataMsg) String() string { return proto.CompactTextString(m) }
func (*MetadataMsg) ProtoMessage()    {}
func (*MetadataMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_77f381d9c6d35a5b, []int{1}
}
func (m *MetadataMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MetadataMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MetadataMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MetadataMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataMsg.Merge(m, src)
}
func (m *MetadataMsg) XXX_Size() int {
	return m.Size()
}
func (m *MetadataMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataMsg.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataMsg proto.InternalMessageInfo

func (m *MetadataMsg) GetMetadataId() string {
	if m != nil {
		return m.MetadataId
	}
	return ""
}

func (m *MetadataMsg) GetMetadataSummary() *MetadataSummary {
	if m != nil {
		return m.MetadataSummary
	}
	return nil
}

func (m *MetadataMsg) GetColumnMetas() []*MetadataColumn {
	if m != nil {
		return m.ColumnMetas
	}
	return nil
}

func (m *MetadataMsg) GetCreateAt() uint64 {
	if m != nil {
		return m.CreateAt
	}
	return 0
}

type MetadataAuthorityMsg struct {
	//*
	//MetadataAuthId string                   `json:"metaDataAuthId"`
	//User           string                   `json:"user"`
	//UserType       apicommonpb.UserType     `json:"userType"`
	//Auth           *types.MetadataAuthority `json:"auth"`
	//Sign           []byte                   `json:"sign"`
	//CreateAt       uint64
	MetadataAuthId       string             `protobuf:"bytes,1,opt,name=metadata_auth_id,json=metadataAuthId,proto3" json:"metadata_auth_id,omitempty"`
	User                 string             `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	UserType             common.UserType    `protobuf:"varint,3,opt,name=user_type,json=userType,proto3,enum=api.protobuf.UserType" json:"user_type,omitempty"`
	Auth                 *MetadataAuthority `protobuf:"bytes,4,opt,name=auth,proto3" json:"auth,omitempty"`
	Sign                 []byte             `protobuf:"bytes,5,opt,name=sign,proto3" json:"sign,omitempty"`
	CreateAt             uint64             `protobuf:"varint,6,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MetadataAuthorityMsg) Reset()         { *m = MetadataAuthorityMsg{} }
func (m *MetadataAuthorityMsg) String() string { return proto.CompactTextString(m) }
func (*MetadataAuthorityMsg) ProtoMessage()    {}
func (*MetadataAuthorityMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_77f381d9c6d35a5b, []int{2}
}
func (m *MetadataAuthorityMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MetadataAuthorityMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MetadataAuthorityMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MetadataAuthorityMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataAuthorityMsg.Merge(m, src)
}
func (m *MetadataAuthorityMsg) XXX_Size() int {
	return m.Size()
}
func (m *MetadataAuthorityMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataAuthorityMsg.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataAuthorityMsg proto.InternalMessageInfo

func (m *MetadataAuthorityMsg) GetMetadataAuthId() string {
	if m != nil {
		return m.MetadataAuthId
	}
	return ""
}

func (m *MetadataAuthorityMsg) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *MetadataAuthorityMsg) GetUserType() common.UserType {
	if m != nil {
		return m.UserType
	}
	return common.UserType_User_Unknown
}

func (m *MetadataAuthorityMsg) GetAuth() *MetadataAuthority {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *MetadataAuthorityMsg) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

func (m *MetadataAuthorityMsg) GetCreateAt() uint64 {
	if m != nil {
		return m.CreateAt
	}
	return 0
}

type TaskMsg struct {
	//*
	//Data          *Task
	//PowerPartyIds []string `json:"powerPartyIds"`
	Data                 *TaskPB  `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	PowerPartyIds        []string `protobuf:"bytes,2,rep,name=power_party_ids,json=powerPartyIds,proto3" json:"power_party_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskMsg) Reset()         { *m = TaskMsg{} }
func (m *TaskMsg) String() string { return proto.CompactTextString(m) }
func (*TaskMsg) ProtoMessage()    {}
func (*TaskMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_77f381d9c6d35a5b, []int{3}
}
func (m *TaskMsg) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TaskMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TaskMsg.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TaskMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskMsg.Merge(m, src)
}
func (m *TaskMsg) XXX_Size() int {
	return m.Size()
}
func (m *TaskMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskMsg.DiscardUnknown(m)
}

var xxx_messageInfo_TaskMsg proto.InternalMessageInfo

func (m *TaskMsg) GetData() *TaskPB {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *TaskMsg) GetPowerPartyIds() []string {
	if m != nil {
		return m.PowerPartyIds
	}
	return nil
}

func init() {
	proto.RegisterType((*PowerMsg)(nil), "types.PowerMsg")
	proto.RegisterType((*MetadataMsg)(nil), "types.MetadataMsg")
	proto.RegisterType((*MetadataAuthorityMsg)(nil), "types.MetadataAuthorityMsg")
	proto.RegisterType((*TaskMsg)(nil), "types.TaskMsg")
}

func init() { proto.RegisterFile("lib/types/messagehandler.proto", fileDescriptor_77f381d9c6d35a5b) }

var fileDescriptor_77f381d9c6d35a5b = []byte{
	// 479 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0xcd, 0x6e, 0xd4, 0x30,
	0x18, 0x54, 0xba, 0x69, 0xbb, 0x71, 0xfa, 0x83, 0x2c, 0x5a, 0x05, 0x90, 0x42, 0xe8, 0x01, 0x45,
	0x02, 0x12, 0x29, 0xbd, 0xc0, 0x71, 0x5b, 0x09, 0x94, 0x43, 0xd1, 0xca, 0x2c, 0x17, 0x2e, 0x91,
	0xb3, 0x36, 0xbb, 0x69, 0xd7, 0x71, 0x64, 0x3b, 0xaa, 0xf2, 0x86, 0x3d, 0xf2, 0x08, 0x68, 0x4f,
	0x3c, 0x06, 0xb2, 0x9d, 0xfd, 0xed, 0x25, 0xf9, 0x32, 0x33, 0xf6, 0xcc, 0x38, 0x06, 0xe1, 0xa2,
	0x2a, 0x53, 0xd5, 0x35, 0x54, 0xa6, 0x8c, 0x4a, 0x89, 0x67, 0x74, 0x8e, 0x6b, 0xb2, 0xa0, 0x22,
	0x69, 0x04, 0x57, 0x1c, 0x1e, 0x1a, 0xee, 0xf5, 0x85, 0x96, 0x4d, 0x39, 0x63, 0xbc, 0x4e, 0x4b,
	0x2c, 0xa9, 0x65, 0x2d, 0x6c, 0x57, 0x9b, 0x67, 0x0f, 0x07, 0xdb, 0x9b, 0x2a, 0x4c, 0xb0, 0xc2,
	0xcf, 0x19, 0x85, 0xe5, 0xc3, 0x86, 0xb9, 0x2a, 0xc1, 0x70, 0xcc, 0x1f, 0xa9, 0xb8, 0x93, 0x33,
	0xf8, 0x0a, 0x0c, 0x1b, 0x3d, 0x17, 0x15, 0x09, 0x9c, 0xc8, 0x89, 0x3d, 0x74, 0x6c, 0xbe, 0x73,
	0x02, 0x43, 0xe0, 0xdf, 0xf3, 0xb2, 0xa8, 0x39, 0xa1, 0x9a, 0x3d, 0x30, 0xac, 0x77, 0xcf, 0xcb,
	0xef, 0x9c, 0xd0, 0x9c, 0xc0, 0x37, 0xc0, 0x9b, 0x0a, 0x8a, 0x15, 0x2d, 0xb0, 0x0a, 0x06, 0x91,
	0x13, 0xbb, 0x68, 0x68, 0x81, 0x91, 0xba, 0x7a, 0x72, 0x80, 0x7f, 0xd7, 0x07, 0xd2, 0x3e, 0x6f,
	0x81, 0xbf, 0xca, 0xb7, 0xb1, 0x02, 0x2b, 0x28, 0x27, 0x70, 0x04, 0x5e, 0xac, 0x05, 0xb2, 0x65,
	0x0c, 0x8b, 0xce, 0x58, 0xfa, 0xd9, 0x65, 0x62, 0x0b, 0xaf, 0xb6, 0xfb, 0x61, 0x59, 0x74, 0xce,
	0x76, 0x01, 0xf8, 0x19, 0x9c, 0x4c, 0xf9, 0xa2, 0x65, 0x75, 0xa1, 0x19, 0x19, 0x0c, 0xa2, 0x41,
	0xec, 0x67, 0x17, 0x7b, 0xcb, 0x6f, 0x8d, 0x04, 0xf9, 0x56, 0xaa, 0x51, 0xb9, 0x5b, 0xc5, 0xdd,
	0xab, 0xf2, 0xcf, 0x01, 0x2f, 0x57, 0x8b, 0x47, 0xad, 0x9a, 0x73, 0x51, 0xa9, 0x4e, 0x77, 0x8a,
	0xb7, 0x22, 0xe3, 0x56, 0xcd, 0x37, 0xc5, 0xce, 0xd8, 0x96, 0x3e, 0x27, 0x10, 0x02, 0xb7, 0x95,
	0x54, 0xf4, 0x67, 0x68, 0x66, 0x78, 0x0d, 0x3c, 0xfd, 0x2e, 0x74, 0x3a, 0x73, 0x7c, 0x67, 0xd9,
	0x65, 0x82, 0x9b, 0xca, 0xfe, 0xa4, 0xb2, 0xfd, 0x9d, 0xfc, 0x94, 0x54, 0x4c, 0xba, 0x86, 0xa2,
	0x61, 0xdb, 0x4f, 0xf0, 0x23, 0x70, 0xb5, 0x93, 0xc9, 0xe8, 0x67, 0xc1, 0x5e, 0xb5, 0x75, 0x3a,
	0x64, 0x54, 0xda, 0x56, 0x56, 0xb3, 0x3a, 0x38, 0x8c, 0x9c, 0xf8, 0x04, 0x99, 0x79, 0xb7, 0xea,
	0xd1, 0x5e, 0xd5, 0x09, 0x38, 0x9e, 0x60, 0xf9, 0xa0, 0xcb, 0xbd, 0x03, 0xae, 0xde, 0xd2, 0x14,
	0xf2, 0xb3, 0xd3, 0xde, 0x49, 0xb3, 0xe3, 0x1b, 0x64, 0x28, 0xf8, 0x1e, 0x9c, 0xdb, 0xbb, 0xd3,
	0x60, 0xa1, 0xba, 0xa2, 0x22, 0x32, 0x38, 0x88, 0x06, 0xb1, 0x87, 0x4e, 0x0d, 0x3c, 0xd6, 0x68,
	0x4e, 0xe4, 0xcd, 0x97, 0xa7, 0x65, 0xe8, 0xfc, 0x59, 0x86, 0xce, 0xdf, 0x65, 0xe8, 0xfc, 0xfa,
	0x30, 0xab, 0xd4, 0xbc, 0x2d, 0x93, 0x29, 0x67, 0x29, 0xe2, 0x92, 0x2a, 0x85, 0xbf, 0x2e, 0xf8,
	0x63, 0x7a, 0x8b, 0x85, 0xa8, 0xa8, 0xf8, 0xf4, 0x8d, 0xa7, 0xeb, 0x9b, 0x5b, 0x1e, 0x99, 0xc3,
	0xb8, 0xfe, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x73, 0x9b, 0x33, 0x29, 0x3c, 0x03, 0x00, 0x00,
}

func (m *PowerMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PowerMsg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PowerMsg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.CreateAt != 0 {
		i = encodeVarintMessagehandler(dAtA, i, uint64(m.CreateAt))
		i--
		dAtA[i] = 0x18
	}
	if len(m.JobNodeId) > 0 {
		i -= len(m.JobNodeId)
		copy(dAtA[i:], m.JobNodeId)
		i = encodeVarintMessagehandler(dAtA, i, uint64(len(m.JobNodeId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PowerId) > 0 {
		i -= len(m.PowerId)
		copy(dAtA[i:], m.PowerId)
		i = encodeVarintMessagehandler(dAtA, i, uint64(len(m.PowerId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MetadataMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MetadataMsg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MetadataMsg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.CreateAt != 0 {
		i = encodeVarintMessagehandler(dAtA, i, uint64(m.CreateAt))
		i--
		dAtA[i] = 0x20
	}
	if len(m.ColumnMetas) > 0 {
		for iNdEx := len(m.ColumnMetas) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ColumnMetas[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMessagehandler(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.MetadataSummary != nil {
		{
			size, err := m.MetadataSummary.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessagehandler(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.MetadataId) > 0 {
		i -= len(m.MetadataId)
		copy(dAtA[i:], m.MetadataId)
		i = encodeVarintMessagehandler(dAtA, i, uint64(len(m.MetadataId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MetadataAuthorityMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MetadataAuthorityMsg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MetadataAuthorityMsg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.CreateAt != 0 {
		i = encodeVarintMessagehandler(dAtA, i, uint64(m.CreateAt))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Sign) > 0 {
		i -= len(m.Sign)
		copy(dAtA[i:], m.Sign)
		i = encodeVarintMessagehandler(dAtA, i, uint64(len(m.Sign)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Auth != nil {
		{
			size, err := m.Auth.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessagehandler(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.UserType != 0 {
		i = encodeVarintMessagehandler(dAtA, i, uint64(m.UserType))
		i--
		dAtA[i] = 0x18
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintMessagehandler(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.MetadataAuthId) > 0 {
		i -= len(m.MetadataAuthId)
		copy(dAtA[i:], m.MetadataAuthId)
		i = encodeVarintMessagehandler(dAtA, i, uint64(len(m.MetadataAuthId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TaskMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TaskMsg) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TaskMsg) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.PowerPartyIds) > 0 {
		for iNdEx := len(m.PowerPartyIds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PowerPartyIds[iNdEx])
			copy(dAtA[i:], m.PowerPartyIds[iNdEx])
			i = encodeVarintMessagehandler(dAtA, i, uint64(len(m.PowerPartyIds[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessagehandler(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessagehandler(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessagehandler(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PowerMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PowerId)
	if l > 0 {
		n += 1 + l + sovMessagehandler(uint64(l))
	}
	l = len(m.JobNodeId)
	if l > 0 {
		n += 1 + l + sovMessagehandler(uint64(l))
	}
	if m.CreateAt != 0 {
		n += 1 + sovMessagehandler(uint64(m.CreateAt))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MetadataMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MetadataId)
	if l > 0 {
		n += 1 + l + sovMessagehandler(uint64(l))
	}
	if m.MetadataSummary != nil {
		l = m.MetadataSummary.Size()
		n += 1 + l + sovMessagehandler(uint64(l))
	}
	if len(m.ColumnMetas) > 0 {
		for _, e := range m.ColumnMetas {
			l = e.Size()
			n += 1 + l + sovMessagehandler(uint64(l))
		}
	}
	if m.CreateAt != 0 {
		n += 1 + sovMessagehandler(uint64(m.CreateAt))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MetadataAuthorityMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MetadataAuthId)
	if l > 0 {
		n += 1 + l + sovMessagehandler(uint64(l))
	}
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovMessagehandler(uint64(l))
	}
	if m.UserType != 0 {
		n += 1 + sovMessagehandler(uint64(m.UserType))
	}
	if m.Auth != nil {
		l = m.Auth.Size()
		n += 1 + l + sovMessagehandler(uint64(l))
	}
	l = len(m.Sign)
	if l > 0 {
		n += 1 + l + sovMessagehandler(uint64(l))
	}
	if m.CreateAt != 0 {
		n += 1 + sovMessagehandler(uint64(m.CreateAt))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TaskMsg) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovMessagehandler(uint64(l))
	}
	if len(m.PowerPartyIds) > 0 {
		for _, s := range m.PowerPartyIds {
			l = len(s)
			n += 1 + l + sovMessagehandler(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMessagehandler(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessagehandler(x uint64) (n int) {
	return sovMessagehandler(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PowerMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessagehandler
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
			return fmt.Errorf("proto: PowerMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PowerMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PowerId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessagehandler
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
				return ErrInvalidLengthMessagehandler
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessagehandler
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PowerId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobNodeId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessagehandler
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
				return ErrInvalidLengthMessagehandler
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessagehandler
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JobNodeId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateAt", wireType)
			}
			m.CreateAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessagehandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreateAt |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessagehandler(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessagehandler
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
func (m *MetadataMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessagehandler
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
			return fmt.Errorf("proto: MetadataMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetadataMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetadataId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessagehandler
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
				return ErrInvalidLengthMessagehandler
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessagehandler
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MetadataId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetadataSummary", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessagehandler
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
				return ErrInvalidLengthMessagehandler
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessagehandler
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MetadataSummary == nil {
				m.MetadataSummary = &MetadataSummary{}
			}
			if err := m.MetadataSummary.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ColumnMetas", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessagehandler
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
				return ErrInvalidLengthMessagehandler
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessagehandler
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ColumnMetas = append(m.ColumnMetas, &MetadataColumn{})
			if err := m.ColumnMetas[len(m.ColumnMetas)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateAt", wireType)
			}
			m.CreateAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessagehandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreateAt |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessagehandler(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessagehandler
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
func (m *MetadataAuthorityMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessagehandler
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
			return fmt.Errorf("proto: MetadataAuthorityMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetadataAuthorityMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetadataAuthId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessagehandler
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
				return ErrInvalidLengthMessagehandler
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessagehandler
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MetadataAuthId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessagehandler
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
				return ErrInvalidLengthMessagehandler
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessagehandler
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserType", wireType)
			}
			m.UserType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessagehandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserType |= common.UserType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Auth", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessagehandler
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
				return ErrInvalidLengthMessagehandler
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessagehandler
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Auth == nil {
				m.Auth = &MetadataAuthority{}
			}
			if err := m.Auth.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sign", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessagehandler
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
				return ErrInvalidLengthMessagehandler
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessagehandler
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sign = append(m.Sign[:0], dAtA[iNdEx:postIndex]...)
			if m.Sign == nil {
				m.Sign = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateAt", wireType)
			}
			m.CreateAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessagehandler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreateAt |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessagehandler(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessagehandler
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
func (m *TaskMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessagehandler
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
			return fmt.Errorf("proto: TaskMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TaskMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessagehandler
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
				return ErrInvalidLengthMessagehandler
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessagehandler
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &TaskPB{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PowerPartyIds", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessagehandler
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
				return ErrInvalidLengthMessagehandler
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessagehandler
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PowerPartyIds = append(m.PowerPartyIds, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessagehandler(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessagehandler
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
func skipMessagehandler(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessagehandler
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
					return 0, ErrIntOverflowMessagehandler
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
					return 0, ErrIntOverflowMessagehandler
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
				return 0, ErrInvalidLengthMessagehandler
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessagehandler
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessagehandler
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessagehandler        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessagehandler          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessagehandler = fmt.Errorf("proto: unexpected end of group")
)
