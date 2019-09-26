package handler

import (
	"context"
	"fmt"

	"github.com/go-xorm/builder"
	"github.com/idoall/MicroService-UserPowerManager/srv/users/v1/models"
	proto "github.com/idoall/MicroService-UserPowerManager/srv/users/v1/proto"
	"github.com/idoall/MicroService-UserPowerManager/utils"
	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
	"github.com/idoall/MicroService-UserPowerManager/utils/jaeger"
	"github.com/idoall/TokenExchangeCommon/commonutils"
	"github.com/idoall/TokenExchangeCommon/commonutils/checkmail"
	"github.com/micro/go-micro/errors"
)

// Users Struct
type Users struct{}

// Add 添加一个用户
func (e *Users) Add(ctx context.Context, req *proto.AddRequest, rep *proto.AddResponse) error {
	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_User_Add_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_SRVUSERS

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s Srv [Users][Add] request", namespaceID)
	}

	if req.Model.UserName == "" {
		return errors.BadRequest(namespaceID, "UserName 不能为空")
	}

	if req.Model.Password == "" {
		return errors.BadRequest(namespaceID, "Password 不能为空")
	}

	if req.Model.Email == "" {
		return errors.BadRequest(namespaceID, "Email 不能为空")
	} else if err := checkmail.ValidateFormat(req.Model.Email); err != nil {
		return errors.BadRequest(namespaceID, "Email 的格式不正确:%s", err.Error())
	}

	// 创建用户结构
	model := new(models.Users)
	model.UserName = req.Model.UserName
	model.RealyName = req.Model.RealyName
	model.Password = req.Model.Password
	model.Email = req.Model.Email
	model.AuthKey = req.Model.AuthKey
	model.IsDel = int(req.Model.IsDel)
	model.ParentId = req.Model.ParentID
	model.Note = req.Model.Note

	// 添加用户
	ctx, dbspan := jaeger.StartSpan(ctx, "Srv_User_Add_WriteDB_Begin")
	if span != nil {
		defer dbspan.Finish()
	}
	if newID, err := model.Add(model); err != nil {
		return errors.BadRequest(namespaceID, "添加用户失败:%s", err.Error())
	} else {

		// 设置返回值
		rep.NewID = newID
	}

	ctx, dbspan = jaeger.StartSpan(ctx, "Srv_User_Add_WriteDB_End")
	if span != nil {
		defer dbspan.Finish()
	}

	// 写入一个 jaeger span
	_, span = jaeger.StartSpan(ctx, "Srv_User_Add_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("NewID", rep.NewID)
	}

	return nil
}

// GetList 获取用户列表,默认 id 倒排序
func (e *Users) GetList(ctx context.Context, req *proto.GetListRequest, rep *proto.GetListResponse) error {
	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_User_GetList_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_SRVUSERS

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s Srv [Users][GetList] request", namespaceID)
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

	userListCond := builder.Eq{"is_del": 0}
	if list, totalcount, err := new(models.Users).GetAll(userListCond, orderBy, int(req.PageSize), int(req.CurrentPageIndex)); err != nil {
		return errors.BadRequest(namespaceID, "Model.Users GetAll Error:%s", err.Error())
	} else {

		rep.TotalCount = totalcount
		for _, v := range list {
			rep.List = append(rep.List, &proto.User{
				ID:             v.Id,
				UserName:       v.UserName,
				RealyName:      v.RealyName,
				Password:       v.Password,
				AuthKey:        v.AuthKey,
				Email:          v.Email,
				Note:           v.Note,
				ParentID:       v.ParentId,
				CreateTime:     v.CreateTime.Unix(),
				LastUpdateTime: v.LastUpdateTime.Unix(),
			})
		}
	}

	// 写入一个 jaeger span
	_, span = jaeger.StartSpan(ctx, "Srv_User_GetList_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("PageSize", req.PageSize)
		span.SetTag("CurrentPageIndex", req.CurrentPageIndex)
		span.SetTag("orderBy", orderBy)
		span.SetTag("TotalCount", rep.TotalCount)
	}

	return nil
}

// Get 获取单个用户
func (e *Users) Get(ctx context.Context, req *proto.GetRequest, rep *proto.GetResponse) error {

	var err error
	var modelUser models.Users

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_User_GetUser_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_SRVUSERS

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s Srv [Users][Get] request", namespaceID)
	}

	// 判断请求参数
	if req.ID == 0 && req.UserName == "" {
		return errors.BadRequest(namespaceID, "UserId 和 UserName 没有赋值")
	}

	var whereCond = builder.NewCond()

	if req.ID != 0 {
		whereCond = whereCond.And(builder.Eq{"id": req.ID})
	}

	if req.UserName != "" {
		whereCond = whereCond.And(builder.Eq{"user_name": req.UserName})
	}

	// 根据 id 获取一个用户
	if modelUser, err = new(models.Users).QueryOne(whereCond, "id desc"); err != nil {
		return errors.BadRequest(namespaceID, "models.Users GetOne Error:%s", err.Error())
	}
	responseModel := &proto.User{
		ID:             modelUser.Id,
		UserName:       modelUser.UserName,
		RealyName:      modelUser.RealyName,
		Password:       modelUser.Password,
		Email:          modelUser.Email,
		AuthKey:        modelUser.AuthKey,
		Note:           modelUser.Note,
		ParentID:       modelUser.ParentId,
		CreateTime:     modelUser.CreateTime.Unix(),
		LastUpdateTime: modelUser.LastUpdateTime.Unix(),
	}

	rep.Model = responseModel

	fmt.Printf("%+v \n", modelUser)

	// 写入一个 jaeger span
	_, span = jaeger.StartSpan(ctx, "Srv_User_GetUser_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("ID", req.ID)
		span.SetTag("UserName", req.UserName)
	}

	return nil
}

// Update 修改用户
func (e *Users) Update(ctx context.Context, req *proto.UpdateRequest, rep *proto.UpdateResponse) error {

	var err error
	var modelUser models.Users

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_User_UpdateUser_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_SRVUSERS

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s Srv [Users][Update] request", namespaceID)
	}

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s Service [SrvUsers][Update] request", namespaceID)
	}

	// 判断请求参数
	if req.Model.ID == 0 {
		return errors.BadRequest(namespaceID, "User.ID 不能为0")
	}

	// 根据 id 获取一个用户
	if modelUser, err = new(models.Users).GetOne(req.Model.ID); err != nil {
		return errors.BadRequest(namespaceID, "UpdateUser models.Users GetOne Error:%s", err.Error())
	}

	modelUser.Id = req.Model.ID
	modelUser.UserName = req.Model.UserName
	modelUser.RealyName = req.Model.RealyName
	modelUser.AuthKey = req.Model.AuthKey
	modelUser.Password = req.Model.Password
	modelUser.Email = req.Model.Email
	modelUser.ParentId = req.Model.ParentID
	modelUser.Note = req.Model.Note
	modelUser.IsDel = int(req.Model.IsDel)
	modelUser.CreateTime = commonutils.TimeFromUnixEscInt64(req.Model.CreateTime)

	// 修改用户
	if ok, err := modelUser.Update(modelUser); err != nil {
		return errors.BadRequest(namespaceID, "UpdateUser models.Users Update Error:%s", err.Error())
	} else {
		rep.Updated = ok
	}

	// 写入一个 jaeger span
	_, span = jaeger.StartSpan(ctx, "Srv_User_UpdateUser_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("Updated", rep.Updated)
	}

	return nil
}

// BatchDelete 批量删除用户（假删除）
func (e *Users) BatchDelete(ctx context.Context, req *proto.DeleteRequest, rep *proto.DeleteResponse) error {

	var err error
	// var modelUser *models.Users

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_User_BatchDelete_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_SRVUSERS

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s Srv [Users][BatchDelete] request", namespaceID)
	}

	// 判断请求参数
	if len(req.IDArray) == 0 {
		return errors.BadRequest(namespaceID, "IDArray 长度不能为0")
	}

	// 批量删除用户
	if rep.Deleted, err = new(models.Users).BatchDelete(req.IDArray); err != nil {
		return errors.BadRequest(namespaceID, "BatchDeleteUser Users Error:%s", err.Error())
	}

	// 写入一个 jaeger span
	_, span = jaeger.StartSpan(ctx, "Srv_User_BatchDelete_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("IDArray", req.IDArray)
	}

	return nil
}
