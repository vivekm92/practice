package stringProblems

/*
	problem : https://www.interviewbit.com/problems/minimum-parantheses/
*/

// T(n) : O(n), S(n) : O(1)
func MinimumParantheses(A string) int {

	cnt, n, res := 0, len(A), 0
	for i := 0; i < n; i++ {

		if A[i] == '(' {
			cnt++
		} else {
			cnt--
		}

		if cnt < 0 {
			res++
			cnt = 0
		}
	}

	return cnt + res
}
