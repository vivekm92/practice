package bitManipulation

import "testing"

func TestCountTotalSetBits(t *testing.T) {
	tests := []struct {
		A        int
		Expected int
	}{
		{5, 7},
		{11, 20},
	}

	for _, test := range tests {
		result := CountTotalSetBits(test.A)
		if result != test.Expected {
			t.Errorf("CountTotalSetBits(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
