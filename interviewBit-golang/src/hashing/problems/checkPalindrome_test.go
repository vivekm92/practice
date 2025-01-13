package hashing

import (
	"testing"
)

func TestCheckPalindrome(t *testing.T) {
	tests := []struct {
		A        string
		Expected int
	}{
		{"abbaee", 1},
		{"abcde", 0},
	}

	for _, test := range tests {
		result := CheckPalindrome(test.A)
		if result != test.Expected {
			t.Errorf("CheckPalindrome(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
