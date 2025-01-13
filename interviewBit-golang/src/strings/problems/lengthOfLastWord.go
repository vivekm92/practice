package stringProblems

import (
	"strings"
)

/*
	problem : https://www.interviewbit.com/problems/length-of-last-word/
*/

// T(n) : O(n), S(n) : O(1)
func LengthOfLastWord(A string) int {
	A = strings.Trim(A, " ")
	v := strings.Split(A, " ")

	n := len(v)
	return len(v[n-1])
}

// T(n) : O(n), S(n) : O(1)
func LengthOfLastWord1(A string) int {

	n := len(A)
	j := n - 1
	for j >= 0 && A[j] == ' ' {
		j--
	}

	cnt := 0
	for j >= 0 && A[j] != ' ' {
		cnt++
		j--
	}

	return cnt
}
