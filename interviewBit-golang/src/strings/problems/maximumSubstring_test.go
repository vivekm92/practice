package stringProblems

import (
	"testing"
)

func TestMaximumSubstring(t *testing.T) {
	tests := []struct {
		A        string
		B        int
		Expected int
	}{
		{"baab", 2, 1},
		{"baa", 2, 1},
	}

	for _, test := range tests {
		result := MaximumSubstring(test.A, test.B)
		if result != test.Expected {
			t.Errorf("MaximumSubstring(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}
}
