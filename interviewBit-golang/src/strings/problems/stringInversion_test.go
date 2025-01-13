package stringProblems

import (
	"testing"
)

func TestStringInversion(t *testing.T) {
	tests := []struct {
		A        string
		Expected string
	}{
		{"abc", "ABC"},
		{"ABC", "abc"},
		{"aBc", "AbC"},
	}

	for _, test := range tests {
		result := StringInversion(test.A)
		if result != test.Expected {
			t.Errorf("StringInversion(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
