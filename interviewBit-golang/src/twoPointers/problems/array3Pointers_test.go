package twoPointers

import (
	"testing"
)

func TestMinimize(t *testing.T) {
	tests := []struct {
		A        []int
		B        []int
		C        []int
		Expected int
	}{
		{[]int{1, 4, 10}, []int{2, 15, 20}, []int{10, 12}, 5},
	}

	for _, test := range tests {
		result := Minimize(test.A, test.B, test.C)
		if result != test.Expected {
			t.Errorf("Minimize(%v, %v, %v) = %v ; want %v", test.A, test.B, test.C, result, test.Expected)
		}
	}

	for _, test := range tests {
		result := Minimize1(test.A, test.B, test.C)
		if result != test.Expected {
			t.Errorf("Minimize1(%v, %v, %v) = %v ; want %v", test.A, test.B, test.C, result, test.Expected)
		}
	}
}
