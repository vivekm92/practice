package mathProblems

/*
	Problem : https://www.interviewbit.com/problems/digital-root/
*/

func getSum(A int) int {
	s := 0
	for A > 0 {
		s += (A % 10)
		A /= 10
	}
	return s
}

// T(n) : O(logn), S(n) : O(1)
func DigitalRoot(A int) int {

	for A/10 > 0 {
		A = getSum(A)
	}
	return A
}
