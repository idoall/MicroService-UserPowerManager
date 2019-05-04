package usersgroup

import (
	"bytes"
	"fmt"
	"net/url"
	"strconv"

	"github.com/idoall/TokenExchangeCommon/commonutils"

	"github.com/idoall/MicroService-UserPowerManager/utils"
	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
	"github.com/idoall/MicroService-UserPowerManager/utils/request"
	"github.com/idoall/MicroService-UserPowerManager/web/controllers/admin"
	"github.com/idoall/MicroService-UserPowerManager/web/models"
)

// UsersGroupController Controller
type UsersGroupController struct {
	admin.AdminBaseController
}

var TemplageBaseURL = "usersgroup"
var baseTitle = "用户组"
var pageSizeDefault = 11

// GetListJSON Default Json
func (e *UsersGroupController) GetListJSON() {
	var err error

	// return json
	jsonList := struct {
		Rows []struct {
			ID             int64  `json:"ID"`
			Name           string `json:"Name"`
			Sorts          int64  `json:"Sorts"`
			ParentID       int64  `json:"ParentID"`
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

	// 发送请求的路径
	path := fmt.Sprintf("%s%s?%s",
		inner.MicroServiceHostProt,
		utils.TConfig.String("MicroServices::ServiceURL_UsersGroup_GetList"),
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

// Get 首页 /v1/admin/usersgroup [get]
func (e *UsersGroupController) Get() {

	//set Data
	versionAdminURL := e.GetVersionAdminBaseURL()
	e.Data["title"] = fmt.Sprintf("%s管理", baseTitle)
	e.Data["URL_Add"] = fmt.Sprintf("%s/%s/add", versionAdminURL, TemplageBaseURL)
	e.Data["URL_Update"] = fmt.Sprintf("%s/%s/update/", versionAdminURL, TemplageBaseURL)
	e.Data["URL_Del"] = fmt.Sprintf("%s/%s/delete/", versionAdminURL, TemplageBaseURL)
	e.Data["URL_ColumnPower"] = fmt.Sprintf("%s/%s/ColumnPower?id=", versionAdminURL, TemplageBaseURL)
	e.Data["URL_BatchDelete"] = fmt.Sprintf("%s/%s/batchdelete", versionAdminURL, TemplageBaseURL)
	e.Data["URL_JsonListUrl"] = fmt.Sprintf("%s/%s/GetListJSON", versionAdminURL, TemplageBaseURL)
	// e.Data["Power_ViewUserGroupPowerBtn"] = e.HasPermissions(e.GetUserID(), 19) //显示配置权限按钮
	// e.Data["Power_ViewDelBtn"] = e.HasPermissions(e.GetUserID(), 25)            //显示删除和批量删除按钮

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

	// 发送请求的路径
	path := fmt.Sprintf("%s%s", inner.MicroServiceHostProt, utils.TConfig.String("MicroServices::ServiceURL_UsersGroup_Add"))

	// 临时 Json解析类
	responseJSON := struct {
		NewID int64 `json:"newid"`
	}{}
	// 发送 http 请求
	if err = request.Request.SendPayload("POST", path, nil, bytes.NewBufferString(params.Encode()), &responseJSON, false, true, false); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	} else {
		e.Ctx.Redirect(302, fmt.Sprintf("%s/%s", e.GetVersionAdminBaseURL(), TemplageBaseURL))
	}
}

// Update_GET 修改  /v1/admin/usersgroup/update/:id [get]
func (e *UsersGroupController) Update_Get() {
	var err error

	// 用于 json 返回的数据
	var result models.Result

	// 拼接要发送的url参数
	paramID := e.Ctx.Input.Param(":id")
	params := url.Values{}
	params.Set("ID", paramID)

	// 发送请求的路径
	path := fmt.Sprintf("%s%s?%s",
		inner.MicroServiceHostProt,
		utils.TConfig.String("MicroServices::ServiceURL_UsersGroup_Get"),
		params.Encode(),
	)

	// 临时 Json解析类
	var responseJSON map[string]interface{}
	// 发送 http 请求
	if err = request.Request.SendPayload("GET", path, nil, nil, &responseJSON, false, true, false); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	} else {
		e.Data["Model"] = responseJSON["Model"]
	}

	//set Data
	versionAdminURL := e.GetVersionAdminBaseURL()
	e.Data["title"] = fmt.Sprintf("修改%s", baseTitle)
	e.Data["UpdateSaveUrl"] = fmt.Sprintf("%s/%s/update/%s", versionAdminURL, TemplageBaseURL, paramID)

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

// Update_Post /v1/admin/usersgroup/update/:id [post]
func (e *UsersGroupController) Update_Post() {
	// 用于 json 返回的数据
	var result models.Result
	var err error
	paramID := e.Ctx.Input.Param(":id")

	// 拼接要发送的url参数
	params := url.Values{}
	params.Set("ID", paramID)
	params.Set("Name", e.GetString("name"))
	params.Set("Sorts", e.GetString("sorts"))
	params.Set("Note", e.GetString("note"))

	// 发送请求的路径
	path := fmt.Sprintf("%s%s", inner.MicroServiceHostProt, utils.TConfig.String("MicroServices::ServiceURL_UsersGroup_Update"))

	// 临时 Json 解析类
	responseJSON := struct {
		Updated int64
	}{}
	// 发送 http 请求
	if err = request.Request.SendPayload("POST", path, nil, bytes.NewBufferString(params.Encode()), &responseJSON, false, true, false); err != nil {
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
	var ID int64

	paramID := e.Ctx.Input.Param(":id")
	if ID, err = commonutils.Int64FromString(paramID); err != nil {
		result.Code = -1
		result.Msg = err.Error()
		e.Data["json"] = result
		e.ServeJSON()
		return
	} else if ID == 0 {
		result.Code = -1
		result.Msg = fmt.Sprintf("id不能为0")
		e.Data["json"] = result
		e.ServeJSON()
		return
	}

	// 拼接要发送的url参数
	params := url.Values{}
	params.Set("IDArray", strconv.FormatInt(ID, 10))

	// 发送请求的路径
	path := fmt.Sprintf("%s%s", inner.MicroServiceHostProt, utils.TConfig.String("MicroServices::ServiceURL_UsersGroup_BatchDelete"))

	// 临时 Json解析类
	responseJSON := struct {
		Deleted int64
	}{}
	// 发送 http 请求
	if err = request.Request.SendPayload("POST", path, nil, bytes.NewBufferString(params.Encode()), &responseJSON, false, true, false); err != nil {
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

	// 发送请求的路径
	path := fmt.Sprintf("%s%s", inner.MicroServiceHostProt, utils.TConfig.String("MicroServices::ServiceURL_UsersGroup_BatchDelete"))

	// 临时 Json解析类
	responseJSON := struct {
		Deleted int64
	}{}
	// 发送 http 请求
	if err = request.Request.SendPayload("POST", path, nil, bytes.NewBufferString(params.Encode()), &responseJSON, false, true, false); err != nil {
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
	// var result models.Result
	// id, _ := e.GetInt64("id", 0)

	// model, err := new(models.UserGroup).GetOne(id)
	// if err != nil {
	// 	result.Code = -1
	// 	result.Msg = err.Error()
	// 	e.Data["json"] = result
	// 	e.ServeJSON()
	// 	return
	// }

	// e.Data["Model"] = model
	// e.Data["title"] = fmt.Sprintf("%s『%s』 权限配置", baseTitle, model.Name)
	// e.Data["ColumnPowerSaveURL"] = fmt.Sprintf("/%s/ColumnPowerSave", BaseURL)
	// e.Data["URL_GetColumnPowerTreeViewJSON"] = fmt.Sprintf("/%s/GetColumnPowerTreeViewJSON", BaseURL)
	// e.Data["JSONTreeViewListSelectedDataUrl"] = fmt.Sprintf("/%s/GetTreeViewJSONSelectedData", BaseURL)

	// e.SetMortStype()
	// e.SetMortScript()
	// e.AppendCustomScripts([]string{
	// 	//Bootstrap table
	// 	"/static/js/hplus/plugins/bootstrap-table/bootstrap-table.min.js",
	// 	"/static/js/hplus/plugins/bootstrap-table/locale/bootstrap-table-zh-CN.min.js",
	// 	// TreeView
	// 	"/static/js/hplus/plugins/treeview/bootstrap-treeview.min.js",
	// 	//layer
	// 	"/static/js/hplus/plugins/layer/layer.min.js",
	// })
	// e.AppendCustomStyles([]string{
	// 	//Bootstrap table
	// 	"/static/css/hplus/bootstrap-table/bootstrap-table.min.css",
	// 	// TreeView
	// 	"/static/css/hplus/treeview/bootstrap-treeview.min.css",
	// })

	// e.Layout = "admin/layout/layout.html"
	// e.LayoutSections = make(map[string]string)
	// e.LayoutSections["CustomHeader"] = "admin/layout/layout-customsheader.html"
	// e.TplName = BaseURL + "/columnpower.html"
}

// GetColumnPowerTreeViewJSON Default Json
func (e *UsersGroupController) GetColumnPowerTreeViewJSON() {
	// var list []*models.TreeView
	// var result models.Result

	// //取出用户组的所有权限
	// columnPowerList := admin.RoleS.GetPermissionsForUser("usergroup_" + e.GetString("id"))

	// // fmt.Println(columnPowerList)

	// list, err := new(models.ColumnPower).GetTreeViewBootstrap()
	// if err != nil {
	// 	result.Code = -1
	// 	result.Msg = err.Error()
	// 	e.Data["json"] = result
	// 	e.ServeJSON()
	// 	return
	// }

	// for _, v := range list {
	// 	for _, cv := range columnPowerList {
	// 		cid, _ := strconv.ParseInt(cv[1], 10, 64)
	// 		if v.ID == cid {
	// 			v.State = &models.TreeViewState{Checked: true, Expanded: true}
	// 		}
	// 		if v.Nodes != nil {
	// 			e.getRecursiveColumnPowerTreeView(columnPowerList, v.Nodes)
	// 		}
	// 	}

	// }
	// // list[0].State = &models.TreeViewState{Checked: true, Expanded: true}

	// e.Data["json"] = list
	// e.ServeJSON()
}

// getRecursive 递归获取下一级
// func (e *UsersGroupController) getRecursiveColumnPowerTreeView(columnPowerList [][]string, list []*models.TreeView) {

// 	// for _, v := range list {
// 	// 	for _, cv := range columnPowerList {
// 	// 		cid, _ := strconv.ParseInt(cv[1], 10, 64)
// 	// 		if v.ID == cid {
// 	// 			v.State = &models.TreeViewState{Checked: true, Expanded: true}
// 	// 		}
// 	// 		if v.Nodes != nil {
// 	// 			e.getRecursiveColumnPowerTreeView(columnPowerList, v.Nodes)
// 	// 		}
// 	// 	}
// 	// }
// }

// ColumnPowerSave 展示用户组-权限页面-保存
func (e *UsersGroupController) ColumnPowerSave() {
	// var result models.Result
	// idArray := strings.Split(e.GetString("ids"), ",")
	// id := e.GetString("id")
	// action := "GET"

	// //先删除所有的权限
	// admin.RoleS.RemoveFilteredPolicy(0, "usergroup_"+id)
	// // if !boo {
	// // 	result.Code = -1
	// // 	result.Msg = "删除权限失败"
	// // 	e.Data["json"] = result
	// // 	e.ServeJSON()
	// // 	return
	// // }

	// // return
	// // fmt.Println("admin.RoleS.GetPermissionsForUser", admin.RoleS.GetPermissionsForUser("usergroup_"+id))

	// for _, v := range idArray {
	// 	mid, _ := strconv.ParseInt(v, 10, 64)
	// 	m, _ := new(models.ColumnPower).GetOne(mid)
	// 	//保存用户组的权限
	// 	admin.RoleS.AddPolicy("usergroup_"+id, strconv.FormatInt(m.ID, 10), m.URL, action)
	// }

	// // admin.RoleS.LoadPolicy()
	// result.Code = 0
	// e.Data["json"] = result
	// e.ServeJSON()
}
