package twoPointers

/*
problem : https://www.interviewbit.com/problems/merge-two-sorted-lists-ii/
*/

// T(n): O(n), S(n): O(n)
func Merge(A []int, B []int) []int {

	n, m := len(A), len(B)
	c := make([]int, 0)
	i, j := 0, 0
	for i < n || j < m {
		if i < n && j < m {
			if A[i] < B[j] {
				c = append(c, A[i])
				i++
			} else {
				c = append(c, B[j])
				j++
			}
		} else if i < n {
			c = append(c, A[i])
			i++
		} else {
			c = append(c, B[j])
			j++
		}
	}
	return c
}
