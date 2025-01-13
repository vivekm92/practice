package mathProblems

import (
	"sort"
)

/*
	Problem : https://www.interviewbit.com/problems/city-tour/
*/

// T(n) : O(nlogn) ; S(n) : O(n)
func CityTour(A int, B []int) int {

	const MOD = 1000000007
	fact := make([]int, A+1)
	fact[0] = 1
	for i := 1; i <= A; i++ {
		fact[i] = (i % MOD * fact[i-1] % MOD) % MOD
	}

	sort.Slice(B, func(i, j int) bool {
		return B[i] < B[j]
	})

	uv, n := make([]int, 0), len(B)
	uv = append(uv, B[0]-1)
	for i := 1; i < n; i++ {
		uv = append(uv, B[i]-B[i-1]-1)
	}
	uv = append(uv, A-B[n-1])

	k := A - n
	n1, p, d := len(uv), fact[k], 1
	for i := 0; i < n1; i++ {
		if i != 0 && i != n1-1 {
			p = (p % MOD * pow(2, uv[i]-1, MOD) % MOD) % MOD
		}
		d = (d % MOD * fact[uv[i]] % MOD) % MOD
	}
	res := (p % MOD * pow(d, MOD-2, MOD) % MOD) % MOD

	return res
}

func pow(A, B, C int) int {
	if B < 1 {
		return 1
	} else if B == 1 {
		return A % C
	}
	t := (A % C * A % C) % C
	if B%2 == 0 {
		return pow(t, B/2, C) % C
	}
	return (A % C * pow(t, B/2, C) % C) % C
}
