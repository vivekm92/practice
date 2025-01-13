package linkedList

import (
	"interviewBit/src/utils"
)

/*
	Problem : https://www.interviewbit.com/problems/even-reverse/
*/

// T(n) : O(n), S(n) : O(1)
func EvenReverse(A *utils.ListNode) *utils.ListNode {

	if A == nil || A.Next == nil {
		return A
	}

	k := 1
	var prev, curr *utils.ListNode = nil, A
	var h2 = A.Next
	for curr != nil {
		currNext := curr.Next
		if k%2 == 0 {
			prev.Next = curr.Next
			if currNext != nil {
				curr.Next = currNext.Next
			}
		}

		k++
		prev = curr
		curr = currNext
	}

	// Reverse second List
	prev, curr = nil, h2
	for curr != nil {
		currNext := curr.Next
		curr.Next = prev
		prev = curr
		curr = currNext
	}
	h2 = prev

	// Merge Both Lists
	var x int = 0
	dummyHead := utils.ListNode_new(-1)
	curr = dummyHead
	for A != nil && h2 != nil {
		if x == 0 {
			curr.Next = A
			A = A.Next
		} else {
			curr.Next = h2
			h2 = h2.Next
		}

		x = x ^ 1
		curr = curr.Next
	}
	if A != nil {
		curr.Next = A
	} else {
		curr.Next = h2
	}

	return dummyHead.Next
}
