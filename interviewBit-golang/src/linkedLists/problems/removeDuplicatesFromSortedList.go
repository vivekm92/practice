package linkedList

import (
	"interviewBit/src/utils"
)

/*
	Problem : https://www.interviewbit.com/problems/remove-duplicates-from-sorted-list/
*/

// T(n) : O(n), S(n) : O(1)
func DeleteDuplicates(A *utils.ListNode) *utils.ListNode {

	if A == nil {
		return A
	}

	var currNode *utils.ListNode = A
	var nextNode *utils.ListNode = A.Next
	var head *utils.ListNode = A
	for nextNode != nil {
		for nextNode != nil && currNode.Value == nextNode.Value {
			nextNode = nextNode.Next
		}
		currNode.Next = nextNode
		if nextNode != nil {
			nextNode = nextNode.Next
		}
		currNode = currNode.Next
	}

	return head
}
