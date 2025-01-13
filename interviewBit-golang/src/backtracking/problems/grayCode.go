package backtracking

/*
	Problem : https://www.interviewbit.com/problems/gray-code/
*/

// T(n) : O() ; S(n) : O()
func GrayCode(A int) []int {

	if A == 0 {
		return []int{}
	} else if A == 1 {
		return []int{0, 1}
	}

	prevGC := GrayCode(A - 1)
	t, n := 1<<(A-1), len(prevGC)
	for i := n - 1; i >= 0; i-- {
		prevGC = append(prevGC, t|prevGC[i])
	}

	return prevGC
}
