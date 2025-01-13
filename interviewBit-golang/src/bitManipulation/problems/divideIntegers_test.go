package bitManipulation

import "testing"

func TestDivideIntegers(t *testing.T) {
	tests := []struct {
		A        int
		B        int
		Expected int
	}{
		{17, 3, 5},
		{-5, 1, -5},
		{-2147483648, -1, 2147483647},
	}

	for _, test := range tests {
		result := DivideIntegers(test.A, test.B)
		if result != test.Expected {
			t.Errorf("DivideIntegers(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}
}
