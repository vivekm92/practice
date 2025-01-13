package stringProblems

/*
	problem : https://www.interviewbit.com/problems/salutes/
*/

// T(n) : O(n), S(n) : O(n)
func CountSalutes(A string) int64 {

	n := len(A)
	t := make([]int64, n)
	if A[n-1] == '<' {
		t[n-1] = 1
	}
	for i := n - 2; i >= 0; i-- {
		if A[i] == '<' {
			t[i] = t[i+1] + 1
		} else {
			t[i] = t[i+1]
		}
	}

	var nClaps int64
	for i := 0; i < n; i++ {
		if A[i] == '>' {
			nClaps += t[i]
		}
	}

	return nClaps
}
