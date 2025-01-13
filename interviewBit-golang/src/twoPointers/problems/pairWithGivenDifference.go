package twoPointers

import (
	"sort"
)

/*
problem : https://www.interviewbit.com/problems/pair-with-given-difference/
*/

// T(n): O(nlogn), S(n) : O(1)
func PairsWithGivenDifference(A []int, B int) int {

	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})

	n := len(A)
	i, j := 1, 0
	for i < n && j < n {
		diff := A[i] - A[j]
		if i != j && (diff == B || diff == -1*B) {
			return 1
		} else if diff < B {
			i++
		} else {
			j++
		}
	}

	return 0
}
