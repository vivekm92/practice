package mathProblems

import (
	"math"
)

/*
	Problem : https://www.interviewbit.com/problems/armstrong-number/
*/

// T(n) : O(logn), S(n) : O(1)
func ArmstrongNumber(A int) int {

	k, n, res := A, 0, 0
	for k > 0 {
		k /= 10
		n++
	}

	k = A
	for k > 0 {
		t := k % 10
		res += int(math.Pow(float64(t), float64(n)))
		k /= 10
	}

	if res == A {
		return 1
	}

	return 0
}
