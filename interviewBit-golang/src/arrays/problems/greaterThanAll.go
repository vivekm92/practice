package arrayProblems

import (
	"interviewBit/src/utils"
)

func GreaterThanPreviousAll(A []int) int {

	count, m := 0, A[0]
	for _, v := range A {
		if v > m {
			count++
		}
		m = utils.MaxOfIntsOrFloats(m, v)
	}

	return count + 1
}
