package types

// TransactionBatchStatus represents the batch settlement status of a transaction.
type TransactionBatchStatus string

const (
	// TransactionBatchStatusN indicates no batch settlement needed
	TransactionBatchStatusN TransactionBatchStatus = "N"

	// TransactionBatchStatusU indicates the batch is not completed yet
	TransactionBatchStatusU TransactionBatchStatus = "U"

	// TransactionBatchStatusC indicates the batch is completed
	TransactionBatchStatusC TransactionBatchStatus = "C"
)

// String returns the transaction batch status code
func (s TransactionBatchStatus) String() string {
	return string(s)
}

// IsValid checks if the status is valid
func (s TransactionBatchStatus) IsValid() bool {
	switch s {
	case TransactionBatchStatusN, TransactionBatchStatusU, TransactionBatchStatusC:
		return true
	default:
		return false
	}
}
