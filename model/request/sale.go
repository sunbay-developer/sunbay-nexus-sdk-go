package request

import "github.com/sunbay-developer/sunbay-nexus-sdk-go/model/common"

// SaleRequest represents a sale transaction request
type SaleRequest struct {
	// AppID is the application ID
	AppID string `json:"appId"`

	// MerchantID is the merchant ID
	MerchantID string `json:"merchantId"`

	// ReferenceOrderID is the reference order ID for the sale transaction. Unique ID assigned by merchant system to identify this sale transaction, 6-32 characters, can only contain numbers, uppercase/lowercase letters, _-\|*
	ReferenceOrderID string `json:"referenceOrderId"`

	// TransactionRequestID is the unique request identifier for this sale transaction. Used as the API idempotency key to ensure that retrying the same request does not create multiple transactions. Can only contain letters, digits, underscore (_) and hyphen (-), with a maximum length of 64 characters
	TransactionRequestID string `json:"transactionRequestId"`

	// Amount is the amount information
	Amount *common.SaleAmount `json:"amount"`

	// PaymentMethod is the payment method information. Optional, recommended to omit for maximum flexibility
	PaymentMethod *common.PaymentMethodInfo `json:"paymentMethod,omitempty"`

	// Description is the product description. Should be a real description representing the product information, may be displayed on some payment App billing pages
	Description string `json:"description"`

	// TerminalSN is the terminal serial number. SUNBAY provided financial POS device serial number for reading bank cards and processing PIN security operations
	TerminalSN string `json:"terminalSn"`

	// Attach is additional data, returned as-is, recommended to use JSON format
	Attach string `json:"attach,omitempty"`

	// NotifyURL is the asynchronous notification URL
	NotifyURL string `json:"notifyUrl,omitempty"`

	// TimeExpire is the transaction expiration time, format: yyyy-MM-DDTHH:mm:ss+TIMEZONE (ISO 8601). Transaction will be closed if payment is not completed after this time. Minimum 3 minutes, maximum 1 day, default 1 day if not provided
	TimeExpire string `json:"timeExpire,omitempty"`
}

