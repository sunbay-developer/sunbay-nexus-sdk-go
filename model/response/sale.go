package response

import "github.com/sunbay-developer/sunbay-nexus-sdk-go/model/common"

// SaleResponse represents a sale transaction response
type SaleResponse struct {
	common.BaseResponse

	// TransactionID is the SUNBAY Nexus transaction ID for this sale transaction, used for subsequent queries and notifications
	TransactionID string `json:"transactionId"`

	// ReferenceOrderID is the reference order ID, returned as-is from request
	ReferenceOrderID string `json:"referenceOrderId"`

	// TransactionRequestID is the transaction request ID, returned as-is from request
	TransactionRequestID string `json:"transactionRequestId"`
}

