package bsearch

/*
	Problem : https://www.interviewbit.com/problems/capacity-to-ship-packages-within-b-days/
*/

func canShipInBDays(A []int, wc, B int) bool {

	n, d, ts := len(A), 1, 0
	for i := 0; i < n; i++ {
		ts += A[i]

		if ts > wc {
			d++
			ts = A[i]
		}
	}

	return d <= B
}

// T(n): O(nlogn), S(n) : O(1)
func FindLeastCapacity(A []int, B int) int {
	l, r := 0, 0
	for _, v := range A {
		r += v
		if l < v {
			l = v
		}
	}

	for l < r {
		mid := l + (r-l)/2
		if canShipInBDays(A, mid, B) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}
