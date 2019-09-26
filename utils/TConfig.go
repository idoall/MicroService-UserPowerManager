package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego/config"
	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
	"github.com/idoall/TokenExchangeCommon/commonutils"
)

var (
	// TConfig is the default config for Application
	TConfig config.Configer
	// AppPath is the absolute path to the app
	AppPath string
	// DataStrategyPath is the absolute path to the app
	DataStrategyPath string
	// appConfigPath is the path to the config files
	appConfigPath string
	// Verbose 是否打印日志信息
	Verbose bool
	// RunMode   dev | prod
	RunMode string
)

func init() {
	var err error
	if AppPath, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		panic(err)
	}

	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	DataStrategyPath = filepath.Join(workPath, "data", "Strategy")

	var filename = "app.conf"
	appConfigPath = filepath.Join(workPath, "conf", filename)
	if !commonutils.PathExists(appConfigPath) {
		panic(fmt.Errorf("not find ./conf/app.conf"))
	}

	TConfig, err = config.NewConfig("ini", appConfigPath)
	if err != nil {
		panic(err)
	}

	RunMode = TConfig.String("default::runmode")
	if RunMode != "dev" && RunMode != "prod" {
		panic("错误的runmode参数， [dev | prod]")
	}

	Verbose = TConfig.DefaultBool("default::Verbose", false)
	if os.Getenv("Verbose") != "" && strings.EqualFold("true", os.Getenv("Verbose")) {
		Verbose = true
	}

	microServiceHostPort := TConfig.String("MicroServices::MicroService_HostPort")
	if os.Getenv("MICROSERVICE_HOSTPORT") != "" {
		inner.MicroServiceHostProt = os.Getenv("MICROSERVICE_HOSTPORT")
	} else {
		inner.MicroServiceHostProt = microServiceHostPort
	}

}
