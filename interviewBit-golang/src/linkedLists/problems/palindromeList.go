package linkedList

import (
	"interviewBit/src/utils"
)

/*
	Problem : https://www.interviewbit.com/problems/palindrome-list/
*/

// T(n) : O(n), S(n) : O(1)
func PalindromeList(A *utils.ListNode) int {

	var mid, r *utils.ListNode = A, A.Next
	for r != nil && r.Next != nil {
		mid = mid.Next
		r = r.Next.Next
	}

	if mid == A && r == nil {
		return 1
	}

	var prev, curr *utils.ListNode = nil, mid.Next
	mid.Next = nil
	for curr != nil {
		var t *utils.ListNode = curr.Next
		curr.Next = prev
		prev = curr
		curr = t
	}

	var h2 *utils.ListNode = prev
	for A != nil && h2 != nil {
		if A.Value != h2.Value {
			return 0
		}

		A = A.Next
		h2 = h2.Next
	}

	return 1
}
