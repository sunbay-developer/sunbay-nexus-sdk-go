package response

import "github.com/sunbay-developer/sunbay-nexus-sdk-go/model/common"

// BatchCloseResponse represents a batch close transaction response
type BatchCloseResponse struct {
	common.BaseResponse

	// BatchNo is the batch number
	BatchNo string `json:"batchNo"`

	// TerminalSN is the terminal serial number
	TerminalSN string `json:"terminalSn"`

	// CloseTime is the batch close time, format: yyyy-MM-DDTHH:mm:ss+TIMEZONE (ISO 8601)
	CloseTime string `json:"closeTime"`

	// TransactionCount is the number of transactions in the batch
	TransactionCount *int `json:"transactionCount"`

	// TotalAmount is the total amount of the batch
	TotalAmount *common.BatchTotalAmount `json:"totalAmount"`
}

