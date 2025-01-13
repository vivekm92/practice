package linkedList

import (
	"interviewBit/src/utils"
)

/*
	Problem : https://www.interviewbit.com/problems/reverse-alternate-k-nodes/
*/

// T(n) : O(n), S(n) : O(1)
func ReverseAlternate(A *utils.ListNode, B int) *utils.ListNode {

	blockCount := 0
	var prev, curr *utils.ListNode = nil, A
	for curr != nil {

		k := B
		var pB, hB *utils.ListNode = prev, curr
		if blockCount%2 == 0 {
			for k > 0 && curr != nil {
				var t *utils.ListNode = curr.Next
				curr.Next = prev
				prev = curr
				curr = t
				k--
			}

			hB.Next = curr
			if blockCount == 0 {
				A = prev
			} else {
				pB.Next = prev
			}
		} else {
			for k > 0 && curr != nil {
				prev = curr
				curr = curr.Next
				k--
			}
		}

		blockCount++
	}

	return A
}
