package twoPointers

/*
	Problem : https://www.interviewbit.com/problems/counting-subarrays/
*/

// T(n) : O(n), S(n) : O(1)
func CountSubArrays(A []int, B int) int {

	n, st, ed := len(A), 0, 0
	count, currSum := 0, A[0]
	for ed < n {
		if currSum < B {
			ed++
			count += (ed - st)
			if ed < n {
				currSum += A[ed]
			}
		} else {
			currSum -= A[st]
			st++
		}
	}
	return count
}
