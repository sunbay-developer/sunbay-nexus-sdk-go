package request

import "github.com/sunbay-developer/sunbay-nexus-sdk-go/model/common"

// OnlineRefundRequest represents an online refund request (POST /v1/checkout/refund).
// Either OriginalTransactionID or OriginalTransactionRequestID must be provided
// to identify the original transaction to refund.
type OnlineRefundRequest struct {
	// AppID is the application ID
	AppID string `json:"appId"`

	// MerchantID is the merchant ID
	MerchantID string `json:"merchantId"`

	// TransactionRequestID is the transaction request ID for this refund transaction.
	// Unique ID to identify this refund request, used as API idempotency control field.
	// Only letters, numbers, underscores and hyphens are supported, max length 64.
	TransactionRequestID string `json:"transactionRequestId"`

	// OriginalTransactionID is the original transaction ID to refund (SUNBAY transaction ID from the payment response).
	// Either OriginalTransactionID or OriginalTransactionRequestID is required.
	// If both are provided, OriginalTransactionID takes priority.
	OriginalTransactionID string `json:"originalTransactionId,omitempty"`

	// OriginalTransactionRequestID is the original transaction request ID to refund.
	// Either OriginalTransactionID or OriginalTransactionRequestID is required.
	// If both are provided, OriginalTransactionID takes priority.
	OriginalTransactionRequestID string `json:"originalTransactionRequestId,omitempty"`

	// Amount is the refund amount information.
	// If TotalAmount is provided, system will validate it equals OrderAmount + TaxAmount + SurchargeAmount + TipAmount.
	Amount *common.OnlineRefundAmount `json:"amount,omitempty"`

	// Description is the refund description
	Description string `json:"description,omitempty"`

	// Attach is additional data, returned as-is, can be used to record refund reason or other custom information
	Attach string `json:"attach,omitempty"`

	// NotifyURL is the asynchronous notification URL (Webhook). Must be a publicly accessible HTTPS address if provided.
	NotifyURL string `json:"notifyUrl,omitempty"`
}
