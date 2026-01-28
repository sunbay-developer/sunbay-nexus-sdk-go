package request

import "github.com/sunbay-developer/sunbay-nexus-sdk-go/model/common"

// IncrementalAuthRequest represents an incremental authorization transaction request
type IncrementalAuthRequest struct {
	// AppID is the application ID
	AppID string `json:"appId"`

	// MerchantID is the merchant ID
	MerchantID string `json:"merchantId"`

	// OriginalTransactionID is the original authorization transaction ID to increase authorization amount. Either originalTransactionId or originalTransactionRequestId is required. If both are provided, originalTransactionId takes priority
	OriginalTransactionID string `json:"originalTransactionId,omitempty"`

	// OriginalTransactionRequestID is the original authorization transaction request ID to increase authorization amount. Either originalTransactionId or originalTransactionRequestId is required. If both are provided, originalTransactionId takes priority
	OriginalTransactionRequestID string `json:"originalTransactionRequestId,omitempty"`

	// TransactionRequestID is the transaction request ID for this incremental authorization transaction. Unique ID to identify this incremental authorization request, used as API idempotency control field
	TransactionRequestID string `json:"transactionRequestId"`

	// Amount is the amount information
	Amount *common.AuthAmount `json:"amount"`

	// Description is the product description. Should be a real description representing the product information, may be displayed on some payment App billing pages
	Description string `json:"description"`

	// TerminalSN is the terminal serial number. SUNBAY provided financial POS device serial number for reading bank cards and processing PIN security operations
	TerminalSN string `json:"terminalSn"`

	// Attach is additional data, returned as-is, recommended to use JSON format
	Attach string `json:"attach,omitempty"`

	// NotifyURL is the asynchronous notification URL
	NotifyURL string `json:"notifyUrl,omitempty"`

	// PrintReceipt is the receipt print option. Possible values: NONE (do not print), MERCHANT (print merchant copy only), CUSTOMER (print customer copy only), BOTH (print both copies). Default: "NONE"
	PrintReceipt string `json:"printReceipt,omitempty"`
}

