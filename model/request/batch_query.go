package request

// BatchQueryRequest represents a batch query request
type BatchQueryRequest struct {
	// AppID is the application ID
	AppID string `json:"appId"`

	// MerchantID is the merchant ID
	MerchantID string `json:"merchantId"`

	// TerminalSN is the terminal serial number. SUNBAY provided financial POS device serial number
	TerminalSN string `json:"terminalSn"`
}
