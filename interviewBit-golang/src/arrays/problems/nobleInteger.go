package arrayProblems

import "sort"

// T(n) : O(nlog(n)) , S(n) : O(1)
func NobleIntegers(A []int) int {

	sort.Slice(A, func(i, j int) bool {
		return A[i] <= A[j]
	})

	n := len(A)
	if A[n-1] == 0 {
		return 1
	}

	for i := 0; i < n-1; i++ {
		if A[i] == n-i-1 && A[i] != A[i+1] {
			return 1
		}
	}

	return -1
}

func NobleInteger(A []int) int {

	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})

	n := len(A)
	for i := 0; i < n; i++ {
		if i == n-1 && A[i] == 0 {
			return 1
		} else if i < n-1 && A[i] != A[i+1] {
			if A[i] == n-1-i {
				return 1
			}
		}
	}

	return -1
}
