// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lib/common/base.proto

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

type UserType int32

const (
	UserType_User_Unknown UserType = 0
	UserType_User_1       UserType = 1
	UserType_User_2       UserType = 2
	UserType_User_3       UserType = 3
)

var UserType_name = map[int32]string{
	0: "User_Unknown",
	1: "User_1",
	2: "User_2",
	3: "User_3",
}

var UserType_value = map[string]int32{
	"User_Unknown": 0,
	"User_1":       1,
	"User_2":       2,
	"User_3":       3,
}

func (x UserType) String() string {
	return proto.EnumName(UserType_name, int32(x))
}

func (UserType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5dab74a081363d1b, []int{0}
}

// the status of data, N means normal, D means deleted.
type DataStatus int32

const (
	DataStatus_DataStatus_Unknown DataStatus = 0
	DataStatus_DataStatus_Normal  DataStatus = 1
	DataStatus_DataStatus_Deleted DataStatus = 2
)

var DataStatus_name = map[int32]string{
	0: "DataStatus_Unknown",
	1: "DataStatus_Normal",
	2: "DataStatus_Deleted",
}

var DataStatus_value = map[string]int32{
	"DataStatus_Unknown": 0,
	"DataStatus_Normal":  1,
	"DataStatus_Deleted": 2,
}

func (x DataStatus) String() string {
	return proto.EnumName(DataStatus_name, int32(x))
}

func (DataStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5dab74a081363d1b, []int{1}
}

// Y : normal, N non-normal
type CommonStatus int32

const (
	CommonStatus_CommonStatus_Unknown   CommonStatus = 0
	CommonStatus_CommonStatus_Normal    CommonStatus = 1
	CommonStatus_CommonStatus_NonNormal CommonStatus = 2
)

var CommonStatus_name = map[int32]string{
	0: "CommonStatus_Unknown",
	1: "CommonStatus_Normal",
	2: "CommonStatus_NonNormal",
}

var CommonStatus_value = map[string]int32{
	"CommonStatus_Unknown":   0,
	"CommonStatus_Normal":    1,
	"CommonStatus_NonNormal": 2,
}

func (x CommonStatus) String() string {
	return proto.EnumName(CommonStatus_name, int32(x))
}

func (CommonStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5dab74a081363d1b, []int{2}
}

type AuditMetadataOption int32

const (
	AuditMetadataOption_Audit_Pending AuditMetadataOption = 0
	AuditMetadataOption_Audit_Passed  AuditMetadataOption = 1
	AuditMetadataOption_Audit_Refused AuditMetadataOption = 2
)

var AuditMetadataOption_name = map[int32]string{
	0: "Audit_Pending",
	1: "Audit_Passed",
	2: "Audit_Refused",
}

var AuditMetadataOption_value = map[string]int32{
	"Audit_Pending": 0,
	"Audit_Passed":  1,
	"Audit_Refused": 2,
}

func (x AuditMetadataOption) String() string {
	return proto.EnumName(AuditMetadataOption_name, int32(x))
}

func (AuditMetadataOption) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5dab74a081363d1b, []int{3}
}

//原始文件类型
type OriginFileType int32

const (
	OriginFileType_FileType_Unknown OriginFileType = 0
	OriginFileType_FileType_CSV     OriginFileType = 1
)

var OriginFileType_name = map[int32]string{
	0: "FileType_Unknown",
	1: "FileType_CSV",
}

var OriginFileType_value = map[string]int32{
	"FileType_Unknown": 0,
	"FileType_CSV":     1,
}

func (x OriginFileType) String() string {
	return proto.EnumName(OriginFileType_name, int32(x))
}

func (OriginFileType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5dab74a081363d1b, []int{4}
}

// 元数据的状态 (0: 未知; 1: 还未发布的新表; 2: 已发布的表; 3: 已撤销的表)
type MetadataState int32

const (
	MetadataState_MetadataState_Unknown  MetadataState = 0
	MetadataState_MetadataState_Created  MetadataState = 1
	MetadataState_MetadataState_Released MetadataState = 2
	MetadataState_MetadataState_Revoked  MetadataState = 3
)

var MetadataState_name = map[int32]string{
	0: "MetadataState_Unknown",
	1: "MetadataState_Created",
	2: "MetadataState_Released",
	3: "MetadataState_Revoked",
}

var MetadataState_value = map[string]int32{
	"MetadataState_Unknown":  0,
	"MetadataState_Created":  1,
	"MetadataState_Released": 2,
	"MetadataState_Revoked":  3,
}

func (x MetadataState) String() string {
	return proto.EnumName(MetadataState_name, int32(x))
}

func (MetadataState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5dab74a081363d1b, []int{5}
}

// 元数据的使用方式类型枚举 (0: 未定义; 1: 按照时间段来使用; 2: 按照次数来使用)
type MetadataUsageType int32

const (
	MetadataUsageType_Usage_Unknown MetadataUsageType = 0
	MetadataUsageType_Usage_Period  MetadataUsageType = 1
	MetadataUsageType_Usage_Times   MetadataUsageType = 2
)

var MetadataUsageType_name = map[int32]string{
	0: "Usage_Unknown",
	1: "Usage_Period",
	2: "Usage_Times",
}

var MetadataUsageType_value = map[string]int32{
	"Usage_Unknown": 0,
	"Usage_Period":  1,
	"Usage_Times":   2,
}

func (x MetadataUsageType) String() string {
	return proto.EnumName(MetadataUsageType_name, int32(x))
}

func (MetadataUsageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5dab74a081363d1b, []int{6}
}

// 算力的状态 (0: 未知; 1: 还未发布的算力; 2: 已发布的算力(算力未被占用); 3: 已发布的算力(算力正在被占用); 4: 已撤销的算力)
type PowerState int32

const (
	PowerState_PowerState_Unknown    PowerState = 0
	PowerState_PowerState_Created    PowerState = 1
	PowerState_PowerState_Released   PowerState = 2
	PowerState_PowerState_Occupation PowerState = 3
	PowerState_PowerState_Revoked    PowerState = 4
)

var PowerState_name = map[int32]string{
	0: "PowerState_Unknown",
	1: "PowerState_Created",
	2: "PowerState_Released",
	3: "PowerState_Occupation",
	4: "PowerState_Revoked",
}

var PowerState_value = map[string]int32{
	"PowerState_Unknown":    0,
	"PowerState_Created":    1,
	"PowerState_Released":   2,
	"PowerState_Occupation": 3,
	"PowerState_Revoked":    4,
}

func (x PowerState) String() string {
	return proto.EnumName(PowerState_name, int32(x))
}

func (PowerState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5dab74a081363d1b, []int{7}
}

// 本组织在task中的角色
type TaskRole int32

const (
	TaskRole_TaskRole_Unknown       TaskRole = 0
	TaskRole_TaskRole_Sender        TaskRole = 1
	TaskRole_TaskRole_DataSupplier  TaskRole = 2
	TaskRole_TaskRole_PowerSupplier TaskRole = 3
	TaskRole_TaskRole_Receiver      TaskRole = 4
	TaskRole_TaskRole_AlgoSupplier  TaskRole = 5
)

var TaskRole_name = map[int32]string{
	0: "TaskRole_Unknown",
	1: "TaskRole_Sender",
	2: "TaskRole_DataSupplier",
	3: "TaskRole_PowerSupplier",
	4: "TaskRole_Receiver",
	5: "TaskRole_AlgoSupplier",
}

var TaskRole_value = map[string]int32{
	"TaskRole_Unknown":       0,
	"TaskRole_Sender":        1,
	"TaskRole_DataSupplier":  2,
	"TaskRole_PowerSupplier": 3,
	"TaskRole_Receiver":      4,
	"TaskRole_AlgoSupplier":  5,
}

func (x TaskRole) String() string {
	return proto.EnumName(TaskRole_name, int32(x))
}

func (TaskRole) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5dab74a081363d1b, []int{8}
}

// task的状态
type TaskState int32

const (
	TaskState_TaskState_Unknown TaskState = 0
	TaskState_TaskState_Pending TaskState = 1
	TaskState_TaskState_Running TaskState = 2
	TaskState_TaskState_Failed  TaskState = 3
	TaskState_TaskState_Succeed TaskState = 4
)

var TaskState_name = map[int32]string{
	0: "TaskState_Unknown",
	1: "TaskState_Pending",
	2: "TaskState_Running",
	3: "TaskState_Failed",
	4: "TaskState_Succeed",
}

var TaskState_value = map[string]int32{
	"TaskState_Unknown": 0,
	"TaskState_Pending": 1,
	"TaskState_Running": 2,
	"TaskState_Failed":  3,
	"TaskState_Succeed": 4,
}

func (x TaskState) String() string {
	return proto.EnumName(TaskState_name, int32(x))
}

func (TaskState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5dab74a081363d1b, []int{9}
}

// 数据授权信息的状态 (0: 未知; 1: 还未发布的数据授权; 2: 已发布的数据授权; 3: 已撤销的数据授权 <失效前主动撤回的>; 4: 已经失效的数据授权 <过期or达到使用上限的or被拒绝的>;)
type MetadataAuthorityState int32

const (
	MetadataAuthorityState_MAState_Unknown  MetadataAuthorityState = 0
	MetadataAuthorityState_MAState_Created  MetadataAuthorityState = 1
	MetadataAuthorityState_MAState_Released MetadataAuthorityState = 2
	MetadataAuthorityState_MAState_Revoked  MetadataAuthorityState = 3
	MetadataAuthorityState_MAState_Invalid  MetadataAuthorityState = 4
)

var MetadataAuthorityState_name = map[int32]string{
	0: "MAState_Unknown",
	1: "MAState_Created",
	2: "MAState_Released",
	3: "MAState_Revoked",
	4: "MAState_Invalid",
}

var MetadataAuthorityState_value = map[string]int32{
	"MAState_Unknown":  0,
	"MAState_Created":  1,
	"MAState_Released": 2,
	"MAState_Revoked":  3,
	"MAState_Invalid":  4,
}

func (x MetadataAuthorityState) String() string {
	return proto.EnumName(MetadataAuthorityState_name, int32(x))
}

func (MetadataAuthorityState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5dab74a081363d1b, []int{10}
}

// A represents the basic information of the organization.
type Organization struct {
	// org name
	NodeName string `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	// the node_id for org
	NodeId string `protobuf:"bytes,2,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// the identity for org
	IdentityId string `protobuf:"bytes,3,opt,name=identity_id,json=identityId,proto3" json:"identity_id,omitempty"`
	// the status for organization(deleted/normal)
	Status               CommonStatus `protobuf:"varint,4,opt,name=status,proto3,enum=api.protobuf.CommonStatus" json:"status,omitempty"`
	UpdateAt             uint64       `protobuf:"varint,5,opt,name=update_at,json=updateAt,proto3" json:"update_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Organization) Reset()         { *m = Organization{} }
func (m *Organization) String() string { return proto.CompactTextString(m) }
func (*Organization) ProtoMessage()    {}
func (*Organization) Descriptor() ([]byte, []int) {
	return fileDescriptor_5dab74a081363d1b, []int{0}
}
func (m *Organization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Organization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Organization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Organization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Organization.Merge(m, src)
}
func (m *Organization) XXX_Size() int {
	return m.Size()
}
func (m *Organization) XXX_DiscardUnknown() {
	xxx_messageInfo_Organization.DiscardUnknown(m)
}

var xxx_messageInfo_Organization proto.InternalMessageInfo

func (m *Organization) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *Organization) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *Organization) GetIdentityId() string {
	if m != nil {
		return m.IdentityId
	}
	return ""
}

func (m *Organization) GetStatus() CommonStatus {
	if m != nil {
		return m.Status
	}
	return CommonStatus_CommonStatus_Unknown
}

func (m *Organization) GetUpdateAt() uint64 {
	if m != nil {
		return m.UpdateAt
	}
	return 0
}

// 任务下的组织信息
type TaskOrganization struct {
	PartyId string `protobuf:"bytes,1,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	// org name
	NodeName string `protobuf:"bytes,2,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	// the node_id for org
	NodeId string `protobuf:"bytes,3,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// the identity for org
	IdentityId           string   `protobuf:"bytes,4,opt,name=identity_id,json=identityId,proto3" json:"identity_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskOrganization) Reset()         { *m = TaskOrganization{} }
func (m *TaskOrganization) String() string { return proto.CompactTextString(m) }
func (*TaskOrganization) ProtoMessage()    {}
func (*TaskOrganization) Descriptor() ([]byte, []int) {
	return fileDescriptor_5dab74a081363d1b, []int{1}
}
func (m *TaskOrganization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TaskOrganization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TaskOrganization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TaskOrganization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskOrganization.Merge(m, src)
}
func (m *TaskOrganization) XXX_Size() int {
	return m.Size()
}
func (m *TaskOrganization) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskOrganization.DiscardUnknown(m)
}

var xxx_messageInfo_TaskOrganization proto.InternalMessageInfo

func (m *TaskOrganization) GetPartyId() string {
	if m != nil {
		return m.PartyId
	}
	return ""
}

func (m *TaskOrganization) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *TaskOrganization) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *TaskOrganization) GetIdentityId() string {
	if m != nil {
		return m.IdentityId
	}
	return ""
}

type SimpleResponse struct {
	// the code for response
	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	// the message for response
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleResponse) Reset()         { *m = SimpleResponse{} }
func (m *SimpleResponse) String() string { return proto.CompactTextString(m) }
func (*SimpleResponse) ProtoMessage()    {}
func (*SimpleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5dab74a081363d1b, []int{2}
}
func (m *SimpleResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SimpleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SimpleResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SimpleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleResponse.Merge(m, src)
}
func (m *SimpleResponse) XXX_Size() int {
	return m.Size()
}
func (m *SimpleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleResponse proto.InternalMessageInfo

func (m *SimpleResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *SimpleResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterEnum("api.protobuf.UserType", UserType_name, UserType_value)
	proto.RegisterEnum("api.protobuf.DataStatus", DataStatus_name, DataStatus_value)
	proto.RegisterEnum("api.protobuf.CommonStatus", CommonStatus_name, CommonStatus_value)
	proto.RegisterEnum("api.protobuf.AuditMetadataOption", AuditMetadataOption_name, AuditMetadataOption_value)
	proto.RegisterEnum("api.protobuf.OriginFileType", OriginFileType_name, OriginFileType_value)
	proto.RegisterEnum("api.protobuf.MetadataState", MetadataState_name, MetadataState_value)
	proto.RegisterEnum("api.protobuf.MetadataUsageType", MetadataUsageType_name, MetadataUsageType_value)
	proto.RegisterEnum("api.protobuf.PowerState", PowerState_name, PowerState_value)
	proto.RegisterEnum("api.protobuf.TaskRole", TaskRole_name, TaskRole_value)
	proto.RegisterEnum("api.protobuf.TaskState", TaskState_name, TaskState_value)
	proto.RegisterEnum("api.protobuf.MetadataAuthorityState", MetadataAuthorityState_name, MetadataAuthorityState_value)
	proto.RegisterType((*Organization)(nil), "api.protobuf.Organization")
	proto.RegisterType((*TaskOrganization)(nil), "api.protobuf.TaskOrganization")
	proto.RegisterType((*SimpleResponse)(nil), "api.protobuf.SimpleResponse")
}

func init() { proto.RegisterFile("lib/common/base.proto", fileDescriptor_5dab74a081363d1b) }

var fileDescriptor_5dab74a081363d1b = []byte{
	// 764 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x95, 0xc1, 0x6e, 0xe3, 0x36,
	0x10, 0x86, 0x97, 0xb2, 0x93, 0x75, 0x66, 0x9d, 0x84, 0x91, 0x13, 0xc7, 0x9b, 0x02, 0x69, 0x90,
	0x53, 0x20, 0xb4, 0x36, 0x9a, 0xbd, 0x14, 0x7b, 0xaa, 0xeb, 0x45, 0x0a, 0x1f, 0xb2, 0x0e, 0xe4,
	0xa4, 0x87, 0x02, 0x85, 0x41, 0x5b, 0xb3, 0x5e, 0x22, 0x12, 0x29, 0x50, 0x54, 0x82, 0x14, 0xed,
	0xa5, 0xe8, 0xa1, 0x4f, 0xd1, 0xc7, 0xe8, 0x33, 0xf4, 0xd8, 0x47, 0x28, 0xf2, 0x24, 0x05, 0x29,
	0x4b, 0xa6, 0xbc, 0x68, 0x6f, 0xc3, 0xef, 0x27, 0x87, 0xff, 0xcc, 0x50, 0x36, 0x1c, 0xc5, 0x7c,
	0x3e, 0x58, 0xc8, 0x24, 0x91, 0x62, 0x30, 0x67, 0x19, 0xf6, 0x53, 0x25, 0xb5, 0xf4, 0xdb, 0x2c,
	0xe5, 0x45, 0x38, 0xcf, 0x3f, 0x9c, 0xff, 0x49, 0xa0, 0x3d, 0x51, 0x4b, 0x26, 0xf8, 0x4f, 0x4c,
	0x73, 0x29, 0xfc, 0xcf, 0x60, 0x47, 0xc8, 0x08, 0x67, 0x82, 0x25, 0xd8, 0x23, 0x67, 0xe4, 0x62,
	0x27, 0x6c, 0x19, 0xf0, 0x9e, 0x25, 0xe8, 0x1f, 0xc3, 0x4b, 0x2b, 0xf2, 0xa8, 0xe7, 0x59, 0x69,
	0xdb, 0x2c, 0xc7, 0x91, 0xff, 0x39, 0xbc, 0xe2, 0x11, 0x0a, 0xcd, 0xf5, 0x93, 0x11, 0x1b, 0x56,
	0x84, 0x12, 0x8d, 0x23, 0xff, 0x12, 0xb6, 0x33, 0xcd, 0x74, 0x9e, 0xf5, 0x9a, 0x67, 0xe4, 0x62,
	0xef, 0xf2, 0xa4, 0xef, 0xda, 0xe8, 0x8f, 0xac, 0xcd, 0xa9, 0xdd, 0x11, 0xae, 0x76, 0x1a, 0x2b,
	0x79, 0x1a, 0x31, 0x8d, 0x33, 0xa6, 0x7b, 0x5b, 0x67, 0xe4, 0xa2, 0x19, 0xb6, 0x0a, 0x30, 0xd4,
	0xe7, 0xbf, 0x11, 0xa0, 0xb7, 0x2c, 0xbb, 0xaf, 0x99, 0x7f, 0x0d, 0xad, 0x94, 0xa9, 0xc2, 0x43,
	0xe1, 0xfd, 0xa5, 0x5d, 0x8f, 0xa3, 0x7a, 0x5d, 0xde, 0x7f, 0xd7, 0xd5, 0xf8, 0xbf, 0xba, 0x9a,
	0x9b, 0x75, 0x9d, 0xbf, 0x85, 0xbd, 0x29, 0x4f, 0xd2, 0x18, 0x43, 0xcc, 0x52, 0x29, 0x32, 0xf4,
	0xbb, 0x55, 0xa5, 0xc6, 0xc1, 0x56, 0x55, 0x0d, 0x85, 0x46, 0x92, 0x2d, 0x57, 0x57, 0x9b, 0x30,
	0xf8, 0x06, 0x5a, 0x77, 0x19, 0xaa, 0xdb, 0xa7, 0x14, 0x7d, 0x0a, 0x6d, 0x13, 0xcf, 0xee, 0xc4,
	0xbd, 0x90, 0x8f, 0x82, 0xbe, 0xf0, 0x01, 0xb6, 0x2d, 0xf9, 0x8a, 0x92, 0x2a, 0xbe, 0xa4, 0x5e,
	0x15, 0xbf, 0xa1, 0x8d, 0x60, 0x0a, 0xf0, 0x8e, 0x69, 0x56, 0xf4, 0xcd, 0xef, 0x82, 0xbf, 0x5e,
	0x39, 0x99, 0x8e, 0xe0, 0xc0, 0xe1, 0xef, 0xa5, 0x4a, 0x58, 0x4c, 0xc9, 0xc6, 0xf6, 0x77, 0x18,
	0xa3, 0xc6, 0x88, 0x7a, 0xc1, 0x8f, 0xd0, 0x76, 0xc7, 0xe1, 0xf7, 0xe0, 0xd0, 0x5d, 0x3b, 0x89,
	0x8f, 0xa1, 0x53, 0x53, 0xaa, 0xd4, 0x27, 0xd0, 0xdd, 0x10, 0xc4, 0x4a, 0xf3, 0x82, 0x6b, 0xe8,
	0x0c, 0xf3, 0x88, 0xeb, 0x6b, 0xd4, 0x2c, 0x62, 0x9a, 0x4d, 0x52, 0x3b, 0xba, 0x03, 0xd8, 0xb5,
	0x78, 0x76, 0x83, 0x22, 0xe2, 0x62, 0x49, 0x5f, 0x98, 0x9e, 0xac, 0x10, 0xcb, 0x32, 0x8c, 0x28,
	0x59, 0x6f, 0x0a, 0xf1, 0x43, 0x9e, 0x59, 0xb7, 0x5f, 0xc3, 0xde, 0x44, 0xf1, 0x25, 0x17, 0x57,
	0x3c, 0x46, 0xdb, 0xca, 0x43, 0xa0, 0x65, 0xec, 0x78, 0xa5, 0xd0, 0xae, 0xe8, 0x68, 0xfa, 0x3d,
	0x25, 0xc1, 0xcf, 0xb0, 0x5b, 0x7a, 0x30, 0x36, 0xd1, 0x7f, 0x0d, 0x47, 0x35, 0xe0, 0x9c, 0xfe,
	0x44, 0x1a, 0x29, 0x64, 0xda, 0x7a, 0x3a, 0x81, 0x6e, 0x5d, 0x0a, 0x31, 0x46, 0x66, 0xcd, 0x7d,
	0x7a, 0x2c, 0xc4, 0x07, 0x79, 0x8f, 0x11, 0x6d, 0x04, 0x63, 0x38, 0x28, 0xa5, 0xbb, 0x8c, 0x2d,
	0x0b, 0xeb, 0x07, 0xb0, 0x6b, 0x17, 0x75, 0xdf, 0x05, 0xba, 0x41, 0xc5, 0xa5, 0xb9, 0x70, 0x1f,
	0x5e, 0x15, 0xe4, 0x96, 0x27, 0x98, 0x51, 0x2f, 0xf8, 0x9d, 0x00, 0xdc, 0xc8, 0x47, 0x54, 0x45,
	0x19, 0x5d, 0xf0, 0xd7, 0x2b, 0x27, 0x53, 0x9d, 0xaf, 0x0b, 0x38, 0x86, 0x8e, 0xc3, 0xeb, 0xee,
	0x1d, 0x61, 0xb2, 0x58, 0xe4, 0xa9, 0xfd, 0xcc, 0x68, 0x63, 0x23, 0x57, 0x59, 0x55, 0x33, 0xf8,
	0x83, 0x40, 0xcb, 0x7c, 0x95, 0xa1, 0x8c, 0xed, 0x20, 0xca, 0xd8, 0xb1, 0xd1, 0x81, 0xfd, 0x8a,
	0x4e, 0x51, 0x44, 0xa8, 0x28, 0x31, 0x57, 0x55, 0xd0, 0x3e, 0xca, 0x3c, 0x4d, 0x63, 0x8e, 0x8a,
	0x7a, 0xa6, 0xbf, 0x95, 0x54, 0xdc, 0x59, 0x6a, 0x0d, 0xf3, 0xb2, 0x2b, 0x2d, 0xc4, 0x05, 0xf2,
	0x07, 0x54, 0xb4, 0x59, 0xcb, 0x36, 0x8c, 0x97, 0xb2, 0x3a, 0xb1, 0x15, 0xfc, 0x02, 0x3b, 0x46,
	0x2a, 0x3a, 0xb5, 0x3a, 0xbe, 0xd9, 0xa8, 0x1a, 0x2e, 0x9f, 0x23, 0xa9, 0xe3, 0x30, 0x17, 0xc2,
	0x60, 0xaf, 0xac, 0xb2, 0xc0, 0x57, 0x8c, 0xc7, 0x66, 0xbc, 0xf5, 0xcd, 0xd3, 0x7c, 0xb1, 0x40,
	0xdb, 0x9f, 0x5f, 0xc9, 0xfa, 0xb5, 0x0c, 0x73, 0xfd, 0x51, 0x2a, 0xae, 0x9f, 0x0a, 0x33, 0x1d,
	0xd8, 0xbf, 0x1e, 0x6e, 0x5a, 0x71, 0xe0, 0x7a, 0x60, 0x87, 0x40, 0x4b, 0xe8, 0x4c, 0xcb, 0xd9,
	0x5a, 0xbd, 0x32, 0x17, 0x8e, 0xc5, 0x03, 0x8b, 0x79, 0x44, 0x9b, 0xdf, 0xbe, 0xfd, 0xeb, 0xf9,
	0x94, 0xfc, 0xfd, 0x7c, 0x4a, 0xfe, 0x79, 0x3e, 0x25, 0x3f, 0x7c, 0xb1, 0xe4, 0xfa, 0x63, 0x3e,
	0xef, 0x2f, 0x64, 0x32, 0x08, 0x65, 0x86, 0x5a, 0xb3, 0xab, 0x58, 0x3e, 0x0e, 0x46, 0x4c, 0x29,
	0x8e, 0xea, 0xcb, 0xef, 0xe4, 0x60, 0xfd, 0x47, 0x32, 0xdf, 0xb6, 0x3f, 0xd9, 0x6f, 0xfe, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0x19, 0x65, 0x72, 0xa2, 0x5d, 0x06, 0x00, 0x00,
}

func (m *Organization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Organization) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Organization) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.UpdateAt != 0 {
		i = encodeVarintBase(dAtA, i, uint64(m.UpdateAt))
		i--
		dAtA[i] = 0x28
	}
	if m.Status != 0 {
		i = encodeVarintBase(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x20
	}
	if len(m.IdentityId) > 0 {
		i -= len(m.IdentityId)
		copy(dAtA[i:], m.IdentityId)
		i = encodeVarintBase(dAtA, i, uint64(len(m.IdentityId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.NodeId) > 0 {
		i -= len(m.NodeId)
		copy(dAtA[i:], m.NodeId)
		i = encodeVarintBase(dAtA, i, uint64(len(m.NodeId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.NodeName) > 0 {
		i -= len(m.NodeName)
		copy(dAtA[i:], m.NodeName)
		i = encodeVarintBase(dAtA, i, uint64(len(m.NodeName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TaskOrganization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TaskOrganization) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TaskOrganization) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.IdentityId) > 0 {
		i -= len(m.IdentityId)
		copy(dAtA[i:], m.IdentityId)
		i = encodeVarintBase(dAtA, i, uint64(len(m.IdentityId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.NodeId) > 0 {
		i -= len(m.NodeId)
		copy(dAtA[i:], m.NodeId)
		i = encodeVarintBase(dAtA, i, uint64(len(m.NodeId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.NodeName) > 0 {
		i -= len(m.NodeName)
		copy(dAtA[i:], m.NodeName)
		i = encodeVarintBase(dAtA, i, uint64(len(m.NodeName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PartyId) > 0 {
		i -= len(m.PartyId)
		copy(dAtA[i:], m.PartyId)
		i = encodeVarintBase(dAtA, i, uint64(len(m.PartyId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SimpleResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SimpleResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SimpleResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		i = encodeVarintBase(dAtA, i, uint64(len(m.Msg)))
		i--
		dAtA[i] = 0x12
	}
	if m.Status != 0 {
		i = encodeVarintBase(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBase(dAtA []byte, offset int, v uint64) int {
	offset -= sovBase(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Organization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NodeName)
	if l > 0 {
		n += 1 + l + sovBase(uint64(l))
	}
	l = len(m.NodeId)
	if l > 0 {
		n += 1 + l + sovBase(uint64(l))
	}
	l = len(m.IdentityId)
	if l > 0 {
		n += 1 + l + sovBase(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovBase(uint64(m.Status))
	}
	if m.UpdateAt != 0 {
		n += 1 + sovBase(uint64(m.UpdateAt))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TaskOrganization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PartyId)
	if l > 0 {
		n += 1 + l + sovBase(uint64(l))
	}
	l = len(m.NodeName)
	if l > 0 {
		n += 1 + l + sovBase(uint64(l))
	}
	l = len(m.NodeId)
	if l > 0 {
		n += 1 + l + sovBase(uint64(l))
	}
	l = len(m.IdentityId)
	if l > 0 {
		n += 1 + l + sovBase(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SimpleResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovBase(uint64(m.Status))
	}
	l = len(m.Msg)
	if l > 0 {
		n += 1 + l + sovBase(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovBase(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBase(x uint64) (n int) {
	return sovBase(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Organization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBase
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
			return fmt.Errorf("proto: Organization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Organization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
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
				return ErrInvalidLengthBase
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBase
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
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
				return ErrInvalidLengthBase
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBase
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdentityId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
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
				return ErrInvalidLengthBase
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBase
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IdentityId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= CommonStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdateAt", wireType)
			}
			m.UpdateAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UpdateAt |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBase(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBase
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
func (m *TaskOrganization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBase
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
			return fmt.Errorf("proto: TaskOrganization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TaskOrganization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PartyId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
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
				return ErrInvalidLengthBase
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBase
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PartyId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
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
				return ErrInvalidLengthBase
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBase
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
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
				return ErrInvalidLengthBase
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBase
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdentityId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
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
				return ErrInvalidLengthBase
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBase
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IdentityId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBase(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBase
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
func (m *SimpleResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBase
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
			return fmt.Errorf("proto: SimpleResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SimpleResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
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
				return ErrInvalidLengthBase
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBase
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBase(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBase
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
func skipBase(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBase
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
					return 0, ErrIntOverflowBase
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
					return 0, ErrIntOverflowBase
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
				return 0, ErrInvalidLengthBase
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBase
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBase
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBase        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBase          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBase = fmt.Errorf("proto: unexpected end of group")
)
