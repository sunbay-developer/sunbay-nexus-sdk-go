package common

// CheckoutProductLine is a line item for hosted checkout or direct checkout requests.
// If productList is sent, the sum of amount × num must equal amount.orderAmount.
type CheckoutProductLine struct {
	// Amount is the line amount in minor units (e.g. cents)
	Amount int64 `json:"amount"`
	// Name is the product name
	Name string `json:"name"`
	// Num is the quantity
	Num int `json:"num"`
}

// CheckoutAddress is billing or shipping address for direct checkout.
type CheckoutAddress struct {
	Country    string `json:"country,omitempty"`
	State      string `json:"state,omitempty"`
	City       string `json:"city,omitempty"`
	Line1      string `json:"line1,omitempty"`
	Line2      string `json:"line2,omitempty"`
	PostalCode string `json:"postalCode,omitempty"`
}
