package mathProblems

/*
	Problem : https://www.interviewbit.com/problems/product-of-digits/
*/

// T(n) : O(logn), S(n) : O(1)
func ProductOfDigits(A int) int {

	if A == 0 {
		return A
	}
	res := 1
	for A > 0 {
		res *= (A % 10)
		A /= 10
	}
	return res
}
