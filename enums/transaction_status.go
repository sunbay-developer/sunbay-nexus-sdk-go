package enums

// TransactionStatus represents transaction status
type TransactionStatus string

const (
	// TransactionStatusInitial is the initial state (code: "I")
	TransactionStatusInitial TransactionStatus = "I"

	// TransactionStatusProcessing is the transaction processing state (code: "P")
	// Channel called but no result obtained, or unexpected exception returned
	TransactionStatusProcessing TransactionStatus = "P"

	// TransactionStatusSuccess is the transaction successful state (code: "S")
	TransactionStatusSuccess TransactionStatus = "S"

	// TransactionStatusFail is the transaction failed state (code: "F")
	TransactionStatusFail TransactionStatus = "F"

	// TransactionStatusClosed is the transaction closed state (code: "C")
	TransactionStatusClosed TransactionStatus = "C"
)

// Code returns the status code (same as the enum value)
func (s TransactionStatus) Code() string {
	return string(s)
}

// Desc returns the status description
func (s TransactionStatus) Desc() string {
	switch s {
	case TransactionStatusInitial:
		return "INITIAL"
	case TransactionStatusProcessing:
		return "PROCESSING"
	case TransactionStatusSuccess:
		return "SUCCESS"
	case TransactionStatusFail:
		return "FAIL"
	case TransactionStatusClosed:
		return "CLOSED"
	default:
		return ""
	}
}

// String returns the string representation (status description)
func (s TransactionStatus) String() string {
	return s.Desc()
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

