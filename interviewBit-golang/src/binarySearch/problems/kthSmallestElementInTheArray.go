package bsearch

import "math"

/*
	Problem : https://www.interviewbit.com/problems/kth-smallest-element-in-the-array/
*/

func countElementsLessThanMid(A []int, t int) int {

	count := 0
	for _, v := range A {
		if v <= t {
			count++
		}
	}
	return count
}

// T(n) : O(klogn), S(n) : O(n)
func Kthsmallest(A []int, B int) int {

	min, max := math.MaxInt32, math.MinInt32
	for _, v := range A {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	l, r := min, max
	for l < r {
		mid := l + (r-l)/2
		count := countElementsLessThanMid(A, mid)
		if count < B {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}
