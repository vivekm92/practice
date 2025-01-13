package bitManipulation

/*
	Problem : https://www.interviewbit.com/problems/bit-flipping/
*/

// T(n) : O(logn) , S(n) : O(1)
func BitFlipping(A int) int {

	res, pos := 0, 0
	for A > 0 {
		t := A & 1
		res |= (t ^ 1) << pos
		A = A >> 1
		pos++
	}

	return res
}
