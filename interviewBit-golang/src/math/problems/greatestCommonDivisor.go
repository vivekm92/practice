package mathProblems

import "testing"

// T(n) : O(log(n)), S(n) : O(1)
func Gcd(A int, B int) int {

	if A < B {
		return Gcd(B, A)
	}

	if B == 0 {
		return A
	}

	return Gcd(B, A%B)
}

func TestGcd(t *testing.T) {

}
