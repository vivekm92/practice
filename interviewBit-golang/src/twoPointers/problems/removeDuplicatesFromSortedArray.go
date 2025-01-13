package twoPointers

/*
problem : https://www.interviewbit.com/problems/remove-duplicates-from-sorted-array/
*/

// T(n) : O(n), S(n) : O(1)
func RemoveDuplicates(A []int) []int {

	n := len(A)
	if n == 0 {
		return []int{}
	} else if n == 1 {
		return A[:1]
	}

	cnt := 1
	i, j := 1, 0
	for i < n {
		for i < n && A[i] == A[i-1] {
			i++
		}

		if i < n {
			cnt++
			A[j+1] = A[i]
		}

		i++
		j++
	}

	return A[:cnt]
}
