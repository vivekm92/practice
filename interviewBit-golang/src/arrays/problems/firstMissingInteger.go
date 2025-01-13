package arrayProblems

// T(n) : O(n), S(n): O(n)
func FirstMissingPositive(A []int) int {

	n := len(A)
	arr := make([]bool, n+1)
	for i := 0; i < n; i++ {
		if A[i] > 0 && A[i] <= n {
			arr[A[i]] = true
		}
	}

	for i := 1; i <= n; i++ {
		if !arr[i] {
			return i
		}
	}

	return n + 1
}
