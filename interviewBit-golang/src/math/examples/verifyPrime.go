package mathExamples

// T(n) : O(sqrt(n)), S(n) : O(1)
func IsPrime(A int) int {

	if A < 2 {
		return 0
	}

	for i := 2; i*i <= A; i++ {
		if A%i == 0 {
			return 0
		}
	}

	return 1
}
