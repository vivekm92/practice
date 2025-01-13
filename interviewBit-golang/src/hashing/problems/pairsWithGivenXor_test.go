package hashing

import "testing"

func TestPairsWithGivenXor(t *testing.T) {

	tests := []struct {
		A        []int
		B        int
		Expected int
	}{
		{[]int{5, 4, 10, 15, 7, 6}, 5, 1},
		{[]int{3, 6, 8, 10, 15, 50}, 5, 2},
	}

	for _, test := range tests {
		result := PairsWithGivenXor(test.A, test.B)
		if result != test.Expected {
			t.Errorf("PairsWithGivenXor(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}

	for _, test := range tests {
		result := PairsWithGivenXorBruteForce(test.A, test.B)
		if result != test.Expected {
			t.Errorf("PairsWithGivenXorBruteForce(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}

	for _, test := range tests {
		result := PairsWithGivenXorOptimised(test.A, test.B)
		if result != test.Expected {
			t.Errorf("PairsWithGivenXorOptimised(%v, %v) = %v ; want %v", test.A, test.B, result, test.Expected)
		}
	}
}
