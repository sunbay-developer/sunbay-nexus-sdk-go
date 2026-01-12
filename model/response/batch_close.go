package response

import "github.com/sunbay-developer/sunbay-nexus-sdk-go/model/common"

// BatchCloseResponse represents a batch close transaction response.
// The fields from the 'data' object are automatically flattened into this struct by the HTTP client.
type BatchCloseResponse struct {
	common.BaseResponse

	// BatchNo is the batch number
	BatchNo string `json:"batchNo"`

	// TerminalSN is the terminal serial number
	TerminalSN string `json:"terminalSn,omitempty"`

	// BatchTime is the batch close time, format: yyyy-MM-DDTHH:mm:ss+TIMEZONE (ISO 8601)
	BatchTime string `json:"batchTime"`

	// TransactionCount is the number of transactions in the batch
	TransactionCount int `json:"transactionCount"`

	// PriceCurrency is the transaction currency (ISO 4217)
	PriceCurrency string `json:"priceCurrency"`

	// NetAmount is the net amount in cents
	NetAmount int64 `json:"netAmount"`

	// TipAmount is the tip amount in cents
	TipAmount int64 `json:"tipAmount"`

	// SurchargeAmount is the surcharge amount in cents
	SurchargeAmount int64 `json:"surchargeAmount"`

	// TaxAmount is the tax amount in cents
	TaxAmount int64 `json:"taxAmount"`
}
