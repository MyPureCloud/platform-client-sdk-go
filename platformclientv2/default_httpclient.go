package platformclientv2

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

type retryableLogger struct {
}

func (rl *retryableLogger) Printf(format string, v ...interface{}) {
}

// DefaultHttpClient wraps retryablehttp.Client to provide HTTP functionality with automatic retries
type DefaultHttpClient struct {
	client     retryablehttp.Client
	proxyAgent *ProxyAgent
}

// SetRetryMax sets the maximum number of retries for failed requests
func (c *DefaultHttpClient) SetRetryMax(max int) {
	c.client.RetryMax = max
}

// SetRetryWaitMax sets the maximum time to wait between retries
func (c *DefaultHttpClient) SetRetryWaitMax(duration time.Duration) {
	c.client.RetryWaitMax = duration
}

// SetRetryWaitMin sets the minimum time to wait between retries
func (c *DefaultHttpClient) SetRetryWaitMin(duration time.Duration) {
	c.client.RetryWaitMin = duration
}

// SetPreHook sets a logging hook that runs before each retry
func (c *DefaultHttpClient) SetPreHook(hook func(retryablehttp.Logger, *http.Request, int)) {
	c.client.RequestLogHook = hook
}

// SetPostHook sets a logging hook that runs after each retry
func (c *DefaultHttpClient) SetPostHook(hook func(retryablehttp.Logger, *http.Response)) {
	c.client.ResponseLogHook = hook
}

// SetCheckRetry sets the retry policy function to determine if a request should be retried
func (c *DefaultHttpClient) SetCheckRetry(checkRetry func(ctx context.Context, resp *http.Response, err error) (bool, error)) {
	c.client.CheckRetry = checkRetry
}

// SetTransport sets the underlying HTTP transport for the client
func (c *DefaultHttpClient) SetTransport(transport *http.Transport) {
	c.client.HTTPClient.Transport = transport
}

func (c *DefaultHttpClient) SetHttpsAgent(agent *ProxyAgent) {
	c.proxyAgent = agent
}

// NewDefaultHttpClient creates and returns a new DefaultHttpClient instance with default settings
func NewDefaultHttpClient(proxyAgent *ProxyAgent) *DefaultHttpClient {
	timeout, err := time.ParseDuration("16s")
	if err != nil {
		panic(err)
	}

	client := retryablehttp.NewClient()
	client.Logger = &retryableLogger{}
	client.HTTPClient.Timeout = timeout
	if proxyAgent != nil && proxyAgent.ProxyURL != "" {
		proxyURL, err := url.Parse(proxyAgent.ProxyURL)
		if err == nil {
			transport := &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			}
			client.HTTPClient.Transport = transport
		}
	}

	return &DefaultHttpClient{
		client:     *client,
		proxyAgent: proxyAgent,
	}
}

// Do executes an HTTP request with the configured retry settings
func (c *DefaultHttpClient) Do(options *HTTPRequestOptions) (*http.Response, error) {
	request, err := options.ToRetryableRequest()
	if err != nil {
		return nil, err
	}
	return c.client.Do(request)
}
