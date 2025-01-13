package stringProblems

import "testing"

func TestConvertToPalindrome(t *testing.T) {
	tests := []struct {
		A        string
		Expected int
	}{
		{"aabbcbbaa", 1},
		{"ivwevi", 1},
		{"ivwerevi", 1},
		{"ivwejhalrevi", 0},
	}

	for _, test := range tests {
		result := ConvertToPalindrome(test.A)
		if result != test.Expected {
			t.Errorf("ConvertToPalindrome(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}

}
