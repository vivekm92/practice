package mathProblems

// T(n) : O(sqrt(n)), S(n) : O(a)
func StepByStep(A int) int {

	if A < 0 {
		return StepByStep(-A)
	}

	step, sum := 0, 0
	for sum < A || (sum-A)%2 != 0 {
		step++
		sum += step
	}

	return step
}
