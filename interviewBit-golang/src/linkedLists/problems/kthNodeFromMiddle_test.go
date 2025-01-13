package linkedList

import (
	"fmt"
	"interviewBit/src/utils"
	"testing"
)

type kthNodeFromMiddleTestCase struct {
	A        *utils.ListNode
	B        int
	Expected int
}

func generateKthNodeFromMiddleTestCase() []kthNodeFromMiddleTestCase {
	var kthNodeFromMiddleTestCases []kthNodeFromMiddleTestCase
	ia1 := utils.GenerateLinkedList([]int{3, 4, 7, 5, 6, 16, 15, 61, 16})
	ib1 := 4
	t1 := kthNodeFromMiddleTestCase{
		A:        ia1,
		B:        ib1,
		Expected: 3,
	}
	kthNodeFromMiddleTestCases = append(kthNodeFromMiddleTestCases, t1)

	ia2 := utils.GenerateLinkedList([]int{3, 4, 7, 5, 6, 16, 15, 61})
	ib2 := 2
	t2 := kthNodeFromMiddleTestCase{
		A:        ia2,
		B:        ib2,
		Expected: 7,
	}
	kthNodeFromMiddleTestCases = append(kthNodeFromMiddleTestCases, t2)

	ia3 := utils.GenerateLinkedList([]int{3, 4, 7, 5, 6, 16, 15, 61})
	ib3 := 5
	t3 := kthNodeFromMiddleTestCase{
		A:        ia3,
		B:        ib3,
		Expected: -1,
	}
	kthNodeFromMiddleTestCases = append(kthNodeFromMiddleTestCases, t3)

	return kthNodeFromMiddleTestCases
}

func TestKthNodeFromMiddle(t *testing.T) {
	for idx, test := range generateKthNodeFromMiddleTestCase() {
		if output := KthNodeFromMiddle(test.A, test.B); output != test.Expected {
			fmt.Println(test.A.Value, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
