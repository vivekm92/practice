package linkedList

import (
	"interviewBit/src/utils"
)

/*
	Problem : https://www.interviewbit.com/problems/partition-list/
*/

// T(n): O(n), S(n) : O(n)
func Partition(A *utils.ListNode, B int) *utils.ListNode {

	var dummyLeftHead *utils.ListNode = utils.ListNode_new(-1)
	var leftHead *utils.ListNode = dummyLeftHead

	var dummyRightHead *utils.ListNode = utils.ListNode_new(-1)
	var rightHead *utils.ListNode = dummyRightHead

	var currNode *utils.ListNode = A
	for currNode != nil {
		if currNode.Value < B {
			leftHead.Next = utils.ListNode_new(currNode.Value)
			leftHead = leftHead.Next
		} else {
			rightHead.Next = utils.ListNode_new(currNode.Value)
			rightHead = rightHead.Next
		}

		currNode = currNode.Next
	}

	leftHead.Next = dummyRightHead.Next
	return dummyLeftHead.Next
}
