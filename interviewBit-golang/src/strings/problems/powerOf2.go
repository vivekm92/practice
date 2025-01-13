package stringProblems

import "strconv"

/*
	Problem : 	https://www.interviewbit.com/problems/power-of-2/
*/

// T(n) : O(nlogn) , S(n) : O(n)
func PowerOf2(A string) int {

	k := 0
	for A != "1" {
		if _A, ok := DivideBy2(A); ok {
			k++
			A = _A
		} else {
			return 0
		}
	}

	if k >= 1 {
		return 1
	}

	return 0
}

func DivideBy2(A string) (string, bool) {

	n := len(A)
	v1 := int(A[n-1] - '0')
	if n == 1 && v1 == 0 {
		return "", false
	} else if v1%2 != 0 {
		return "", false
	}

	res, i, carry := "", 0, 0
	for i < n {
		v1 := int(A[i] - '0')
		v1 = 10*carry + v1

		res += strconv.Itoa(v1 / 2)
		carry = v1 % 2
		i++
	}

	i = 0
	for i < len(res) && res[i] == '0' {
		i++
	}

	return res[i:], true
}
