package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/dgrijalva/jwt-go"
	"github.com/idoall/TokenExchangeCommon/commonutils"

	"github.com/idoall/MicroService-UserPowerManager/utils"

	"github.com/idoall/MicroService-UserPowerManager/utils/encrypt"
	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
	"github.com/idoall/MicroService-UserPowerManager/utils/jaeger"

	srvhistoryuserlogin "github.com/idoall/MicroService-UserPowerManager/srv/historyuserlogin/v1/proto"
	srvusers "github.com/idoall/MicroService-UserPowerManager/srv/users/v1/proto"

	"github.com/idoall/TokenExchangeCommon/commonutils/checkmail"
	api "github.com/micro/go-api/proto"
	"github.com/micro/go-micro/errors"
)

// Users struct
type Users struct {
	Client        srvusers.ProtoUsersService
	ClientHistory srvhistoryuserlogin.ProtoHistoryUserLoginService
}

// 添加一个用户
func (e *Users) Add(ctx context.Context, req *api.Request, rsp *api.Response) error {

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_User_Add_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_APIUSERS

	// debug
	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [Users][Add] request", namespaceID)
	}

	// 获取请求参数 - 开始
	var username, password, email, realyname, note string
	if req.Post["UserName"] == nil || req.Post["UserName"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "UserName 不能为空")
	} else {
		username = req.Post["UserName"].Values[0]
	}

	if req.Post["PassWord"] == nil || req.Post["PassWord"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "PassWord 不能为空")
	} else {
		password = req.Post["PassWord"].Values[0]
	}

	if req.Post["Email"] == nil || req.Post["Email"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "Email 不能为空.")
	} else if err := checkmail.ValidateFormat(req.Post["Email"].Values[0]); err != nil {
		return errors.InternalServerError(namespaceID, "Email 的格式不正确:%s", err.Error())
	} else {
		email = req.Post["Email"].Values[0]
	}

	if req.Post["RealyName"] != nil {
		realyname = req.Post["RealyName"].Values[0]
	}

	if req.Post["Note"] != nil {
		note = req.Post["Note"].Values[0]
	}
	// 获取请求参数 - 结束

	// 生成加密盐 - 对密码加密
	node, err := snowflake.NewNode(2)
	if err != nil {
		return errors.BadRequest(namespaceID, "snowflake.NewNode Error:%+v", err)
	}
	authKey := node.Generate()
	password = commonutils.HexEncodeToString(commonutils.GetHMAC(commonutils.HashSHA256, []byte(password), authKey.Bytes()))

	// make request
	resUser := &srvusers.User{
		UserName:  username,
		Password:  password,
		RealyName: realyname,
		AuthKey:   authKey.String(),
		Email:     email,
		Note:      note,
	}

	// 调用服务端方法
	response, err := e.Client.Add(ctx, &srvusers.AddRequest{Model: resUser})
	if err != nil {
		return errors.InternalServerError(namespaceID, err.Error())
	}

	// 输出的 json
	responseJSON := struct {
		NewID int64 `json:"newid"`
	}{}
	responseJSON.NewID = response.NewID
	b, _ := commonutils.JSONEncode(responseJSON)
	rsp.StatusCode = 200
	rsp.Body = string(b)

	// debug
	if utils.RunMode == "dev" {
		inner.Mlogger.Info(response)
		inner.Mlogger.Info("rsp.Body", rsp.Body)
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Api_User_Add_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("NewID", response.NewID)
	}

	return nil
}

// 获取用户列表,默认 id 倒排序
func (e *Users) GetList(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_User_GetList_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_APIUSERS

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [Users][GetList] request", namespaceID)
	}

	// 获取请求参数 - 开始
	var pageSize, currentPageIndex int64
	var orderBy string
	if req.Get["PageSize"] == nil || req.Get["PageSize"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "PageSize 不能为空")
	} else if pageSize, err = commonutils.Int64FromString(req.Get["PageSize"].Values[0]); err != nil {
		return errors.InternalServerError(namespaceID, "PageSize Format Error:%s", err.Error())
	}

	if req.Get["CurrentPageIndex"] == nil || req.Get["CurrentPageIndex"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "CurrentPageIndex 不能为空")
	} else if currentPageIndex, err = commonutils.Int64FromString(req.Get["CurrentPageIndex"].Values[0]); err != nil {
		return errors.InternalServerError(namespaceID, "CurrentPageIndex Format Error:%s", err.Error())

	}

	if req.Get["OrderBy"] != nil {
		orderBy = req.Get["OrderBy"].Values[0]
	}
	// 获取请求参数 - 结束

	// 调用服务端方法
	response, err := e.Client.GetList(ctx, &srvusers.GetListRequest{
		CurrentPageIndex: currentPageIndex,
		PageSize:         pageSize,
		OrderBy:          orderBy,
	})
	if err != nil {
		return errors.InternalServerError(namespaceID, err.Error())
	}

	// return json
	jsonList := struct {
		Rows  []*srvusers.User `json:"rows"`
		Total int64            `json:"total"`
	}{}
	jsonList.Rows = response.List
	jsonList.Total = response.TotalCount

	// 对 json 序列化并输出
	b, _ := json.Marshal(jsonList)
	rsp.StatusCode = 200
	rsp.Body = string(b)

	// debug
	if utils.RunMode == "dev" {
		// inner.Mlogger.Info("rsp.Body", rsp.Body)
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Api_User_GetList_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("PageSize", pageSize)
		span.SetTag("CurrentPageIndex", currentPageIndex)
		span.SetTag("orderBy", orderBy)
		span.SetTag("TotalCount", response.TotalCount)
	}

	return nil
}

// 获取单个用户，根据Id或用户名
func (e *Users) GetUser(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_User_GetUser_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_APIUSERS

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [Users][GetUser] request", namespaceID)
	}

	// 获取请求参数 - 开始
	var ID int64
	var userName string
	if req.Get["ID"] != nil && req.Get["ID"].Values[0] != "0" {
		if ID, err = commonutils.Int64FromString(req.Get["ID"].Values[0]); err != nil {
			return errors.InternalServerError(namespaceID, "ID Format Error:%s", err.Error())
		}
	}

	if req.Get["UserName"] != nil && req.Get["UserName"].Values[0] != "" {
		userName = req.Get["UserName"].Values[0]
	}

	// 获取请求参数 - 结束

	// 调用服务端方法
	response, err := e.Client.Get(ctx, &srvusers.GetRequest{
		ID:       ID,
		UserName: userName,
	})
	if err != nil {
		return errors.InternalServerError(namespaceID, err.Error())
	}

	// 对 json 序列化并输出
	b, _ := json.Marshal(response)
	rsp.StatusCode = 200
	rsp.Body = string(b)

	// debug
	if utils.RunMode == "dev" {
		inner.Mlogger.Info("rsp.Body", rsp.Body)
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Api_User_GetUser_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("ID", ID)
	}

	return nil
}

// 修改用户信息
func (e *Users) Update(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_User_Update_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_APIUSERS

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [Users][Update] request", namespaceID)
	}

	// 获取请求参数 - 开始
	var ID int64
	var username, password, email, realyname, note, authKey string
	if req.Post["ID"] == nil || req.Post["ID"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "ID 不能为空")
	} else if ID, err = commonutils.Int64FromString(req.Post["ID"].Values[0]); err != nil {
		return errors.InternalServerError(namespaceID, "ID Format Error:%s", err.Error())
	}
	if req.Post["UserName"] == nil || req.Post["UserName"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "UserName 不能为空")
	} else {
		username = req.Post["UserName"].Values[0]
	}

	if req.Post["PassWord"] != nil || req.Post["PassWord"].Values[0] != "" {
		password = req.Post["PassWord"].Values[0]
	}

	if req.Post["Email"] == nil || req.Post["Email"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "Email 不能为空")
	} else if err := checkmail.ValidateFormat(req.Post["Email"].Values[0]); err != nil {
		return errors.InternalServerError(namespaceID, "Email 的格式不正确:%s", err.Error())
	} else {
		email = req.Post["Email"].Values[0]
	}

	if req.Post["RealyName"] != nil {
		realyname = req.Post["RealyName"].Values[0]
	}

	if req.Post["Note"] != nil {
		note = req.Post["Note"].Values[0]
	}
	// 获取请求参数 - 结束

	// 调用服务端方法获取用户
	responseGetUser, err := e.Client.Get(ctx, &srvusers.GetRequest{
		ID: ID,
	})
	if err != nil {
		return errors.InternalServerError(namespaceID, err.Error())
	}

	// 如果密码不为空，重新对密码加密
	if password != "" {
		// 重置加密盐 - 对密码加密
		node, err := snowflake.NewNode(2)
		if err != nil {
			return errors.BadRequest(namespaceID, "snowflake.NewNode Error:%+v", err)
		}
		authKey = node.Generate().String()
		password = commonutils.HexEncodeToString(commonutils.GetHMAC(commonutils.HashSHA256, []byte(password), []byte(authKey)))
	} else {
		password = responseGetUser.Model.Password
		authKey = responseGetUser.Model.AuthKey
	}

	// 调用服务端方法 - 修改用户
	responseUpdateUser := &srvusers.User{
		ID:             ID,
		UserName:       username,
		Password:       password,
		RealyName:      realyname,
		AuthKey:        authKey,
		Email:          email,
		Note:           note,
		IsDel:          responseGetUser.Model.IsDel,
		ParentID:       responseGetUser.Model.ParentID,
		CreateTime:     responseGetUser.Model.CreateTime,
		LastUpdateTime: responseGetUser.Model.LastUpdateTime,
	}
	response, err := e.Client.Update(ctx, &srvusers.UpdateRequest{Model: responseUpdateUser})
	if err != nil {
		return errors.InternalServerError(namespaceID, err.Error())
	}

	// 输出的 json
	responseJSON := struct {
		Updated int64
	}{}
	responseJSON.Updated = response.Updated
	b, _ := commonutils.JSONEncode(responseJSON)
	rsp.StatusCode = 200
	rsp.Body = string(b)

	// debug
	if utils.RunMode == "dev" {
		inner.Mlogger.Info("rsp.Body", rsp.Body)
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Api_User_Update_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("ID", ID)
	}

	return nil
}

// 批量删除用户信息
func (e *Users) BatchDelete(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_User_BatchDelete_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_APIUSERS

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [Users][BatchDelete] request", namespaceID)
	}

	// 获取请求参数 - 开始
	var IDArray []string
	if req.Post["IDArray"] == nil || req.Post["IDArray"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "IDArray 不能为空")
	} else {
		IDArray = strings.Split(req.Post["IDArray"].Values[0], ",")
	}

	// 获取请求参数 - 结束

	// 调用服务端方法获取用户
	response, err := e.Client.BatchDelete(ctx, &srvusers.DeleteRequest{
		IDArray: IDArray,
	})
	if err != nil {
		return errors.InternalServerError(namespaceID, err.Error())
	}

	// 输出的 json
	responseJSON := struct {
		Deleted int64 `json:"deleted"`
	}{}
	responseJSON.Deleted = response.Deleted
	b, _ := commonutils.JSONEncode(responseJSON)
	rsp.StatusCode = 200
	rsp.Body = string(b)

	// debug
	if utils.RunMode == "dev" {
		inner.Mlogger.Info("rsp.Body", rsp.Body)
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Api_User_BatchDelete_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("IDArray", strings.Join(IDArray, ","))
	}

	return nil
}

// 用户登录
func (e *Users) UserLogin(ctx context.Context, req *api.Request, rsp *api.Response) error {
	var err error
	var tokenString string

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_User_UserLogin_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespaceID := inner.NAMESPACE_MICROSERVICE_APIUSERS

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [Users][UserLogin] request", namespaceID)
	}

	// 获取请求参数 -------------------- 开始
	var userName, password, deviceDetector, geoRemoteAddr, geoCountry, geoCity string
	if req.Post["UserName"] == nil || req.Post["UserName"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "UserName 不能为空")
	} else {
		userName = req.Post["UserName"].Values[0]
	}

	if req.Post["PassWord"] == nil || req.Post["PassWord"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "PassWord 不能为空")
	} else {
		password = req.Post["PassWord"].Values[0]
	}

	if req.Post["DeviceDetector"] == nil || req.Post["DeviceDetector"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "DeviceDetector 不能为空")
	} else {
		deviceDetector = req.Post["DeviceDetector"].Values[0]
	}

	if req.Post["GeoRemoteAddr"] == nil || req.Post["GeoRemoteAddr"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "GeoRemoteAddr 不能为空")
	} else {
		geoRemoteAddr = req.Post["GeoRemoteAddr"].Values[0]
	}

	if req.Post["GeoCountry"] == nil || req.Post["GeoCountry"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "GeoCountry 不能为空")
	} else {
		geoCountry = req.Post["GeoCountry"].Values[0]
	}

	if req.Post["GeoCity"] == nil || req.Post["GeoCity"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "GeoCity 不能为空")
	} else {
		geoCity = req.Post["GeoCity"].Values[0]
	}
	// 获取请求参数 -------------------- 结束

	// 调用服务端方法获取用户
	response, err := e.Client.Get(ctx, &srvusers.GetRequest{
		UserName: userName,
	})
	if err != nil {
		if commonutils.StringContains(err.Error(), "no row found") {
			return errors.InternalServerError(namespaceID, "用户不存在")
		} else {
			return errors.InternalServerError(namespaceID, err.Error())
		}
	}

	if commonutils.HexEncodeToString(commonutils.GetHMAC(commonutils.HashSHA256, []byte(password), []byte(response.Model.AuthKey))) != response.Model.Password {
		return errors.InternalServerError(namespaceID, "密码不正确")
	}

	// 记录登录历史 -------------------- 开始
	srvhistoryuser := &srvusers.User{ID: response.Model.ID}
	if _, err := e.ClientHistory.Add(ctx, &srvhistoryuserlogin.AddRequest{
		User:           srvhistoryuser,
		DeviceDetector: deviceDetector,
		GeoRemoteAddr:  geoRemoteAddr,
		GeoCountry:     geoCountry,
		GeoCity:        geoCity,
	}); err != nil {
		return errors.InternalServerError(namespaceID, "ClientHistory.Add Error:%s", err.Error())
	}
	// 记录登录历史 -------------------- 结束

	// 写入 JWT
	timeDuration := time.Duration(time.Minute * 10) //10分钟后过期

	// 要写入 JWT 的加密参数
	mapClaims := jwt.MapClaims{
		"UserId":         strconv.FormatInt(response.Model.ID, 10),
		"UserName":       response.Model.UserName,
		"ExpirationDate": time.Now().Add(timeDuration).Format("2006-01-02 15:04:05"),
	}

	//如果加密失败，则返回错误
	if tokenString, err = encrypt.JWTEncrypt(mapClaims); err != nil {
		return errors.InternalServerError(namespaceID, "token.SignedString Error:%s", err.Error())
	} else {
		// 输出的 json
		responseJSON := struct {
			TokenString string `json:"tokenstring"` //返回的 Token
		}{}
		responseJSON.TokenString = tokenString
		b, _ := commonutils.JSONEncode(responseJSON)
		rsp.StatusCode = 200
		rsp.Body = string(b)
	}

	// debug
	if utils.RunMode == "dev" {
		inner.Mlogger.Info("rsp.Body", rsp.Body)
	}

	// 写入一个 jaeger span
	ctx, span = jaeger.StartSpan(ctx, "Api_User_UserLogin_End")
	if span != nil {
		defer span.Finish()
		span.SetTag("Token", tokenString)
	}

	return nil
}

// ValidToken 验证 Token
func (e *Users) ValidToken(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 注册原因：感觉验证Token就一个方法，写入jaeger没意义也不好看
	// 写入一个 jaeger span
	// ctx, span := jaeger.StartSpan(ctx, "Api_User_ValidToken_Begin")
	// if span != nil {
	// 	defer span.Finish()
	// }

	namespaceID := inner.NAMESPACE_MICROSERVICE_APIUSERS

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [Users][ValidToken] request", namespaceID)
	}

	// 获取请求参数 -------------------- 开始
	var tokenString string
	var updateExpirationDate bool // 是否更新token过期时间
	if req.Post["Token"] == nil || req.Post["Token"].Values[0] == "" {
		return errors.InternalServerError(namespaceID, "Token 不能为空")
	} else {
		tokenString = req.Post["Token"].Values[0]
	}

	//获取是否要更新新的过期时间 1或0
	if req.Post["UpdateExpirationDate"] != nil && req.Post["UpdateExpirationDate"].Values[0] == "1" {
		updateExpirationDate = true
	}
	// 获取请求参数 -------------------- 结束

	// 解析 Token
	var expirationDate time.Time

	// 解析 JWT 的 struct
	var mapClaims jwt.MapClaims

	// 如果解析失败，返回错误，可能编码被改变
	if mapClaims, err = encrypt.JWTDecrypt(tokenString); err != nil {
		return errors.InternalServerError(namespaceID, fmt.Sprintf("JWTDecrypt Error:%s", err.Error()))
	}

	// 打印解析后的值
	fmt.Println("mapClaims", mapClaims)

	// 如果日期字段解析错误，也返回
	if expirationDate, err = time.ParseInLocation("2006-01-02 15:04:05", mapClaims["ExpirationDate"].(string), time.Local); err != nil {
		return errors.InternalServerError(namespaceID, fmt.Sprintf("ExpirationDate ParseInLocation Error:%s", err.Error()))
	}

	// 输出的 json
	responseJSON := struct {
		Vaild       int    `json:"vaild"`       // 是否验证通过
		UserID      string `json:"userid"`      //用户ID
		UserName    string `json:"username"`    //用户登录名
		TokenString string `json:"tokenstring"` //返回的 Token
	}{}

	// 判断 token 是否过期
	if expirationDate.Before(time.Now()) {
		responseJSON.TokenString = "Token已过期"
		responseJSON.Vaild = 0
		b, _ := commonutils.JSONEncode(responseJSON)
		rsp.StatusCode = 200
		rsp.Body = string(b)
	} else {
		//如果没过期，重新返回原token
		responseJSON.TokenString = tokenString
		responseJSON.UserID = mapClaims["UserId"].(string)
		responseJSON.UserName = mapClaims["UserName"].(string)
		responseJSON.Vaild = 1
		b, _ := commonutils.JSONEncode(responseJSON)
		rsp.StatusCode = 200
		rsp.Body = string(b)
	}

	// 是否要更新Token的过期时间
	if updateExpirationDate {
		timeDuration := time.Duration(time.Minute * 10) //重新设置10分钟后过期
		mapClaims := jwt.MapClaims{
			"UserId":         mapClaims["UserId"].(string),
			"UserName":       mapClaims["UserName"].(string),
			"ExpirationDate": time.Now().Add(timeDuration),
		}
		if tokenString, err = encrypt.JWTEncrypt(mapClaims); err != nil {
			return errors.InternalServerError(namespaceID, "token.SignedString Error:%s", err.Error())
		} else {
			rsp.StatusCode = 200
			rsp.Body = tokenString
		}
	}

	// debug
	if utils.RunMode == "dev" {
		inner.Mlogger.Info("rsp.Body", rsp.Body)
	}

	// 写入一个 jaeger span
	// ctx, span = jaeger.StartSpan(ctx, "Api_User_ValidToken_End")
	// if span != nil {
	// 	defer span.Finish()
	// 	span.SetTag("Token", tokenString)
	// }

	return nil

}
