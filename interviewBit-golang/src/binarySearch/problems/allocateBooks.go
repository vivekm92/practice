package bsearch

/*
	problem : https://www.interviewbit.com/problems/allocate-books/
*/

// T(n) : (nlogn), S(n) : O(1)
func AllocateBooks(A []int, B int) int {

	n := len(A)
	if B > n {
		return -1
	}

	l, r := 0, 0
	for _, v := range A {
		if l < v {
			l = v
		}
		r += v
	}

	res := -1
	for l <= r {
		mid := l + (r-l)/2
		cnt, curr := 1, 0
		for _, v := range A {
			if curr+v > mid {
				cnt++
				curr = v
			} else {
				curr += v
			}
		}

		if cnt <= B {
			res = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return res
}
