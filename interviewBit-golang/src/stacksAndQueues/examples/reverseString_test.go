package stacksAndQueuesExamples

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		A        string
		Expected string
	}{
		{"abcdef", "fedcba"},
	}

	for _, test := range tests {
		result := ReverseString(test.A)
		if result != test.Expected {
			t.Errorf("ReverseString(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}

	for _, test := range tests {
		result := ReverseString1(test.A)
		if result != test.Expected {
			t.Errorf("ReverseString(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}

	for _, test := range tests {
		result := ReverseString2(test.A)
		if result != test.Expected {
			t.Errorf("ReverseString(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
