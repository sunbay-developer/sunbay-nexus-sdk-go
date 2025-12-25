package common

import (
	"encoding/json"
	"strconv"
)

// Amount represents transaction amount information
// Used in query response
type Amount struct {
	// PriceCurrency is the price currency (used in query response, ISO 4217)
	PriceCurrency string `json:"priceCurrency,omitempty"`

	// TransAmount is the transaction amount (calculated field in response)
	TransAmount *float64 `json:"transAmount,omitempty"`

	// OrderAmount is the order amount
	OrderAmount *float64 `json:"orderAmount,omitempty"`

	// TaxAmount is the tax amount
	TaxAmount *float64 `json:"taxAmount,omitempty"`

	// SurchargeAmount is the surcharge amount
	SurchargeAmount *float64 `json:"surchargeAmount,omitempty"`

	// TipAmount is the tip amount
	TipAmount *float64 `json:"tipAmount,omitempty"`

	// CashbackAmount is the cashback amount
	CashbackAmount *float64 `json:"cashbackAmount,omitempty"`

	// PricingCurrency is the pricing currency (ISO 4217, used in request)
	PricingCurrency string `json:"pricingCurrency,omitempty"`
}

// UnmarshalJSON implements custom JSON unmarshaling to handle string-to-float conversion
func (a *Amount) UnmarshalJSON(data []byte) error {
	// Define a temporary struct with string fields for amount values
	type Alias struct {
		PriceCurrency   string  `json:"priceCurrency,omitempty"`
		TransAmount     *string `json:"transAmount,omitempty"`
		OrderAmount     *string `json:"orderAmount,omitempty"`
		TaxAmount       *string `json:"taxAmount,omitempty"`
		SurchargeAmount *string `json:"surchargeAmount,omitempty"`
		TipAmount       *string `json:"tipAmount,omitempty"`
		CashbackAmount  *string `json:"cashbackAmount,omitempty"`
		PricingCurrency string  `json:"pricingCurrency,omitempty"`
	}

	var alias Alias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}

	// Copy non-amount fields
	a.PriceCurrency = alias.PriceCurrency
	a.PricingCurrency = alias.PricingCurrency

	// Convert string amounts to float64
	if alias.TransAmount != nil {
		if val, err := strconv.ParseFloat(*alias.TransAmount, 64); err == nil {
			a.TransAmount = &val
		}
	}
	if alias.OrderAmount != nil {
		if val, err := strconv.ParseFloat(*alias.OrderAmount, 64); err == nil {
			a.OrderAmount = &val
		}
	}
	if alias.TaxAmount != nil {
		if val, err := strconv.ParseFloat(*alias.TaxAmount, 64); err == nil {
			a.TaxAmount = &val
		}
	}
	if alias.SurchargeAmount != nil {
		if val, err := strconv.ParseFloat(*alias.SurchargeAmount, 64); err == nil {
			a.SurchargeAmount = &val
		}
	}
	if alias.TipAmount != nil {
		if val, err := strconv.ParseFloat(*alias.TipAmount, 64); err == nil {
			a.TipAmount = &val
		}
	}
	if alias.CashbackAmount != nil {
		if val, err := strconv.ParseFloat(*alias.CashbackAmount, 64); err == nil {
			a.CashbackAmount = &val
		}
	}

	return nil
}

