package timeComplexity

import (
	"testing"
)

func TestCollatz(t *testing.T) {
	tests := []struct {
		A        int
		B        int
		Expected int64
	}{
		{1, 3, 2},
		{5, 6, 1},
	}

	for _, test := range tests {
		result := Collatz(test.A, test.B)
		if result != test.Expected {
			t.Errorf("Collatz(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}

}
