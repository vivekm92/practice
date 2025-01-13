package linkedListExamples

import (
	"interviewBit/src/utils"
)

/*
	Problem : https://www.interviewbit.com/problems/reverse-linked-list/
*/

func reverse(prev *utils.ListNode, curr *utils.ListNode, currNext *utils.ListNode) *utils.ListNode {

	if currNext == nil {
		return curr
	}

	var temp *utils.ListNode = currNext.Next
	currNext.Next = curr
	curr.Next = prev

	return reverse(curr, currNext, temp)
}

// T(n) : O(n), S(n) : O(n)
func ReverseList(A *utils.ListNode) *utils.ListNode {
	return reverse(nil, A, A.Next)
}

// T(n) : O(n), S(n) : O(1)
func ReverseList1(A *utils.ListNode) *utils.ListNode {

	var prev *utils.ListNode = nil
	var curr *utils.ListNode = A
	var currNext *utils.ListNode = A.Next

	for currNext != nil {
		var t *utils.ListNode = currNext.Next
		currNext.Next = curr
		curr.Next = prev

		prev = curr
		curr = currNext
		currNext = t
	}

	return curr
}
