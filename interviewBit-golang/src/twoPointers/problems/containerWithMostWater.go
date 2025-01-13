package twoPointers

/*
problem : https://www.interviewbit.com/problems/container-with-most-water/
*/

// T(n) : O(n); S(n) : O(1)
func MaxArea(A []int) int {

	n, res := len(A), 0
	i, j := 0, n-1
	for i < j {
		curr := 0
		if A[i] < A[j] {
			curr += A[i] * (j - i)
			i++
		} else {
			curr += A[j] * (j - i)
			j--
		}

		if res < curr {
			res = curr
		}
	}

	return res
}
