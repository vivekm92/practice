package arrayProblems

/*
	Problem : https://www.interviewbit.com/problems/leaders-in-an-array/
*/

// T(n) : O(n), S(n) : O(1)
func FindLeaders(A []int) []int {

	n, res := len(A), make([]int, 0)
	suffixMax := A[n-1]
	res = append(res, suffixMax)
	for i := n - 2; i >= 0; i-- {
		if A[i] > suffixMax {
			res = append(res, A[i])
			suffixMax = A[i]
		}
	}

	return res
}
