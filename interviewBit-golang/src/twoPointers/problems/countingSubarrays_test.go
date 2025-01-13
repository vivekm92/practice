package twoPointers

import "testing"

func TestCountSubArrays(t *testing.T) {
	tests := []struct {
		input1 []int
		input2 int
		output int
	}{
		{[]int{2, 5, 6}, 10, 4},
		{[]int{1, 11, 2, 3, 15}, 10, 4},
		{[]int{8, 5, 1, 10, 5, 9, 9, 3, 5, 6, 6, 2, 8, 2, 2, 6, 3, 8, 7, 2, 5, 3, 4, 3, 3, 2, 7, 9, 6, 8, 7, 2, 9, 10, 3, 8, 10, 6, 5, 4, 2, 3, 4, 4, 5, 2, 2, 4, 9, 8, 5}, 32, 280},
	}

	for _, test := range tests {
		result := CountSubArrays(test.input1, test.input2)
		if result != test.output {
			t.Errorf("CountSubArrays(%v, %v) = %v ; want %v", test.input1, test.input2, result, test.output)
		}
	}
}
