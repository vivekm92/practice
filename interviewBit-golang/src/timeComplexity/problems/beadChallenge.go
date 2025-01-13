package timeComplexity

import "math"

// T(n) : O(nlogn) , S(n) : O(n)
func BeadChallenge(A []int) int {

	t := math.MinInt32
	for _, a := range A {
		if a > t {
			t = a
		}
	}

	t = int(math.Sqrt(float64(t))) + 2
	sieve := make([]bool, t)
	sieve[0] = true
	sieve[1] = true
	for i := 2; i*i < t; i++ {
		for j := i * i; j < t; j += i {
			if !sieve[j] {
				sieve[j] = true
			}
		}
	}

	lookup := make(map[int]int)
	for _, a := range A {
		lookup = getPrimeFactors(a, sieve, lookup)
	}

	n := len(A)
	for _, v := range lookup {
		if v%n != 0 {
			return 0
		}
	}

	return 1
}

func getPrimeFactors(A int, sieve []bool, lookup map[int]int) map[int]int {

	for i := 2; i < len(sieve); i++ {
		if !sieve[i] && A%i == 0 {
			for A%i == 0 {
				lookup[i]++
				A /= i
			}
		}
	}

	if A > 1 {
		lookup[A]++
	}

	return lookup
}
