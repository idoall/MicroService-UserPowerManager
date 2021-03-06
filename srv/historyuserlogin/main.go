package main

import (
	"time"

	"github.com/micro/go-log"
	"github.com/micro/go-micro"

	"github.com/idoall/MicroService-UserPowerManager/srv/historyuserlogin/v1/handler"
	proto "github.com/idoall/MicroService-UserPowerManager/srv/historyuserlogin/v1/proto"
	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
)

func main() {

	var err error

	// New Service
	service := micro.NewService(
		micro.Name(inner.NAMESPACE_MICROSERVICE_SRVHISTORYUSERLOGIN),
		micro.RegisterTTL(time.Second*30),      // 注册服务的过期时间
		micro.RegisterInterval(time.Second*20), //间隔多久再次注册服务
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	if err = proto.RegisterProtoHistoryUserLoginHandler(service.Server(), new(handler.HistoryUserLogin)); err != nil {
		log.Fatal(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
