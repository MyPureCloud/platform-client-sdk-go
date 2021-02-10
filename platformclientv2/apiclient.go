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
)

// APIClient provides functions for making API requests
type APIClient struct {
	client        http.Client
	configuration *Configuration
}

// NewAPIClient creates a new API client
func NewAPIClient(c *Configuration) APIClient {
	timeout, err := time.ParseDuration("16s")
	if err != nil {
		panic(err)
	}
	return APIClient{
		client:        http.Client{Timeout: timeout},
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
			urlString += fmt.Sprintf("%v=%v&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
		}
	}
	urlString = urlString[:len(urlString)-1]
	u, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}
	request := http.Request{
		URL:    u,
		Close:  true,
		Method: strings.ToUpper(method),
		Header: make(map[string][]string),
	}

	// Set default headers
	if c.configuration.DefaultHeader != nil {
		for k, v := range c.configuration.DefaultHeader {
			fmt.Printf("  %v=%v", k, v)
			request.Header.Set(k, v)
		}
	}

	// Set form data
	if formParams != nil {
		// request.Form = make(map[string][]string)
		// request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		// for k, v := range formParams {
		// 	request.Form.Set(k, v)
		// }

		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		request.Body = ioutil.NopCloser(strings.NewReader(formParams.Encode()))
	}

	// Set post body
	if postBody != nil {
		request.Header.Set("Content-Type", "application/json")
		j, _ := json.Marshal(postBody)
		request.Body = ioutil.NopCloser(bytes.NewReader(j))
	}

	// Set provided headers
	if headerParams != nil {
		for k, v := range headerParams {
			request.Header.Set(k, v)
		}
	}

	// Debug request
	if c.configuration.GetDebug() {
		c.configuration.Debugf("======== %v\n", time.Now())
		c.configuration.Debugf("%v %v\n", request.Method, u.Path)
		c.configuration.Debugf("HOST: %v:%v\n", u.Host, u.Port())
		c.configuration.Debug("HEADERS:")
		for k, v := range request.Header {
			c.configuration.Debugf("  %v=%v\n", k, v)
		}
		c.configuration.Debug(request)
	}

	// Execute request
	reqStart := time.Now()
	res, err := c.client.Do(&request)
	reqEnd := time.Now()
	duration := reqEnd.Sub(reqStart)

	// Read body
	body, _ := ioutil.ReadAll(res.Body)

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

	// Debug response
	if c.configuration.GetDebug() {
		c.configuration.Debugf("==== RESPONSE %v (%v) ====\n", reqEnd, duration)
		c.configuration.Debugf("STATUS: %v\n", res.Status)
		c.configuration.Debug("HEADERS:")
		for k, v := range res.Header {
			c.configuration.Debugf("  %v=%v\n", k, v)
		}
		if body != nil {
			c.configuration.Debugf("BODY:\n%v\n", string(body))
		}
		c.configuration.Debug("========")
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
		// Check if we've gone over the wait threshhold
		for time.Now().Unix() - startTime < int64(c.configuration.RefreshTokenWaitTime) {
			time.Sleep(sleepDuration) // Sleep for 200ms on every iteration
			if atomic.LoadInt64(&c.configuration.RefreshInProgress) == 0 {
				return nil
			}
		}
		return fmt.Errorf("Token refresh took longer than %d seconds", c.configuration.RefreshTokenWaitTime)
	}

	return nil
}

// Int32 is an easy way to get a pointer
func Int32(v int) *int32 {
	p := int32(v)
	return &p
}

// String is an easy way to get a pointer
func String(v string) *string {
	return &v
}

// Bool is an easy way to get a pointer
func Bool(v bool) *bool {
	return &v
}
