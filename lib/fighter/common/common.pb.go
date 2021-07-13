// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lib/fighter/common/common.proto

package common

import (
	fmt "fmt"
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

type TaskReadyGoReq struct {
	TaskId               string                 `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	ContractId           string                 `protobuf:"bytes,2,opt,name=contract_id,json=contractId,proto3" json:"contract_id,omitempty"`
	DataId               string                 `protobuf:"bytes,3,opt,name=data_id,json=dataId,proto3" json:"data_id,omitempty"`
	PartyId              string                 `protobuf:"bytes,4,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	EnvId                string                 `protobuf:"bytes,5,opt,name=env_id,json=envId,proto3" json:"env_id,omitempty"`
	Peers                []*TaskReadyGoReq_Peer `protobuf:"bytes,6,rep,name=peers,proto3" json:"peers,omitempty"`
	ContractCfg          string                 `protobuf:"bytes,7,opt,name=contract_cfg,json=contractCfg,proto3" json:"contract_cfg,omitempty"`
	DataParty            []string               `protobuf:"bytes,8,rep,name=data_party,json=dataParty,proto3" json:"data_party,omitempty"`
	ComputationParty     []string               `protobuf:"bytes,9,rep,name=computation_party,json=computationParty,proto3" json:"computation_party,omitempty"`
	ResultParty          []string               `protobuf:"bytes,10,rep,name=result_party,json=resultParty,proto3" json:"result_party,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *TaskReadyGoReq) Reset()         { *m = TaskReadyGoReq{} }
func (m *TaskReadyGoReq) String() string { return proto.CompactTextString(m) }
func (*TaskReadyGoReq) ProtoMessage()    {}
func (*TaskReadyGoReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fb4ac3629666f03, []int{0}
}
func (m *TaskReadyGoReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TaskReadyGoReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TaskReadyGoReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TaskReadyGoReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskReadyGoReq.Merge(m, src)
}
func (m *TaskReadyGoReq) XXX_Size() int {
	return m.Size()
}
func (m *TaskReadyGoReq) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskReadyGoReq.DiscardUnknown(m)
}

var xxx_messageInfo_TaskReadyGoReq proto.InternalMessageInfo

func (m *TaskReadyGoReq) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *TaskReadyGoReq) GetContractId() string {
	if m != nil {
		return m.ContractId
	}
	return ""
}

func (m *TaskReadyGoReq) GetDataId() string {
	if m != nil {
		return m.DataId
	}
	return ""
}

func (m *TaskReadyGoReq) GetPartyId() string {
	if m != nil {
		return m.PartyId
	}
	return ""
}

func (m *TaskReadyGoReq) GetEnvId() string {
	if m != nil {
		return m.EnvId
	}
	return ""
}

func (m *TaskReadyGoReq) GetPeers() []*TaskReadyGoReq_Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

func (m *TaskReadyGoReq) GetContractCfg() string {
	if m != nil {
		return m.ContractCfg
	}
	return ""
}

func (m *TaskReadyGoReq) GetDataParty() []string {
	if m != nil {
		return m.DataParty
	}
	return nil
}

func (m *TaskReadyGoReq) GetComputationParty() []string {
	if m != nil {
		return m.ComputationParty
	}
	return nil
}

func (m *TaskReadyGoReq) GetResultParty() []string {
	if m != nil {
		return m.ResultParty
	}
	return nil
}

type TaskReadyGoReq_Peer struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Party                string   `protobuf:"bytes,3,opt,name=party,proto3" json:"party,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskReadyGoReq_Peer) Reset()         { *m = TaskReadyGoReq_Peer{} }
func (m *TaskReadyGoReq_Peer) String() string { return proto.CompactTextString(m) }
func (*TaskReadyGoReq_Peer) ProtoMessage()    {}
func (*TaskReadyGoReq_Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fb4ac3629666f03, []int{0, 0}
}
func (m *TaskReadyGoReq_Peer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TaskReadyGoReq_Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TaskReadyGoReq_Peer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TaskReadyGoReq_Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskReadyGoReq_Peer.Merge(m, src)
}
func (m *TaskReadyGoReq_Peer) XXX_Size() int {
	return m.Size()
}
func (m *TaskReadyGoReq_Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskReadyGoReq_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_TaskReadyGoReq_Peer proto.InternalMessageInfo

func (m *TaskReadyGoReq_Peer) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *TaskReadyGoReq_Peer) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *TaskReadyGoReq_Peer) GetParty() string {
	if m != nil {
		return m.Party
	}
	return ""
}

func (m *TaskReadyGoReq_Peer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type TaskReadyGoReply struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskReadyGoReply) Reset()         { *m = TaskReadyGoReply{} }
func (m *TaskReadyGoReply) String() string { return proto.CompactTextString(m) }
func (*TaskReadyGoReply) ProtoMessage()    {}
func (*TaskReadyGoReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fb4ac3629666f03, []int{1}
}
func (m *TaskReadyGoReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TaskReadyGoReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TaskReadyGoReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TaskReadyGoReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskReadyGoReply.Merge(m, src)
}
func (m *TaskReadyGoReply) XXX_Size() int {
	return m.Size()
}
func (m *TaskReadyGoReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskReadyGoReply.DiscardUnknown(m)
}

var xxx_messageInfo_TaskReadyGoReply proto.InternalMessageInfo

func (m *TaskReadyGoReply) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *TaskReadyGoReply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*TaskReadyGoReq)(nil), "common.TaskReadyGoReq")
	proto.RegisterType((*TaskReadyGoReq_Peer)(nil), "common.TaskReadyGoReq.Peer")
	proto.RegisterType((*TaskReadyGoReply)(nil), "common.TaskReadyGoReply")
}

func init() { proto.RegisterFile("lib/fighter/common/common.proto", fileDescriptor_3fb4ac3629666f03) }

var fileDescriptor_3fb4ac3629666f03 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0xdd, 0xce, 0xd2, 0x30,
	0x18, 0xce, 0x36, 0x36, 0xd8, 0xcb, 0x97, 0x2f, 0xd8, 0x68, 0x9c, 0x1a, 0xf9, 0x80, 0x23, 0x12,
	0x23, 0x8b, 0xca, 0x15, 0x40, 0x22, 0xd9, 0x19, 0x69, 0x38, 0xf2, 0xc4, 0x94, 0xb5, 0x8c, 0x66,
	0x3f, 0xad, 0x5d, 0xc1, 0x70, 0x57, 0x5e, 0x86, 0x87, 0x5e, 0x82, 0xe1, 0x4a, 0x4c, 0xdb, 0x49,
	0x24, 0x1e, 0xed, 0x7d, 0x7e, 0xda, 0xf7, 0xc9, 0xb3, 0xc2, 0x53, 0xc5, 0xf7, 0xe9, 0x81, 0x17,
	0x47, 0xcd, 0x54, 0x9a, 0x8b, 0xba, 0x16, 0x4d, 0xf7, 0x59, 0x48, 0x25, 0xb4, 0x40, 0x91, 0x43,
	0xb3, 0x1f, 0x01, 0x3c, 0xee, 0x48, 0x5b, 0x62, 0x46, 0xe8, 0x65, 0x23, 0x30, 0xfb, 0x86, 0x5e,
	0x42, 0x5f, 0x93, 0xb6, 0xfc, 0xca, 0x69, 0xe2, 0x4d, 0xbc, 0x79, 0x8c, 0x23, 0x03, 0x33, 0x8a,
	0x9e, 0x60, 0x98, 0x8b, 0x46, 0x2b, 0x92, 0x6b, 0x23, 0xfa, 0x56, 0x84, 0xbf, 0x54, 0x46, 0xcd,
	0x49, 0x4a, 0x34, 0x31, 0x62, 0xe0, 0x4e, 0x1a, 0x98, 0x51, 0xf4, 0x0a, 0x06, 0x92, 0x28, 0x7d,
	0x31, 0x4a, 0xcf, 0x2a, 0x7d, 0x8b, 0x33, 0x8a, 0x5e, 0x40, 0xc4, 0x9a, 0xb3, 0x11, 0x42, 0x2b,
	0x84, 0xac, 0x39, 0x67, 0x14, 0x7d, 0x80, 0x50, 0x32, 0xa6, 0xda, 0x24, 0x9a, 0x04, 0xf3, 0xe1,
	0xc7, 0x37, 0x8b, 0x2e, 0xfd, 0x7d, 0xd6, 0xc5, 0x96, 0x31, 0x85, 0x9d, 0x13, 0x4d, 0xe1, 0xe1,
	0x16, 0x2f, 0x3f, 0x14, 0x49, 0xdf, 0xde, 0x77, 0x8b, 0xbc, 0x3e, 0x14, 0xe8, 0x2d, 0x80, 0x0d,
	0x68, 0x97, 0x27, 0x83, 0x49, 0x30, 0x8f, 0x71, 0x6c, 0x98, 0xad, 0x21, 0xd0, 0x3b, 0x78, 0x96,
	0x8b, 0x5a, 0x9e, 0x34, 0xd1, 0x5c, 0x34, 0x9d, 0x2b, 0xb6, 0xae, 0xd1, 0x3f, 0x82, 0x33, 0x4f,
	0xe1, 0x41, 0xb1, 0xf6, 0x54, 0xe9, 0xce, 0x07, 0xd6, 0x37, 0x74, 0x9c, 0xb5, 0xbc, 0xde, 0x41,
	0xcf, 0x04, 0x44, 0x8f, 0xe0, 0x73, 0xd9, 0x95, 0xe9, 0x73, 0x89, 0x10, 0xf4, 0xa4, 0x50, 0xda,
	0x36, 0x18, 0x62, 0x3b, 0xa3, 0xe7, 0x10, 0xba, 0x7b, 0x5c, 0x73, 0x0e, 0x18, 0x67, 0x43, 0x6a,
	0xd6, 0x95, 0x66, 0xe7, 0xd9, 0x12, 0x46, 0x77, 0x2d, 0xc8, 0xea, 0x62, 0x36, 0x88, 0xd2, 0x6e,
	0x18, 0x60, 0x5f, 0x94, 0x68, 0x04, 0x41, 0xdd, 0x16, 0xdd, 0x2f, 0x32, 0xe3, 0x6a, 0xf5, 0xf3,
	0x3a, 0xf6, 0x7e, 0x5d, 0xc7, 0xde, 0xef, 0xeb, 0xd8, 0xfb, 0xb2, 0x2c, 0xb8, 0x3e, 0x9e, 0xf6,
	0xa6, 0xd5, 0x14, 0x8b, 0x96, 0x69, 0x4d, 0x3e, 0x57, 0xe2, 0x7b, 0xba, 0x26, 0x4a, 0x71, 0xa6,
	0xde, 0x6f, 0x44, 0xfa, 0xff, 0x0b, 0xda, 0x47, 0xf6, 0xed, 0x7c, 0xfa, 0x13, 0x00, 0x00, 0xff,
	0xff, 0x7a, 0xb6, 0x5c, 0xf9, 0x5e, 0x02, 0x00, 0x00,
}

func (m *TaskReadyGoReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TaskReadyGoReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TaskReadyGoReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ResultParty) > 0 {
		for iNdEx := len(m.ResultParty) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ResultParty[iNdEx])
			copy(dAtA[i:], m.ResultParty[iNdEx])
			i = encodeVarintCommon(dAtA, i, uint64(len(m.ResultParty[iNdEx])))
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.ComputationParty) > 0 {
		for iNdEx := len(m.ComputationParty) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ComputationParty[iNdEx])
			copy(dAtA[i:], m.ComputationParty[iNdEx])
			i = encodeVarintCommon(dAtA, i, uint64(len(m.ComputationParty[iNdEx])))
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.DataParty) > 0 {
		for iNdEx := len(m.DataParty) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.DataParty[iNdEx])
			copy(dAtA[i:], m.DataParty[iNdEx])
			i = encodeVarintCommon(dAtA, i, uint64(len(m.DataParty[iNdEx])))
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.ContractCfg) > 0 {
		i -= len(m.ContractCfg)
		copy(dAtA[i:], m.ContractCfg)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.ContractCfg)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Peers) > 0 {
		for iNdEx := len(m.Peers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Peers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCommon(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.EnvId) > 0 {
		i -= len(m.EnvId)
		copy(dAtA[i:], m.EnvId)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.EnvId)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.PartyId) > 0 {
		i -= len(m.PartyId)
		copy(dAtA[i:], m.PartyId)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.PartyId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.DataId) > 0 {
		i -= len(m.DataId)
		copy(dAtA[i:], m.DataId)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.DataId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ContractId) > 0 {
		i -= len(m.ContractId)
		copy(dAtA[i:], m.ContractId)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.ContractId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TaskId) > 0 {
		i -= len(m.TaskId)
		copy(dAtA[i:], m.TaskId)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.TaskId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TaskReadyGoReq_Peer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TaskReadyGoReq_Peer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TaskReadyGoReq_Peer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Party) > 0 {
		i -= len(m.Party)
		copy(dAtA[i:], m.Party)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Party)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Port != 0 {
		i = encodeVarintCommon(dAtA, i, uint64(m.Port))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Ip) > 0 {
		i -= len(m.Ip)
		copy(dAtA[i:], m.Ip)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Ip)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TaskReadyGoReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TaskReadyGoReply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TaskReadyGoReply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Msg) > 0 {
		i -= len(m.Msg)
		copy(dAtA[i:], m.Msg)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Msg)))
		i--
		dAtA[i] = 0x12
	}
	if m.Ok {
		i--
		if m.Ok {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCommon(dAtA []byte, offset int, v uint64) int {
	offset -= sovCommon(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TaskReadyGoReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TaskId)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.ContractId)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.DataId)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.PartyId)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.EnvId)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	if len(m.Peers) > 0 {
		for _, e := range m.Peers {
			l = e.Size()
			n += 1 + l + sovCommon(uint64(l))
		}
	}
	l = len(m.ContractCfg)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	if len(m.DataParty) > 0 {
		for _, s := range m.DataParty {
			l = len(s)
			n += 1 + l + sovCommon(uint64(l))
		}
	}
	if len(m.ComputationParty) > 0 {
		for _, s := range m.ComputationParty {
			l = len(s)
			n += 1 + l + sovCommon(uint64(l))
		}
	}
	if len(m.ResultParty) > 0 {
		for _, s := range m.ResultParty {
			l = len(s)
			n += 1 + l + sovCommon(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TaskReadyGoReq_Peer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Ip)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.Port != 0 {
		n += 1 + sovCommon(uint64(m.Port))
	}
	l = len(m.Party)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TaskReadyGoReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ok {
		n += 2
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCommon(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCommon(x uint64) (n int) {
	return sovCommon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TaskReadyGoReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: TaskReadyGoReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TaskReadyGoReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaskId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TaskId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PartyId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PartyId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnvId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EnvId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Peers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Peers = append(m.Peers, &TaskReadyGoReq_Peer{})
			if err := m.Peers[len(m.Peers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractCfg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractCfg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataParty", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataParty = append(m.DataParty, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ComputationParty", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ComputationParty = append(m.ComputationParty, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResultParty", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ResultParty = append(m.ResultParty, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommon
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
func (m *TaskReadyGoReq_Peer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: Peer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Peer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ip", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ip = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			m.Port = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Port |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Party", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Party = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommon
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
func (m *TaskReadyGoReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: TaskReadyGoReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TaskReadyGoReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ok", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Ok = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommon
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
func skipCommon(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommon
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
					return 0, ErrIntOverflowCommon
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
					return 0, ErrIntOverflowCommon
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
				return 0, ErrInvalidLengthCommon
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCommon
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCommon
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCommon        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommon          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCommon = fmt.Errorf("proto: unexpected end of group")
)
