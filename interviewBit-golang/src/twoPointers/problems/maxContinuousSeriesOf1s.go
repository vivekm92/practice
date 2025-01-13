package twoPointers

/*
	Problem : https://www.interviewbit.com/problems/max-continuous-series-of-1s/
*/

// T(n) : O(n), S(n) : O(1)
func MaxContinuos1s(A []int, B int) []int {

	i, n := 0, len(A)
	st1, currCount, maxCount := -1, 0, 0
	for i < n {
		st, j, ed := i, B, i
		for i < n && (A[i] == 1 || A[i] == 0 && j > 0) {
			if A[i] == 0 {
				j--
			}
			i++
		}
		ed = i
		currCount = ed - st
		if currCount > maxCount {
			st1 = st
			maxCount = currCount
		}
		i = st + 1
	}

	res := make([]int, 0)
	for i := 0; i < maxCount; i++ {
		res = append(res, st1)
		st1++
	}

	return res
}
