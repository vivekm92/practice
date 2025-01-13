package arrayProblems

import (
	"sort"
)

// T(n) : O(nlog(n)), S(n) : O(n)
func Wave(A []int) []int {

	n := len(A)
	sort.Slice(A, func(i, j int) bool {
		return A[i] <= A[j]
	})

	res := make([]int, 0)
	i, j, c := 0, 1, 0
	for c <= n {
		if c%2 == 0 && j < n {
			res = append(res, A[j])
			j += 2
		} else if c%2 == 1 && i < n {
			res = append(res, A[i])
			i += 2
		}

		c++
	}

	return res
}

// T(n) : O(nlogn), S(n) : O(1)
func WaveArray(A []int) []int {

	sort.Ints(A)

	n := len(A)
	for i := 0; i < n-1; i += 2 {
		A[i], A[i+1] = A[i+1], A[i]
	}

	return A
}
