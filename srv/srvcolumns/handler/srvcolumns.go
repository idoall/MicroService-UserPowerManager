package handler

import (
	"context"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/idoall/MicroService-UserPowerManager/models"
	proto "github.com/idoall/MicroService-UserPowerManager/srv/srvcolumns/proto"
	"github.com/idoall/MicroService-UserPowerManager/utils"
	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
	"github.com/idoall/MicroService-UserPowerManager/utils/jaeger"
	"github.com/micro/go-micro/errors"
)

type SrvColumns struct{}

// 添加一条记录
func (e *SrvColumns) Add(ctx context.Context, req *proto.AddRequest, rep *proto.AddResponse) error {
	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_Columns_Add_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_id := inner.NAMESPACE_MICROSERVICE_SRVCOLUMNS

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s Service [Add] request", namespace_id)
	}

	if req.Model.Name == "" {
		return errors.BadRequest(namespace_id, "Name 不能为空")
	}

	// 创建结构
	model := new(models.Columns)
	model.Name = req.Model.Name
	model.URL = req.Model.URL
	model.ParentId = int(req.Model.ParentId)
	model.Sorts = int(req.Model.Sorts)
	model.IsShowNav = req.Model.IsShowNav
	model.CssIcon = req.Model.CssIceo

	// 添加到数据库
	ctx, dbspan := jaeger.StartSpan(ctx, "Srv_Columns_Add_WriteDB_Begin")
	if span != nil {
		defer dbspan.Finish()
	}
	if newID, err := model.Add(model); err != nil {
		return errors.BadRequest(namespace_id, "添加到数据库失败:%s", err.Error())
	} else {

		// 设置返回值
		rep.NewId = newID
	}
	ctx, dbspan = jaeger.StartSpan(ctx, "Srv_Columns_Add_WriteDB_End")
	if span != nil {
		defer dbspan.Finish()
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Srv_Columns_Add_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("NewId", rep.NewId)
	}

	return nil
}

// 获取列表,默认 id 倒排序
func (e *SrvColumns) GetList(ctx context.Context, req *proto.GetListRequest, rep *proto.GetListResponse) error {
	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_Columns_GetList_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_id := inner.NAMESPACE_MICROSERVICE_SRVCOLUMNS

	// 判断请求参数
	if req.PageSize == 0 {
		return errors.BadRequest(namespace_id, "PageSize 不能为0")
	}

	if req.CurrentPageIndex == 0 {
		return errors.BadRequest(namespace_id, "CurrentPageIndex 不能0")
	}
	orderBy := "-id"
	if req.OrderBy != "" {
		orderBy = req.OrderBy
	}

	cond := orm.NewCondition()
	if list, totalcount, err := new(models.Columns).GetAll(cond, int(req.PageSize), int(req.CurrentPageIndex), orderBy); err != nil {
		return errors.BadRequest(namespace_id, "Model.Columns GetAll Error:%s", err.Error())
	} else {
		rep.TotalCount = totalcount
		for _, v := range list {
			rep.List = append(rep.List, &proto.Columns{
				Id:             int64(v.Id),
				Name:           v.Name,
				URL:            v.URL,
				Sorts:          int64(v.Sorts),
				ParentId:       int64(v.ParentId),
				IsShowNav:      v.IsShowNav,
				CssIceo:        v.CssIcon,
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
func (e *SrvColumns) Get(ctx context.Context, req *proto.GetRequest, rep *proto.GetResponse) error {

	var err error
	var modelGet *models.Columns

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_Columns_GetUser_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_id := inner.NAMESPACE_MICROSERVICE_SRVCOLUMNS

	// 判断请求参数
	if req.Id == 0 {
		return errors.BadRequest(namespace_id, "Id 没有赋值")
	}

	cond := orm.NewCondition().And("id", req.Id)

	// 根据 id 获取一个栏目
	if modelGet, err = new(models.Columns).QueryOne(cond, "-id"); err != nil {
		return errors.BadRequest(namespace_id, "models.Columns QueryOne Error:%s", err.Error())
	}
	responseModel := &proto.Columns{
		Id:             int64(modelGet.Id),
		Name:           modelGet.Name,
		URL:            modelGet.URL,
		Sorts:          int64(modelGet.Sorts),
		ParentId:       int64(modelGet.ParentId),
		IsShowNav:      modelGet.IsShowNav,
		CssIceo:        modelGet.CssIcon,
		CreateTime:     modelGet.CreateTime.Unix(),
		LastUpdateTime: modelGet.LastUpdateTime.Unix(),
	}

	rep.Model = responseModel

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Srv_Columns_GetUser_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("Id", req.Id)
	}

	return nil
}

// 修改
func (e *SrvColumns) Update(ctx context.Context, req *proto.UpdateRequest, rep *proto.UpdateResponse) error {

	var err error
	var modelGet *models.Columns

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_Columns_Update_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_id := inner.NAMESPACE_MICROSERVICE_SRVCOLUMNS

	// 判断请求参数
	if req.Model.Id == 0 {
		return errors.BadRequest(namespace_id, "Id 不能为0")
	}

	// 根据 id 获取一个栏目
	if modelGet, err = new(models.Columns).GetOne(req.Model.Id); err != nil {
		return errors.BadRequest(namespace_id, "Update Columns GetOne Error:%s", err.Error())
	}

	modelGet.Id = int(req.Model.Id)
	modelGet.Name = req.Model.Name
	modelGet.Sorts = int(req.Model.Sorts)
	modelGet.ParentId = int(req.Model.ParentId)
	modelGet.URL = req.Model.URL
	modelGet.IsShowNav = req.Model.IsShowNav
	modelGet.CssIcon = req.Model.CssIceo

	// 修改
	if ok, err := modelGet.Update(modelGet); err != nil {
		return errors.BadRequest(namespace_id, "Update Columns Update Error:%s", err.Error())
	} else {
		rep.Updated = ok
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Srv_Columns_Update_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("Id", modelGet.Id)
		span.SetTag("Name", modelGet.Name)
		span.SetTag("Updated", rep.Updated)
	}

	return nil
}

// 批量删除
func (e *SrvColumns) BatchDelete(ctx context.Context, req *proto.DeleteRequest, rep *proto.DeleteResponse) error {

	var err error
	// var modelUser *models.Users

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_Columns_BatchDelete_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_id := inner.NAMESPACE_MICROSERVICE_SRVCOLUMNS

	// 判断请求参数
	if len(req.IdArray) == 0 {
		return errors.BadRequest(namespace_id, "IdArray 长度不能为0")
	}

	// 批量删除
	if _, err = new(models.Columns).BatchDelete(req.IdArray); err != nil {
		return errors.BadRequest(namespace_id, "BatchDelete Columns Error:%s", err.Error())
	} else {
		rep.Deleted = 1
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Srv_Columns_BatchDelete_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("UserIdArray", strings.Join(req.IdArray, ","))
	}

	return nil
}
