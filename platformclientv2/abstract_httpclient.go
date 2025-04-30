package platformclientv2

import (
	"context"
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

// AbstractHttpClient defines the interface for an HTTP client with retry capabilities
type AbstractHttpClient interface {
	// Do executes an HTTP request with the given options
	Do(options *HTTPRequestOptions) (*http.Response, error)

	// SetRetryMax sets the maximum number of retries for failed requests
	SetRetryMax(max int)

	// SetRetryWaitMax sets the maximum wait time between retries
	SetRetryWaitMax(duration time.Duration)

	// SetRetryWaitMin sets the minimum wait time between retries
	SetRetryWaitMin(duration time.Duration)

	// SetRequestLogHook sets a logging hook that is called before each retry attempt
	SetRequestLogHook(hook func(retryablehttp.Logger, *http.Request, int))

	// SetResponseLogHook sets a logging hook that is called after each response
	SetResponseLogHook(hook func(retryablehttp.Logger, *http.Response))

	// SetCheckRetry sets the retry policy function to determine if a request should be retried
	SetCheckRetry(checkRetry func(ctx context.Context, resp *http.Response, err error) (bool, error))

	// SetTransport sets the underlying HTTP transport for the client
	SetTransport(transport *http.Transport)

	// SetHttpsAgent
	SetHttpsAgent(proxy *ProxyAgent)
}

// ProxyAgent holds proxy configuration
type ProxyAgent struct {
	ProxyURL string
}
