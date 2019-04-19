// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: srv/srvusersgroup/proto/srvusersgroup.proto

// import public "google/protobuf/timestamp.proto";

package SrvUsersGroup

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for SrvUsersGroup service

type SrvUsersGroupService interface {
	// 添回用户组
	Add(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error)
	// 获取用户组列表
	GetList(ctx context.Context, in *GetListRequest, opts ...client.CallOption) (*GetListResponse, error)
	// 获取单个用户组
	Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
	// 修改用户组信息
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	// 批量删除用户组
	BatchDelete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
}

type srvUsersGroupService struct {
	c    client.Client
	name string
}

func NewSrvUsersGroupService(name string, c client.Client) SrvUsersGroupService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "SrvUsersGroup"
	}
	return &srvUsersGroupService{
		c:    c,
		name: name,
	}
}

func (c *srvUsersGroupService) Add(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error) {
	req := c.c.NewRequest(c.name, "SrvUsersGroup.Add", in)
	out := new(AddResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvUsersGroupService) GetList(ctx context.Context, in *GetListRequest, opts ...client.CallOption) (*GetListResponse, error) {
	req := c.c.NewRequest(c.name, "SrvUsersGroup.GetList", in)
	out := new(GetListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvUsersGroupService) Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.name, "SrvUsersGroup.Get", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvUsersGroupService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "SrvUsersGroup.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *srvUsersGroupService) BatchDelete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "SrvUsersGroup.BatchDelete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SrvUsersGroup service

type SrvUsersGroupHandler interface {
	// 添回用户组
	Add(context.Context, *AddRequest, *AddResponse) error
	// 获取用户组列表
	GetList(context.Context, *GetListRequest, *GetListResponse) error
	// 获取单个用户组
	Get(context.Context, *GetRequest, *GetResponse) error
	// 修改用户组信息
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	// 批量删除用户组
	BatchDelete(context.Context, *DeleteRequest, *DeleteResponse) error
}

func RegisterSrvUsersGroupHandler(s server.Server, hdlr SrvUsersGroupHandler, opts ...server.HandlerOption) error {
	type srvUsersGroup interface {
		Add(ctx context.Context, in *AddRequest, out *AddResponse) error
		GetList(ctx context.Context, in *GetListRequest, out *GetListResponse) error
		Get(ctx context.Context, in *GetRequest, out *GetResponse) error
		Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
		BatchDelete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
	}
	type SrvUsersGroup struct {
		srvUsersGroup
	}
	h := &srvUsersGroupHandler{hdlr}
	return s.Handle(s.NewHandler(&SrvUsersGroup{h}, opts...))
}

type srvUsersGroupHandler struct {
	SrvUsersGroupHandler
}

func (h *srvUsersGroupHandler) Add(ctx context.Context, in *AddRequest, out *AddResponse) error {
	return h.SrvUsersGroupHandler.Add(ctx, in, out)
}

func (h *srvUsersGroupHandler) GetList(ctx context.Context, in *GetListRequest, out *GetListResponse) error {
	return h.SrvUsersGroupHandler.GetList(ctx, in, out)
}

func (h *srvUsersGroupHandler) Get(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.SrvUsersGroupHandler.Get(ctx, in, out)
}

func (h *srvUsersGroupHandler) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.SrvUsersGroupHandler.Update(ctx, in, out)
}

func (h *srvUsersGroupHandler) BatchDelete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.SrvUsersGroupHandler.BatchDelete(ctx, in, out)
}
