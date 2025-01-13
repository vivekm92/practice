package mathProblems

/*
	Problem : https://www.interviewbit.com/problems/last-digit-k-count/
*/

// T(n) : O(1) , S(n) : O(1)
func LastDigitKCount(A int, B int, C int) int {

	if A%10 > C {
		A += 10 - ((A % 10) - C)
	} else if C > A%10 {
		A += (C - (A % 10))
	}

	if B%10 > C {
		B -= ((B % 10) - C)
	} else if C > A%10 {
		B -= 10 - (C - (B % 10))
	}

	if A > B {
		return 0
	}

	return (B-A)/10 + 1
}
