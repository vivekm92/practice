package bitManipulation

/*
problem : https://www.interviewbit.com/problems/number-of-1-bits/
*/

// T(n) : O(logN), S(n) : O(1)
func CountNumberOfSetBits(A int) int {

	var cnt int
	for A > 0 {
		cnt += (A & 1)
		A = A >> 1
	}

	return cnt
}
