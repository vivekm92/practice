package bsearch

/*
	Problem : https://www.interviewbit.com/problems/woodcutting-made-easy/
*/

// T(n): O(n)
func check(A []int, k, th, B int) bool {

	rh := 0
	for _, v := range A {
		if v < k {
			rh += v
		} else {
			rh += k
		}
	}

	acq := th - rh

	return acq >= B
}

// T(n): (nlogn), S(n): O(1)
func MaximumSawBladeHeight(A []int, B int) int {

	l, r, th := 0, 0, 0
	for _, v := range A {
		if v > r {
			r = v
		}
		th += v
	}

	for l < r {
		mid := l + (r-l)/2
		if check(A, mid, th, B) {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l - 1
}
