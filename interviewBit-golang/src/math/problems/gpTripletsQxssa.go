package mathProblems

import (
	"sort"
)

/*
	Problem : https://www.interviewbit.com/problems/gp-triplets-qxssa/
*/

// T(n) : O(n2longn) , S(n) : O(n)
func GpTriplets(A []int) int {

	lookup := make(map[int]int)
	for _, v := range A {
		lookup[v]++
	}

	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})

	count, n := 0, len(A)
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

func binomial(A int) int {
	r := 1
	for i := A; i > A-3; i-- {
		r = r * i
	}
	return r
}
