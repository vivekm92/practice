package stacksAndQueuesExamples

/*
	Problem : https://www.interviewbit.com/problems/generate-all-parentheses/
*/

// T(n) : O(n), S(n) : O(n)
func IsValidParentheses(A string) int {

	n := len(A)
	lookup := make(map[rune]rune, 0)
	lookup[')'], lookup['}'], lookup[']'] = '(', '{', '['
	stack := make([]rune, 0)
	for i := 0; i < n; i++ {
		if len(stack) == 0 || A[i] == '(' || A[i] == '{' || A[i] == '[' {
			stack = append(stack, rune(A[i]))
		} else if len(stack) != 0 && lookup[rune(A[i])] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			return 0
		}
	}

	if len(stack) == 0 {
		return 1
	}
	return 0
}
