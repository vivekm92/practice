package backtracking

/*
	Problem : https://www.interviewbit.com/problems/all-unique-permutations/
*/

func solveAllUniquePermutations(A map[int]int, B []int, C *[][]int, D int) {
	if len(B) == D {
		E := make([]int, D)
		copy(E, B)
		*C = append(*C, E)
		return
	}

	for k, v := range A {
		if v == 0 {
			continue
		}
		A[k]--
		B = append(B, k)
		solveAllUniquePermutations(A, B, C, D)
		A[k]++
		B = B[:len(B)-1]
	}
}

// T(n) : O() ; S(n) : O()
func AllUniquePermutations(A []int) [][]int {

	n := len(A)
	res := make([][]int, 0)
	if n == 0 {
		return res
	}

	m := make(map[int]int, 0)
	for _, v := range A {
		m[v]++
	}
	t := make([]int, 0)
	solveAllUniquePermutations(m, t, &res, n)
	return res
}
