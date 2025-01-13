package bsearch

import "math"

/*
	Problem : https://www.interviewbit.com/problems/matrix-median
*/

func lowerBound(A []int, t int) int {
	l, r := 0, len(A)
	for l < r {
		mid := l + (r-l)/2
		if A[mid] <= t {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

// T(n) : O(logN), S(n) : O(1)
func MatrixMedian(A [][]int) int {

	n, m := len(A), len(A[0])
	l, r := math.MaxInt32, math.MinInt32
	for _, arr := range A {
		if arr[0] < l {
			l = arr[0]
		}
		if arr[m-1] > r {
			r = arr[m-1]
		}
	}

	target := (n * m) / 2
	for l <= r {
		mid := l + (r-l)/2
		count := 0
		for _, arr := range A {
			count += lowerBound(arr, mid)
		}

		if count <= target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return l
}
