package bitManipulation

import "testing"

func TestBitFlipping(t *testing.T) {
	tests := []struct {
		A        int
		Expected int
	}{
		{7, 0},
		{5, 2},
	}

	for _, test := range tests {
		result := BitFlipping(test.A)
		if result != test.Expected {
			t.Errorf("BitFlipping(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
