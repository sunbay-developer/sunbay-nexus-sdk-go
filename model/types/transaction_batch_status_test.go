package types

import "testing"

func TestTransactionBatchStatusValues(t *testing.T) {
	cases := []struct {
		value TransactionBatchStatus
		want  string
	}{
		{TransactionBatchStatusN, "N"},
		{TransactionBatchStatusU, "U"},
		{TransactionBatchStatusC, "C"},
	}

	for _, tc := range cases {
		if tc.value.String() != tc.want {
			t.Fatalf("String() = %q, want %q", tc.value.String(), tc.want)
		}
		if !tc.value.IsValid() {
			t.Fatalf("%q should be valid", tc.value)
		}
	}
}
