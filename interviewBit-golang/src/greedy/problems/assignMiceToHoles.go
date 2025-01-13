package greedy

import (
	"sort"
)

/*
	Problem : https://www.interviewbit.com/problems/assign-mice-to-holes/
*/

// T(n) : O(nlong) ; S(n) : O(1)
func AssignMiceToHoles(A []int, B []int) int {

	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})

	sort.Slice(B, func(i, j int) bool {
		return B[i] < B[j]
	})

	n, res := len(A), 0
	for i := 0; i < n; i++ {
		diff := A[i] - B[i]
		if diff < 0 {
			diff = -1 * diff
		}
		if res < diff {
			res = diff
		}
	}

	return res
}
