package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/sunbay-developer/sunbay-nexus-sdk-go/constant"
	"github.com/sunbay-developer/sunbay-nexus-sdk-go/exception"
	"github.com/sunbay-developer/sunbay-nexus-sdk-go/model/common"
	"github.com/sunbay-developer/sunbay-nexus-sdk-go/util"
)

const (
	headerAuthorization = "Authorization"
	headerContentType   = "Content-Type"
	headerRequestID     = "X-Client-Request-Id"
	headerTimestamp     = "X-Timestamp"
	headerUserAgent     = "User-Agent"
	contentTypeJSON     = "application/json"
	retryDelayBase      = 1 * time.Second
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

// defaultLogger is the default logger
type defaultLogger struct{}

func (l *defaultLogger) Debug(args ...interface{}) {
	fmt.Println("[DEBUG]", fmt.Sprint(args...))
}

func (l *defaultLogger) Info(args ...interface{}) {
	fmt.Println("[INFO]", fmt.Sprint(args...))
}

func (l *defaultLogger) Warn(args ...interface{}) {
	fmt.Println("[WARN]", fmt.Sprint(args...))
}

func (l *defaultLogger) Error(args ...interface{}) {
	fmt.Println("[ERROR]", fmt.Sprint(args...))
}

// getLogger gets the logger (returns default logger if nil)
func getLogger(logger Logger) Logger {
	if logger == nil {
		return &defaultLogger{}
	}
	return logger
}

// Client is the HTTP client
type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
	maxRetries int
	retryDelay time.Duration
	logger     Logger
}

// NewClient creates a new HTTP client
func NewClient(apiKey, baseURL string, connectTimeout, readTimeout, maxRetries, maxTotal, maxPerRoute int, logger Logger) *Client {
	transport := &http.Transport{
		MaxIdleConns:        maxTotal,
		MaxIdleConnsPerHost: maxPerRoute,
		IdleConnTimeout:     90 * time.Second,
	}

	httpClient := &http.Client{
		Transport: transport,
		Timeout:   time.Duration(readTimeout) * time.Millisecond,
	}

	return &Client{
		apiKey:     apiKey,
		baseURL:    baseURL,
		httpClient: httpClient,
		maxRetries: maxRetries,
		retryDelay: retryDelayBase,
		logger:     logger,
	}
}

// Post executes a POST request
func (c *Client) Post(path string, requestBody interface{}, responseType interface{}) error {
	url := c.baseURL + path
	requestJSON := util.ToJSON(requestBody)

	req, err := http.NewRequest("POST", url, strings.NewReader(requestJSON))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	c.addCommonHeaders(req, "POST")

	// Log request
	headers := make(map[string]string)
	for k, v := range req.Header {
		if len(v) > 0 {
			headers[k] = v[0]
		}
	}
	c.logRequest("POST", url, headers, requestJSON)

	return c.executeRequest(req, responseType, false)
}

// Get executes a GET request
func (c *Client) Get(path string, request interface{}, responseType interface{}) error {
	baseURL := c.baseURL + path
	urlStr := c.buildQueryURL(baseURL, request)

	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	c.addCommonHeaders(req, "GET")

	// Log request
	headers := make(map[string]string)
	for k, v := range req.Header {
		if len(v) > 0 {
			headers[k] = v[0]
		}
	}
	c.logRequest("GET", urlStr, headers, "")

	return c.executeRequest(req, responseType, true)
}

// addCommonHeaders adds common request headers
func (c *Client) addCommonHeaders(req *http.Request, method string) {
	req.Header.Set(headerAuthorization, constant.AuthorizationBearerPrefix+c.apiKey)
	req.Header.Set(headerRequestID, util.GenerateRequestID())
	req.Header.Set(headerTimestamp, strconv.FormatInt(time.Now().UnixMilli(), 10))
	req.Header.Set(headerUserAgent, util.GetUserAgent())

	if method == "POST" {
		req.Header.Set(headerContentType, contentTypeJSON)
	}
}

// buildQueryURL builds query URL
func (c *Client) buildQueryURL(baseURL string, request interface{}) string {
	if request == nil {
		return baseURL
	}

	u, err := url.Parse(baseURL)
	if err != nil {
		return baseURL
	}

	q := u.Query()

	// Extract query parameters using reflection
	v := reflect.ValueOf(request)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return baseURL
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// Get JSON tag
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		// Remove omitempty and other options
		jsonTag = strings.Split(jsonTag, ",")[0]

		// Skip zero values
		if !value.IsValid() || value.IsZero() {
			continue
		}

		// Convert to string
		var strValue string
		switch value.Kind() {
		case reflect.String:
			strValue = value.String()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			strValue = strconv.FormatInt(value.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			strValue = strconv.FormatUint(value.Uint(), 10)
		case reflect.Float32, reflect.Float64:
			strValue = strconv.FormatFloat(value.Float(), 'f', -1, 64)
		case reflect.Bool:
			strValue = strconv.FormatBool(value.Bool())
		default:
			continue
		}

		if strValue != "" {
			q.Set(jsonTag, strValue)
		}
	}

	u.RawQuery = q.Encode()
	return u.String()
}

// executeRequest executes HTTP request with retry
func (c *Client) executeRequest(req *http.Request, responseType interface{}, retryable bool) error {
	maxAttempts := 1
	if retryable {
		maxAttempts = c.maxRetries
	}

	var lastErr error
	method := req.Method
	urlStr := req.URL.String()

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		resp, err := c.httpClient.Do(req)
		if err != nil {
			lastErr = err
			if !retryable || attempt >= maxAttempts {
				c.logError(method, urlStr, err)
				return exception.NewNetworkError("Network error: "+err.Error(), true, err)
			}
			c.logRetry(attempt, maxAttempts, err.Error())
			time.Sleep(c.retryDelay * time.Duration(attempt))
			continue
		}

		// Read response body for logging
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			lastErr = err
			if !retryable || attempt >= maxAttempts {
				c.logError(method, urlStr, err)
				return exception.NewNetworkError("Failed to read response body: "+err.Error(), false, err)
			}
			c.logRetry(attempt, maxAttempts, err.Error())
			time.Sleep(c.retryDelay * time.Duration(attempt))
			continue
		}
		resp.Body.Close()
		bodyStr := string(bodyBytes)

		// Log response
		c.logResponse(method, urlStr, resp.StatusCode, bodyStr)

		// Recreate response body for parsing
		resp.Body = io.NopCloser(bytes.NewReader(bodyBytes))

		// Parse response
		err = c.parseResponse(resp, responseType)
		resp.Body.Close()

		if err != nil {
			// If network error and retryable, continue retrying
			if netErr, ok := err.(*exception.NetworkError); ok && netErr.IsRetryable() && retryable && attempt < maxAttempts {
				lastErr = err
				c.logRetry(attempt, maxAttempts, err.Error())
				time.Sleep(c.retryDelay * time.Duration(attempt))
				continue
			}
			c.logError(method, urlStr, err)
			return err
		}

		return nil
	}

	finalErr := exception.NewNetworkError(
		fmt.Sprintf("Request failed after %d attempts", maxAttempts),
		retryable,
		lastErr,
	)
	c.logError(method, urlStr, finalErr)
	return finalErr
}

// parseResponse parses HTTP response
func (c *Client) parseResponse(resp *http.Response, result interface{}) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return exception.NewNetworkError("Failed to read response body", false, err)
	}

	// Check HTTP status code
	if resp.StatusCode < constant.HTTPStatusOKStart || resp.StatusCode >= constant.HTTPStatusOKEnd {
		return exception.NewNetworkError(
			fmt.Sprintf("HTTP %d: %s", resp.StatusCode, string(body)),
			resp.StatusCode >= constant.HTTPStatusServerErrorStart,
			nil,
		)
	}

	// Parse JSON (supports data field wrapping)
	var wrapper struct {
		Code    string          `json:"code"`
		Msg     string          `json:"msg"`
		TraceID string          `json:"traceId"`
		Data    json.RawMessage `json:"data"`
	}

	if err := json.Unmarshal(body, &wrapper); err != nil {
		return exception.NewNetworkError("Failed to parse response", false, err)
	}

	// If data field exists, parse data; otherwise parse entire response
	var dataToParse []byte
	if len(wrapper.Data) > 0 && string(wrapper.Data) != "null" {
		dataToParse = wrapper.Data
	} else {
		dataToParse = body
	}

	if err := json.Unmarshal(dataToParse, result); err != nil {
		return exception.NewNetworkError("Failed to unmarshal result", false, err)
	}

	// Set base fields
	if baseResp, ok := result.(interface {
		SetCode(code string)
		SetMsg(msg string)
		SetTraceID(traceID string)
	}); ok {
		baseResp.SetCode(wrapper.Code)
		baseResp.SetMsg(wrapper.Msg)
		baseResp.SetTraceID(wrapper.TraceID)
	}

	// Check business error
	if baseResp, ok := result.(*common.BaseResponse); ok {
		if !baseResp.IsSuccess() {
			return exception.NewBusinessError(
				baseResp.Code,
				baseResp.Msg,
				baseResp.TraceID,
			)
		}
	}

	return nil
}

// logRequest logs request
func (c *Client) logRequest(method, url string, headers map[string]string, body string) {
	logger := getLogger(c.logger)

	// Mask Authorization header
	maskedHeaders := make(map[string]string)
	for k, v := range headers {
		if k == headerAuthorization {
			maskedHeaders[k] = maskAuthorization(v)
		} else {
			maskedHeaders[k] = v
		}
	}

	// Convert headers to JSON format
	headersJSON, _ := json.Marshal(maskedHeaders)

	if body != "" {
		logger.Info(fmt.Sprintf("Request %s %s - Headers: %s - Body: %s", method, url, string(headersJSON), body))
	} else {
		logger.Info(fmt.Sprintf("Request %s %s - Headers: %s", method, url, string(headersJSON)))
	}
}

// logResponse logs response
func (c *Client) logResponse(method, url string, statusCode int, body string) {
	logger := getLogger(c.logger)
	logger.Info(fmt.Sprintf("Response %s %s - Status: %d, Body: %s", method, url, statusCode, body))
}

// logRetry logs retry
func (c *Client) logRetry(attempt, maxAttempts int, reason string) {
	logger := getLogger(c.logger)
	logger.Debug(fmt.Sprintf("Request failed, retrying (%d/%d) after delay: %s", attempt, maxAttempts, reason))
}

// logError logs error
func (c *Client) logError(method, url string, err error) {
	logger := getLogger(c.logger)
	if netErr, ok := err.(*exception.NetworkError); ok {
		logger.Warn(fmt.Sprintf("Network error %s %s: %v (retryable: %v)", method, url, netErr, netErr.IsRetryable()))
	} else if bizErr, ok := err.(*exception.BusinessError); ok {
		logger.Error(fmt.Sprintf("API error %s %s - code: %s, msg: %s, traceID: %s",
			method, url, bizErr.Code(), bizErr.Message(), bizErr.TraceID()))
	} else {
		logger.Error(fmt.Sprintf("Unknown error %s %s: %v", method, url, err))
	}
}

// maskAuthorization masks Authorization header
func maskAuthorization(authValue string) string {
	if authValue == "" {
		return ""
	}

	prefix := constant.AuthorizationBearerPrefix
	if strings.HasPrefix(authValue, prefix) {
		token := authValue[len(prefix):]
		if len(token) > 8 {
			return prefix + token[:4] + "****" + token[len(token)-4:]
		}
		return prefix + "****"
	}

	return "****"
}

// Close closes HTTP client and releases resources
func (c *Client) Close() error {
	if transport, ok := c.httpClient.Transport.(*http.Transport); ok {
		transport.CloseIdleConnections()
	}
	return nil
}

