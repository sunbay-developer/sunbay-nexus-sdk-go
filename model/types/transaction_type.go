package types

// TransactionType represents transaction type
type TransactionType string

const (
	// TransactionTypeSale SALE (code: "SALE")
	TransactionTypeSale TransactionType = "SALE"

	// TransactionTypeAuth AUTH (code: "AUTH")
	TransactionTypeAuth TransactionType = "AUTH"

	// TransactionTypeForcedAuth FORCED_AUTH (code: "FORCED_AUTH")
	TransactionTypeForcedAuth TransactionType = "FORCED_AUTH"

	// TransactionTypeIncremental INCREMENTAL (code: "INCREMENTAL")
	TransactionTypeIncremental TransactionType = "INCREMENTAL"

	// TransactionTypePostAuth POST_AUTH (code: "POST_AUTH")
	TransactionTypePostAuth TransactionType = "POST_AUTH"

	// TransactionTypeRefund REFUND (code: "REFUND")
	TransactionTypeRefund TransactionType = "REFUND"

	// TransactionTypeVoid VOID (code: "VOID")
	TransactionTypeVoid TransactionType = "VOID"
)

// String returns the transaction type code
func (t TransactionType) String() string {
	return string(t)
}

// IsValid checks if the type is valid
func (t TransactionType) IsValid() bool {
	switch t {
	case TransactionTypeSale, TransactionTypeAuth, TransactionTypeForcedAuth,
		TransactionTypeIncremental, TransactionTypePostAuth, TransactionTypeRefund,
		TransactionTypeVoid:
		return true
	default:
		return false
	}
}

