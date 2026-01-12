package request

// BatchCloseRequest represents a batch close transaction request
type BatchCloseRequest struct {
	// AppID is the application ID
	AppID string `json:"appId"`

	// MerchantID is the merchant ID
	MerchantID string `json:"merchantId"`

	// TransactionRequestID is the batch close request unique identifier. Unique ID to identify this batch close request, used as API idempotency control field, can be used later to query batch close results
	TransactionRequestID string `json:"transactionRequestId"`

	// TerminalSN is the terminal serial number. SUNBAY provided financial POS device serial number for reading bank cards and processing PIN security operations
	TerminalSN string `json:"terminalSn"`

	// Description is the batch close description
	Description string `json:"description"`

	// ChannelCode is the channel code (optional)
	ChannelCode string `json:"channelCode,omitempty"`

	// Attach is additional data, returned as-is, recommended to use JSON format
	Attach string `json:"attach,omitempty"`
}
