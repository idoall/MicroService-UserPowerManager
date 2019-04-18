package main

import (
	"github.com/idoall/MicroService-UserPowerManager/utils"
	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
	"github.com/idoall/MicroService-UserPowerManager/utils/jaeger"
	"github.com/micro/go-log"

	"github.com/micro/go-micro"

	"github.com/idoall/MicroService-UserPowerManager/api/apiusers/handler"
	srvhistoryuserlogin "github.com/idoall/MicroService-UserPowerManager/srv/srvhistoryuserlogin/proto"
	srvusers "github.com/idoall/MicroService-UserPowerManager/srv/srvusers/proto"
)

func main() {

	// 是否开启 Jaeger
	if enableJaeger, err := utils.TConfig.Bool("JaegerServices::Enable"); err != nil {
		inner.Mlogger.Fatalf("utils.TConfig Jaeger:%s", err)
	} else if enableJaeger {
		// 初始化 Jaeger 配置
		_, closer := jaeger.InitJaeger("UserPowerManager")
		defer closer.Close()
	}

	// 名称  say 一定要和在proto中定义的一样
	service := micro.NewService(
		micro.Name(inner.NAMESPACE_MICROSERVICE_API),
	)

	// parse command line flags
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&handler.ApiUsers{
				ClientUser:    srvusers.NewSrvUsersService(inner.NAMESPACE_MICROSERVICE_SRVUSERS, service.Client()),
				ClientHistory: srvhistoryuserlogin.NewSrvHistoryUserLoginService(inner.NAMESPACE_MICROSERVICE_SRVHISTORYUSERLOGIN, service.Client()),
			},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
