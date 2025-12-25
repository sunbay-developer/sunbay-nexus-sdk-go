package common

// BatchTotalAmount represents batch total amount information
type BatchTotalAmount struct {
	// PriceCurrency is the price currency (ISO 4217)
	PriceCurrency string `json:"priceCurrency"`

	// Amount is the total amount
	Amount *float64 `json:"amount"`
}

