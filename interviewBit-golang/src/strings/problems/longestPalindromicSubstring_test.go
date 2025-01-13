package stringProblems

import "testing"

func TestLongestPalindromicSubstring(t *testing.T) {
	tests := []struct {
		A        string
		Expected string
	}{
		{"aaabaaa", "aaabaaa"},
		{"vivek", "viv"},
	}

	for _, test := range tests {
		result := LongestPalindromicSubstring(test.A)
		if result != test.Expected {
			t.Errorf("LongestPalindromicSubstring(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}

}
