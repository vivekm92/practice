package bsearch

/*
	problem : https://www.interviewbit.com/problems/rotated-array/
*/

// T(n) : O(log(n)), S(n) : O(1)
func FindMin(A []int) int {

	n := len(A)
	l, r := 0, n
	for l < r {
		mid := l + (r-l)/2
		if A[mid] > A[n-1] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return A[l]
}

// T(n) : O(log(n)), S(n) : O(1)
func FindMin2(A []int) int {

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

	return A[l]
}
