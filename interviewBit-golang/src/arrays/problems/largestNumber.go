package arrayProblems

import (
	"sort"
	"strconv"
)

/*
	Problem : https://www.interviewbit.com/problems/largest-number/
	LeetCode : https://leetcode.com/problems/largest-number/
*/

// T(n) : O(nlogn), S(n) : O(n)
func LargestNumber(A []int) string {
	sort.Slice(A, func(i, j int) bool {
		return strconv.Itoa(A[i])+strconv.Itoa(A[j]) > strconv.Itoa(A[j])+strconv.Itoa(A[i])
	})

	if A[0] == 0 {
		return "0"
	}

	res := ""
	for _, v := range A {
		res += strconv.Itoa(v)
	}
	return res
}

func LargestNumber1(A []int) string {

	n := len(A)
	for i := 0; i < n; i++ {
		
	}

	res := ""
	return res
}

/*
Approach 0:

1.

*/
