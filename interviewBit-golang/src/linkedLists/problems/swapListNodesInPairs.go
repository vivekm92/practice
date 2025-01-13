package linkedList

import (
	"interviewBit/src/utils"
)

/*
	Problem : https://www.interviewbit.com/problems/swap-list-nodes-in-pairs/
*/

// T(n) : O(n), S(n) : O(1)
func SwapPairs(A *utils.ListNode) *utils.ListNode {

	var prev, curr *utils.ListNode = nil, A
	for curr != nil {
		currNext := curr.Next
		if currNext == nil {
			return A
		}

		if prev == nil {
			A = currNext
		} else {
			prev.Next = currNext
		}

		currNext = currNext.Next
		curr.Next.Next = curr
		curr.Next = currNext
		prev = curr
		curr = currNext
	}

	return A
}
