package response

import "github.com/sunbay-developer/sunbay-nexus-sdk-go/model/common"

// OnlineRefundResponse represents an online refund response (POST /v1/checkout/refund).
type OnlineRefundResponse struct {
	common.BaseResponse

	// TransactionID is the SUNBAY Nexus transaction ID for this refund transaction
	TransactionID string `json:"transactionId"`

	// TransactionRequestID is the transaction request ID, returned as-is from request
	TransactionRequestID string `json:"transactionRequestId"`

	// OriginalTransactionID is the original transaction ID
	OriginalTransactionID string `json:"originalTransactionId,omitempty"`

	// TransactionStatus is the transaction status: INITIAL/PROCESSING/SUCCESS/FAIL/CLOSED
	TransactionStatus string `json:"transactionStatus"`

	// TransactionType is the transaction type, fixed as REFUND
	TransactionType string `json:"transactionType,omitempty"`

	// Amount is the refund amount information (smallest currency unit)
	Amount *common.OnlineRefundAmount `json:"amount,omitempty"`

	// CreateTime is the refund creation time, ISO 8601 format
	CreateTime string `json:"createTime,omitempty"`

	// CompleteTime is the refund completion time, returned when transaction reaches terminal state. ISO 8601 format
	CompleteTime string `json:"completeTime,omitempty"`

	// TransactionResultCode is the transaction result code
	TransactionResultCode string `json:"transactionResultCode,omitempty"`

	// TransactionResultMsg is the transaction result message
	TransactionResultMsg string `json:"transactionResultMsg,omitempty"`

	// Description is the refund description (returned as-is from request)
	Description string `json:"description,omitempty"`
}
