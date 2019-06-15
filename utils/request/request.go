package request

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/idoall/MicroService-UserPowerManager/utils"
	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
	"github.com/idoall/TokenExchangeCommon/commonutils"
)

var Request *Requester

var supportedMethods = []string{"GET", "POST", "HEAD", "PUT", "DELETE", "OPTIONS", "CONNECT"}

// Requester struct for the request client
type Requester struct {
	HTTPClient *http.Client
	Name       string
	UserAgent  string
}

// New returns a new Requester
func New(name string, httpRequester *http.Client) *Requester {
	return &Requester{
		HTTPClient: httpRequester,
		Name:       name,
	}
}

// IsValidMethod returns whether the supplied method is supported
func IsValidMethod(method string) bool {
	return commonutils.StringDataCompareUpper(supportedMethods, method)
}

// WebPOSTSendPayload 封装统一请求微服务的 POST 方法
func (r *Requester) WebPOSTSendPayload(configParam string, body io.Reader, result interface{}) error {
	path := fmt.Sprintf("%s%s", inner.MicroServiceHostProt, utils.TConfig.String("MicroServices::"+configParam))
	return r.SendPayload("POST", path, nil, body, result, false, true, false)
}

// WebGETSendPayload 封装统一请求微服务的 GET 方法
func (r *Requester) WebGETSendPayload(configParam string, params url.Values, result interface{}) error {
	// 发送请求的路径
	path := fmt.Sprintf("%s%s?%s",
		inner.MicroServiceHostProt,
		utils.TConfig.String("MicroServices::"+configParam),
		params.Encode(),
	)
	return r.SendPayload("GET", path, nil, nil, result, false, true, false)
}

// Set header
func (r *Requester) checkRequest(method, path string, body io.Reader, headers map[string]string, debug bool) (*http.Request, error) {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	if r.UserAgent != "" && req.Header.Get("User-Agent") == "" {
		req.Header.Add("User-Agent", r.UserAgent)
	}

	if method == "POST" && req.Header.Get("Content-Type") == "" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	if debug {
		r.debug(httputil.DumpRequestOut(req, true))
	}
	return req, nil
}

func (r *Requester) debug(data []byte, err error) {
	if err == nil {
		inner.Mlogger.Warningf("%s\n\n", data)
	} else {
		inner.Mlogger.Fatalf("%s\n\n", err)
	}
}

// DoRequest performs a HTTP/HTTPS request with the supplied params
func (r *Requester) DoRequest(req *http.Request, method, path string, headers map[string]string, body io.Reader, result interface{}, authRequest, verbose, debug bool) error {
	if verbose {
		inner.Mlogger.Warning(fmt.Sprintf("%s request path: %s", r.Name, path))
	}

	// 发送请求
	resp, err := r.HTTPClient.Do(req)

	if err != nil {
		return err
	} else if debug {
		r.debug(httputil.DumpRequestOut(req, true))
	}

	if resp == nil {
		return errors.New("resp is nil")
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	resp.Body.Close()
	if verbose {
		inner.Mlogger.Warningf("%s raw response: %s", r.Name, string(contents[:]))
	}
	if debug {
		r.debug(httputil.DumpRequestOut(req, true))
	}

	if resp.StatusCode == 500 {
		var rawResponse map[string]interface{}
		if err := commonutils.JSONDecode(contents, &rawResponse); err != nil {
			return fmt.Errorf("commonutils.JSONDecode Err:%s", err.Error())
		} else {
			return fmt.Errorf("StatusCode:%d %s %s", resp.StatusCode, rawResponse["detail"].(string), rawResponse["status"].(string))
		}
	} else if resp.StatusCode == 200 {
		if result != nil {
			err := commonutils.JSONDecode(contents, result)
			if err != nil {
				return fmt.Errorf("Err:%s, Content:%s", err.Error(), string(contents))
			}
		}
	} else {
		return fmt.Errorf("[%d]未知错误:%s", resp.StatusCode, string(contents))
	}

	return nil
}

// SendPayload handles sending HTTP/HTTPS requests
func (r *Requester) SendPayload(method, path string, headers map[string]string, body io.Reader, result interface{}, authRequest, verbose, debug bool) error {
	if r == nil || r.Name == "" {
		return errors.New("not initiliased, Name?")
	}

	// 验证请求方法的合法性
	if !IsValidMethod(method) {
		return fmt.Errorf("incorrect method supplied %s: supported %s", method, supportedMethods)
	}

	// 验证请求路径
	if path == "" {
		return errors.New("invalid path")
	}

	// 设置 header
	req, err := r.checkRequest(method, path, body, headers, debug)
	if err != nil {
		return err
	}

	return r.DoRequest(req, method, path, headers, body, result, authRequest, verbose, debug)
}
