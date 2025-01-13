package mathProblems

/*
	Problem : https://www.interviewbit.com/problems/sorted-permutation-rank/
*/

// T(n) : O(n2), S(n) : O(1)
func SortedPermutationRank(A string) int {

	const MOD = 1000003
	n := len(A)
	fact := make([]int, n+1)
	fact[0] = 1
	for i := 1; i <= n; i++ {
		fact[i] = ((i % MOD) * (fact[i-1] % MOD)) % MOD
	}

	res := 1
	for i := 0; i < n; i++ {
		count := 0
		for j := i + 1; j < n; j++ {
			if A[i] > A[j] {
				count++
			}
		}
		res = ((res % MOD) + ((count%MOD)*(fact[n-i-1]%MOD))%MOD) % MOD
	}

	return res
}
