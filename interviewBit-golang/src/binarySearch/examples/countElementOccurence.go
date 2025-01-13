package bsearch

/*
	Problem : https://www.interviewbit.com/problems/count-element-occurence/
*/

// T(n) : O(logn), S(n): O(1)
func FindElementCount(A []int, B int) int {

	n := len(A)
	l, r := 0, n
	for l < r {
		mid := l + (r-l)/2
		if A[mid] < B {
			l = mid + 1
		} else {
			r = mid
		}
	}
	idx1 := l

	l, r = 0, n
	for l < r {
		mid := l + (r-l)/2
		if A[mid] <= B {
			l = mid + 1
		} else {
			r = mid
		}
	}

	idx2 := l

	if idx2 == n && idx1 == 0 && A[idx1] != B {
		return 0
	}

	return idx2 - idx1
}
