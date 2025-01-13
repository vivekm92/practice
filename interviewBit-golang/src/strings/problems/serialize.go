package stringProblems

/*
problem : https://www.interviewbit.com/problems/serialize/
*/

import (
	"strconv"
)

// T(n) : O(n), S(n) : O(1)
func Serialize(A []string) string {

	var res string
	for _, v := range A {
		res += v + strconv.Itoa(len(v)) + "~"
	}

	return res
}
