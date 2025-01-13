package arrayProblems

import (
	"math"
)

// T(n) : O(n), S(n) : O(1)
func MaxSubArray(A []int) int {

	currSum := math.Inf(-1)
	res := math.Inf(-1)

	for i := 0; i < len(A); i++ {
		currVal := float64(A[i])

		if currSum+currVal < currVal {
			currSum = currVal
		} else {
			currSum += currVal
		}

		if res < currSum {
			res = currSum
		}
	}

	return int(res)
}
