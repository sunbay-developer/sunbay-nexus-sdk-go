package enums

// TransactionType represents transaction type
type TransactionType string

const (
	// TransactionTypeSale is the sale transaction type
	TransactionTypeSale TransactionType = "SALE"

	// TransactionTypeAuth is the authorization (pre-auth) transaction type
	TransactionTypeAuth TransactionType = "AUTH"

	// TransactionTypeForcedAuth is the forced authorization transaction type
	TransactionTypeForcedAuth TransactionType = "FORCED_AUTH"

	// TransactionTypeIncremental is the incremental authorization transaction type
	TransactionTypeIncremental TransactionType = "INCREMENTAL"

	// TransactionTypePostAuth is the post authorization (pre-auth completion) transaction type
	TransactionTypePostAuth TransactionType = "POST_AUTH"

	// TransactionTypeRefund is the refund transaction type
	TransactionTypeRefund TransactionType = "REFUND"

	// TransactionTypeVoid is the void transaction type
	TransactionTypeVoid TransactionType = "VOID"
)

// String returns the string representation
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

