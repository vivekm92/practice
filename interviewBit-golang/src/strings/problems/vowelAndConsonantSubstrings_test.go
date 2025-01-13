package stringProblems

import "testing"

func TestVowelAndConsonantSubstrings(t *testing.T) {
	tests := []struct {
		A        string
		Expected int
	}{
		{"aba", 2},
		{"a", 0},
	}

	for _, test := range tests {
		result := VowelAndConsonantSubstrings(test.A)
		if result != test.Expected {
			t.Errorf("VowelAndConsonantSubstrings(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
