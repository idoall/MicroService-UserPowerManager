package handler

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/idoall/TokenExchangeCommon/commonutils"

	"github.com/idoall/MicroService-UserPowerManager/utils"

	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
	"github.com/idoall/MicroService-UserPowerManager/utils/jaeger"

	srvProto "github.com/idoall/MicroService-UserPowerManager/srv/srvusersgroup/v1/proto"

	api "github.com/micro/go-api/proto"
	"github.com/micro/go-micro/errors"
)

// ApiUsersGroup struct
type ApiUsersGroup struct {
	Client srvProto.SrvUsersGroupService
}

// swagger:route POST /mshk/api/v1/ApiUsersGroup/add users addPet
// Add 添加一个用户组
func (e *ApiUsersGroup) Add(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_UserGroup_Add_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_ID := inner.NAMESPACE_MICROSERVICE_API

	// debug
	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [ApiUsersGroup][Add] request", namespace_ID)
	}

	// 获取请求参数 - 开始
	var name, note string
	var sorts int64
	if req.Post["Name"] == nil || req.Post["Name"].Values[0] == "" {
		return errors.InternalServerError(namespace_ID, "Name 不能为空")
	} else {
		name = req.Post["Name"].Values[0]
	}

	if req.Post["Sorts"] == nil || req.Post["Sorts"].Values[0] == "" {
		return errors.InternalServerError(namespace_ID, "Sorts 不能为空")
	} else {
		if sorts, err = commonutils.Int64FromString(req.Post["Sorts"].Values[0]); err != nil {
			return errors.InternalServerError(namespace_ID, "Sorts Int64FromString Error:"+err.Error())
		}
	}
	if req.Post["Note"] != nil {
		note = req.Post["Note"].Values[0]
	}
	// 获取请求参数 - 结束

	// make request
	requestModel := &srvProto.UsersGroup{
		Name:     name,
		ParentID: 0,
		Sorts:    sorts,
		Note:     note,
	}

	// 调用服务端方法
	srvResponse, err := e.Client.Add(ctx, &srvProto.AddRequest{Model: requestModel})
	if err != nil {
		return errors.InternalServerError(namespace_ID, err.Error())
	}

	// 输出的 json
	srvResponseJSON := struct {
		NewID int64 `json:"newid"`
	}{}
	srvResponseJSON.NewID = srvResponse.NewID
	b, _ := commonutils.JSONEncode(srvResponseJSON)
	rsp.StatusCode = 200
	rsp.Body = string(b)

	// debug
	if utils.RunMode == "dev" {
		inner.Mlogger.Info(srvResponse)
		inner.Mlogger.Info("rsp.Body", rsp.Body)
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Api_UserGroup_Add_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("NewID", srvResponse.NewID)
	}

	return nil
}

// GetList 获取用户列表,默认 id 倒排序
func (e *ApiUsersGroup) GetList(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_UserGroup_GetList_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_ID := inner.NAMESPACE_MICROSERVICE_API

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [ApiUsersGroup][GetList] request", namespace_ID)
	}

	// 获取请求参数 - 开始
	var pageSize, currentPageIndex int64
	var orderBy string
	if req.Get["PageSize"] == nil || req.Get["PageSize"].Values[0] == "" {
		return errors.InternalServerError(namespace_ID, "PageSize 不能为空")
	} else if pageSize, err = commonutils.Int64FromString(req.Get["PageSize"].Values[0]); err != nil {
		return errors.InternalServerError(namespace_ID, "PageSize Format Error:%s", err.Error())
	}

	if req.Get["CurrentPageIndex"] == nil || req.Get["CurrentPageIndex"].Values[0] == "" {
		return errors.InternalServerError(namespace_ID, "CurrentPageIndex 不能为空")
	} else if currentPageIndex, err = commonutils.Int64FromString(req.Get["CurrentPageIndex"].Values[0]); err != nil {
		return errors.InternalServerError(namespace_ID, "CurrentPageIndex Format Error:%s", err.Error())

	}

	if req.Get["OrderBy"] != nil {
		orderBy = req.Get["OrderBy"].Values[0]
	}
	// 获取请求参数 - 结束

	// return json
	jsonList := struct {
		Rows  []*srvProto.UsersGroup `json:"rows"`
		Total int64                  `json:"total"`
	}{}

	// 调用服务端方法
	srvResponse, err := e.Client.GetList(ctx, &srvProto.GetListRequest{
		CurrentPageIndex: currentPageIndex,
		PageSize:         pageSize,
		OrderBy:          orderBy,
	})
	if err != nil {
		if !commonutils.StringContains(err.Error(), "no row found") {
			return errors.InternalServerError(namespace_ID, err.Error())
		}
	} else {
		jsonList.Rows = srvResponse.List
		jsonList.Total = srvResponse.TotalCount
	}

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
		// span.SetTag("TotalCount", srvResponse.TotalCount)
	}

	return nil
}

// Get 获取单个用户组，根据Id
func (e *ApiUsersGroup) Get(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_UserGroup_GetUser_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_ID := inner.NAMESPACE_MICROSERVICE_API

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [Get] request", namespace_ID)
	}

	// 获取请求参数 - 开始
	var ID int64
	if req.Get["ID"] != nil && req.Get["ID"].Values[0] != "0" {
		if ID, err = commonutils.Int64FromString(req.Get["ID"].Values[0]); err != nil {
			return errors.InternalServerError(namespace_ID, "ID Format Error:%s", err.Error())
		}
	}

	// 获取请求参数 - 结束

	// 调用服务端方法
	srvResponse, err := e.Client.Get(ctx, &srvProto.GetRequest{
		ID: ID,
	})
	if err != nil {
		return errors.InternalServerError(namespace_ID, err.Error())
	}

	// 对 json 序列化并输出
	b, _ := json.Marshal(&responseJSONRow{
		ID:             srvResponse.Model.ID,
		Name:           srvResponse.Model.Name,
		ParentID:       srvResponse.Model.ParentID,
		Sorts:          srvResponse.Model.Sorts,
		Note:           srvResponse.Model.Note,
		CreateTime:     srvResponse.Model.CreateTime,
		LastUpdateTime: srvResponse.Model.LastUpdateTime,
	})
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
		span.SetTag("ID", ID)
	}

	return nil
}

// Update 修改用户信息
func (e *ApiUsersGroup) Update(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_UserGroup_Update_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_ID := inner.NAMESPACE_MICROSERVICE_API

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [ApiUsersGroup][Update] request", namespace_ID)
	}

	// 获取请求参数 - 开始
	var ID, sorts int64
	var name, note string
	if req.Post["ID"] == nil || req.Post["ID"].Values[0] == "" {
		return errors.InternalServerError(namespace_ID, "ID 不能为空")
	} else if ID, err = commonutils.Int64FromString(req.Post["ID"].Values[0]); err != nil {
		return errors.InternalServerError(namespace_ID, "ID Format Error:%s", err.Error())
	}
	if req.Post["Name"] == nil || req.Post["Name"].Values[0] == "" {
		return errors.InternalServerError(namespace_ID, "Name 不能为空")
	} else {
		name = req.Post["Name"].Values[0]
	}

	if req.Post["Sorts"] == nil || req.Post["Sorts"].Values[0] == "" {
		return errors.InternalServerError(namespace_ID, "Sorts 不能为空")
	} else {
		if sorts, err = commonutils.Int64FromString(req.Post["Sorts"].Values[0]); err != nil {
			return errors.InternalServerError(namespace_ID, "Sorts Int64FromString Error:"+err.Error())
		}
	}
	if req.Post["Note"] != nil {
		note = req.Post["Note"].Values[0]
	}
	// 获取请求参数 - 结束

	// 调用服务端方法 - 修改用户组
	srvResponseUpdateUser := &srvProto.UsersGroup{
		ID:       ID,
		Name:     name,
		ParentID: 0,
		Sorts:    sorts,
		Note:     note,
	}
	srvsrvResponse, err := e.Client.Update(ctx, &srvProto.UpdateRequest{Model: srvResponseUpdateUser})
	if err != nil {
		return errors.InternalServerError(namespace_ID, err.Error())
	}

	// 输出的 json
	srvResponseJSON := struct {
		Updated int64
	}{}
	srvResponseJSON.Updated = srvsrvResponse.Updated
	b, _ := commonutils.JSONEncode(srvResponseJSON)
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
		span.SetTag("ID", ID)
		span.SetTag("Updated", srvsrvResponse.Updated)
	}

	return nil
}

// BatchDelete 批量删除用户组信息
func (e *ApiUsersGroup) BatchDelete(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_UserGroup_BatchDelete_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_ID := inner.NAMESPACE_MICROSERVICE_API

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [ApiUsersGroup][BatchDelete] request", namespace_ID)
	}

	// 获取请求参数 - 开始
	var IdArray []string
	if req.Post["IDArray"] == nil || req.Post["IDArray"].Values[0] == "" {
		return errors.InternalServerError(namespace_ID, "IDArray 不能为空")
	} else {
		IdArray = strings.Split(req.Post["IDArray"].Values[0], ",")
	}

	// 获取请求参数 - 结束

	// 调用服务端方法获取用户
	srvResponse, err := e.Client.BatchDelete(ctx, &srvProto.DeleteRequest{
		IdArray: IdArray,
	})
	if err != nil {
		return errors.InternalServerError(namespace_ID, err.Error())
	}

	// 输出的 json
	respJSON := struct {
		Deleted int64
	}{}
	respJSON.Deleted = srvResponse.Deleted
	b, _ := commonutils.JSONEncode(respJSON)
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

// responseJSONRow 这里重点讲一下，proto自动生成的每一个字段都会加上omitempty，当ParentID和Sorts为0时，会不显示该字段。所以重新转义
type responseJSONRow struct {
	ID             int64  `json:"ID"`
	Name           string `json:"Name"`
	ParentID       int64  `json:"ParentID"`
	Sorts          int64  `json:"Sorts"`
	Note           string `json:"Note"`
	CreateTime     int64  `json:"CreateTime"`
	LastUpdateTime int64  `json:"LastUpdateTime"`
}
