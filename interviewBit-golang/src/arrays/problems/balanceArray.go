package arrayProblems

// T(n) : O(n), S(n) : O(n)
func BalanceArray(A []int) int {

	n, even, odd := len(A), 0, 0
	rightEven, rightOdd := make([]int, n), make([]int, n)
	for i := n - 1; i >= 0; i-- {
		rightEven[i] = even
		rightOdd[i] = odd

		if i%2 == 0 {
			even += A[i]
		} else {
			odd += A[i]
		}
	}

	count := 0
	even, odd = 0, 0
	for i := 0; i < n; i++ {
		if rightEven[i]+odd == even+rightOdd[i] {
			count++
		}

		if i%2 == 0 {
			even += A[i]
		} else {
			odd += A[i]
		}
	}

	return count
}
