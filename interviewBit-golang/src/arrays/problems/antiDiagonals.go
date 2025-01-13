package arrayProblems

/*
	problem : https://www.interviewbit.com/problems/anti-diagonals/
*/

// T(n) : O(n*m); S(n) : O(1)
func AntiDiagonal(A [][]int) [][]int {

	n, x, y := len(A), 0, 0
	res := make([][]int, 0)
	for i := 0; i < 2*n-1; i++ {
		ix, iy := x, y
		t := make([]int, 0)
		for ix < n && iy >= 0 {
			t = append(t, A[ix][iy])
			iy--
			ix++
		}
		res = append(res, t)
		if i < n-1 {
			y++
		} else {
			x++
		}
	}

	return res
}
