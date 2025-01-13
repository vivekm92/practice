package linkedList

import (
	"interviewBit/src/utils"
)

/*
	Problem : https://www.interviewbit.com/problems/remove-duplicates-from-sorted-list-ii/
*/

// T(n) : O(n), S(n) : O(1)
func DeleteDuplicatesII(A *utils.ListNode) *utils.ListNode {

	var prev, curr *utils.ListNode = nil, A
	for curr != nil {
		currNext := curr.Next
		if currNext != nil && curr.Value == currNext.Value {

			t := currNext
			for t != nil && curr.Value == t.Value {
				t = t.Next
			}

			if prev == nil {
				A = t
			} else {
				prev.Next = t
			}

			curr = t
		} else {
			prev = curr
			curr = curr.Next
		}
	}
	return A
}
