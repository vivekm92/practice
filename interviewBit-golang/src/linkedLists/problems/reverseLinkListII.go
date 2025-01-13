package linkedList

import (
	"interviewBit/src/utils"
)

/*
	Problem : https://www.interviewbit.com/problems/reverse-link-list-ii/
*/

// T(n) : O(n), S(n) : O(1)
func ReverseBetween(A *utils.ListNode, B, C int) *utils.ListNode {

	if B == C {
		return A
	}

	nB := B
	var prev, curr *utils.ListNode = nil, A
	for B > 1 && curr != nil {
		prev = curr
		curr = curr.Next
		B--
		C--
	}

	var pB, cB *utils.ListNode = prev, curr
	for C > 1 && curr != nil {
		var t *utils.ListNode = curr.Next
		curr.Next = prev
		prev = curr
		curr = t
		B--
		C--
	}
	t := curr.Next
	curr.Next = prev

	if nB == 1 {
		A = curr
	}

	cB.Next = t
	if pB != nil {
		pB.Next = curr
	}

	return A
}
