package mathProblems

/*
	Problem : https://www.interviewbit.com/problems/sum-of-pairwise-hamming-distance/
*/

// T(n): O(n) , S(n) : O(1)
func HammingDistance(A []int) int {

	if len(A) == 1 {
		return 0
	}

	const MOD = 1000000007
	totalDistance, n := 0, len(A)
	for i := 0; i < 32; i++ {
		oneCount, zeroCount := 0, 0
		for j := 0; j < n; j++ {
			if A[j]%2 != 0 {
				oneCount++
			} else {
				zeroCount++
			}
			A[j] = A[j] / 2
		}
		totalDistance = (totalDistance%MOD + ((oneCount%MOD)*(zeroCount%MOD))%MOD) % MOD
	}

	return ((totalDistance % MOD) * (2 % MOD)) % MOD
}
