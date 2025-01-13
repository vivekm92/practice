package hashing

import (
	"testing"
)

func TestColorfulNumbers(t *testing.T) {
	tests := []struct {
		A        int
		Expected int
	}{
		{1234, 0},
		{23, 1},
		{3245, 1},
	}

	for _, test := range tests {
		result := ColorfulNumber(test.A)
		if result != test.Expected {
			t.Errorf("ColorfulNumber(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
