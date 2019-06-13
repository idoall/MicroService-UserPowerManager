// Code generated by protoc-gen-go. DO NOT EDIT.
// source: srv/srvusersgroup/v1/proto/proto.proto

// import public "google/protobuf/timestamp.proto";

package SrvUsersGroup

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

type UsersGroup struct {
	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// 用户组名称
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	// 所属组Id
	ParentID int64 `protobuf:"varint,3,opt,name=ParentID,proto3" json:"ParentID,omitempty"`
	// 排序
	Sorts int64 `protobuf:"varint,4,opt,name=Sorts,proto3" json:"Sorts,omitempty"`
	// 备注
	Note string `protobuf:"bytes,5,opt,name=Note,proto3" json:"Note,omitempty"`
	// 创建时间
	CreateTime int64 `protobuf:"varint,6,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	// 最后是错新时间
	LastUpdateTime       int64    `protobuf:"varint,7,opt,name=LastUpdateTime,proto3" json:"LastUpdateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsersGroup) Reset()         { *m = UsersGroup{} }
func (m *UsersGroup) String() string { return proto.CompactTextString(m) }
func (*UsersGroup) ProtoMessage()    {}
func (*UsersGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ee67adc06d7873a, []int{0}
}

func (m *UsersGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsersGroup.Unmarshal(m, b)
}
func (m *UsersGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsersGroup.Marshal(b, m, deterministic)
}
func (m *UsersGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsersGroup.Merge(m, src)
}
func (m *UsersGroup) XXX_Size() int {
	return xxx_messageInfo_UsersGroup.Size(m)
}
func (m *UsersGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_UsersGroup.DiscardUnknown(m)
}

var xxx_messageInfo_UsersGroup proto.InternalMessageInfo

func (m *UsersGroup) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *UsersGroup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UsersGroup) GetParentID() int64 {
	if m != nil {
		return m.ParentID
	}
	return 0
}

func (m *UsersGroup) GetSorts() int64 {
	if m != nil {
		return m.Sorts
	}
	return 0
}

func (m *UsersGroup) GetNote() string {
	if m != nil {
		return m.Note
	}
	return ""
}

func (m *UsersGroup) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *UsersGroup) GetLastUpdateTime() int64 {
	if m != nil {
		return m.LastUpdateTime
	}
	return 0
}

type AddRequest struct {
	Model                *UsersGroup `protobuf:"bytes,1,opt,name=Model,proto3" json:"Model,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ee67adc06d7873a, []int{1}
}

func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRequest.Unmarshal(m, b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
}
func (m *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(m, src)
}
func (m *AddRequest) XXX_Size() int {
	return xxx_messageInfo_AddRequest.Size(m)
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetModel() *UsersGroup {
	if m != nil {
		return m.Model
	}
	return nil
}

type AddResponse struct {
	NewID                int64    `protobuf:"varint,1,opt,name=NewID,proto3" json:"NewID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddResponse) Reset()         { *m = AddResponse{} }
func (m *AddResponse) String() string { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()    {}
func (*AddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ee67adc06d7873a, []int{2}
}

func (m *AddResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddResponse.Unmarshal(m, b)
}
func (m *AddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddResponse.Marshal(b, m, deterministic)
}
func (m *AddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddResponse.Merge(m, src)
}
func (m *AddResponse) XXX_Size() int {
	return xxx_messageInfo_AddResponse.Size(m)
}
func (m *AddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddResponse proto.InternalMessageInfo

func (m *AddResponse) GetNewID() int64 {
	if m != nil {
		return m.NewID
	}
	return 0
}

type GetListRequest struct {
	PageSize             int64    `protobuf:"varint,1,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
	CurrentPageIndex     int64    `protobuf:"varint,2,opt,name=CurrentPageIndex,proto3" json:"CurrentPageIndex,omitempty"`
	OrderBy              string   `protobuf:"bytes,3,opt,name=OrderBy,proto3" json:"OrderBy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetListRequest) Reset()         { *m = GetListRequest{} }
func (m *GetListRequest) String() string { return proto.CompactTextString(m) }
func (*GetListRequest) ProtoMessage()    {}
func (*GetListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ee67adc06d7873a, []int{3}
}

func (m *GetListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetListRequest.Unmarshal(m, b)
}
func (m *GetListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetListRequest.Marshal(b, m, deterministic)
}
func (m *GetListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetListRequest.Merge(m, src)
}
func (m *GetListRequest) XXX_Size() int {
	return xxx_messageInfo_GetListRequest.Size(m)
}
func (m *GetListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetListRequest proto.InternalMessageInfo

func (m *GetListRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *GetListRequest) GetCurrentPageIndex() int64 {
	if m != nil {
		return m.CurrentPageIndex
	}
	return 0
}

func (m *GetListRequest) GetOrderBy() string {
	if m != nil {
		return m.OrderBy
	}
	return ""
}

type GetListResponse struct {
	TotalCount           int64         `protobuf:"varint,1,opt,name=TotalCount,proto3" json:"TotalCount,omitempty"`
	List                 []*UsersGroup `protobuf:"bytes,2,rep,name=List,proto3" json:"List,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetListResponse) Reset()         { *m = GetListResponse{} }
func (m *GetListResponse) String() string { return proto.CompactTextString(m) }
func (*GetListResponse) ProtoMessage()    {}
func (*GetListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ee67adc06d7873a, []int{4}
}

func (m *GetListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetListResponse.Unmarshal(m, b)
}
func (m *GetListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetListResponse.Marshal(b, m, deterministic)
}
func (m *GetListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetListResponse.Merge(m, src)
}
func (m *GetListResponse) XXX_Size() int {
	return xxx_messageInfo_GetListResponse.Size(m)
}
func (m *GetListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetListResponse proto.InternalMessageInfo

func (m *GetListResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *GetListResponse) GetList() []*UsersGroup {
	if m != nil {
		return m.List
	}
	return nil
}

type GetRequest struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ee67adc06d7873a, []int{5}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

type GetResponse struct {
	Model                *UsersGroup `protobuf:"bytes,1,opt,name=Model,proto3" json:"Model,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ee67adc06d7873a, []int{6}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetModel() *UsersGroup {
	if m != nil {
		return m.Model
	}
	return nil
}

type UpdateRequest struct {
	Model                *UsersGroup `protobuf:"bytes,1,opt,name=Model,proto3" json:"Model,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ee67adc06d7873a, []int{7}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetModel() *UsersGroup {
	if m != nil {
		return m.Model
	}
	return nil
}

type UpdateResponse struct {
	// 是否更新成功
	Updated              int64    `protobuf:"varint,1,opt,name=Updated,proto3" json:"Updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ee67adc06d7873a, []int{8}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

type DeleteRequest struct {
	IdArray              []string `protobuf:"bytes,1,rep,name=IdArray,proto3" json:"IdArray,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ee67adc06d7873a, []int{9}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetIdArray() []string {
	if m != nil {
		return m.IdArray
	}
	return nil
}

type DeleteResponse struct {
	// 是否删除成功,批量删除不需要返回值
	Deleted              int64    `protobuf:"varint,1,opt,name=Deleted,proto3" json:"Deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ee67adc06d7873a, []int{10}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetDeleted() int64 {
	if m != nil {
		return m.Deleted
	}
	return 0
}

func init() {
	proto.RegisterType((*UsersGroup)(nil), "SrvUsersGroup.UsersGroup")
	proto.RegisterType((*AddRequest)(nil), "SrvUsersGroup.AddRequest")
	proto.RegisterType((*AddResponse)(nil), "SrvUsersGroup.AddResponse")
	proto.RegisterType((*GetListRequest)(nil), "SrvUsersGroup.GetListRequest")
	proto.RegisterType((*GetListResponse)(nil), "SrvUsersGroup.GetListResponse")
	proto.RegisterType((*GetRequest)(nil), "SrvUsersGroup.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "SrvUsersGroup.GetResponse")
	proto.RegisterType((*UpdateRequest)(nil), "SrvUsersGroup.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "SrvUsersGroup.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "SrvUsersGroup.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "SrvUsersGroup.DeleteResponse")
}

func init() {
	proto.RegisterFile("srv/srvusersgroup/v1/proto/proto.proto", fileDescriptor_4ee67adc06d7873a)
}

var fileDescriptor_4ee67adc06d7873a = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x5f, 0x6f, 0xd3, 0x3e,
	0x14, 0x6d, 0x9b, 0xfe, 0xf9, 0xed, 0x56, 0xed, 0x0f, 0x59, 0x3c, 0x98, 0x6a, 0xab, 0x2a, 0x23,
	0x4d, 0x63, 0x12, 0xad, 0x18, 0xcf, 0x54, 0x74, 0xad, 0x14, 0x05, 0x95, 0x81, 0xd2, 0xed, 0x9d,
	0x80, 0xaf, 0x46, 0xa5, 0xae, 0x2e, 0xb6, 0x53, 0x18, 0x1f, 0x8e, 0x0f, 0xc5, 0x27, 0x40, 0xb1,
	0x9d, 0xa4, 0x49, 0x18, 0x12, 0xbc, 0x44, 0x39, 0xf7, 0x9e, 0x9e, 0x7b, 0x7d, 0x8e, 0x1b, 0x38,
	0x55, 0x72, 0x3f, 0x51, 0x72, 0x1f, 0x2b, 0x94, 0xea, 0x56, 0x8a, 0x78, 0x37, 0xd9, 0xbf, 0x98,
	0xec, 0xa4, 0xd0, 0xc2, 0x3e, 0xc7, 0xe6, 0x49, 0x7a, 0x2b, 0xb9, 0xbf, 0x49, 0x38, 0x7e, 0xc2,
	0x61, 0x3f, 0xea, 0x00, 0x39, 0x24, 0x7d, 0x68, 0x04, 0x0b, 0x5a, 0x1f, 0xd5, 0xcf, 0xbc, 0xb0,
	0x11, 0x2c, 0x08, 0x81, 0xe6, 0x55, 0x74, 0x87, 0xb4, 0x31, 0xaa, 0x9f, 0x1d, 0x85, 0xe6, 0x9d,
	0x0c, 0xe0, 0xbf, 0xf7, 0x91, 0xc4, 0xad, 0x0e, 0x16, 0xd4, 0x33, 0xcc, 0x0c, 0x93, 0xc7, 0xd0,
	0x5a, 0x09, 0xa9, 0x15, 0x6d, 0x9a, 0x86, 0x05, 0x46, 0x45, 0x68, 0xa4, 0x2d, 0xa7, 0x22, 0x34,
	0x92, 0x21, 0xc0, 0x5c, 0x62, 0xa4, 0xf1, 0x7a, 0x7d, 0x87, 0xb4, 0x6d, 0xe8, 0x07, 0x15, 0x72,
	0x0a, 0xfd, 0x65, 0xa4, 0xf4, 0xcd, 0x8e, 0xa7, 0x9c, 0x8e, 0xe1, 0x94, 0xaa, 0xec, 0x15, 0xc0,
	0x8c, 0xf3, 0x10, 0xbf, 0xc4, 0xa8, 0x34, 0x99, 0x40, 0xeb, 0xad, 0xe0, 0xb8, 0x31, 0x47, 0xe8,
	0x5e, 0x3c, 0x19, 0x17, 0x4e, 0x3b, 0xce, 0x5f, 0x43, 0xcb, 0x63, 0x4f, 0xa1, 0x6b, 0x7e, 0xae,
	0x76, 0x62, 0xab, 0x30, 0xd9, 0xff, 0x0a, 0xbf, 0x66, 0x16, 0x58, 0xc0, 0x24, 0xf4, 0x7d, 0xd4,
	0xcb, 0xb5, 0xd2, 0xe9, 0x1c, 0xe3, 0xc1, 0x2d, 0xae, 0xd6, 0xdf, 0xd1, 0x51, 0x33, 0x4c, 0xce,
	0xe1, 0xd1, 0x3c, 0x96, 0x89, 0x21, 0x49, 0x29, 0xd8, 0x72, 0xfc, 0x66, 0xfc, 0xf3, 0xc2, 0x4a,
	0x9d, 0x50, 0xe8, 0xbc, 0x93, 0x1c, 0xe5, 0xe5, 0xbd, 0xb1, 0xf2, 0x28, 0x4c, 0x21, 0xfb, 0x00,
	0xff, 0x67, 0x33, 0xdd, 0x72, 0x43, 0x80, 0x6b, 0xa1, 0xa3, 0xcd, 0x5c, 0xc4, 0x5b, 0xed, 0xc6,
	0x1e, 0x54, 0xc8, 0x73, 0x68, 0x26, 0x7c, 0xda, 0x18, 0x79, 0x7f, 0x3e, 0xbb, 0xa1, 0xb1, 0x63,
	0x00, 0x1f, 0xb3, 0x13, 0x95, 0x92, 0x67, 0x53, 0xe8, 0x9a, 0xae, 0x9b, 0xfd, 0xd7, 0xc6, 0xbe,
	0x86, 0x9e, 0x4d, 0xe9, 0x9f, 0xa3, 0x39, 0x87, 0x7e, 0xaa, 0xe0, 0x96, 0xa0, 0xd0, 0xb1, 0x15,
	0xee, 0x16, 0x4d, 0x21, 0x7b, 0x06, 0xbd, 0x05, 0x6e, 0x30, 0x9f, 0x46, 0xa1, 0x13, 0xf0, 0x99,
	0x94, 0xd1, 0x3d, 0xad, 0x8f, 0xbc, 0xc4, 0x58, 0x07, 0x13, 0xd9, 0x94, 0x9a, 0xcb, 0xda, 0x4a,
	0x26, 0xeb, 0xe0, 0xc5, 0xcf, 0x06, 0x14, 0xff, 0x2f, 0x64, 0x0a, 0xde, 0x8c, 0x73, 0x52, 0xde,
	0x3e, 0xbf, 0x82, 0x83, 0xc1, 0xef, 0x5a, 0x76, 0x12, 0xab, 0x91, 0x37, 0xd0, 0x71, 0xb1, 0x92,
	0x93, 0x12, 0xb1, 0x78, 0xc5, 0x06, 0xc3, 0x87, 0xda, 0x99, 0xd6, 0x14, 0x3c, 0x1f, 0x75, 0x65,
	0x97, 0x3c, 0xd4, 0xca, 0x2e, 0x07, 0x89, 0xb2, 0x1a, 0xf1, 0xa1, 0x6d, 0xfd, 0x23, 0xc7, 0xe5,
	0x30, 0x0e, 0x93, 0x1b, 0x9c, 0x3c, 0xd0, 0xcd, 0x84, 0x96, 0xd0, 0xbd, 0x8c, 0xf4, 0xa7, 0xcf,
	0xd6, 0xb6, 0x8a, 0x5a, 0x21, 0x99, 0x8a, 0x5a, 0x31, 0x0c, 0x56, 0xfb, 0xd8, 0x36, 0x1f, 0xaa,
	0x97, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x64, 0xe8, 0x61, 0x8d, 0xd2, 0x04, 0x00, 0x00,
}