package common

// OnlineRefundAmount represents online refund amount information (smallest currency unit).
type OnlineRefundAmount struct {
	// PriceCurrency is the price currency (ISO 4217)
	PriceCurrency string `json:"priceCurrency"`

	// TotalAmount is the total transaction amount (smallest currency unit)
	TotalAmount *int64 `json:"totalAmount,omitempty"`

	// OrderAmount is the order amount (smallest currency unit)
	OrderAmount *int64 `json:"orderAmount,omitempty"`

	// TaxAmount is the tax amount (smallest currency unit)
	TaxAmount *int64 `json:"taxAmount,omitempty"`

	// SurchargeAmount is the surcharge amount (smallest currency unit)
	SurchargeAmount *int64 `json:"surchargeAmount,omitempty"`

	// TipAmount is the tip amount (smallest currency unit)
	TipAmount *int64 `json:"tipAmount,omitempty"`
}
