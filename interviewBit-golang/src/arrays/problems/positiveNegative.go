package arrayProblems

func CountPositiveNegative(A []int) []int {

	res := make([]int, 2)
	for i := 0; i < len(A); i++ {
		if A[i] > 0 {
			res[0]++
		} else if A[i] < 0 {
			res[1]++
		}
	}

	return res
}
