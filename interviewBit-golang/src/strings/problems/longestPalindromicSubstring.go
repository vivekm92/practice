package stringProblems

/*
	Problem : https://www.interviewbit.com/problems/longest-palindromic-substring/
*/

func checkForPalindrome(A string, left, right int) string {
	n, i, j := len(A), left, right
	for i >= 0 && j < n && A[i] == A[j] {
		i--
		j++
	}
	return A[i+1 : j]
}

// T(n) : O(n2) , S(n) : O(1)
func LongestPalindromicSubstring(A string) string {

	n := len(A)
	var res string
	for i := 0; i < n; i++ {
		odd := checkForPalindrome(A, i, i)
		if len(odd) > len(res) {
			res = odd
		}
		even := checkForPalindrome(A, i, i+1)
		if len(even) > len(res) {
			res = even
		}
	}
	return res
}
