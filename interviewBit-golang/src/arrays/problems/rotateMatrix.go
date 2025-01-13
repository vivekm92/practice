package arrayProblems

// T(n) : O(n^2), S(n): O(1)
func Rotate(A [][]int) [][]int {

	n := len(A)
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			A[j][i], A[n-1-j][i] = A[n-1-j][i], A[j][i]
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			t := A[i][j]
			A[i][j] = A[j][i]
			A[j][i] = t
		}
	}

	return A
}

// T(n) : O(n2) ; S(n) : O(n2)
func RotateMatrix(A [][]int) [][]int {

	n := len(A)
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res[j][n-i-1] = A[i][j]
		}
	}

	return res
}
