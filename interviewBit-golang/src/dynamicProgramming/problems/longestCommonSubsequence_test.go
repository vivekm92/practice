package dpProblems

import (
	"testing"
)

func TestLCSLength(t *testing.T) {

	tests := []struct {
		input1 string
		input2 string
		output int
	}{
		{"ABCBDAB", "BDCABA", 4},
		{"ABCDEF", "ABCDEFGHI", 6},
		{"AAAAA", "AAAA", 4},
	}

	// for _, test := range tests {
	// 	result := LCSLengthNaiveRecursive(test.input1, test.input2)
	// 	if result != test.output {
	// 		t.Errorf("LCSLengthNaiveRecursive(%v, %v) = %v ; want %v", test.input1, test.input2, result, test.output)
	// 	}
	// }

	for _, test := range tests {
		result := LCSLength(test.input1, test.input2)
		if result != test.output {
			t.Errorf("LCSLength(%v, %v) = %v ; want %v", test.input1, test.input2, result, test.output)
		}
	}
}
