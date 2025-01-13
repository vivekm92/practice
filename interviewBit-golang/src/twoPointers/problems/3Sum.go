package twoPointers

import (
	"math"
	"sort"
)

/*
problem : https://www.interviewbit.com/problems/3-sum/
*/

// T(n) : O(nlogn), S(n) : O(1)
func ThreeSumClosest(A []int, B int) int {

	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})

	n := len(A)
	res, maxAbsDiff := 0, math.MaxInt32
	for i := 0; i < n; i++ {
		diff := B - A[i]
		l, r := i+1, n-1

		for l < r {
			sum := A[l] + A[r]
			currAbsDiff := diff - sum
			if currAbsDiff < 0 {
				currAbsDiff = -1 * currAbsDiff
			}

			if maxAbsDiff > currAbsDiff {
				maxAbsDiff = currAbsDiff
				res = A[i] + A[l] + A[r]
			}
			if sum == diff {
				return B
			} else if sum > diff {
				r--
			} else {
				l++
			}
		}
	}

	return res
}
