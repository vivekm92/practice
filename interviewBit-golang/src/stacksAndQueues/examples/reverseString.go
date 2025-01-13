package stacksAndQueuesExamples

/*
	Problem : https://www.interviewbit.com/problems/reverse-string/

	Approach 0 : Reverse Iteration
	Approach 1 : Two Pointers, Swap Iteration
	Approach 2 : Using Stacks
*/

// T(n) : O(n), S(n) : O(n)
func ReverseString(A string) string {

	res, n := "", len(A)
	for i := n - 1; i >= 0; i-- {
		res += string(A[i])
	}

	return res
}

// T(n) : O(n), S(n) : O(1)
func ReverseString1(A string) string {

	n := len(A)
	for i := 0; i < n/2; i++ {
		t1, t2 := A[i], A[n-1-i]
		A = A[:i] + string(t2) + A[i+1:n-1-i] + string(t1) + A[n-i:]
	}

	return A
}

// T(n) : O(n), S(n) : O(n)
func ReverseString2(A string) string {

	stack := make([]rune, 0)
	for _, s := range A {
		stack = append(stack, s)
	}

	res := ""
	for len(stack) > 0 {
		res += string(stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return res
}
