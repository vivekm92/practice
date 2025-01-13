package twoPointers

/*
	Problem : https://www.interviewbit.com/problems/lexicographically-largest-array/
*/

// T(n) : O(n) ; S(n) : O(1)
func LexicographicallyLargestArray(A []int) []int {

	i, n := 0, len(A)
	for i < n && A[i] > A[n-1] {
		i++
	}

	res := make([]int, 0)
	for k := 0; k < i; k++ {
		res = append(res, A[k])
	}

	for k := n - 1; k >= i; k-- {
		res = append(res, A[k])
	}

	return res
}
