package linkedList

import (
	"fmt"
	"interviewBit/src/utils"
	"testing"
)

type rotateRightTestCase struct {
	A        *utils.ListNode
	B        int
	Expected *utils.ListNode
}

func getRotateRightTestCases() []rotateRightTestCase {
	var tc []rotateRightTestCase
	ia1 := utils.GenerateLinkedList([]int{1, 2, 3, 4, 5})
	o1 := utils.GenerateLinkedList([]int{4, 5, 1, 2, 3})
	var tc1 rotateRightTestCase
	tc1.A, tc1.B, tc1.Expected = ia1, 2, o1
	tc = append(tc, tc1)

	ia2 := utils.GenerateLinkedList([]int{1})
	o2 := utils.GenerateLinkedList([]int{1})
	var tc2 rotateRightTestCase
	tc2.A, tc2.B, tc2.Expected = ia2, 1, o2
	tc = append(tc, tc2)

	ia3 := utils.GenerateLinkedList([]int{1, 2, 3, 4, 5, 6})
	o3 := utils.GenerateLinkedList([]int{6, 1, 2, 3, 4, 5})
	var tc3 rotateRightTestCase
	tc3.A, tc3.B, tc3.Expected = ia3, 1, o3
	tc = append(tc, tc3)

	return tc
}

func TestRotateRight(t *testing.T) {
	for idx, test := range getRotateRightTestCases() {
		if output := RotateRight(test.A, test.B); !utils.CompareLinkedLists(output, test.Expected) {
			fmt.Println(test.A.Value, test.Expected.Value, output.Value)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output.Value, test.Expected.Value)
		}
	}
}
