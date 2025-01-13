package backtracking

/*
	Problem : https://www.interviewbit.com/problems/maximal-string/
*/

func swap(A string, B int, C int) string {
	b := string(A[B])
	c := string(A[C])
	return A[:B] + c + A[B+1:C] + b + A[C+1:]
}

func solveMaximalString(A string, B int, C *string) {

	if B == 0 {
		if A > *C {
			*C = A
		}
		return
	}

	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			A = swap(A, i, j)
			solveMaximalString(A, B-1, C)
			A = swap(A, i, j)
		}
	}
}

// T(n) : O() ; S(n) : O()
func MaximalString(A string, B int) string {
	C := ""
	solveMaximalString(A, B, &C)
	return C
}
