package enums

// TransactionStatus represents transaction status
type TransactionStatus string

const (
	// TransactionStatusInitial INITIAL (code: "I")
	TransactionStatusInitial TransactionStatus = "I"

	// TransactionStatusProcessing PROCESSING (code: "P")
	// Channel called but no result obtained, or unexpected exception returned
	TransactionStatusProcessing TransactionStatus = "P"

	// TransactionStatusSuccess SUCCESS (code: "S")
	TransactionStatusSuccess TransactionStatus = "S"

	// TransactionStatusFail FAIL (code: "F")
	TransactionStatusFail TransactionStatus = "F"

	// TransactionStatusClosed CLOSED (code: "C")
	TransactionStatusClosed TransactionStatus = "C"
)

// String returns the status code
func (s TransactionStatus) String() string {
	return string(s)
}

// IsValid checks if the status is valid
func (s TransactionStatus) IsValid() bool {
	switch s {
	case TransactionStatusInitial, TransactionStatusProcessing, TransactionStatusSuccess,
		TransactionStatusFail, TransactionStatusClosed:
		return true
	default:
		return false
	}
}

