package bsearch

/*
	Problem : https://www.interviewbit.com/problems/search-for-a-range/
*/

// T(n): O(logn), S(n): O(1)
func SearchRange(A []int, B int) []int {

	n := len(A)
	l, r := 0, n
	for l < r {
		mid := l + (r-l)/2
		if A[mid] >= B {
			r = mid
		} else {
			l = mid + 1
		}
	}

	if l == 0 && A[l] != B {
		return []int{-1, -1}
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

	if idx1 >= l {
		return []int{-1, -1}
	}

	return []int{idx1, l - 1}
}
