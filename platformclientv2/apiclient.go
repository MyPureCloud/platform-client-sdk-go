package platformclientv2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
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

// CallAPI invokes an API endpoint
func (c *APIClient) CallAPI(path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams map[string]string,
	formParams url.Values,
	fileName string,
	fileBytes []byte) (*APIResponse, error) {

	// Build request URL w/query params
	urlString := path + "?"
	if queryParams != nil {
		for k, v := range queryParams {
			if v != "" {
				urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
		}
	}
	urlString = urlString[:len(urlString)-1]
	u, err := url.Parse(urlString)
	if err != nil {
		return nil, err
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
		j, _ := json.Marshal(postBody)
		request.SetBody(ioutil.NopCloser(bytes.NewReader(j)))
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

	requestBody, _ := request.BodyBytes()

	// Execute request
	res, err := c.client.Do(&request)
	if err != nil {
		return nil, err
	}

	// Read body
	body, _ := ioutil.ReadAll(res.Body)
	c.configuration.LoggingConfiguration.trace(method, urlString, requestBody, res.StatusCode, request.Header, res.Header)
	c.configuration.LoggingConfiguration.debug(method, urlString, requestBody, res.StatusCode, request.Header)

	if res.StatusCode == http.StatusUnauthorized && c.configuration.ShouldRefreshAccessToken && c.configuration.RefreshToken != "" {
		err := c.handleExpiredAccessToken()
		if err != nil {
			return nil, err
		}
		if headerParams != nil {
			headerParams["Authorization"] = "Bearer " + c.configuration.AccessToken
		}

		return c.CallAPI(path, method, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	}

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		c.configuration.LoggingConfiguration.error(method, urlString, requestBody, body, res.StatusCode, request.Header, res.Header)
	}

	return NewAPIResponse(res, body)
}

func buildURL(basePath string, path string, queryParams map[string]string) (*url.URL, error) {
	urlString := basePath + path
	if len(queryParams) > 0 {
		urlString += "?"
		for k, v := range queryParams {
			urlString += fmt.Sprintf("%v=%v&", strings.TrimSpace(k), strings.TrimSpace(v))
		}
		urlString = urlString[:len(urlString)-1]
	}

	u, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}
	return u, nil
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
		_, err := c.configuration.RefreshAuthorizationCodeGrant(c.configuration.ClientID, c.configuration.ClientSecret, c.configuration.RefreshToken)
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
