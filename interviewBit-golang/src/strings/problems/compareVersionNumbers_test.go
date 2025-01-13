package stringProblems

import (
	"testing"
)

func TestCompareVersionNumbers(t *testing.T) {
	tests := []struct {
		A        string
		B        string
		Expected int
	}{
		{"1", "01", 0},
		{"13.0", "13.0.8", -1},
		{"444444444444444444444444", "4444444444444444444444444", -1},
	}

	for _, test := range tests {
		result := CompareVersionNumbers(test.A, test.B)
		if result != test.Expected {
			t.Errorf("CompareVersionNumbers(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}
}
