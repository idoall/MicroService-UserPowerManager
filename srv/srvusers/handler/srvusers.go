package handler

import (
	"context"
	"strings"

	"github.com/idoall/TokenExchangeCommon/commonutils"

	"github.com/astaxie/beego/orm"
	"github.com/idoall/MicroService-UserPowerManager/models"
	proto "github.com/idoall/MicroService-UserPowerManager/srv/srvusers/proto"
	"github.com/idoall/MicroService-UserPowerManager/utils"
	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
	"github.com/idoall/MicroService-UserPowerManager/utils/jaeger"
	"github.com/idoall/TokenExchangeCommon/commonutils/checkmail"
	"github.com/micro/go-micro/errors"
)

type SrvUsers struct{}

// 添加一个用户
func (e *SrvUsers) Add(ctx context.Context, req *proto.AddRequest, rep *proto.AddResponse) error {
	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_User_Add_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_id := inner.NAMESPACE_MICROSERVICE_SRVUSERS

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s Service [Add] request", namespace_id)
	}

	if req.Model.UserName == "" {
		return errors.BadRequest(namespace_id, "UserName 不能为空")
	}

	if req.Model.Password == "" {
		return errors.BadRequest(namespace_id, "Password 不能为空")
	}

	if req.Model.Email == "" {
		return errors.BadRequest(namespace_id, "Email 不能为空")
	} else if err := checkmail.ValidateFormat(req.Model.Email); err != nil {
		return errors.BadRequest(namespace_id, "Email 的格式不正确:%s", err.Error())
	}

	// 创建用户结构
	model := new(models.Users)
	model.UserName = req.Model.UserName
	model.RealyName = req.Model.RealyName
	model.Password = req.Model.Password
	model.Email = req.Model.Email
	model.AuthKey = req.Model.AuthKey
	model.IsDel = req.Model.IsDel
	model.ParentId = int(req.Model.ParentId)
	model.Note = req.Model.Note

	// 添加用户
	ctx, dbspan := jaeger.StartSpan(ctx, "Srv_User_Add_WriteDB_Begin")
	if span != nil {
		defer dbspan.Finish()
	}
	if newID, err := model.Add(model); err != nil {
		return errors.BadRequest(namespace_id, "添加用户失败:%s", err.Error())
	} else {

		// 设置返回值
		rep.NewUserId = newID
	}
	ctx, dbspan = jaeger.StartSpan(ctx, "Srv_User_Add_WriteDB_End")
	if span != nil {
		defer dbspan.Finish()
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Srv_User_Add_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("NewUserId", rep.NewUserId)
	}

	return nil
}

// 获取用户列表,默认 id 倒排序
func (e *SrvUsers) GetList(ctx context.Context, req *proto.GetListRequest, rep *proto.GetListResponse) error {
	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_User_GetList_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_id := inner.NAMESPACE_MICROSERVICE_SRVUSERS

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

	userListCond := orm.NewCondition().And("isdel", 0)
	if list, totalcount, err := new(models.Users).GetAll(userListCond, int(req.PageSize), int(req.CurrentPageIndex), orderBy); err != nil {
		return errors.BadRequest(namespace_id, "Model.Users GetAll Error:%s", err.Error())
	} else {
		rep.TotalCount = totalcount
		for _, v := range list {
			rep.List = append(rep.List, &proto.User{
				Id:             int64(v.Id),
				UserName:       v.UserName,
				RealyName:      v.RealyName,
				Password:       v.Password,
				AuthKey:        v.AuthKey,
				Email:          v.Email,
				Note:           v.Note,
				ParentId:       int64(v.ParentId),
				CreateTime:     v.CreateTime.Unix(),
				LastUpdateTime: v.LastUpdateTime.Unix(),
			})
		}
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Srv_User_GetList_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("PageSize", req.PageSize)
		span.SetTag("CurrentPageIndex", req.CurrentPageIndex)
		span.SetTag("orderBy", orderBy)
		span.SetTag("TotalCount", rep.TotalCount)
	}

	return nil
}

// 获取单个用户
func (e *SrvUsers) Get(ctx context.Context, req *proto.GetRequest, rep *proto.GetResponse) error {

	var err error
	var modelUser *models.Users

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_User_GetUser_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_id := inner.NAMESPACE_MICROSERVICE_SRVUSERS

	// 判断请求参数
	if req.UserId == 0 && req.UserName == "" {
		return errors.BadRequest(namespace_id, "UserId 和 UserName 没有赋值")
	}

	cond := orm.NewCondition()
	if req.UserId != 0 {
		cond = cond.And("id", req.UserId)
	}

	if req.UserName != "" {
		cond = cond.And("user_name", req.UserName)
	}

	// 根据 id 获取一个用户
	if modelUser, err = new(models.Users).QueryOne(cond, "-id"); err != nil {
		return errors.BadRequest(namespace_id, "models.Users GetOne Error:%s", err.Error())
	}
	responseModel := &proto.User{
		Id:             int64(modelUser.Id),
		UserName:       modelUser.UserName,
		RealyName:      modelUser.RealyName,
		Password:       modelUser.Password,
		Email:          modelUser.Email,
		AuthKey:        modelUser.AuthKey,
		Note:           modelUser.Note,
		ParentId:       int64(modelUser.ParentId),
		CreateTime:     modelUser.CreateTime.Unix(),
		LastUpdateTime: modelUser.LastUpdateTime.Unix(),
	}

	rep.Model = responseModel

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Srv_User_GetUser_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("UserId", req.UserId)
	}

	return nil
}

// 修改用户
func (e *SrvUsers) Update(ctx context.Context, req *proto.UpdateRequest, rep *proto.UpdateResponse) error {

	var err error
	var modelUser *models.Users

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_User_UpdateUser_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_id := inner.NAMESPACE_MICROSERVICE_SRVUSERS

	// 判断请求参数
	if req.Model.Id == 0 {
		return errors.BadRequest(namespace_id, "User.Id 不能为0")
	}

	// 根据 id 获取一个用户
	if modelUser, err = new(models.Users).GetOne(req.Model.Id); err != nil {
		return errors.BadRequest(namespace_id, "UpdateUser models.Users GetOne Error:%s", err.Error())
	}

	modelUser.Id = int(req.Model.Id)
	modelUser.UserName = req.Model.UserName
	modelUser.RealyName = req.Model.RealyName
	modelUser.Password = req.Model.Password
	modelUser.Email = req.Model.Email
	modelUser.ParentId = int(req.Model.ParentId)
	modelUser.Note = req.Model.Note
	modelUser.IsDel = req.Model.IsDel
	modelUser.ParentId = int(req.Model.ParentId)
	modelUser.CreateTime = commonutils.TimeFromUnixEscInt64(req.Model.CreateTime)

	// 修改用户
	if ok, err := modelUser.Update(modelUser); err != nil {
		return errors.BadRequest(namespace_id, "UpdateUser models.Users Update Error:%s", err.Error())
	} else {
		rep.Updated = ok
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Srv_User_UpdateUser_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("Updated", rep.Updated)
	}

	return nil
}

// 批量删除用户（假删除）
func (e *SrvUsers) BatchDelete(ctx context.Context, req *proto.DeleteRequest, rep *proto.DeleteResponse) error {

	var err error
	// var modelUser *models.Users

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_User_BatchDelete_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_id := inner.NAMESPACE_MICROSERVICE_SRVUSERS

	// 判断请求参数
	if len(req.UserIdArray) == 0 {
		return errors.BadRequest(namespace_id, "User.Id 长度不能为0")
	}

	// 批量删除用户
	if _, err = new(models.Users).BatchDelete(req.UserIdArray); err != nil {
		return errors.BadRequest(namespace_id, "BatchDeleteUser models.Users BatchDelete Error:%s", err.Error())
	} else {
		rep.Deleted = 1
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Srv_User_BatchDelete_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("UserIdArray", strings.Join(req.UserIdArray, ","))
	}

	return nil
}
