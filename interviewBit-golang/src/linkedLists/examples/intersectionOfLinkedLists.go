package linkedListExamples

import (
	"interviewBit/src/utils"
)

/*
	Problem : https://www.interviewbit.com/problems/intersection-of-linked-lists/
*/

// T(n) : O(n), S(n) : O(1)
func GetIntersectionNode(A *utils.ListNode, B *utils.ListNode) *utils.ListNode {

	var hA, hB *utils.ListNode = A, B
	var cntA, cntB int = 0, 0
	for hA != nil {
		hA = hA.Next
		cntA++
	}

	for hB != nil {
		hB = hB.Next
		cntB++
	}

	for cntA > cntB {
		A = A.Next
		cntA--
	}

	for cntB > cntA {
		B = B.Next
		cntB--
	}

	for cntA > 0 && cntB > 0 {
		if A == B {
			return A
		}

		A = A.Next
		B = B.Next
	}

	return nil
}
