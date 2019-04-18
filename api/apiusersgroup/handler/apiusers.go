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

	srvhistoryuserlogin "github.com/idoall/MicroService-UserPowerManager/srv/srvhistoryuserlogin/proto"
	srvusers "github.com/idoall/MicroService-UserPowerManager/srv/srvusers/proto"

	"github.com/idoall/TokenExchangeCommon/commonutils/checkmail"
	api "github.com/micro/go-api/proto"
	"github.com/micro/go-micro/errors"
)

// Apiusers struct
type ApiUsers struct {
	ClientUser    srvusers.SrvUsersService
	ClientHistory srvhistoryuserlogin.SrvHistoryUserLoginService
}

// swagger:route POST /mshk/api/v1/ApiUsers/add users addPet
// 添加一个用户
func (e *ApiUsers) Add(ctx context.Context, req *api.Request, rsp *api.Response) error {

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_User_Add_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_id := inner.NAMESPACE_MICROSERVICE_API

	// debug
	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [Add] request", namespace_id)
	}

	// 获取请求参数 - 开始
	var username, password, email, realyname, note string
	if req.Post["UserName"] == nil || req.Post["UserName"].Values[0] == "" {
		return errors.InternalServerError(namespace_id, "UserName 不能为空")
	} else {
		username = req.Post["UserName"].Values[0]
	}

	if req.Post["PassWord"] == nil || req.Post["PassWord"].Values[0] == "" {
		return errors.InternalServerError(namespace_id, "PassWord 不能为空")
	} else {
		password = req.Post["PassWord"].Values[0]
	}

	if req.Post["Email"] == nil || req.Post["Email"].Values[0] == "" {
		return errors.InternalServerError(namespace_id, "Email 不能为空")
	} else if err := checkmail.ValidateFormat(req.Post["Email"].Values[0]); err != nil {
		return errors.InternalServerError(namespace_id, "Email 的格式不正确:%s", err.Error())
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
		return errors.BadRequest(namespace_id, "snowflake.NewNode Error:%+v", err)
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
	response, err := e.ClientUser.Add(ctx, &srvusers.AddRequest{Model: resUser})
	if err != nil {
		return errors.InternalServerError(namespace_id, err.Error())
	}

	// 输出的 json
	respJson := struct {
		NewUserId int64
	}{}
	respJson.NewUserId = response.NewUserId
	b, _ := commonutils.JSONEncode(respJson)
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
		span.SetTag("NewUserId", response.NewUserId)
	}

	return nil
}

// 获取用户列表,默认 id 倒排序
func (e *ApiUsers) GetList(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_User_GetList_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_id := inner.NAMESPACE_MICROSERVICE_API

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [GetList] request", namespace_id)
	}

	// 获取请求参数 - 开始
	var pageSize, currentPageIndex int64
	var orderBy string
	if req.Get["PageSize"] == nil || req.Get["PageSize"].Values[0] == "" {
		return errors.InternalServerError(namespace_id, "PageSize 不能为空")
	} else if pageSize, err = commonutils.Int64FromString(req.Get["PageSize"].Values[0]); err != nil {
		return errors.InternalServerError(namespace_id, "PageSize Format Error:%s", err.Error())
	}

	if req.Get["CurrentPageIndex"] == nil || req.Get["CurrentPageIndex"].Values[0] == "" {
		return errors.InternalServerError(namespace_id, "CurrentPageIndex 不能为空")
	} else if currentPageIndex, err = commonutils.Int64FromString(req.Get["CurrentPageIndex"].Values[0]); err != nil {
		return errors.InternalServerError(namespace_id, "CurrentPageIndex Format Error:%s", err.Error())

	}

	if req.Get["OrderBy"] != nil {
		orderBy = req.Get["OrderBy"].Values[0]
	}
	// 获取请求参数 - 结束

	// 调用服务端方法
	response, err := e.ClientUser.GetList(ctx, &srvusers.GetListRequest{
		CurrentPageIndex: currentPageIndex,
		PageSize:         pageSize,
		OrderBy:          orderBy,
	})
	if err != nil {
		return errors.InternalServerError(namespace_id, err.Error())
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
func (e *ApiUsers) GetUser(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_User_GetUser_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_id := inner.NAMESPACE_MICROSERVICE_API

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [GetUser] request", namespace_id)
	}

	// 获取请求参数 - 开始
	var userId int64
	var userName string
	if req.Get["UserId"] != nil && req.Get["UserId"].Values[0] != "0" {
		if userId, err = commonutils.Int64FromString(req.Get["UserId"].Values[0]); err != nil {
			return errors.InternalServerError(namespace_id, "UserId Format Error:%s", err.Error())
		}
	}

	if req.Get["UserName"] != nil && req.Get["UserName"].Values[0] != "" {
		userName = req.Get["UserName"].Values[0]
	}

	// 获取请求参数 - 结束

	// 调用服务端方法
	response, err := e.ClientUser.Get(ctx, &srvusers.GetRequest{
		UserId:   userId,
		UserName: userName,
	})
	if err != nil {
		return errors.InternalServerError(namespace_id, err.Error())
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
		span.SetTag("UserId", userId)
	}

	return nil
}

// 修改用户信息
func (e *ApiUsers) Update(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_User_Update_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_id := inner.NAMESPACE_MICROSERVICE_API

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [Update] request", namespace_id)
	}

	// 获取请求参数 - 开始
	var userId int64
	var username, password, email, realyname, note string
	if req.Post["UserId"] == nil || req.Post["UserId"].Values[0] == "" {
		return errors.InternalServerError(namespace_id, "UserId 不能为空")
	} else if userId, err = commonutils.Int64FromString(req.Post["UserId"].Values[0]); err != nil {
		return errors.InternalServerError(namespace_id, "UserId Format Error:%s", err.Error())
	}
	if req.Post["UserName"] == nil || req.Post["UserName"].Values[0] == "" {
		return errors.InternalServerError(namespace_id, "UserName 不能为空")
	} else {
		username = req.Post["UserName"].Values[0]
	}

	if req.Post["PassWord"] != nil || req.Post["PassWord"].Values[0] != "" {
		password = req.Post["PassWord"].Values[0]
	}

	if req.Post["Email"] == nil || req.Post["Email"].Values[0] == "" {
		return errors.InternalServerError(namespace_id, "Email 不能为空")
	} else if err := checkmail.ValidateFormat(req.Post["Email"].Values[0]); err != nil {
		return errors.InternalServerError(namespace_id, "Email 的格式不正确:%s", err.Error())
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
	responseGetUser, err := e.ClientUser.Get(ctx, &srvusers.GetRequest{
		UserId: userId,
	})
	if err != nil {
		return errors.InternalServerError(namespace_id, err.Error())
	}

	// 如果密码不为空，重新对密码加密
	if password != "" {
		password = commonutils.HexEncodeToString(commonutils.GetHMAC(commonutils.HashSHA256, []byte(password), []byte(responseGetUser.Model.AuthKey)))
	} else {
		password = responseGetUser.Model.Password
	}

	// 调用服务端方法 - 修改用户
	responseUpdateUser := &srvusers.User{
		Id:             userId,
		UserName:       username,
		Password:       password,
		RealyName:      realyname,
		AuthKey:        responseGetUser.Model.AuthKey,
		Email:          email,
		Note:           note,
		IsDel:          responseGetUser.Model.IsDel,
		ParentId:       responseGetUser.Model.ParentId,
		CreateTime:     responseGetUser.Model.CreateTime,
		LastUpdateTime: responseGetUser.Model.LastUpdateTime,
	}
	response, err := e.ClientUser.Update(ctx, &srvusers.UpdateRequest{Model: responseUpdateUser})
	if err != nil {
		return errors.InternalServerError(namespace_id, err.Error())
	}

	// 输出的 json
	respJson := struct {
		Updated int64
	}{}
	respJson.Updated = response.Updated
	b, _ := commonutils.JSONEncode(respJson)
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
		span.SetTag("UserId", userId)
	}

	return nil
}

// 批量删除用户信息
func (e *ApiUsers) BatchDelete(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_User_BatchDelete_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_id := inner.NAMESPACE_MICROSERVICE_API

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [BatchDelete] request", namespace_id)
	}

	// 获取请求参数 - 开始
	var userIdArray []string
	if req.Post["UserIds"] == nil || req.Post["UserIds"].Values[0] == "" {
		return errors.InternalServerError(namespace_id, "UserIds 不能为空")
	} else {
		userIdArray = strings.Split(req.Post["UserIds"].Values[0], ",")
	}

	fmt.Println("userIdArray", userIdArray)
	// 获取请求参数 - 结束

	// 调用服务端方法获取用户
	response, err := e.ClientUser.BatchDelete(ctx, &srvusers.DeleteRequest{
		UserIdArray: userIdArray,
	})
	if err != nil {
		return errors.InternalServerError(namespace_id, err.Error())
	}

	// 输出的 json
	respJson := struct {
		Deleted int64
	}{}
	respJson.Deleted = response.Deleted
	b, _ := commonutils.JSONEncode(respJson)
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
		span.SetTag("UserIdArray", strings.Join(userIdArray, ","))
	}

	return nil
}

// 用户登录
func (e *ApiUsers) UserLogin(ctx context.Context, req *api.Request, rsp *api.Response) error {
	var err error
	var tokenString string

	// 写入一个 jaeger span
	ctx, span := jaeger.StartSpan(ctx, "Api_User_UserLogin_Begin")
	if span != nil {
		defer span.Finish()
	}

	namespace_id := inner.NAMESPACE_MICROSERVICE_API

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [UserLogin] request", namespace_id)
	}

	// 获取请求参数 -------------------- 开始
	var userName, password, deviceDetector, geoRemoteAddr, geoCountry, geoCity string
	if req.Post["UserName"] == nil || req.Post["UserName"].Values[0] == "" {
		return errors.InternalServerError(namespace_id, "UserName 不能为空")
	} else {
		userName = req.Post["UserName"].Values[0]
	}

	if req.Post["PassWord"] == nil || req.Post["PassWord"].Values[0] == "" {
		return errors.InternalServerError(namespace_id, "PassWord 不能为空")
	} else {
		password = req.Post["PassWord"].Values[0]
	}

	if req.Post["DeviceDetector"] == nil || req.Post["DeviceDetector"].Values[0] == "" {
		return errors.InternalServerError(namespace_id, "DeviceDetector 不能为空")
	} else {
		deviceDetector = req.Post["DeviceDetector"].Values[0]
	}

	if req.Post["GeoRemoteAddr"] == nil || req.Post["GeoRemoteAddr"].Values[0] == "" {
		return errors.InternalServerError(namespace_id, "GeoRemoteAddr 不能为空")
	} else {
		geoRemoteAddr = req.Post["GeoRemoteAddr"].Values[0]
	}

	if req.Post["GeoCountry"] == nil || req.Post["GeoCountry"].Values[0] == "" {
		return errors.InternalServerError(namespace_id, "GeoCountry 不能为空")
	} else {
		geoCountry = req.Post["GeoCountry"].Values[0]
	}

	if req.Post["GeoCity"] == nil || req.Post["GeoCity"].Values[0] == "" {
		return errors.InternalServerError(namespace_id, "GeoCity 不能为空")
	} else {
		geoCity = req.Post["GeoCity"].Values[0]
	}
	// 获取请求参数 -------------------- 结束

	// 调用服务端方法获取用户
	response, err := e.ClientUser.Get(ctx, &srvusers.GetRequest{
		UserName: userName,
	})
	if err != nil {
		if commonutils.StringContains(err.Error(), "no row found") {
			return errors.InternalServerError(namespace_id, "用户不存在")
		} else {
			return errors.InternalServerError(namespace_id, err.Error())
		}
	}

	if commonutils.HexEncodeToString(commonutils.GetHMAC(commonutils.HashSHA256, []byte(password), []byte(response.Model.AuthKey))) != response.Model.Password {
		return errors.InternalServerError(namespace_id, "密码不正确")
	}

	// 记录登录历史 -------------------- 开始
	srvhistoryuser := &srvusers.User{Id: response.Model.Id}
	if _, err := e.ClientHistory.Add(ctx, &srvhistoryuserlogin.AddRequest{
		User:           srvhistoryuser,
		DeviceDetector: deviceDetector,
		GeoRemoteAddr:  geoRemoteAddr,
		GeoCountry:     geoCountry,
		GeoCity:        geoCity,
	}); err != nil {
		return errors.InternalServerError(namespace_id, "ClientHistory.Add Error:%s", err.Error())
	}
	// 记录登录历史 -------------------- 结束

	// 写入 JWT
	timeDuration := time.Duration(time.Minute * 10) //10分钟后过期
	mapClaims := jwt.MapClaims{
		"UserId":         strconv.FormatInt(response.Model.Id, 10),
		"ExpirationDate": time.Now().Add(timeDuration).Format("2006-01-02 15:04:05"),
	}
	if tokenString, err = encrypt.JWTEncrypt(mapClaims); err != nil {
		return errors.InternalServerError(namespace_id, "token.SignedString Error:%s", err.Error())
	} else {
		// 输出的 json
		respJson := struct {
			TokenString string `json:"tokenstring"` //返回的 Token
		}{}
		respJson.TokenString = tokenString
		b, _ := commonutils.JSONEncode(respJson)
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

// 验证 Token
func (e *ApiUsers) ValidToken(ctx context.Context, req *api.Request, rsp *api.Response) error {

	var err error
	// 写入一个 jaeger span
	// ctx, span := jaeger.StartSpan(ctx, "Api_User_ValidToken_Begin")
	// if span != nil {
	// 	defer span.Finish()
	// }

	namespace_id := inner.NAMESPACE_MICROSERVICE_API

	if utils.RunMode == "dev" {
		inner.Mlogger.Infof("Received %s API [ValidToken] request", namespace_id)
	}

	// 获取请求参数 -------------------- 开始
	var tokenString string
	var updateExpirationDate bool // 是否更新token过期时间
	if req.Post["Token"] == nil || req.Post["Token"].Values[0] == "" {
		return errors.InternalServerError(namespace_id, "Token 不能为空")
	} else {
		tokenString = req.Post["Token"].Values[0]
	}
	if req.Post["UpdateExpirationDate"] != nil && req.Post["UpdateExpirationDate"].Values[0] == "1" {
		updateExpirationDate = true
	}
	// 获取请求参数 -------------------- 结束

	// 解析 Token
	var expirationDate time.Time
	var mapClaims jwt.MapClaims
	if mapClaims, err = encrypt.JWTDecrypt(tokenString); err != nil {
		return errors.InternalServerError(namespace_id, fmt.Sprintf("JWTDecrypt Error:%s", err.Error()))
	}
	fmt.Println("mapClaims", mapClaims)
	if expirationDate, err = time.ParseInLocation("2006-01-02 15:04:05", mapClaims["ExpirationDate"].(string), time.Local); err != nil {
		return errors.InternalServerError(namespace_id, fmt.Sprintf("ExpirationDate ParseInLocation Error:%s", err.Error()))
	}

	// 输出的 json
	respJson := struct {
		Vaild       int    // 是否验证通过
		TokenString string //返回的 Token
	}{}

	// 判断 token 是否过期
	if expirationDate.Before(time.Now()) {
		respJson.TokenString = "Token已过期"
		respJson.Vaild = 0
		b, _ := commonutils.JSONEncode(respJson)
		rsp.StatusCode = 200
		rsp.Body = string(b)
	} else {
		//如果没过期，重新返回原token
		respJson.TokenString = tokenString
		respJson.Vaild = 1
		b, _ := commonutils.JSONEncode(respJson)
		rsp.StatusCode = 200
		rsp.Body = string(b)
	}

	// 是否要更新Token的过期时间
	if updateExpirationDate {
		timeDuration := time.Duration(time.Minute * 10) //10分钟后过期
		mapClaims := jwt.MapClaims{
			"UserId":         mapClaims["UserId"].(string),
			"ExpirationDate": time.Now().Add(timeDuration),
		}
		if tokenString, err = encrypt.JWTEncrypt(mapClaims); err != nil {
			return errors.InternalServerError(namespace_id, "token.SignedString Error:%s", err.Error())
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
