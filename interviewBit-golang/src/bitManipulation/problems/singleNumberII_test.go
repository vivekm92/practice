package bitManipulation

import "testing"

func TestSingleNumberII(t *testing.T) {
	tests := []struct {
		A        []int
		Expected int
	}{
		{[]int{0, 0, 0, 1}, 1},
		{[]int{1, 1, 1, 2}, 2},
		{[]int{1, 2, 4, 3, 3, 2, 2, 3, 1, 1}, 4},
	}

	for _, test := range tests {
		result := SingleNumberII(test.A)
		if result != test.Expected {
			t.Errorf("SingleNumberII(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
