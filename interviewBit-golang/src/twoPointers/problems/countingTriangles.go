package twoPointers

import (
	"sort"
)

/*
	Problem : https://www.interviewbit.com/problems/counting-triangles/
*/

// T(n) : O(n2logn) ; S(n) : O(1)
func CountingTriangles(A []int) int {

	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})

	const MOD = 1000000007
	count, n := 0, len(A)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			count = (count%MOD + findCount(A, j+1, A[i]+A[j])%MOD) % MOD
		}
	}

	return count
}

func findCount(A []int, B int, C int) int {

	l, r := B, len(A)
	for l < r {
		m := l + (r-l)/2
		if A[m] < C {
			l = m + 1
		} else {
			r = m
		}
	}

	return l - B
}
