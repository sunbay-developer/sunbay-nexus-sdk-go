package types

// CardNetworkType represents card network type
type CardNetworkType string

const (
	// CardNetworkTypeCredit CREDIT (code: "CREDIT")
	CardNetworkTypeCredit CardNetworkType = "CREDIT"

	// CardNetworkTypeDebit DEBIT (code: "DEBIT")
	CardNetworkTypeDebit CardNetworkType = "DEBIT"

	// CardNetworkTypeEBT EBT (code: "EBT")
	CardNetworkTypeEBT CardNetworkType = "EBT"

	// CardNetworkTypeEGC EGC (code: "EGC")
	CardNetworkTypeEGC CardNetworkType = "EGC"

	// CardNetworkTypeUnknown UNKNOWN (code: "UNKNOWN")
	CardNetworkTypeUnknown CardNetworkType = "UNKNOWN"
)

// String returns the card network type code
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

