package main

import (
	"time"

	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
	"github.com/micro/go-log"

	"github.com/micro/go-micro"

	"github.com/idoall/MicroService-UserPowerManager/api/columns/v1/handler"
	srvProto "github.com/idoall/MicroService-UserPowerManager/srv/columns/v1/proto"
)

func main() {

	var err error

	// 名称  say 一定要和在proto中定义的一样
	service := micro.NewService(
		micro.Name(inner.NAMESPACE_MICROSERVICE_APICOLUMNS),
		micro.RegisterTTL(time.Second*30),      // 注册服务的过期时间
		micro.RegisterInterval(time.Second*20), //间隔多久再次注册服务
	)

	// parse command line flags
	service.Init()

	if err = service.Server().Handle(
		service.Server().NewHandler(
			&handler.Columns{
				Client: srvProto.NewProtoColumnsService(inner.NAMESPACE_MICROSERVICE_SRVCOLUMNS, service.Client()),
			},
		),
	); err != nil {
		log.Fatal(err)
	}

	if err = service.Run(); err != nil {
		log.Fatal(err)
	}
}
