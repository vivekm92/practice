package hashing

/*
	Problem : https://www.interviewbit.com/problems/longest-subarray-length/
*/

// Brute-Force Solution
// T(n) : O(n2) ; S(n) : O(1)
func LongestSubarrayLengthBruteForce(A []int) int {

	n, maxLen := len(A), 0
	for i := 0; i < n; i++ {
		cnt0, cnt1 := 0, 0
		for j := i; j < n; j++ {
			if A[j] == 0 {
				cnt0++
			} else {
				cnt1++
			}

			if cnt1-cnt0 == 1 {
				l := j - i + 1
				if l > maxLen {
					maxLen = l
				}
			}
		}
	}

	return maxLen
}

// Optimised Solution
// T(n) : O(n) ; S(n) : O(n)
func LongestSubarrayLength(A []int) int {

	maxLen, currSum, m := 0, 0, make(map[int]int, 0)
	for i, v := range A {
		if v == 0 {
			currSum--
		} else {
			currSum++
		}

		if currSum == 1 {
			maxLen = i + 1
		} else if _, ok := m[currSum]; !ok {
			m[currSum] = i
		}

		if idx, ok := m[currSum-1]; ok && maxLen < i-idx {
			maxLen = i - idx
		}
	}

	return maxLen
}
