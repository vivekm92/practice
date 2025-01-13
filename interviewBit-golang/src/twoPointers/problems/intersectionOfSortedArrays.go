package twoPointers

/*
problem : https://www.interviewbit.com/problems/intersection-of-sorted-arrays/
*/

// T(n) : O(n), S(n) : O(n)
func Intersect(A []int, B []int) []int {

	res := make([]int, 0)
	i, j := 0, 0
	n, m := len(A), len(B)
	for i < n && j < m {
		if A[i] == B[j] {
			res = append(res, A[i])
			i++
			j++
		} else if A[i] < B[j] {
			i++
		} else {
			j++
		}
	}

	return res
}
