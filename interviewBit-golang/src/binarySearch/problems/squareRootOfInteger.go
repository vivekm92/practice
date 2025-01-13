package bsearch

/*
	Problem : https://www.interviewbit.com/problems/square-root-of-integer/
*/

// T(n): O(logn), S(n): O(1)
func Sqrt(A int) int {

	if A < 2 {
		return A
	}

	l, r := 1, A/2
	for l < r {
		mid := l + (r-l)/2

		if A == mid*mid || (mid*mid < A && (mid+1)*(mid+1) > A) {
			return mid
		} else if A < mid*mid {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}
