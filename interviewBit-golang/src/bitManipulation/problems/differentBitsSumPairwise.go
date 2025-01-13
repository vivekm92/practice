package bitManipulation

/*
	Problem : https://www.interviewbit.com/problems/different-bits-sum-pairwise/
*/

// T(n) : O(n) , S(n) : O(1)
func DifferentBitsSumPairwise(A []int) int {

	const MOD = 1000000007
	res, n := 0, len(A)
	for i := 0; i < 32; i++ {
		count := 0
		for _, v := range A {
			if v&(1<<i) > 0 {
				count++
			}
		}
		res = (res%MOD + ((count%MOD)*(((n-count)%MOD)*2)%MOD)%MOD) % MOD
	}

	return res
}
