package arrayProblems

// T(n): O(n^2), S(n): O(1)
func MaximumSum(A [][]int, B int) int {

	n, m := len(A), len(A[0])
	for i := 0; i < n; i++ {
		for j := 1; j < m; j++ {
			A[i][j] += A[i][j-1]
		}
	}

	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			A[i][j] += A[i-1][j]
		}
	}

	maxSum := -100000001
	for i := B - 1; i < n; i++ {
		for j := B - 1; j < m; j++ {
			currSum := A[i][j]
			if i >= B {
				currSum -= A[i-B][j]
			}
			if j >= B {
				currSum -= A[i][j-B]
			}
			if i >= B && j >= B {
				currSum += A[i-B][j-B]
			}

			if maxSum < currSum {
				maxSum = currSum
			}
		}

	}

	return maxSum
}
