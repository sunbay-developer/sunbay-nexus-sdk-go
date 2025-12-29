package response

import "github.com/sunbay-developer/sunbay-nexus-sdk-go/model/common"

// TipAdjustResponse represents a tip adjust transaction response
type TipAdjustResponse struct {
	common.BaseResponse

	// OriginalTransactionID is the original transaction's SUNBAY Nexus transaction ID (only returned when provided in request)
	OriginalTransactionID string `json:"originalTransactionId,omitempty"`

	// OriginalTransactionRequestID is the original transaction's request ID (only returned when provided in request)
	OriginalTransactionRequestID string `json:"originalTransactionRequestId,omitempty"`

	// TipAmount is the adjusted tip amount in cents, returned as-is from request
	TipAmount *int64 `json:"tipAmount"`
}

