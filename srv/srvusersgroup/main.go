package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"

	"github.com/idoall/MicroService-UserPowerManager/srv/srvusers/handler"
	proto "github.com/idoall/MicroService-UserPowerManager/srv/srvusers/proto"
	"github.com/idoall/MicroService-UserPowerManager/utils"
	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
	"github.com/idoall/MicroService-UserPowerManager/utils/jaeger"
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

	// New Service
	service := micro.NewService(
		micro.Name(inner.NAMESPACE_MICROSERVICE_SRVUSERS),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	proto.RegisterSrvUsersHandler(service.Server(), new(handler.SrvUsers))

	// Register Struct as Subscriber
	// micro.RegisterSubscriber("go.micro.srv.template", service.Server(), new(subscriber.Example))

	// // Register Function as Subscriber
	// micro.RegisterSubscriber("go.micro.srv.template", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
