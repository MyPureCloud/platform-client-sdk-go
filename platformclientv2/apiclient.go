package platformclientv2

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

// APIClient provides functions for making API requests
type APIClient struct {
	client        retryablehttp.Client
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
	timeout, err := time.ParseDuration("16s")
	if err != nil {
		panic(err)
	}

	client := retryablehttp.NewClient()
	client.Logger = nil
	client.HTTPClient.Timeout = timeout

	return APIClient{
		client:        *client,
		configuration: c,
	}
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

	request := retryablehttp.Request{
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
		return false, nil
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
