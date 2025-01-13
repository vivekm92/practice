package mathProblems

import (
	"math"
)

/*
	Problem : https://www.interviewbit.com/problems/pythagorean-triplets/
*/

// T(n) : O(n2) , S(n) : O(1)
func PythagoreanTriplets(A int) int {

	count := 0
	for i := 1; i <= A; i++ {
		for j := i + 1; j <= A; j++ {
			t := i*i + j*j
			s := int(math.Sqrt(float64(t)))
			if s <= A && s*s == t {
				count++
			}
		}
	}

	return count
}
