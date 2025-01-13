package backtracking

import "sort"

/*
	Problem : https://www.interviewbit.com/problems/subset/
*/

func solveSubsets(A []int, B []int, C *[][]int, D int) {
	if D == len(A) {
		E := make([]int, len(B))
		copy(E, B)
		*C = append(*C, E)
		return
	}

	B = append(B, A[D])
	solveSubsets(A, B, C, D+1)
	B = B[:len(B)-1]
	solveSubsets(A, B, C, D+1)
}

func Subset(A []int) [][]int {

	res := make([][]int, 0)
	sort.Ints(A)
	temp := make([]int, 0)
	solveSubsets(A, temp, &res, 0)

	sort.Slice(res, func(i, j int) bool {
		for x := range res[i] {
			if x < len(res[j]) {
				if res[i][x] == res[j][x] {
					continue
				}
				return res[i][x] < res[j][x]
			}
			return false
		}
		return len(res[i]) < len(res[j])
	})
	return res
}
