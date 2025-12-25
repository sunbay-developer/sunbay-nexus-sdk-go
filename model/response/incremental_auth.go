package response

import "github.com/sunbay-developer/sunbay-nexus-sdk-go/model/common"

// IncrementalAuthResponse represents an incremental authorization transaction response
type IncrementalAuthResponse struct {
	common.BaseResponse

	// TransactionID is the SUNBAY Nexus transaction ID for this incremental authorization transaction, used for subsequent queries and notifications
	TransactionID string `json:"transactionId"`

	// TransactionRequestID is the transaction request ID for this incremental authorization, returned as-is from request
	TransactionRequestID string `json:"transactionRequestId"`

	// OriginalTransactionID is the original authorization transaction ID, returned as-is from request (only returned when provided in request)
	OriginalTransactionID string `json:"originalTransactionId,omitempty"`

	// OriginalTransactionRequestID is the original authorization transaction request ID, returned as-is from request (only returned when provided in request)
	OriginalTransactionRequestID string `json:"originalTransactionRequestId,omitempty"`
}

