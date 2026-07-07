package types

// RelatedTransactionStatus indicates the lifecycle change of the current transaction
// due to subsequent transactions (e.g., a sale transaction that was later refunded).
type RelatedTransactionStatus string

const (
	// RelatedTransactionStatusVoided indicates the transaction has been voided
	RelatedTransactionStatusVoided RelatedTransactionStatus = "VOIDED"

	// RelatedTransactionStatusIncremental indicates the transaction has incremental authorization
	RelatedTransactionStatusIncremental RelatedTransactionStatus = "INCREMENTAL"

	// RelatedTransactionStatusRefunded indicates the transaction has been fully refunded
	RelatedTransactionStatusRefunded RelatedTransactionStatus = "REFUNDED"

	// RelatedTransactionStatusCapture indicates the transaction has been captured (post-auth)
	RelatedTransactionStatusCapture RelatedTransactionStatus = "CAPTURE"

	// RelatedTransactionStatusPartRefunded indicates the transaction has been partially refunded
	RelatedTransactionStatusPartRefunded RelatedTransactionStatus = "PART_REFUNDED"
)

// String returns the related transaction status code
func (s RelatedTransactionStatus) String() string {
	return string(s)
}

// IsValid checks if the status is valid
func (s RelatedTransactionStatus) IsValid() bool {
	switch s {
	case RelatedTransactionStatusVoided, RelatedTransactionStatusIncremental,
		RelatedTransactionStatusRefunded, RelatedTransactionStatusCapture,
		RelatedTransactionStatusPartRefunded:
		return true
	default:
		return false
	}
}
