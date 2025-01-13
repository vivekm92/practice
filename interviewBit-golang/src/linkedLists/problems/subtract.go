package linkedList

import (
	"interviewBit/src/utils"
)

/*
	Problem : https://www.interviewbit.com/problems/subtract/
*/

// T(n) : O(n), S(n) :  O(n)
func Subtract(A *utils.ListNode) *utils.ListNode {

	count := 0
	var curr *utils.ListNode = A
	for curr != nil {
		curr = curr.Next
		count++
	}

	curr = A
	k := count / 2
	for k > 0 && curr != nil {
		k--
		curr = curr.Next
	}

	if count%2 == 1 {
		curr = curr.Next
	}

	var stack []*utils.ListNode
	for curr != nil {
		stack = append(stack, curr)
		curr = curr.Next
	}

	t := len(stack)
	curr = A
	for i := t - 1; i >= 0; i-- {
		curr.Value = stack[i].Value - curr.Value
		curr = curr.Next
	}

	return A
}
