package handler

import (
	"context"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/idoall/MicroService-UserPowerManager/models"
	proto "github.com/idoall/MicroService-UserPowerManager/srv/srvusersgroup/proto"
	"github.com/idoall/MicroService-UserPowerManager/utils"
	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
	"github.com/idoall/MicroService-UserPowerManager/utils/jaeger"
	"github.com/micro/go-micro/errors"
)

type SrvUsersGroup struct{}

// 添加一个用户组
func (e *SrvUsersGroup) Add(ctx context.Context, req *proto.AddRequest, rep *proto.AddResponse) error {
	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_UserGroup_Add_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_id := inner.NAMESPACE_MICROSERVICE_SRVUSERSGROUP

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s Service [Add] request", namespace_id)
	}

	if req.Model.Name == "" {
		return errors.BadRequest(namespace_id, "Name 不能为空")
	}

	// 创建结构
	model := new(models.UsersGroup)
	model.Name = req.Model.Name
	model.ParentId = int(req.Model.ParentId)
	model.Sorts = int(req.Model.Sorts)
	model.Note = req.Model.Note

	// 添加用户组
	ctx, dbspan := jaeger.StartSpan(ctx, "Srv_UserGroup_Add_WriteDB_Begin")
	if span != nil {
		defer dbspan.Finish()
	}
	if newID, err := model.Add(model); err != nil {
		return errors.BadRequest(namespace_id, "添加用户组失败:%s", err.Error())
	} else {

		// 设置返回值
		rep.NewId = newID
	}
	ctx, dbspan = jaeger.StartSpan(ctx, "Srv_UserGroup_Add_WriteDB_End")
	if span != nil {
		defer dbspan.Finish()
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Srv_UserGroup_Add_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("NewId", rep.NewId)
	}

	return nil
}

// 获取用户组列表,默认 id 倒排序
func (e *SrvUsersGroup) GetList(ctx context.Context, req *proto.GetListRequest, rep *proto.GetListResponse) error {
	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_UserGroup_GetList_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_id := inner.NAMESPACE_MICROSERVICE_SRVUSERSGROUP

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
	if list, totalcount, err := new(models.UsersGroup).GetAll(cond, int(req.PageSize), int(req.CurrentPageIndex), orderBy); err != nil {
		return errors.BadRequest(namespace_id, "Model.UsersGroup GetAll Error:%s", err.Error())
	} else {
		rep.TotalCount = totalcount
		for _, v := range list {
			rep.List = append(rep.List, &proto.UsersGroup{
				Id:             int64(v.Id),
				Name:           v.Name,
				Note:           v.Note,
				Sorts:          int64(v.Sorts),
				ParentId:       int64(v.ParentId),
				CreateTime:     v.CreateTime.Unix(),
				LastUpdateTime: v.LastUpdateTime.Unix(),
			})
		}
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Srv_UserGroup_GetList_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("PageSize", req.PageSize)
		span.SetTag("CurrentPageIndex", req.CurrentPageIndex)
		span.SetTag("orderBy", orderBy)
		span.SetTag("TotalCount", rep.TotalCount)
	}

	return nil
}

// 获取单个用户组
func (e *SrvUsersGroup) Get(ctx context.Context, req *proto.GetRequest, rep *proto.GetResponse) error {

	var err error
	var modelUserGroup *models.UsersGroup

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_UserGroup_GetUser_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_id := inner.NAMESPACE_MICROSERVICE_SRVUSERSGROUP

	// 判断请求参数
	if req.Id == 0 {
		return errors.BadRequest(namespace_id, "Id 没有赋值")
	}

	cond := orm.NewCondition().And("id", req.Id)

	// 根据 id 获取一个用户组
	if modelUserGroup, err = new(models.UsersGroup).QueryOne(cond, "-id"); err != nil {
		return errors.BadRequest(namespace_id, "models.UsersGroup QueryOne Error:%s", err.Error())
	}
	responseModel := &proto.UsersGroup{
		Id:             int64(modelUserGroup.Id),
		Name:           modelUserGroup.Name,
		Sorts:          int64(modelUserGroup.Sorts),
		Note:           modelUserGroup.Note,
		ParentId:       int64(modelUserGroup.ParentId),
		CreateTime:     modelUserGroup.CreateTime.Unix(),
		LastUpdateTime: modelUserGroup.LastUpdateTime.Unix(),
	}

	rep.Model = responseModel

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Srv_UserGroup_GetUser_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("Id", req.Id)
	}

	return nil
}

// 修改用户组
func (e *SrvUsersGroup) Update(ctx context.Context, req *proto.UpdateRequest, rep *proto.UpdateResponse) error {

	var err error
	var modelUserGroup *models.UsersGroup

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_UserGroup_UpdateUser_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_id := inner.NAMESPACE_MICROSERVICE_SRVUSERSGROUP

	// 判断请求参数
	if req.Model.Id == 0 {
		return errors.BadRequest(namespace_id, "Id 不能为0")
	}

	// 根据 id 获取一个用户
	if modelUserGroup, err = new(models.UsersGroup).GetOne(req.Model.Id); err != nil {
		return errors.BadRequest(namespace_id, "UpdateUserGroup GetOne Error:%s", err.Error())
	}

	modelUserGroup.Id = int(req.Model.Id)
	modelUserGroup.Name = req.Model.Name
	modelUserGroup.Sorts = int(req.Model.Sorts)
	modelUserGroup.ParentId = int(req.Model.ParentId)
	modelUserGroup.Note = req.Model.Note

	// 修改用户
	if ok, err := modelUserGroup.Update(modelUserGroup); err != nil {
		return errors.BadRequest(namespace_id, "UpdateUser Update Error:%s", err.Error())
	} else {
		rep.Updated = ok
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Srv_UserGroup_UpdateUser_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("Updated", rep.Updated)
	}

	return nil
}

// 批量删除用户组
func (e *SrvUsersGroup) BatchDelete(ctx context.Context, req *proto.DeleteRequest, rep *proto.DeleteResponse) error {

	var err error
	// var modelUser *models.Users

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_UserGroup_BatchDelete_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_id := inner.NAMESPACE_MICROSERVICE_SRVUSERSGROUP

	// 判断请求参数
	if len(req.IdArray) == 0 {
		return errors.BadRequest(namespace_id, "IdArray 长度不能为0")
	}

	// 批量删除用户
	if _, err = new(models.Users).BatchDelete(req.IdArray); err != nil {
		return errors.BadRequest(namespace_id, "BatchDeleteUserGroup BatchDelete Error:%s", err.Error())
	} else {
		rep.Deleted = 1
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Srv_UserGroup_BatchDelete_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("UserIdArray", strings.Join(req.IdArray, ","))
	}

	return nil
}
