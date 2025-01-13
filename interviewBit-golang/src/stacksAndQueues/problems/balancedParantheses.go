package stacksAndQueues

/*
	problem : https://www.interviewbit.com/problems/balanced-parantheses/
*/

// T(n) : O(n), S(n) : O(1)
func BalancedParantheses(A string) int {

	n, cnt := len(A), 0
	for i := 0; i < n; i++ {
		if cnt < 0 {
			return 0
		}

		if A[i] == '(' {
			cnt++
		} else {
			cnt--
		}
	}

	if cnt == 0 {
		return 1
	}

	return 0
}

// T(n) : O(n), S(n) : O(n)
func BalancedParantheses1(A string) int {

	stack := make([]rune, 0)
	for _, val := range A {
		if val == '(' {
			stack = append(stack, val)
		} else {
			if len(stack) == 0 {
				return 0
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		return 1
	}

	return 0
}
