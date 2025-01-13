package mathProblems

import (
	"math"
)

/*
	Problem : https://www.interviewbit.com/problems/prime-sum/
*/

func isPrime(n int) bool {
	k := int(math.Ceil(math.Sqrt(float64(n + 1))))
	for i := 2; i < k; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// refer : https://www.geeksforgeeks.org/how-is-the-time-complexity-of-sieve-of-eratosthenes-is-nloglogn/
// T(n) : O(NlogLogN), S(n) : O(NloglogN)
func PrimeSum(A int) []int {

	n := int(math.Ceil(math.Sqrt(float64(20000002))))
	sieve := make([]bool, n)
	for i := 2; i < n; i++ {
		for j := i * i; j < n; j += i {
			if !sieve[j] {
				sieve[j] = true
			}
		}
	}
	sieve[0], sieve[1] = true, true

	for i := 2; i < n; i++ {
		if !sieve[i] {
			if A-i <= n && !sieve[A-i] {
				return []int{i, A - i}
			} else if isPrime(A - i) {
				return []int{i, A - i}
			}
		}
	}

	return []int{}
}
