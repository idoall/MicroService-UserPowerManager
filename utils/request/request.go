package request

import (
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"

	"github.com/idoall/MicroService-UserPowerManager/utils"
	"github.com/idoall/MicroService-UserPowerManager/utils/inner"
	"github.com/idoall/TokenExchangeCommon/commonutils"
	"github.com/idoall/gocryptotrader/common"
)

var (
	supportedMethods = []string{http.MethodGet, http.MethodPost, http.MethodHead,
		http.MethodPut, http.MethodDelete, http.MethodOptions, http.MethodConnect}
	Request *Requester
)

const (
	maxRequestJobs              = 50
	proxyTLSTimeout             = 15 * time.Second
	defaultTimeoutRetryAttempts = 10
)

// Requester struct for the request client
type Requester struct {
	HTTPClient           *http.Client
	UnauthLimit          *RateLimit
	AuthLimit            *RateLimit
	Name                 string
	UserAgent            string
	Cycle                time.Time
	timeoutRetryAttempts int
	m                    sync.Mutex
	Jobs                 chan Job
	disengage            chan struct{}
	WorkerStarted        bool
	fifoLock             sync.Mutex
}

// RateLimit struct
type RateLimit struct {
	Duration time.Duration
	Rate     int
	Requests int
	Mutex    sync.Mutex
}

// JobResult holds a request job result
type JobResult struct {
	Error  error
	Result interface{}
}

// Job holds a request job
type Job struct {
	Request       *http.Request
	Method        string
	Path          string
	Headers       map[string]string
	Body          io.Reader
	Result        interface{}
	JobResult     chan *JobResult
	AuthRequest   bool
	Verbose       bool
	HTTPDebugging bool
}

// NewRateLimit creates a new RateLimit
func NewRateLimit(d time.Duration, rate int) *RateLimit {
	return &RateLimit{Duration: d, Rate: rate}
}

// ToString returns the rate limiter in string notation
func (r *RateLimit) ToString() string {
	return fmt.Sprintf("Rate limiter set to %d requests per %v", r.Rate, r.Duration)
}

// GetRate returns the ratelimit rate
func (r *RateLimit) GetRate() int {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	return r.Rate
}

// SetRate sets the ratelimit rate
func (r *RateLimit) SetRate(rate int) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	r.Rate = rate
}

// GetRequests returns the number of requests for the ratelimit
func (r *RateLimit) GetRequests() int {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	return r.Requests
}

// SetRequests sets requests counter for the rateliit
func (r *RateLimit) SetRequests(l int) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	r.Requests = l
}

// SetDuration sets the duration for the ratelimit
func (r *RateLimit) SetDuration(d time.Duration) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	r.Duration = d
}

// GetDuration gets the duration for the ratelimit
func (r *RateLimit) GetDuration() time.Duration {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	return r.Duration
}

// StartCycle restarts the cycle time and requests counters
func (r *Requester) StartCycle() {
	r.Cycle = time.Now()
	r.AuthLimit.SetRequests(0)
	r.UnauthLimit.SetRequests(0)
}

// IsRateLimited returns whether or not the request Requester is rate limited
func (r *Requester) IsRateLimited(auth bool) bool {
	if auth {
		// fmt.Println(r.AuthLimit.GetRequests())
		if r.AuthLimit.GetRequests() >= r.AuthLimit.GetRate() && r.IsValidCycle(auth) {
			return true
		}
	} else {
		if r.UnauthLimit.GetRequests() >= r.UnauthLimit.GetRate() && r.IsValidCycle(auth) {
			return true
		}
	}
	return false
}

// RequiresRateLimiter returns whether or not the request Requester requires a rate limiter
func (r *Requester) RequiresRateLimiter() bool {
	if r.AuthLimit.GetRate() != 0 || r.UnauthLimit.GetRate() != 0 {
		return true
	}
	return false
}

// IncrementRequests increments the ratelimiter request counter for either auth or unauth
// requests 增量计数器
func (r *Requester) IncrementRequests(auth bool) {
	if auth {
		reqs := r.AuthLimit.GetRequests()
		reqs++
		r.AuthLimit.SetRequests(reqs)
		return
	}

	reqs := r.UnauthLimit.GetRequests()
	reqs++
	r.UnauthLimit.SetRequests(reqs)
}

// DecrementRequests decrements the ratelimiter request counter for either auth or unauth
// requests 减量计数器
func (r *Requester) DecrementRequests(auth bool) {
	if auth {

		reqs := r.AuthLimit.GetRequests()
		reqs--
		r.AuthLimit.SetRequests(reqs)
		return
	}

	reqs := r.AuthLimit.GetRequests()
	reqs--
	r.UnauthLimit.SetRequests(reqs)
}

// SetRateLimit sets the request Requester ratelimiter
func (r *Requester) SetRateLimit(auth bool, duration time.Duration, rate int) {
	if auth {
		r.AuthLimit.SetRate(rate)
		r.AuthLimit.SetDuration(duration)
		return
	}
	r.UnauthLimit.SetRate(rate)
	r.UnauthLimit.SetDuration(duration)
}

// GetRateLimit gets the request Requester ratelimiter
func (r *Requester) GetRateLimit(auth bool) *RateLimit {
	if auth {
		return r.AuthLimit
	}
	return r.UnauthLimit
}

// SetTimeoutRetryAttempts sets the amount of times the job will be retried
// if it times out
func (r *Requester) SetTimeoutRetryAttempts(n int) error {
	if n < 0 {
		return errors.New("routines.go error - timeout retry attempts cannot be less than zero")
	}
	r.timeoutRetryAttempts = n
	return nil
}

// New returns a new Requester
func New(name string, authLimit, unauthLimit *RateLimit, httpRequester *http.Client) *Requester {
	return &Requester{
		HTTPClient:           httpRequester,
		UnauthLimit:          unauthLimit,
		AuthLimit:            authLimit,
		Name:                 name,
		Jobs:                 make(chan Job, maxRequestJobs),
		disengage:            make(chan struct{}, 1),
		timeoutRetryAttempts: defaultTimeoutRetryAttempts,
	}
}

// IsValidMethod returns whether the supplied method is supported
func IsValidMethod(method string) bool {
	return commonutils.StringDataCompareInsensitive(supportedMethods, method)
}

// IsValidCycle checks to see whether the current request cycle is valid or not
func (r *Requester) IsValidCycle(auth bool) bool {
	if auth {
		if time.Since(r.Cycle) < r.AuthLimit.GetDuration() {
			return true
		}
	} else {
		if time.Since(r.Cycle) < r.UnauthLimit.GetDuration() {
			return true
		}
	}

	r.StartCycle()
	return false
}

func (r *Requester) checkRequest(method, path string, body io.Reader, headers map[string]string) (*http.Request, error) {
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

	return req, nil
}

// DoRequest performs a HTTP/HTTPS request with the supplied params
func (r *Requester) DoRequest(req *http.Request, path string, body io.Reader, result interface{}, authRequest, verbose, httpDebug bool) error {
	if verbose {
		inner.Mlogger.Debugf("%s request path: %s requires rate limiter: %v", r.Name, path, r.RequiresRateLimiter())
		for k, d := range req.Header {
			inner.Mlogger.Debugf("%s request header [%s]: %s", r.Name, k, d)
		}
		inner.Mlogger.Debug(body)
	}

	var timeoutError error
	for i := 0; i < r.timeoutRetryAttempts+1; i++ {
		// 发送请求
		resp, err := r.HTTPClient.Do(req)
		if err != nil {
			if timeoutErr, ok := err.(net.Error); ok && timeoutErr.Timeout() {
				// if verbose {
				inner.Mlogger.Errorf("%s request has timed-out retrying request, count %d",
					r.Name,
					i)
				// }
				timeoutError = err
				continue
			}

			if r.RequiresRateLimiter() {
				r.DecrementRequests(authRequest)
			}
			return err
		}
		if resp == nil {
			if r.RequiresRateLimiter() {
				r.DecrementRequests(authRequest)
			}
			return errors.New("resp is nil")
		}

		var reader io.ReadCloser
		switch resp.Header.Get("Content-Encoding") {
		case "gzip":
			reader, err = gzip.NewReader(resp.Body)
			defer reader.Close()
			if err != nil {
				return err
			}

		case "json":
			reader = resp.Body

		default:
			switch {
			case commonutils.StringContains(resp.Header.Get("Content-Type"), "application/json"):
				reader = resp.Body

			default:
				inner.Mlogger.Warningf("%s request response content type differs from JSON; received %v [path: %s]",
					r.Name, resp.Header.Get("Content-Type"), path)
				reader = resp.Body
			}
		}

		contents, err := ioutil.ReadAll(reader)
		if err != nil {
			return err
		}

		if resp.StatusCode != 200 && resp.StatusCode != 201 && resp.StatusCode != 202 {
			err = fmt.Errorf("unsuccessful HTTP status code: %d %s", resp.StatusCode, string(contents))
			if verbose {
				err = fmt.Errorf("%s\n%s", err.Error(),
					fmt.Sprintf("%s exchange raw response: %s", r.Name, string(contents)))
			}

			return err
		}

		if httpDebug {
			dump, err := httputil.DumpResponse(resp, false)
			if err != nil {
				inner.Mlogger.Errorf("DumpResponse invalid response: %v:", err)
			}
			inner.Mlogger.Debugf("DumpResponse Headers (%v):\n%s", path, dump)
			inner.Mlogger.Debugf("DumpResponse Body (%v):\n %s", path, string(contents))
		}

		resp.Body.Close()
		if verbose {
			inner.Mlogger.Debugf("HTTP status: %s, Code: %v", resp.Status, resp.StatusCode)
			if !httpDebug {
				inner.Mlogger.Debugf("%s exchange raw response: %s", r.Name, string(contents))
			}
		}

		if result != nil {
			return common.JSONDecode(contents, result)
		}

		return nil
	}
	return fmt.Errorf("request.go error - failed to retry request %s",
		timeoutError)
}

func (r *Requester) worker() {
	for {
		for x := range r.Jobs {
			if !r.IsRateLimited(x.AuthRequest) {
				r.IncrementRequests(x.AuthRequest)

				err := r.DoRequest(x.Request, x.Path, x.Body, x.Result, x.AuthRequest, x.Verbose, x.HTTPDebugging)
				x.JobResult <- &JobResult{
					Error:  err,
					Result: x.Result,
				}
			} else {
				limit := r.GetRateLimit(x.AuthRequest)
				diff := limit.GetDuration() - time.Since(r.Cycle)
				if x.Verbose {
					inner.Mlogger.Debugf("%s request. Rate limited! Sleeping for %v", r.Name, diff)
				}
				time.Sleep(diff)

				for {
					if r.IsRateLimited(x.AuthRequest) {
						time.Sleep(time.Millisecond)
						continue
					}
					r.IncrementRequests(x.AuthRequest)

					if x.Verbose {
						inner.Mlogger.Debugf("%s request. No longer rate limited! Doing request", r.Name)
					}

					err := r.DoRequest(x.Request, x.Path, x.Body, x.Result, x.AuthRequest, x.Verbose, x.HTTPDebugging)
					x.JobResult <- &JobResult{
						Error:  err,
						Result: x.Result,
					}
					break
				}
			}
		}
	}
}

// WebPOSTSendPayload 封装统一请求微服务的 POST 方法
func (r *Requester) WebPOSTSendPayload(configParam string, body io.Reader, result interface{}, authRequest, nonceEnabled, verbose, httpDebugging bool) error {
	path := fmt.Sprintf("%s%s", inner.MicroServiceHostProt, utils.TConfig.String("MicroServices::"+configParam))
	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	return r.SendPayload("POST", path, headers, body, result, authRequest, nonceEnabled, verbose, httpDebugging)
}

// WebGETSendPayload 封装统一请求微服务的 GET 方法
func (r *Requester) WebGETSendPayload(configParam string, params url.Values, result interface{}, authRequest, nonceEnabled, verbose, httpDebugging bool) error {
	// 发送请求的路径
	path := fmt.Sprintf("%s%s?%s",
		inner.MicroServiceHostProt,
		utils.TConfig.String("MicroServices::"+configParam),
		params.Encode(),
	)

	return r.SendPayload("GET", path, nil, nil, result, authRequest, nonceEnabled, verbose, httpDebugging)
}

// SendPayload handles sending HTTP/HTTPS requests
func (r *Requester) SendPayload(method, path string, headers map[string]string, body io.Reader, result interface{}, authRequest, nonceEnabled, verbose, httpDebugging bool) error {
	if !nonceEnabled {
		r.lock()
	}

	if r == nil || r.Name == "" {
		r.unlock()
		return errors.New("not initiliased, SetDefaults() called before making request?")
	}

	// 验证请求方法的合法性
	if !IsValidMethod(method) {
		r.unlock()
		return fmt.Errorf("incorrect method supplied %s: supported %s", method, supportedMethods)
	}

	// 验证请求路径
	if path == "" {
		r.unlock()
		return errors.New("invalid path")
	}

	// 设置 header
	req, err := r.checkRequest(method, path, body, headers)
	if err != nil {
		r.unlock()
		return err
	}

	if httpDebugging {
		dump, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			inner.Mlogger.Errorf("DumpRequest invalid response %v:", err)
		}
		inner.Mlogger.Debugf("DumpRequest:\n%s", dump)
	}

	if !r.RequiresRateLimiter() {
		r.unlock()
		return r.DoRequest(req, path, body, result, authRequest, verbose, httpDebugging)
	}

	if len(r.Jobs) == maxRequestJobs {
		r.unlock()
		return errors.New("max request jobs reached")
	}

	r.m.Lock()
	if !r.WorkerStarted {
		r.StartCycle()
		r.WorkerStarted = true
		go r.worker()
	}
	r.m.Unlock()

	jobResult := make(chan *JobResult)

	newJob := Job{
		Request:       req,
		Method:        method,
		Path:          path,
		Headers:       headers,
		Body:          body,
		Result:        result,
		JobResult:     jobResult,
		AuthRequest:   authRequest,
		Verbose:       verbose,
		HTTPDebugging: httpDebugging,
	}

	if verbose {
		inner.Mlogger.Debugf("%s request. Attaching new job.", r.Name)
	}
	r.Jobs <- newJob
	r.unlock()

	if verbose {
		inner.Mlogger.Debugf("%s request. Waiting for job to complete.", r.Name)
	}
	resp := <-newJob.JobResult

	if verbose {
		inner.Mlogger.Debugf("%s request. Job complete.", r.Name)
	}

	return resp.Error
}

// SetProxy sets a proxy address to the client transport
func (r *Requester) SetProxy(p *url.URL) error {
	if p.String() == "" {
		return errors.New("no proxy URL supplied")
	}

	r.HTTPClient.Transport = &http.Transport{
		Proxy:               http.ProxyURL(p),
		TLSHandshakeTimeout: proxyTLSTimeout,
	}
	return nil
}

// lock 锁定并设置问题计时器，如果出现超出范围的错误自动解锁
func (r *Requester) lock() {
	if r.disengage == nil {
		r.disengage = make(chan struct{}, 1)
	}
	var wg sync.WaitGroup
	r.fifoLock.Lock()
	wg.Add(1)
	go func() {
		timer := time.NewTimer(50 * time.Millisecond)
		wg.Done()
		select {
		case <-timer.C:
			inner.Mlogger.Error("由于可能的错误而解锁:" + r.Name)
			r.fifoLock.Unlock()

		case <-r.disengage:
			return
		}
	}()
	wg.Wait()
}

// 解锁解锁MTX并关闭计时器
func (r *Requester) unlock() {
	r.disengage <- struct{}{}
	r.fifoLock.Unlock()
}
