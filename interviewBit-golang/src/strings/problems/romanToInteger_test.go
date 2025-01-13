package stringProblems

import "testing"

func TestRomanToInteger(t *testing.T) {
	tests := []struct {
		A        string
		Expected int
	}{
		{"XIV", 14},
		{"XVI", 16},
	}

	for _, test := range tests {
		result := RomanToInteger(test.A)
		if result != test.Expected {
			t.Errorf("RomanToInteger(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}

}
