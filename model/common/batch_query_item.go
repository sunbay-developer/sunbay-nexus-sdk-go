package common

// BatchQueryItem represents batch query item information
// Statistics grouped by channel code and transaction currency
type BatchQueryItem struct {
	// BatchNo is the batch number
	BatchNo string `json:"batchNo"`

	// StartTime is the batch start time, format: yyyy-MM-DDTHH:mm:ss+TIMEZONE (ISO 8601)
	StartTime string `json:"startTime"`

	// ChannelCode is the payment channel code
	ChannelCode string `json:"channelCode"`

	// PriceCurrency is the transaction currency (ISO 4217)
	// Note: Java uses transactionCurrency, Go uses priceCurrency for consistency
	PriceCurrency string `json:"priceCurrency"`

	// TotalCount is the total count of transactions in the batch
	TotalCount int `json:"totalCount"`

	// NetAmount is the net amount in cents
	NetAmount int64 `json:"netAmount"`

	// TipAmount is the tip amount in cents
	TipAmount int64 `json:"tipAmount"`

	// SurchargeAmount is the surcharge amount in cents
	SurchargeAmount int64 `json:"surchargeAmount"`

	// TaxAmount is the tax amount in cents
	TaxAmount int64 `json:"taxAmount"`
}
