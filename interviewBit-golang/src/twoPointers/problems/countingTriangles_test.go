package twoPointers

import "testing"

func TestCountingTriangles(t *testing.T) {
	tests := []struct {
		A        []int
		Expected int
	}{
		{[]int{2, 1, 1, 1, 2}, 4},
	}

	for _, test := range tests {
		result := CountingTriangles(test.A)
		if result != test.Expected {
			t.Errorf("CountingTriangles(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
