package stringProblems

import (
	"testing"
)

func TestAddBinary(t *testing.T) {
	tests := []struct {
		A        string
		B        string
		Expected string
	}{
		{"100", "11", "111"},
		{"110", "10", "1000"},
	}

	for _, test := range tests {
		result := AddBinary(test.A, test.B)
		if result != test.Expected {
			t.Errorf("AddBinary(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}
}
