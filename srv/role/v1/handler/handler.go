package handler

import (
	"context"

	"github.com/casbin/casbin/v2"
	proto "github.com/idoall/MicroService-UserPowerManager/srv/role/v1/proto"
	"github.com/idoall/MicroService-UserPowerManager/utils"
	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
	"github.com/idoall/MicroService-UserPowerManager/utils/jaeger"
	"github.com/micro/go-micro/errors"
)

// RoleS 全局角色对象
var RoleS *casbin.Enforcer

// SrvRole struct
type SrvRole struct{}

// GetPermissionsForUser 根据用户获取权限
func (e *SrvRole) GetPermissionsForUser(ctx context.Context, req *proto.ForUserRequest, rep *proto.GetPermissionsForUserResponse) error {
	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_Role_GetPermissionsForUser_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_SRVROLE

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s Service [ROLE][GetPermissionsForUser] request", namespaceID)
	}

	//取出用户组的所有权限
	columnPowerList := RoleS.GetPermissionsForUser(req.User)

	// 输出权限列表
	for _, cv := range columnPowerList {
		rep.One = append(rep.One, &proto.TwoString{Two: cv})
	}

	// 写入一个 jaeger span
	_, span = jaeger.StartSpan(ctx, "Srv_Role_GetPermissionsForUser_End")
	if span != nil {
		defer span.Finish()
		// span.SetTag("NewID", rep.NewID)
	}

	return nil
}

// DeletePermissionsForUser 根据用户删除权限
func (e *SrvRole) DeletePermissionsForUser(ctx context.Context, req *proto.ForUserRequest, rep *proto.DeletePermissionsForUserResponse) error {
	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_Role_DeletePermissionsForUser_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_SRVROLE

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s Service [ROLE][DeletePermissionsForUser] request", namespaceID)
	}

	//取出用户组的所有权限
	delName := "usergroup_" + req.User
	isDel, _ := RoleS.DeletePermissionsForUser(delName)
	if !isDel {
		return errors.BadRequest(namespaceID, "删除失败")
	}

	// 写入一个 jaeger span
	_, span = jaeger.StartSpan(ctx, "Srv_Role_DeletePermissionsForUser_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("delName", delName)
	}

	return nil
}

// RemoveFilteredPolicy 删除权限
func (e *SrvRole) RemoveFilteredPolicy(ctx context.Context, req *proto.RemoveFilteredPolicyRequest, rep *proto.Empty) error {
	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_Role_RemoveFilteredPolicy_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_SRVROLE

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s Service [ROLE][RemoveFilteredPolicy] request", namespaceID)
	}

	//移除资源
	removeName := "usergroup_" + req.Role
	if _, err := RoleS.RemoveFilteredPolicy(0, removeName); err != nil {
		return errors.New("RemoveFilteredPolicy", err.Error(), 500)
	}

	// 写入一个 jaeger span
	_, span = jaeger.StartSpan(ctx, "Srv_Role_RemoveFilteredPolicy_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("removeName", removeName)
	}

	return nil
}

// AddPolicy 添加权限
func (e *SrvRole) AddPolicy(ctx context.Context, req *proto.AddPolicyRequest, rep *proto.Empty) error {
	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_Role_AddPolicy_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_SRVROLE

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s Service [ROLE][AddPolicy] request", namespaceID)
	}

	//添加权限
	isADD, err := RoleS.AddPolicy(req.S1, req.S2, req.S3, req.S4)
	if !isADD {
		return errors.BadRequest(namespaceID, "添加失败:"+err.Error())
	}

	// 写入一个 jaeger span
	_, span = jaeger.StartSpan(ctx, "Srv_Role_AddPolicy_End")
	if span != nil {
		defer span.Finish()
		// span.SetTag("NewID", rep.NewID)
	}

	return nil
}

// AddGroupingPolicy 添加用户和角色（组）
func (e *SrvRole) AddGroupingPolicy(ctx context.Context, req *proto.AddGroupingPolicyRequest, rep *proto.Empty) error {
	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_Role_AddGroupingPolicy_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_SRVROLE

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s Service [ROLE][AddGroupingPolicy] request", namespaceID)
	}

	//添加
	for _, v := range req.UserGroup {
		groupName := "usergroup_" + v
		_, _ = RoleS.AddGroupingPolicy(req.User, groupName)
	}

	// 写入一个 jaeger span
	_, span = jaeger.StartSpan(ctx, "Srv_Role_AddGroupingPolicy_End")
	if span != nil {
		defer span.Finish()
		// span.SetTag("NewID", rep.NewID)
	}

	return nil
}

// GetRolesForUser 根据用户获取角色
func (e *SrvRole) GetRolesForUser(ctx context.Context, req *proto.GetRolesForUserRequest, rep *proto.GetRolesForUserResponse) error {
	var err error
	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_Role_GetRolesForUser_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_SRVROLE

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s Service [ROLE][GetRolesForUser][%s] request", namespaceID, req.Name)
	}

	//取出用户组的所有权限
	if rep.Roles, err = RoleS.GetRolesForUser(req.Name); err != nil {
		return err
	}

	// 写入一个 jaeger span
	_, span = jaeger.StartSpan(ctx, "Srv_Role_GetRolesForUser_End")
	if span != nil {
		defer span.Finish()
		// span.SetTag("NewID", rep.NewID)
	}

	return nil
}

// DeleteRolesForUser 根据用户删除角色
func (e *SrvRole) DeleteRolesForUser(ctx context.Context, req *proto.ForUserRequest, rep *proto.Empty) error {

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_Role_DeleteRolesForUser_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_SRVROLE

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s Service [ROLE][DeleteRolesForUser] request", namespaceID)
	}

	//取出用户组的所有权限
	if _, err := RoleS.DeleteRolesForUser(req.User); err != nil {
		return errors.New("DeleteRolesForUser", err.Error(), 500)
	}

	// 写入一个 jaeger span
	_, span = jaeger.StartSpan(ctx, "Srv_Role_DeleteRolesForUser_End")
	if span != nil {
		defer span.Finish()
		// span.SetTag("NewID", rep.NewID)
	}

	return nil
}
