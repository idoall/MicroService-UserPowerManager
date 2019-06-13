package handler

import (
	"context"

	"github.com/casbin/casbin"
	proto "github.com/idoall/MicroService-UserPowerManager/srv/role/v1/proto"
	"github.com/idoall/MicroService-UserPowerManager/utils"
	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
	"github.com/idoall/MicroService-UserPowerManager/utils/jaeger"
	"github.com/micro/go-micro/errors"
)

var RoleS *casbin.Enforcer

type SrvRole struct{}

// GetPermissionsForUser 根据用户获取权限
func (e *SrvRole) GetPermissionsForUser(ctx context.Context, req *proto.ForUserRequest, rep *proto.GetPermissionsForUserResponse) error {
	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_Role_GetPermissionsForUser_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_ID := inner.NAMESPACE_MICROSERVICE_SRVROLE

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s Service [ROLE][GetPermissionsForUser] request", namespace_ID)
	}

	//取出用户组的所有权限
	columnPowerList := RoleS.GetPermissionsForUser("usergroup_" + req.User)
	for _, cv := range columnPowerList {
		twoString := []*proto.TwoString{
			&proto.TwoString{Two: cv},
		}
		rep.One = twoString
		// append(rep.One, twoString)
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Srv_Role_GetPermissionsForUser_End")
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

	namespace_ID := inner.NAMESPACE_MICROSERVICE_SRVROLE

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s Service [ROLE][DeletePermissionsForUser] request", namespace_ID)
	}

	//取出用户组的所有权限
	delName := "usergroup_" + req.User
	if !RoleS.DeletePermissionsForUser(delName) {
		return errors.BadRequest(namespace_ID, "删除失败")
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Srv_Role_DeletePermissionsForUser_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("delName", delName)
	}

	return nil
}

func (e *SrvRole) RemoveFilteredPolicy(ctx context.Context, req *proto.RemoveFilteredPolicyRequest, rep *proto.Empty) error {
	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_Role_RemoveFilteredPolicy_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_ID := inner.NAMESPACE_MICROSERVICE_SRVROLE

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s Service [ROLE][RemoveFilteredPolicy] request", namespace_ID)
	}

	//移除资源
	removeName := "usergroup_" + req.Role
	if !RoleS.RemoveFilteredPolicy(0, removeName) {
		return errors.BadRequest(namespace_ID, "删除失败")
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Srv_Role_RemoveFilteredPolicy_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("removeName", removeName)
	}

	return nil
}

func (e *SrvRole) AddPolicy(ctx context.Context, req *proto.AddPolicyRequest, rep *proto.Empty) error {
	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_Role_AddPolicy_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_ID := inner.NAMESPACE_MICROSERVICE_SRVROLE

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s Service [ROLE][AddPolicy] request", namespace_ID)
	}

	//添加权限
	if !RoleS.AddPolicy(req.S1, req.S2, req.S3, req.S4) {
		return errors.BadRequest(namespace_ID, "添加失败")
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Srv_Role_AddPolicy_End")
	if span != nil {
		defer span.Finish()
		// span.SetTag("NewID", rep.NewID)
	}

	return nil
}

func (e *SrvRole) GetRolesForUser(ctx context.Context, req *proto.GetRolesForUserRequest, rep *proto.GetRolesForUserResponse) error {
	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_Role_GetRolesForUser_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_ID := inner.NAMESPACE_MICROSERVICE_SRVROLE

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s Service [ROLE][GetRolesForUser] request", namespace_ID)
	}

	//取出用户组的所有权限
	rep.Roles = RoleS.GetRolesForUser(req.Name)

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Srv_Role_GetRolesForUser_End")
	if span != nil {
		defer span.Finish()
		// span.SetTag("NewID", rep.NewID)
	}

	return nil
}
