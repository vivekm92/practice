package mathExamples

/*
	problem : https://www.interviewbit.com/problems/prime-numbers/
*/

import (
	"math"
)

func isPrime(n int) bool {

	s := int(math.Sqrt(float64(n)))
	for i := 2; i <= s; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

// T(n) : O(nSqrt(n)), S(n) : O(n)
func PrimeNumbers(A int) []int {

	res := make([]int, 0)
	for i := 2; i <= A; i++ {
		if isPrime(i) {
			res = append(res, i)
		}
	}

	return res
}

// T(n) : O(nloglogn), S(n) : O(n)
func Sieve(A int) []int {

	arr, s := make([]bool, A+1), int(math.Sqrt(float64(A)))
	for i := 2; i <= s; i++ {
		for j := 2 * i; j <= A; j += i {
			arr[j] = true
		}
	}

	res := make([]int, 0)
	for i := 2; i <= A; i++ {
		if !arr[i] {
			res = append(res, i)
		}
	}

	return res
}
