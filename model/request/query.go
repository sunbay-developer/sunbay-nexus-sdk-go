package request

// QueryRequest represents a query transaction request
type QueryRequest struct {
	// AppID is the application ID
	AppID string `json:"appId"`

	// MerchantID is the merchant ID
	MerchantID string `json:"merchantId"`

	// TransactionID is the SUNBAY Nexus transaction ID. Either transactionId or transactionRequestId is required. If both are provided, transactionId takes priority
	TransactionID string `json:"transactionId,omitempty"`

	// TransactionRequestID is the transaction request ID. Either transactionId or transactionRequestId is required. If both are provided, transactionId takes priority
	TransactionRequestID string `json:"transactionRequestId,omitempty"`
}

