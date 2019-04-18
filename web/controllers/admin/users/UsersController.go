package users

import (
	"bytes"
	"fmt"
	"net/url"
	"strconv"

	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
	"github.com/idoall/MicroService-UserPowerManager/utils/request"

	"github.com/idoall/MicroService-UserPowerManager/utils"
	"github.com/idoall/MicroService-UserPowerManager/web/controllers/admin"
	"github.com/idoall/MicroService-UserPowerManager/web/models"
)

// UsersController Controller
type UsersController struct {
	admin.AdminBaseController
}

var TemplageBaseURL = "users"
var baseTitle = "用户"
var pageSizeDefault = 11

// GetListJSON Default Json
func (e *UsersController) GetListJSON() {

	var err error

	// return json
	jsonList := struct {
		Rows []struct {
			Id             int64  `json:"Id"`
			UserName       string `json:"UserName"`
			RealyName      string `json:"RealyName"`
			Password       string `json:"Password"`
			Email          string `json:"Email"`
			CreateTime     int64  `json:"CreateTime"`
			LastUpdateTime int64  `json:"LastUpdateTime"`
		} `json:"rows"`
		Total      int64  `json:"total"`
		ErrMessage string `json:"errmsg"`
	}{}

	// get page params
	var pageSize, currentPageIndex int
	if pageSize, err = e.GetInt("pagesize", pageSizeDefault); err != nil {
		jsonList.ErrMessage = err.Error()
		e.Data["json"] = jsonList
		e.ServeJSON()
		return
	}
	if currentPageIndex, err = e.GetInt("currentpage", 1); err != nil {
		jsonList.ErrMessage = err.Error()
		e.Data["json"] = jsonList
		e.ServeJSON()
		return
	}

	// 拼接要发送的url参数
	params := url.Values{}
	params.Set("PageSize", fmt.Sprintf("%d", pageSize))
	params.Set("CurrentPageIndex", fmt.Sprintf("%d", currentPageIndex))

	// 发送请求的路径
	path := fmt.Sprintf("%s%s?%s",
		inner.MicroServiceHostProt,
		utils.TConfig.String("MicroServices::ServiceURL_User_GetList"),
		params.Encode(),
	)

	// 发送 http 请求
	err = request.Request.SendPayload("GET", path, nil, nil, &jsonList, false, false, false)
	if err != nil {
		jsonList.ErrMessage = err.Error()
		e.Data["json"] = jsonList
		e.ServeJSON()
		return
	} else {
		e.Data["json"] = jsonList
		e.ServeJSON()
		return
	}

}

// Get 首页
func (e *UsersController) Get() {

	//set Data
	versionAdminURL := e.GetVersionAdminURL()
	e.Data["title"] = fmt.Sprintf("%s管理", baseTitle)
	e.Data["AddUrl"] = fmt.Sprintf("%s/%s/add", versionAdminURL, TemplageBaseURL)
	e.Data["UpdateUrl"] = fmt.Sprintf("%s/%s/update?id=", versionAdminURL, TemplageBaseURL)
	e.Data["DelUrl"] = fmt.Sprintf("%s/%s/delete?id=", versionAdminURL, TemplageBaseURL)
	e.Data["BatchDelUrl"] = fmt.Sprintf("%s/%s/batchdelete", versionAdminURL, TemplageBaseURL)
	e.Data["JSONListUrl"] = fmt.Sprintf("%s/%s/GetListJSON", versionAdminURL, TemplageBaseURL)

	e.SetMortStype()
	e.SetMortScript()
	e.AppendCustomScripts([]string{
		//Bootstrap table
		"/static/js/hplus/plugins/bootstrap-table/bootstrap-table.min.js",
		"/static/js/hplus/plugins/bootstrap-table/bootstrap-table-mobile.min.js",
		"/static/js/hplus/plugins/bootstrap-table/locale/bootstrap-table-zh-CN.min.js",
	})
	e.AppendCustomStyles([]string{
		//Bootstrap table
		"/static/css/hplus/bootstrap-table/bootstrap-table.min.css",
	})
	e.Layout = "admin/layout/layout.html"
	e.LayoutSections = make(map[string]string)
	e.LayoutSections["CustomHeader"] = "admin/layout/layout-customsheader.html"
	e.TplName = fmt.Sprintf("%s/%s/index.html", admin.TemplageAdminBaseURL, TemplageBaseURL)
}

// Add 添加用户
func (e *UsersController) Add() {
	versionAdminURL := e.GetVersionAdminURL()
	e.Data["title"] = fmt.Sprintf("添加%s", baseTitle)
	e.Data["DefaultUrl"] = fmt.Sprintf("%s/%s", versionAdminURL, TemplageBaseURL)
	e.Data["AddSaveUrl"] = fmt.Sprintf("%s/%s/addsave", versionAdminURL, TemplageBaseURL)

	//公用设置，样式、脚本、layout
	e.SetMortStype()
	e.SetMortScript()
	e.AppendCustomScripts(nil)
	e.AppendCustomStyles(nil)
	e.Layout = "admin/layout/layout.html"
	e.LayoutSections = make(map[string]string)
	e.LayoutSections["CustomHeader"] = "admin/layout/layout-customsheader.html"
	e.TplName = fmt.Sprintf("%s/%s/add.html", admin.TemplageAdminBaseURL, TemplageBaseURL)
}

// AddSave 添加用户-保存
func (e *UsersController) AddSave() {

	// 用于 json 返回的数据
	var result models.Result
	var err error

	// 拼接要发送的url参数
	params := url.Values{}
	params.Set("UserName", e.GetString("username"))
	params.Set("PassWord", e.GetString("realyname"))
	params.Set("RealyName", e.GetString("username"))
	params.Set("Email", e.GetString("email"))
	params.Set("Note", e.GetString("note"))

	// 发送请求的路径
	path := fmt.Sprintf("%s%s", inner.MicroServiceHostProt, utils.TConfig.String("MicroServices::ServiceURL_User_Add"))

	// 临时 Json解析类
	responseJson := struct {
		NewUserId int64
	}{}
	// 发送 http 请求
	if err = request.Request.SendPayload("POST", path, nil, bytes.NewBufferString(params.Encode()), &responseJson, false, true, false); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	} else {
		result.Code = 0
		result.Msg = "添加成功, 用户ID:" + strconv.FormatInt(responseJson.NewUserId, 10)
		e.Data["json"] = result
		e.ServeJSON()
	}

}

// Update 修改用户
func (e *UsersController) Update() {
	var err error
	var userId int64
	// 用于 json 返回的数据
	var result models.Result

	if userId, err = e.GetInt64("id", 0); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	}

	// 拼接要发送的url参数
	params := url.Values{}
	params.Set("UserId", strconv.FormatInt(userId, 10))

	// 发送请求的路径
	path := fmt.Sprintf("%s%s?%s",
		inner.MicroServiceHostProt,
		utils.TConfig.String("MicroServices::ServiceURL_User_GetUser"),
		params.Encode(),
	)

	// 临时 Json解析类
	var responseJson interface{}
	// 发送 http 请求
	if err = request.Request.SendPayload("GET", path, nil, nil, &responseJson, false, true, false); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	} else {
		e.Data["Model"] = responseJson
	}

	//set Data
	versionAdminURL := e.GetVersionAdminURL()
	e.Data["title"] = fmt.Sprintf("修改%s", baseTitle)
	e.Data["UpdateSaveUrl"] = fmt.Sprintf("%s/%s/updatesave", versionAdminURL, TemplageBaseURL)

	//公用设置，样式、脚本、layout
	e.SetMortStype()
	e.SetMortScript()
	e.AppendCustomScripts(nil)
	e.AppendCustomStyles(nil)
	e.Layout = "admin/layout/layout.html"
	e.LayoutSections = make(map[string]string)
	e.LayoutSections["CustomHeader"] = "admin/layout/layout-customsheader.html"
	e.TplName = fmt.Sprintf("%s/%s/update.html", admin.TemplageAdminBaseURL, TemplageBaseURL)
}

// UpdateSave 保存修改用户
func (e *UsersController) UpdateSave() {
	// 用于 json 返回的数据
	var result models.Result
	var err error
	var userId int64
	if userId, err = e.GetInt64("userid", 0); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	}

	// 拼接要发送的url参数
	params := url.Values{}
	params.Set("UserId", strconv.FormatInt(userId, 10))
	params.Set("UserName", e.GetString("username"))
	params.Set("PassWord", e.GetString("password"))
	params.Set("RealyName", e.GetString("username"))
	params.Set("Email", e.GetString("email"))
	params.Set("Note", e.GetString("note"))

	// 发送请求的路径
	path := fmt.Sprintf("%s%s", inner.MicroServiceHostProt, utils.TConfig.String("MicroServices::ServiceURL_User_Update"))

	// 临时 Json解析类
	responseJson := struct {
		UpdateUserId int64
	}{}
	// 发送 http 请求
	if err = request.Request.SendPayload("POST", path, nil, bytes.NewBufferString(params.Encode()), &responseJson, false, true, false); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	} else {
		e.Ctx.Redirect(302, fmt.Sprintf("/%s/%s", admin.AdminBaseRoterVersion, TemplageBaseURL))
	}
}

// Delete 删除
func (e *UsersController) Delete() {
	// 用于 json 返回的数据
	var result models.Result
	var err error
	var userId int64

	if userId, err = e.GetInt64("id", 0); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	} else if userId == 0 {
		result.Code = -1
		result.Msg = fmt.Sprintf("id不能为0")
		e.Data["json"] = result
		e.ServeJSON()
		return
	}

	// 拼接要发送的url参数
	params := url.Values{}
	params.Set("UserIds", strconv.FormatInt(userId, 10))

	// 发送请求的路径
	path := fmt.Sprintf("%s%s", inner.MicroServiceHostProt, utils.TConfig.String("MicroServices::ServiceURL_User_BatchDelete"))

	// 临时 Json解析类
	responseJson := struct {
		Deleted int64
	}{}
	// 发送 http 请求
	if err = request.Request.SendPayload("POST", path, nil, bytes.NewBufferString(params.Encode()), &responseJson, false, true, false); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	} else {
		e.Ctx.Redirect(302, fmt.Sprintf("/%s/%s", admin.AdminBaseRoterVersion, TemplageBaseURL))
	}
}

// BatchDelete 批量删除
func (e *UsersController) BatchDelete() {

	// 用于 json 返回的数据
	var result models.Result
	var err error

	userIds := e.GetString("ids")
	if userIds == "" {
		result.Code = -1
		result.Msg = "ids 不能为空"
		e.Data["json"] = result
		e.ServeJSON()
		return
	}

	// 拼接要发送的url参数
	params := url.Values{}
	params.Set("UserIds", userIds)
	fmt.Println(params.Encode())

	// 发送请求的路径
	path := fmt.Sprintf("%s%s", inner.MicroServiceHostProt, utils.TConfig.String("MicroServices::ServiceURL_User_BatchDelete"))

	// 临时 Json解析类
	responseJson := struct {
		Deleted int64
	}{}
	// 发送 http 请求
	if err = request.Request.SendPayload("POST", path, nil, bytes.NewBufferString(params.Encode()), &responseJson, false, true, false); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
	} else {
		result.Code = 0
		e.Data["json"] = result
	}

	e.ServeJSON()
}
