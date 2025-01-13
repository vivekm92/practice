package mathProblems

import "testing"

// T(n) : O(log(n)), S(n) : O(1)
func IsPalindrome(A int) int {

	if A < 0 {
		return 0
	}

	n, curr := 0, A
	for curr > 0 {
		n = n*10 + curr%10
		curr /= 10
	}

	if n == A {
		return 1
	}

	return 0
}

func TestIsPalindrome(t *testing.T) {

}
