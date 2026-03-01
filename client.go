package aurora

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	defaultBaseURL = "https://localhost:3000"
	basePath       = "/api/v1"
	defaultTimeout = 30 * time.Second
)

// Client is the Aurora API client.
type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client

	Process *ProcessService
}

// Option is a functional option for configuring the Client.
type Option func(*Client)

// WithBaseURL sets the host URL (e.g. "https://api.example.com").
// The API path (/api/v1) is appended automatically.
func WithBaseURL(url string) Option {
	return func(c *Client) {
		c.baseURL = strings.TrimRight(url, "/")
	}
}

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) {
		c.httpClient = hc
	}
}

// WithTimeout sets the HTTP client timeout.
func WithTimeout(d time.Duration) Option {
	return func(c *Client) {
		c.httpClient.Timeout = d
	}
}

// NewClient creates a new Aurora API client.
func NewClient(apiKey string, opts ...Option) *Client {
	c := &Client{
		baseURL:    defaultBaseURL,
		apiKey:     apiKey,
		httpClient: &http.Client{Timeout: defaultTimeout},
	}

	for _, opt := range opts {
		opt(c)
	}

	c.Process = &ProcessService{client: c}

	return c
}

// newRequest builds an authenticated HTTP request.
func (c *Client) newRequest(ctx context.Context, method, path string, body any) (*http.Request, error) {
	url := c.baseURL + basePath + path

	var buf io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("aurora: marshal request body: %w", err)
		}
		buf = bytes.NewReader(data)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, buf)
	if err != nil {
		return nil, fmt.Errorf("aurora: create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	return req, nil
}

const maxResponseSize = 10 << 20 // 10 MB

// drainBody reads and discards remaining body bytes so the underlying
// TCP connection can be reused by the HTTP transport pool.
func drainBody(body io.ReadCloser) {
	_, _ = io.Copy(io.Discard, body)
	body.Close()
}

// do executes an HTTP request and decodes the response into v.
func (c *Client) do(req *http.Request, v any) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("aurora: send request: %w", err)
	}
	defer drainBody(resp.Body)

	limited := io.LimitReader(resp.Body, maxResponseSize)
	respBody, err := io.ReadAll(limited)
	if err != nil {
		return fmt.Errorf("aurora: read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		var apiErr ErrorResponse
		if err := json.Unmarshal(respBody, &apiErr); err == nil && apiErr.Error != "" {
			return &APIError{
				StatusCode: resp.StatusCode,
				Message:    apiErr.Error,
			}
		}
		return &APIError{
			StatusCode: resp.StatusCode,
			Message:    string(respBody),
		}
	}

	if v != nil {
		if err := json.Unmarshal(respBody, v); err != nil {
			return fmt.Errorf("aurora: decode response: %w", err)
		}
	}

	return nil
}
