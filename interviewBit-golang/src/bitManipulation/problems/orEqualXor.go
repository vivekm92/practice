package bitManipulation

/*
	Problem : https://www.interviewbit.com/problems/or-equal-xor/
*/

// T(n) : O(logn) , S(n) : O(1)
func OrEqualXor(A int) int {

	const MOD = 1000000007
	res := 1
	for A > 0 {
		res = ((res % MOD) * (3 % MOD)) % MOD
		A--
	}

	return res
}
