package twoPointers

import "testing"

func TestMaximumOnesAfterModification(t *testing.T) {
	tests := []struct {
		input1 []int
		input2 int
		output int
	}{
		{[]int{1, 0, 0, 1, 1, 0, 1}, 1, 4},
		{[]int{1, 0, 0, 1, 0, 1, 0, 1, 0, 1}, 2, 5},
	}

	for _, test := range tests {
		result := MaximumOnesAfterModification(test.input1, test.input2)
		if result != test.output {
			t.Errorf("MaximumOnesAfterModification(%v, %v) = %v ; want %v", test.input1, test.input2, result, test.output)
		}
	}
}
