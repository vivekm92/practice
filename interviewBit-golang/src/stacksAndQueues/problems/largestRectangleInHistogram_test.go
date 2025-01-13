package stacksAndQueues

import (
	"testing"
)

func TestLargestRectangleInHistogram(t *testing.T) {

	tests := []struct {
		A        []int
		Expected int
	}{
		{[]int{2, 1, 5, 6, 2, 3}, 10},
		{[]int{90, 58, 69, 70, 82, 100, 13, 57, 47, 18}, 348},
		{[]int{6, 2, 5, 4, 5, 1, 6}, 12},
	}

	for _, test := range tests {
		result := LargestRectangleInHistogram(test.A)
		if result != test.Expected {
			t.Errorf("LargestRectangleInHistogram(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}

	for _, test := range tests {
		result := LargestRectangleInHistogram1(test.A)
		if result != test.Expected {
			t.Errorf("LargestRectangleInHistogram1(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
