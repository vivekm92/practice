package stacksAndQueues

/*
	problem : https://www.interviewbit.com/problems/nearest-smaller-element/
*/

// T(n) : O(n), S(n) : O(n)
func PrevSmaller(A []int) []int {

	res, stack := make([]int, 0), make([]int, 0)
	for _, v := range A {
		if len(stack) == 0 {
			res = append(res, -1)
			stack = append(stack, v)
		} else {
			for len(stack) > 0 && stack[len(stack)-1] >= v {
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 {
				res = append(res, -1)
				stack = append(stack, v)
			} else {
				res = append(res, stack[len(stack)-1])
				stack = append(stack, v)
			}
		}
	}

	return res
}
