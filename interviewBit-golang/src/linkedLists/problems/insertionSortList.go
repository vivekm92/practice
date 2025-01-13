package linkedList

import (
	"interviewBit/src/utils"
)

/*
	Problem : https://www.interviewbit.com/problems/insertion-sort-list/
*/

// T(n) : O(n2), S(n) : O(1)
func InsertionSortList(A *utils.ListNode) *utils.ListNode {

	var start *utils.ListNode = utils.ListNode_new(-1)
	start.Next = A

	var prev, curr *utils.ListNode = start, start.Next
	for curr != nil {
		var pt, t *utils.ListNode = prev, curr
		var rpt, rt *utils.ListNode = prev, curr
		for t != nil {
			if t.Value < rt.Value {
				rpt = pt
				rt = t
			}

			pt = t
			t = t.Next
		}

		if rt == curr {
			prev = curr
			curr = curr.Next
			continue
		}

		rpt.Next = rt.Next
		prev.Next = rt
		rt.Next = curr
		prev = rt
	}

	return start.Next
}
