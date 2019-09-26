package users

import (
	"bytes"
	"fmt"
	"net/url"
	"strconv"

	"github.com/idoall/MicroService-UserPowerManager/utils/request"
	"github.com/idoall/MicroService-UserPowerManager/web/controllers/admin/users/models"

	"github.com/idoall/MicroService-UserPowerManager/web/controllers/admin"
	"github.com/idoall/MicroService-UserPowerManager/web/controllers/admin/usersgroup"
)

// UsersController Controller
type UsersController struct {
	admin.AdminBaseController
}

var (
	TemplageBaseURL = "users"
	baseTitle       = "用户"
	pageSizeDefault = 11
)

// GetListJSON Default Json
func (e *UsersController) GetListJSON() {

	var err error

	// return json
	jsonList := struct {
		Rows []struct {
			ID             int64  `json:"ID"`
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

	// 发送 http 请求
	err = request.Request.WebGETSendPayload("ServiceURL_User_GetList", params, &jsonList, true, false, false, false)
	if err != nil {
		jsonList.ErrMessage = err.Error()
		e.Data["json"] = jsonList
		e.ServeJSON()
	} else {
		e.Data["json"] = jsonList
		e.ServeJSON()
	}

}

// Get 首页
func (e *UsersController) Get() {

	//set Data
	versionAdminURL := e.GetVersionAdminBaseURL()
	e.Data["title"] = fmt.Sprintf("%s管理", baseTitle)
	e.Data["AddUrl"] = fmt.Sprintf("%s/%s/add", versionAdminURL, TemplageBaseURL)
	e.Data["UpdateUrl"] = fmt.Sprintf("%s/%s/update?id=", versionAdminURL, TemplageBaseURL)
	e.Data["DelUrl"] = fmt.Sprintf("%s/%s/delete?id=", versionAdminURL, TemplageBaseURL)
	e.Data["BatchDelUrl"] = fmt.Sprintf("%s/%s/batchdelete", versionAdminURL, TemplageBaseURL)
	e.Data["User2UserGroupURL"] = fmt.Sprintf("%s/%s/user2usersgroup?id=", versionAdminURL, TemplageBaseURL)
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
	versionAdminURL := e.GetVersionAdminBaseURL()
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

	// 临时 Json解析类
	responseJSON := struct {
		NewID int64 `json:"newid"`
	}{}
	// 发送 http 请求
	if err = request.Request.WebPOSTSendPayload("ServiceURL_User_Add", bytes.NewBufferString(params.Encode()), &responseJSON, true, false, false, false); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	}

	result.Code = 0
	result.Msg = "添加成功, 用户ID:" + strconv.FormatInt(responseJSON.NewID, 10)
	e.Data["json"] = result
	e.ServeJSON()

}

// Update 修改用户
func (e *UsersController) Update() {
	var err error
	var userID int64
	// 用于 json 返回的数据
	var result models.Result

	if userID, err = e.GetInt64("id", 0); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	}

	// 拼接要发送的url参数
	params := url.Values{}
	params.Set("ID", strconv.FormatInt(userID, 10))

	// 临时 Json解析类
	var responseJSON map[string]interface{}
	// 发送 http 请求
	if err = request.Request.WebGETSendPayload("ServiceURL_User_GetUser", params, &responseJSON, true, false, false, false); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	}

	e.Data["Model"] = responseJSON["Model"]

	//set Data
	versionAdminURL := e.GetVersionAdminBaseURL()
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
	var userID int64
	if userID, err = e.GetInt64("userid", 0); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	}

	// 拼接要发送的url参数
	params := url.Values{}
	params.Set("ID", strconv.FormatInt(userID, 10))
	params.Set("UserName", e.GetString("username"))
	params.Set("PassWord", e.GetString("password"))
	params.Set("RealyName", e.GetString("username"))
	params.Set("Email", e.GetString("email"))
	params.Set("Note", e.GetString("note"))

	// 临时 Json 解析类
	responseJSON := struct {
		Updated int64
	}{}
	// 发送 http 请求
	if err = request.Request.WebPOSTSendPayload("ServiceURL_User_Update", bytes.NewBufferString(params.Encode()), &responseJSON, true, false, false, false); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	}
	e.Ctx.Redirect(302, fmt.Sprintf("%s/%s", e.GetVersionAdminBaseURL(), TemplageBaseURL))

}

// Delete 删除
func (e *UsersController) Delete() {
	// 用于 json 返回的数据
	var result models.Result
	var err error
	var userID int64

	if userID, err = e.GetInt64("id", 0); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	} else if userID == 0 {
		result.Code = -1
		result.Msg = fmt.Sprintf("id不能为0")
		e.Data["json"] = result
		e.ServeJSON()
		return
	}

	// 拼接要发送的url参数
	params := url.Values{}
	params.Set("IDArray", strconv.FormatInt(userID, 10))

	// 临时 Json 解析类
	responseJSON := struct {
		Deleted int64
	}{}
	// 发送 http 请求
	if err = request.Request.WebPOSTSendPayload("ServiceURL_User_BatchDelete", bytes.NewBufferString(params.Encode()), &responseJSON, true, false, false, false); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	}

	e.Ctx.Redirect(302, fmt.Sprintf("%s/%s", e.GetVersionAdminBaseURL(), TemplageBaseURL))

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
	params.Set("IDArray", userIds)

	// 临时 JSON 解析类
	responseJSON := struct {
		Deleted int64
	}{}
	// 发送 http 请求
	if err = request.Request.WebPOSTSendPayload("ServiceURL_User_BatchDelete", bytes.NewBufferString(params.Encode()), &responseJSON, true, false, false, false); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
	} else {
		result.Code = 0
		e.Data["json"] = result
	}

	e.ServeJSON()
}

// User2UserGroup View
func (e *UsersController) User2UserGroup() {

	var err error
	var result models.Result
	userID, _ := e.GetInt64("id", 0)

	// 拼接要发送的url参数
	params := url.Values{}
	params.Set("ID", strconv.FormatInt(userID, 10))

	// 临时 Json解析类
	var responseUserJSON map[string]map[string]interface{}
	// 发送 http 请求
	if err = request.Request.WebGETSendPayload("ServiceURL_User_GetUser", params, &responseUserJSON, true, false, false, false); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	}

	//set Data
	versionAdminURL := e.GetVersionAdminBaseURL()
	e.Data["Model"] = responseUserJSON["Model"]
	e.Data["title"] = fmt.Sprintf("%s[%s] -  用户组配置", baseTitle, responseUserJSON["Model"]["UserName"])
	e.Data["URL_UserGroupListJSON"] = fmt.Sprintf("%s/%s/GetListJSON", versionAdminURL, usersgroup.TemplageBaseURL)
	e.Data["URL_User2UserGroupSave"] = fmt.Sprintf("%s/%s/user2usersgroupsave", versionAdminURL, TemplageBaseURL)
	e.Data["URL_GetUserGroupsByUser"] = fmt.Sprintf("%s/%s/GetUserGroupsByUser", versionAdminURL, TemplageBaseURL)
	e.Data["URL_ColumnPowerUrl"] = fmt.Sprintf("%s/%s/ColumnPower?id=", versionAdminURL, usersgroup.TemplageBaseURL)

	//公用设置，样式、脚本、layout
	e.SetMortStype()
	e.SetMortScript()
	e.AppendCustomScripts([]string{
		//Bootstrap table
		"/static/js/hplus/plugins/bootstrap-table/bootstrap-table.min.js",
		// "/static/js/hplus/plugins/bootstrap-table/bootstrap-table-mobile.min.js",
		"/static/js/hplus/plugins/bootstrap-table/locale/bootstrap-table-zh-CN.min.js",
		"/static/js/hplus/plugins/layer/layer.min.js",
	})
	e.AppendCustomStyles([]string{
		//Bootstrap table
		"/static/css/hplus/bootstrap-table/bootstrap-table.min.css",
	})
	e.Layout = "admin/layout/layout.html"
	e.LayoutSections = make(map[string]string)
	e.LayoutSections["CustomHeader"] = "admin/layout/layout-customsheader.html"
	e.TplName = fmt.Sprintf("%s/%s/user2group.html", admin.TemplageAdminBaseURL, TemplageBaseURL)
}

// User2UserGroupSave View
func (e *UsersController) User2UserGroupSave() {
	var err error
	var result models.Result
	userID, _ := e.GetInt64("id", 0)
	// 拼接要发送的url参数
	params := url.Values{}
	params.Set("ID", strconv.FormatInt(userID, 10))

	// 临时 Json解析类
	var responseUserJSON map[string]map[string]interface{}
	// 发送 http 请求
	if err = request.Request.WebGETSendPayload("ServiceURL_User_GetUser", params, &responseUserJSON, true, false, false, false); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	}

	//先删除用户的权限
	params = url.Values{}
	params.Set("User", fmt.Sprintf("%.f", responseUserJSON["Model"]["ID"].(float64)))
	if err = request.Request.WebPOSTSendPayload("ServiceURL_Role_DeleteRolesForUser", bytes.NewBufferString(params.Encode()), nil, true, false, false, false); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	}

	// 添加用户和角色（组）的关系
	params = url.Values{}
	params.Set("User", fmt.Sprintf("%.f", responseUserJSON["Model"]["ID"].(float64)))
	params.Set("UserGroup", e.GetString("ids"))
	if err = request.Request.WebPOSTSendPayload("ServiceURL_Role_AddGroupingPolicy", bytes.NewBufferString(params.Encode()), nil, true, false, false, false); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	}
	result.Code = 0
	e.Data["json"] = result
	e.ServeJSON()
}

// GetUserGroupsByUser View
func (e *UsersController) GetUserGroupsByUser() {
	var err error
	var result models.Result
	userID, _ := e.GetInt64("id", 0)

	// 拼接要发送的url参数
	params := url.Values{}
	params.Set("ID", strconv.FormatInt(userID, 10))

	// 临时 Json解析类
	var responseUserJSON map[string]map[string]interface{}
	// 发送 http 请求
	if err = request.Request.WebGETSendPayload("ServiceURL_User_GetUser", params, &responseUserJSON, true, false, false, false); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	}

	// 获取用户所属用户组
	params = url.Values{}
	params.Set("Name", fmt.Sprintf("%.f", responseUserJSON["Model"]["ID"].(float64)))

	var roleArray []string
	// 发送 http 请求
	if err = request.Request.WebPOSTSendPayload("ServiceURL_Role_GetRolesForUser", bytes.NewBufferString(params.Encode()), &roleArray, true, false, false, false); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
	} else {
		result.Code = 0
		e.Data["json"] = roleArray
	}

	e.ServeJSON()
}
