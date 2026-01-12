package common

// SaleAmount represents sale transaction amount information
// Supports: orderAmount, tipAmount, taxAmount, surchargeAmount, cashbackAmount
type SaleAmount struct {
	// OrderAmount is the order amount in cents (required)
	OrderAmount *int64 `json:"orderAmount"`

	// TipAmount is the tip amount in cents (optional)
	TipAmount *int64 `json:"tipAmount,omitempty"`

	// TaxAmount is the tax amount in cents (optional)
	TaxAmount *int64 `json:"taxAmount,omitempty"`

	// SurchargeAmount is the surcharge amount in cents (optional)
	SurchargeAmount *int64 `json:"surchargeAmount,omitempty"`

	// CashbackAmount is the cashback amount in cents (optional)
	CashbackAmount *int64 `json:"cashbackAmount,omitempty"`

	// PriceCurrency is the price currency (ISO 4217, required)
	PriceCurrency string `json:"priceCurrency"`
}
