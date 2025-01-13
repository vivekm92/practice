package stringProblems

/*
	problem : https://www.interviewbit.com/problems/longest-common-prefix/
*/

// T(n) : O(n), S(n) : O(n), where n is sum of length of all strings
func LongestCommonPrefix(A []string) string {

	n := len(A)
	if n == 0 {
		return ""
	}

	nC, res := len(A[0]), ""
	for j := 0; j < nC; j++ {
		for i := 1; i < n; i++ {
			if j >= len(A[i]) || A[0][j] != A[i][j] {
				return res
			}
		}
		res += string(A[0][j])
	}

	return res
}
