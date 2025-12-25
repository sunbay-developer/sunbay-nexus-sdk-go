package enums

// PaymentCategory represents payment category
type PaymentCategory string

const (
	// PaymentCategoryCard CARD (code: "CARD")
	PaymentCategoryCard PaymentCategory = "CARD"

	// PaymentCategoryCardCredit CARD-CREDIT (code: "CARD-CREDIT")
	PaymentCategoryCardCredit PaymentCategory = "CARD-CREDIT"

	// PaymentCategoryCardDebit CARD-DEBIT (code: "CARD-DEBIT")
	PaymentCategoryCardDebit PaymentCategory = "CARD-DEBIT"

	// PaymentCategoryQRMPM QR-MPM (code: "QR-MPM")
	PaymentCategoryQRMPM PaymentCategory = "QR-MPM"

	// PaymentCategoryQRCPM QR-CPM (code: "QR-CPM")
	PaymentCategoryQRCPM PaymentCategory = "QR-CPM"
)

// String returns the payment category code
func (p PaymentCategory) String() string {
	return string(p)
}

// IsValid checks if the payment category is valid
func (p PaymentCategory) IsValid() bool {
	switch p {
	case PaymentCategoryCard, PaymentCategoryCardCredit, PaymentCategoryCardDebit,
		PaymentCategoryQRMPM, PaymentCategoryQRCPM:
		return true
	default:
		return false
	}
}

