package request

// VoidRequest represents a void transaction request
type VoidRequest struct {
	// AppID is the application ID
	AppID string `json:"appId"`

	// MerchantID is the merchant ID
	MerchantID string `json:"merchantId"`

	// OriginalTransactionID is the original transaction ID to void. Either originalTransactionId or originalTransactionRequestId is required. If both are provided, originalTransactionId takes priority
	OriginalTransactionID string `json:"originalTransactionId,omitempty"`

	// OriginalTransactionRequestID is the original transaction request ID to void. Either originalTransactionId or originalTransactionRequestId is required. If both are provided, originalTransactionId takes priority
	OriginalTransactionRequestID string `json:"originalTransactionRequestId,omitempty"`

	// TransactionRequestID is the transaction request ID for this void transaction. Unique ID to identify this void request, used as API idempotency control field
	TransactionRequestID string `json:"transactionRequestId"`

	// Description is the void reason description
	Description string `json:"description"`

	// TerminalSN is the terminal serial number. SUNBAY provided financial POS device serial number for reading bank cards and processing PIN security operations
	TerminalSN string `json:"terminalSn"`

	// Attach is additional data, returned as-is, recommended to use JSON format
	Attach string `json:"attach,omitempty"`

	// NotifyURL is the asynchronous notification URL
	NotifyURL string `json:"notifyUrl,omitempty"`
}

