package stringProblems

import "testing"

func TestFlippingString(t *testing.T) {
	tests := []struct {
		A        string
		Expected int
	}{
		{"000", 3},
		{"10010", 4},
	}

	for _, test := range tests {
		result := FlippingString(test.A)
		if result != test.Expected {
			t.Errorf("FlippingString(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
