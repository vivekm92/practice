package bsearch

/*
	Problem : https://www.interviewbit.com/problems/search-in-bitonic-array/
*/

// T(n) : O(logn), S(n) : O(1)
func SearchInBitnoicArray(A []int, B int) int {

	n := len(A)
	l, r := 0, n

	// find Peak
	for l < r {
		mid := l + (r-l)/2
		if A[mid] > A[mid-1] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	// Search in Left Array
	x, y := 0, l
	for x < y {
		mid := x + (y-x)/2
		if A[mid] == B {
			return mid
		} else if A[mid] < B {
			x = mid + 1
		} else {
			y = mid
		}
	}

	// Search in Right Array
	x, y = l, n
	for x < y {
		mid := x + (y-x)/2
		if A[mid] == B {
			return mid
		} else if A[mid] > B {
			x = mid + 1
		} else {
			y = mid
		}
	}

	return -1
}
