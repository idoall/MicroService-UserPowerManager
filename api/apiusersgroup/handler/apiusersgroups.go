package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/idoall/TokenExchangeCommon/commonutils"

	"github.com/idoall/MicroService-UserPowerManager/utils"

	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
	"github.com/idoall/MicroService-UserPowerManager/utils/jaeger"

	srvProto "github.com/idoall/MicroService-UserPowerManager/srv/srvusersgroup/proto"

	api "github.com/micro/go-api/proto"
	"github.com/micro/go-micro/errors"
)

// ApiUsersGroup struct
type ApiUsersGroup struct {
	Client srvProto.SrvUsersGroupService
}

// Add 添加一个用户组
// swagger:route POST /mshk/api/v1/ApiUsersGroup/add users addPet
func (e *ApiUsersGroup) Add(ctx context.Context, req *api.Request, rsp *api.Response) error {

	namespaceID := inner.NAMESPACE_MICROSERVICE_API

	// 验证提交的方法是否符合要求
	if req.Method != "POST" {
		supportedMethods := "POST"
		return errors.InternalServerError(namespaceID, fmt.Sprintf("incorrect method supplied %s: supported %s", req.Method, supportedMethods))
	}

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_UsersGroup_Add_Begin")
	if span != nil {
		defer span.Finish()
	}

	// debug
	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [ApiUsersGroup][Add] request", namespaceID)
	}

	// 获取请求参数 - 开始
	var name, note string
	var sorts int64
	if req.Post["Name"] == nil || req.Post["Name"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "Name 不能为空")
	} else {
		name = req.Post["Name"].Values[0]
	}

	if req.Post["Sorts"] == nil || req.Post["Sorts"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "Sorts 不能为空")
	} else {
		if sorts, err = commonutils.Int64FromString(req.Post["Sorts"].Values[0]); err != nil {
			return errors.InternalServerError(namespaceID, "Sorts Int64FromString Error:"+err.Error())
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
	response, err := e.Client.Add(ctx, &srvProto.AddRequest{Model: requestModel})
	if err != nil {
		return errors.InternalServerError(namespaceID, err.Error())
	}

	// 输出的 json
	responseJSON := struct {
		NewID int64 `json:"newid"`
	}{}
	responseJSON.NewID = response.NewID
	b, _ := commonutils.JSONEncode(responseJSON)
	rsp.StatusCode = 200
	rsp.Body = string(b)

	// debug
	if utils.RunMode == "dev" {
		inner.Mlogger.Info(response)
		inner.Mlogger.Info("rsp.Body", rsp.Body)
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Api_UsersGroup_Add_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("NewID", response.NewID)
	}

	return nil
}

// GetList 获取用户列表,默认 id 倒排序
func (e *ApiUsersGroup) GetList(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_UsersGroup_GetList_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_API

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [ApiUsersGroup][GetList] request", namespaceID)
	}

	// 获取请求参数 - 开始
	var pageSize, currentPageIndex int64
	var orderBy string
	if req.Get["PageSize"] == nil || req.Get["PageSize"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "PageSize 不能为空")
	} else if pageSize, err = commonutils.Int64FromString(req.Get["PageSize"].Values[0]); err != nil {
		return errors.InternalServerError(namespaceID, "PageSize Format Error:%s", err.Error())
	}

	if req.Get["CurrentPageIndex"] == nil || req.Get["CurrentPageIndex"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "CurrentPageIndex 不能为空")
	} else if currentPageIndex, err = commonutils.Int64FromString(req.Get["CurrentPageIndex"].Values[0]); err != nil {
		return errors.InternalServerError(namespaceID, "CurrentPageIndex Format Error:%s", err.Error())

	}

	if req.Get["OrderBy"] != nil {
		orderBy = req.Get["OrderBy"].Values[0]
	}
	// 获取请求参数 - 结束

	// 调用服务端方法
	srvResponse, err := e.Client.GetList(ctx, &srvProto.GetListRequest{
		CurrentPageIndex: currentPageIndex,
		PageSize:         pageSize,
		OrderBy:          orderBy,
	})
	if err != nil {
		return errors.InternalServerError(namespaceID, err.Error())
	}

	// return json
	responseJSON := struct {
		Rows  []*responseJSONRow `json:"rows"`
		Total int64              `json:"total"`
	}{}
	for _, v := range srvResponse.List {
		r := &responseJSONRow{
			ID:             v.ID,
			Name:           v.Name,
			ParentID:       v.ParentID,
			Sorts:          v.Sorts,
			Note:           v.Note,
			CreateTime:     v.CreateTime,
			LastUpdateTime: v.LastUpdateTime,
		}
		responseJSON.Rows = append(responseJSON.Rows, r)
	}
	responseJSON.Total = srvResponse.TotalCount

	// 对 json 序列化并输出
	b, _ := json.Marshal(responseJSON)
	rsp.StatusCode = 200
	rsp.Body = string(b)

	// debug
	if utils.RunMode == "dev" {
		// inner.Mlogger.Info("rsp.Body", rsp.Body)
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Api_UsersGroup_GetList_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("PageSize", pageSize)
		span.SetTag("CurrentPageIndex", currentPageIndex)
		span.SetTag("orderBy", orderBy)
		span.SetTag("TotalCount", srvResponse.TotalCount)
	}

	return nil
}

// Get 获取单个用户组，根据Id
func (e *ApiUsersGroup) Get(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_UsersGroup_Get_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_API

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [ApiUsersGroup][Get] request", namespaceID)
	}

	// 获取请求参数 - 开始
	var ID int64
	if req.Get["ID"] != nil && req.Get["ID"].Values[0] != "0" {
		if ID, err = commonutils.Int64FromString(req.Get["ID"].Values[0]); err != nil {
			return errors.InternalServerError(namespaceID, "ID Format Error:%s", err.Error())
		}
	}

	// 获取请求参数 - 结束

	// 调用服务端方法
	response, err := e.Client.Get(ctx, &srvProto.GetRequest{
		ID: ID,
	})
	if err != nil {
		return errors.InternalServerError(namespaceID, err.Error())
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
	ctx, span = jaeger.StartSpan(ctx, "Api_UsersGroup_Get_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("ID", ID)
	}

	return nil
}

// 修改用户信息
func (e *ApiUsersGroup) Update(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_UsersGroup_Update_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_API

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [ApiUsersGroup][Update] request", namespaceID)
	}

	// 获取请求参数 - 开始
	var ID, sorts int64
	var name, note string
	if req.Post["ID"] == nil || req.Post["ID"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "ID 不能为空")
	} else if ID, err = commonutils.Int64FromString(req.Post["ID"].Values[0]); err != nil {
		return errors.InternalServerError(namespaceID, "ID Format Error:%s", err.Error())
	}
	if req.Post["Name"] == nil || req.Post["Name"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "Name 不能为空")
	} else {
		name = req.Post["Name"].Values[0]
	}

	if req.Post["Sorts"] == nil || req.Post["Sorts"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "Sorts 不能为空")
	} else {
		if sorts, err = commonutils.Int64FromString(req.Post["Sorts"].Values[0]); err != nil {
			return errors.InternalServerError(namespaceID, "Sorts Int64FromString Error:"+err.Error())
		}
	}
	if req.Post["Note"] != nil {
		note = req.Post["Note"].Values[0]
	}
	// 获取请求参数 - 结束

	// 调用服务端方法 - 修改用户组
	responseUpdateUser := &srvProto.UsersGroup{
		ID:       ID,
		Name:     name,
		ParentID: 0,
		Sorts:    sorts,
		Note:     note,
	}
	srvResponse, err := e.Client.Update(ctx, &srvProto.UpdateRequest{Model: responseUpdateUser})
	if err != nil {
		return errors.InternalServerError(namespaceID, err.Error())
	}

	// 输出的 json
	responseJSON := struct {
		Updated int64
	}{}
	responseJSON.Updated = srvResponse.Updated
	b, _ := commonutils.JSONEncode(responseJSON)
	rsp.StatusCode = 200
	rsp.Body = string(b)

	// debug
	if utils.RunMode == "dev" {
		inner.Mlogger.Info("rsp.Body", rsp.Body)
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Api_UsersGroup_Update_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("Id", ID)
		span.SetTag("Updated", srvResponse.Updated)
	}

	return nil
}

// 批量删除用户组信息
func (e *ApiUsersGroup) BatchDelete(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_UsersGroup_BatchDelete_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_API

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [ApiUsersGroup][BatchDelete] request", namespaceID)
	}

	// 获取请求参数 - 开始
	var IDArray []string
	if req.Post["IDArray"] == nil || req.Post["IDArray"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "IDArray 不能为空")
	} else {
		IDArray = strings.Split(req.Post["IDArray"].Values[0], ",")
	}

	// 获取请求参数 - 结束

	// 调用服务端方法获取用户
	response, err := e.Client.BatchDelete(ctx, &srvProto.DeleteRequest{
		IDArray: IDArray,
	})
	if err != nil {
		return errors.InternalServerError(namespaceID, err.Error())
	}

	// 输出的 json
	responseJSON := struct {
		Deleted int64
	}{}
	responseJSON.Deleted = response.Deleted
	b, _ := commonutils.JSONEncode(responseJSON)
	rsp.StatusCode = 200
	rsp.Body = string(b)

	// debug
	if utils.RunMode == "dev" {
		inner.Mlogger.Info("rsp.Body", rsp.Body)
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Api_UsersGroup_BatchDelete_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("IdArray", strings.Join(IDArray, ","))
	}

	return nil
}

// 这里重点讲一下，proto自动生成的每一个字段都会加上omitempty，当ParentID和Sorts为0时，会不显示该字段。所以重新转义
type responseJSONRow struct {
	ID             int64  `json:"ID"`
	Name           string `json:"Name"`
	URL            string `json:"URL"`
	ParentID       int64  `json:"ParentID"`
	Sorts          int64  `json:"Sorts"`
	Note           string `json:"Note"`
	CreateTime     int64  `json:"CreateTime"`
	LastUpdateTime int64  `json:"LastUpdateTime"`
}
