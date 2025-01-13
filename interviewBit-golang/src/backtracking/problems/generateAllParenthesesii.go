package backtracking

/*
	Problem : https://www.interviewbit.com/problems/generate-all-parentheses-ii/
*/

func solveGenerateAllParenthesesii(A int, B int, C int, D string, E *[]string) {

	if A == B && A == C {
		*E = append(*E, D)
		return
	}

	if A < C {
		solveGenerateAllParenthesesii(A+1, B, C, D+"(", E)
	}

	if B < A {
		solveGenerateAllParenthesesii(A, B+1, C, D+")", E)
	}
}

// T(n) : O() ; S(n) : O()
func GenerateAllParenthesesii(A int) []string {
	res := make([]string, 0)
	if A == 0 {
		return res
	}

	solveGenerateAllParenthesesii(0, 0, A, "", &res)
	return res
}
