package handler

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/idoall/TokenExchangeCommon/commonutils"

	"github.com/idoall/MicroService-UserPowerManager/utils"

	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
	"github.com/idoall/MicroService-UserPowerManager/utils/jaeger"

	srvproto "github.com/idoall/MicroService-UserPowerManager/srv/srvusersgroup/proto"

	api "github.com/micro/go-api/proto"
	"github.com/micro/go-micro/errors"
)

// ApiUsersGroup struct
type ApiUsersGroup struct {
	Client srvproto.SrvUsersGroupService
}

// swagger:route POST /mshk/api/v1/ApiUsersGroup/add users addPet
// 添加一个用户组
func (e *ApiUsersGroup) Add(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_UserGroup_Add_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_id := inner.NAMESPACE_MICROSERVICE_API

	// debug
	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [Add] request", namespace_id)
	}

	// 获取请求参数 - 开始
	var name, note string
	var sorts int64
	if req.Post["Name"] == nil || req.Post["Name"].Values[0] == "" {
		return errors.InternalServerError(namespace_id, "Name 不能为空")
	} else {
		name = req.Post["Name"].Values[0]
	}

	if req.Post["Sorts"] == nil || req.Post["Sorts"].Values[0] == "" {
		return errors.InternalServerError(namespace_id, "Sorts 不能为空")
	} else {
		if sorts, err = commonutils.Int64FromString(req.Post["Sorts"]); err != nil {
			return errors.InternalServerError(namespace_id, "Sorts Int64FromString Error:"+err.Error())
		}
	}
	if req.Post["Note"] != nil {
		note = req.Post["Note"].Values[0]
	}
	// 获取请求参数 - 结束

	// make request
	requestModel := &srvproto.UsersGroup{
		Name:     name,
		ParentId: 0,
		Sorts:    sorts,
		Note:     note,
	}

	// 调用服务端方法
	response, err := e.Client.Add(ctx, &srvproto.AddRequest{Model: requestModel})
	if err != nil {
		return errors.InternalServerError(namespace_id, err.Error())
	}

	// 输出的 json
	responseJson := struct {
		NewId int64
	}{}
	responseJson.NewId = response.NewId
	b, _ := commonutils.JSONEncode(responseJson)
	rsp.StatusCode = 200
	rsp.Body = string(b)

	// debug
	if utils.RunMode == "dev" {
		inner.Mlogger.Info(response)
		inner.Mlogger.Info("rsp.Body", rsp.Body)
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Api_UserGroup_Add_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("NewId", response.NewId)
	}

	return nil
}

// 获取用户列表,默认 id 倒排序
func (e *ApiUsersGroup) GetList(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_UserGroup_GetList_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_id := inner.NAMESPACE_MICROSERVICE_API

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [GetList] request", namespace_id)
	}

	// 获取请求参数 - 开始
	var pageSize, currentPageIndex int64
	var orderBy string
	if req.Get["PageSize"] == nil || req.Get["PageSize"].Values[0] == "" {
		return errors.InternalServerError(namespace_id, "PageSize 不能为空")
	} else if pageSize, err = commonutils.Int64FromString(req.Get["PageSize"].Values[0]); err != nil {
		return errors.InternalServerError(namespace_id, "PageSize Format Error:%s", err.Error())
	}

	if req.Get["CurrentPageIndex"] == nil || req.Get["CurrentPageIndex"].Values[0] == "" {
		return errors.InternalServerError(namespace_id, "CurrentPageIndex 不能为空")
	} else if currentPageIndex, err = commonutils.Int64FromString(req.Get["CurrentPageIndex"].Values[0]); err != nil {
		return errors.InternalServerError(namespace_id, "CurrentPageIndex Format Error:%s", err.Error())

	}

	if req.Get["OrderBy"] != nil {
		orderBy = req.Get["OrderBy"].Values[0]
	}
	// 获取请求参数 - 结束

	// 调用服务端方法
	response, err := e.Client.GetList(ctx, &srvproto.GetListRequest{
		CurrentPageIndex: currentPageIndex,
		PageSize:         pageSize,
		OrderBy:          orderBy,
	})
	if err != nil {
		return errors.InternalServerError(namespace_id, err.Error())
	}

	// return json
	jsonList := struct {
		Rows  []*srvproto.UsersGroup `json:"rows"`
		Total int64                  `json:"total"`
	}{}
	jsonList.Rows = response.List
	jsonList.Total = response.TotalCount

	// 对 json 序列化并输出
	b, _ := json.Marshal(jsonList)
	rsp.StatusCode = 200
	rsp.Body = string(b)

	// debug
	if utils.RunMode == "dev" {
		// inner.Mlogger.Info("rsp.Body", rsp.Body)
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Api_UserGroup_GetList_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("PageSize", pageSize)
		span.SetTag("CurrentPageIndex", currentPageIndex)
		span.SetTag("orderBy", orderBy)
		span.SetTag("TotalCount", response.TotalCount)
	}

	return nil
}

// 获取单个用户组，根据Id
func (e *ApiUsersGroup) Get(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_UserGroup_GetUser_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_id := inner.NAMESPACE_MICROSERVICE_API

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [Get] request", namespace_id)
	}

	// 获取请求参数 - 开始
	var Id int64
	if req.Get["Id"] != nil && req.Get["Id"].Values[0] != "0" {
		if Id, err = commonutils.Int64FromString(req.Get["Id"].Values[0]); err != nil {
			return errors.InternalServerError(namespace_id, "Id Format Error:%s", err.Error())
		}
	}

	// 获取请求参数 - 结束

	// 调用服务端方法
	response, err := e.Client.Get(ctx, &srvproto.GetRequest{
		Id: Id,
	})
	if err != nil {
		return errors.InternalServerError(namespace_id, err.Error())
	}

	// 对 json 序列化并输出
	b, _ := json.Marshal(response)
	rsp.StatusCode = 200
	rsp.Body = string(b)

	// debug
	if utils.RunMode == "dev" {
		inner.Mlogger.Info("rsp.Body", rsp.Body)
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Api_UserGroup_GetUser_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("Id", Id)
	}

	return nil
}

// 修改用户信息
func (e *ApiUsersGroup) Update(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_UserGroup_Update_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_id := inner.NAMESPACE_MICROSERVICE_API

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [Update] request", namespace_id)
	}

	// 获取请求参数 - 开始
	var Id, sorts int64
	var name, note string
	if req.Post["Id"] == nil || req.Post["Id"].Values[0] == "" {
		return errors.InternalServerError(namespace_id, "Id 不能为空")
	} else if Id, err = commonutils.Int64FromString(req.Post["Id"].Values[0]); err != nil {
		return errors.InternalServerError(namespace_id, "Id Format Error:%s", err.Error())
	}
	if req.Post["Name"] == nil || req.Post["Name"].Values[0] == "" {
		return errors.InternalServerError(namespace_id, "Name 不能为空")
	} else {
		name = req.Post["Name"].Values[0]
	}

	if req.Post["Sorts"] == nil || req.Post["Sorts"].Values[0] == "" {
		return errors.InternalServerError(namespace_id, "Sorts 不能为空")
	} else {
		if sorts, err = commonutils.Int64FromString(req.Post["Sorts"]); err != nil {
			return errors.InternalServerError(namespace_id, "Sorts Int64FromString Error:"+err.Error())
		}
	}
	if req.Post["Note"] != nil {
		note = req.Post["Note"].Values[0]
	}
	// 获取请求参数 - 结束

	// 调用服务端方法 - 修改用户组
	responseUpdateUser := &srvproto.UsersGroup{
		Id:       Id,
		Name:     name,
		ParentId: 0,
		Sorts:    sorts,
		Note:     note,
	}
	srvResponse, err := e.Client.Update(ctx, &srvproto.UpdateRequest{Model: responseUpdateUser})
	if err != nil {
		return errors.InternalServerError(namespace_id, err.Error())
	}

	// 输出的 json
	responseJson := struct {
		Updated int64
	}{}
	responseJson.Updated = srvResponse.Updated
	b, _ := commonutils.JSONEncode(responseJson)
	rsp.StatusCode = 200
	rsp.Body = string(b)

	// debug
	if utils.RunMode == "dev" {
		inner.Mlogger.Info("rsp.Body", rsp.Body)
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Api_UserGroup_Update_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("Id", Id)
		span.SetTag("Updated", srvResponse.Updated)
	}

	return nil
}

// 批量删除用户组信息
func (e *ApiUsersGroup) BatchDelete(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_UserGroup_BatchDelete_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_id := inner.NAMESPACE_MICROSERVICE_API

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [BatchDelete] request", namespace_id)
	}

	// 获取请求参数 - 开始
	var IdArray []string
	if req.Post["Ids"] == nil || req.Post["Ids"].Values[0] == "" {
		return errors.InternalServerError(namespace_id, "Ids 不能为空")
	} else {
		IdArray = strings.Split(req.Post["Ids"].Values[0], ",")
	}

	// 获取请求参数 - 结束

	// 调用服务端方法获取用户
	response, err := e.Client.BatchDelete(ctx, &srvproto.DeleteRequest{
		IdArray: IdArray,
	})
	if err != nil {
		return errors.InternalServerError(namespace_id, err.Error())
	}

	// 输出的 json
	respJson := struct {
		Deleted int64
	}{}
	respJson.Deleted = response.Deleted
	b, _ := commonutils.JSONEncode(respJson)
	rsp.StatusCode = 200
	rsp.Body = string(b)

	// debug
	if utils.RunMode == "dev" {
		inner.Mlogger.Info("rsp.Body", rsp.Body)
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Api_UserGroup_BatchDelete_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("IdArray", strings.Join(IdArray, ","))
	}

	return nil
}
