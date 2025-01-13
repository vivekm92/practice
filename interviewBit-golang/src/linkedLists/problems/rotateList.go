package linkedList

import (
	"interviewBit/src/utils"
)

/*
	Problem : https://www.interviewbit.com/problems/rotate-list/
*/

// T(n) : O(n), S(n) : O(1)
func RotateRight(A *utils.ListNode, B int) *utils.ListNode {

	n := 0
	var curr *utils.ListNode = A
	for curr != nil {
		curr = curr.Next
		n++
	}

	B = B % n
	if B == 0 {
		return A
	}

	curr = A
	k := n - B
	var prev *utils.ListNode = nil
	for k > 0 {
		prev = curr
		curr = curr.Next
		k--
	}

	var last *utils.ListNode = curr
	for last.Next != nil {
		last = last.Next
	}

	if prev != nil {
		prev.Next = nil
	}
	last.Next = A

	return curr
}
