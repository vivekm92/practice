package linkedList

import (
	"interviewBit/src/utils"
)

/*
	Problem : https://www.interviewbit.com/problems/replace-with-factors/
*/

// T(n) : O(n), S(n) : O(1)
func ReplaceWithFactors(A *utils.ListNode, B int) *utils.ListNode {

	var head *utils.ListNode = utils.ListNode_new(-1)
	var curr *utils.ListNode = head
	for A != nil {
		A.Value = (A.Value - (A.Value % B))
		curr.Next = A
		curr = curr.Next
		A = A.Next
	}

	return head.Next
}
