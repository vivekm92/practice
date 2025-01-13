package arrayProblems

import (
	"math"
)

/*
	Problem : https://www.interviewbit.com/problems/find-duplicate-in-array/
*/

// T(n) : O(n) ; S(n) : O(n)
func RepeatedNumber(A []int) int {

	x := make(map[int]int, 0)
	for _, v := range A {
		if _, ok := x[v]; ok {
			return v
		} else {
			x[v] = 1
		}
	}

	return -1
}

// T(n) : O(n), S(n) : O(1)
func RepeatedNumber1(A []int) int {

	for _, v := range A {
		idx := int(math.Abs(float64(v)))
		if A[idx-1] < 0 {
			return idx
		}
		A[idx-1] *= -1
	}

	return -1
}
