package hashing

/*
	Problem : https://www.interviewbit.com/problems/an-increment-problem/
*/

// Brute Force Solution
// T(n) : O(n2) ; S(n) : O(1)
func AnIncrementProblem(A []int) []int {

	for i := 0; i < len(A); i++ {
		for j := 0; j < i; j++ {
			if A[i] == A[j] {
				A[j] += 1
				break
			}
		}
	}

	return A
}

// Optimised Solution
// T(n) : O() ; S(n) : O()
// func AnIncrementProblem1(A []int) []int {

// 	return A
// }
