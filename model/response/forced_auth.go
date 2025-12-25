package response

import "github.com/sunbay-developer/sunbay-nexus-sdk-go/model/common"

// ForcedAuthResponse represents a forced authorization transaction response
type ForcedAuthResponse struct {
	common.BaseResponse

	// TransactionID is the SUNBAY Nexus transaction ID for this forced authorization transaction, used for subsequent queries and notifications
	TransactionID string `json:"transactionId"`

	// ReferenceOrderID is the reference order ID, returned as-is from request
	ReferenceOrderID string `json:"referenceOrderId"`

	// TransactionRequestID is the transaction request ID, returned as-is from request
	TransactionRequestID string `json:"transactionRequestId"`
}

