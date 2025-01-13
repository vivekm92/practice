package linkedList

import (
	"interviewBit/src/utils"
)

/*
	Problem : https://www.interviewbit.com/problems/reorder-list/
*/

// T(n) : O(n), S(n) : O(1)
func ReorderList(A *utils.ListNode) *utils.ListNode {

	var mid, r *utils.ListNode = A, A.Next
	for r != nil && r.Next != nil {
		mid = mid.Next
		r = r.Next.Next
	}

	var prev, curr *utils.ListNode = nil, mid.Next
	mid.Next = nil
	for curr != nil {
		var t *utils.ListNode = curr.Next
		curr.Next = prev
		prev = curr
		curr = t
	}

	var k int = 0
	var h1, h2 *utils.ListNode = A, prev
	var dummy *utils.ListNode = utils.ListNode_new(-1)
	curr = dummy
	for h1 != nil && h2 != nil {
		if k^1 == 1 {
			curr.Next = h1
			h1 = h1.Next
		} else {
			curr.Next = h2
			h2 = h2.Next
		}

		k = k ^ 1
		curr = curr.Next
	}

	if h1 == nil {
		curr.Next = h2
	} else {
		curr.Next = h1
	}

	return dummy.Next
}
