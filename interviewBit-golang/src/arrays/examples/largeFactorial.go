package arrayExamples

/*
	problem : https://www.interviewbit.com/problems/large-factorial/
*/

import (
	"strconv"
)

func Multiply(n int, t []int) []int {

	var carry int
	for i := 0; i < len(t); i++ {
		p := t[i]*n + carry
		t[i] = p % 10
		carry = p / 10
	}

	for carry > 0 {
		t = append(t, carry%10)
		carry = carry / 10
	}

	return t
}

// T(n) : O(N*K), S(n) : O(K), where k is number of digits in N!
func LargeFactorial(A int) string {

	t := make([]int, 0)
	t = append(t, 1)
	for i := 2; i <= A; i++ {
		t = Multiply(i, t)
	}

	var res string
	for i := len(t) - 1; i >= 0; i-- {
		res += strconv.Itoa(t[i])
	}

	return res
}
