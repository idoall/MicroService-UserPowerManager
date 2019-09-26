package inner

import "github.com/idoall/MicroService-UserPowerManager/utils/log4"

var (
	// 全局 log
	Mlogger log4.Logger4
	// 微服务的地址，单独给web程序使用
	MicroServiceHostProt string
)

const (
	LOGFILENAME = "microservice"

	// 注册的命名空间 ID - go micro api
	NAMESPACE_MICROSERVICE_API = "go.micro.api.mshk.api.v1"

	// 注册的命名空间 ID - go micro api Users
	NAMESPACE_MICROSERVICE_APIUSERS = "go.micro.api.mshk.v1.users"
	// 注册的命名空间 ID - go micro api UsersGroup
	NAMESPACE_MICROSERVICE_APIUSERSGROUP = "go.micro.api.mshk.v1.usersgroup"
	// 注册的命名空间 ID - go micro api role
	NAMESPACE_MICROSERVICE_APIROLE = "go.micro.api.mshk.v1.role"
	// 注册的命名空间 ID - go micro api column
	NAMESPACE_MICROSERVICE_APICOLUMNS = "go.micro.api.mshk.v1.columns"

	// NAMESPACE_MICROSERVICE_APIEMAIL 注册的命名空间 ID - go micro api Check Science Online 检查是否能够科学上网
	NAMESPACE_MICROSERVICE_APICHECKSCIENCEONLINE = "go.micro.api.mshk.v1.checkscienceonline"
	// NAMESPACE_MICROSERVICE_APIEMAIL 注册的命名空间 ID - go micro api SMTP service
	NAMESPACE_MICROSERVICE_APISMTPSERVER = "go.micro.api.mshk.v1.smtpservice"
	// NAMESPACE_MICROSERVICE_APIEXCHANGE 注册的命名空间 ID - go micro api Exchange
	NAMESPACE_MICROSERVICE_APIEXCHANGE = "go.micro.api.mshk.v1.exchange"
	// NAMESPACE_MICROSERVICE_APIV8STRATEGY 注册的命名空间 ID - go micro api strategy
	NAMESPACE_MICROSERVICE_APISTRATEGY = "go.micro.api.mshk.v1.strategy"

	// 注册的命名空间 ID - go micro server - Users
	NAMESPACE_MICROSERVICE_SRVUSERS = "go.micro.srv.mshk.v1.users"
	// 注册的命名空间 ID - go micro server - UsersGroup
	NAMESPACE_MICROSERVICE_SRVUSERSGROUP = "go.micro.srv.mshk.v1.usersgroup"
	// 注册的命名空间 ID - go micro server - Columns
	NAMESPACE_MICROSERVICE_SRVCOLUMNS = "go.micro.srv.mshk.v1.columns"
	// 注册的命名空间 ID - go micro server - HistoryUserLogin
	NAMESPACE_MICROSERVICE_SRVHISTORYUSERLOGIN = "go.micro.srv.mshk.v1.historyuserlogin"
	// 注册的命名空间 ID - go micro server - Role
	NAMESPACE_MICROSERVICE_SRVROLE = "go.micro.srv.mshk.v1.role"
)
