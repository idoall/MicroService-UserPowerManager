package handler

import (
	"context"

	"github.com/go-xorm/builder"
	"github.com/idoall/MicroService-UserPowerManager/srv/usersgroup/v1/models"
	proto "github.com/idoall/MicroService-UserPowerManager/srv/usersgroup/v1/proto"
	"github.com/idoall/MicroService-UserPowerManager/utils"
	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
	"github.com/idoall/MicroService-UserPowerManager/utils/jaeger"
	"github.com/micro/go-micro/errors"
)

// UsersGroup struct
type UsersGroup struct{}

// Add 添加一个用户组
func (e *UsersGroup) Add(ctx context.Context, req *proto.AddRequest, rep *proto.AddResponse) error {
	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_UsersGroup_Add_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_SRVUSERSGROUP

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s Service [UsersGroup][Add] request", namespaceID)
	}

	if req.Model.Name == "" {
		return errors.BadRequest(namespaceID, "Name 不能为空")
	}

	// 创建结构
	model := new(models.UsersGroup)
	model.Name = req.Model.Name
	model.ParentId = req.Model.ParentID
	model.Sorts = req.Model.Sorts
	model.Note = req.Model.Note

	// 添加用户组
	ctx, dbspan := jaeger.StartSpan(ctx, "Srv_UsersGroup_Add_WriteDB_Begin")
	if span != nil {
		defer dbspan.Finish()
	}
	if newID, err := model.Add(model); err != nil {
		return errors.BadRequest(namespaceID, "添加用户组失败:%s", err.Error())
	} else {

		// 设置返回值
		rep.NewID = newID
	}
	ctx, dbspan = jaeger.StartSpan(ctx, "Srv_UsersGroup_Add_WriteDB_End")
	if span != nil {
		defer dbspan.Finish()
	}

	// 写入一个 jaeger span
	_, span = jaeger.StartSpan(ctx, "Srv_UsersGroup_Add_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("NewID", rep.NewID)
	}

	return nil
}

// GetList 获取用户组列表,默认 id 倒排序
func (e *UsersGroup) GetList(ctx context.Context, req *proto.GetListRequest, rep *proto.GetListResponse) error {
	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_UsersGroup_GetList_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_SRVUSERSGROUP
	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s Service [UsersGroup][GetList] request", namespaceID)
	}

	// 判断请求参数
	if req.PageSize == 0 {
		return errors.BadRequest(namespaceID, "PageSize 不能为0")
	}

	if req.CurrentPageIndex == 0 {
		return errors.BadRequest(namespaceID, "CurrentPageIndex 不能0")
	}
	orderBy := "id desc"
	if req.OrderBy != "" {
		orderBy = req.OrderBy
	}

	var whereCond = builder.NewCond()

	for k, v := range req.Where {
		whereCond = whereCond.And(builder.Eq{k: v})
	}

	if list, totalcount, err := new(models.UsersGroup).GetAll(whereCond, orderBy, int(req.PageSize), int(req.CurrentPageIndex)); err != nil {
		return errors.BadRequest(namespaceID, "Model.UsersGroup GetAll Error:%s", err.Error())
	} else {
		rep.TotalCount = totalcount
		for _, v := range list {
			rep.List = append(rep.List, &proto.UsersGroup{
				ID:             v.Id,
				Name:           v.Name,
				Note:           v.Note,
				Sorts:          v.Sorts,
				ParentID:       v.ParentId,
				CreateTime:     v.CreateTime.Unix(),
				LastUpdateTime: v.LastUpdateTime.Unix(),
			})
		}
	}

	// 写入一个 jaeger span
	_, span = jaeger.StartSpan(ctx, "Srv_UsersGroup_GetList_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("PageSize", req.PageSize)
		span.SetTag("CurrentPageIndex", req.CurrentPageIndex)
		span.SetTag("orderBy", orderBy)
		span.SetTag("TotalCount", rep.TotalCount)
	}

	return nil
}

// Get 获取单个用户组
func (e *UsersGroup) Get(ctx context.Context, req *proto.GetRequest, rep *proto.GetResponse) error {

	var err error
	var modelUsersGroup models.UsersGroup

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_UsersGroup_GetUser_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_SRVUSERSGROUP

	// 判断请求参数
	if req.ID == 0 {
		return errors.BadRequest(namespaceID, "ID 没有赋值")
	}

	// 根据 id 获取一个用户组
	if modelUsersGroup, err = new(models.UsersGroup).GetOne(req.ID); err != nil {
		return errors.BadRequest(namespaceID, "models.UsersGroup QueryOne Error:%s", err.Error())
	}
	responseModel := &proto.UsersGroup{
		ID:             modelUsersGroup.Id,
		Name:           modelUsersGroup.Name,
		Sorts:          modelUsersGroup.Sorts,
		Note:           modelUsersGroup.Note,
		ParentID:       modelUsersGroup.ParentId,
		CreateTime:     modelUsersGroup.CreateTime.Unix(),
		LastUpdateTime: modelUsersGroup.LastUpdateTime.Unix(),
	}

	rep.Model = responseModel

	// 写入一个 jaeger span
	_, span = jaeger.StartSpan(ctx, "Srv_UsersGroup_GetUser_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("ID", req.ID)
	}

	return nil
}

// Update 修改用户组
func (e *UsersGroup) Update(ctx context.Context, req *proto.UpdateRequest, rep *proto.UpdateResponse) error {

	var err error
	var modelUsersGroup models.UsersGroup

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_UsersGroup_UpdateUser_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_SRVUSERSGROUP

	// 判断请求参数
	if req.Model.ID == 0 {
		return errors.BadRequest(namespaceID, "Id 不能为0")
	}

	// 根据 id 获取一个用户
	if modelUsersGroup, err = new(models.UsersGroup).GetOne(req.Model.ID); err != nil {
		return errors.BadRequest(namespaceID, "UpdateUsersGroup GetOne Error:%s", err.Error())
	}

	modelUsersGroup.Id = req.Model.ID
	modelUsersGroup.Name = req.Model.Name
	modelUsersGroup.Sorts = req.Model.Sorts
	modelUsersGroup.ParentId = req.Model.ParentID
	modelUsersGroup.Note = req.Model.Note

	// 修改用户
	if ok, err := modelUsersGroup.Update(modelUsersGroup); err != nil {
		return errors.BadRequest(namespaceID, "UpdateUser Update Error:%s", err.Error())
	} else {
		rep.Updated = ok
	}

	// 写入一个 jaeger span
	_, span = jaeger.StartSpan(ctx, "Srv_UsersGroup_UpdateUser_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("Updated", rep.Updated)
	}

	return nil
}

// BatchDelete 批量删除用户组
func (e *UsersGroup) BatchDelete(ctx context.Context, req *proto.DeleteRequest, rep *proto.DeleteResponse) error {

	var err error
	// var modelUser *models.Users

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_UsersGroup_BatchDelete_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_SRVUSERSGROUP

	// 判断请求参数
	if len(req.IDArray) == 0 {
		return errors.BadRequest(namespaceID, "IDArray 长度不能为0")
	}

	// 批量删除用户
	if _, err = new(models.UsersGroup).BatchDelete(req.IDArray); err != nil {
		return errors.BadRequest(namespaceID, "BatchDelete UsersGroup Error:%s", err.Error())
	} else {
		rep.Deleted = 1
	}

	// 写入一个 jaeger span
	_, span = jaeger.StartSpan(ctx, "Srv_UsersGroup_BatchDelete_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("UserIDArray", req.IDArray)
	}

	return nil
}
