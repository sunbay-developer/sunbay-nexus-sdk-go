package constant

// ErrorCodeParameterError is the parameter error code (C17)
const ErrorCodeParameterError = "C17"

// HTTP methods
const (
	HTTPMethodPOST = "POST"
	HTTPMethodGET  = "GET"
)

// HTTP status codes
const (
	HTTPStatusOKStart          = 200
	HTTPStatusOKEnd            = 300
	HTTPStatusClientErrorStart = 400
	HTTPStatusClientErrorEnd   = 500
	HTTPStatusServerErrorStart = 500
)

// ResponseSuccessCode is the response success code
const ResponseSuccessCode = "0"

// AuthorizationBearerPrefix is the Authorization header prefix
const AuthorizationBearerPrefix = "Bearer "

// JSON field names
const (
	JSONFieldCode    = "code"
	JSONFieldMsg     = "msg"
	JSONFieldData    = "data"
	JSONFieldTraceID = "traceId"
)

// GetterMethodPrefixLength is the getter method name prefix length
const GetterMethodPrefixLength = 3

// API path prefixes
const (
	SemiIntegrationPrefix = "/v1/semi-integration"
	CommonPrefix          = "/v1"
)

// Semi-integration transaction API paths
const (
	PathSale            = SemiIntegrationPrefix + "/transaction/sale"
	PathAuth            = SemiIntegrationPrefix + "/transaction/auth"
	PathForcedAuth      = SemiIntegrationPrefix + "/transaction/forced-auth"
	PathIncrementalAuth = SemiIntegrationPrefix + "/transaction/incremental-auth"
	PathPostAuth        = SemiIntegrationPrefix + "/transaction/post-auth"
	PathRefund          = SemiIntegrationPrefix + "/transaction/refund"
	PathVoid            = SemiIntegrationPrefix + "/transaction/void"
	PathAbort           = SemiIntegrationPrefix + "/transaction/abort"
	PathTipAdjust       = SemiIntegrationPrefix + "/transaction/tip-adjust"
	PathQuery           = CommonPrefix + "/transaction/query"
)

// Semi-integration settlement API paths
const (
	PathBatchClose = CommonPrefix + "/settlement/batch-close"
	PathBatchQuery = CommonPrefix + "/settlement/batch-query"
)
