package request

// TipAdjustRequest represents a tip adjust transaction request
type TipAdjustRequest struct {
	// AppID is the application ID
	AppID string `json:"appId"`

	// MerchantID is the merchant ID
	MerchantID string `json:"merchantId"`

	// TerminalSN is the terminal serial number
	TerminalSN string `json:"terminalSn"`

	// OriginalTransactionID is the original transaction ID to adjust tip. Either originalTransactionId or originalTransactionRequestId is required. If both are provided, originalTransactionId takes priority
	OriginalTransactionID string `json:"originalTransactionId,omitempty"`

	// OriginalTransactionRequestID is the original transaction request ID to adjust tip. Either originalTransactionId or originalTransactionRequestId is required. If both are provided, originalTransactionId takes priority
	OriginalTransactionRequestID string `json:"originalTransactionRequestId,omitempty"`

	// TipAmount is the new tip amount after adjustment, in basic currency unit
	TipAmount *float64 `json:"tipAmount"`

	// Attach is additional data, returned as-is, recommended to use JSON format
	Attach string `json:"attach,omitempty"`
}

