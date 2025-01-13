package mathProblems

// T(n) : O(n), S(n) = O(b); where n equals to (A + B)
func UniquePaths(A int, B int) int {
	if A == 0 || B == 0 {
		return 0
	}

	dp := make([][]int, A)
	for i := 0; i < A; i++ {
		dp[i] = make([]int, B)
	}

	dp[0][0] = 1
	for i := 1; i < A; i++ {
		dp[i][0] = 1
	}

	for i := 1; i < B; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < A; i++ {
		for j := 1; j < B; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[A-1][B-1]
}
