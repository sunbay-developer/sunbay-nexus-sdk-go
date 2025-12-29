package common

// RefundAmount represents refund amount information
// Supports: orderAmount, tipAmount, taxAmount, surchargeAmount, cashbackAmount
type RefundAmount struct {
	// OrderAmount is the order amount in cents (required)
	OrderAmount *int64 `json:"orderAmount"`

	// TipAmount is the tip amount in cents (optional, must be greater than or equal to 0)
	TipAmount *int64 `json:"tipAmount,omitempty"`

	// TaxAmount is the tax amount in cents (optional, must be greater than or equal to 0)
	TaxAmount *int64 `json:"taxAmount,omitempty"`

	// SurchargeAmount is the surcharge amount in cents (optional, must be greater than or equal to 0)
	// Note: Some processors may require surcharge to be refunded proportionally. Please contact technical support for detailed policies.
	SurchargeAmount *int64 `json:"surchargeAmount,omitempty"`

	// CashbackAmount is the cashback amount in cents (optional, must be greater than or equal to 0)
	CashbackAmount *int64 `json:"cashbackAmount,omitempty"`

	// PricingCurrency is the pricing currency (ISO 4217, required)
	PricingCurrency string `json:"pricingCurrency"`
}

