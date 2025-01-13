package stringProblems

/*
	Problem : https://www.interviewbit.com/problems/maximum-substring/
*/

// T(n) : O(n) , S(n) : O(1)
func MaximumSubstring(A string, B int) int {

	n, i, maxCount := len(A), 0, 0
	for i < n {
		cnt, j := 0, 0
		for j < B && i+j < n {
			if A[i+j] == 'a' {
				cnt++
			}
			j++
		}
		if cnt > maxCount {
			maxCount = cnt
		}
		i += j
	}

	return maxCount
}
