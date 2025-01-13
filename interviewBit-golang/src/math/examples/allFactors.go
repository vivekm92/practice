package mathExamples

/*
	problem : https://www.interviewbit.com/problems/all-factors/
*/

import (
	"math"
	"sort"
)

// T(n) : O(sqrt(n)), S(n) : O(n)
func AllFactors(A int) []int {

	factors := make([]int, 0)
	s := int(math.Sqrt(float64(A)))

	for i := 1; i <= s; i++ {
		if A%i == 0 {
			factors = append(factors, i)
			v := A / i
			if i != v {
				factors = append(factors, v)
			}
		}
	}

	sort.Slice(factors, func(i, j int) bool {
		return factors[i] < factors[j]
	})

	return factors
}
