package stacksAndQueuesExamples

import (
	"testing"
)

func TestIsValidParentheses(t *testing.T) {
	tests := []struct {
		A        string
		Expected int
	}{
		{"[{()}]()[]", 1},
		{"{([", 0},
		{"))}]", 0},
	}

	for _, test := range tests {
		result := IsValidParentheses(test.A)
		if result != test.Expected {
			t.Errorf("IsValidParentheses(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
