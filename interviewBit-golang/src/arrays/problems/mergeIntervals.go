package arrayProblems

import "interviewBit/src/utils"

type Interval struct {
	start, end int
}

// T(n): O(n), S(n): O(n)
func MergeInterval(intervals []Interval, newInterval Interval) []Interval {

	res := make([]Interval, 0)
	res = append(res, newInterval)
	for _, interval := range intervals {
		k := len(res)
		s, e := res[k-1].start, res[k-1].end
		if interval.start > e {
			res = append(res, interval)
		} else if interval.end < s {
			t := res[k-1]
			res = res[:k-1]             // pop_back
			res = append(res, interval) // push_back
			res = append(res, t)        // push_back
		} else {
			res = res[:k-1] //pop_back
			s = utils.MinOfIntsOrFloats[int](s, interval.start)
			e = utils.MaxOfIntsOrFloats[int](e, interval.end)

			i := new(Interval)
			i.start, i.end = s, e
			res = append(res, *i) //push_back
		}
	}

	return res
}
