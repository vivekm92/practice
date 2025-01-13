package stringProblems

import "testing"

func TestStrStr(t *testing.T) {
	tests := []struct {
		A        string
		B        string
		Expected int
	}{
		{"strstr", "str", 0},
		{"bighit", "bit", -1},
	}

	for _, test := range tests {
		result := StrStr(test.A, test.B)
		if result != test.Expected {
			t.Errorf("StrStr(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}
}
