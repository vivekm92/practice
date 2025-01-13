package bsearch

/*
	problem : https://www.interviewbit.com/problems/rotated-sorted-array-search/
*/

// T(n) : O(logn), S(n) : O(logn), we can have atmost logn recursive calls
func binarySearch(A []int, B int, l int, r int) int {
	if l >= r {
		return -1
	}
	mid := l + (r-l)/2
	if A[mid] == B {
		return mid
	} else if A[mid] > B {
		return binarySearch(A, B, l, mid)
	}
	return binarySearch(A, B, mid+1, r)
}

// T(n) : O(2*logn), S(n) : O(1)
func RotatedSearch(A []int, B int) int {

	n := len(A)
	l, r := 0, n
	for l < r {
		mid := l + (r-l)/2
		if A[mid] >= A[0] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	if l == n {
		l = 0
	}

	if B <= A[n-1] {
		return binarySearch(A, B, l, n)
	}
	return binarySearch(A, B, 0, l)
}
