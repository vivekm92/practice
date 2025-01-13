package hashing

/*
	Problem : https://www.interviewbit.com/problems/subarray-with-b-odd-numbers/
*/

// Brute-Force Solution
// T(n) : O(n2) ; S(n) : O(1)
func SubarrayWithBOddNumbersBruteForce(A []int, B int) int {

	n, res := len(A), 0
	for i := 0; i < n; i++ {
		cnt := 0
		for j := i; j < n; j++ {
			if A[j]%2 != 0 {
				cnt++
			}

			if cnt == B {
				res++
			} else if cnt > B {
				break
			}
		}
	}

	return res
}

// Optimised-Solution
// T(n) : O(n) ; S(n) : O(n)
func SubarrayWithBOddNumbers(A []int, B int) int {

	currSum, res := 0, 0
	m := make(map[int]int, 0)
	m[currSum] = 1
	for _, v := range A {
		if v%2 != 0 {
			currSum++
		}

		if _val, ok := m[currSum-B]; ok {
			res += _val
		}
		m[currSum]++
	}

	return res
}
