package backtracking

/*
	Problem : https://www.interviewbit.com/problems/permutations/
*/

func solvePermutations(A []int, B *[][]int, idx, n int) {
	if idx == n {
		C := make([]int, n)
		copy(C, A)
		*B = append(*B, C)
		return
	}

	for i := idx; i < n; i++ {
		A[i], A[idx] = A[idx], A[i]
		solvePermutations(A, B, idx+1, n)
		A[i], A[idx] = A[idx], A[i]
	}
}

// T(n) : O() ; S(n) : O()
func Permutations(A []int) [][]int {

	n := len(A)
	res := make([][]int, 0)
	if n == 0 {
		return res
	}

	solvePermutations(A, &res, 0, n)
	return res
}
