package stacksAndQueues

/*
	Problem : https://www.interviewbit.com/problems/nextgreater/
*/

// T(n) : O(n), S(n) : O(n)
func NextGreater(A []int) []int {

	n := len(A)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = -1
	}

	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		if len(stack) == 0 || A[stack[len(stack)-1]] >= A[i] {
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && A[stack[len(stack)-1]] < A[i] {
				res[stack[len(stack)-1]] = A[i]
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, i)
		}
	}

	return res
}
