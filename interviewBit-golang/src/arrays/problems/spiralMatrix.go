package arrayProblems

/*
	Problem : https://www.interviewbit.com/problems/spiral-matrix/
*/

// T(n) : O(n), S(n) : O(1)
func SpiralMatrix(A []int, B int, C int) [][]int {

	res := make([][]int, 0)
	for i := 0; i < B; i++ {
		arr := make([]int, C)
		res = append(res, arr)
	}

	i, n := 0, len(A)
	rs, re, cs, ce := 0, B-1, 0, C-1
	for i < n {
		for j := cs; j <= ce && i < n; j++ {
			res[rs][j] = A[i]
			i++
		}
		rs++

		for j := rs; j <= re && i < n; j++ {
			res[j][ce] = A[i]
			i++
		}
		ce--

		for j := ce; j >= cs && i < n; j-- {
			res[re][j] = A[i]
			i++
		}
		re--

		for j := re; j >= rs && i < n; j-- {
			res[j][cs] = A[i]
			i++
		}
		cs++
	}

	return res
}
