package mathProblems

/*
	Problem : https://www.interviewbit.com/problems/rearrange-array/
*/

// T(n) : O(n), S(n) : O(1)
func Arrange(A []int) []int {

	n := len(A)
	for i, v := range A {
		A[i] = v + (A[v]%n)*n
	}

	for i := range A {
		A[i] = A[i] / n
	}

	return A
}
