package common

// AuthAmount represents authorization amount information
// Supports: orderAmount, pricingCurrency only
// Used for: Auth, ForcedAuth, IncrementalAuth
type AuthAmount struct {
	// OrderAmount is the order amount in cents (required)
	OrderAmount *int64 `json:"orderAmount"`

	// PricingCurrency is the pricing currency (ISO 4217, required)
	PricingCurrency string `json:"pricingCurrency"`
}

