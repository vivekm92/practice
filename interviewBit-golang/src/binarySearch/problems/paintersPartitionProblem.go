package bsearch

/*
	problem : https://www.interviewbit.com/problems/painters-partition-problem/
*/

// T(n) : O(nlong), S(n) : O(1)
func Paint(A int, B int, C []int) int {

	const modulo = 10000003
	l, r := 0, 0
	for _, v := range C {
		if l < v {
			l = v
		}
		r += v
	}

	res := -1
	for l <= r {
		mid := l + (r-l)/2

		p, curr := 1, 0
		for _, v := range C {
			if curr+v > mid {
				p++
				curr = v
			} else {
				curr += v
			}
		}

		if p <= A {
			res = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return ((res % modulo) * (B % modulo)) % modulo
}
