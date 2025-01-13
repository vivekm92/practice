package greedy

import "sort"

/*
	Problem : https://www.interviewbit.com/problems/disjoint-intervals/
*/

// T(n) : O(nlogn) ; S(n) : O(1)
func DisjointIntervals(A [][]int) int {

	n := len(A)
	if n == 0 {
		return 0
	}

	sort.Slice(A, func(i, j int) bool {
		if A[i][1] == A[j][1] {
			return A[i][0] < A[j][0]
		}
		return A[i][1] < A[j][1]
	})

	i, cnt, e := 0, 0, A[0][1]
	for i < n {
		for i < n && e >= A[i][0] {
			i++
		}
		cnt++
		if i < n {
			e = A[i][1]
		}
	}

	return cnt
}
