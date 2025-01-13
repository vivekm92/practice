package linkedList

import (
	"interviewBit/src/utils"
)

/*
	Problem : https://www.interviewbit.com/problems/kth-node-from-middle/
*/

// T(n) : O(n), S(n) : O(1)
func KthNodeFromMiddle(A *utils.ListNode, B int) int {

	var head *utils.ListNode = A
	var currNode *utils.ListNode = head
	n := 0
	for currNode != nil {
		n++
		currNode = currNode.Next
	}

	middle := n / 2
	target := middle - B
	if target < 0 {
		return -1
	}
	for target > 0 {
		head = head.Next
		target--
	}

	return head.Value
}
