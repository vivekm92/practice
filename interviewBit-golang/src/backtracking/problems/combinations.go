package backtracking

/*
	Problem : https://www.interviewbit.com/problems/combinations/
*/

func solveCombinations(A, B, C int, D []int, E *[][]int) {

	if B == len(D) {
		F := make([]int, len(D))
		copy(F, D)
		*E = append(*E, F)
		return
	}

	for i := C; i <= A; i++ {
		D = append(D, i)
		solveCombinations(A, B, i+1, D, E)
		D = D[:len(D)-1]
	}
}

// T(n) : O() ; S(n) : O()
func Combinations(A int, B int) [][]int {

	res := make([][]int, 0)
	if A == 0 {
		return res
	}

	temp := make([]int, 0)
	solveCombinations(A, B, 1, temp, &res)
	return res
}
