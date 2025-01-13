package twoPointers

import (
	"interviewBit/src/utils"
	"math"
)

/*
	Problem : https://www.interviewbit.com/problems/array-3-pointers/
*/

// T(n) : O(n3) , S(n) : O(1)
func Minimize(A []int, B []int, C []int) int {

	res := math.MaxInt32
	n, m, o := len(A), len(B), len(C)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < o; k++ {
				v1 := absDiff(A[i], B[j])
				v2 := absDiff(B[j], C[k])
				v3 := absDiff(C[k], A[i])

				currMax := utils.MaxOfIntsOrFloats[int](v1, v2, v3)
				if res > currMax {
					res = currMax
				}
			}
		}
	}

	return res
}

// T(n) : O(nlogn) , S(n) : O(1)
func Minimize1(A []int, B []int, C []int) int {

	v1 := solve(A, B, C)
	v2 := solve(B, A, C)
	v3 := solve(C, A, B)

	return utils.MinOfIntsOrFloats(v1, v2, v3)
}

func solve(A []int, B []int, C []int) int {

	res := math.MaxInt32
	for _, v := range A {
		v1 := bsearch(B, v)
		v2 := bsearch(C, v)
		k1 := absDiff(v1, v2)
		k2 := absDiff(v, v1)
		k3 := absDiff(v, v2)

		currMax := utils.MaxOfIntsOrFloats(k1, k2, k3)
		if res > currMax {
			res = currMax
		}
	}

	return res
}

func bsearch(A []int, B int) int {
	n := len(A)
	l, r := 0, n
	for l < r {
		mid := l + (r-l)/2
		if A[mid] == B {
			return A[mid]
		} else if A[mid] > B {
			r = mid
		} else {
			l = mid + 1
		}
	}

	if l == n {
		return A[l-1]
	}

	return A[l]
}

func absDiff(A, B int) int {
	if A > B {
		return A - B
	}
	return B - A
}
