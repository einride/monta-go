package monta

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"gotest.tools/v3/assert"
)

// roundTripFunc adapts a function to an [http.RoundTripper].
type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func newTestResponse(statusCode int, body string) *http.Response {
	return &http.Response{
		StatusCode: statusCode,
		Status:     http.StatusText(statusCode),
		Body:       io.NopCloser(strings.NewReader(body)),
		Header:     make(http.Header),
	}
}

func newTestClient(transport http.RoundTripper, maxRetries int) *clientImpl {
	return &clientImpl{
		httpClient: &http.Client{Transport: transport},
		config: clientConfig{
			maxRetries:     maxRetries,
			retryBaseDelay: time.Millisecond,
		},
	}
}

func closeResponse(resp *http.Response) {
	if resp != nil {
		_ = resp.Body.Close()
	}
}

func TestDoWithRetry_RetriesGETOn5xxThenSucceeds(t *testing.T) {
	var calls atomic.Int32
	client := newTestClient(roundTripFunc(func(*http.Request) (*http.Response, error) {
		if calls.Add(1) <= 2 {
			return newTestResponse(http.StatusInternalServerError, `{"message":"Read Timeout"}`), nil
		}
		return newTestResponse(http.StatusOK, `{}`), nil
	}), 3)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, apiHost+"/foo", nil)
	assert.NilError(t, err)
	resp, err := client.doWithRetry(context.Background(), req, http.MethodGet)
	defer closeResponse(resp)
	assert.NilError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, int32(3), calls.Load())
}

func TestDoWithRetry_RetriesGETOn429(t *testing.T) {
	var calls atomic.Int32
	client := newTestClient(roundTripFunc(func(*http.Request) (*http.Response, error) {
		if calls.Add(1) == 1 {
			return newTestResponse(http.StatusTooManyRequests, "slow down"), nil
		}
		return newTestResponse(http.StatusOK, `{}`), nil
	}), 3)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, apiHost+"/foo", nil)
	assert.NilError(t, err)
	resp, err := client.doWithRetry(context.Background(), req, http.MethodGet)
	defer closeResponse(resp)
	assert.NilError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, int32(2), calls.Load())
}

func TestDoWithRetry_ExhaustsRetriesAndReturnsStatusError(t *testing.T) {
	var calls atomic.Int32
	client := newTestClient(roundTripFunc(func(*http.Request) (*http.Response, error) {
		calls.Add(1)
		return newTestResponse(http.StatusInternalServerError, "boom"), nil
	}), 3)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, apiHost+"/foo", nil)
	assert.NilError(t, err)
	resp, err := client.doWithRetry(context.Background(), req, http.MethodGet)
	defer closeResponse(resp)
	var statusErr *StatusError
	assert.Assert(t, errors.As(err, &statusErr))
	assert.Equal(t, http.StatusInternalServerError, statusErr.StatusCode)
	// 1 initial attempt + 3 retries.
	assert.Equal(t, int32(4), calls.Load())
}

func TestDoWithRetry_DoesNotRetryNon5xx(t *testing.T) {
	var calls atomic.Int32
	client := newTestClient(roundTripFunc(func(*http.Request) (*http.Response, error) {
		calls.Add(1)
		return newTestResponse(http.StatusNotFound, "not found"), nil
	}), 3)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, apiHost+"/foo", nil)
	assert.NilError(t, err)
	resp, err := client.doWithRetry(context.Background(), req, http.MethodGet)
	defer closeResponse(resp)
	var statusErr *StatusError
	assert.Assert(t, errors.As(err, &statusErr))
	assert.Equal(t, http.StatusNotFound, statusErr.StatusCode)
	assert.Equal(t, int32(1), calls.Load())
}

func TestDoWithRetry_DoesNotRetryNonGET(t *testing.T) {
	var calls atomic.Int32
	client := newTestClient(roundTripFunc(func(*http.Request) (*http.Response, error) {
		calls.Add(1)
		return newTestResponse(http.StatusInternalServerError, "boom"), nil
	}), 3)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, apiHost+"/foo", nil)
	assert.NilError(t, err)
	resp, err := client.doWithRetry(context.Background(), req, http.MethodPost)
	defer closeResponse(resp)
	var statusErr *StatusError
	assert.Assert(t, errors.As(err, &statusErr))
	assert.Equal(t, int32(1), calls.Load())
}

func TestDoWithRetry_RetriesNetworkErrorOnGET(t *testing.T) {
	var calls atomic.Int32
	netErr := errors.New("connection reset")
	client := newTestClient(roundTripFunc(func(*http.Request) (*http.Response, error) {
		if calls.Add(1) == 1 {
			return nil, netErr
		}
		return newTestResponse(http.StatusOK, `{}`), nil
	}), 3)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, apiHost+"/foo", nil)
	assert.NilError(t, err)
	resp, err := client.doWithRetry(context.Background(), req, http.MethodGet)
	defer closeResponse(resp)
	assert.NilError(t, err)
	assert.Equal(t, int32(2), calls.Load())
}

func TestDoWithRetry_Disabled(t *testing.T) {
	var calls atomic.Int32
	client := newTestClient(roundTripFunc(func(*http.Request) (*http.Response, error) {
		calls.Add(1)
		return newTestResponse(http.StatusInternalServerError, "boom"), nil
	}), 0)
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, apiHost+"/foo", nil)
	assert.NilError(t, err)
	resp, err := client.doWithRetry(context.Background(), req, http.MethodGet)
	defer closeResponse(resp)
	assert.Assert(t, err != nil)
	assert.Equal(t, int32(1), calls.Load())
}

func TestDoWithRetry_RespectsContextCancellation(t *testing.T) {
	var calls atomic.Int32
	client := newTestClient(roundTripFunc(func(*http.Request) (*http.Response, error) {
		calls.Add(1)
		return newTestResponse(http.StatusInternalServerError, "boom"), nil
	}), 5)
	client.config.retryBaseDelay = time.Hour // force the backoff to block until cancellation
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiHost+"/foo", nil)
	assert.NilError(t, err)
	resp, err := client.doWithRetry(ctx, req, http.MethodGet)
	defer closeResponse(resp)
	assert.Assert(t, errors.Is(err, context.Canceled))
	// One attempt before the (cancelled) backoff wait stops the loop.
	assert.Equal(t, int32(1), calls.Load())
}

func TestBackoffDelay(t *testing.T) {
	base := 100 * time.Millisecond
	assert.Equal(t, base, backoffDelay(base, 1))
	assert.Equal(t, 2*base, backoffDelay(base, 2))
	assert.Equal(t, 4*base, backoffDelay(base, 3))
	// Exponent is capped to avoid overflow on large attempt counts.
	assert.Equal(t, 64*base, backoffDelay(base, 100))
}

func TestNewClient_RetryDefaults(t *testing.T) {
	client, ok := NewClient().(*clientImpl)
	assert.Assert(t, ok)
	assert.Equal(t, defaultMaxRetries, client.config.maxRetries)
	assert.Equal(t, defaultRetryBaseDelay, client.config.retryBaseDelay)
}

func TestWithRetry(t *testing.T) {
	client, ok := NewClient(WithRetry(1, 5*time.Second)).(*clientImpl)
	assert.Assert(t, ok)
	assert.Equal(t, 1, client.config.maxRetries)
	assert.Equal(t, 5*time.Second, client.config.retryBaseDelay)
}
