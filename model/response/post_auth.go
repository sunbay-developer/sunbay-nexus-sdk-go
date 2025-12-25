package response

import "github.com/sunbay-developer/sunbay-nexus-sdk-go/model/common"

// PostAuthResponse represents a post authorization (pre-auth completion) transaction response
type PostAuthResponse struct {
	common.BaseResponse

	// TransactionID is the SUNBAY Nexus transaction ID for this post authorization transaction, used for subsequent queries and notifications
	TransactionID string `json:"transactionId"`

	// TransactionRequestID is the transaction request ID for this post authorization, returned as-is from request
	TransactionRequestID string `json:"transactionRequestId"`

	// OriginalTransactionID is the original authorization transaction ID, returned as-is from request (only returned when provided in request)
	OriginalTransactionID string `json:"originalTransactionId,omitempty"`

	// OriginalTransactionRequestID is the original authorization transaction request ID, returned as-is from request (only returned when provided in request)
	OriginalTransactionRequestID string `json:"originalTransactionRequestId,omitempty"`
}

