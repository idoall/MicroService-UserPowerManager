package admin

import (
	"fmt"

	"github.com/idoall/MicroService-UserPowerManager/web/controllers"
	"github.com/idoall/TokenExchangeCommon/commonutils"
)

// Result struct
type Result struct {
	Code  int    `json:"code"`
	Title string `json:"title"`
	Msg   string `json:"msg"`
}

// AdminBaseController struct
type AdminBaseController struct {
	controllers.BaseController
}

var (
	AdminBaseRoterVersion = "v1"
	TemplageAdminBaseURL  = "admin"
)

// 返回版本号+admin拼接的url
func (e *AdminBaseController) GetVersionAdminURL() string {
	return fmt.Sprintf("/%s/%s", AdminBaseRoterVersion, TemplageAdminBaseURL)
}

// AppendCustomStyles 注册自定义脚样式
func (e *AdminBaseController) AppendCustomStyles(list []string) {
	styleList := []string{
		"/static/css/hplus/iCheck/custom.css",
		"/static/css/hplus/chosen/chosen.css",
		"/static/css/awesome-bootstrap-checkbox/awesome-bootstrap-checkbox.css",
		"/static/css/hplus/sweetalert/sweetalert.css",
	}
	for _, v := range list {
		if !commonutils.StringDataContains(styleList, v) {
			styleList = append(styleList, v)
		}
	}
	e.Data["customStyles"] = styleList
}

// AppendCustomScripts 注册自定义脚本
func (e *AdminBaseController) AppendCustomScripts(list []string) {
	scriptList := []string{
		"/static/js/hplus/plugins/validate/jquery.validate.min.js",
		"/static/js/hplus/plugins/validate/messages_zh.min.js",
		// "/static/js/hplus/content.js?v=1.0.0",
		"/static/js/hplus/plugins/iCheck/1.0.2/icheck.min.js",
		// "/static/js/ianPager.js",
		"/static/js/hplus/plugins/sweetalert/2.1.0/sweetalert.min.js", //第三方插件
		// "/static/js/hplus/plugins/echarts/4.1.0/echarts.min.js",       //百度echarets
		//layer
		"/static/js/hplus/plugins/layer/layer.min.js",
	}
	for _, v := range list {
		if !commonutils.StringDataContains(scriptList, v) {
			scriptList = append(scriptList, v)
		}
	}
	e.Data["customScripts"] = scriptList
}
