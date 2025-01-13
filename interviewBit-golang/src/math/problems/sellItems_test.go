package mathProblems

import (
	"testing"
)

func TestSellItems(t *testing.T) {
	tests := []struct {
		A, B, Expected int
	}{
		{3, 17, 2},
		{1, 1, 1},
	}

	for _, test := range tests {
		result := SellItems(test.A, test.B)
		if result != test.Expected {
			t.Errorf("SellItems(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}
}
