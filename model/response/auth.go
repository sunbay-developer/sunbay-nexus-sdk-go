package response

import "github.com/sunbay-developer/sunbay-nexus-sdk-go/model/common"

// AuthResponse represents an authorization (pre-auth) transaction response
type AuthResponse struct {
	common.BaseResponse

	// TransactionID is the SUNBAY Nexus transaction ID for this authorization transaction, used for subsequent queries, notifications, and post authorization/incremental authorization operations
	TransactionID string `json:"transactionId"`

	// ReferenceOrderID is the reference order ID, returned as-is from request
	ReferenceOrderID string `json:"referenceOrderId"`

	// TransactionRequestID is the transaction request ID, returned as-is from request
	TransactionRequestID string `json:"transactionRequestId"`
}

