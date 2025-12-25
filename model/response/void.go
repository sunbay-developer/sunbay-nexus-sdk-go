package response

import "github.com/sunbay-developer/sunbay-nexus-sdk-go/model/common"

// VoidResponse represents a void transaction response
type VoidResponse struct {
	common.BaseResponse

	// TransactionID is the SUNBAY Nexus transaction ID for this void transaction, used for subsequent queries and notifications
	TransactionID string `json:"transactionId"`

	// TransactionRequestID is the transaction request ID for this void, returned as-is from request
	TransactionRequestID string `json:"transactionRequestId"`

	// OriginalTransactionID is the original transaction ID
	OriginalTransactionID string `json:"originalTransactionId"`

	// OriginalTransactionRequestID is the original transaction request ID
	OriginalTransactionRequestID string `json:"originalTransactionRequestId"`
}

