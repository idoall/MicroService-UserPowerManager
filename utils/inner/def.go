package inner

import "github.com/idoall/MicroService-UserPowerManager/log4"

var (
	// 全局 log
	Mlogger log4.Logger4
	// 微服务的地址，单独给web程序使用
	MicroServiceHostProt string
)

const (

	// 注册的命名空间 ID
	// NAMESPACE_ID = "mshk.top.api.v1.cusers"
	NAMESPACE_MICROSERVICE_APIUSERS = "go.micro.api.mshk.api.v1"
	// 注册的命名空间 ID
	NAMESPACE_MICROSERVICE_SRVUSERS = "go.micro.srv.mshk.v1.srvusers"
	// 注册的命名空间 ID
	NAMESPACE_MICROSERVICE_SRVHISTORYUSERLOGIN = "go.micro.srv.mshk.v1.srvhistoryuserlogin"
)
