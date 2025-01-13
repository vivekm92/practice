package bitManipulation

/*
problem : https://www.interviewbit.com/problems/trailing-zeroes/
*/

// T(n): O(logN), S(n) : O(1)
func CountTrailingZeros(A int) int {

	cnt := 0
	for A&1 == 0 {
		cnt++
		A = A >> 1
	}

	return cnt
}
