package timeComplexity

import "strconv"

/*
	Problem : https://www.interviewbit.com/problems/extracting-numbers/
*/

// T(n) : O(n), S(n) : O(1)
func ExtractNumbers(A string) int64 {

	t, res := "", 0
	for _, v := range A {
		if v >= '0' && v <= '9' {
			t += string(v)
		} else {
			value, _ := strconv.Atoi(t)
			res += value
			t = ""
		}
	}

	if len(t) > 0 {
		value, _ := strconv.Atoi(t)
		res += value
	}

	return int64(res)
}
