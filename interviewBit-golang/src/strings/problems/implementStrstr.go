package stringProblems

/*
	Prroblem : https://www.interviewbit.com/problems/implement-strstr/
*/

// T(n) : O(n) , S(n) : O(1)
func StrStr(A string, B string) int {

	n, m := len(A), len(B)
	for i := 0; i < n-m+1; i++ {
		if A[i:i+m] == B {
			return i
		}
	}

	return -1
}
