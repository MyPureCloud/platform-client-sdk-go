package platformclientv2

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/go-retryablehttp"
)

var (
	// Default retry configuration
	defaultRetryWaitMin = 1 * time.Second
	defaultRetryWaitMax = 30 * time.Second
	defaultRetryMax     = 4

	// defaultLogger is the logger provided with defaultClient
	defaultLogger = log.New(os.Stderr, "", log.LstdFlags)

	// defaultClient is used for performing requests without explicitly making
	// a new client. It is purposely private to avoid modifications.
	defaultClient = NewClient()

	// We need to consume response bodies to maintain http connections, but
	// limit the size we consume to respReadLimit.
	respReadLimit = int64(4096)
)

// RetryableHTTPClient is an extension of the retryablehttp.Client
type RetryableHTTPClient struct {
	*retryablehttp.Client
	loggerInit sync.Once
	clientInit sync.Once
}

// Do wraps calling an HTTP method with retries.
func (c *RetryableHTTPClient) Do(req *Request) (*http.Response, error) {
	c.clientInit.Do(func() {
		if c.HTTPClient == nil {
			c.HTTPClient = cleanhttp.DefaultPooledClient()
		}
	})

	logger := c.logger()

	if logger != nil {
		switch v := logger.(type) {
		case retryablehttp.LeveledLogger:
			v.Debug("performing request", "method", req.Method, "url", req.URL)
		case retryablehttp.Logger:
			v.Printf("[DEBUG] %s %s", req.Method, req.URL)
		}
	}

	var resp *http.Response
	var attempt int
	var shouldRetry bool
	var doErr, checkErr error

	for i := 0; ; i++ {
		attempt++

		var code int // HTTP response code

		// Always rewind the request body when non-nil.
		if req.body != nil {
			body, err := req.body()
			if err != nil {
				c.HTTPClient.CloseIdleConnections()
				return resp, err
			}
			if c, ok := body.(io.ReadCloser); ok {
				req.Body = c
			} else {
				req.Body = ioutil.NopCloser(body)
			}
		}

		if c.RequestLogHook != nil {
			switch v := logger.(type) {
			case retryablehttp.LeveledLogger:
				c.RequestLogHook(hookLogger{v}, req.Request, i)
			case retryablehttp.Logger:
				c.RequestLogHook(v, req.Request, i)
			default:
				c.RequestLogHook(nil, req.Request, i)
			}
		}

		// Attempt the request
		resp, doErr = c.HTTPClient.Do(req.Request)
		if resp != nil {
			code = resp.StatusCode
		}

		// Check if we should continue with retries.
		shouldRetry, checkErr = c.CheckRetry(req.Context(), resp, doErr)

		if doErr != nil {
			switch v := logger.(type) {
			case retryablehttp.LeveledLogger:
				v.Error("request failed", "error", doErr, "method", req.Method, "url", req.URL)
			case retryablehttp.Logger:
				v.Printf("[ERR] %s %s request failed: %v", req.Method, req.URL, doErr)
			}
		} else {
			// Call this here to maintain the behavior of logging all requests,
			// even if CheckRetry signals to stop.
			if c.ResponseLogHook != nil {
				// Call the response logger function if provided.
				switch v := logger.(type) {
				case retryablehttp.LeveledLogger:
					c.ResponseLogHook(hookLogger{v}, resp)
				case retryablehttp.Logger:
					c.ResponseLogHook(v, resp)
				default:
					c.ResponseLogHook(nil, resp)
				}
			}
		}

		if !shouldRetry {
			break
		}

		// We do this before drainBody because there's no need for the I/O if
		// we're breaking out
		remain := c.RetryMax - i
		if remain <= 0 {
			break
		}

		// We're going to retry, consume any response to reuse the connection.
		if doErr == nil {
			c.drainBody(resp.Body)
		}

		wait := c.Backoff(c.RetryWaitMin, c.RetryWaitMax, i, resp)
		desc := fmt.Sprintf("%s %s", req.Method, req.URL)
		if code > 0 {
			desc = fmt.Sprintf("%s (status: %d)", desc, code)
		}
		if logger != nil {
			switch v := logger.(type) {
			case retryablehttp.LeveledLogger:
				v.Debug("retrying request", "request", desc, "timeout", wait, "remaining", remain)
			case retryablehttp.Logger:
				v.Printf("[DEBUG] %s: retrying in %s (%d left)", desc, wait, remain)
			}
		}
		select {
		case <-req.Context().Done():
			c.HTTPClient.CloseIdleConnections()
			return nil, req.Context().Err()
		case <-time.After(wait):
		}

		// Make shallow copy of http Request so that we can modify its body
		// without racing against the closeBody call in persistConn.writeLoop.
		httpreq := *req.Request
		req.Request = &httpreq
	}

	// this is the closest we have to success criteria
	if doErr == nil && checkErr == nil {
		return resp, nil
	}

	defer c.HTTPClient.CloseIdleConnections()

	err := doErr
	if checkErr != nil {
		err = checkErr
	}

	if c.ErrorHandler != nil {
		return c.ErrorHandler(resp, err, attempt)
	}

	// By default, we close the response body and return an error without
	// returning the response
	if resp != nil {
		c.drainBody(resp.Body)
	}

	// this means CheckRetry thought the request was a failure, but didn't
	// communicate why
	if err == nil {
		return nil, fmt.Errorf("%s %s giving up after %d attempt(s)",
			req.Method, req.URL, attempt)
	}

	return nil, fmt.Errorf("%s %s giving up after %d attempt(s): %w",
		req.Method, req.URL, attempt, err)
}

// hookLogger adapts an LeveledLogger to Logger for use by the existing hook functions
// without changing the API.
type hookLogger struct {
	retryablehttp.LeveledLogger
}

func (h hookLogger) Printf(s string, args ...interface{}) {
	h.Info(fmt.Sprintf(s, args...))
}

// Request wraps the metadata needed to create HTTP requests.
type Request struct {
	// body is a seekable reader over the request body payload. This is
	// used to rewind the request data in between retries.
	body retryablehttp.ReaderFunc

	// Embed an HTTP request directly. This makes a *Request act exactly
	// like an *http.Request so that all meta methods are supported.
	*http.Request
}

// BodyBytes allows accessing the request body. It is an analogue to
// http.Request's Body variable, but it returns a copy of the underlying data
// rather than consuming it.
//
// This function is not thread-safe; do not call it at the same time as another
// call, or at the same time this request is being used with Client.Do.
func (r *Request) BodyBytes() ([]byte, error) {
	if r.body == nil {
		return nil, nil
	}
	body, err := r.body()
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(body)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// SetBody allows setting the request body.
//
// It is useful if a new body needs to be set without constructing a new Request.
func (r *Request) SetBody(rawBody interface{}) error {
	bodyReader, contentLength, err := getBodyReaderAndContentLength(rawBody)
	if err != nil {
		return err
	}
	r.body = bodyReader
	r.ContentLength = contentLength
	return nil
}

// WriteTo allows copying the request body into a writer.
//
// It writes data to w until there's no more data to write or
// when an error occurs. The return int64 value is the number of bytes
// written. Any error encountered during the write is also returned.
// The signature matches io.WriterTo interface.
func (r *Request) WriteTo(w io.Writer) (int64, error) {
	body, err := r.body()
	if err != nil {
		return 0, err
	}
	if c, ok := body.(io.Closer); ok {
		defer c.Close()
	}
	return io.Copy(w, body)
}

// NewRequest creates a new wrapped request.
func NewRequest(method, url string, rawBody interface{}) (*Request, error) {
	bodyReader, contentLength, err := getBodyReaderAndContentLength(rawBody)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	httpReq.ContentLength = contentLength

	return &Request{bodyReader, httpReq}, nil
}

func getBodyReaderAndContentLength(rawBody interface{}) (retryablehttp.ReaderFunc, int64, error) {
	var bodyReader retryablehttp.ReaderFunc
	var contentLength int64

	switch body := rawBody.(type) {
	// If they gave us a function already, great! Use it.
	case retryablehttp.ReaderFunc:
		bodyReader = body
		tmp, err := body()
		if err != nil {
			return nil, 0, err
		}
		if lr, ok := tmp.(retryablehttp.LenReader); ok {
			contentLength = int64(lr.Len())
		}
		if c, ok := tmp.(io.Closer); ok {
			c.Close()
		}

	case func() (io.Reader, error):
		bodyReader = body
		tmp, err := body()
		if err != nil {
			return nil, 0, err
		}
		if lr, ok := tmp.(retryablehttp.LenReader); ok {
			contentLength = int64(lr.Len())
		}
		if c, ok := tmp.(io.Closer); ok {
			c.Close()
		}

	// If a regular byte slice, we can read it over and over via new
	// readers
	case []byte:
		buf := body
		bodyReader = func() (io.Reader, error) {
			return bytes.NewReader(buf), nil
		}
		contentLength = int64(len(buf))

	// If a bytes.Buffer we can read the underlying byte slice over and
	// over
	case *bytes.Buffer:
		buf := body
		bodyReader = func() (io.Reader, error) {
			return bytes.NewReader(buf.Bytes()), nil
		}
		contentLength = int64(buf.Len())

	// We prioritize *bytes.Reader here because we don't really want to
	// deal with it seeking so want it to match here instead of the
	// io.ReadSeeker case.
	case *bytes.Reader:
		buf, err := ioutil.ReadAll(body)
		if err != nil {
			return nil, 0, err
		}
		bodyReader = func() (io.Reader, error) {
			return bytes.NewReader(buf), nil
		}
		contentLength = int64(len(buf))

	// Compat case
	case io.ReadSeeker:
		raw := body
		bodyReader = func() (io.Reader, error) {
			_, err := raw.Seek(0, 0)
			return ioutil.NopCloser(raw), err
		}
		if lr, ok := raw.(retryablehttp.LenReader); ok {
			contentLength = int64(lr.Len())
		}

	// Read all in so we can reset
	case io.Reader:
		buf, err := ioutil.ReadAll(body)
		if err != nil {
			return nil, 0, err
		}
		bodyReader = func() (io.Reader, error) {
			return bytes.NewReader(buf), nil
		}
		contentLength = int64(len(buf))

	// No body provided, nothing to do
	case nil:

	// Unrecognized type
	default:
		return nil, 0, fmt.Errorf("cannot handle type %T", rawBody)
	}
	return bodyReader, contentLength, nil
}

// Try to read the response body so we can reuse this connection.
func (c *RetryableHTTPClient) drainBody(body io.ReadCloser) {
	defer body.Close()
	_, err := io.Copy(ioutil.Discard, io.LimitReader(body, respReadLimit))
	if err != nil {
		if c.logger() != nil {
			switch v := c.logger().(type) {
			case retryablehttp.LeveledLogger:
				v.Error("error reading response body", "error", err)
			case retryablehttp.Logger:
				v.Printf("[ERR] error reading response body: %v", err)
			}
		}
	}
}

func (c *RetryableHTTPClient) logger() interface{} {
	c.loggerInit.Do(func() {
		if c.Logger == nil {
			return
		}

		switch c.Logger.(type) {
		case retryablehttp.Logger, retryablehttp.LeveledLogger:
			// ok
		default:
			// This should happen in dev when they are setting Logger and work on code, not in prod.
			panic(fmt.Sprintf("invalid logger type passed, must be Logger or LeveledLogger, was %T", c.Logger))
		}
	})

	return c.Logger
}

// APIClient provides functions for making API requests
type APIClient struct {
	client        RetryableHTTPClient
	configuration *Configuration
}

var (
	redirectsErrorRe = regexp.MustCompile(`stopped after \d+ redirects\z`)

	schemeErrorRe = regexp.MustCompile(`unsupported protocol scheme`)

	invalidHeaderErrorRe = regexp.MustCompile(`invalid header`)

	notTrustedErrorRe = regexp.MustCompile(`certificate is not trusted`)
)

// NewAPIClient creates a new API client
func NewAPIClient(c *Configuration) APIClient {
	timeout, err := time.ParseDuration("32s")
	if err != nil {
		panic(err)
	}

	client := NewClient()
	client.Logger = nil
	client.HTTPClient.Timeout = timeout

	return APIClient{
		client:        *client,
		configuration: c,
	}
}

// NewClient creates a new Client with default settings.
func NewClient() *RetryableHTTPClient {
	return &RetryableHTTPClient{
		Client: &retryablehttp.Client{
			HTTPClient:   cleanhttp.DefaultPooledClient(),
			Logger:       defaultLogger,
			RetryWaitMin: defaultRetryWaitMin,
			RetryWaitMax: defaultRetryWaitMax,
			RetryMax:     defaultRetryMax,
			CheckRetry:   DefaultRetryPolicy,
			Backoff:      DefaultBackoff,
		},
	}
}

// DefaultBackoff provides a default callback for Client.Backoff which
// will perform exponential backoff based on the attempt number and limited
// by the provided minimum and maximum durations.
//
// It also tries to parse Retry-After response header when a http.StatusTooManyRequests
// (HTTP Code 429) is found in the resp parameter. Hence it will return the number of
// seconds the server states it may be ready to process more requests from this client.
func DefaultBackoff(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
	if resp != nil {
		if resp.StatusCode == http.StatusTooManyRequests ||
			resp.StatusCode == http.StatusServiceUnavailable {
			if s, ok := resp.Header["Retry-After"]; ok {
				if sleep, err := strconv.ParseInt(s[0], 10, 64); err == nil {
					return time.Second * time.Duration(sleep)
				}
			}
		}
	}

	mult := math.Pow(2, float64(attemptNum)) * float64(min)
	sleep := time.Duration(mult)
	if float64(sleep) != mult || sleep > max {
		sleep = max
	}
	return sleep
}

// SelectHeaderContentType selects the header content type
func (c *APIClient) SelectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// SelectHeaderAccept selects the header accept
func (c *APIClient) SelectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}
	if contains(accepts, "application/json") {
		return "application/json"
	}
	return strings.Join(accepts, ",")
}

func contains(source []string, containvalue string) bool {
	for _, a := range source {
		if strings.ToLower(a) == strings.ToLower(containvalue) {
			return true
		}
	}
	return false
}

// Function to get PathValue for a given PathName
func getPathValue(pc *ProxyConfiguration, key string) (string, bool) {
	for _, param := range pc.PathParams {
		if param.PathName == key {
			return param.PathValue, true
		}
	}
	return "", false
}

func getConfUrl(c *Configuration, path string, extendedPath string) *url.URL {
	var uri *url.URL
	gateWayConfiguration := c.GateWayConfiguration

	pathValue := ""
	if len(gateWayConfiguration.PathParams) > 0 {
		for _, param := range gateWayConfiguration.PathParams {
			if param.PathName == path {
				tempValue := param.PathValue

				if !strings.HasPrefix(tempValue, "/") {
					pathValue = fmt.Sprintf("/%v", tempValue)
				} else {
					pathValue = fmt.Sprintf("%v", tempValue)
				}
				break
			}
		}
	}
	if gateWayConfiguration.Port == "80" || gateWayConfiguration.Port == "443" {
		uri, _ = url.Parse(
			fmt.Sprintf(
				"%s://%s%s%s",
				gateWayConfiguration.Protocol,
				gateWayConfiguration.Host,
				pathValue,
				extendedPath,
			),
		)
	} else {
		uri, _ = url.Parse(fmt.Sprintf("%s://%s:%s/%s%s", gateWayConfiguration.Protocol, gateWayConfiguration.Host, gateWayConfiguration.Port, pathValue, extendedPath))
	}

	return uri
}

// CallAPI invokes an API endpoint
func (c *APIClient) CallAPI(path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams map[string]string,
	formParams url.Values,
	fileName string,
	fileBytes []byte,
	pathName string) (*APIResponse, error) {
	var u *url.URL

	urlString := path + "?"
	if queryParams != nil {
		for k, v := range queryParams {
			if v != "" {
				urlString += fmt.Sprintf(
					"%s=%s&",
					url.QueryEscape(strings.TrimSpace(k)),
					url.QueryEscape(strings.TrimSpace(v)),
				)
			}
		}
	}
	urlString = urlString[:len(urlString)-1] // Remove the trailing '&'

	if c.configuration.GateWayConfiguration != nil &&
		c.configuration.GateWayConfiguration.Host != "" {
		if index := strings.Index(urlString, "/api/v2"); index != -1 {
			urlString = urlString[index:] // Get substring from /api onward
		}
		if pathName == "login" {
			u = getConfUrl(c.configuration, pathName, "/oauth/token")
		} else {
			u = getConfUrl(c.configuration, "api", urlString)
		}
	} else {
		u, _ = url.Parse(urlString)
	}

	request := Request{
		Request: &http.Request{
			URL:    u,
			Close:  true,
			Method: strings.ToUpper(method),
			Header: make(map[string][]string),
		},
	}

	// Set default headers
	if c.configuration.DefaultHeader != nil {
		for k, v := range c.configuration.DefaultHeader {
			request.Header.Set(k, v)
		}
	}

	// Set form data
	if len(formParams) > 0 {
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		request.SetBody(ioutil.NopCloser(strings.NewReader(formParams.Encode())))
	}

	// Set post body
	if postBody != nil {
		request.Header.Set("Content-Type", "application/json")
		// Check if postBody is of type *string
		if body, ok := postBody.(*string); ok {
			j := []byte(*body)
			request.SetBody(ioutil.NopCloser(bytes.NewReader(j)))
		} else {
			j, _ := json.Marshal(postBody)
			request.SetBody(ioutil.NopCloser(bytes.NewReader(j)))
		}
	}

	// Set provided headers
	if len(headerParams) > 0 {
		for k, v := range headerParams {
			request.Header.Set(k, v)
		}
	}

	if c.configuration.RetryConfiguration == nil {
		c.client.RetryMax = 0
		c.client.RetryWaitMax = 0
	} else {
		c.client.RetryWaitMax = c.configuration.RetryConfiguration.RetryWaitMax
		c.client.RetryWaitMin = c.configuration.RetryConfiguration.RetryWaitMin
		c.client.RetryMax = c.configuration.RetryConfiguration.RetryMax
		if c.configuration.RetryConfiguration.RequestLogHook != nil {
			c.client.RequestLogHook = func(_ retryablehttp.Logger, req *http.Request, retryNumber int) {
				c.configuration.RetryConfiguration.RequestLogHook(req, retryNumber)
			}
		}
		if c.configuration.RetryConfiguration.ResponseLogHook != nil {
			c.client.ResponseLogHook = func(_ retryablehttp.Logger, res *http.Response) {
				c.configuration.RetryConfiguration.ResponseLogHook(res)
			}
		}
	}

	c.client.CheckRetry = DefaultRetryPolicy

	if c.configuration.ProxyConfiguration != nil {
		var proxyUrl *url.URL
		pathValue, exists := getPathValue(c.configuration.ProxyConfiguration, pathName)

		if c.configuration.ProxyConfiguration.Auth != nil &&
			c.configuration.ProxyConfiguration.Auth.UserName != "" &&
			c.configuration.ProxyConfiguration.Auth.Password != "" {
			proxyUrl = &url.URL{
				Scheme: c.configuration.ProxyConfiguration.Protocol,
				User: url.UserPassword(c.configuration.ProxyConfiguration.Auth.UserName,
					c.configuration.ProxyConfiguration.Auth.Password),
				Host: c.configuration.ProxyConfiguration.Host + ":" + c.configuration.ProxyConfiguration.Port,
			}
		} else {

			urlString := c.configuration.ProxyConfiguration.Protocol + "://" +
				c.configuration.ProxyConfiguration.Host + ":" +
				c.configuration.ProxyConfiguration.Port
			proxyUrl, _ = url.Parse(urlString)
		}

		if exists {
			proxyUrl.Path = pathValue
		}

		tr := &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		}

		c.client.HTTPClient.Transport = tr
	}

	requestBody, _ := request.BodyBytes()

	// Execute request
	res, err := c.client.Do(&request)
	if err != nil {
		return nil, err
	}

	// Read body
	body, _ := ioutil.ReadAll(res.Body)
	c.configuration.LoggingConfiguration.trace(
		method,
		urlString,
		requestBody,
		res.StatusCode,
		request.Header,
		res.Header,
	)
	c.configuration.LoggingConfiguration.debug(
		method,
		urlString,
		requestBody,
		res.StatusCode,
		request.Header,
	)

	if res.StatusCode == http.StatusUnauthorized && c.configuration.ShouldRefreshAccessToken &&
		c.configuration.RefreshToken != "" {
		err := c.handleExpiredAccessToken()
		if err != nil {
			return nil, err
		}
		if headerParams != nil {
			headerParams["Authorization"] = "Bearer " + c.configuration.AccessToken
		}

		return c.CallAPI(
			path,
			method,
			postBody,
			headerParams,
			queryParams,
			formParams,
			fileName,
			fileBytes,
			pathName,
		)
	}

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		c.configuration.LoggingConfiguration.error(
			method,
			urlString,
			requestBody,
			body,
			res.StatusCode,
			request.Header,
			res.Header,
		)
	}

	return NewAPIResponse(res, body)
}

func DefaultRetryPolicy(ctx context.Context, resp *http.Response, err error) (bool, error) {
	// do not retry on context.Canceled or context.DeadlineExceeded
	if ctx.Err() != nil {
		return false, ctx.Err()
	}

	// don't propagate other errors
	shouldRetry, _ := newRetryPolicy(resp, err)
	return shouldRetry, nil
}

func newRetryPolicy(resp *http.Response, err error) (bool, error) {
	if err != nil {
		if v, ok := err.(*url.Error); ok {
			// Don't retry if the error was due to too many redirects.
			if redirectsErrorRe.MatchString(v.Error()) {
				return false, v
			}

			// Don't retry if the error was due to an invalid protocol scheme.
			if schemeErrorRe.MatchString(v.Error()) {
				return false, v
			}

			// Don't retry if the error was due to an invalid header.
			if invalidHeaderErrorRe.MatchString(v.Error()) {
				return false, v
			}

			// Don't retry if the error was due to TLS cert verification failure.
			if notTrustedErrorRe.MatchString(v.Error()) {
				return false, v
			}
			if isCertError(v.Err) {
				return false, v
			}
		}

		// The error is likely recoverable so retry.
		return true, nil
	}

	// 429 Too Many Requests is recoverable. Sometimes the server puts
	// a Retry-After response header to indicate when the server is
	// available to start processing request from client.
	if resp.StatusCode == http.StatusTooManyRequests {
		return verifyRetryAfterHeader(resp.Header["Retry-After"], 180), nil
	}

	// Check the response code. We retry on 500-range responses to allow
	// the server time to recover, as 500's are typically not permanent
	// errors and may relate to outages on the server side. This will catch
	// invalid response codes as well, like 0 and 999.
	if resp.StatusCode == 0 ||
		(resp.StatusCode >= 500 && resp.StatusCode != http.StatusNotImplemented) {
		return true, fmt.Errorf("unexpected HTTP status %s", resp.Status)
	}

	return false, nil
}

func isCertError(err error) bool {
	_, ok := err.(*tls.CertificateVerificationError)
	return ok
}

func verifyRetryAfterHeader(headers []string, defaultMaxRetry int64) bool {
	if len(headers) == 0 || headers[0] == "" {
		return true
	}
	header := headers[0]

	if sleep, err := strconv.ParseInt(header, 10, 64); err == nil {
		if sleep > defaultMaxRetry {
			return false
		}
		return true
	}
	return true
}

// ParameterToString joins a parameter in the desired format
func (c *APIClient) ParameterToString(obj interface{}, collectionFormat string) string {
	sep := ","
	switch collectionFormat {
	case "pipes":
		sep = "|"
	case "ssv":
		sep = " "
	case "tsv":
		sep = "\t"
	}

	switch t := reflect.TypeOf(obj).String(); t {
	case "[]string":
		return strings.Join(obj.([]string), sep)
	default:
		return fmt.Sprintf("%v", obj)
	}
}

func (c *APIClient) handleExpiredAccessToken() error {
	if atomic.CompareAndSwapInt64(&c.configuration.RefreshInProgress, 0, 1) {
		defer atomic.StoreInt64(&c.configuration.RefreshInProgress, 0)
		_, err := c.configuration.RefreshAuthorizationCodeGrant(
			c.configuration.ClientID,
			c.configuration.ClientSecret,
			c.configuration.RefreshToken,
		)
		return err
	} else {
		// Wait maximum of RefreshTokenWaitTime seconds for other thread to complete refresh
		startTime := time.Now().Unix()
		sleepDuration := time.Millisecond * 200
		// Check if we've gone over the wait threshold
		for time.Now().Unix()-startTime < int64(c.configuration.RefreshTokenWaitTime) {
			time.Sleep(sleepDuration) // Sleep for 200ms on every iteration
			if atomic.LoadInt64(&c.configuration.RefreshInProgress) == 0 {
				return nil
			}
		}
		return fmt.Errorf("token refresh took longer than %d seconds", c.configuration.RefreshTokenWaitTime)
	}
}

// Int32 is an easy way to get an int32 pointer
func Int32(v int) *int32 {
	p := int32(v)
	return &p
}

// Int64 is an easy way to get an int64 pointer
func Int64(v int) *int64 {
	p := int64(v)
	return &p
}

// Int is an easy way to get an int pointer
func Int(v int) *int {
	return &v
}

// Float32 is an easy way to get a float32 pointer
func Float32(v float32) *float32 {
	return &v
}

// Float64 is an easy way to get a float64 pointer
func Float64(v float64) *float64 {
	return &v
}

// String is an easy way to get a string pointer
func String(v string) *string {
	return &v
}

// Bool is an easy way to get a pointer
func Bool(v bool) *bool {
	return &v
}

func copy(data []byte, v interface{}) {
	dataS := string(data)

	reflect.ValueOf(v).Elem().Set(reflect.ValueOf(&dataS))
}

func getFieldName(t reflect.Type, field string) string {
	// Find JSON prop name
	structField, ok := t.Elem().FieldByName(field)
	if ok {
		tag := structField.Tag.Get("json")
		if tag != "" {
			tagParts := strings.Split(tag, ",")
			if len(tagParts) > 0 {
				return tagParts[0]
			}
		}
	}

	return field
}

// toTime is a helper function for models to prevent a static import of the "time" package duplicating dynamic imports determined by the codegen process
func toTime(o interface{}) *time.Time {
	return o.(*time.Time)
}
