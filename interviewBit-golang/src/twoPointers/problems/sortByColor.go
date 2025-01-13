package twoPointers

/*
	problem : https://www.interviewbit.com/problems/sort-by-color/
*/

import (
	"sort"
)

// T(n) : O(nlogn), S(n) : O(1)
func SortColor(A []int) []int {

	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})

	return A
}

// T(n) : O(n), S(n) : O(1)
func SortColor2(A []int) []int {

	cntRed, cntWhite, cntBlue := 0, 0, 0
	for _, v := range A {
		switch v {
		case 0:
			cntRed++
		case 1:
			cntWhite++
		default:
			cntBlue++
		}
	}

	n := len(A)
	for i := 0; i < n; i++ {
		if cntRed > 0 {
			A[i] = 0
			cntRed--
		} else if cntWhite > 0 {
			A[i] = 1
			cntWhite--
		} else {
			A[i] = 2
			cntBlue--
		}
	}

	return A
}

// T(n) : O(n), S(n) : O(1)
func SortColor3(A []int) []int {

	n := len(A)
	i, j, k := 0, n-1, n-1
	for i < k {
		if A[i] == 0 {
			i += 1
		} else if A[i] == 1 {
			if i < j {
				t1 := A[i]
				A[i] = A[j]
				A[j] = t1
				j -= 1
			} else {
				i += 1
			}
		} else {
			t1 := A[k]
			A[k] = A[i]
			A[i] = t1
			k -= 1
			if j > k {
				j = k
			}
		}
	}
	return A
}
