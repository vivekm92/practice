package arrayProblems

/*
	Problem : https://www.interviewbit.com/problems/diagonal-flip/
*/

// T(n) : O(n) , S(n) : O(1)
func DiagonalFlip(A [][]int) [][]int {

	n := len(A)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			A[i][j], A[j][i] = A[j][i], A[i][j]
		}
	}

	return A
}
