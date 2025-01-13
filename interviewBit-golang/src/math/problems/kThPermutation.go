package mathProblems

/*
	Problem : https://www.interviewbit.com/problems/k-th-permutation
*/

// T(n) : O(n) ; S(n) : O(n)
func KThPermutation(A int, B int64) []int {

	res := make([]int, A)
	for i := 1; i <= A; i++ {
		res[i-1] = i
	}

	fact := make([]int64, 21)
	fact[0] = 1
	for i := 1; i < 21; i++ {
		fact[i] = fact[i-1] * int64(i)
	}

	var idx int
	if A > 20 {
		idx = A - 20 - 1
	}

	var totalIterationsLeft int64 = B
	var currentElmentPos int = 0
	for idx < A {
		currentElmentPos, totalIterationsLeft = div(totalIterationsLeft, fact[A-idx-1])
		if currentElmentPos == 0 {
			idx++
			continue
		}

		val := res[idx+currentElmentPos-1]
		for j := idx + currentElmentPos - 1; j > idx; j-- {
			res[j] = res[j-1]
		}
		res[idx] = val
		idx++
	}

	return res
}

func div(divident, divisor int64) (int, int64) {
	if divident < divisor {
		return 0, divident
	}
	if divident%divisor == 0 {
		return int(divident / divisor), divisor
	}
	return int(divident/divisor) + 1, divident % divisor
}
