package arrayProblems

// T(n) : O(n), S(n) : O(1)
func MaxMin(A []int) int {

	n := len(A)
	maxElement, minElement := A[0], A[0]
	for i := 1; i < n; i++ {
		if A[i] > maxElement {
			maxElement = A[i]
		}
		if A[i] < minElement {
			minElement = A[i]
		}
	}

	return maxElement + minElement
}
