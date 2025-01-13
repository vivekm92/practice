package mathProblems

/*
	Problem : https://www.interviewbit.com/problems/lowest-common-multiple-lcm/
*/

func gcd(A int, B int) int {
	if A < B {
		return gcd(B, A)
	} else if B == 0 {
		return A
	}
	return gcd(B, A%B)
}

// T(n) : O(logn); S(n) : O(1)
func Lcm(A int, B int) int64 {
	g := gcd(A, B)
	return int64(A) * int64(B) / int64(g)
}
