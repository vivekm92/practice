package mathProblems

/*
	Problem : https://www.interviewbit.com/problems/largest-coprime-divisor/
*/

// T(n) : O(logn), S(n) : O(1)
func LargestCoprimeDivisor(A, B int) int {

	X := A
	g := gcd(X, B)
	for g != 1 {
		X = X / g
		g = gcd(X, B)
	}
	return X
}
