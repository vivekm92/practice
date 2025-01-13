package linkedList

import (
	"interviewBit/src/utils"
)

/*
	Problem : https://www.interviewbit.com/problems/k-reverse-linked-list/
*/

// T(n) : O(n), S(n) : O(1)
func ReverseList(A *utils.ListNode, B int) *utils.ListNode {

	k := 0
	var prev, curr *utils.ListNode = nil, A
	for curr != nil {

		pB, hB, pt, t := prev, curr, prev, curr
		for k < B && t != nil {
			nt := t.Next
			t.Next = pt
			pt = t
			t = nt
			k++
		}

		if k == B {
			k = 0
		}

		if pB == nil {
			A = pt
		} else {
			pB.Next = pt
		}
		hB.Next = t

		prev = hB
		curr = t
	}

	return A
}
