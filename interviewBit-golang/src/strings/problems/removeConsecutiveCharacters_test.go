package stringProblems

import (
	"testing"
)

func TestRemoveConsecutiveCharacters(t *testing.T) {
	tests := []struct {
		A        string
		B        int
		Expected string
	}{
		{"aabbccd", 2, "d"},
		{"aaaabbccd", 2, "aaaad"},
	}

	for _, test := range tests {
		result := RemoveConsecutiveCharacters(test.A, test.B)
		if result != test.Expected {
			t.Errorf("RemoveConsecutiveCharacters(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}
}
