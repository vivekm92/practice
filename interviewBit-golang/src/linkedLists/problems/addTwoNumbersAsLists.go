package linkedList

import (
	"interviewBit/src/utils"
)

/*
	Problem : https://www.interviewbit.com/problems/add-two-numbers-as-lists/
*/

// T(n) : O(n), S(n) : O(n)
func AddTwoNumbers(A *utils.ListNode, B *utils.ListNode) *utils.ListNode {

	if A == nil {
		return B
	} else if B == nil {
		return A
	}

	var carry int
	var currNode *utils.ListNode = utils.ListNode_new(-1)
	var head *utils.ListNode = currNode
	for A != nil && B != nil {
		var sum int = A.Value + B.Value + carry
		currNode.Next = utils.ListNode_new(sum % 10)
		carry = sum / 10

		currNode = currNode.Next
		A = A.Next
		B = B.Next
	}

	for A != nil {
		var sum int = A.Value + carry
		currNode.Next = utils.ListNode_new(sum % 10)
		carry = sum / 10
		currNode = currNode.Next
		A = A.Next
	}

	for B != nil {
		var sum int = B.Value + carry
		currNode.Next = utils.ListNode_new(sum % 10)
		carry = sum / 10
		currNode = currNode.Next
		B = B.Next
	}

	if carry != 0 {
		currNode.Next = utils.ListNode_new(carry)
	}

	return head.Next
}
