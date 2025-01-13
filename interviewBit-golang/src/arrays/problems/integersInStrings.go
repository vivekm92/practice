package arrayProblems

import (
	"strconv"
	"strings"
)

/*
	problem : https://www.interviewbit.com/problems/integers-in-strings/
*/

// T(n) : O(n), S(n) : O(n)
func IntegersInStrings(A string) []int {

	t := strings.Split(A, ",")
	rs := make([]int, 0)
	for _, v := range t {
		if v1, err := strconv.Atoi(v); err == nil {
			rs = append(rs, v1)
		}
	}

	return rs
}
