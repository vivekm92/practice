package dpProblems

/*
	Problem : https://www.interviewbit.com/problems/0-1-knapsack/
*/

// T(n) : O(2^N) , S(n) : O(n)
func solveKnapSackRecursiveNaive(A []int, B []int, C int, N int) int {
	if N == 0 || C == 0 {
		return 0
	}

	v1 := solveKnapSackRecursiveNaive(A, B, C, N-1)
	if C < B[N-1] {
		return v1
	}

	v2 := solveKnapSackRecursiveNaive(A, B, C-B[N-1], N-1) + A[N-1]
	if v1 > v2 {
		return v1
	}
	return v2
}

// Driver Method for solveKnapSackRecursiveNaive
func KnapSackRecursiveNaive(A []int, B []int, C int) int {
	return solveKnapSackRecursiveNaive(A, B, C, len(A))
}

// T(n) : O(C*N) , S(n) : O(C*N)
func solveKnapSackRecursive(A []int, B []int, C int, N int, dp [][]int) int {
	if N == 0 || C == 0 {
		return 0
	}

	if dp[C][N] != -1 {
		return dp[C][N]
	} else if C < B[N-1] {
		return solveKnapSackRecursive(A, B, C, N-1, dp)
	}

	v1 := solveKnapSackRecursive(A, B, C, N-1, dp)
	v2 := solveKnapSackRecursive(A, B, C-B[N-1], N-1, dp) + A[N-1]
	if v1 > v2 {
		dp[C][N] = v1
	} else {
		dp[C][N] = v2
	}

	return dp[C][N]
}

// Driver Method for solveKnapSackRecursive
func KnapSackRecursive(A []int, B []int, C int) int {

	n := len(A)
	dp := make([][]int, C+1)
	for i := 0; i <= C; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= C; i++ {
		for j := 0; j <= n; j++ {
			dp[i][j] = -1
		}
	}

	return solveKnapSackRecursive(A, B, C, n, dp)
}

// T(n) : O(C*N) , S(n) : O(C*N)
func KnapSack(A []int, B []int, C int) int {

	n := len(A)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, C+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= C; j++ {
			if B[i-1] > j {
				dp[i][j] = dp[i-1][j]
				continue
			}

			if dp[i-1][j] < A[i-1]+dp[i-1][j-B[i-1]] {
				dp[i][j] = A[i-1] + dp[i-1][j-B[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[n][C]
}

// func KnapSackOptimised(A []int, B []int, C int) int {

// 	n := len(A)
// 	dp := make([][]int, 2)
// 	for i := 0; i < 2; i++ {
// 		dp[i] = make([]int, C+1)
// 	}

// 	for i := 0; i <= n; i++ {
// 		for j := 0; j <= C; j++ {
// 			if
// 		}
// 	}

// 	return dp[1][C]
// }
