package linkedList

import (
	"interviewBit/src/utils"
)

/*
	Problem : https://www.interviewbit.com/problems/remove-nth-node-from-list-end/
*/

// T(n) : O(n), S(n) : O(1)
func RemoveNthFromEnd(A *utils.ListNode, B int) *utils.ListNode {

	n := 0
	var curr *utils.ListNode = A
	for curr != nil {
		n++
		curr = curr.Next
	}

	k := n - B
	if k <= 0 {
		return A.Next
	}

	i := 1
	curr = A
	for i < k && curr != nil {
		i++
		curr = curr.Next
	}
	curr.Next = curr.Next.Next

	return A
}
