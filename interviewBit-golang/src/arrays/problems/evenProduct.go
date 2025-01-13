package arrayProblems

/*
	Problem : https://www.interviewbit.com/problems/even-product/
*/

// T(n) : O(logn), S(n) : O(1)
func EvenProduct(A []int) int {

	n := len(A)
	const MOD = 1000000007
	var t int64 = 1
	for i := 1; i <= n; i++ {
		t = ((t % MOD) * (2 % MOD)) % MOD
	}

	return int(t) - 1
}


