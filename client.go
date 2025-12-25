package nexus

import (
	"context"
	"time"

	"github.com/sunbay-developer/sunbay-nexus-sdk-go/constant"
	"github.com/sunbay-developer/sunbay-nexus-sdk-go/errors"
	"github.com/sunbay-developer/sunbay-nexus-sdk-go/http"
	"github.com/sunbay-developer/sunbay-nexus-sdk-go/model/request"
	"github.com/sunbay-developer/sunbay-nexus-sdk-go/model/response"
)

const (
	defaultBaseURL        = "https://open.sunbay.us"
	defaultConnectTimeout = 30 * time.Second
	defaultReadTimeout    = 60 * time.Second
	defaultMaxRetries     = 3
	defaultMaxTotal        = 200
	defaultMaxPerRoute     = 20
)

// Logger is the logging interface that allows integration with any logging library
// Users can implement this interface for logrus, zap, zerolog, or any other logging library
// SDK uses this interface internally, allowing users to inject their preferred logging implementation
type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
}

// NexusClient is the main client for Sunbay Nexus SDK
// The client is thread-safe and can be safely used by multiple goroutines
type NexusClient struct {
	httpClient *http.Client
}

// Config holds the configuration for creating a NexusClient
type Config struct {
	// APIKey is the API key for authentication (required)
	APIKey string

	// BaseURL is the base URL for API requests (optional, defaults to "https://open.sunbay.us")
	BaseURL string

	// ConnectTimeout is the connection timeout (optional, defaults to 30s)
	ConnectTimeout time.Duration

	// ReadTimeout is the read timeout (optional, defaults to 60s)
	ReadTimeout time.Duration

	// MaxRetries is the maximum number of retries for GET requests (optional, defaults to 3)
	MaxRetries int

	// MaxTotal is the maximum total connections in the connection pool (optional, defaults to 200)
	MaxTotal int

	// MaxPerRoute is the maximum connections per route (optional, defaults to 20)
	MaxPerRoute int

	// Logger is a custom logger implementation (optional, defaults to console logger)
	Logger Logger
}

// NewNexusClient creates a new NexusClient with the given configuration
// This follows the common Go SDK pattern used by Stripe, Alibaba Cloud, etc.
func NewNexusClient(config *Config) (*NexusClient, error) {
	if config == nil {
		return nil, errors.NewBusinessError(
			constant.ErrorCodeParameterError,
			"Config cannot be nil",
			"",
		)
	}

	if config.APIKey == "" {
		return nil, errors.NewBusinessError(
			constant.ErrorCodeParameterError,
			"API key cannot be empty",
			"",
		)
	}

	// Apply defaults
	baseURL := config.BaseURL
	if baseURL == "" {
		baseURL = defaultBaseURL
	}

	connectTimeout := config.ConnectTimeout
	if connectTimeout == 0 {
		connectTimeout = defaultConnectTimeout
	}

	readTimeout := config.ReadTimeout
	if readTimeout == 0 {
		readTimeout = defaultReadTimeout
	}

	maxRetries := config.MaxRetries
	if maxRetries == 0 {
		maxRetries = defaultMaxRetries
	}

	maxTotal := config.MaxTotal
	if maxTotal == 0 {
		maxTotal = defaultMaxTotal
	}

	maxPerRoute := config.MaxPerRoute
	if maxPerRoute == 0 {
		maxPerRoute = defaultMaxPerRoute
	}

	// Create HTTP client wrapper
	httpClientWrapper := http.NewClient(
		config.APIKey,
		baseURL,
		int(connectTimeout.Milliseconds()),
		int(readTimeout.Milliseconds()),
		maxRetries,
		maxTotal,
		maxPerRoute,
		config.Logger,
	)

	return &NexusClient{
		httpClient: httpClientWrapper,
	}, nil
}

// Sale executes a sale transaction
func (c *NexusClient) Sale(ctx context.Context, req *request.SaleRequest) (*response.SaleResponse, error) {
	if req == nil {
		return nil, errors.NewBusinessError(
			constant.ErrorCodeParameterError,
			"SaleRequest cannot be nil",
			"",
		)
	}

	resp := &response.SaleResponse{}
	err := c.httpClient.Post(constant.PathSale, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Auth executes an authorization (pre-auth) transaction
func (c *NexusClient) Auth(ctx context.Context, req *request.AuthRequest) (*response.AuthResponse, error) {
	if req == nil {
		return nil, errors.NewBusinessError(
			constant.ErrorCodeParameterError,
			"AuthRequest cannot be nil",
			"",
		)
	}

	resp := &response.AuthResponse{}
	err := c.httpClient.Post(constant.PathAuth, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ForcedAuth executes a forced authorization transaction
func (c *NexusClient) ForcedAuth(ctx context.Context, req *request.ForcedAuthRequest) (*response.ForcedAuthResponse, error) {
	if req == nil {
		return nil, errors.NewBusinessError(
			constant.ErrorCodeParameterError,
			"ForcedAuthRequest cannot be nil",
			"",
		)
	}

	resp := &response.ForcedAuthResponse{}
	err := c.httpClient.Post(constant.PathForcedAuth, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// IncrementalAuth executes an incremental authorization transaction
func (c *NexusClient) IncrementalAuth(ctx context.Context, req *request.IncrementalAuthRequest) (*response.IncrementalAuthResponse, error) {
	if req == nil {
		return nil, errors.NewBusinessError(
			constant.ErrorCodeParameterError,
			"IncrementalAuthRequest cannot be nil",
			"",
		)
	}

	resp := &response.IncrementalAuthResponse{}
	err := c.httpClient.Post(constant.PathIncrementalAuth, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// PostAuth executes a post authorization (pre-auth completion) transaction
func (c *NexusClient) PostAuth(ctx context.Context, req *request.PostAuthRequest) (*response.PostAuthResponse, error) {
	if req == nil {
		return nil, errors.NewBusinessError(
			constant.ErrorCodeParameterError,
			"PostAuthRequest cannot be nil",
			"",
		)
	}

	resp := &response.PostAuthResponse{}
	err := c.httpClient.Post(constant.PathPostAuth, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Refund executes a refund transaction
func (c *NexusClient) Refund(ctx context.Context, req *request.RefundRequest) (*response.RefundResponse, error) {
	if req == nil {
		return nil, errors.NewBusinessError(
			constant.ErrorCodeParameterError,
			"RefundRequest cannot be nil",
			"",
		)
	}

	resp := &response.RefundResponse{}
	err := c.httpClient.Post(constant.PathRefund, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Void executes a void transaction
func (c *NexusClient) Void(ctx context.Context, req *request.VoidRequest) (*response.VoidResponse, error) {
	if req == nil {
		return nil, errors.NewBusinessError(
			constant.ErrorCodeParameterError,
			"VoidRequest cannot be nil",
			"",
		)
	}

	resp := &response.VoidResponse{}
	err := c.httpClient.Post(constant.PathVoid, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Abort executes an abort transaction
func (c *NexusClient) Abort(ctx context.Context, req *request.AbortRequest) (*response.AbortResponse, error) {
	if req == nil {
		return nil, errors.NewBusinessError(
			constant.ErrorCodeParameterError,
			"AbortRequest cannot be nil",
			"",
		)
	}

	resp := &response.AbortResponse{}
	err := c.httpClient.Post(constant.PathAbort, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// TipAdjust executes a tip adjust transaction
func (c *NexusClient) TipAdjust(ctx context.Context, req *request.TipAdjustRequest) (*response.TipAdjustResponse, error) {
	if req == nil {
		return nil, errors.NewBusinessError(
			constant.ErrorCodeParameterError,
			"TipAdjustRequest cannot be nil",
			"",
		)
	}

	resp := &response.TipAdjustResponse{}
	err := c.httpClient.Post(constant.PathTipAdjust, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Query queries a transaction
func (c *NexusClient) Query(ctx context.Context, req *request.QueryRequest) (*response.QueryResponse, error) {
	if req == nil {
		return nil, errors.NewBusinessError(
			constant.ErrorCodeParameterError,
			"QueryRequest cannot be nil",
			"",
		)
	}

	resp := &response.QueryResponse{}
	err := c.httpClient.Get(constant.PathQuery, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// BatchClose executes a batch close transaction
func (c *NexusClient) BatchClose(ctx context.Context, req *request.BatchCloseRequest) (*response.BatchCloseResponse, error) {
	if req == nil {
		return nil, errors.NewBusinessError(
			constant.ErrorCodeParameterError,
			"BatchCloseRequest cannot be nil",
			"",
		)
	}

	resp := &response.BatchCloseResponse{}
	err := c.httpClient.Post(constant.PathBatchClose, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

