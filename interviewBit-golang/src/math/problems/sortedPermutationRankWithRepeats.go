package mathProblems

/*
	Problem : https://www.interviewbit.com/problems/sorted-permutation-rank-with-repeats/
*/

// T(n) : O(n2) ; S(n) : O(1)
func SortedPermutationRankWithRepeats(A string) int {

	n := len(A)
	const MOD = 1000003
	fact := make([]int, n+1)
	fact[0] = 1
	for i := 1; i <= n; i++ {
		fact[i] = (i % MOD * fact[i-1] % MOD) % MOD
	}

	res := 1
	for i := 0; i < n; i++ {
		count := 0
		for j := i + 1; j < n; j++ {
			if A[i] > A[j] {
				count++
			}
		}

		freq := make(map[rune]int)
		for j := i; j < n; j++ {
			freq[rune(A[j])]++
		}

		d := 1
		for _, v := range freq {
			d = ((d % MOD) * (fact[v] % MOD)) % MOD
		}

		inv := inverseModulo(d, MOD-2, MOD) % MOD
		p := (inv % MOD * fact[n-i-1] % MOD) % MOD
		res = (res%MOD + ((p%MOD)*(count%MOD))%MOD) % MOD
	}

	return res
}

func inverseModulo(A, B, C int) int {
	if B == 1 {
		return A % C
	}

	t := (A % C * A % C) % C
	if B%2 == 0 {
		return inverseModulo(t, B/2, C) % C
	}
	return ((A % C) * (inverseModulo(t, B/2, C) % C)) % C
}
