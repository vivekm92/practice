package mathProblems

import (
	"testing"
)

func TestProductOfDigits(t *testing.T) {
	tests := []struct {
		A        int
		Expected int
	}{
		{0, 0},
		{1, 1},
		{10, 0},
		{123, 6},
		{123347, 504},
	}

	for _, test := range tests {
		result := ProductOfDigits(test.A)
		if result != test.Expected {
			t.Errorf("ProductOfDigits(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
