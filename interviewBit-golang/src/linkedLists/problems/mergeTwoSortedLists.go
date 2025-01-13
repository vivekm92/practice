package linkedList

import (
	"interviewBit/src/utils"
)

/*
	Problem : https://www.interviewbit.com/problems/merge-two-sorted-lists/
*/

// T(n) : O(n), S(n) : O(1)
func MergeTwoLists(A *utils.ListNode, B *utils.ListNode) *utils.ListNode {
	var dummyHead *utils.ListNode = utils.ListNode_new(-1)
	var curr *utils.ListNode = dummyHead
	for A != nil && B != nil {
		if A.Value < B.Value {
			curr.Next = A
			A = A.Next
		} else {
			curr.Next = B
			B = B.Next
		}
		curr = curr.Next
	}
	if A == nil {
		curr.Next = B
	} else {
		curr.Next = A
	}

	return dummyHead.Next
}
