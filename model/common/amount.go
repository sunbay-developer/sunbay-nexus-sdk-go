package common

import (
	"encoding/json"
	"strconv"
)

// Amount represents transaction amount information
// Used in query response
type Amount struct {
	// PriceCurrency is the price currency (ISO 4217)
	PriceCurrency string `json:"priceCurrency,omitempty"`

	// TransAmount is the transaction amount in cents (calculated field in response)
	TransAmount *int64 `json:"transAmount,omitempty"`

	// OrderAmount is the order amount in cents
	OrderAmount *int64 `json:"orderAmount,omitempty"`

	// TaxAmount is the tax amount in cents
	TaxAmount *int64 `json:"taxAmount,omitempty"`

	// SurchargeAmount is the surcharge amount in cents
	SurchargeAmount *int64 `json:"surchargeAmount,omitempty"`

	// TipAmount is the tip amount in cents
	TipAmount *int64 `json:"tipAmount,omitempty"`

	// CashbackAmount is the cashback amount in cents
	CashbackAmount *int64 `json:"cashbackAmount,omitempty"`
}

// UnmarshalJSON implements custom JSON unmarshaling to handle number/string-to-int conversion
// API may return amounts as numbers or strings (in cents), we convert them to int64
func (a *Amount) UnmarshalJSON(data []byte) error {
	// Define a temporary struct with interface{} fields to handle both number and string
	type Alias struct {
		PriceCurrency   string      `json:"priceCurrency,omitempty"`
		TransAmount     interface{} `json:"transAmount,omitempty"`
		OrderAmount     interface{} `json:"orderAmount,omitempty"`
		TaxAmount       interface{} `json:"taxAmount,omitempty"`
		SurchargeAmount interface{} `json:"surchargeAmount,omitempty"`
		TipAmount       interface{} `json:"tipAmount,omitempty"`
		CashbackAmount  interface{} `json:"cashbackAmount,omitempty"`
	}

	var alias Alias
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}

	// Copy non-amount fields
	a.PriceCurrency = alias.PriceCurrency

	// Convert amount values to int64 (cents)
	// API may return as number (int/float) or string (integer or float), handle all cases
	// If it's a float (e.g., 222.00 or "222.00"), convert to cents by multiplying by 100
	convertAmount := func(v interface{}) *int64 {
		if v == nil {
			return nil
		}

		switch val := v.(type) {
		case float64:
			// Number type: if it's already an integer (e.g., 10000), use it directly
			// If it's a float (e.g., 222.00), convert to cents
			if val == float64(int64(val)) {
				// Already an integer (in cents)
				result := int64(val)
				return &result
			} else {
				// Float (in currency units), convert to cents
				result := int64(val * 100)
				return &result
			}
		case int64:
			// Already int64 (in cents)
			return &val
		case int:
			// int (in cents)
			result := int64(val)
			return &result
		case string:
			// String type: try parsing as integer first (already in cents)
			if intVal, err := strconv.ParseInt(val, 10, 64); err == nil {
				return &intVal
			}
			// Try parsing as float (in currency units), then convert to cents
			if floatVal, err := strconv.ParseFloat(val, 64); err == nil {
				result := int64(floatVal * 100)
				return &result
			}
		}

		return nil
	}

	a.TransAmount = convertAmount(alias.TransAmount)
	a.OrderAmount = convertAmount(alias.OrderAmount)
	a.TaxAmount = convertAmount(alias.TaxAmount)
	a.SurchargeAmount = convertAmount(alias.SurchargeAmount)
	a.TipAmount = convertAmount(alias.TipAmount)
	a.CashbackAmount = convertAmount(alias.CashbackAmount)

	return nil
}
