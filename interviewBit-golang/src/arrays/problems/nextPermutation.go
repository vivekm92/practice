package arrayProblems

func NextPermutation(A []int) []int {

	n, k := len(A), -1
	for i := n - 2; i >= 0; i-- {
		if A[i] < A[i+1] {
			k = i
			break
		}
	}

	if k == -1 {
		for i := 0; i < n/2; i++ {
			A[i], A[n-1-i] = A[n-1-i], A[i]
		}

		return A
	}

	x, mini := A[k+1], k+1
	for i := k + 1; i < n; i++ {
		if A[i] > A[k] {
			if x > A[i] {
				x = A[i]
			}
			if x == A[i] {
				mini = i
			}
		}
	}

	A[k], A[mini] = A[mini], A[k]

	for i := 0; i < (n-(k))/2; i++ {
		A[i+k+1], A[n-1-i] = A[n-1-i], A[i+k+1]
	}

	return A
}
