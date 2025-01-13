package stringProblems

/*
	Problem : https://www.interviewbit.com/problems/flipping-string/
*/

// T(n) : O(n) , S(n) : O(1)
func FlippingString(A string) int {

	i, n, maxCount := 0, len(A), 0
	for i < n {
		var count int
		for i < n && A[i] == '1' {
			count++
			i++
		}

		for i < n && A[i] == '0' {
			count++
			i++
		}

		start := i
		for i < n && A[i] == '1' {
			count++
			i++
		}

		if maxCount < count {
			maxCount = count
		}
		i = start
	}

	return maxCount
}
