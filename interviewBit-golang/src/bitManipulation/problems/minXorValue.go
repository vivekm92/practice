package bitManipulation

import (
	"math"
	"sort"
)

/*
	problem : https://www.interviewbit.com/problems/min-xor-value/
*/

// T(n) : O(nlong), S(n) : O(1)
func FindMinXor(A []int) int {
	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})

	minXor, n := math.MaxInt32, len(A)
	for i := 0; i < n-1; i++ {
		xorVal := A[i] ^ A[i+1]
		if minXor > xorVal {
			minXor = xorVal
		}
	}

	return minXor
}
