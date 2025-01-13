package bsearch

import (
	"interviewBit/src/utils"
	"math"
)

/*
	Problem : https://www.interviewbit.com/problems/median-of-array/
*/

// T(n) : O(logN), S(n) : O(1)
func FindMedianSortedArrays(A []int, B []int) float64 {

	n, m := len(A), len(B)
	if n > m {
		return FindMedianSortedArrays(B, A)
	}

	l, r := 0, n
	mF := (n + m + 1) / 2 // Final mid element of merged array
	for l <= r {
		mid := l + (r-l)/2
		lAsz, lBsz := mid, mF-mid // Since lAsz + lBsz == mF
		lA := math.MinInt32
		if lAsz > 0 {
			lA = A[lAsz-1]
		}
		lB := math.MinInt32
		if lBsz > 0 {
			lB = B[lBsz-1]
		}

		rA := math.MaxInt32
		if lAsz < n {
			rA = A[lAsz]
		}

		rB := math.MaxInt32
		if lBsz < m {
			rB = B[lBsz]
		}

		if lA <= rB && lB <= rA {
			if (n+m)%2 == 0 {
				a := utils.MaxOfIntsOrFloats[int](lA, lB)
				b := utils.MinOfIntsOrFloats[int](rA, rB)
				return (float64(a) + float64(b)) / 2.0
			}
			return float64(utils.MaxOfIntsOrFloats[int](lA, lB))
		} else if lA > rB {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return 0.0
}
