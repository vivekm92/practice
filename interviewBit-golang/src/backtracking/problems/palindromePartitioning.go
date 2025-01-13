package backtracking

/*
	Problem : https://www.interviewbit.com/problems/palindrome-partitioning/
*/

func isPalindrome(A string, B, C int) bool {
	for B < C && A[B] == A[C] {
		B++
		C--
	}
	return B >= C
}

func solvePalindromePartitioning(A string, B []string, C *[][]string, D int) {
	if D >= len(A) {
		E := make([]string, len(B))
		copy(E, B)
		*C = append(*C, E)
		return
	}

	for i := D; i < len(A); i++ {
		if isPalindrome(A, D, i) {
			B = append(B, A[D:i+1])
			solvePalindromePartitioning(A, B, C, i+1)
			B = B[:len(B)-1]
		}
	}
}

// T(n) : O() ; S(n) : O()
func PalindromePartitioning(A string) [][]string {

	temp := make([]string, 0)
	res := make([][]string, 0)
	solvePalindromePartitioning(A, temp, &res, 0)
	return res
}
