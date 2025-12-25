package enums

// CardNetworkType represents card network type
type CardNetworkType string

const (
	// CardNetworkTypeCredit is the credit card type
	CardNetworkTypeCredit CardNetworkType = "CREDIT"

	// CardNetworkTypeDebit is the debit card type
	CardNetworkTypeDebit CardNetworkType = "DEBIT"

	// CardNetworkTypeEBT is the EBT (Electronic Benefit Transfer) type
	CardNetworkTypeEBT CardNetworkType = "EBT"

	// CardNetworkTypeEGC is the EGC (Electronic Gift Card) type
	CardNetworkTypeEGC CardNetworkType = "EGC"

	// CardNetworkTypeUnknown is the unknown card type
	CardNetworkTypeUnknown CardNetworkType = "UNKNOWN"
)

// String returns the string representation
func (c CardNetworkType) String() string {
	return string(c)
}

// IsValid checks if the card network type is valid
func (c CardNetworkType) IsValid() bool {
	switch c {
	case CardNetworkTypeCredit, CardNetworkTypeDebit, CardNetworkTypeEBT,
		CardNetworkTypeEGC, CardNetworkTypeUnknown:
		return true
	default:
		return false
	}
}

