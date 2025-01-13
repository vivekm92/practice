package mathProblems

import "testing"

// T(n) : O(log(n)), S(n) : O(1)
func TrailingZeroes(A int) int {

	numZeros := 0
	for A > 0 {
		numZeros += A / 5
		A /= 5
	}

	return numZeros
}

func TestTrailingZeroes(t *testing.T) {

}
