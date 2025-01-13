package dpProblems

/*
	Problem : https://www.interviewbit.com/problems/subset-sum-problem/
*/

// T(n) : O(2^n) , S(n) : O(n)
func solveSubsetSumRecursiveNaive(A []int, B int, idx int) int {
	if B == 0 {
		return 1
	} else if idx >= len(A) {
		return 0
	}

	v1 := solveSubsetSumRecursiveNaive(A, B-A[idx], idx+1)
	v2 := solveSubsetSumRecursiveNaive(A, B, idx+1)

	if v1+v2 >= 1 {
		return 1
	}

	return 0
}

func SubsetSumRecursiveNaive(A []int, B int) int {
	return solveSubsetSumRecursiveNaive(A, B, 0)
}

// T(n) : O(n*x) ; S(n) : O(n*x)
func solveSubSetSumRecusive(A []int, B int, idx int, total int, dp [][]int) int {

	if total == B {
		return 1
	} else if idx == len(A) {
		return 0
	}

	if dp[idx][total] != -1 {
		return dp[idx][total]
	}

	v1 := solveSubSetSumRecusive(A, B, idx+1, total+A[idx], dp)
	v2 := solveSubSetSumRecusive(A, B, idx+1, total, dp)
	if v1+v2 > 0 {
		dp[idx][total] = 1
	} else {
		dp[idx][total] = 0
	}

	return dp[idx][total]
}

func SubSetSumRecusive(A []int, B int) int {

	n, x := len(A), 0
	for _, v := range A {
		x += v
	}

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, x+1)
	}
	for i := 0; i <= n; i++ {
		for j := 0; j <= x; j++ {
			dp[i][j] = -1
		}
	}

	return solveSubSetSumRecusive(A, B, 0, 0, dp)
}

// T(n) : O(n*B) , S(n) : O(n*B)
func SubSetSum(A []int, B int) int {

	n := len(A)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, B+1)
	}

	for i := 0; i <= n; i++ {
		for j := 0; j <= B; j++ {
			if j == 0 {
				dp[i][j] = 1
			} else if i == 0 {
				dp[i][j] = 0
			} else if j < A[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				v1 := dp[i-1][j]
				v2 := dp[i-1][j-A[i-1]]
				if v1+v2 > 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = 0
				}
			}
		}
	}

	return dp[n][B]
}

// T(n) : O(n*B) ; S(n) : O(B)
func SubSetSumOptimised(A []int, B int) int {
	return -1
}
