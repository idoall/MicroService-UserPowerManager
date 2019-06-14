package handler

import (
	"context"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/idoall/MicroService-UserPowerManager/models"
	proto "github.com/idoall/MicroService-UserPowerManager/srv/columns/v1/proto"
	"github.com/idoall/MicroService-UserPowerManager/utils"
	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
	"github.com/idoall/MicroService-UserPowerManager/utils/jaeger"
	"github.com/micro/go-micro/errors"
)

type Columns struct{}

// 添加一条记录
func (e *Columns) Add(ctx context.Context, req *proto.AddRequest, rep *proto.AddResponse) error {
	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_Columns_Add_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_SRVCOLUMNS

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s Service [SrvColumns][Add] request", namespaceID)
	}

	if req.Model.Name == "" {
		return errors.BadRequest(namespaceID, "Name 不能为空")
	}

	// 创建结构
	model := new(models.Columns)
	model.Name = req.Model.Name
	model.URL = req.Model.URL
	model.ParentID = req.Model.ParentID
	model.Sorts = req.Model.Sorts
	model.IsShowNav = req.Model.IsShowNav
	model.CssIcon = req.Model.CssIcon

	// 添加到数据库
	ctx, dbspan := jaeger.StartSpan(ctx, "Srv_Columns_Add_WriteDB_Begin")
	if span != nil {
		defer dbspan.Finish()
	}
	if newID, err := model.Add(model); err != nil {
		return errors.BadRequest(namespaceID, "添加到数据库失败:%s", err.Error())
	} else {

		// 设置返回值
		rep.NewID = newID
	}
	ctx, dbspan = jaeger.StartSpan(ctx, "Srv_Columns_Add_WriteDB_End")
	if span != nil {
		defer dbspan.Finish()
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Srv_Columns_Add_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("NewID", rep.NewID)
	}

	return nil
}

// 获取列表,默认 id 倒排序
func (e *Columns) GetList(ctx context.Context, req *proto.GetListRequest, rep *proto.GetListResponse) error {
	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_Columns_GetList_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_SRVCOLUMNS

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s Service [SrvColumns][GetList] request", namespaceID)
	}

	// 判断请求参数
	if req.PageSize == 0 {
		return errors.BadRequest(namespaceID, "PageSize 不能为0")
	}

	if req.CurrentPageIndex == 0 {
		return errors.BadRequest(namespaceID, "CurrentPageIndex 不能0")
	}
	orderBy := "-sorts"
	if req.OrderBy != "" {
		orderBy = req.OrderBy
	}

	cond := orm.NewCondition()
	if list, totalcount, err := new(models.Columns).GetAll(cond, int(req.PageSize), int(req.CurrentPageIndex), orderBy); err != nil {
		return errors.BadRequest(namespaceID, "Model.Columns GetAll Error:%s", err.Error())
	} else {
		rep.TotalCount = totalcount
		for _, v := range list {
			rep.List = append(rep.List, &proto.Columns{
				ID:             v.ID,
				Name:           v.Name,
				URL:            v.URL,
				Sorts:          v.Sorts,
				ParentID:       v.ParentID,
				IsShowNav:      v.IsShowNav,
				CssIcon:        v.CssIcon,
				CreateTime:     v.CreateTime.Unix(),
				LastUpdateTime: v.LastUpdateTime.Unix(),
			})
		}
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Srv_Columns_GetList_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("PageSize", req.PageSize)
		span.SetTag("CurrentPageIndex", req.CurrentPageIndex)
		span.SetTag("orderBy", orderBy)
		span.SetTag("TotalCount", rep.TotalCount)
	}

	return nil
}

// 获取单条记录
func (e *Columns) Get(ctx context.Context, req *proto.GetRequest, rep *proto.GetResponse) error {

	var err error
	var modelGet *models.Columns

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_Columns_GetUser_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_SRVCOLUMNS

	// 判断请求参数
	if req.ID == 0 {
		return errors.BadRequest(namespaceID, "Id 没有赋值")
	}

	cond := orm.NewCondition().And("id", req.ID)

	// 根据 id 获取一个栏目
	if modelGet, err = new(models.Columns).QueryOne(cond, "-id"); err != nil {
		return errors.BadRequest(namespaceID, "models.Columns QueryOne Error:%s", err.Error())
	}
	responseModel := &proto.Columns{
		ID:             int64(modelGet.ID),
		Name:           modelGet.Name,
		URL:            modelGet.URL,
		Sorts:          int64(modelGet.Sorts),
		ParentID:       int64(modelGet.ParentID),
		IsShowNav:      modelGet.IsShowNav,
		CssIcon:        modelGet.CssIcon,
		CreateTime:     modelGet.CreateTime.Unix(),
		LastUpdateTime: modelGet.LastUpdateTime.Unix(),
	}

	rep.Model = responseModel

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Srv_Columns_GetUser_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("ID", req.ID)
	}

	return nil
}

// 修改
func (e *Columns) Update(ctx context.Context, req *proto.UpdateRequest, rep *proto.UpdateResponse) error {

	var err error
	var modelGet *models.Columns

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_Columns_Update_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_SRVCOLUMNS

	// 判断请求参数
	if req.Model.ID == 0 {
		return errors.BadRequest(namespaceID, "ID 不能为0")
	}

	// 根据 id 获取一个栏目
	if modelGet, err = new(models.Columns).GetOne(req.Model.ID); err != nil {
		return errors.BadRequest(namespaceID, "Update Columns GetOne Error:%s", err.Error())
	}

	modelGet.ID = req.Model.ID
	modelGet.Name = req.Model.Name
	modelGet.Sorts = req.Model.Sorts
	modelGet.ParentID = req.Model.ParentID
	modelGet.URL = req.Model.URL
	modelGet.IsShowNav = req.Model.IsShowNav
	modelGet.CssIcon = req.Model.CssIcon

	// 修改
	if ok, err := modelGet.Update(modelGet); err != nil {
		return errors.BadRequest(namespaceID, "Update Columns Update Error:%s", err.Error())
	} else {
		rep.Updated = ok
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Srv_Columns_Update_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("ID", modelGet.ID)
		span.SetTag("Name", modelGet.Name)
		span.SetTag("Updated", rep.Updated)
	}

	return nil
}

// 批量删除
func (e *Columns) BatchDelete(ctx context.Context, req *proto.DeleteRequest, rep *proto.DeleteResponse) error {

	var err error
	// var modelUser *models.Users

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_Columns_BatchDelete_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_SRVCOLUMNS

	// 判断请求参数
	if len(req.IDArray) == 0 {
		return errors.BadRequest(namespaceID, "IDArray 长度不能为0")
	}

	// 批量删除
	if _, err = new(models.Columns).BatchDelete(req.IDArray); err != nil {
		return errors.BadRequest(namespaceID, "BatchDelete Columns Error:%s", err.Error())
	} else {
		rep.Deleted = 1
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Srv_Columns_BatchDelete_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("IDArray", strings.Join(req.IDArray, ","))
	}

	return nil
}
