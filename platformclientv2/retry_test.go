package platformclientv2

import (
	"net"
	"net/http"
	"net/url"
	"sync/atomic"
	"testing"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

func TestClientRetry(t *testing.T) {
	testRetryErrorCode(t, http.StatusTooManyRequests)

	// retryablehttp retries all 500 codes except for 501, we'll test the first 10 for brevity
	for code := http.StatusInternalServerError; code <= http.StatusNotExtended; code++ {
		if code == http.StatusNotImplemented {
			continue
		}
		testRetryErrorCode(t, code)
	}
}

func TestClientDoNotRetry(t *testing.T) {
	testDoNotRetryErrorCode(t, http.StatusTooManyRequests)

	// retryablehttp retries all 500 codes except for 501, we'll test the first 10 for brevity
	for code := http.StatusInternalServerError; code <= http.StatusNotExtended; code++ {
		if code == http.StatusNotImplemented {
			continue
		}
		testDoNotRetryErrorCode(t, code)
	}
}

func testRetryErrorCode(t *testing.T, errorCode int) {
	// Test implementation adapted from https://github.com/hashicorp/go-retryablehttp/blob/master/client_test.go
	// Track the number of times the logging hook was called
	retryCount := -1
	hook := func(logger retryablehttp.Logger, req *http.Request, retryNumber int) {
		retryCount = retryNumber
	}
	responseHook := func(logger retryablehttp.Logger, res *http.Response) {}

	APIClient := NewAPIClient(&Configuration{})
	APIClient.client.SetRetryWaitMin(10 * time.Millisecond)
	APIClient.client.SetRetryWaitMax(50 * time.Millisecond)
	APIClient.client.SetRetryMax(50)
	APIClient.client.SetRequestLogHook(hook)
	APIClient.client.SetResponseLogHook(responseHook)

	// Create a request
	testBytes := []byte("hello")

	u, _ := url.Parse("http://127.0.0.1:28934/v1/foo")

	// Initialize HTTP request options
	httpRequestOptions := &HTTPRequestOptions{}
	httpRequestOptions.SetUrl(u)
	httpRequestOptions.SetMethod("GET")
	httpRequestOptions.SetBody(testBytes)

	// Send the request
	var resp *http.Response
	doneCh := make(chan struct{})
	go func() {
		defer close(doneCh)
		var err error
		resp, err = APIClient.client.Do(httpRequestOptions)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
	}()

	select {
	case <-doneCh:
		t.Fatalf("should retry on error")
	case <-time.After(200 * time.Millisecond):
		// Client should still be retrying due to connection failure.
	}

	// Create the mock handler. First we return error responses to ensure
	// that we power through and keep retrying in the face of recoverable
	// errors.
	code := int64(errorCode)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(int(atomic.LoadInt64(&code)))
	})

	// Create a test server
	list, err := net.Listen("tcp", ":28934")
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	defer list.Close()
	go http.Serve(list, handler)

	// Wait again
	select {
	case <-doneCh:
		t.Fatalf("should retry on error codes")
	case <-time.After(200 * time.Millisecond):
		// Client should still be retrying due to error codes.
	}

	// Start returning 200's
	atomic.StoreInt64(&code, http.StatusOK)

	// Wait again
	select {
	case <-doneCh:
	case <-time.After(time.Second):
		t.Fatalf("timed out")
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("exected %v, got: %d", http.StatusOK, resp.StatusCode)
	}

	if retryCount < 0 {
		t.Fatal("request log hook was not called")
	}
}

func testDoNotRetryErrorCode(t *testing.T, errorCode int) {
	// Test implementation adapted from https://github.com/hashicorp/go-retryablehttp/blob/master/client_test.go
	// Track the number of times the logging hook was called
	retryCount := -1
	hook := func(logger retryablehttp.Logger, req *http.Request, retryNumber int) {
		retryCount = retryNumber
	}
	responseHook := func(logger retryablehttp.Logger, res *http.Response) {}

	APIClient := NewAPIClient(&Configuration{})
	APIClient.client.SetRetryWaitMax(0)
	APIClient.client.SetRetryMax(0)
	APIClient.client.SetRequestLogHook(hook)
	APIClient.client.SetResponseLogHook(responseHook)

	// Create a request
	testBytes := []byte("hello")

	u, _ := url.Parse("http://127.0.0.1:28934/v1/foo")

	// Initialize HTTP request options
	httpRequestOptions := &HTTPRequestOptions{}
	httpRequestOptions.SetUrl(u)
	httpRequestOptions.SetMethod("GET")
	httpRequestOptions.SetBody(testBytes)

	// Send the request
	doneCh := make(chan struct{})
	go func() {
		defer close(doneCh)
		var err error
		_, err = APIClient.client.Do(httpRequestOptions)
		if err == nil {
			t.Fatalf("expected error response")
		}
	}()

	select {
	case <-doneCh:
		//no-op
	case <-time.After(200 * time.Millisecond):
		t.Fatalf("should not retry on error")
	}

	// Create the mock handler. First we return error responses to ensure
	// that we power through and keep retrying in the face of recoverable
	// errors.
	code := int64(errorCode)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(int(atomic.LoadInt64(&code)))
	})

	// Create a test server
	list, err := net.Listen("tcp", ":28934")
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	defer list.Close()
	go http.Serve(list, handler)

	// Wait again
	select {
	case <-doneCh:
		//no-op
	case <-time.After(200 * time.Millisecond):
		t.Fatalf("should not retry on error")
	}

	// Start returning 200's
	atomic.StoreInt64(&code, http.StatusOK)

	// Wait again
	select {
	case <-doneCh:
	case <-time.After(time.Second):
		t.Fatalf("timed out")
	}

	if retryCount > 0 {
		t.Fatal("request log hook should not have been called")
	}
}
