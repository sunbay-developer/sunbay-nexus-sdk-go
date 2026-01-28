package types

// PrintReceipt represents receipt print option
type PrintReceipt string

const (
	// PrintReceiptNone NONE (code: "NONE") - Do not print receipt
	PrintReceiptNone PrintReceipt = "NONE"

	// PrintReceiptMerchant MERCHANT (code: "MERCHANT") - Print merchant copy only
	PrintReceiptMerchant PrintReceipt = "MERCHANT"

	// PrintReceiptCustomer CUSTOMER (code: "CUSTOMER") - Print customer copy only
	PrintReceiptCustomer PrintReceipt = "CUSTOMER"

	// PrintReceiptBoth BOTH (code: "BOTH") - Print both merchant and customer copies
	PrintReceiptBoth PrintReceipt = "BOTH"
)

// String returns the print receipt option code
func (p PrintReceipt) String() string {
	return string(p)
}

// IsValid checks if the print receipt option is valid
func (p PrintReceipt) IsValid() bool {
	switch p {
	case PrintReceiptNone, PrintReceiptMerchant, PrintReceiptCustomer, PrintReceiptBoth:
		return true
	default:
		return false
	}
}
