package timeComplexity

import (
	"testing"
)

func TestPalindromicTime(t *testing.T) {
	tests := []struct {
		A        string
		Expected int
	}{
		{"23:59", 1},
		{"21:00", 12},
		{"03:00", 30},
	}

	for _, test := range tests {
		result := PalindromicTime(test.A)
		if result != test.Expected {
			t.Errorf("PalindromicTime(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
