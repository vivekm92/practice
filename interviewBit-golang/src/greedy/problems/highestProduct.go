package greedy

// Problem : https://www.interviewbit.com/problems/highest-product/

import (
	"sort"
)

// T(n) : O(nlogn) , S(n) : O(1)
func HighestProduct(A []int) int {

	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})

	n := len(A)
	v1 := A[0] * A[1] * A[n-1]
	v2 := A[n-1] * A[n-2] * A[n-3]
	if v1 < v2 {
		return v2
	}

	return v1
}
