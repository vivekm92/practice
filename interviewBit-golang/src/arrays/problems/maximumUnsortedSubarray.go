package arrayProblems

import "interviewBit/src/utils"

func SubUnsort(A []int) []int {

	n, idx1 := len(A), -1
	for i := 0; i < n-1; i++ {
		if A[i] > A[i+1] {
			idx1 = i
			break
		}
	}

	if idx1 == -1 {
		return []int{-1}
	}

	idx2 := -1
	for i := n - 1; i > 0; i-- {
		if A[i] < A[i-1] {
			idx2 = i
			break
		}
	}

	mini, maxi := A[idx1], A[idx1]
	for i := idx1; i <= idx2; i++ {
		maxi = utils.MaxOfIntsOrFloats[int](maxi, A[i])
		mini = utils.MinOfIntsOrFloats[int](mini, A[i])
	}

	for i := 0; i < idx1; i++ {
		if A[i] > mini {
			idx1 = i
			break
		}
	}

	for i := n - 1; i > idx2; i-- {
		if A[i] < maxi {
			idx2 = i
			break
		}
	}

	return []int{idx1, idx2}
}
