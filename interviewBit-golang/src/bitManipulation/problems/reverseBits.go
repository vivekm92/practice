package bitManipulation

/*
	problem : https://www.interviewbit.com/problems/reverse-bits/
*/

// T(n) : O(1), S(n) : O(1)
func ReverseBits(A uint32) uint32 {

	t := make([]uint32, 0)
	for i := 0; i < 32; i++ {
		t = append(t, A&1)
		A = A >> 1
	}

	var res uint32
	for i := 0; i < 32; i++ {
		res = res*2 + t[i]
	}

	return res
}
