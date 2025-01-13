package mathProblems

import "testing"

func TestNextSmallestPalindrome(t *testing.T) {

	tests := []struct {
		A        string
		Expected string
	}{
		{"74094882455", "74094949047"},
		{"61423221", "61433416"},
		{"2222", "2332"},
		{"2773", "2882"},
		{"37223", "37273"},
		{"999", "1001"},
		{"9", "11"},
		{"99", "101"},
		{"9999", "10001"},
		{"3723", "3773"},
		{"37223", "37273"},
	}

	for _, test := range tests {
		result := NextSmallestPalindrome(test.A)
		if result != test.Expected {
			t.Errorf("NextSmallestPalindrome(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}

}
