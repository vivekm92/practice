package stacksAndQueues

import (
	"strconv"
)

/*
	problem : https://www.interviewbit.com/problems/evaluate-expression/
*/

func isSymbol(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/"
}

// T(n) : O(n), S(n) : On()
func EvalRPN(A []string) int {

	n := len(A)
	st := make([]int, 0)
	for i := 0; i < n; i++ {
		if !isSymbol(A[i]) {
			if v, err := strconv.Atoi(A[i]); err == nil {
				st = append(st, v)
			}
		} else {
			v1, v2 := st[len(st)-1], st[len(st)-2]
			st = st[:len(st)-2]
			if A[i] == "+" {
				st = append(st, v2+v1)
			} else if A[i] == "-" {
				st = append(st, v2-v1)
			} else if A[i] == "*" {
				st = append(st, v2*v1)
			} else {
				st = append(st, v2/v1)
			}
		}
	}

	return st[0]
}
