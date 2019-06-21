package handler

import (
	"context"
	"strings"

	"github.com/idoall/TokenExchangeCommon/commonutils"

	"github.com/idoall/MicroService-UserPowerManager/utils"

	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
	"github.com/idoall/MicroService-UserPowerManager/utils/jaeger"

	srvProto "github.com/idoall/MicroService-UserPowerManager/srv/role/v1/proto"

	api "github.com/micro/go-api/proto"
	"github.com/micro/go-micro/errors"
)

// Role struct
type Role struct {
	Client srvProto.SrvRoleService
}

// route POST /mshk/api/v1/Role/add
// 根据用户获取权限
func (e *Role) GetPermissionsForRole(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_Role_GetPermissionsForRole_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_APIROLE

	// debug
	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [Role][GetPermissionsForRole] request", namespaceID)
	}

	// 获取请求参数 - 开始
	var user string
	if req.Get["User"] == nil || req.Get["User"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "User 不能为空")
	} else {
		user = req.Get["User"].Values[0]
	}

	// 调用服务端方法
	srvResponse, err := e.Client.GetPermissionsForUser(ctx, &srvProto.ForUserRequest{User: user})
	if err != nil {
		return errors.InternalServerError(namespaceID, err.Error())
	}

	// 输出的 json
	b, _ := commonutils.JSONEncode(srvResponse.One)
	rsp.StatusCode = 200
	rsp.Body = string(b)

	// debug
	if utils.RunMode == "dev" {
		inner.Mlogger.Info(srvResponse)
		inner.Mlogger.Info("rsp.Body", rsp.Body)
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Api_Role_GetPermissionsForRole_End")
	if span != nil {
		defer span.Finish()
	}

	return nil
}

// route POST /mshk/api/v1/Role/add
// 根据用户获取权限
func (e *Role) GetPermissionsForUser(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_Role_GetPermissionsForUser_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_APIROLE

	// debug
	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [Role][GetPermissionsForUser] request", namespaceID)
	}

	// 获取请求参数 - 开始
	var user string
	if req.Get["User"] == nil || req.Get["User"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "User 不能为空")
	} else {
		user = "usergroup_" + req.Get["User"].Values[0]
	}

	// 调用服务端方法
	srvResponse, err := e.Client.GetPermissionsForUser(ctx, &srvProto.ForUserRequest{User: user})
	if err != nil {
		return errors.InternalServerError(namespaceID, err.Error())
	}

	// 输出的 json
	b, _ := commonutils.JSONEncode(srvResponse.One)
	rsp.StatusCode = 200
	rsp.Body = string(b)

	// debug
	if utils.RunMode == "dev" {
		inner.Mlogger.Info(srvResponse)
		inner.Mlogger.Info("rsp.Body", rsp.Body)
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Api_Role_GetPermissionsForUser_End")
	if span != nil {
		defer span.Finish()
	}

	return nil
}

// DeletePermissionsForUser 删除用户所属权限
func (e *Role) DeletePermissionsForUser(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_Role_DeletePermissionsForUser_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_APIROLE

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [Role][DeletePermissionsForUser] request", namespaceID)
	}

	// 获取请求参数 - 开始
	var user string
	if req.Post["User"] == nil || req.Post["User"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "User 不能为空")
	} else {
		user = req.Post["User"].Values[0]
	}

	// 调用服务端方法
	srvResponse, err := e.Client.DeletePermissionsForUser(ctx, &srvProto.ForUserRequest{User: user})
	if err != nil {
		return errors.InternalServerError(namespaceID, err.Error())
	}

	// 输出的 json
	b, _ := commonutils.JSONEncode(srvResponse.IsDel)
	rsp.StatusCode = 200
	rsp.Body = string(b)

	// debug
	if utils.RunMode == "dev" {
		// inner.Mlogger.Info("rsp.Body", rsp.Body)
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Api_Role_DeletePermissionsForUser_End")
	if span != nil {
		defer span.Finish()
	}

	return nil
}

// RemoveFilteredPolicy 删除
func (e *Role) RemoveFilteredPolicy(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_Role_RemoveFilteredPolicy_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_APIROLE

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [Role][RemoveFilteredPolicy] request", namespaceID)
	}

	// 获取请求参数 - 开始
	var role string
	if req.Post["Role"] == nil || req.Post["Role"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "Role 不能为空")
	} else {
		role = req.Post["Role"].Values[0]
	}

	// 调用服务端方法
	_, err = e.Client.RemoveFilteredPolicy(ctx, &srvProto.RemoveFilteredPolicyRequest{Role: role})
	if err != nil {
		return errors.InternalServerError(namespaceID, err.Error())
	}

	// 输出的 json
	b, _ := commonutils.JSONEncode("")
	rsp.StatusCode = 200
	rsp.Body = string(b)

	// debug
	if utils.RunMode == "dev" {
		// inner.Mlogger.Info("rsp.Body", rsp.Body)
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Api_Role_RemoveFilteredPolicy_End")
	if span != nil {
		defer span.Finish()
	}

	return nil
}

// AddPolicy 添加权限
func (e *Role) AddPolicy(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_Role_AddPolicy_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_APIROLE

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [Role][AddPolicy] request", namespaceID)
	}

	// 获取请求参数 - 开始
	var s1, s2, s3, s4 string
	if req.Post["S1"] == nil || req.Post["S1"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "S1 不能为空")
	} else {
		s1 = req.Post["S1"].Values[0]
	}
	if req.Post["S2"] == nil || req.Post["S2"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "S2 不能为空")
	} else {
		s2 = req.Post["S2"].Values[0]
	}
	// if req.Post["S3"] == nil || req.Post["S3"].Values[0] == "" {
	// 	return errors.InternalServerError(namespaceID, "S3 不能为空")
	// } else {
	// 	s3 = req.Post["S3"].Values[0]
	// }
	if req.Post["S4"] == nil || req.Post["S4"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "S4 不能为空")
	} else {
		s4 = req.Post["S4"].Values[0]
	}

	// 调用服务端方法
	srvResponse, err := e.Client.AddPolicy(ctx, &srvProto.AddPolicyRequest{S1: s1, S2: s2, S3: s3, S4: s4})
	if err != nil {
		return errors.InternalServerError(namespaceID, err.Error())
	}

	// 输出的 json
	b, _ := commonutils.JSONEncode(srvResponse)
	rsp.StatusCode = 200
	rsp.Body = string(b)

	// debug
	if utils.RunMode == "dev" {
		// inner.Mlogger.Info("rsp.Body", rsp.Body)
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Api_Role_AddPolicy_End")
	if span != nil {
		defer span.Finish()
	}

	return nil
}

// AddGroupingPolicy 添加用户和角色（组）的关系
func (e *Role) AddGroupingPolicy(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_Role_AddGroupingPolicy_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_APIROLE

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [Role][AddGroupingPolicy] request", namespaceID)
	}

	// 获取请求参数 - 开始
	var user string
	var userGroup []string
	if req.Post["User"] == nil || req.Post["User"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "User 不能为空")
	} else {
		user = req.Post["User"].Values[0]
	}
	if req.Post["UserGroup"] == nil || req.Post["UserGroup"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "UserGroup 不能为空")
	} else {
		userGroup = strings.Split(req.Post["UserGroup"].Values[0], ",")
	}

	// 调用服务端方法
	srvResponse, err := e.Client.AddGroupingPolicy(ctx, &srvProto.AddGroupingPolicyRequest{User: user, UserGroup: userGroup})
	if err != nil {
		return errors.InternalServerError(namespaceID, err.Error())
	}

	// 输出的 json
	b, _ := commonutils.JSONEncode(srvResponse)
	rsp.StatusCode = 200
	rsp.Body = string(b)

	// debug
	if utils.RunMode == "dev" {
		// inner.Mlogger.Info("rsp.Body", rsp.Body)
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Api_Role_AddGroupingPolicy_End")
	if span != nil {
		defer span.Finish()
	}

	return nil
}

// GetRolesForUser 获取角色
func (e *Role) GetRolesForUser(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_Role_GetRolesForUser_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_APIROLE

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [Role][GetRolesForUser] request", namespaceID)
	}

	// 获取请求参数 - 开始
	var name string
	if req.Post["Name"] == nil || req.Post["Name"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "Name 不能为空")
	} else {
		name = req.Post["Name"].Values[0]
	}

	// 调用服务端方法
	srvResponse, err := e.Client.GetRolesForUser(ctx, &srvProto.GetRolesForUserRequest{Name: name})
	if err != nil {
		return errors.InternalServerError(namespaceID, err.Error())
	}

	// 输出的 json
	b, _ := commonutils.JSONEncode(srvResponse.Roles)
	rsp.StatusCode = 200
	rsp.Body = string(b)

	// debug
	if utils.RunMode == "dev" {
		// inner.Mlogger.Info("rsp.Body", rsp.Body)
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Api_Role_GetRolesForUser_End")
	if span != nil {
		defer span.Finish()
	}

	return nil
}

// DeleteRolesForUser 根据用户删除角色
func (e *Role) DeleteRolesForUser(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_Role_DeleteRolesForUser_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_APIROLE

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [Role][DeleteRolesForUser] request", namespaceID)
	}

	// 获取请求参数 - 开始
	var user string
	if req.Post["User"] == nil || req.Post["User"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "User 不能为空")
	} else {
		user = req.Post["User"].Values[0]
	}

	// 调用服务端方法
	srvResponse, err := e.Client.DeleteRolesForUser(ctx, &srvProto.ForUserRequest{User: user})
	if err != nil {
		return errors.InternalServerError(namespaceID, err.Error())
	}

	// 输出的 json
	b, _ := commonutils.JSONEncode(srvResponse)
	rsp.StatusCode = 200
	rsp.Body = string(b)

	// debug
	if utils.RunMode == "dev" {
		// inner.Mlogger.Info("rsp.Body", rsp.Body)
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Api_Role_DeleteRolesForUser_End")
	if span != nil {
		defer span.Finish()
	}

	return nil
}
