package stringProblems

import "strconv"

/*
	Problem : https://www.interviewbit.com/problems/add-binary-strings/
*/

// T(n) : O(n), S(n) : O(n)
func AddBinary(A string, B string) string {

	n, m, res, carry := len(A), len(B), "", 0
	for n > 0 || m > 0 {
		var v1, v2 int
		if n > 0 {
			v1, _ = strconv.Atoi(string(A[n-1]))
		}

		if m > 0 {
			v2, _ = strconv.Atoi(string(B[m-1]))
		}

		sum := v1 + v2 + carry
		res += strconv.Itoa(sum % 2)
		carry = sum / 2

		n -= 1
		m -= 1
	}

	if carry > 0 {
		res += "1"
	}

	r := []rune(res)
	n = len(r)
	for i := 0; i < n/2; i++ {
		r[i], r[n-1-i] = r[n-1-i], r[i]
	}

	return string(r)
}
