package stringProblems

/*
	problem : https://www.interviewbit.com/problems/amazing-subarrays/
*/

// T(n) : O(n), S(n) : O(1)
func AmazingSubarrays(A string) int {

	n, res := len(A), 0
	for i := 0; i < n; i++ {
		v := A[i]
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' || v == 'A' || v == 'E' || v == 'I' || v == 'O' || v == 'U' {
			res += n - i
			res = res % 10003
		}
	}

	return res % 10003
}
