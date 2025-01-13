package mathProblems

import (
	"math"
)

/*
	Problem : https://www.interviewbit.com/problems/powerful-divisors/
*/

func isPowerOfTwo(A int) bool {
	for A%2 == 0 {
		A /= 2
	}
	return A == 1
}

// T(n) : O(n^1.5), S(n) : O(n)
func PowerfulDivisors(A []int) []int {

	m := math.MinInt32
	for _, v := range A {
		if v > m {
			m = v
		}
	}

	t := make([]int, m+1)
	for i := 1; i <= m; i++ {
		for j := 1; j*i <= m; j++ {
			t[i*j]++
		}
	}

	r := make([]int, m+1)
	for i := 1; i <= m; i++ {
		if isPowerOfTwo(t[i]) {
			r[i] = 1
		}
		r[i] += r[i-1]
	}

	for i, v := range A {
		A[i] = r[v]
	}

	return A
}
