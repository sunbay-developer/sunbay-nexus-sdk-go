package request

import "github.com/sunbay-developer/sunbay-nexus-sdk-go/model/common"

// RefundRequest represents a refund transaction request
type RefundRequest struct {
	// AppID is the application ID
	AppID string `json:"appId"`

	// MerchantID is the merchant ID
	MerchantID string `json:"merchantId"`

	// OriginalTransactionID is the original transaction ID to refund. Either originalTransactionId or originalTransactionRequestId is required for refund with reference. Both must be empty for refund without reference. If both are provided, originalTransactionId takes priority
	OriginalTransactionID string `json:"originalTransactionId,omitempty"`

	// OriginalTransactionRequestID is the original transaction request ID to refund. Either originalTransactionId or originalTransactionRequestId is required for refund with reference. Both must be empty for refund without reference. If both are provided, originalTransactionId takes priority
	OriginalTransactionRequestID string `json:"originalTransactionRequestId,omitempty"`

	// ReferenceOrderID is the reference order ID. Required for refund without reference, used to associate business records in merchant system. Not required for refund with reference, system will automatically associate with original transaction's reference order ID
	ReferenceOrderID string `json:"referenceOrderId,omitempty"`

	// TransactionRequestID is the transaction request ID for this refund transaction. Unique ID to identify this refund request, used as API idempotency control field
	TransactionRequestID string `json:"transactionRequestId"`

	// Amount is the amount information
	Amount *common.RefundAmount `json:"amount"`

	// PaymentMethod is the payment method information. Only available for refund without reference. Optional, recommended to omit for maximum flexibility
	PaymentMethod *common.PaymentMethodInfo `json:"paymentMethod,omitempty"`

	// Description is the refund reason description. Should be a real description representing the refund reason
	Description string `json:"description"`

	// TerminalSN is the terminal serial number. SUNBAY provided financial POS device serial number for reading bank cards and processing PIN security operations
	TerminalSN string `json:"terminalSn"`

	// Attach is additional data, returned as-is, recommended to use JSON format
	Attach string `json:"attach,omitempty"`

	// NotifyURL is the asynchronous notification URL
	NotifyURL string `json:"notifyUrl,omitempty"`

	// TimeExpire is the transaction expiration time, format: yyyy-MM-DDTHH:mm:ss+TIMEZONE (ISO 8601). Transaction will be closed if payment is not completed after this time. Minimum 3 minutes, maximum 1 day, default 1 day if not provided. Only used for refund without reference (requires customer operation on terminal), not needed for refund with reference
	TimeExpire string `json:"timeExpire,omitempty"`
}

