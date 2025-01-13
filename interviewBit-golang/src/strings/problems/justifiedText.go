package stringProblems

/*
	Problem : https://www.interviewbit.com/problems/justified-text/
*/

// T(n) : O(n2), S(n) : O(n)
func JustifiedText(A []string, B int) []string {

	res, n, currRow := make([]string, 0), len(A), ""
	for i := 0; i < n; i++ {
		if len(currRow) == 0 {
			currRow += A[i]
		} else if len(currRow) > B {
			currRow = ""
		}

	}
	return res
}
