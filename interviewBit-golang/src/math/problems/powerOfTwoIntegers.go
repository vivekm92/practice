package mathProblems

import (
	"math"
)

/*
	Problem : https://www.interviewbit.com/problems/power-of-two-integers/
*/

// T(n) : O(sqrt(n)), S(n) : O(logn)
func IsPower(A int) int {

	if A == 1 {
		return 1
	}

	n := int(math.Sqrt(float64(A))) + 1
	sieve := make([]bool, n+1)
	sieve[0], sieve[1] = true, true
	for i := 2; i <= n; i++ {
		for j := i * i; j <= n; j += i {
			if !sieve[j] {
				sieve[j] = true
			}
		}
	}

	B := A
	lookup := make(map[int]int, 0)
	for i, v := range sieve {
		if !v && A%i == 0 {
			for A%i == 0 {
				lookup[i]++
				A = A / i
			}
		}
	}
	if A != 1 {
		lookup[A] = 1
	}

	eventCount, oddCount := 0, 0
	for k, v := range lookup {
		if k == B || v == 1 { // return 0 if Prime
			return 0
		}
		if v%2 == 0 {
			eventCount++
		} else {
			oddCount++
		}
	}

	if eventCount == 0 && len(lookup) == 1 {
		return 1
	} else if eventCount > 0 && oddCount == 0 {
		return 1
	}

	return 0
}
