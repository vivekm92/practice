package mathProblems

import "math"

/*
	Problem : https://www.interviewbit.com/problems/maximum-height-of-staircase/
*/

// T(n) : O(1) , S(n) : O(1)
func MaxStaircaseHeight(A int) int {

	k := A * 2
	v := int(math.Sqrt(float64(k)))
	if v*(v+1) <= k {
		return v
	}

	return v - 1
}
