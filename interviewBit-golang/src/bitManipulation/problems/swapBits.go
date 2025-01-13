package bitManipulation

/*
	Problem : https://www.interviewbit.com/problems/swap-bits/
*/

// T(n) : O(1) , S(n) : O(1)
func SwapBits(A int, B int, C int) int {

	B -= 1
	C -= 1
	b1 := (A >> B) & 1
	c1 := (A >> C) & 1

	if b1 == c1 {
		return A
	}

	A = A ^ (1 << C)
	A = A ^ (1 << B)
	return A
}
