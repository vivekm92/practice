package bitManipulation

/*
	Problem : https://www.interviewbit.com/problems/count-total-set-bits/
*/

const MOD = 1000000007

// T(n) : O(logn) , S(n) : O(1)
func CountTotalSetBits(A int) int {
	if A <= 1 {
		return A
	}
	x := max2(A)
	res := ((x % MOD * (1 << (x - 1) % MOD) % MOD) + (A-(1<<x)+1)%MOD) % MOD
	res = (res + CountTotalSetBits(A-(1<<x))) % MOD
	return res
}

func max2(A int) int {
	c := 0
	for (1 << c) <= A {
		c++
	}
	return c - 1
}
