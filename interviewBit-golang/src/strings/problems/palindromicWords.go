package stringProblems

/*
	Problem : https://www.interviewbit.com/problems/palindromic-words/
*/

func IsPalindromeString(A string) bool {

	n := len(A)
	for i := 0; i < n/2; i++ {
		if A[i] != A[n-1-i] {
			return false
		}
	}
	return true
}

// T(n) : O(n), S(n) : O(1)
func CountPalidromicWords(A string) int {

	n, count := len(A), 0
	startIdx, endIdx := 0, 0
	for i := 0; i < n; i++ {

		for i < n && A[i] == ' ' {
			i++
		}
		startIdx = i

		for i < n && A[i] != ' ' {
			i++
		}
		endIdx = i

		if IsPalindromeString(A[startIdx:endIdx]) {
			count++
		}

	}

	return count
}
