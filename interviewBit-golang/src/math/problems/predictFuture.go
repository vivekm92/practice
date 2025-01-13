package mathProblems

/*
	Problem : https://www.interviewbit.com/problems/predict-future/
*/

const MOD = 1000000007

func findModPow(A int) int {
	if A <= 1 {
		return 5
	}

	v := findModPow(A/2) % MOD
	if A%2 == 0 {
		return ((v % MOD) * (v % MOD)) % MOD
	}
	return ((5 % MOD) * (v % MOD) * (v % MOD)) % MOD
}

// O(n) : O(logn), S(n) : O(1)
func PredictFuture(A int, B int, C int) []int {

	if C == 0 {
		return []int{A, B}
	} else if C == 1 {
		return []int{B - 2*A, 2*B + A}
	}

	v := findModPow(C / 2)
	if C%2 == 0 {
		t := PredictFuture(A, B, 0)
		v1, v2 := (t[0]%MOD*v%MOD)%MOD, (t[1]%MOD*v%MOD)%MOD
		if v1 < 0 {
			v1 += MOD
		}
		if v2 < 0 {
			v2 += MOD
		}
		return []int{v1, v2}
	}

	t := PredictFuture(A, B, 1)
	v1, v2 := (t[0]%MOD*v%MOD)%MOD, (t[1]%MOD*v%MOD)%MOD
	if v1 < 0 {
		v1 += MOD
	}
	if v2 < 0 {
		v2 += MOD
	}
	return []int{v1, v2}
}
