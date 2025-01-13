package twoPointers

import "testing"

func TestNumRange(t *testing.T) {
	tests := []struct {
		A        []int
		B        int
		C        int
		Expected int
	}{
		{[]int{10, 5, 1, 0, 2}, 6, 8, 3},
	}

	for _, test := range tests {
		result := NumRange(test.A, test.B, test.C)
		if result != test.Expected {
			t.Errorf("NumRange(%v, %v, %v) = %v ; want %v", test.A, test.B, test.C, result, test.Expected)
		}
	}
}
