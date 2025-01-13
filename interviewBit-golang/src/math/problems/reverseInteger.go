package mathProblems

import "testing"

// T(n) : O(log(n)), S(n) : O(1)
func Reverse(A int) int {

	rev := 0
	const INT_MAX = (1 << 31) - 1
	const INT_MIN = (-1 << 31)

	for A != 0 {
		val := A % 10
		A /= 10

		if (rev > INT_MAX/10) || (rev == INT_MAX/10 && val > INT_MAX%10) {
			return 0
		}

		if (rev < INT_MIN/10) || (rev == INT_MIN/10 && val < INT_MIN%10) {
			return 0
		}

		rev = rev*10 + val
	}

	return rev
}

func TestReverse(t *testing.T) {

}
