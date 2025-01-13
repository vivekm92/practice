package backtracking

import (
	"sort"
)

/*
	Problem : https://www.interviewbit.com/problems/subsets-ii/
*/

func solveSubsetsii(A []int, B []int, C *[][]int, D int) {
	E := make([]int, len(B))
	copy(E, B)
	*C = append(*C, E)

	for i := D; i < len(A); i++ {
		if i != D && A[i] == A[i-1] {
			continue
		}
		B = append(B, A[i])
		solveSubsetsii(A, B, C, i+1)
		B = B[:len(B)-1]
	}
}

// T(n) : O() ; S(n) : O()
func Subsetsii(A []int) [][]int {

	res := make([][]int, 0)
	temp := make([]int, 0)
	sort.Ints(A)
	solveSubsetsii(A, temp, &res, 0)
	return res
}
