package arrayExamples

/*
	Problem : https://www.interviewbit.com/problems/max-non-negative-subarray/
*/

// MaxNonNegativeSubarray returns the contiguous subarray with the largest sum
// among all subarrays in A. If there are multiple subarrays with the largest
// sum, any one of them is returned. If no subarray exists, an empty slice is
// returned.
//
// T(n): O(n), S(n): O(n)
func MaxSet(A []int) []int {
	// Initialize variables.
	n, k := len(A), 0
	i1, i2 := -1, -1
	currSum, maxSum := 0, 0

	// Iterate through each element in A.
	for i := 0; i < n; i++ {
		// If the current element is non-negative, add it to the current sum
		// and increment the length of the current subarray.
		if A[i] >= 0 {
			currSum += A[i]
			k++
		} else {
			// If the current element is negative, set the current sum to 0
			// and reset the length of the current subarray to 0.
			currSum = 0
			k = 0
		}

		// If the current sum is the largest sum seen so far or there is a tie
		// and the current subarray is longer than the previous longest
		// subarray, update the longest subarray indices.
		if currSum > maxSum || currSum == maxSum && i2-i1+1 < k {
			maxSum = currSum
			i2 = i
			i1 = i2 - k + 1
		}
	}

	// Create a slice to store the subarray with the largest sum.
	res := make([]int, 0)

	// If there are no subarrays with non-negative elements, return an empty
	// slice.
	if i1 == -1 {
		return res
	}

	// Add the elements in the subarray with the largest sum to the result slice.
	for i := i1; i <= i2; i++ {
		res = append(res, A[i])
	}

	// Return the result slice.
	return res
}
