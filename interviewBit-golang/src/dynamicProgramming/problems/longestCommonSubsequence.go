package dpProblems

/*
	Problem : https://www.interviewbit.com/problems/longest-common-subsequence/
*/

// func solveLCSLengthNaiveRecursive(A string, B string, i1 int, i2 int) int {
// 	if i1 == len(A) || i2 == len(B) {
// 		return 0
// 	}

// 	v1 := 0
// 	if A[i1] != B[i2] {
// 		return solveLCSLengthNaiveRecursive(A, B, i1+1, i2) + solveLCSLengthNaiveRecursive(A, B, i1, i2+1)
// 	}
// 	return solveLCSLengthNaiveRecursive(A, B, i1+1, i2+1) + 1
// }

// // Approach : Naive Recursive
// // T(n) : O() ; S(n) : O()
// func LCSLengthNaiveRecursive(A string, B string) int {
// 	return solveLCSLengthNaiveRecursive(A, B, 0, 0)
// }

// Approach : Iterative DP
// T(n) : O(n*m), S(n) : O(n*m)
func LCSLength(A string, B string) int {

	n, m := len(A), len(B)
	dp := make([][]int, 0)
	for i := 0; i <= n; i++ {
		t := make([]int, m+1)
		dp = append(dp, t)
	}

	for i := 0; i <= m; i++ {
		dp[0][i] = 0
	}

	for i := 0; i <= n; i++ {
		dp[i][0] = 0
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j]
				if dp[i][j-1] > dp[i-1][j] {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}

	return dp[n][m]
}
