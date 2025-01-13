package mathPrimers

import (
	"testing"
)

func TestIsPrime(t *testing.T) {
	tests := []struct {
		A        int
		Expected int
	}{
		{0, 0}, {1, 0}, {2, 1}, {3, 1}, {4, 0}, {5, 1}, {6, 0}, {7, 1}, {8, 0}, {9, 0}, {10, 0}, {11, 1},
		{12, 0}, {13, 1}, {14, 0}, {15, 0}, {16, 0}, {17, 1}, {18, 0}, {19, 1}, {20, 0}, {21, 0}, {22, 0}, {23, 1},
	}

	for _, test := range tests {
		result := IsPrime(test.A)
		if result != test.Expected {
			t.Errorf("IsPrime(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
