package types

// TransactionBatchStatus represents the batch settlement status of a transaction.
type TransactionBatchStatus string

const (
	// TransactionBatchStatusNB indicates no batch settlement needed
	TransactionBatchStatusNB TransactionBatchStatus = "NB"

	// TransactionBatchStatusUB indicates waiting for batch close
	TransactionBatchStatusUB TransactionBatchStatus = "UB"

	// TransactionBatchStatusBC indicates batch closed
	TransactionBatchStatusBC TransactionBatchStatus = "BC"
)

// String returns the transaction batch status code
func (s TransactionBatchStatus) String() string {
	return string(s)
}

// IsValid checks if the status is valid
func (s TransactionBatchStatus) IsValid() bool {
	switch s {
	case TransactionBatchStatusNB, TransactionBatchStatusUB, TransactionBatchStatusBC:
		return true
	default:
		return false
	}
}
