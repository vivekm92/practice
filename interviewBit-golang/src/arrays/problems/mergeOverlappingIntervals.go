package arrayProblems

import (
	"interviewBit/src/utils"
	"sort"
)

// T(n): O(n), S(n): O(n)
func MergeOverlappingIntervals(A []Interval) []Interval {

	sort.Slice(A, func(i, j int) bool {
		return A[i].start <= A[j].start
	})

	res, n := make([]Interval, 0), len(A)
	if n == 0 {
		return res
	}

	res = append(res, A[0])

	for i := 1; i < n; i++ {

		k := len(res) - 1
		s, e := res[k].start, res[k].end
		if A[i].start > e {
			res = append(res, A[i])
		} else {
			s = utils.MinOfIntsOrFloats[int](s, A[i].start)
			e = utils.MaxOfIntsOrFloats[int](e, A[i].end)

			res[k].start = s
			res[k].end = e
		}
	}

	return res
}
