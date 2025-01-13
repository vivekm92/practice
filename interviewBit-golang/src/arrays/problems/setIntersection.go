package arrayProblems

import (
	"sort"
)

// T(n): O(nlogn), S(n): O(1)
func SetIntersection(A [][]int) int {

	sort.Slice(A, func(i, j int) bool {
		if A[i][1] == A[j][1] {
			return A[i][0] > A[j][0]
		}
		return A[i][1] < A[j][1]
	})

	n := len(A)
	if n == 1 {
		return 2
	}

	cnt := 2
	rs, re := A[0][1]-1, A[0][1]
	for i := 0; i < n; i++ {
		s, e := A[i][0], A[i][1]
		if s > rs {
			if s > re {
				cnt += 2
				re = e
				rs = e - 1
			} else if s <= re {
				cnt += 1
				rs = re
				re = e
			}
		}
	}

	return cnt
}
