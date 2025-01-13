package timeComplexity

/*
	Problem : https://www.interviewbit.com/problems/word-count/
*/

// T(n) : O(n) , S(n) : O(1)
func WordCount(A string) int {

	i, n, cnt := 0, len(A), 0
	for i < n {
		for i < n && A[i] == ' ' {
			i++
		}

		e := false
		for i < n && A[i] != ' ' {
			e = true
			i++
		}

		if e {
			cnt++
		}
	}

	return cnt
}
