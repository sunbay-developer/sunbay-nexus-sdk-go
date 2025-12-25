package response

import "github.com/sunbay-developer/sunbay-nexus-sdk-go/model/common"

// RefundResponse represents a refund transaction response
type RefundResponse struct {
	common.BaseResponse

	// TransactionID is the SUNBAY Nexus transaction ID for this refund transaction, used for subsequent queries and notifications
	TransactionID string `json:"transactionId"`

	// ReferenceOrderID is the reference order ID (same as original transaction for refund with reference, new refund reference order ID for refund without reference)
	ReferenceOrderID string `json:"referenceOrderId"`

	// TransactionRequestID is the transaction request ID for this refund, returned as-is from request
	TransactionRequestID string `json:"transactionRequestId"`

	// OriginalTransactionID is the original transaction ID (only returned for refund with reference)
	OriginalTransactionID string `json:"originalTransactionId,omitempty"`

	// OriginalTransactionRequestID is the original transaction request ID (only returned for refund with reference)
	OriginalTransactionRequestID string `json:"originalTransactionRequestId,omitempty"`
}

