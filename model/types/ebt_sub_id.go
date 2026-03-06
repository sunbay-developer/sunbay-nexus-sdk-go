package types

// EBTSubID represents EBT sub payment method. Only applicable when payment method category is EBT and id is EBT
type EBTSubID string

const (
	// EBTSubIDSnap SNAP
	EBTSubIDSnap EBTSubID = "SNAP"
	// EBTSubIDVoucher VOUCHER
	EBTSubIDVoucher EBTSubID = "VOUCHER"
	// EBTSubIDBenefit BENEFIT
	EBTSubIDBenefit EBTSubID = "BENEFIT"
)

// String returns the EBT sub payment method code
func (e EBTSubID) String() string {
	return string(e)
}

// IsValid checks if the EBT sub ID is valid
func (e EBTSubID) IsValid() bool {
	switch e {
	case EBTSubIDSnap, EBTSubIDVoucher, EBTSubIDBenefit:
		return true
	default:
		return false
	}
}
