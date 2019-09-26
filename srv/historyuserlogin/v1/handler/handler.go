package handler

import (
	"context"

	"github.com/idoall/MicroService-UserPowerManager/srv/historyuserlogin/v1/models"
	proto "github.com/idoall/MicroService-UserPowerManager/srv/historyuserlogin/v1/proto"
	"github.com/idoall/MicroService-UserPowerManager/utils"
	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
	"github.com/idoall/MicroService-UserPowerManager/utils/jaeger"
	"github.com/micro/go-micro/errors"
)

type HistoryUserLogin struct{}

// 添加一条记录
func (e *HistoryUserLogin) Add(ctx context.Context, req *proto.AddRequest, rep *proto.AddResponse) error {
	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Srv_HistoryUserLogin_Add_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_id := inner.NAMESPACE_MICROSERVICE_SRVHISTORYUSERLOGIN

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s Service [Add] request", namespace_id)
	}

	if req.User.ID == 0 {
		return errors.BadRequest(namespace_id, "User.ID 不能为0")
	}

	// if req.GeoRemoteAddr == "" {
	// 	return errors.BadRequest(namespace_id, "GeoRemoteAddr 不能为空")
	// }

	// if req.GeoCountry == "" {
	// 	return errors.BadRequest(namespace_id, "GeoCountry 不能为空")
	// }

	// if req.GeoCity == "" {
	// 	return errors.BadRequest(namespace_id, "GeoCity 不能为空")
	// }

	// if req.DeviceDetector == "" {
	// 	return errors.BadRequest(namespace_id, "DeviceDetector 不能为空")
	// }

	// 创建数据库记录
	model := models.HistoryUserLogin{}
	model.UserId = req.User.ID
	model.GeoRemoteAddr = req.GeoRemoteAddr
	model.GeoCountry = req.GeoCountry
	model.GeoCity = req.GeoCity
	model.DeviceDetector = req.DeviceDetector

	// 添加用户
	ctx, dbspan := jaeger.StartSpan(ctx, "Srv_HistoryUserLogin_Add_WriteDB_Begin")
	if span != nil {
		defer dbspan.Finish()
	}
	if _, err := model.Add(model); err != nil {
		return errors.BadRequest(namespace_id, "添加失败:%s", err.Error())
	}
	// 设置返回值
	rep.NewID = 1

	ctx, dbspan = jaeger.StartSpan(ctx, "Srv_HistoryUserLogin_Add_WriteDB_End")
	if span != nil {
		defer dbspan.Finish()
	}

	// 写入一个 jaeger span
	_, span = jaeger.StartSpan(ctx, "Srv_HistoryUserLogin_Add_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("NewId", rep.NewID)
		span.SetTag("GeoRemoteAddr", req.GeoRemoteAddr)
		span.SetTag("GeoCountry", req.GeoCountry)
		span.SetTag("GeoCity", req.GeoCity)
	}

	return nil
}
