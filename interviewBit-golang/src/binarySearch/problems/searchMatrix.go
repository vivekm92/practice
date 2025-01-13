package bsearch

/*
	Problem : https://www.interviewbit.com/problems/matrix-search/
*/

// n : #rows, k : max(len(__all__rows__))
// T(n): O(nlogk), S(n): O(1)
func SearchMatrix(A [][]int, B int) int {

	for _, v := range A {

		l, r := 0, len(v)
		for l < r {
			mid := l + (r-l)/2
			if v[mid] == B {
				return 1
			} else if v[mid] < B {
				l = mid + 1
			} else {
				r = mid
			}
		}
	}

	return 0
}
