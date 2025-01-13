package arrayProblems

// T(n): O(n^2), S(n): O(1)
func SetZeros(A [][]int) [][]int {

	n, m := len(A), len(A[0])
	lr, lc := make(map[int]bool, 0), make(map[int]bool, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if A[i][j] == 0 {
				lr[i] = true
				lc[j] = true
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if _, ok := lr[i]; ok {
				A[i][j] = 0
			}
			if _, ok := lc[j]; ok {
				A[i][j] = 0
			}
		}
	}

	return A
}
