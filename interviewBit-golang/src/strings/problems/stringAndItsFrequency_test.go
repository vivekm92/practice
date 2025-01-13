package stringProblems

import (
	"testing"
)

func TestStringAndItsFrequency(t *testing.T) {

	tests := []struct {
		A        string
		Expected string
	}{
		{"abbhuabcfghh", "a2b3h3u1c1f1g1"},
		{"a", "a1"},
	}

	for _, test := range tests {
		result := StringAndItsFrequency(test.A)
		if result != test.Expected {
			t.Errorf("StringAndItsFrequency(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
