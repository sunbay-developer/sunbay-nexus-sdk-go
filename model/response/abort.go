package response

import "github.com/sunbay-developer/sunbay-nexus-sdk-go/model/common"

// AbortResponse represents an abort transaction response
type AbortResponse struct {
	common.BaseResponse

	// OriginalTransactionID is the aborted transaction's SUNBAY Nexus transaction ID
	OriginalTransactionID string `json:"originalTransactionId"`

	// OriginalTransactionRequestID is the aborted transaction's request ID (only returned when provided in request)
	OriginalTransactionRequestID string `json:"originalTransactionRequestId,omitempty"`
}

