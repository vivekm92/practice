package bsearch

/*
	Problem : https://www.interviewbit.com/problems/find-a-peak-element/
*/

// T(n) : O(logN), S(n) : O(1)
func FindPeak(A []int) int {

	n := len(A)
	l, r := 0, n
	for l < r {
		mid := l + (r-l)/2
		if (mid == 0 || A[mid] >= A[mid-1]) && (mid == n-1 || A[mid] >= A[mid+1]) {
			return A[mid]
		} else if A[mid] >= A[mid-1] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return -1
}

// T(n): O(n), S(n) : O(1)
func FindPeak1(A []int) int {

	n := len(A)
	for i := 1; i < n-1; i++ {
		if A[i] >= A[i-1] && A[i] >= A[i+1] {
			return A[i]
		}
	}

	if A[0] >= A[1] {
		return A[0]
	}

	return A[n-1]
}
