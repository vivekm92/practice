package mathProblems

import (
	"sort"
)

/*
	Problem : https://www.interviewbit.com/problems/gp-triplets/
*/

// T(n) : O(n2) ; S(n) : O(n)
func GpTriplets1(A []int) int {

	lookup := make(map[int]int)
	for _, v := range A {
		lookup[v]++
	}

	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})

	n, count := len(A), 0
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if d := A[i] / A[j]; d > 1 && A[i]%A[j] == 0 {
				if v, found := lookup[A[i]*d]; found {
					count += v
				}
			}
		}
	}

	for _, v := range lookup {
		if v > 2 {
			count += binomial(v)
		}
	}

	return count
}
