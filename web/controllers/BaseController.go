package controllers

import (
	"github.com/astaxie/beego"
	"github.com/idoall/TokenExchangeCommon/commonutils"
)

// BaseController struct
type BaseController struct {
	beego.Controller
}

// SetMortStype 设置本地
func (e *BaseController) SetMortStype() {
	e.Data["moreStyles"] = []string{
		"/static/css/bootstrap-3.3.6.min.css?v=3.3.6",
		"/static/css/hplus/font-awesome.min.css?v=4.4.0",
		"/static/css/hplus/animate.css",
		"/static/css/hplus/style.css?v=4.1.0",
	}
}

// SetMortScript 设置本地
func (e *BaseController) SetMortScript() {
	e.Data["moreScripts"] = []string{
		"/static/js/jquery-2.2.4.min.js?v=2.2.4",
		"/static/js/bootstrap-3.3.6.min.js?v=3.3.6",
	}
}

// AppendCustomScripts 注册自定义脚本
func (e *BaseController) AppendCustomScripts(list []string) {
	scriptList := []string{}
	for _, v := range list {
		if !commonutils.StringDataContains(scriptList, v) {
			scriptList = append(scriptList, v)
		}
	}
	e.Data["customScripts"] = scriptList
}
