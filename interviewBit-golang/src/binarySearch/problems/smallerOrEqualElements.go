package bsearch

/*
	Problem : https://www.interviewbit.com/problems/smaller-or-equal-elements/
*/

// T(n) : O(logn), S(n): O(1)
func SmallerOrEqualElements(A []int, B int) int {

	n := len(A)
	l, r := 0, n

	for l < r {
		mid := l + (r-l)/2
		if A[mid] <= B {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}
