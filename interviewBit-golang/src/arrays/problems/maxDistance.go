package arrayProblems

import (
	"interviewBit/src/utils"
	"sort"
)

type distancePair struct {
	val int
	idx int
}

// T(n): O(nlogn), S(n): O(n)
func MaxDistance(A []int) int {

	n := len(A)
	var distancePairs []distancePair
	for i := 0; i < n; i++ {
		dPair := new(distancePair)
		dPair.idx = i
		dPair.val = A[i]
		distancePairs = append(distancePairs, *dPair)
	}

	sort.Slice(distancePairs, func(i, j int) bool {
		if distancePairs[i].val == distancePairs[j].val {
			return distancePairs[i].idx < distancePairs[j].idx
		}
		return distancePairs[i].val < distancePairs[j].val
	})

	maxDistance := 0
	ix := distancePairs[0].idx
	for i := 1; i < n; i++ {
		curr := distancePairs[i].idx - ix
		ix = utils.MinOfIntsOrFloats(ix, distancePairs[i].idx)
		maxDistance = utils.MaxOfIntsOrFloats(maxDistance, curr)
	}

	return maxDistance
}
