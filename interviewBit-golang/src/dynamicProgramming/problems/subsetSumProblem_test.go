package dpProblems

import (
	"testing"
)

func TestSubsetSumProblem(t *testing.T) {
	tests := []struct {
		A        []int
		B        int
		Expected int
	}{
		{[]int{4, 1, 3, 12, 2}, 9, 1},
		{[]int{4, 1, 3, 12, 2}, 11, 0},
	}

	for _, test := range tests {
		result := SubsetSumRecursiveNaive(test.A, test.B)
		if result != test.Expected {
			t.Errorf("SubsetSumRecursiveNaive(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}

	for _, test := range tests {
		result := SubSetSumRecusive(test.A, test.B)
		if result != test.Expected {
			t.Errorf("SubSetSumRecusive(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}

	for _, test := range tests {
		result := SubSetSum(test.A, test.B)
		if result != test.Expected {
			t.Errorf("SubSetSum(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}

	for _, test := range tests {
		result := SubSetSumOptimised(test.A, test.B)
		if result != test.Expected {
			t.Errorf("SubSetSumOptimised(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}
}
