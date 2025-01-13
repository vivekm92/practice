package stringProblems

import (
	"testing"
)

func TestCountAndSay(t *testing.T) {
	tests := []struct {
		A        int
		Expected string
	}{
		{1, "1"},
		{2, "11"},
		{3, "21"},
		{4, "1211"},
		{5, "111221"},
	}

	for _, test := range tests {
		result := CountAndSay(test.A)
		if result != test.Expected {
			t.Errorf("CountAndSay(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
