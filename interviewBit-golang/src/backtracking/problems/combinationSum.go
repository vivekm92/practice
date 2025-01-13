package backtracking

import "sort"

/*
	Problem : https://www.interviewbit.com/problems/combination-sum/
*/

func solveCombinationSum(A []int, B int, C []int, D *[][]int, i int) {

	if B == 0 {
		E := make([]int, len(C))
		copy(E, C)
		*D = append(*D, E)
		return
	}

	for i < len(A) && A[i] <= B {
		C = append(C, A[i])
		solveCombinationSum(A, B-A[i], C, D, i)
		i += 1
		C = C[:len(C)-1]
	}
}

// T(n) : O() , S(n) : O()
func CombinationSum(A []int, B int) [][]int {

	D := make([][]int, 0)
	n := len(A)
	if n == 0 {
		return D
	}

	sort.Ints(A)
	t := make([]int, 0)
	t = append(t, A[0])
	for i := 1; i < n; i++ {
		if A[i] != A[i-1] {
			t = append(t, A[i])
		}
	}

	C := make([]int, 0)
	solveCombinationSum(t, B, C, &D, 0)
	return D
}
