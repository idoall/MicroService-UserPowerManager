package inner

import "github.com/idoall/MicroService-UserPowerManager/log4"

var (
	// 全局 log
	Mlogger log4.Logger4
	// 微服务的地址，单独给web程序使用
	MicroServiceHostProt string
)

const (

	// 注册的命名空间 ID - go micro api
	NAMESPACE_MICROSERVICE_API = "go.micro.api.mshk.api.v1"
	// 注册的命名空间 ID - go micro server - Users
	NAMESPACE_MICROSERVICE_SRVUSERS = "go.micro.srv.mshk.v1.srvusers"
	// 注册的命名空间 ID - go micro server - UsersGroup
	NAMESPACE_MICROSERVICE_SRVUSERSGROUP = "go.micro.srv.mshk.v1.srvusersgroup"
	// 注册的命名空间 ID - go micro server - HistoryUserLogin
	NAMESPACE_MICROSERVICE_SRVHISTORYUSERLOGIN = "go.micro.srv.mshk.v1.srvhistoryuserlogin"
)
