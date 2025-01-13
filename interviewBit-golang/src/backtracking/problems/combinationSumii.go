package backtracking

import "sort"

/*
	Problem : https://www.interviewbit.com/problems/combination-sum-ii/
*/

func solveCombinationSumii(A []int, B []int, C *[][]int, D int, E int) {
	if D == 0 {
		F := make([]int, len(B))
		copy(F, B)
		*C = append(*C, F)
		return
	}

	for i := E; i < len(A); i++ {
		if A[i] <= D && (i == E || A[i] != A[i-1]) {
			B = append(B, A[i])
			solveCombinationSumii(A, B, C, D-A[i], i+1)
			B = B[:len(B)-1]
		}
	}
}

// T(n) : O() ; S(n) : O()
func CombinationSumii(A []int, B int) [][]int {

	n := len(A)
	res := make([][]int, 0)
	if n == 0 {
		return res
	}

	sort.Ints(A)
	temp := make([]int, 0)
	solveCombinationSumii(A, temp, &res, B, 0)
	return res
}
