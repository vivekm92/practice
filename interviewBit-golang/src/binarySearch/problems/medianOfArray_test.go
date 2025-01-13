package bsearch

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {

	tests := []struct {
		input1 []int
		input2 []int
		output float64
	}{
		{[]int{}, []int{1}, 1.0},
		{[]int{1, 4, 5}, []int{2, 3}, 3.0},
	}

	for _, test := range tests {
		result := FindMedianSortedArrays(test.input1, test.input2)
		if result != test.output {
			t.Errorf("FindMedianSortedArrays(%v, %v) = %v ; want %v", test.input1, test.input2, result, test.output)
		}
	}
}
