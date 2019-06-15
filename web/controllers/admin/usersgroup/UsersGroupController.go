package usersgroup

import (
	"bytes"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/idoall/MicroService-UserPowerManager/utils/request"
	"github.com/idoall/MicroService-UserPowerManager/web/controllers/admin"
	"github.com/idoall/MicroService-UserPowerManager/web/controllers/admin/columns"
	"github.com/idoall/MicroService-UserPowerManager/web/models"
)

// UsersGroupController Controller
type UsersGroupController struct {
	admin.AdminBaseController
}

var (
	TemplageBaseURL = "usersgroup"
	baseTitle       = "用户组"
	pageSizeDefault = 11
)

// GetListJSON Default Json
func (e *UsersGroupController) GetListJSON() {
	var err error

	// return json
	jsonList := struct {
		Rows []struct {
			ID             int64  `json:"ID"`
			Name           string `json:"Name"`
			ParentID       int64  `json:"ParentID"`
			Sorts          int    `json:"Sorts"`
			Note           string `json:"Note"`
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
	err = request.Request.WebGETSendPayload("ServiceURL_UsersGroup_GetList", params, &jsonList)
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

// Get 首页 /v1/admin/usersgroup [get]
func (e *UsersGroupController) Get() {

	//set Data
	versionAdminURL := e.GetVersionAdminBaseURL()
	e.Data["title"] = fmt.Sprintf("%s管理", baseTitle)
	e.Data["AddUrl"] = fmt.Sprintf("%s/%s/add", versionAdminURL, TemplageBaseURL)
	e.Data["UpdateUrl"] = fmt.Sprintf("%s/%s/update?id=", versionAdminURL, TemplageBaseURL)
	e.Data["DelUrl"] = fmt.Sprintf("%s/%s/delete?id=", versionAdminURL, TemplageBaseURL)
	e.Data["URL_ColumnPower"] = fmt.Sprintf("%s/%s/ColumnPower?id=", versionAdminURL, TemplageBaseURL)
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

// Add 添加 /v1/admin/usersgroup/add [get]
func (e *UsersGroupController) Add() {
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

// AddSave 添加 - 保存 /v1/admin/usersgroup/addsave [post]
func (e *UsersGroupController) AddSave() {

	// 用于 json 返回的数据
	var result models.Result
	var err error

	// 拼接要发送的url参数
	params := url.Values{}
	params.Set("Name", e.GetString("name"))
	params.Set("Sorts", e.GetString("sorts"))
	params.Set("Note", e.GetString("note"))

	// 临时 Json解析类
	responseJSON := struct {
		NewID int64 `json:"newid"`
	}{}
	// 发送 http 请求
	if err = request.Request.WebPOSTSendPayload("ServiceURL_UsersGroup_Add", bytes.NewBufferString(params.Encode()), &responseJSON); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	} else {
		e.Ctx.Redirect(302, fmt.Sprintf("%s/%s", e.GetVersionAdminBaseURL(), TemplageBaseURL))
	}
}

// Update 修改
func (e *UsersGroupController) Update() {
	var err error
	var ID int64
	// 用于 json 返回的数据
	var result models.Result

	if ID, err = e.GetInt64("id", 0); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	}

	// 拼接要发送的url参数
	params := url.Values{}
	params.Set("ID", strconv.FormatInt(ID, 10))

	// 临时 Json解析类
	var responseJSON map[string]interface{}
	// 发送 http 请求
	if err = request.Request.WebGETSendPayload("ServiceURL_UsersGroup_Get", params, &responseJSON); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	} else {
		e.Data["Model"] = responseJSON
	}

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

// UpdateSave 保存修改
func (e *UsersGroupController) UpdateSave() {
	// 用于 json 返回的数据
	var result models.Result
	var err error
	var ID int64
	if ID, err = e.GetInt64("ID", 0); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	}

	// 拼接要发送的url参数
	params := url.Values{}
	params.Set("ID", strconv.FormatInt(ID, 10))
	params.Set("Name", e.GetString("name"))
	params.Set("Sorts", e.GetString("sorts"))
	params.Set("Note", e.GetString("note"))

	// 临时 Json 解析类
	responseJSON := struct {
		Updated int64
	}{}
	// 发送 http 请求
	if err = request.Request.WebPOSTSendPayload("ServiceURL_UsersGroup_Update", bytes.NewBufferString(params.Encode()), &responseJSON); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	} else {
		e.Ctx.Redirect(302, fmt.Sprintf("%s/%s", e.GetVersionAdminBaseURL(), TemplageBaseURL))
	}
}

// Delete 删除
func (e *UsersGroupController) Delete() {

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

	// 临时 Json解析类
	responseJSON := struct {
		Deleted int64
	}{}
	// 发送 http 请求
	if err = request.Request.WebPOSTSendPayload("ServiceURL_UsersGroup_BatchDelete", bytes.NewBufferString(params.Encode()), &responseJSON); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	} else {
		e.Ctx.Redirect(302, fmt.Sprintf("%s/%s", e.GetVersionAdminBaseURL(), TemplageBaseURL))
	}
}

// BatchDelete 批量删除
func (e *UsersGroupController) BatchDelete() {

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
	fmt.Println(params.Encode())

	// 临时 Json解析类
	responseJSON := struct {
		Deleted int64
	}{}
	// 发送 http 请求
	if err = request.Request.WebPOSTSendPayload("ServiceURL_UsersGroup_BatchDelete", bytes.NewBufferString(params.Encode()), &responseJSON); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
	} else {
		result.Code = 0
		e.Data["json"] = result
	}

	e.ServeJSON()
}

//-------------------权限配置

// ColumnPower 展示用户组-权限页面
func (e *UsersGroupController) ColumnPower() {
	var err error
	var ID int64
	// 用于 json 返回的数据
	var result models.Result

	if ID, err = e.GetInt64("id", 0); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	}

	// 拼接要发送的url参数
	params := url.Values{}
	params.Set("ID", strconv.FormatInt(ID, 10))

	// 临时 Json解析类
	var responseJSON map[string]interface{}
	// 发送 http 请求
	if err = request.Request.WebGETSendPayload("ServiceURL_UsersGroup_Get", params, &responseJSON); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	} else {
		e.Data["Model"] = responseJSON
	}

	//set Data
	versionAdminURL := e.GetVersionAdminBaseURL()
	e.Data["Model"] = responseJSON
	e.Data["title"] = fmt.Sprintf("%s『%s』 权限配置", baseTitle, responseJSON["Name"].(string))
	e.Data["ColumnPowerSaveURL"] = fmt.Sprintf("%s/%s/ColumnPowerSave", versionAdminURL, TemplageBaseURL)
	e.Data["URL_GetColumnPowerTreeViewJSON"] = fmt.Sprintf("%s/%s/GetColumnPowerTreeViewJSON", versionAdminURL, TemplageBaseURL)
	e.Data["JSONTreeViewListSelectedDataUrl"] = fmt.Sprintf("%s/%s/GetTreeViewJSONSelectedData", versionAdminURL, TemplageBaseURL)

	//公用设置，样式、脚本、layout
	e.SetMortStype()
	e.SetMortScript()
	e.AppendCustomScripts([]string{
		//Bootstrap table
		"/static/js/hplus/plugins/bootstrap-table/bootstrap-table.min.js",
		"/static/js/hplus/plugins/bootstrap-table/locale/bootstrap-table-zh-CN.min.js",
		// TreeView
		"/static/js/hplus/plugins/treeview/bootstrap-treeview.min.js",
		//layer
		"/static/js/hplus/plugins/layer/layer.min.js",
	})
	e.AppendCustomStyles([]string{
		//Bootstrap table
		"/static/css/hplus/bootstrap-table/bootstrap-table.min.css",
		// TreeView
		"/static/css/hplus/treeview/bootstrap-treeview.min.css",
	})
	e.Layout = "admin/layout/layout.html"
	e.LayoutSections = make(map[string]string)
	e.LayoutSections["CustomHeader"] = "admin/layout/layout-customsheader.html"
	e.TplName = fmt.Sprintf("%s/%s/columnpower.html", admin.TemplageAdminBaseURL, TemplageBaseURL)

}

// GetColumnPowerTreeViewJSON Default Json
func (e *UsersGroupController) GetColumnPowerTreeViewJSON() {
	var err error
	var list []*models.TreeView
	var result models.Result

	// 拼接要发送的url参数
	params := url.Values{}
	params.Set("User", e.GetString("id"))

	// 临时 Json解析类
	var responseJSON []map[string][]string
	// 发送 http 请求
	if err = request.Request.WebGETSendPayload("ServiceURL_Role_GetPermissionsForUser", params, &responseJSON); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	}

	//取出用户组的所有权限
	columnPowerList := responseJSON
	// columnPowerList := admin.RoleS.GetPermissionsForUser("usergroup_" + e.GetString("id"))

	list, err = new(columns.ColumnsController).GetTreeViewBootstrap()
	if err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	}

	for _, v := range list {
		for _, cv := range columnPowerList {
			cid, _ := strconv.ParseInt(cv["Two"][1], 10, 64)
			if v.ID == cid {
				v.State = &models.TreeViewState{Checked: true, Expanded: true}
			}
			if v.Nodes != nil {
				e.getRecursiveColumnPowerTreeView(columnPowerList, v.Nodes)
			}
		}

	}
	// list[0].State = &models.TreeViewState{Checked: true, Expanded: true}

	e.Data["json"] = list
	e.ServeJSON()
}

// getRecursiveColumnPowerTreeView 递归获取下一级
func (e *UsersGroupController) getRecursiveColumnPowerTreeView(columnPowerList []map[string][]string, list []*models.TreeView) {

	for _, v := range list {
		for _, cv := range columnPowerList {
			cid, _ := strconv.ParseInt(cv["Two"][1], 10, 64)
			if v.ID == cid {
				v.State = &models.TreeViewState{Checked: true, Expanded: true}
			}
			if v.Nodes != nil {
				e.getRecursiveColumnPowerTreeView(columnPowerList, v.Nodes)
			}
		}
	}
}

// ColumnPowerSave 展示用户组-权限页面-保存
func (e *UsersGroupController) ColumnPowerSave() {
	var err error
	var result models.Result
	idArray := strings.Split(e.GetString("ids"), ",")
	id := e.GetString("id")
	action := "GET"

	//---------------先删除所有的权限
	// 拼接要发送的url参数
	params := url.Values{}
	params.Set("Role", e.GetString("id"))

	// 发送 http 请求
	if err = request.Request.WebPOSTSendPayload("ServiceURL_Role_RemoveFilteredPolicy", bytes.NewBufferString(params.Encode()), nil); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	}

	for _, v := range idArray {
		mid, _ := strconv.ParseInt(v, 10, 64)
		m, _ := new(columns.ColumnsController).GetColumnsByID(mid)
		//保存用户组的权限
		params = url.Values{}
		params.Set("S1", "usergroup_"+id)
		params.Set("S2", v)
		params.Set("S3", m["URL"].(string))
		params.Set("S4", action)

		// 发送 http 请求
		if err = request.Request.WebPOSTSendPayload("ServiceURL_Role_AddPolicy", bytes.NewBufferString(params.Encode()), nil); err != nil {
			result.Code = -1
			result.Msg = err.Error()
			e.Data["json"] = result
			e.ServeJSON()
			return
		}
	}

	// admin.RoleS.LoadPolicy()
	result.Code = 0
	e.Data["json"] = result
	e.ServeJSON()
}
