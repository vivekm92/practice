package mathPrimers

import (
	"math"
)

/*
	Problem : https://www.interviewbit.com/problems/mathbug01/
*/

// T(n) : O(sqrt(n)), S(n) : O(1)
func IsPrime(A int) int {

	if A < 2 {
		return 0
	}

	upperLimit := int(math.Sqrt(float64(A)))
	for i := 2; i <= upperLimit; i++ {
		if i < A && A%i == 0 {
			return 0
		}
	}

	return 1
}
