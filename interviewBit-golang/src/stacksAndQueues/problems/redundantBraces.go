package stacksAndQueues

/*
	Problem : https://www.interviewbit.com/problems/redundant-braces/
*/

// T(n) : O(n) , S(n) : O(n)
func Braces(A string) int {

	n := len(A)
	stack := make([]rune, 0)
	for i := 0; i < n; i++ {
		if A[i] == '(' || A[i] == '+' || A[i] == '-' || A[i] == '*' || A[i] == '/' {
			stack = append(stack, rune(A[i]))
		} else if A[i] == ')' {
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				return 1
			}
			for len(stack) > 0 && stack[len(stack)-1] != '(' {
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		}
	}

	return 0
}
