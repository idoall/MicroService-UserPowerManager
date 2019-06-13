// Code generated by protoc-gen-go. DO NOT EDIT.
// source: srv/srvrole/v1/proto/proto.proto

// import public "google/protobuf/timestamp.proto";

package SrvRole

import (
	fmt "fmt"
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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_47a92f34644fd9d5, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type ForUserRequest struct {
	User                 string   `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ForUserRequest) Reset()         { *m = ForUserRequest{} }
func (m *ForUserRequest) String() string { return proto.CompactTextString(m) }
func (*ForUserRequest) ProtoMessage()    {}
func (*ForUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47a92f34644fd9d5, []int{1}
}

func (m *ForUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ForUserRequest.Unmarshal(m, b)
}
func (m *ForUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ForUserRequest.Marshal(b, m, deterministic)
}
func (m *ForUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ForUserRequest.Merge(m, src)
}
func (m *ForUserRequest) XXX_Size() int {
	return xxx_messageInfo_ForUserRequest.Size(m)
}
func (m *ForUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ForUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ForUserRequest proto.InternalMessageInfo

func (m *ForUserRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

type GetPermissionsForUserResponse struct {
	One                  []*TwoString `protobuf:"bytes,1,rep,name=One,proto3" json:"One,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetPermissionsForUserResponse) Reset()         { *m = GetPermissionsForUserResponse{} }
func (m *GetPermissionsForUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetPermissionsForUserResponse) ProtoMessage()    {}
func (*GetPermissionsForUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_47a92f34644fd9d5, []int{2}
}

func (m *GetPermissionsForUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPermissionsForUserResponse.Unmarshal(m, b)
}
func (m *GetPermissionsForUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPermissionsForUserResponse.Marshal(b, m, deterministic)
}
func (m *GetPermissionsForUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPermissionsForUserResponse.Merge(m, src)
}
func (m *GetPermissionsForUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetPermissionsForUserResponse.Size(m)
}
func (m *GetPermissionsForUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPermissionsForUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPermissionsForUserResponse proto.InternalMessageInfo

func (m *GetPermissionsForUserResponse) GetOne() []*TwoString {
	if m != nil {
		return m.One
	}
	return nil
}

type TwoString struct {
	Two                  []string `protobuf:"bytes,1,rep,name=Two,proto3" json:"Two,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TwoString) Reset()         { *m = TwoString{} }
func (m *TwoString) String() string { return proto.CompactTextString(m) }
func (*TwoString) ProtoMessage()    {}
func (*TwoString) Descriptor() ([]byte, []int) {
	return fileDescriptor_47a92f34644fd9d5, []int{3}
}

func (m *TwoString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TwoString.Unmarshal(m, b)
}
func (m *TwoString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TwoString.Marshal(b, m, deterministic)
}
func (m *TwoString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TwoString.Merge(m, src)
}
func (m *TwoString) XXX_Size() int {
	return xxx_messageInfo_TwoString.Size(m)
}
func (m *TwoString) XXX_DiscardUnknown() {
	xxx_messageInfo_TwoString.DiscardUnknown(m)
}

var xxx_messageInfo_TwoString proto.InternalMessageInfo

func (m *TwoString) GetTwo() []string {
	if m != nil {
		return m.Two
	}
	return nil
}

type DeletePermissionsForUserResponse struct {
	IsDel                bool     `protobuf:"varint,1,opt,name=IsDel,proto3" json:"IsDel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePermissionsForUserResponse) Reset()         { *m = DeletePermissionsForUserResponse{} }
func (m *DeletePermissionsForUserResponse) String() string { return proto.CompactTextString(m) }
func (*DeletePermissionsForUserResponse) ProtoMessage()    {}
func (*DeletePermissionsForUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_47a92f34644fd9d5, []int{4}
}

func (m *DeletePermissionsForUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePermissionsForUserResponse.Unmarshal(m, b)
}
func (m *DeletePermissionsForUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePermissionsForUserResponse.Marshal(b, m, deterministic)
}
func (m *DeletePermissionsForUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePermissionsForUserResponse.Merge(m, src)
}
func (m *DeletePermissionsForUserResponse) XXX_Size() int {
	return xxx_messageInfo_DeletePermissionsForUserResponse.Size(m)
}
func (m *DeletePermissionsForUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePermissionsForUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePermissionsForUserResponse proto.InternalMessageInfo

func (m *DeletePermissionsForUserResponse) GetIsDel() bool {
	if m != nil {
		return m.IsDel
	}
	return false
}

type RemoveFilteredPolicyRequest struct {
	Role                 string   `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveFilteredPolicyRequest) Reset()         { *m = RemoveFilteredPolicyRequest{} }
func (m *RemoveFilteredPolicyRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveFilteredPolicyRequest) ProtoMessage()    {}
func (*RemoveFilteredPolicyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47a92f34644fd9d5, []int{5}
}

func (m *RemoveFilteredPolicyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveFilteredPolicyRequest.Unmarshal(m, b)
}
func (m *RemoveFilteredPolicyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveFilteredPolicyRequest.Marshal(b, m, deterministic)
}
func (m *RemoveFilteredPolicyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveFilteredPolicyRequest.Merge(m, src)
}
func (m *RemoveFilteredPolicyRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveFilteredPolicyRequest.Size(m)
}
func (m *RemoveFilteredPolicyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveFilteredPolicyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveFilteredPolicyRequest proto.InternalMessageInfo

func (m *RemoveFilteredPolicyRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

type AddPolicyRequest struct {
	S1                   string   `protobuf:"bytes,1,opt,name=s1,proto3" json:"s1,omitempty"`
	S2                   string   `protobuf:"bytes,2,opt,name=s2,proto3" json:"s2,omitempty"`
	S3                   string   `protobuf:"bytes,3,opt,name=s3,proto3" json:"s3,omitempty"`
	S4                   string   `protobuf:"bytes,4,opt,name=s4,proto3" json:"s4,omitempty"`
	S5                   string   `protobuf:"bytes,5,opt,name=s5,proto3" json:"s5,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddPolicyRequest) Reset()         { *m = AddPolicyRequest{} }
func (m *AddPolicyRequest) String() string { return proto.CompactTextString(m) }
func (*AddPolicyRequest) ProtoMessage()    {}
func (*AddPolicyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47a92f34644fd9d5, []int{6}
}

func (m *AddPolicyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPolicyRequest.Unmarshal(m, b)
}
func (m *AddPolicyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPolicyRequest.Marshal(b, m, deterministic)
}
func (m *AddPolicyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPolicyRequest.Merge(m, src)
}
func (m *AddPolicyRequest) XXX_Size() int {
	return xxx_messageInfo_AddPolicyRequest.Size(m)
}
func (m *AddPolicyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPolicyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddPolicyRequest proto.InternalMessageInfo

func (m *AddPolicyRequest) GetS1() string {
	if m != nil {
		return m.S1
	}
	return ""
}

func (m *AddPolicyRequest) GetS2() string {
	if m != nil {
		return m.S2
	}
	return ""
}

func (m *AddPolicyRequest) GetS3() string {
	if m != nil {
		return m.S3
	}
	return ""
}

func (m *AddPolicyRequest) GetS4() string {
	if m != nil {
		return m.S4
	}
	return ""
}

func (m *AddPolicyRequest) GetS5() string {
	if m != nil {
		return m.S5
	}
	return ""
}

type GetRolesForUserRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRolesForUserRequest) Reset()         { *m = GetRolesForUserRequest{} }
func (m *GetRolesForUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetRolesForUserRequest) ProtoMessage()    {}
func (*GetRolesForUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_47a92f34644fd9d5, []int{7}
}

func (m *GetRolesForUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRolesForUserRequest.Unmarshal(m, b)
}
func (m *GetRolesForUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRolesForUserRequest.Marshal(b, m, deterministic)
}
func (m *GetRolesForUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRolesForUserRequest.Merge(m, src)
}
func (m *GetRolesForUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetRolesForUserRequest.Size(m)
}
func (m *GetRolesForUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRolesForUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRolesForUserRequest proto.InternalMessageInfo

func (m *GetRolesForUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetRolesForUserResponse struct {
	Roles                []string `protobuf:"bytes,1,rep,name=Roles,proto3" json:"Roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRolesForUserResponse) Reset()         { *m = GetRolesForUserResponse{} }
func (m *GetRolesForUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetRolesForUserResponse) ProtoMessage()    {}
func (*GetRolesForUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_47a92f34644fd9d5, []int{8}
}

func (m *GetRolesForUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRolesForUserResponse.Unmarshal(m, b)
}
func (m *GetRolesForUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRolesForUserResponse.Marshal(b, m, deterministic)
}
func (m *GetRolesForUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRolesForUserResponse.Merge(m, src)
}
func (m *GetRolesForUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetRolesForUserResponse.Size(m)
}
func (m *GetRolesForUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRolesForUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetRolesForUserResponse proto.InternalMessageInfo

func (m *GetRolesForUserResponse) GetRoles() []string {
	if m != nil {
		return m.Roles
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "SrvRole.Empty")
	proto.RegisterType((*ForUserRequest)(nil), "SrvRole.ForUserRequest")
	proto.RegisterType((*GetPermissionsForUserResponse)(nil), "SrvRole.GetPermissionsForUserResponse")
	proto.RegisterType((*TwoString)(nil), "SrvRole.TwoString")
	proto.RegisterType((*DeletePermissionsForUserResponse)(nil), "SrvRole.DeletePermissionsForUserResponse")
	proto.RegisterType((*RemoveFilteredPolicyRequest)(nil), "SrvRole.RemoveFilteredPolicyRequest")
	proto.RegisterType((*AddPolicyRequest)(nil), "SrvRole.AddPolicyRequest")
	proto.RegisterType((*GetRolesForUserRequest)(nil), "SrvRole.GetRolesForUserRequest")
	proto.RegisterType((*GetRolesForUserResponse)(nil), "SrvRole.GetRolesForUserResponse")
}

func init() { proto.RegisterFile("srv/srvrole/v1/proto/proto.proto", fileDescriptor_47a92f34644fd9d5) }

var fileDescriptor_47a92f34644fd9d5 = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x5f, 0xcb, 0xda, 0x30,
	0x14, 0xc6, 0xd5, 0xea, 0x5c, 0xcf, 0xc0, 0x49, 0x70, 0xb3, 0x73, 0xc8, 0x4a, 0x90, 0xe1, 0x60,
	0x58, 0x5a, 0x15, 0xbc, 0x1d, 0xf8, 0x87, 0xdd, 0x38, 0xa9, 0x6e, 0x17, 0xbb, 0xda, 0x9c, 0x87,
	0x51, 0x68, 0x1b, 0x97, 0x64, 0x15, 0x3f, 0xd2, 0xbe, 0xe5, 0x68, 0x5a, 0xf3, 0xfa, 0x6a, 0xf5,
	0xbd, 0x09, 0xe7, 0xc9, 0x39, 0x49, 0xce, 0x79, 0x7e, 0x2d, 0xd8, 0x82, 0x27, 0x8e, 0xe0, 0x09,
	0x67, 0x21, 0x3a, 0x89, 0xeb, 0xec, 0x39, 0x93, 0x2c, 0x5b, 0x07, 0x6a, 0x25, 0xf5, 0x35, 0x4f,
	0x7c, 0x16, 0x22, 0xad, 0x43, 0x6d, 0x16, 0xed, 0xe5, 0x91, 0xf6, 0xa0, 0x31, 0x67, 0xfc, 0xab,
	0x40, 0xee, 0xe3, 0x9f, 0xbf, 0x28, 0x24, 0x21, 0x50, 0x4d, 0xa5, 0x55, 0xb6, 0xcb, 0x7d, 0xd3,
	0x57, 0x31, 0x9d, 0x41, 0x77, 0x81, 0x72, 0x85, 0x3c, 0x0a, 0x84, 0x08, 0x58, 0x2c, 0xf4, 0x19,
	0xb1, 0x67, 0xb1, 0x40, 0xd2, 0x03, 0xe3, 0x4b, 0x8c, 0x56, 0xd9, 0x36, 0xfa, 0x2f, 0x3c, 0x32,
	0xc8, 0x9f, 0x19, 0x6c, 0x0e, 0x6c, 0x2d, 0x79, 0x10, 0xff, 0xf6, 0xd3, 0x34, 0xed, 0x82, 0xa9,
	0x77, 0x48, 0x13, 0x8c, 0xcd, 0x81, 0xa9, 0x23, 0xa6, 0x9f, 0x86, 0x74, 0x02, 0xf6, 0x14, 0x43,
	0x94, 0x78, 0xe7, 0xa1, 0x16, 0xd4, 0x3e, 0x8b, 0x29, 0x86, 0xaa, 0xbd, 0xe7, 0x7e, 0x26, 0xa8,
	0x0b, 0x6f, 0x7d, 0x8c, 0x58, 0x82, 0xf3, 0x20, 0x94, 0xc8, 0x71, 0xb7, 0x62, 0x61, 0xf0, 0xeb,
	0x78, 0x36, 0x52, 0xea, 0xc9, 0x69, 0xa4, 0x34, 0xa6, 0x5b, 0x68, 0x7e, 0xda, 0x5d, 0xd4, 0x35,
	0xa0, 0x22, 0xdc, 0xbc, 0xaa, 0x22, 0x5c, 0xa5, 0x3d, 0xab, 0x92, 0x6b, 0x4f, 0xe9, 0xa1, 0x65,
	0xe4, 0x7a, 0xa8, 0xf4, 0xc8, 0xaa, 0xe6, 0x7a, 0xa4, 0xf4, 0xd8, 0xaa, 0xe5, 0x7a, 0x4c, 0x3f,
	0xc2, 0xeb, 0x05, 0xca, 0xd4, 0x09, 0x71, 0x6d, 0xf2, 0xf2, 0x67, 0xa4, 0x3b, 0x4a, 0x63, 0xea,
	0x40, 0xfb, 0xaa, 0xfa, 0x61, 0x6a, 0xb5, 0x9f, 0xbb, 0x95, 0x09, 0xef, 0x9f, 0x01, 0x27, 0xa0,
	0xe4, 0x3b, 0xbc, 0x2a, 0x24, 0x44, 0xda, 0x1a, 0xc6, 0xe3, 0x16, 0x3a, 0xef, 0x75, 0xe2, 0x2e,
	0x5a, 0x5a, 0x22, 0x3f, 0xc0, 0xba, 0xc5, 0xe5, 0xf6, 0xf5, 0x1f, 0x74, 0xe2, 0x29, 0xa6, 0xb4,
	0x44, 0x96, 0xd0, 0x2a, 0xe2, 0x47, 0x7a, 0xfa, 0x92, 0x3b, 0x78, 0x3b, 0x0d, 0x5d, 0x95, 0x7d,
	0xd3, 0x25, 0x32, 0x01, 0x53, 0xc3, 0x25, 0x6f, 0x74, 0xfa, 0x12, 0x78, 0xc1, 0xc9, 0x6f, 0xf0,
	0xf2, 0x02, 0x02, 0x79, 0x77, 0x6e, 0x54, 0x01, 0xcc, 0x8e, 0x7d, 0xbb, 0xe0, 0x34, 0xe1, 0xf6,
	0x99, 0xfa, 0x01, 0x87, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0x95, 0xc2, 0x06, 0x93, 0xa4, 0x03,
	0x00, 0x00,
}