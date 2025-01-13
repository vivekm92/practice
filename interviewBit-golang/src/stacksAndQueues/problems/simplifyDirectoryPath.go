package stacksAndQueues

import (
	"strings"
)

/*
	Problem : https://www.interviewbit.com/problems/simplify-directory-path/
*/

// T(n) : O(n) , S(n) : O(n)
func SimplifyDirectoryPath(A string) string {

	t := strings.Split(A, "/")
	res := make([]string, 0)
	for _, v := range t {
		v = strings.Trim(v, " ")
		if v == "." || v == "" {
			continue
		} else if v == ".." {
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
		} else {
			res = append(res, v)
		}
	}

	return "/" + strings.Join(res, "/")
}
