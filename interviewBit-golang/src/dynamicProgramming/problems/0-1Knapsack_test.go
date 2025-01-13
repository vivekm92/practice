package dpProblems

import (
	"testing"
)

func TestKnapSack(t *testing.T) {

	tests := []struct {
		A        []int
		B        []int
		C        int
		Expected int
	}{
		{[]int{4, 5, 1}, []int{1, 2, 3}, 4, 9},
		{[]int{60, 100, 120}, []int{10, 20, 30}, 50, 220},
	}

	for _, test := range tests {
		result := KnapSackRecursive(test.A, test.B, test.C)
		if result != test.Expected {
			t.Errorf("KnapSack(%v, %v, %v) = %v ; want %v", test.A, test.B, test.C, result, test.Expected)
		}
	}

	for _, test := range tests {
		result := KnapSackRecursiveNaive(test.A, test.B, test.C)
		if result != test.Expected {
			t.Errorf("KnapSackRecursiveNaive(%v, %v, %v) = %v ; want %v", test.A, test.B, test.C, result, test.Expected)
		}
	}

	for _, test := range tests {
		result := KnapSack(test.A, test.B, test.C)
		if result != test.Expected {
			t.Errorf("KnapSack(%v, %v, %v) = %v ; want %v", test.A, test.B, test.C, result, test.Expected)
		}
	}
}
