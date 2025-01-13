package bitManipulation

import (
	"testing"
)

func TestSwapBits(t *testing.T) {
	tests := []struct {
		A        int
		B        int
		C        int
		Expected int
	}{
		{5, 1, 2, 6},
		{9, 1, 2, 10},
	}

	for _, test := range tests {
		result := SwapBits(test.A, test.B, test.C)
		if result != test.Expected {
			t.Errorf("SwapBits(%v, %v, %v) = %v ; want %v", test.A, test.B, test.C, result, test.Expected)
		}
	}
}
