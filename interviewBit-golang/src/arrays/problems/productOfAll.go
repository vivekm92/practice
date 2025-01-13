package arrayProblems

/*
	Problem : https://www.interviewbit.com/problems/product-of-all/
*/

// T(n) : O(n) , S(n) : O(n)
func ProductOfAll(A []int) []int {

	const MOD = 1000000007
	n := len(A)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			arr[i] = A[i] % MOD
		} else {
			arr[i] = ((arr[i-1] % MOD) * (A[i] % MOD)) % MOD
		}
	}

	brr := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			brr[i] = A[i] % MOD
		} else {
			brr[i] = ((brr[i+1] % MOD) * (A[i] % MOD)) % MOD
		}
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			res[i] = brr[i+1] % MOD
		} else if i == n-1 {
			res[i] = arr[i-1] % MOD
		} else {
			res[i] = ((arr[i-1] % MOD) * (brr[i+1] % MOD)) % MOD
		}
	}

	return res
}
