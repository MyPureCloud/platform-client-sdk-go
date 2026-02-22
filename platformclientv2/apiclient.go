package platformclientv2

import (
	"context"
    "crypto/tls"
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
	"crypto/x509"
	"os"

	"github.com/hashicorp/go-retryablehttp"
)

// APIClient provides functions for making API requests
type APIClient struct {
	client        AbstractHttpClient
	configuration *Configuration
	proxyAgent    *ProxyAgent
}

var (

redirectsErrorRe = regexp.MustCompile(`stopped after \d+ redirects\z`)


schemeErrorRe = regexp.MustCompile(`unsupported protocol scheme`)


invalidHeaderErrorRe = regexp.MustCompile(`invalid header`)


notTrustedErrorRe = regexp.MustCompile(`certificate is not trusted`)
)


// SetHttpClient sets the HTTP client
func (c *APIClient) SetHttpClient(client interface{}) error {
    // check if client implements AbstractHttpClient
    httpClient, ok := client.(AbstractHttpClient)
    if !ok {
        return fmt.Errorf("httpClient must implement AbstractHttpClient interface. See DefaultHttpClient for an example")
    }
    c.client = httpClient
    return nil
}

// GetHttpClient gets or creates the HTTP client
func (c *APIClient) GetHttpClient() AbstractHttpClient {
    if c.client != nil {
        return c.client
    }
    
    // Create new default client if none exists
    c.client = NewDefaultHttpClient(c.proxyAgent)
    return c.client
}

func (c *APIClient) SetProxyAgent(proxyAgent *ProxyAgent) {
    c.proxyAgent = proxyAgent
	client := c.GetHttpClient()
	client.SetHttpsAgent(proxyAgent)
}

// NewAPIClient creates a new API client
func NewAPIClient(c *Configuration) APIClient {
	defaultClient := c.APIClient.GetHttpClient()

	return APIClient{
		client:        defaultClient,
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
				if tempValue != "" {
					tempValue = strings.TrimSpace(tempValue)
				}
				if tempValue != "" {
					tempValue = strings.TrimRight(tempValue, "/")
				}
				if tempValue != "" {
					tempValue = strings.TrimLeft(tempValue, "/")
				}

				if tempValue != "" {
					if !strings.HasPrefix(tempValue, "/") {
						pathValue = fmt.Sprintf("/%v", tempValue)
					} else {
						pathValue = fmt.Sprintf("%v", tempValue)
					}
				}
				break
			}
		}
	}
	if gateWayConfiguration.Port == "80" || gateWayConfiguration.Port == "443" {
		uri, _ = url.Parse(fmt.Sprintf("%s://%s%s%s", gateWayConfiguration.Protocol, gateWayConfiguration.Host, pathValue, extendedPath))
	} else {
		uri, _ = url.Parse(fmt.Sprintf("%s://%s:%s%s%s", gateWayConfiguration.Protocol, gateWayConfiguration.Host, gateWayConfiguration.Port, pathValue, extendedPath))
	}

	return uri
}

// Handles the Proxy Configuration
func (c *APIClient) configureProxy(pathName string) (*http.Transport, error) {
	if c.configuration.ProxyConfiguration == nil {
		return nil, nil
	}

	var proxyUrl *url.URL
    pathValue, exists :=  getPathValue(c.configuration.ProxyConfiguration, pathName)

    if c.configuration.ProxyConfiguration.Auth != nil && c.configuration.ProxyConfiguration.Auth.UserName != "" && c.configuration.ProxyConfiguration.Auth.Password != "" {
        proxyUrl = &url.URL{
            Scheme: c.configuration.ProxyConfiguration.Protocol,
            User: url.UserPassword(
				c.configuration.ProxyConfiguration.Auth.UserName,
            	c.configuration.ProxyConfiguration.Auth.Password,
			),
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

	return &http.Transport{
        Proxy: http.ProxyURL(proxyUrl),
    }, nil
}

// SetMTLSContents sets the MTLS certificate contents directly rather than loading from files.
func (c *APIClient) SetMTLSContents(certContent, keyContent, caContent string) error {
    if certContent == "" {
        return fmt.Errorf("certificate content cannot be empty")
    }

    if keyContent == "" {
        return fmt.Errorf("private key content cannot be empty") 
    }

    if caContent == "" {
        return fmt.Errorf("CA certificate content cannot be empty")
    }

    c.configuration.MTLSConfiguration = &MTLSConfiguration {
		CertFile: []byte(certContent),
		KeyFile: []byte(keyContent),
		CAFile: []byte(caContent),
	}
	return nil
}

// Handles the MTLS Gateway Configuration
func (c *APIClient) configureMTLS() (*tls.Config, error) {
	if c.configuration.MTLSConfiguration == nil {
		return nil, nil
	}

	if c.configuration.MTLSConfiguration.CertFile == nil {
		return nil, fmt.Errorf("failed to load certificate content")
	}

	if c.configuration.MTLSConfiguration.KeyFile == nil {
		return nil, fmt.Errorf("failed to load Key file content")
	}

	if c.configuration.MTLSConfiguration.CAFile == nil {
        return nil, fmt.Errorf("Failed to load CA certificate content")
    }

	cert, err := tls.X509KeyPair(
		[]byte(c.configuration.MTLSConfiguration.CertFile), 
		[]byte(c.configuration.MTLSConfiguration.KeyFile),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to load client certificate: %v", err)
	}

	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM([]byte(c.configuration.MTLSConfiguration.CAFile)) {
		return nil, fmt.Errorf("failed to append CA certificate")
	}

	// Create TLS config with custom verification
    tlsConfig := &tls.Config{
        Certificates: []tls.Certificate{cert},
        RootCAs:     caCertPool,
        MinVersion:  tls.VersionTLS12,
		InsecureSkipVerify: true,  // Required for custom verification
    }

    return tlsConfig, nil
}

// Handles the transport configurations (Proxy, Gateway)
func (c *APIClient) configureTransport(pathName string) (*http.Transport, error) {
	// Proxy Configuration
	proxyTr, err := c.configureProxy(pathName)
	if err != nil {
		return nil, err
	}

	// MTLS Configuration
	mtlsConfig, err := c.configureMTLS()
	if err != nil {
		return nil, err
	}

	transport := &http.Transport{}

	if proxyTr != nil {
		transport.Proxy = proxyTr.Proxy
	}

	if mtlsConfig != nil {
		transport.TLSClientConfig = mtlsConfig
	}

	return transport, nil
}

// Sets the MTLS Certificates for the file paths
func (c *APIClient) SetMTLSCertificates(certFile, keyFile, caFile string) error {
	if certFile == "" {
		return fmt.Errorf("Failed to load the Certificate")
	}

	if keyFile == "" {
		return fmt.Errorf("Failed to load Key File")
	}

	if caFile == "" {
		return fmt.Errorf("Failed to load CA certificate")
	}
	
	certPEMBlock, err := os.ReadFile(certFile)
	if err != nil {
		return err
	}
	keyPEMBlock, err := os.ReadFile(keyFile)
	if err != nil {
		return err
	}

	caPEMBlock, err := os.ReadFile(caFile)
	if err != nil {
		return err
	}

	c.configuration.MTLSConfiguration = &MTLSConfiguration {
		CertFile: certPEMBlock,
		KeyFile: keyPEMBlock,
		CAFile: caPEMBlock,
	}

	return nil
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

	// Build initial URL with query parameters
	urlString := path + "?"
	if queryParams != nil {
		for k, v := range queryParams {
			if v != "" {
				urlString += fmt.Sprintf("%s=%s&", url.QueryEscape(strings.TrimSpace(k)), url.QueryEscape(strings.TrimSpace(v)))
			}
		}
	}
	urlString = urlString[:len(urlString)-1] // Remove the trailing '&'

	// Handle gateway configuration if present
	if c.configuration.GateWayConfiguration != nil && c.configuration.GateWayConfiguration.Host != "" {
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

	// Initialize HTTP request options
	httpRequestOptions := &HTTPRequestOptions{}
	httpRequestOptions.SetUrl(u)
	httpRequestOptions.SetMethod(method)

	// Set default headers from configuration
	if c.configuration.DefaultHeader != nil {
		for k, v := range c.configuration.DefaultHeader {
			httpRequestOptions.SetHeaders(k, v)
		}
	}

	// Set form parameters if present
	if len(formParams) > 0 {
		httpRequestOptions.SetHeaders("Content-Type", "application/x-www-form-urlencoded")
		httpRequestOptions.SetFormParams(formParams)
	}

	// Set JSON body if present
	if postBody != nil {
		httpRequestOptions.SetHeaders("Content-Type", "application/json")
	}

	// Set additional headers from parameters
	if len(headerParams) > 0 {
		for k, v := range headerParams {
			if v != "" {
				httpRequestOptions.SetHeaders(k, v)
			}
		}
	}

	httpRequestOptions.SetBody(postBody)

	// Configure retry behavior
	if c.configuration.RetryConfiguration == nil {
		c.client.SetRetryMax(0)
		c.client.SetRetryWaitMax(0)
	} else {
		c.client.SetRetryWaitMax(c.configuration.RetryConfiguration.RetryWaitMax)
		c.client.SetRetryWaitMin(c.configuration.RetryConfiguration.RetryWaitMin)
		c.client.SetRetryMax(c.configuration.RetryConfiguration.RetryMax)
		if c.configuration.RetryConfiguration.RequestLogHook != nil {
			c.client.SetPreHook(func(_ retryablehttp.Logger, req *http.Request, retryNumber int) {
				c.configuration.RetryConfiguration.RequestLogHook(req, retryNumber)
			})
		}
		if c.configuration.RetryConfiguration.ResponseLogHook != nil {
			c.client.SetPostHook(func(_ retryablehttp.Logger, res *http.Response) {
				c.configuration.RetryConfiguration.ResponseLogHook(res)
			})
		}
	}


	// Set retry policy
	c.client.SetCheckRetry(DefaultRetryPolicy)

	// Configure proxy if specified
	transport, err := c.configureTransport(pathName)
	if err != nil {
		return nil, fmt.Errorf("failed to configure transport: %v", err)
	}

	if transport.TLSClientConfig != nil || transport.Proxy != nil {
		c.client.SetTransport(transport)
	}

	// Build the request
	buildReq, err := httpRequestOptions.ToRetryableRequest()
	if err != nil {
		return nil, err
	}
	//For logging
	requestBody, _ := buildReq.BodyBytes()

	// Execute request
	res, err := c.client.Do(httpRequestOptions)
	if err != nil {
		return nil, err
	}

	// Read response body and log request/response details
	body, _ := ioutil.ReadAll(res.Body)
	c.configuration.LoggingConfiguration.trace(method, urlString, requestBody, res.StatusCode, httpRequestOptions.GetHeaders(), res.Header)
	c.configuration.LoggingConfiguration.debug(method, urlString, requestBody, res.StatusCode, httpRequestOptions.GetHeaders())

	// Handle unauthorized response by refreshing access token if configured
	if res.StatusCode == http.StatusUnauthorized && c.configuration.ShouldRefreshAccessToken && c.configuration.RefreshToken != "" {
		err := c.handleExpiredAccessToken()
		if err != nil {
			return nil, err
		}
		if headerParams != nil {
			headerParams["Authorization"] = "Bearer " + c.configuration.AccessToken
		}

		return c.CallAPI(path, method, postBody, headerParams, queryParams, formParams, fileName, fileBytes, pathName)
	}

	// Log errors for non-successful status codes
	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		c.configuration.LoggingConfiguration.error(method, urlString, requestBody, body, res.StatusCode, httpRequestOptions.GetHeaders(), res.Header)
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
	if resp.StatusCode == 0 || (resp.StatusCode >= 500 && resp.StatusCode != http.StatusNotImplemented) {
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
			return  false
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
