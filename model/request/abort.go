package request

// AbortRequest represents an abort transaction request
type AbortRequest struct {
	// AppID is the application ID
	AppID string `json:"appId"`

	// MerchantID is the merchant ID
	MerchantID string `json:"merchantId"`

	// OriginalTransactionID is the original transaction ID to abort. Either originalTransactionId or originalTransactionRequestId is required. If both are provided, originalTransactionId takes priority
	OriginalTransactionID string `json:"originalTransactionId,omitempty"`

	// OriginalTransactionRequestID is the original transaction request ID to abort. Either originalTransactionId or originalTransactionRequestId is required. If both are provided, originalTransactionId takes priority
	OriginalTransactionRequestID string `json:"originalTransactionRequestId,omitempty"`

	// TerminalSN is the terminal serial number. SUNBAY provided financial POS device serial number for reading bank cards and processing PIN security operations
	TerminalSN string `json:"terminalSn"`

	// Description is the abort reason description
	Description string `json:"description"`

	// Attach is additional data, returned as-is, recommended to use JSON format
	Attach string `json:"attach,omitempty"`
}

