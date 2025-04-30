package platformclientv2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/hashicorp/go-retryablehttp"
)

// HTTPRequestOptions contains the configuration for making HTTP requests
type HTTPRequestOptions struct {
	url         *url.URL            // The request URL
	method      string              // HTTP method (GET, POST, etc)
	headers     map[string][]string // HTTP headers
	queryParams map[string]string   // URL query parameters
	body        interface{}         // Request body
	formParams  url.Values          // Form parameters
}

// GetUrl returns the request URL
func (o *HTTPRequestOptions) GetUrl() *url.URL {
	return o.url
}

// GetMethod returns the HTTP method
func (o *HTTPRequestOptions) GetMethod() string {
	return o.method
}

// GetHeaders returns the HTTP headers
func (o *HTTPRequestOptions) GetHeaders() map[string][]string {
	return o.headers
}

// GetQueryParams returns the URL query parameters
func (o *HTTPRequestOptions) GetQueryParams() map[string]string {
	return o.queryParams
}

// GetBody returns the request body
func (o *HTTPRequestOptions) GetBody() interface{} {
	return o.body
}

// GetFormParams returns the form parameters
func (o *HTTPRequestOptions) GetFormParams() url.Values {
	return o.formParams
}

// SetUrl sets the request URL, returns error if URL is nil
func (o *HTTPRequestOptions) SetUrl(url *url.URL) error {
	if url == nil {
		return fmt.Errorf("Invalid URL")
	}
	o.url = url
	return nil
}

// SetMethod sets the HTTP method, validates against allowed methods
func (o *HTTPRequestOptions) SetMethod(method string) error {
	if method == "" {
		return fmt.Errorf("Method cannot be empty")
	}

	method = strings.ToUpper(method)

	// Define valid HTTP methods
	validMethods := map[string]bool{
		"GET":     true,
		"HEAD":    true,
		"POST":    true,
		"PUT":     true,
		"PATCH":   true,
		"DELETE":  true,
		"OPTIONS": true,
	}

	if validMethods[method] {
		o.method = method
		return nil
	}
	return fmt.Errorf("Invalid method: %s", method)
}

// SetHeaders sets a single header key-value pair
func (o *HTTPRequestOptions) SetHeaders(key string, value string) error {
	if key == "" {
		return fmt.Errorf("header key cannot be empty")
	}
	if o.headers == nil {
		o.headers = make(map[string][]string)
	}
	o.headers[key] = []string{value}
	return nil
}

// SetQueryParams sets the URL query parameters
func (o *HTTPRequestOptions) SetQueryParams(queryParams map[string]string) error {
	if queryParams == nil {
		return fmt.Errorf("Query Params cannot be empty")
	}
	o.queryParams = queryParams
	return nil
}

// SetBody sets the request body
func (o *HTTPRequestOptions) SetBody(body interface{}) error {
	if body == nil {
		return fmt.Errorf("Body cannot be empty")
	}
	o.body = body
	return nil
}

// SetFormParams sets the form parameters
func (o *HTTPRequestOptions) SetFormParams(formParams url.Values) error {
	if formParams == nil {
		return fmt.Errorf("Form Params cannot be empty")
	}
	o.formParams = formParams
	return nil
}

// ToRetryableRequest converts HTTPRequestOptions to a retryablehttp.Request
func (o *HTTPRequestOptions) ToRetryableRequest() (*retryablehttp.Request, error) {
	// Validate required fields
	if o.GetUrl() == nil {
		return nil, fmt.Errorf("URL is required")
	}
	if o.GetMethod() == "" {
		return nil, fmt.Errorf("Method is required")
	}

	// Create base request
	request := retryablehttp.Request{
		Request: &http.Request{
			URL:    o.GetUrl(),
			Close:  true,
			Method: strings.ToUpper(o.GetMethod()),
			Header: o.GetHeaders(),
		},
	}

	// Set form data if present
	if len(o.GetFormParams()) > 0 {
		request.SetBody(ioutil.NopCloser(strings.NewReader(o.GetFormParams().Encode())))
	}

	// Set request body if present
	if o.GetBody() != nil {
		// Handle string body type differently
		if body, ok := o.GetBody().(*string); ok {
			j := []byte(*body)
			request.SetBody(ioutil.NopCloser(bytes.NewReader(j)))
		} else {
			// Marshal non-string body to JSON
			j, _ := json.Marshal(o.GetBody())
			request.SetBody(ioutil.NopCloser(bytes.NewReader(j)))
		}
	}

	return &request, nil
}
