package bsearch

/*
	Problem : https://www.interviewbit.com/problems/first-index/
*/

// T(n) : O(nlong), S(n) : O(1)
func FirstIndex(A []int, B []int) []int {

	var p pair
	arr, n := make([]pair, 0), len(A)
	for i := 0; i < n; i++ {
		if i == 0 || p.val <= A[i] {
			p.idx, p.val = i, A[i]
			arr = append(arr, p)
		}
	}

	n = len(B)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = bsearch(arr, B[i])
	}
	return res
}

func bsearch(arr []pair, t int) int {

	l, r, ans := 0, len(arr), -1
	for l < r {
		mid := l + (r-l)/2
		if arr[mid].val >= t {
			ans = arr[mid].idx
			r = mid
		} else {
			l = mid + 1
		}
	}
	return ans
}

type pair struct {
	idx, val int
}
